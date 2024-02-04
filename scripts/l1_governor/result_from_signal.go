package l1governor

import (
	"context"
	"fmt"
	"math/big"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/ethclient/gethclient"
	"github.com/ethereum/go-ethereum/rpc"
	governor "github.com/offch-eth/offch/contracts/l1_governor/OffchGovernor"
	l2voting "github.com/offch-eth/offch/contracts/l2_voting/ProposalVoteFactory"
	signalService "github.com/offch-eth/offch/contracts/thirdparty/SignalService/SignalService"
	"github.com/offch-eth/offch/scripts"
	"github.com/stretchr/testify/require"
)

func ResultFromSignal(
	t *testing.T,
	l1client *ethclient.Client,
	l2rpcClient *rpc.Client,
	l1auth *bind.TransactOpts,
	l1governorAddress, l2factoryAddress common.Address,
	l2BlockNumberSignalSent *big.Int,
) {
	ctx := context.Background()
	r := require.New(t)

	l1blockNumber, err := l1client.BlockNumber(ctx)
	r.NoError(err)

	l1SignalService, err := signalService.NewContracts(scripts.L1TaikoSignalService, l1client)
	r.NoError(err)

	g, err := governor.NewContracts(l1governorAddress, l1client)
	r.NoError(err)

	l2client := ethclient.NewClient(l2rpcClient)
	l2gethClient := gethclient.New(l2rpcClient)

	f, err := l2voting.NewContracts(l2factoryAddress, l2client)
	r.NoError(err)

	var l2MessageHash common.Hash
	var l2Message governor.GovernorCountingFromSignalProposalVoteMessage
	for {
		fmt.Println("Looking for message sent event from", l2BlockNumberSignalSent, "...")
		iterator, err := f.FilterMessageSent(
			&bind.FilterOpts{
				Start: l2BlockNumberSignalSent.Uint64(),
			},
		)
		r.NoError(err)

		if ok := iterator.Next(); !ok {
			time.Sleep(2 * time.Second)
			continue
		}

		if iterator.Event != nil {
			fmt.Println("Message sent found: ", iterator.Event.String())
			fmt.Println("Message:", iterator.Event.Message.String())
			l2MessageHash = iterator.Event.MessageHash
			l2Message = (governor.GovernorCountingFromSignalProposalVoteMessage)(iterator.Event.Message)
			break
		}
	}

	slot, err := l1SignalService.GetSignalSlot(
		nil, uint64(scripts.L2ChainID), l2factoryAddress, l2MessageHash,
	)
	r.NoError(err)

	proof, err := l2gethClient.GetProof(
		ctx,
		scripts.L2TaikoSignalService,
		[]string{common.Bytes2Hex(slot[:])},
		l2BlockNumberSignalSent,
	)
	r.NoError(err)

	encodedSignalProof, err := scripts.EncodeSignalProof(scripts.L1Taiko, l1blockNumber, proof.StorageProof[0].Proof)
	r.NoError(err)

	tx, err := g.ResultFromSignal(l1auth, l2factoryAddress, l2MessageHash, l2Message, encodedSignalProof)
	r.NoError(err)

	fmt.Println("Save result from signal tx hash:", tx.Hash().Hex())
	_, err = bind.WaitMined(ctx, l1client, tx)
	r.NoError(err)
}
