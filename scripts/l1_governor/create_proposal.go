package l1governor

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
	governor "github.com/offch-eth/offch/contracts/l1_governor/OffchGovernor"
	l1token "github.com/offch-eth/offch/contracts/l1_token/OffchToken"
	crossChainSync "github.com/offch-eth/offch/contracts/thirdparty/CrossChainSync/CrossChainSync"
	"github.com/offch-eth/offch/scripts"
	"github.com/stretchr/testify/require"
)

func CreateProposal(
	t *testing.T,
	l1client, l2client *ethclient.Client,
	l1auth *bind.TransactOpts,
	l1TokenAddress, l1GovernodAddress common.Address) (
	proposalId, snapshotBlockNumber, voteEnd *big.Int,
) {
	ctx := context.TODO()
	r := require.New(t)

	blockNumber, err := l1client.BlockNumber(ctx)
	r.NoError(err)

	l2CrossChainSync, err := crossChainSync.NewContracts(scripts.L2Taiko, l2client)
	r.NoError(err)

	g, err := governor.NewContracts(l1GovernodAddress, l1client)
	r.NoError(err)

	abi, err := l1token.ContractsMetaData.GetAbi()
	r.NoError(err)

	transferCallData, err := abi.Pack("transfer", common.HexToAddress("0x5C2e64714a5Ff72A011e3ddC82dA45B9630d11eE"), big.NewInt(10))
	r.NoError(err)

	tenTokens, _ := big.NewInt(0).SetString("10000000000000000000", 10)
	var (
		targets     = []common.Address{l1TokenAddress}
		values      = []*big.Int{tenTokens}
		calldatas   = [][]byte{transferCallData}
		description = "Proposal #1: Grant 10 tokens to team"
	)

	wg := &sync.WaitGroup{}
	wg.Add(1)
	go func() {
		for {
			fmt.Println("Waiting for proposal created...")
			iterator, err := g.FilterProposalCreated(&bind.FilterOpts{Start: blockNumber})
			r.NoError(err)

			if ok := iterator.Next(); !ok {
				time.Sleep(2 * time.Second)
				continue
			}

			if iterator.Event != nil {
				proposalId = iterator.Event.ProposalId
				fmt.Println("L1 Governor created proposal: ", proposalId.String())
				wg.Done()
				break
			}
		}
	}()

	wg.Add(1)
	var message governor.OffchGovernorMessage
	go func() {
		for {
			fmt.Println("Waiting for message sent...")
			iterator, err := g.FilterMessageSent(&bind.FilterOpts{Start: blockNumber})
			r.NoError(err)

			if ok := iterator.Next(); !ok {
				time.Sleep(2 * time.Second)
				continue
			}

			if iterator.Event != nil {
				messageSent := iterator.Event
				message = iterator.Event.Message
				voteEnd = big.NewInt(int64(iterator.Event.Raw.BlockNumber + message.Duration.Uint64()))
				fmt.Println("L1 Governor sent message: ", messageSent.String())
				wg.Done()
				break
			}
		}
	}()

	snippet, err := l2CrossChainSync.GetSyncedSnippet(nil, 0)
	r.NoError(err)

	snapshotBlockNumber = big.NewInt(int64(snippet.RemoteBlockId))
	fmt.Println("Snapshot block number:", snapshotBlockNumber)

	fmt.Println("Creating proposal...")
	tx, err := g.ProposeWithSnapshotBlockNumber(l1auth, targets, values, calldatas, description, snapshotBlockNumber)
	r.NoError(err)

	fmt.Println("Create proposal tx hash:", tx.Hash().Hex())
	_, err = bind.WaitMined(ctx, l1client, tx)
	r.NoError(err)

	wg.Wait()

	return message.ProposalId, snapshotBlockNumber, voteEnd
}
