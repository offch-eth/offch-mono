package l1governor

import (
	"testing"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/offch-eth/offch/scripts"
	"github.com/stretchr/testify/require"
)

func TestDeployGovernor(t *testing.T) {
	r := require.New(t)

	client, err := ethclient.Dial("https://1rpc.io/holesky")
	r.NoError(err)

	auth := scripts.GetAuth(t, client)

	DeployGovernor(t, client, auth, scripts.L1TokenAddress)
}
