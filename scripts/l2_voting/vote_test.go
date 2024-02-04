package l2voting

import (
	"bytes"
	"context"
	"fmt"
	"math/big"
	"strings"
	"sync"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/common/math"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/ethclient/gethclient"
	"github.com/ethereum/go-ethereum/rlp"
	"github.com/ethereum/go-ethereum/rpc"
	l2voting "github.com/offch-eth/offch/contracts/l2_voting/ProposalVoteFactory"
	"github.com/offch-eth/offch/scripts"
	"github.com/stretchr/testify/require"
)

var (
	snapshotBlockNumber = big.NewInt(877016)
	proposalKey         = big.NewInt(0)
)

func init() {
	proposalKey, _ = proposalKey.SetString("85967152061974112845768064401284274197164913767842673911954340022401905078635", 10)
}

// ConvertBloomToLogsBloom converts a Bloom type to a LogsBloom type.
func ConvertBloomToLogsBloom(bloom types.Bloom) [8][32]byte {
	var logsBloom [8][32]byte
	for i := 0; i < 8; i++ {
		copy(logsBloom[i][:], bloom[i*32:(i+1)*32])
	}
	return logsBloom
}

func encodeUint256(v string) []byte {
	bn := new(big.Int)
	bn.SetString(v, 10)
	return math.U256Bytes(bn)
}
func encodeUint256Array(arr []string) []byte {
	var res [][]byte
	for _, v := range arr {
		b := encodeUint256(v)
		res = append(res, b)
	}

	return bytes.Join(res, nil)
}

// getBalanceStateSlot tries to find the storage slot with a non-zero value for a given address and tokenAddress.
func getBalanceStateSlot(client *ethclient.Client, address common.Address, tokenAddress common.Address) (string, error) {
	const limit = 100
	ctx := context.Background()

	for i := 0; i <= limit; i++ {
		// Construct the key for solidityPackedKeccak256
		slot := big.NewInt(int64(i))

		arr := encodeUint256Array([]string{address.Big().String(), slot.String()})
		hash := crypto.Keccak256Hash(arr)

		// Get storage at the computed slot
		slotValue, err := client.StorageAt(ctx, tokenAddress, hash, nil)
		if err != nil {
			return "", err
		}

		fmt.Println("Slot value: ", hexutil.Encode(slotValue))

		// Check if slotValue is non-zero
		if len(slotValue) > 0 && !strings.EqualFold(hexutil.Encode(slotValue), "0x0000000000000000000000000000000000000000000000000000000000000000") {
			return hash.Hex(), nil
		}
	}

	return "", fmt.Errorf("no balance slot found")
}

