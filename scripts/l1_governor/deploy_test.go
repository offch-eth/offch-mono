package l1governor

import (
	"context"
	"fmt"
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/ethclient"
	governor "github.com/offch-eth/offch/contracts/l1_governor/OffchGovernor"
	"github.com/offch-eth/offch/scripts"
	"github.com/stretchr/testify/require"
)

func TestDeployGovernor(t *testing.T) {
	ctx := context.Background()
	r := require.New(t)

	client, err := ethclient.Dial("https://1rpc.io/holesky")
	r.NoError(err)

	auth := scripts.GetAuth(t, client)

	fmt.Println("Deploying governor...")
	addr, tx, governor, err := governor.DeployContracts(auth, client, scripts.L1TokenAddress, scripts.L1TaikoSignalService)
	r.NoError(err)

	fmt.Println("Deploy tx hash:", tx.Hash().Hex())

	_, err = bind.WaitMined(ctx, client, tx)
	r.NoError(err)

	fmt.Println("Governor deployed at:", addr.Hex(), "with tx hash:", tx.Hash().Hex())

	name, err := governor.Name(nil)
	r.NoError(err)
	fmt.Println("Governor name:", name)
}
