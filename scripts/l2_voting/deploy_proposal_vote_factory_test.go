package l2voting

import (
	"testing"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/offch-eth/offch/scripts"
	"github.com/stretchr/testify/require"
)

func TestDeployProposalVoteFactory(t *testing.T) {
	r := require.New(t)

	client, err := ethclient.Dial("https://rpc.katla.taiko.xyz")
	r.NoError(err)

	auth := scripts.GetAuthTaikoL2(t, client)

	DeployVoteFactory(t, client, auth)
}
