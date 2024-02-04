package l1token

import (
	"context"
	"fmt"
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	l1token "github.com/offch-eth/offch/contracts/l1_token/OffchToken"
	"github.com/stretchr/testify/require"
)

func Deploy(t *testing.T, client *ethclient.Client, auth *bind.TransactOpts) common.Address {
	r := require.New(t)

	fmt.Println("Deploying L1Token...")
	addr, tx, l1Token, err := l1token.DeployContracts(auth, client)
	r.NoError(err)

	fmt.Println("Deploy tx hash:", tx.Hash().Hex())

	_, err = bind.WaitMined(context.TODO(), client, tx)
	r.NoError(err)

	fmt.Println("L1Token deployed at:", addr.Hex(), "with tx hash:", tx.Hash().Hex())

	name, err := l1Token.Name(nil)
	r.NoError(err)
	fmt.Println("L1Token name:", name)

	return addr
}
