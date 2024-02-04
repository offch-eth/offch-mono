package l2voting

import (
	"context"
	"testing"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/offch-eth/offch/scripts"
	"github.com/stretchr/testify/require"
)

func TestCreateVoteFromSignal(t *testing.T) {
	ctx := context.Background()
	r := require.New(t)

	l1RpcClient, err := rpc.Dial("https://ethereum-holesky.publicnode.com")
	r.NoError(err)
	l1client := ethclient.NewClient(l1RpcClient)
	r.NoError(err)
	l1blockNumber, err := l1client.BlockNumber(ctx)
	r.NoError(err)
	// TODO: define the block number to get events not from the latest block
	// l1blockNumber = 869396 - 10

	client, err := ethclient.Dial("https://rpc.katla.taiko.xyz")
	r.NoError(err)

	auth := scripts.GetAuthTaikoL2(t, client)

	CreateVoteFromSignal(t, l1blockNumber, l1RpcClient, client, auth,
		scripts.L1TokenAddress, scripts.L1GovernorAddress, scripts.L2FactoryAddress,
	)
}
