package l1governor

import (
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/offch-eth/offch/scripts"
	"github.com/stretchr/testify/require"
)

var (
	l2BlockNumberSignalSent = big.NewInt(157797)
)

func TestResultFromSignal(t *testing.T) {
	r := require.New(t)

	l1client, err := ethclient.Dial("https://ethereum-holesky.publicnode.com")
	r.NoError(err)

	auth := scripts.GetAuth(t, l1client)

	l2rpcClient, err := rpc.Dial("https://rpc.katla.taiko.xyz")
	r.NoError(err)

	ResultFromSignal(
		t, l1client, l2rpcClient, auth,
		scripts.L1GovernorAddress, scripts.L2FactoryAddress,
		l2BlockNumberSignalSent,
	)
}
