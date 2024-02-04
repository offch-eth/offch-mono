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
	"github.com/ethereum/go-ethereum/rpc"
	governor "github.com/offch-eth/offch/contracts/l1_governor/OffchGovernor"
	l1token "github.com/offch-eth/offch/contracts/l1_token/OffchToken"
	crossChainSync "github.com/offch-eth/offch/contracts/thirdparty/CrossChainSync/CrossChainSync"
	"github.com/offch-eth/offch/scripts"
	"github.com/stretchr/testify/require"
)

func TestCreateProposal(t *testing.T) {
	ctx := context.Background()
	r := require.New(t)

	client, err := ethclient.Dial("https://ethereum-holesky.publicnode.com")
	r.NoError(err)
	blockNumber, err := client.BlockNumber(ctx)
	r.NoError(err)

	l2rpcClient, err := rpc.Dial("https://rpc.katla.taiko.xyz")
	r.NoError(err)
	l2client := ethclient.NewClient(l2rpcClient)
	l2CrossChainSync, err := crossChainSync.NewContracts(scripts.L2Taiko, l2client)
	r.NoError(err)

	auth := scripts.GetAuth(t, client)

	g, err := governor.NewContracts(scripts.L1GovernorAddress, client)
	r.NoError(err)

	abi, err := l1token.ContractsMetaData.GetAbi()
	r.NoError(err)

	transferCallData, err := abi.Pack("transfer", common.HexToAddress("0x5C2e64714a5Ff72A011e3ddC82dA45B9630d11eE"), big.NewInt(10))
	r.NoError(err)

	tenTokens, _ := big.NewInt(0).SetString("10000000000000000000", 10)
	var (
		targets     = []common.Address{scripts.L1TokenAddress}
		values      = []*big.Int{tenTokens}
		calldatas   = [][]byte{transferCallData}
		description = "Proposal #1: Grant 10 tokens to team"
	)

	wg := &sync.WaitGroup{}
	wg.Add(1)
	proposalCreatedChan := make(chan *governor.ContractsProposalCreated)
	psub, err := g.WatchProposalCreated(&bind.WatchOpts{Start: &blockNumber}, proposalCreatedChan)
	r.NoError(err)

	defer psub.Unsubscribe()

	go func() {
		for proposalCreated := range proposalCreatedChan {
			fmt.Println("Proposal Created Logs: ", proposalCreated.String())
			wg.Done()
			return
		}
	}()

	wg.Add(1)
	messageSentChan := make(chan *governor.ContractsMessageSent)
	msub, err := g.WatchMessageSent(&bind.WatchOpts{Start: &blockNumber}, messageSentChan)
	r.NoError(err)

	defer msub.Unsubscribe()

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
				fmt.Println("L1 Governor sent message: ", messageSent.String())
				wg.Done()
				break
			}
		}
	}()

	snippet, err := l2CrossChainSync.GetSyncedSnippet(nil, 0)
	r.NoError(err)

	snapshotBlockNumber := big.NewInt(int64(snippet.RemoteBlockId))
	fmt.Println("Snapshot block number:", snapshotBlockNumber)

	fmt.Println("Creating proposal...")
	tx, err := g.ProposeWithSnapshotBlockNumber(auth, targets, values, calldatas, description, snapshotBlockNumber)
	r.NoError(err)

	fmt.Println("Create proposal tx hash:", tx.Hash().Hex())
	_, err = bind.WaitMined(ctx, client, tx)
	r.NoError(err)

	wg.Wait()
}
