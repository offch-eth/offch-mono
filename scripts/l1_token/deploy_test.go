package l1token

import (
	"testing"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/offch-eth/offch/scripts"
	"github.com/stretchr/testify/require"
)

func TestDeployL1Token(t *testing.T) {
	r := require.New(t)

	client, err := ethclient.Dial("https://1rpc.io/holesky")
	r.NoError(err)

	auth := scripts.GetAuth(t, client)

	Deploy(t, client, auth)
}
