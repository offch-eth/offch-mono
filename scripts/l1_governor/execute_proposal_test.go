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
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	governor "github.com/offch-eth/offch/contracts/l1_governor/OffchGovernor"
	l1token "github.com/offch-eth/offch/contracts/l1_token/OffchToken"
	"github.com/offch-eth/offch/scripts"
	"github.com/stretchr/testify/require"
)

func TestExecuteProposal(t *testing.T) {
	ctx := context.Background()
	wg := &sync.WaitGroup{}
	r := require.New(t)

	l1client, err := ethclient.Dial("https://ethereum-holesky.publicnode.com")
	r.NoError(err)
	l1blockNumber, err := l1client.BlockNumber(ctx)
	r.NoError(err)

	auth := scripts.GetAuth(t, l1client)

	g, err := governor.NewContracts(scripts.L1GovernorAddress, l1client)
	r.NoError(err)

	abi, err := l1token.ContractsMetaData.GetAbi()
	r.NoError(err)

	transferCallData, err := abi.Pack("transfer", common.HexToAddress("0x5C2e64714a5Ff72A011e3ddC82dA45B9630d11eE"), big.NewInt(10))
	r.NoError(err)

	var (
		targets     = []common.Address{scripts.L1TokenAddress}
		values      = []*big.Int{big.NewInt(0)}
		calldatas   = [][]byte{transferCallData}
		description = "Proposal #1: Grant 10 tokens to team"
	)

	descriptionHash := crypto.Keccak256Hash([]byte(description))

	wg.Add(1)
	go func() {
		for {
			fmt.Println("Waiting for proposal to be executed event from", l1blockNumber, "...")
			iterator, err := g.FilterProposalExecuted(
				&bind.FilterOpts{
					Start: l1blockNumber,
				},
			)
			r.NoError(err)

			if ok := iterator.Next(); !ok {
				time.Sleep(2 * time.Second)
				continue
			}

			if iterator.Event != nil {
				fmt.Println("Proposal executed: ", iterator.Event.String())
				break
			}
		}

		wg.Done()
	}()

	tx, err := g.Execute(auth, targets, values, calldatas, descriptionHash)
	r.NoError(err)

	fmt.Println("Execute proposal tx hash:", tx.Hash().Hex())
	_, err = bind.WaitMined(ctx, l1client, tx)
	r.NoError(err)

	wg.Wait()
}
