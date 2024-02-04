package l2voting

import (
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/offch-eth/offch/scripts"
	"github.com/stretchr/testify/require"
)

func TestSendResultsToL1(t *testing.T) {
	proposalId, _ := big.NewInt(0).SetString("85967152061974112845768064401284274197164913767842673911954340022401905078635", 10)

	r := require.New(t)

	client, err := ethclient.Dial("https://rpc.katla.taiko.xyz")
	r.NoError(err)

	auth := scripts.GetAuthTaikoL2(t, client)

	SendResultsToL1(
		t, client, auth,
		scripts.L2FactoryAddress, proposalId,
	)
}
