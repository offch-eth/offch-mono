package l1governor

import (
	"context"
	"fmt"
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	governor "github.com/offch-eth/offch/contracts/l1_governor/OffchGovernor"
	"github.com/offch-eth/offch/scripts"
	"github.com/stretchr/testify/require"
)

func DeployGovernor(t *testing.T, client *ethclient.Client, auth *bind.TransactOpts, tokenAddress common.Address) common.Address {
	r := require.New(t)

	fmt.Println("Deploying governor...")
	addr, tx, governor, err := governor.DeployContracts(auth, client, tokenAddress, scripts.L1TaikoSignalService)
	r.NoError(err)

	fmt.Println("Deploy tx hash:", tx.Hash().Hex())

	_, err = bind.WaitMined(context.TODO(), client, tx)
	r.NoError(err)

	fmt.Println("Governor deployed at:", addr.Hex(), "with tx hash:", tx.Hash().Hex())

	name, err := governor.Name(nil)
	r.NoError(err)
	fmt.Println("Governor name:", name)

	return addr
}
