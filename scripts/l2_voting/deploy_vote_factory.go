package l2voting

import (
	"context"
	"fmt"
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	l2voting "github.com/offch-eth/offch/contracts/l2_voting/ProposalVoteFactory"
	"github.com/offch-eth/offch/scripts"
	"github.com/stretchr/testify/require"
)

func DeployVoteFactory(t *testing.T, client *ethclient.Client, auth *bind.TransactOpts) common.Address {
	r := require.New(t)
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

	_, err = bind.WaitMined(context.TODO(), client, tx)
	r.NoError(err)

	fmt.Println("L2 Factory deployed at:", addr.Hex(), "with tx hash:", tx.Hash().Hex())

	return addr
}
