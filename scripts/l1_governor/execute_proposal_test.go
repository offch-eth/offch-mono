package l1governor

import (
	"testing"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/offch-eth/offch/scripts"
	"github.com/stretchr/testify/require"
)

func TestExecuteProposal(t *testing.T) {
	r := require.New(t)

	l1client, err := ethclient.Dial("https://ethereum-holesky.publicnode.com")
	r.NoError(err)

	auth := scripts.GetAuth(t, l1client)

	ExecuteProposal(
		t, l1client, auth,
		scripts.L1TokenAddress, scripts.L1GovernorAddress,
	)
}
