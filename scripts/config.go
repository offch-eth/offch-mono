package scripts

import (
	"embed"
	"os"
	"strings"

	"github.com/caarlos0/env/v10"
	"github.com/ethereum/go-ethereum/common"
	"github.com/joho/godotenv"
)

var (
	// Taiko L1
	L1Taiko              = common.HexToAddress("0xB20BB9105e007Bd3E0F73d63D4D3dA2c8f736b77")
	L1TaikoSignalService = common.HexToAddress("0x08a3f537c4bbe8B6176420f4Cd0C84b02172dC65")

	// Taiko L2
	L2TaikoSignalService = common.HexToAddress("0x1670080000000000000000000000000000000005")
	L2Taiko              = common.HexToAddress("0x1670080000000000000000000000000000010001")

	// Offch L1
	L1GovernorAddress = common.HexToAddress("0x669A036455Cd7dc2a49f7A67ABFd113ac249A765")
	L1TokenAddress    = common.HexToAddress("0x90f7d2A09D0c66327c39bd9513f348ee885E89e6")

	// Offch L2
	L2FactoryAddress = common.HexToAddress("0xB7AC63071F7377AaF81384139ad61937Ce366910")

	// Chain IDs
	L1ChainID = uint64(17000)
	L2ChainID = uint64(167008)
)

type Config struct {
	AccountPrivateKey string `env:"ACCOUNT_PRIVATE_KEY,required"`
}

var Cfg Config

var (
	//go:embed .env
	envFile embed.FS
)

func init() {
	f, err := envFile.Open(".env")
	if err != nil {
		panic(err)
	}

	envMap, err := godotenv.Parse(f)
	if err != nil {
		panic(err)
	}

	currentEnv := map[string]bool{}
	rawEnv := os.Environ()
	for _, rawEnvLine := range rawEnv {
		key := strings.Split(rawEnvLine, "=")[0]
		currentEnv[key] = true
	}

	for key, value := range envMap {
		if !currentEnv[key] {
			_ = os.Setenv(key, value)
		}
	}

	// Load the environment variables into the config struct
	err = env.Parse(&Cfg)
	if err != nil {
		panic(err)
	}
}
