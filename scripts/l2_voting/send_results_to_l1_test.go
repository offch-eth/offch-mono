package l2voting

import (
	"context"
	"fmt"
	"math/big"
	"sync"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/ethclient"
	l2voting "github.com/offch-eth/offch/contracts/l2_voting/ProposalVoteFactory"
	"github.com/offch-eth/offch/scripts"
	"github.com/stretchr/testify/require"
)

func TestSendResultsToL1(t *testing.T) {
	proposalKey, _ := big.NewInt(0).SetString("85967152061974112845768064401284274197164913767842673911954340022401905078635", 10)

	ctx := context.Background()
	r := require.New(t)

	wg := &sync.WaitGroup{}

	client, err := ethclient.Dial("https://rpc.katla.taiko.xyz")
	r.NoError(err)
	l2blockNumber, err := client.BlockNumber(ctx)
	r.NoError(err)

	auth := scripts.GetAuthTaikoL2(t, client)

	f, err := l2voting.NewContracts(scripts.L2FactoryAddress, client)
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
				break
			}
		}

		wg.Done()
	}()

	tx, err := f.SendSignalAboutResults(auth, proposalKey)
	r.NoError(err)

	fmt.Println("Send results to l1 tx hash:", tx.Hash().Hex())
	_, err = bind.WaitMined(ctx, client, tx)
	r.NoError(err)

	wg.Wait()
}
