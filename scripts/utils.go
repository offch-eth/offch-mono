package scripts

import (
	"context"
	"crypto/ecdsa"
	"math/big"
	"strings"
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rlp"
	signalservice "github.com/offch-eth/offch/contracts/thirdparty/SignalService/SignalService"
	"github.com/stretchr/testify/require"
)

func GetAuth(t *testing.T, client *ethclient.Client) *bind.TransactOpts {
	r := require.New(t)

	privateKey, err := crypto.HexToECDSA(Cfg.AccountPrivateKey)
	r.NoError(err)

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	r.True(ok)

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	r.NoError(err)

	gasPrice, err := client.SuggestGasPrice(context.Background())
	r.NoError(err)

	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(int64(L1ChainID)))
	r.NoError(err)

	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0) // in wei
	auth.GasLimit = uint64(0)  // in units
	auth.GasPrice = gasPrice.Mul(gasPrice, big.NewInt(2))

	return auth
}

func GetAuthTaikoL2(t *testing.T, client *ethclient.Client) *bind.TransactOpts {
	r := require.New(t)

	privateKey, err := crypto.HexToECDSA(Cfg.AccountPrivateKey)
	r.NoError(err)

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	r.True(ok)

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	r.NoError(err)

	gasPrice, err := client.SuggestGasPrice(context.Background())
	r.NoError(err)

	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(int64(L2ChainID)))
	r.NoError(err)

	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0) // in wei
	auth.GasLimit = uint64(0)  // in units
	auth.GasPrice = gasPrice.Mul(gasPrice, big.NewInt(2))

	return auth
}

var hopComponents = []abi.ArgumentMarshaling{
	{
		Name: "signalRootRelay",
		Type: "address",
	},
	{
		Name: "signalRoot",
		Type: "bytes32",
	},
	{
		Name: "storageProof",
		Type: "bytes",
	},
}

var signalProofT, _ = abi.NewType("tuple", "", []abi.ArgumentMarshaling{
	{
		Name: "crossChainSync",
		Type: "address",
	},
	{
		Name: "height",
		Type: "uint64",
	},
	{
		Name: "storageProof",
		Type: "bytes",
	},
	{
		Name:       "hops",
		Type:       "tuple[]",
		Components: hopComponents,
	},
})

func EncodeSignalProof(crossChainSync common.Address, height uint64, storageProof []string) ([]byte, error) {
	var bytesProof [][]byte
	for _, v := range storageProof {
		bytesProof = append(bytesProof, common.Hex2Bytes(strings.TrimPrefix(v, "0x")))
	}

	encodedProof, err := rlp.EncodeToBytes(bytesProof)
	if err != nil {
		return nil, err
	}

	signalProof := signalservice.ISignalServiceProof{
		CrossChainSync: crossChainSync,
		Height:         height,
		StorageProof:   encodedProof,
	}
	args := abi.Arguments{
		{
			Type: signalProofT,
		},
	}

	return args.Pack(signalProof)
}