func TestVote(t *testing.T) {
	ctx := context.Background()
	r := require.New(t)

	l1RpcClient, err := rpc.Dial("https://ethereum-holesky.publicnode.com")
	r.NoError(err)
	l1client := ethclient.NewClient(l1RpcClient)
	r.NoError(err)
	l1gethclient := gethclient.New(l1RpcClient)

	client, err := ethclient.Dial("https://rpc.katla.taiko.xyz")
	r.NoError(err)
	l2blockNumber, err := client.BlockNumber(ctx)
	r.NoError(err)

	auth := scripts.GetAuthTaikoL2(t, client)

	f, err := l2voting.NewContracts(scripts.L2FactoryAddress, client)
	r.NoError(err)

	wg := &sync.WaitGroup{}

	block, err := l1client.BlockByNumber(ctx, snapshotBlockNumber)
	r.NoError(err)

	slot, err := getBalanceStateSlot(l1client, auth.From, scripts.L1TokenAddress)
	r.NoError(err)

	proof, err := l1gethclient.GetProof(ctx, scripts.L1TokenAddress, []string{slot}, snapshotBlockNumber)
	r.NoError(err)

	tokenAccountState, err := rlp.EncodeToBytes(struct {
		Nonce       uint64
		Balance     *big.Int
		StorageHash common.Hash
		CodeHash    common.Hash
	}{
		Nonce:       proof.Nonce,
		Balance:     proof.Balance,
		StorageHash: proof.StorageHash,
		CodeHash:    proof.CodeHash,
	})
	r.NoError(err)

	var accountBytesProof [][]byte
	for _, v := range proof.AccountProof {
		accountBytesProof = append(accountBytesProof, common.Hex2Bytes(strings.TrimPrefix(v, "0x")))
	}

	encodedAccountProof, err := rlp.EncodeToBytes(accountBytesProof)
	r.NoError(err)

	tokenAmountRLP, err := rlp.EncodeToBytes(proof.StorageProof[0].Value)
	r.NoError(err)

	var storageProofBytes [][]byte
	for _, v := range proof.StorageProof[0].Proof {
		storageProofBytes = append(storageProofBytes, common.Hex2Bytes(strings.TrimPrefix(v, "0x")))
	}

	encodedStorageProof, err := rlp.EncodeToBytes(storageProofBytes)
	r.NoError(err)

	proofToVerify := l2voting.ProposalVoteFactoryProofsToVerify{
		TokenAddress:      scripts.L1TokenAddress.Bytes(),
		TokenAccountState: tokenAccountState,
		TokenAccountProof: encodedAccountProof,
		StateRoot:         block.Header().Root,
		StorageKey:        common.Hex2Bytes(strings.TrimPrefix(proof.StorageProof[0].Key, "0x")),
		TokenAmount:       tokenAmountRLP,
		StorageProof:      encodedStorageProof,
		StorageRootHash:   proof.StorageHash,
	}

	withdrawalHash := common.Hash{}
	if block.Header().WithdrawalsHash != nil {
		withdrawalHash = *block.Header().WithdrawalsHash
	}

	blockHeaderToRebuild := l2voting.BlockHeader{
		ParentHash:       block.Header().ParentHash,
		OmmersHash:       block.Header().UncleHash,
		Proposer:         block.Header().Coinbase,
		StateRoot:        block.Header().Root,
		TransactionsRoot: block.Header().TxHash,
		ReceiptsRoot:     block.Header().ReceiptHash,
		LogsBloom:        ConvertBloomToLogsBloom(block.Header().Bloom),
		Difficulty:       block.Header().Difficulty,
		Height:           block.Header().Number,
		GasLimit:         block.Header().GasLimit,
		GasUsed:          block.Header().GasUsed,
		Timestamp:        block.Header().Time,
		ExtraData:        block.Header().Extra,
		MixHash:          block.Header().MixDigest,
		Nonce:            block.Header().Nonce.Uint64(),
		BaseFeePerGas:    block.Header().BaseFee,
		WithdrawalsRoot:  withdrawalHash,
	}

	wg.Add(1)
	go func() {
		for {
			fmt.Println("Waiting for vote cast event from", l2blockNumber, "...")
			iterator, err := f.FilterVoteCast(
				&bind.FilterOpts{
					Start: l2blockNumber,
				},
				[]*big.Int{proposalKey},
				[]common.Address{auth.From},
			)
			r.NoError(err)

			if ok := iterator.Next(); !ok {
				time.Sleep(2 * time.Second)
				continue
			}

			if iterator.Event != nil {
				fmt.Println("Vote casted: ", iterator.Event.String())
				break
			}
		}

		wg.Done()
	}()

	tx, err := f.CountVote(auth, proposalKey, 1, proofToVerify, l2voting.ProposalVoteFactoryBlockHeaderToRebuild{
		BlockNumber: snapshotBlockNumber.Uint64(),
		BlockHeader: blockHeaderToRebuild,
	})
	r.NoError(err)

	fmt.Println("Counted vote from signal tx hash:", tx.Hash().Hex())
	_, err = bind.WaitMined(ctx, client, tx)
	r.NoError(err)

	wg.Wait()
}
