package l2voting

import (
	"context"
	"fmt"
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/ethclient"
	l2voting "github.com/offch-eth/offch/contracts/l2_voting/ProposalVoteFactory"
	"github.com/offch-eth/offch/scripts"
	"github.com/stretchr/testify/require"
)

func TestDeployProposalVoteFactory(t *testing.T) {
	ctx := context.Background()
	r := require.New(t)

	client, err := ethclient.Dial("https://rpc.katla.taiko.xyz")
	r.NoError(err)

	auth := scripts.GetAuthTaikoL2(t, client)

	fmt.Println("Deploying L2 factory...")
	addr, tx, _, err := l2voting.DeployContracts(
		auth,
		client,
		scripts.L1ChainID,
		scripts.L2TaikoSignalService,
		scripts.L2Taiko,
	)
	r.NoError(err)

	fmt.Println("Deploy tx hash:", tx.Hash().Hex())

	_, err = bind.WaitMined(ctx, client, tx)
	r.NoError(err)

	fmt.Println("L2 Factory deployed at:", addr.Hex(), "with tx hash:", tx.Hash().Hex())
}
