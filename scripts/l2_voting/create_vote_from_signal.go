package l2voting

import (
	"context"
	"fmt"
	"math/big"
	"sync"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/ethclient/gethclient"
	"github.com/ethereum/go-ethereum/rpc"
	governor "github.com/offch-eth/offch/contracts/l1_governor/OffchGovernor"
	l2voting "github.com/offch-eth/offch/contracts/l2_voting/ProposalVoteFactory"
	crossChainSync "github.com/offch-eth/offch/contracts/thirdparty/CrossChainSync/CrossChainSync"
	signalService "github.com/offch-eth/offch/contracts/thirdparty/SignalService/SignalService"
	"github.com/offch-eth/offch/scripts"
	"github.com/stretchr/testify/require"
)

func CreateVoteFromSignal(
	t *testing.T,
	l1blockNumber uint64,
	l1rpcClient *rpc.Client,
	l2client *ethclient.Client,
	l2auth *bind.TransactOpts,
	l1tokenAddress, l1governorAddress, l2factoryAddress common.Address) (proposalKey *big.Int) {
	r := require.New(t)
	ctx := context.TODO()

	l1client := ethclient.NewClient(l1rpcClient)
	l1gethclient := gethclient.New(l1rpcClient)

	l2blockNumber, err := l2client.BlockNumber(ctx)
	r.NoError(err)

	g, err := governor.NewContracts(l1governorAddress, l1client)
	r.NoError(err)
	l1SignalService, err := signalService.NewContracts(scripts.L1TaikoSignalService, l1client)
	r.NoError(err)

	wg := &sync.WaitGroup{}
	wg.Add(1)

	var messageHash common.Hash
	var message *l2voting.OffchGovernorMessage
	var proof *gethclient.AccountResult
	var blockNumber uint64

	go func() {
		for {
			fmt.Println("Waiting for message sent event...")
			iterator, err := g.FilterMessageSent(&bind.FilterOpts{Start: l1blockNumber})
			r.NoError(err)

			if ok := iterator.Next(); !ok {
				time.Sleep(2 * time.Second)
				continue
			}

			if iterator.Event != nil {
				messageSent := iterator.Event
				fmt.Println("L1 Governor sent message: ", messageSent.String())
				messageHash = messageSent.MessageHash
				message = (*l2voting.OffchGovernorMessage)(&messageSent.Message)
				blockNumber = messageSent.Raw.BlockNumber
				wg.Done()
				break
			}
		}
	}()
	wg.Wait()

	l2CrossChainSync, err := crossChainSync.NewContracts(scripts.L2Taiko, l2client)
	r.NoError(err)

	time.Sleep(5 * time.Second)

	for {
		height, err := l2CrossChainSync.LatestSyncedL1Height(nil)
		r.NoError(err)

		if height < blockNumber {
			fmt.Println("Waiting for L1 block to be synced...")
			time.Sleep(2 * time.Second)
			continue
		}

		snippet, err := l2CrossChainSync.GetSyncedSnippet(nil, height)
		r.NoError(err)
		if snippet.BlockHash != (common.Hash{}) {
			blockNumber = height
			fmt.Println("L1 block synced!",
				"Number: ", blockNumber,
				"Hash: ", common.Hash(snippet.BlockHash).Hex(),
				"Signal root: ", common.Hash(snippet.SignalRoot).Hex())
			break
		}

		fmt.Println("Hash not found, waiting for L1 block to be synced...")
		time.Sleep(1 * time.Second)
	}
	f, err := l2voting.NewContracts(l2factoryAddress, l2client)
	r.NoError(err)

	wg.Add(1)
	go func() {
		for {
			fmt.Println("Waiting for proposal created event from", l2blockNumber, "...")
			iterator, err := f.FilterProposalCreated(
				&bind.FilterOpts{
					Start: l2blockNumber,
				},
				[]common.Address{l1governorAddress},
				[]common.Address{l1tokenAddress},
			)
			r.NoError(err)

			if ok := iterator.Next(); !ok {
				time.Sleep(2 * time.Second)
				continue
			}

			if iterator.Event != nil {
				fmt.Println("Proposal Created Logs: ", iterator.Event.String())
				proposalKey = iterator.Event.VoteKey
				break
			}
		}

		wg.Done()
	}()

	slot, err := l1SignalService.GetSignalSlot(
		&bind.CallOpts{BlockNumber: big.NewInt(int64(blockNumber))}, scripts.L1ChainID, l1governorAddress, messageHash,
	)
	r.NoError(err)

	sent, err := l1SignalService.IsSignalSent(&bind.CallOpts{BlockNumber: big.NewInt(int64(blockNumber))}, l1governorAddress, messageHash)
	r.NoError(err)
	fmt.Println("Signal sent:", sent)

	proof, err = l1gethclient.GetProof(
		ctx,
		scripts.L1TaikoSignalService,
		[]string{common.Bytes2Hex(slot[:])},
		big.NewInt(int64(blockNumber)),
	)
	r.NoError(err)
	fmt.Println("proof generated", "value", proof.StorageProof[0].Value.Int64())

	encodedSignalProof, err := scripts.EncodeSignalProof(scripts.L2Taiko, blockNumber, proof.StorageProof[0].Proof)
	r.NoError(err)

	tx, err := f.CreateProposalVoteFromSignal(
		l2auth,
		l1governorAddress,
		messageHash,
		*message,
		encodedSignalProof,
	)
	r.NoError(err)

	fmt.Println("Create proposal vote from signal tx hash:", tx.Hash().Hex())
	_, err = bind.WaitMined(ctx, l2client, tx)
	r.NoError(err)

	wg.Wait()

	return
}
