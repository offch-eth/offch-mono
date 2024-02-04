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
	l2voting "github.com/offch-eth/offch/contracts/l2_voting/ProposalVoteFactory"
	"github.com/stretchr/testify/require"
)

func SendResultsToL1(
	t *testing.T,
	l2client *ethclient.Client,
	l2auth *bind.TransactOpts,
	l2factoryAddress common.Address,
	proposalID *big.Int,
) (signalSentBlockNumber *big.Int) {
	ctx := context.Background()
	r := require.New(t)
	wg := &sync.WaitGroup{}

	l2blockNumber, err := l2client.BlockNumber(ctx)
	r.NoError(err)

	f, err := l2voting.NewContracts(l2factoryAddress, l2client)
	r.NoError(err)

	wg.Add(1)
	go func() {
		for {
			fmt.Println("Waiting for message sent event from", l2blockNumber, "...")
			iterator, err := f.FilterMessageSent(
				&bind.FilterOpts{
					Start: l2blockNumber,
				},
			)
			r.NoError(err)

			if ok := iterator.Next(); !ok {
				time.Sleep(2 * time.Second)
				continue
			}

			if iterator.Event != nil {
				fmt.Println("Message sent casted: ", iterator.Event.String())
				fmt.Println("Message: ", iterator.Event.Message.String())
				fmt.Println("Event block number", iterator.Event.Raw.BlockNumber)
				signalSentBlockNumber = big.NewInt(int64(iterator.Event.Raw.BlockNumber))
				break
			}
		}

		wg.Done()
	}()

	tx, err := f.SendSignalAboutResults(l2auth, proposalID)
	r.NoError(err)

	fmt.Println("Send results to l1 tx hash:", tx.Hash().Hex())
	_, err = bind.WaitMined(ctx, l2client, tx)
	r.NoError(err)

	wg.Wait()
	return
}
