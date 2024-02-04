package l2voting

import (
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/offch-eth/offch/scripts"
	"github.com/stretchr/testify/require"
)

var (
	snapshotBlockNumber = big.NewInt(877016)
	proposalKey         = big.NewInt(0)
)

func init() {
	proposalKey, _ = proposalKey.SetString("85967152061974112845768064401284274197164913767842673911954340022401905078635", 10)
}

func TestVote(t *testing.T) {
	r := require.New(t)

	l1RpcClient, err := rpc.Dial("https://ethereum-holesky.publicnode.com")
	r.NoError(err)

	client, err := ethclient.Dial("https://rpc.katla.taiko.xyz")
	r.NoError(err)

	auth := scripts.GetAuthTaikoL2(t, client)

	Vote(
		t, l1RpcClient, client, auth,
		scripts.L1TokenAddress, scripts.L1GovernorAddress, scripts.L2FactoryAddress,
		proposalKey, snapshotBlockNumber,
	)
}
