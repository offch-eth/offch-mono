package l1governor

import (
	"testing"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/offch-eth/offch/scripts"
	"github.com/stretchr/testify/require"
)

func TestCreateProposal(t *testing.T) {
	r := require.New(t)

	client, err := ethclient.Dial("https://ethereum-holesky.publicnode.com")
	r.NoError(err)

	l2rpcClient, err := rpc.Dial("https://rpc.katla.taiko.xyz")
	r.NoError(err)
	l2client := ethclient.NewClient(l2rpcClient)

	auth := scripts.GetAuth(t, client)

	CreateProposal(t, client, l2client, auth,
		scripts.L1TokenAddress, scripts.L1GovernorAddress)
}
