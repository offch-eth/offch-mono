package integration_test

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
	crossChainSync "github.com/offch-eth/offch/contracts/thirdparty/CrossChainSync/CrossChainSync"
	"github.com/offch-eth/offch/scripts"
	l1governor "github.com/offch-eth/offch/scripts/l1_governor"
	l1token "github.com/offch-eth/offch/scripts/l1_token"
	l2voting "github.com/offch-eth/offch/scripts/l2_voting"
	"github.com/stretchr/testify/require"
)

func TestOffch(t *testing.T) {
	r := require.New(t)
	l1rpcClient, err := rpc.Dial("https://ethereum-holesky.publicnode.com")
	l1client := ethclient.NewClient(l1rpcClient)
	r.NoError(err)

	l2rpcClient, err := rpc.Dial("https://rpc.katla.taiko.xyz")
	r.NoError(err)
	l2client := ethclient.NewClient(l2rpcClient)

	l1TokenAddress := l1token.Deploy(t, l1client, scripts.GetAuth(t, l1client))
	fmt.Printf("\n\n\n")
	l1governorAddress := l1governor.DeployGovernor(t, l1client, scripts.GetAuth(t, l1client), l1TokenAddress)
	fmt.Printf("\n\n\n")
	l2voteFactoryAddress := l2voting.DeployVoteFactory(t, l2client, scripts.GetAuthTaikoL2(t, l2client))
	fmt.Printf("\n\n\n")

	_, snapshotBlockNumber, voteEnd := l1governor.CreateProposal(t, l1client, l2client, scripts.GetAuth(t, l1client), l1TokenAddress, l1governorAddress)
	fmt.Printf("\n\n\n")

	l2proposalKey := l2voting.CreateVoteFromSignal(
		t, snapshotBlockNumber.Uint64(),
		l1rpcClient, l2client, scripts.GetAuthTaikoL2(t, l2client),
		l1TokenAddress, l1governorAddress, l2voteFactoryAddress,
	)
	fmt.Printf("\n\n\n")

	l2voting.Vote(
		t, l1rpcClient, l2client, scripts.GetAuthTaikoL2(t, l2client),
		l1TokenAddress, l1governorAddress, l2voteFactoryAddress,
		l2proposalKey, snapshotBlockNumber,
	)
	fmt.Printf("\n\n\n")

	for {
		fmt.Println("Waiting for vote end...")
		blockNumber, err := l1client.BlockNumber(context.TODO())
		r.NoError(err)

		if blockNumber > voteEnd.Uint64() {
			break
		}

		time.Sleep(5 * time.Second)
	}

	signalSentBlockNumber := l2voting.SendResultsToL1(
		t, l2client, scripts.GetAuthTaikoL2(t, l2client),
		l2voteFactoryAddress, l2proposalKey,
	)
	fmt.Printf("\n\n\n")

	l1CrossChainSync, err := crossChainSync.NewContracts(scripts.L1Taiko, l1client)
	r.NoError(err)

	for {
		fmt.Println("Waiting for synced snippet...")
		snippet, err := l1CrossChainSync.GetSyncedSnippet(
			nil,
			signalSentBlockNumber.Uint64(),
		)
		if err != nil {
			time.Sleep(5 * time.Second)
			continue
		}

		if snippet.RemoteBlockId == signalSentBlockNumber.Uint64() {
			break
		}
	}

	l1governor.ResultFromSignal(
		t, l1client, l2rpcClient, scripts.GetAuth(t, l1client),
		l1governorAddress, l2voteFactoryAddress,
		signalSentBlockNumber,
	)
	fmt.Printf("\n\n\n")

	l1governor.ExecuteProposal(
		t, l1client, scripts.GetAuth(t, l1client),
		l1TokenAddress, l1governorAddress,
	)

	fmt.Printf("\n\n\n")

	fmt.Println("Integration test passed")
}
