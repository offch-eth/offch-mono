package l1token

import (
	"context"
	"fmt"
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/ethclient"
	l1token "github.com/offch-eth/offch/contracts/l1_token/OffchToken"
	"github.com/offch-eth/offch/scripts"
	"github.com/stretchr/testify/require"
)

func TestDeployL1Token(t *testing.T) {
	ctx := context.Background()
	r := require.New(t)

	client, err := ethclient.Dial("https://1rpc.io/holesky")
	r.NoError(err)

	auth := scripts.GetAuth(t, client)

	fmt.Println("Deploying L1Token...")
	addr, tx, l1Token, err := l1token.DeployContracts(auth, client)
	r.NoError(err)

	fmt.Println("Deploy tx hash:", tx.Hash().Hex())

	_, err = bind.WaitMined(ctx, client, tx)
	r.NoError(err)

	fmt.Println("L1Token deployed at:", addr.Hex(), "with tx hash:", tx.Hash().Hex())

	name, err := l1Token.Name(nil)
	r.NoError(err)
	fmt.Println("L1Token name:", name)
}
