// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contracts

import (
	"errors"
	"fmt"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
	_ = fmt.Sprintf
)

// GovernorCountingFromSignalProposalVoteMessage is an auto generated low-level Go binding around an user-defined struct.
type GovernorCountingFromSignalProposalVoteMessage struct {
	Id           uint64
	From         common.Address
	SrcChainId   uint64
	DestChainId  uint64
	ProposalId   *big.Int
	AgainstVotes *big.Int
	ForVotes     *big.Int
	AbstainVotes *big.Int
}

// String GovernorCountingFromSignalProposalVoteMessage returns a readable string representing the user-defined struct.
func (st *GovernorCountingFromSignalProposalVoteMessage) String() string {
	s := "User defined struct: GovernorCountingFromSignalProposalVoteMessage {\n"

	s += fmt.Sprintf("Id: %v,\n", st.Id)
	s += fmt.Sprintf("From: %v,\n", st.From)
	s += fmt.Sprintf("SrcChainId: %v,\n", st.SrcChainId)
	s += fmt.Sprintf("DestChainId: %v,\n", st.DestChainId)
	s += fmt.Sprintf("ProposalId: %v,\n", st.ProposalId)
	s += fmt.Sprintf("AgainstVotes: %v,\n", st.AgainstVotes)
	s += fmt.Sprintf("ForVotes: %v,\n", st.ForVotes)
	s += fmt.Sprintf("AbstainVotes: %v,\n", st.AbstainVotes)
	s += "}"

	return s
}

// OffchGovernorMessage is an auto generated low-level Go binding around an user-defined struct.
type OffchGovernorMessage struct {
	Id                  uint64
	From                common.Address
	SrcChainId          uint64
	DestChainId         uint64
	ProposalId          *big.Int
	SnapshotBlockNumber *big.Int
	Duration            *big.Int
	VotingTokenAddress  common.Address
}

// String OffchGovernorMessage returns a readable string representing the user-defined struct.
func (st *OffchGovernorMessage) String() string {
	s := "User defined struct: OffchGovernorMessage {\n"

	s += fmt.Sprintf("Id: %v,\n", st.Id)
	s += fmt.Sprintf("From: %v,\n", st.From)
	s += fmt.Sprintf("SrcChainId: %v,\n", st.SrcChainId)
	s += fmt.Sprintf("DestChainId: %v,\n", st.DestChainId)
	s += fmt.Sprintf("ProposalId: %v,\n", st.ProposalId)
	s += fmt.Sprintf("SnapshotBlockNumber: %v,\n", st.SnapshotBlockNumber)
	s += fmt.Sprintf("Duration: %v,\n", st.Duration)
	s += fmt.Sprintf("VotingTokenAddress: %v,\n", st.VotingTokenAddress)
	s += "}"

	return s
}

// ContractsMetaData contains all meta data concerning the Contracts contract.
var ContractsMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractIVotes\",\"name\":\"_votingToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_signalService\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"CheckpointUnorderedInsertion\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FailedInnerCall\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"voter\",\"type\":\"address\"}],\"name\":\"GovernorAlreadyCastVote\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"}],\"name\":\"GovernorAlreadyQueuedProposal\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"GovernorDisabledDeposit\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"proposer\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"votes\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"threshold\",\"type\":\"uint256\"}],\"name\":\"GovernorInsufficientProposerVotes\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"targets\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"calldatas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"values\",\"type\":\"uint256\"}],\"name\":\"GovernorInvalidProposalLength\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"quorumNumerator\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"quorumDenominator\",\"type\":\"uint256\"}],\"name\":\"GovernorInvalidQuorumFraction\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"voter\",\"type\":\"address\"}],\"name\":\"GovernorInvalidSignature\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"GovernorInvalidVoteType\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"votingPeriod\",\"type\":\"uint256\"}],\"name\":\"GovernorInvalidVotingPeriod\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"}],\"name\":\"GovernorNonexistentProposal\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"}],\"name\":\"GovernorNotQueuedProposal\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"GovernorOnlyExecutor\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"GovernorOnlyProposer\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"GovernorQueueNotImplemented\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"proposer\",\"type\":\"address\"}],\"name\":\"GovernorRestrictedProposer\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"},{\"internalType\":\"enumIGovernor.ProposalState\",\"name\":\"current\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"expectedStates\",\"type\":\"bytes32\"}],\"name\":\"GovernorUnexpectedProposalState\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"currentNonce\",\"type\":\"uint256\"}],\"name\":\"InvalidAccountNonce\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidShortString\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"QueueEmpty\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"QueueFull\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"bits\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"SafeCastOverflowedUintDowncast\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"str\",\"type\":\"string\"}],\"name\":\"StringTooLong\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"EIP712DomainChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"messageHash\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"id\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"srcChainId\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"destChainId\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"snapshotBlockNumber\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"duration\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"votingTokenAddress\",\"type\":\"address\"}],\"indexed\":false,\"internalType\":\"structOffchGovernor.Message\",\"name\":\"message\",\"type\":\"tuple\"}],\"name\":\"MessageSent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"}],\"name\":\"ProposalCanceled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"proposer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"targets\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"values\",\"type\":\"uint256[]\"},{\"indexed\":false,\"internalType\":\"string[]\",\"name\":\"signatures\",\"type\":\"string[]\"},{\"indexed\":false,\"internalType\":\"bytes[]\",\"name\":\"calldatas\",\"type\":\"bytes[]\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"voteStart\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"voteEnd\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"}],\"name\":\"ProposalCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"}],\"name\":\"ProposalExecuted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"etaSeconds\",\"type\":\"uint256\"}],\"name\":\"ProposalQueued\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldQuorumNumerator\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newQuorumNumerator\",\"type\":\"uint256\"}],\"name\":\"QuorumNumeratorUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"voter\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"support\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"weight\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"reason\",\"type\":\"string\"}],\"name\":\"VoteCast\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"voter\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"support\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"weight\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"reason\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"params\",\"type\":\"bytes\"}],\"name\":\"VoteCastWithParams\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"BALLOT_TYPEHASH\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"CLOCK_MODE\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"COUNTING_MODE\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"EXTENDED_BALLOT_TYPEHASH\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"targets\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"values\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes[]\",\"name\":\"calldatas\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes32\",\"name\":\"descriptionHash\",\"type\":\"bytes32\"}],\"name\":\"cancel\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"support\",\"type\":\"uint8\"}],\"name\":\"castVote\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"support\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"voter\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"name\":\"castVoteBySig\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"support\",\"type\":\"uint8\"},{\"internalType\":\"string\",\"name\":\"reason\",\"type\":\"string\"}],\"name\":\"castVoteWithReason\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"support\",\"type\":\"uint8\"},{\"internalType\":\"string\",\"name\":\"reason\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"params\",\"type\":\"bytes\"}],\"name\":\"castVoteWithReasonAndParams\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"support\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"voter\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"reason\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"params\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"name\":\"castVoteWithReasonAndParamsBySig\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"clock\",\"outputs\":[{\"internalType\":\"uint48\",\"name\":\"\",\"type\":\"uint48\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"eip712Domain\",\"outputs\":[{\"internalType\":\"bytes1\",\"name\":\"fields\",\"type\":\"bytes1\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"version\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"verifyingContract\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"salt\",\"type\":\"bytes32\"},{\"internalType\":\"uint256[]\",\"name\":\"extensions\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"targets\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"values\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes[]\",\"name\":\"calldatas\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes32\",\"name\":\"descriptionHash\",\"type\":\"bytes32\"}],\"name\":\"execute\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"timepoint\",\"type\":\"uint256\"}],\"name\":\"getVotes\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"timepoint\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"params\",\"type\":\"bytes\"}],\"name\":\"getVotesWithParams\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasVoted\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"targets\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"values\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes[]\",\"name\":\"calldatas\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes32\",\"name\":\"descriptionHash\",\"type\":\"bytes32\"}],\"name\":\"hashProposal\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"nonces\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"onERC1155BatchReceived\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"onERC1155Received\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"onERC721Received\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"}],\"name\":\"proposalDeadline\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"}],\"name\":\"proposalEta\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"proposalNeedsQueuing\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"}],\"name\":\"proposalProposer\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"}],\"name\":\"proposalSnapshot\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proposalThreshold\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"}],\"name\":\"proposalVotes\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"againstVotes\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"forVotes\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"abstainVotes\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"targets\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"values\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes[]\",\"name\":\"calldatas\",\"type\":\"bytes[]\"},{\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"}],\"name\":\"propose\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"targets\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"values\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes[]\",\"name\":\"calldatas\",\"type\":\"bytes[]\"},{\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"},{\"internalType\":\"uint48\",\"name\":\"_snapshotBlockNumber\",\"type\":\"uint48\"}],\"name\":\"proposeWithSnapshotBlockNumber\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"targets\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"values\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes[]\",\"name\":\"calldatas\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes32\",\"name\":\"descriptionHash\",\"type\":\"bytes32\"}],\"name\":\"queue\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"timepoint\",\"type\":\"uint256\"}],\"name\":\"quorum\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"quorumDenominator\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"timepoint\",\"type\":\"uint256\"}],\"name\":\"quorumNumerator\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"quorumNumerator\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"relay\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"app\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"_messageHash\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"id\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"srcChainId\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"destChainId\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"againstVotes\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"forVotes\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"abstainVotes\",\"type\":\"uint256\"}],\"internalType\":\"structGovernorCountingFromSignal.ProposalVoteMessage\",\"name\":\"message\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"proof\",\"type\":\"bytes\"}],\"name\":\"resultFromSignal\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"}],\"name\":\"state\",\"outputs\":[{\"internalType\":\"enumIGovernor.ProposalState\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"token\",\"outputs\":[{\"internalType\":\"contractIERC5805\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newQuorumNumerator\",\"type\":\"uint256\"}],\"name\":\"updateQuorumNumerator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"votingDelay\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"votingPeriod\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
	Bin: "0x610180604052600a8054600160a01b600160e01b0319169055600b805465ffffffffffff60a01b1916905534801562000036575f80fd5b50604051620044a5380380620044a58339810160408190526200005991620005d0565b600f82826040518060400160405280600d81526020016c27b33331b423b7bb32b93737b960991b8152508062000094620001b660201b60201c565b620000a0825f620001d1565b61012052620000b1816001620001d1565b61014052815160208084019190912060e052815190820120610100524660a0526200013e60e05161010051604080517f8b73c3c69bb8fe3d512ecc4cf759cc79239f7b179b0ffacaa9a75d522b39400f60208201529081019290925260608201524660808201523060a08201525f9060c00160405160208183030381529060405280519060200120905090565b60805250503060c0526003620001558282620006ab565b5050600780546001600160a01b0319166001600160a01b039283161790551661016052620001838162000209565b50600b80546001600160a01b039384166001600160a01b031991821617909155600a805492909316911617905562000809565b6040805180820190915260018152603160f81b602082015290565b5f602083511015620001f057620001e883620002ac565b905062000203565b81620001fd8482620006ab565b5060ff90505b92915050565b6064808211156200023c5760405163243e544560e01b815260048101839052602481018290526044015b60405180910390fd5b5f62000247620002ee565b90506200026c6200025762000309565b620002628562000348565b6009919062000381565b505060408051828152602081018590527f0553476bf02ef2726e8ce5ced78d63e26e602e4a2257b1f559418e24b4633997910160405180910390a1505050565b5f80829050601f81511115620002d9578260405163305a27a960e01b815260040162000233919062000777565b8051620002e682620007c5565b179392505050565b5f620002fb60096200039d565b6001600160d01b0316905090565b600b545f90600160a01b900465ffffffffffff1615620003395750600b54600160a01b900465ffffffffffff1690565b62000343620003ea565b905090565b5f6001600160d01b038211156200037d576040516306dfcc6560e41b815260d060048201526024810183905260440162000233565b5090565b5f8062000390858585620003f6565b915091505b935093915050565b80545f908015620003e157620003c783620003ba600184620007e9565b5f91825260209091200190565b54660100000000000090046001600160d01b0316620003e3565b5f5b9392505050565b5f620003434362000584565b82545f908190801562000526575f6200041687620003ba600185620007e9565b60408051808201909152905465ffffffffffff80821680845266010000000000009092046001600160d01b0316602084015291925090871610156200046e57604051632520601d60e01b815260040160405180910390fd5b805165ffffffffffff808816911603620004c257846200049588620003ba600186620007e9565b80546001600160d01b039290921666010000000000000265ffffffffffff90921691909117905562000515565b6040805180820190915265ffffffffffff80881682526001600160d01b0380881660208085019182528b54600181018d555f8d815291909120945191519092166601000000000000029216919091179101555b602001519250839150620003959050565b50506040805180820190915265ffffffffffff80851682526001600160d01b0380851660208085019182528854600181018a555f8a815291822095519251909316660100000000000002919093161792019190915590508162000395565b5f65ffffffffffff8211156200037d576040516306dfcc6560e41b8152603060048201526024810183905260440162000233565b6001600160a01b0381168114620005cd575f80fd5b50565b5f8060408385031215620005e2575f80fd5b8251620005ef81620005b8565b60208401519092506200060281620005b8565b809150509250929050565b634e487b7160e01b5f52604160045260245ffd5b600181811c908216806200063657607f821691505b6020821081036200065557634e487b7160e01b5f52602260045260245ffd5b50919050565b601f821115620006a657805f5260205f20601f840160051c81016020851015620006825750805b601f840160051c820191505b81811015620006a3575f81556001016200068e565b50505b505050565b81516001600160401b03811115620006c757620006c76200060d565b620006df81620006d8845462000621565b846200065b565b602080601f83116001811462000715575f8415620006fd5750858301515b5f19600386901b1c1916600185901b1785556200076f565b5f85815260208120601f198616915b82811015620007455788860151825594840194600190910190840162000724565b50858210156200076357878501515f19600388901b60f8161c191681555b505060018460011b0185555b505050505050565b5f602080835283518060208501525f5b81811015620007a55785810183015185820160400152820162000787565b505f604082860101526040601f19601f8301168501019250505092915050565b8051602080830151919081101562000655575f1960209190910360031b1b16919050565b818103818111156200020357634e487b7160e01b5f52601160045260245ffd5b60805160a05160c05160e05161010051610120516101405161016051613c2b6200087a5f395f81816108ea01528181610d22015281816115c50152611ea201525f611e6f01525f611e4301525f61212701525f6120ff01525f61205a01525f61208401525f6120ae0152613c2b5ff3fe60806040526004361061026d575f3560e01c80637b3c71d31161014a578063ab58fb8e116100be578063dd4e2ba511610078578063dd4e2ba5146107fb578063deaaa7cc14610840578063eb9019d414610873578063f23a6e6114610892578063f8ce560a146108bd578063fc0c546a146108dc575f80fd5b8063ab58fb8e14610749578063b58131b014610439578063bc197c811461077f578063c01f9e37146107aa578063c28bc2fa146107c9578063c59057e4146107dc575f80fd5b806391ddadf41161010f57806391ddadf41461069a57806397c3d334146106c55780639a802a6d146106d85780639fbd2818146106f7578063a7713a7014610716578063a9a952941461072a575f80fd5b80637b3c71d3146105e25780637d5e81e2146106015780637ecebe001461062057806384b0196e146106545780638ff262e31461067b575f80fd5b80633e4f49e6116101e157806354fd4d50116101a657806354fd4d501461051e57806356781388146105475780635b8d0e0d146105665780635f398a141461058557806360c4247f146105a457806365218612146105c3575f80fd5b80633e4f49e61461044b5780634385963214610477578063452115d6146104985780634bf5d7e9146104b7578063544ffc9c146104cb575f80fd5b8063150b7a0211610232578063150b7a0214610357578063160cbed71461039a5780632656227d146103b95780632d63f693146103cc5780632fe3e261146104065780633932abb114610439575f80fd5b806301ffc9a71461027a57806302a251a3146102ae57806306f3f9e6146102cb57806306fdde03146102ea578063143489d01461030b575f80fd5b3661027657005b005b5f80fd5b348015610285575f80fd5b50610299610294366004612acd565b61090e565b60405190151581526020015b60405180910390f35b3480156102b9575f80fd5b5060145b6040519081526020016102a5565b3480156102d6575f80fd5b506102746102e5366004612af4565b61095f565b3480156102f5575f80fd5b506102fe610973565b6040516102a59190612b58565b348015610316575f80fd5b5061033f610325366004612af4565b5f908152600460205260409020546001600160a01b031690565b6040516001600160a01b0390911681526020016102a5565b348015610362575f80fd5b50610381610371366004612c65565b630a85bd0160e11b949350505050565b6040516001600160e01b031990911681526020016102a5565b3480156103a5575f80fd5b506102bd6103b4366004612e2e565b610a03565b6102bd6103c7366004612e2e565b610a41565b3480156103d7575f80fd5b506102bd6103e6366004612af4565b5f90815260046020526040902054600160a01b900465ffffffffffff1690565b348015610411575f80fd5b506102bd7f3e83946653575f9a39005e1545185629e92736b7528ab20ca3816f315424a81181565b348015610444575f80fd5b505f6102bd565b348015610456575f80fd5b5061046a610465366004612af4565b610b69565b6040516102a59190612eeb565b348015610482575f80fd5b50610299610491366004612ef9565b5f92915050565b3480156104a3575f80fd5b506102bd6104b2366004612e2e565b610cb2565b3480156104c2575f80fd5b506102fe610d1e565b3480156104d6575f80fd5b506105036104e5366004612af4565b5f908152600860205260409020805460018201546002909201549092565b604080519384526020840192909252908201526060016102a5565b348015610529575f80fd5b506040805180820190915260018152603160f81b60208201526102fe565b348015610552575f80fd5b506102bd610561366004612f33565b610dde565b348015610571575f80fd5b506102bd610580366004612f98565b610e05565b348015610590575f80fd5b506102bd61059f366004613048565b610f61565b3480156105af575f80fd5b506102bd6105be366004612af4565b610fb4565b3480156105ce575f80fd5b506102bd6105dd3660046130c5565b611040565b3480156105ed575f80fd5b506102bd6105fc366004613188565b6110ed565b34801561060c575f80fd5b506102bd61061b3660046131dd565b611133565b34801561062b575f80fd5b506102bd61063a366004613276565b6001600160a01b03165f9081526002602052604090205490565b34801561065f575f80fd5b50610668611192565b6040516102a597969594939291906132c9565b348015610686575f80fd5b506102bd610695366004613338565b6111d4565b3480156106a5575f80fd5b506106ae6112a3565b60405165ffffffffffff90911681526020016102a5565b3480156106d0575f80fd5b5060646102bd565b3480156106e3575f80fd5b506102bd6106f2366004613383565b6112df565b348015610702575f80fd5b506102746107113660046133eb565b6112f5565b348015610721575f80fd5b506102bd61146c565b348015610735575f80fd5b50610299610744366004612af4565b505f90565b348015610754575f80fd5b506102bd610763366004612af4565b5f9081526004602052604090206001015465ffffffffffff1690565b34801561078a575f80fd5b506103816107993660046134d1565b63bc197c8160e01b95945050505050565b3480156107b5575f80fd5b506102bd6107c4366004612af4565b611485565b6102746107d7366004613559565b6114c7565b3480156107e7575f80fd5b506102bd6107f6366004612e2e565b611543565b348015610806575f80fd5b506040805180820190915260208082527f737570706f72743d627261766f2671756f72756d3d666f722c6162737461696e908201526102fe565b34801561084b575f80fd5b506102bd7ff2aad550cf55f045cb27e9c559f9889fdfb6e6cdaa032301d6ea397784ae51d781565b34801561087e575f80fd5b506102bd61088d366004613596565b61157c565b34801561089d575f80fd5b506103816108ac3660046135be565b63f23a6e6160e01b95945050505050565b3480156108c8575f80fd5b506102bd6108d7366004612af4565b61159b565b3480156108e7575f80fd5b507f000000000000000000000000000000000000000000000000000000000000000061033f565b5f6001600160e01b031982166332a2ad4360e11b148061093e57506001600160e01b03198216630271189760e51b145b8061095957506301ffc9a760e01b6001600160e01b03198316145b92915050565b610967611642565b61097081611679565b50565b6060600380546109829061361d565b80601f01602080910402602001604051908101604052809291908181526020018280546109ae9061361d565b80156109f95780601f106109d0576101008083540402835291602001916109f9565b820191905f5260205f20905b8154815290600101906020018083116109dc57829003601f168201915b5050505050905090565b5f80610a1186868686611543565b9050610a2681610a21600461170e565b611730565b505f604051634844252360e11b815260040160405180910390fd5b5f80610a4f86868686611543565b9050610a6f81610a5f600561170e565b610a69600461170e565b17611730565b505f818152600460205260409020805460ff60f01b1916600160f01b17905530610a963090565b6001600160a01b031614610b1f575f5b8651811015610b1d57306001600160a01b0316878281518110610acb57610acb613655565b60200260200101516001600160a01b031603610b1557610b15858281518110610af657610af6613655565b602002602001015180519060200120600561176d90919063ffffffff16565b600101610aa6565b505b610b2c81878787876117dd565b6040518181527f712ae1383f79ac853f8d882153778e0260ef8f03b504e2866e0593e04d2b291f906020015b60405180910390a195945050505050565b5f818152600460205260408120805460ff600160f01b8204811691600160f81b9004168115610b9d57506007949350505050565b8015610bae57506002949350505050565b5f85815260046020526040812054600160a01b900465ffffffffffff169050805f03610bf557604051636ad0607560e01b8152600481018790526024015b60405180910390fd5b5f610bfe6112a3565b65ffffffffffff169050808210610c1b57505f9695505050505050565b5f610c2588611485565b9050818110610c3c57506001979650505050505050565b610c45886118b2565b1580610c6457505f888152600860205260409020805460019091015411155b15610c7757506003979650505050505050565b5f8881526004602052604090206001015465ffffffffffff165f03610ca457506004979650505050505050565b506005979650505050505050565b5f80610cc086868686611543565b9050610ccf81610a215f61170e565b505f818152600460205260409020546001600160a01b03163314610d085760405163233d98e360e01b8152336004820152602401610bec565b610d1486868686611903565b9695505050505050565b60607f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316634bf5d7e96040518163ffffffff1660e01b81526004015f60405180830381865afa925050508015610d9d57506040513d5f823e601f3d908101601f19168201604052610d9a9190810190613669565b60015b610dd9575060408051808201909152601d81527f6d6f64653d626c6f636b6e756d6265722666726f6d3d64656661756c74000000602082015290565b919050565b5f80339050610dfd84828560405180602001604052805f8152506119b2565b949350505050565b5f80610ee687610ee07f3e83946653575f9a39005e1545185629e92736b7528ab20ca3816f315424a8118c8c8c610e588e6001600160a01b03165f90815260026020526040902080546001810190915590565b8d8d604051610e689291906136d1565b60405180910390208c80519060200120604051602001610ec59796959493929190968752602087019590955260ff9390931660408601526001600160a01b03919091166060850152608084015260a083015260c082015260e00190565b604051602081830303815290604052805190602001206119dc565b85611a08565b905080610f11576040516394ab6c0760e01b81526001600160a01b0388166004820152602401610bec565b610f5489888a89898080601f0160208091040260200160405190810160405280939291908181526020018383808284375f920191909152508b9250611a5d915050565b9998505050505050505050565b5f80339050610fa987828888888080601f0160208091040260200160405190810160405280939291908181526020018383808284375f920191909152508a9250611a5d915050565b979650505050505050565b600980545f918290610fc76001846136f4565b81548110610fd757610fd7613655565b5f918252602090912001805490915065ffffffffffff811690600160301b90046001600160d01b0316858211611019576001600160d01b031695945050505050565b61102d61102587611abf565b600990611af5565b6001600160d01b03169695505050505050565b600b805465ffffffffffff60a01b1916600160a01b65ffffffffffff8416021790555f3361106e8185611ba4565b6110965760405163d9b3955760e01b81526001600160a01b0382166004820152602401610bec565b5f6110bc8260016110a56112a3565b6110af9190613707565b65ffffffffffff1661157c565b90505f806110cd8a8a8a8a88611c8c565b600b805465ffffffffffff60a01b191690559a9950505050505050505050565b5f80339050610d1486828787878080601f0160208091040260200160405190810160405280939291908181526020018383808284375f920191909152506119b292505050565b5f3361113f8184611ba4565b6111675760405163d9b3955760e01b81526001600160a01b0382166004820152602401610bec565b5f6111768260016110a56112a3565b90505f6111868888888887611c8c565b98975050505050505050565b5f6060805f805f60606111a3611e3c565b6111ab611e68565b604080515f80825260208201909252600f60f81b9b939a50919850469750309650945092509050565b5f8061125e84610ee07ff2aad550cf55f045cb27e9c559f9889fdfb6e6cdaa032301d6ea397784ae51d78989896112278b6001600160a01b03165f90815260026020526040902080546001810190915590565b60408051602081019690965285019390935260ff90911660608401526001600160a01b0316608083015260a082015260c001610ec5565b905080611289576040516394ab6c0760e01b81526001600160a01b0385166004820152602401610bec565b610d1486858760405180602001604052805f8152506119b2565b600b545f90600160a01b900465ffffffffffff16156112d25750600b54600160a01b900465ffffffffffff1690565b6112da611e95565b905090565b5f6112eb848484611e9f565b90505b9392505050565b600754604080850151905163910af6ed60e01b81526001600160a01b039092169163910af6ed9161133091899089908890889060040161372d565b602060405180830381865afa15801561134b573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061136f9190613784565b6113ad5760405162461bcd60e51b815260206004820152600f60248201526e24b73b30b634b21039b4b3b730b61760891b6044820152606401610bec565b60808301515f9081526008602052604090206003015460ff16156114265760405162461bcd60e51b815260206004820152602a60248201527f526573756c74732061726520616c72656164792073657420666f72207468697360448201526910383937b837b9b0b61760b11b6064820152608401610bec565b505060808101515f90815260086020526040902060a0820151815560c082015160018083019190915560e0909201516002820155600301805460ff191690911790555050565b5f6114776009611f32565b6001600160d01b0316905090565b5f818152600460205260408120546114b990600160d01b810463ffffffff1690600160a01b900465ffffffffffff166137a3565b65ffffffffffff1692915050565b6114cf611642565b5f80856001600160a01b03168585856040516114ec9291906136d1565b5f6040518083038185875af1925050503d805f8114611526576040519150601f19603f3d011682016040523d82523d5f602084013e61152b565b606091505b509150915061153a8282611f69565b50505050505050565b5f8484848460405160200161155b9493929190613852565b60408051601f19818403018152919052805160209091012095945050505050565b5f6112ee838361159660408051602081019091525f815290565b611e9f565b5f60646115a783610fb4565b604051632394e7a360e21b8152600481018590526001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001690638e539e8c90602401602060405180830381865afa15801561160a573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061162e919061389c565b61163891906138b3565b61095991906138de565b303314611664576040516347096e4760e01b8152336004820152602401610bec565b565b806116716005611f85565b036116665750565b6064808211156116a65760405163243e544560e01b81526004810183905260248101829052604401610bec565b5f6116af61146c565b90506116ce6116bc6112a3565b6116c585612001565b60099190612034565b505060408051828152602081018590527f0553476bf02ef2726e8ce5ced78d63e26e602e4a2257b1f559418e24b4633997910160405180910390a1505050565b5f81600781111561172157611721612eb7565b600160ff919091161b92915050565b5f8061173b84610b69565b90505f836117488361170e565b16036112ee578381846040516331b75e4d60e01b8152600401610bec939291906138fd565b81546001600160801b03600160801b8204811691811660018301909116036117a857604051638acb5f2760e01b815260040160405180910390fd5b6001600160801b038082165f90815260018086016020526040909120939093558354919092018216600160801b029116179055565b5f5b84518110156118aa575f808683815181106117fc576117fc613655565b60200260200101516001600160a01b031686848151811061181f5761181f613655565b602002602001015186858151811061183957611839613655565b602002602001015160405161184e919061391f565b5f6040518083038185875af1925050503d805f8114611888576040519150601f19603f3d011682016040523d82523d5f602084013e61188d565b606091505b509150915061189c8282611f69565b5050508060010190506117df565b505050505050565b5f818152600860205260408120600281015460018201546118d3919061393a565b6118fa6108d7855f90815260046020526040902054600160a01b900465ffffffffffff1690565b11159392505050565b5f8061191186868686611543565b905061195f81611921600761170e565b61192b600661170e565b611935600261170e565b600161194260078261394d565b61194d906002613a46565b61195791906136f4565b181818611730565b505f818152600460205260409081902080546001600160f81b0316600160f81b179055517f789cf55be980739dad1d0699b93b58e806b51c9d96619bfa8fe0a28abaa7b30c90610b589083815260200190565b5f6119d3858585856119ce60408051602081019091525f815290565b611a5d565b95945050505050565b5f6109596119e861204e565b8360405161190160f01b8152600281019290925260228201526042902090565b5f805f611a158585612177565b5090925090505f816003811115611a2e57611a2e612eb7565b148015611a4c5750856001600160a01b0316826001600160a01b0316145b80610d145750610d148686866121c0565b60405162461bcd60e51b815260206004820152602f60248201527f4f66666368476f7665726e65723a20796f75206861766520746f20636173742060448201526e3b37ba32903a343937bab3b410261960891b60648201525f90608401610bec565b5f65ffffffffffff821115611af1576040516306dfcc6560e41b81526030600482015260248101839052604401610bec565b5090565b81545f9081816005811115611b51575f611b0e84612296565b611b1890856136f4565b5f8881526020902090915081015465ffffffffffff9081169087161015611b4157809150611b4f565b611b4c81600161393a565b92505b505b5f611b5e8787858561237a565b90508015611b9857611b8287611b756001846136f4565b5f91825260209091200190565b54600160301b90046001600160d01b0316610fa9565b5f979650505050505050565b80515f906034811015611bbb576001915050610959565b82810160131901516001600160a01b031981166b046e0e4dee0dee6cae47a60f60a31b14611bee57600192505050610959565b5f80611bfb6028856136f4565b90505b83811015611c6b575f80611c31888481518110611c1d57611c1d613655565b01602001516001600160f81b0319166123d9565b9150915081611c495760019650505050505050610959565b8060ff166004856001600160a01b0316901b1793505050806001019050611bfe565b50856001600160a01b0316816001600160a01b031614935050505092915050565b5f611c9a8686868686612469565b600b549091504390600160a01b900465ffffffffffff1615611cca5750600b54600160a01b900465ffffffffffff165b6040805161010081018252600a8054600160a01b90046001600160401b03168083523060208401526142689383019390935262028c6060608301526080820185905260a08201849052601460c08301819052600b546001600160a01b031660e0840152919291611d3983613a54565b91906101000a8154816001600160401b0302191690836001600160401b03160217905550505f81604051602001611d709190613ae4565b60408051808303601f190181529082905280516020820120600a5463019b28af60e61b845260048401829052919350916001600160a01b03909116906366ca2bc0906024016020604051808303815f875af1158015611dd1573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190611df5919061389c565b507f3d151231d49a27ffabff6c55740d7f90f6fa45e81fcfc2dec9ecf1637488e1718184604051611e27929190613af3565b60405180910390a15050505095945050505050565b60606112da7f00000000000000000000000000000000000000000000000000000000000000005f612667565b60606112da7f00000000000000000000000000000000000000000000000000000000000000006001612667565b5f6112da43611abf565b5f7f0000000000000000000000000000000000000000000000000000000000000000604051630748d63560e31b81526001600160a01b038681166004830152602482018690529190911690633a46b1a890604401602060405180830381865afa158015611f0e573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906112eb919061389c565b80545f908015611f6157611f4b83611b756001846136f4565b54600160301b90046001600160d01b03166112ee565b5f9392505050565b606082611f7e57611f7982612710565b610959565b5080610959565b80545f906001600160801b0380821691600160801b9004168103611fbc576040516375e52f4f60e01b815260040160405180910390fd5b6001600160801b038181165f908152600185810160205260408220805492905585546fffffffffffffffffffffffffffffffff19169301909116919091179092555090565b5f6001600160d01b03821115611af1576040516306dfcc6560e41b815260d0600482015260248101839052604401610bec565b5f80612041858585612739565b915091505b935093915050565b5f306001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000161480156120a657507f000000000000000000000000000000000000000000000000000000000000000046145b156120d057507f000000000000000000000000000000000000000000000000000000000000000090565b6112da604080517f8b73c3c69bb8fe3d512ecc4cf759cc79239f7b179b0ffacaa9a75d522b39400f60208201527f0000000000000000000000000000000000000000000000000000000000000000918101919091527f000000000000000000000000000000000000000000000000000000000000000060608201524660808201523060a08201525f9060c00160405160208183030381529060405280519060200120905090565b5f805f83516041036121ae576020840151604085015160608601515f1a6121a0888285856128af565b9550955095505050506121b9565b505081515f91506002905b9250925092565b5f805f856001600160a01b031685856040516024016121e0929190613b08565b60408051601f198184030181529181526020820180516001600160e01b0316630b135d3f60e11b17905251612215919061391f565b5f60405180830381855afa9150503d805f811461224d576040519150601f19603f3d011682016040523d82523d5f602084013e612252565b606091505b509150915081801561226657506020815110155b8015610d1457508051630b135d3f60e11b9061228b908301602090810190840161389c565b149695505050505050565b5f815f036122a557505f919050565b5f60016122b184612977565b901c6001901b905060018184816122ca576122ca6138ca565b048201901c905060018184816122e2576122e26138ca565b048201901c905060018184816122fa576122fa6138ca565b048201901c90506001818481612312576123126138ca565b048201901c9050600181848161232a5761232a6138ca565b048201901c90506001818481612342576123426138ca565b048201901c9050600181848161235a5761235a6138ca565b048201901c90506112ee81828581612374576123746138ca565b04612a0a565b5f5b818310156123d1575f61238f8484612a1f565b5f8781526020902090915065ffffffffffff86169082015465ffffffffffff1611156123bd578092506123cb565b6123c881600161393a565b93505b5061237c565b509392505050565b5f8060f883901c602f811180156123f35750603a8160ff16105b1561240857600194602f199091019350915050565b8060ff16604010801561241e575060478160ff16105b15612433576001946036199091019350915050565b8060ff166060108015612449575060678160ff16105b1561245e576001946056199091019350915050565b505f93849350915050565b5f61247d8686868680519060200120611543565b90508451865114158061249257508351865114155b8061249c57508551155b156124d157855184518651604051630447b05d60e41b8152600481019390935260248301919091526044820152606401610bec565b5f81815260046020526040902054600160a01b900465ffffffffffff161561251a57806124fd82610b69565b6040516331b75e4d60e01b8152610bec9291905f906004016138fd565b5f806125246112a3565b65ffffffffffff16612536919061393a565b90505f60145f84815260046020526040902080546001600160a01b0319166001600160a01b03871617815590915061256d83611abf565b815465ffffffffffff91909116600160a01b0265ffffffffffff60a01b1990911617815561259a82612a39565b815463ffffffff91909116600160d01b0263ffffffff60d01b1990911617815588517f7d84a6263ae0d98d3329bd7b46bb4e8d6f98cd35a7adb45c274c8b7fd5ebd5e090859087908c908c906001600160401b038111156125fd576125fd612b80565b60405190808252806020026020018201604052801561263057816020015b606081526020019060019003908161261b5790505b508c8961263d8a8261393a565b8e60405161265399989796959493929190613b20565b60405180910390a150505095945050505050565b606060ff83146126815761267a83612a69565b9050610959565b81805461268d9061361d565b80601f01602080910402602001604051908101604052809291908181526020018280546126b99061361d565b80156127045780601f106126db57610100808354040283529160200191612704565b820191905f5260205f20905b8154815290600101906020018083116126e757829003601f168201915b50505050509050610959565b8051156127205780518082602001fd5b604051630a12f52160e11b815260040160405180910390fd5b82545f9081908015612855575f61275587611b756001856136f4565b60408051808201909152905465ffffffffffff808216808452600160301b9092046001600160d01b0316602084015291925090871610156127a957604051632520601d60e01b815260040160405180910390fd5b805165ffffffffffff8088169116036127f557846127cc88611b756001866136f4565b80546001600160d01b0392909216600160301b0265ffffffffffff909216919091179055612845565b6040805180820190915265ffffffffffff80881682526001600160d01b0380881660208085019182528b54600181018d555f8d81529190912094519151909216600160301b029216919091179101555b6020015192508391506120469050565b50506040805180820190915265ffffffffffff80851682526001600160d01b0380851660208085019182528854600181018a555f8a815291822095519251909316600160301b029190931617920191909155905081612046565b5f80807f7fffffffffffffffffffffffffffffff5d576e7357a4501ddfe92f46681b20a08411156128e857505f9150600390508261296d565b604080515f808252602082018084528a905260ff891692820192909252606081018790526080810186905260019060a0016020604051602081039080840390855afa158015612939573d5f803e3d5ffd5b5050604051601f1901519150506001600160a01b03811661296457505f92506001915082905061296d565b92505f91508190505b9450945094915050565b5f80608083901c1561298b57608092831c92015b604083901c1561299d57604092831c92015b602083901c156129af57602092831c92015b601083901c156129c157601092831c92015b600883901c156129d357600892831c92015b600483901c156129e557600492831c92015b600283901c156129f757600292831c92015b600183901c156109595760010192915050565b5f818310612a1857816112ee565b5090919050565b5f612a2d60028484186138de565b6112ee9084841661393a565b5f63ffffffff821115611af1576040516306dfcc6560e41b81526020600482015260248101839052604401610bec565b60605f612a7583612aa6565b6040805160208082528183019092529192505f91906020820181803683375050509182525060208101929092525090565b5f60ff8216601f81111561095957604051632cd44ac360e21b815260040160405180910390fd5b5f60208284031215612add575f80fd5b81356001600160e01b0319811681146112ee575f80fd5b5f60208284031215612b04575f80fd5b5035919050565b5f5b83811015612b25578181015183820152602001612b0d565b50505f910152565b5f8151808452612b44816020860160208601612b0b565b601f01601f19169290920160200192915050565b602081525f6112ee6020830184612b2d565b80356001600160a01b0381168114610dd9575f80fd5b634e487b7160e01b5f52604160045260245ffd5b60405161010081016001600160401b0381118282101715612bb757612bb7612b80565b60405290565b604051601f8201601f191681016001600160401b0381118282101715612be557612be5612b80565b604052919050565b5f6001600160401b03821115612c0557612c05612b80565b50601f01601f191660200190565b5f82601f830112612c22575f80fd5b8135612c35612c3082612bed565b612bbd565b818152846020838601011115612c49575f80fd5b816020850160208301375f918101602001919091529392505050565b5f805f8060808587031215612c78575f80fd5b612c8185612b6a565b9350612c8f60208601612b6a565b92506040850135915060608501356001600160401b03811115612cb0575f80fd5b612cbc87828801612c13565b91505092959194509250565b5f6001600160401b03821115612ce057612ce0612b80565b5060051b60200190565b5f82601f830112612cf9575f80fd5b81356020612d09612c3083612cc8565b8083825260208201915060208460051b870101935086841115612d2a575f80fd5b602086015b84811015612d4d57612d4081612b6a565b8352918301918301612d2f565b509695505050505050565b5f82601f830112612d67575f80fd5b81356020612d77612c3083612cc8565b8083825260208201915060208460051b870101935086841115612d98575f80fd5b602086015b84811015612d4d5780358352918301918301612d9d565b5f82601f830112612dc3575f80fd5b81356020612dd3612c3083612cc8565b82815260059290921b84018101918181019086841115612df1575f80fd5b8286015b84811015612d4d5780356001600160401b03811115612e12575f80fd5b612e208986838b0101612c13565b845250918301918301612df5565b5f805f8060808587031215612e41575f80fd5b84356001600160401b0380821115612e57575f80fd5b612e6388838901612cea565b95506020870135915080821115612e78575f80fd5b612e8488838901612d58565b94506040870135915080821115612e99575f80fd5b50612ea687828801612db4565b949793965093946060013593505050565b634e487b7160e01b5f52602160045260245ffd5b60088110612ee757634e487b7160e01b5f52602160045260245ffd5b9052565b602081016109598284612ecb565b5f8060408385031215612f0a575f80fd5b82359150612f1a60208401612b6a565b90509250929050565b803560ff81168114610dd9575f80fd5b5f8060408385031215612f44575f80fd5b82359150612f1a60208401612f23565b5f8083601f840112612f64575f80fd5b5081356001600160401b03811115612f7a575f80fd5b602083019150836020828501011115612f91575f80fd5b9250929050565b5f805f805f805f60c0888a031215612fae575f80fd5b87359650612fbe60208901612f23565b9550612fcc60408901612b6a565b945060608801356001600160401b0380821115612fe7575f80fd5b612ff38b838c01612f54565b909650945060808a013591508082111561300b575f80fd5b6130178b838c01612c13565b935060a08a013591508082111561302c575f80fd5b506130398a828b01612c13565b91505092959891949750929550565b5f805f805f6080868803121561305c575f80fd5b8535945061306c60208701612f23565b935060408601356001600160401b0380821115613087575f80fd5b61309389838a01612f54565b909550935060608801359150808211156130ab575f80fd5b506130b888828901612c13565b9150509295509295909350565b5f805f805f60a086880312156130d9575f80fd5b85356001600160401b03808211156130ef575f80fd5b6130fb89838a01612cea565b96506020880135915080821115613110575f80fd5b61311c89838a01612d58565b95506040880135915080821115613131575f80fd5b61313d89838a01612db4565b94506060880135915080821115613152575f80fd5b5061315f88828901612c13565b925050608086013565ffffffffffff8116811461317a575f80fd5b809150509295509295909350565b5f805f806060858703121561319b575f80fd5b843593506131ab60208601612f23565b925060408501356001600160401b038111156131c5575f80fd5b6131d187828801612f54565b95989497509550505050565b5f805f80608085870312156131f0575f80fd5b84356001600160401b0380821115613206575f80fd5b61321288838901612cea565b95506020870135915080821115613227575f80fd5b61323388838901612d58565b94506040870135915080821115613248575f80fd5b61325488838901612db4565b93506060870135915080821115613269575f80fd5b50612cbc87828801612c13565b5f60208284031215613286575f80fd5b6112ee82612b6a565b5f815180845260208085019450602084015f5b838110156132be578151875295820195908201906001016132a2565b509495945050505050565b60ff60f81b8816815260e060208201525f6132e760e0830189612b2d565b82810360408401526132f98189612b2d565b606084018890526001600160a01b038716608085015260a0840186905283810360c0850152905061332a818561328f565b9a9950505050505050505050565b5f805f806080858703121561334b575f80fd5b8435935061335b60208601612f23565b925061336960408601612b6a565b915060608501356001600160401b03811115612cb0575f80fd5b5f805f60608486031215613395575f80fd5b61339e84612b6a565b92506020840135915060408401356001600160401b038111156133bf575f80fd5b6133cb86828701612c13565b9150509250925092565b80356001600160401b0381168114610dd9575f80fd5b5f805f805f858703610160811215613401575f80fd5b61340a87612b6a565b95506020870135945061010080603f1983011215613426575f80fd5b61342e612b94565b915061343c604089016133d5565b825261344a60608901612b6a565b602083015261345b608089016133d5565b604083015261346c60a089016133d5565b606083015260c088810135608084015260e0808a013560a085015291890135908301526101208801359082015292506101408601356001600160401b038111156134b4575f80fd5b6134c088828901612f54565b969995985093965092949392505050565b5f805f805f60a086880312156134e5575f80fd5b6134ee86612b6a565b94506134fc60208701612b6a565b935060408601356001600160401b0380821115613517575f80fd5b61352389838a01612d58565b94506060880135915080821115613538575f80fd5b61354489838a01612d58565b935060808801359150808211156130ab575f80fd5b5f805f806060858703121561356c575f80fd5b61357585612b6a565b93506020850135925060408501356001600160401b038111156131c5575f80fd5b5f80604083850312156135a7575f80fd5b6135b083612b6a565b946020939093013593505050565b5f805f805f60a086880312156135d2575f80fd5b6135db86612b6a565b94506135e960208701612b6a565b9350604086013592506060860135915060808601356001600160401b03811115613611575f80fd5b6130b888828901612c13565b600181811c9082168061363157607f821691505b60208210810361364f57634e487b7160e01b5f52602260045260245ffd5b50919050565b634e487b7160e01b5f52603260045260245ffd5b5f60208284031215613679575f80fd5b81516001600160401b0381111561368e575f80fd5b8201601f8101841361369e575f80fd5b80516136ac612c3082612bed565b8181528560208385010111156136c0575f80fd5b6119d3826020830160208601612b0b565b818382375f9101908152919050565b634e487b7160e01b5f52601160045260245ffd5b81810381811115610959576109596136e0565b65ffffffffffff828116828216039080821115613726576137266136e0565b5092915050565b6001600160401b03861681526001600160a01b0385166020820152604081018490526080606082018190528101829052818360a08301375f81830160a090810191909152601f909201601f19160101949350505050565b5f60208284031215613794575f80fd5b815180151581146112ee575f80fd5b65ffffffffffff818116838216019080821115613726576137266136e0565b5f815180845260208085019450602084015f5b838110156132be5781516001600160a01b0316875295820195908201906001016137d5565b5f8282518085526020808601955060208260051b840101602086015f5b8481101561384557601f19868403018952613833838351612b2d565b98840198925090830190600101613817565b5090979650505050505050565b608081525f61386460808301876137c2565b8281036020840152613876818761328f565b9050828103604084015261388a81866137fa565b91505082606083015295945050505050565b5f602082840312156138ac575f80fd5b5051919050565b8082028115828204841417610959576109596136e0565b634e487b7160e01b5f52601260045260245ffd5b5f826138f857634e487b7160e01b5f52601260045260245ffd5b500490565b838152606081016139116020830185612ecb565b826040830152949350505050565b5f8251613930818460208701612b0b565b9190910192915050565b80820180821115610959576109596136e0565b60ff8181168382160190811115610959576109596136e0565b600181815b808511156139a057815f1904821115613986576139866136e0565b8085161561399357918102915b93841c939080029061396b565b509250929050565b5f826139b657506001610959565b816139c257505f610959565b81600181146139d857600281146139e2576139fe565b6001915050610959565b60ff8411156139f3576139f36136e0565b50506001821b610959565b5060208310610133831016604e8410600b8410161715613a21575081810a610959565b613a2b8383613966565b805f1904821115613a3e57613a3e6136e0565b029392505050565b5f6112ee60ff8416836139a8565b5f6001600160401b03808316818103613a6f57613a6f6136e0565b6001019392505050565b6001600160401b03808251168352602082015160018060a01b0380821660208601528260408501511660408601528260608501511660608601526080840151608086015260a084015160a086015260c084015160c08601528060e08501511660e08601525050505050565b61010081016109598284613a79565b82815261012081016112ee6020830184613a79565b828152604060208201525f6112eb6040830184612b2d565b5f6101208b8352602060018060a01b038c1681850152816040850152613b488285018c6137c2565b91508382036060850152613b5c828b61328f565b915083820360808501528189518084528284019150828160051b850101838c015f5b83811015613bac57601f19878403018552613b9a838351612b2d565b94860194925090850190600101613b7e565b505086810360a0880152613bc0818c6137fa565b9450505050508560c08401528460e0840152828103610100840152613be58185612b2d565b9c9b50505050505050505050505056fea2646970667358221220eeb622d7c12ddb4ba8f25736b5ed0b9c70e8dc636fdacf869fbae4e3e11a82c764736f6c63430008180033",
}

// ContractsABI is the input ABI used to generate the binding from.
// Deprecated: Use ContractsMetaData.ABI instead.
var ContractsABI = ContractsMetaData.ABI

// ContractsBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ContractsMetaData.Bin instead.
var ContractsBin = ContractsMetaData.Bin

// DeployContracts deploys a new Ethereum contract, binding an instance of Contracts to it.
func DeployContracts(auth *bind.TransactOpts, backend bind.ContractBackend, _votingToken common.Address, _signalService common.Address) (common.Address, *types.Transaction, *Contracts, error) {
	parsed, err := ContractsMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ContractsBin), backend, _votingToken, _signalService)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Contracts{ContractsCaller: ContractsCaller{contract: contract}, ContractsTransactor: ContractsTransactor{contract: contract}, ContractsFilterer: ContractsFilterer{contract: contract}}, nil
}

// Contracts is an auto generated Go binding around an Ethereum contract.
type Contracts struct {
	ContractsCaller     // Read-only binding to the contract
	ContractsTransactor // Write-only binding to the contract
	ContractsFilterer   // Log filterer for contract events
}

// ContractsCaller is an auto generated read-only Go binding around an Ethereum contract.
type ContractsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ContractsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ContractsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ContractsSession struct {
	Contract     *Contracts        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ContractsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ContractsCallerSession struct {
	Contract *ContractsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// ContractsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ContractsTransactorSession struct {
	Contract     *ContractsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// ContractsRaw is an auto generated low-level Go binding around an Ethereum contract.
type ContractsRaw struct {
	Contract *Contracts // Generic contract binding to access the raw methods on
}

// ContractsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ContractsCallerRaw struct {
	Contract *ContractsCaller // Generic read-only contract binding to access the raw methods on
}

// ContractsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ContractsTransactorRaw struct {
	Contract *ContractsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewContracts creates a new instance of Contracts, bound to a specific deployed contract.
func NewContracts(address common.Address, backend bind.ContractBackend) (*Contracts, error) {
	contract, err := bindContracts(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Contracts{ContractsCaller: ContractsCaller{contract: contract}, ContractsTransactor: ContractsTransactor{contract: contract}, ContractsFilterer: ContractsFilterer{contract: contract}}, nil
}

// NewContractsCaller creates a new read-only instance of Contracts, bound to a specific deployed contract.
func NewContractsCaller(address common.Address, caller bind.ContractCaller) (*ContractsCaller, error) {
	contract, err := bindContracts(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ContractsCaller{contract: contract}, nil
}

// NewContractsTransactor creates a new write-only instance of Contracts, bound to a specific deployed contract.
func NewContractsTransactor(address common.Address, transactor bind.ContractTransactor) (*ContractsTransactor, error) {
	contract, err := bindContracts(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ContractsTransactor{contract: contract}, nil
}

// NewContractsFilterer creates a new log filterer instance of Contracts, bound to a specific deployed contract.
func NewContractsFilterer(address common.Address, filterer bind.ContractFilterer) (*ContractsFilterer, error) {
	contract, err := bindContracts(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ContractsFilterer{contract: contract}, nil
}

// bindContracts binds a generic wrapper to an already deployed contract.
func bindContracts(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ContractsMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Contracts *ContractsRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Contracts.Contract.ContractsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Contracts *ContractsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contracts.Contract.ContractsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Contracts *ContractsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Contracts.Contract.ContractsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Contracts *ContractsCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Contracts.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Contracts *ContractsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contracts.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Contracts *ContractsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Contracts.Contract.contract.Transact(opts, method, params...)
}

// BALLOTTYPEHASH is a free data retrieval call binding the contract method 0xdeaaa7cc.
//
// Solidity: function BALLOT_TYPEHASH() view returns(bytes32)
func (_Contracts *ContractsCaller) BALLOTTYPEHASH(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "BALLOT_TYPEHASH")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// BALLOTTYPEHASH is a free data retrieval call binding the contract method 0xdeaaa7cc.
//
// Solidity: function BALLOT_TYPEHASH() view returns(bytes32)
func (_Contracts *ContractsSession) BALLOTTYPEHASH() ([32]byte, error) {
	return _Contracts.Contract.BALLOTTYPEHASH(&_Contracts.CallOpts)
}

// BALLOTTYPEHASH is a free data retrieval call binding the contract method 0xdeaaa7cc.
//
// Solidity: function BALLOT_TYPEHASH() view returns(bytes32)
func (_Contracts *ContractsCallerSession) BALLOTTYPEHASH() ([32]byte, error) {
	return _Contracts.Contract.BALLOTTYPEHASH(&_Contracts.CallOpts)
}

// CLOCKMODE is a free data retrieval call binding the contract method 0x4bf5d7e9.
//
// Solidity: function CLOCK_MODE() view returns(string)
func (_Contracts *ContractsCaller) CLOCKMODE(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "CLOCK_MODE")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// CLOCKMODE is a free data retrieval call binding the contract method 0x4bf5d7e9.
//
// Solidity: function CLOCK_MODE() view returns(string)
func (_Contracts *ContractsSession) CLOCKMODE() (string, error) {
	return _Contracts.Contract.CLOCKMODE(&_Contracts.CallOpts)
}

// CLOCKMODE is a free data retrieval call binding the contract method 0x4bf5d7e9.
//
// Solidity: function CLOCK_MODE() view returns(string)
func (_Contracts *ContractsCallerSession) CLOCKMODE() (string, error) {
	return _Contracts.Contract.CLOCKMODE(&_Contracts.CallOpts)
}

// COUNTINGMODE is a free data retrieval call binding the contract method 0xdd4e2ba5.
//
// Solidity: function COUNTING_MODE() pure returns(string)
func (_Contracts *ContractsCaller) COUNTINGMODE(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "COUNTING_MODE")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// COUNTINGMODE is a free data retrieval call binding the contract method 0xdd4e2ba5.
//
// Solidity: function COUNTING_MODE() pure returns(string)
func (_Contracts *ContractsSession) COUNTINGMODE() (string, error) {
	return _Contracts.Contract.COUNTINGMODE(&_Contracts.CallOpts)
}

// COUNTINGMODE is a free data retrieval call binding the contract method 0xdd4e2ba5.
//
// Solidity: function COUNTING_MODE() pure returns(string)
func (_Contracts *ContractsCallerSession) COUNTINGMODE() (string, error) {
	return _Contracts.Contract.COUNTINGMODE(&_Contracts.CallOpts)
}

// EXTENDEDBALLOTTYPEHASH is a free data retrieval call binding the contract method 0x2fe3e261.
//
// Solidity: function EXTENDED_BALLOT_TYPEHASH() view returns(bytes32)
func (_Contracts *ContractsCaller) EXTENDEDBALLOTTYPEHASH(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "EXTENDED_BALLOT_TYPEHASH")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// EXTENDEDBALLOTTYPEHASH is a free data retrieval call binding the contract method 0x2fe3e261.
//
// Solidity: function EXTENDED_BALLOT_TYPEHASH() view returns(bytes32)
func (_Contracts *ContractsSession) EXTENDEDBALLOTTYPEHASH() ([32]byte, error) {
	return _Contracts.Contract.EXTENDEDBALLOTTYPEHASH(&_Contracts.CallOpts)
}

// EXTENDEDBALLOTTYPEHASH is a free data retrieval call binding the contract method 0x2fe3e261.
//
// Solidity: function EXTENDED_BALLOT_TYPEHASH() view returns(bytes32)
func (_Contracts *ContractsCallerSession) EXTENDEDBALLOTTYPEHASH() ([32]byte, error) {
	return _Contracts.Contract.EXTENDEDBALLOTTYPEHASH(&_Contracts.CallOpts)
}

// Clock is a free data retrieval call binding the contract method 0x91ddadf4.
//
// Solidity: function clock() view returns(uint48)
func (_Contracts *ContractsCaller) Clock(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "clock")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Clock is a free data retrieval call binding the contract method 0x91ddadf4.
//
// Solidity: function clock() view returns(uint48)
func (_Contracts *ContractsSession) Clock() (*big.Int, error) {
	return _Contracts.Contract.Clock(&_Contracts.CallOpts)
}

// Clock is a free data retrieval call binding the contract method 0x91ddadf4.
//
// Solidity: function clock() view returns(uint48)
func (_Contracts *ContractsCallerSession) Clock() (*big.Int, error) {
	return _Contracts.Contract.Clock(&_Contracts.CallOpts)
}

// Eip712Domain is a free data retrieval call binding the contract method 0x84b0196e.
//
// Solidity: function eip712Domain() view returns(bytes1 fields, string name, string version, uint256 chainId, address verifyingContract, bytes32 salt, uint256[] extensions)
func (_Contracts *ContractsCaller) Eip712Domain(opts *bind.CallOpts) (struct {
	Fields            [1]byte
	Name              string
	Version           string
	ChainId           *big.Int
	VerifyingContract common.Address
	Salt              [32]byte
	Extensions        []*big.Int
}, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "eip712Domain")

	outstruct := new(struct {
		Fields            [1]byte
		Name              string
		Version           string
		ChainId           *big.Int
		VerifyingContract common.Address
		Salt              [32]byte
		Extensions        []*big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Fields = *abi.ConvertType(out[0], new([1]byte)).(*[1]byte)
	outstruct.Name = *abi.ConvertType(out[1], new(string)).(*string)
	outstruct.Version = *abi.ConvertType(out[2], new(string)).(*string)
	outstruct.ChainId = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.VerifyingContract = *abi.ConvertType(out[4], new(common.Address)).(*common.Address)
	outstruct.Salt = *abi.ConvertType(out[5], new([32]byte)).(*[32]byte)
	outstruct.Extensions = *abi.ConvertType(out[6], new([]*big.Int)).(*[]*big.Int)

	return *outstruct, err

}

// Eip712Domain is a free data retrieval call binding the contract method 0x84b0196e.
//
// Solidity: function eip712Domain() view returns(bytes1 fields, string name, string version, uint256 chainId, address verifyingContract, bytes32 salt, uint256[] extensions)
func (_Contracts *ContractsSession) Eip712Domain() (struct {
	Fields            [1]byte
	Name              string
	Version           string
	ChainId           *big.Int
	VerifyingContract common.Address
	Salt              [32]byte
	Extensions        []*big.Int
}, error) {
	return _Contracts.Contract.Eip712Domain(&_Contracts.CallOpts)
}

// Eip712Domain is a free data retrieval call binding the contract method 0x84b0196e.
//
// Solidity: function eip712Domain() view returns(bytes1 fields, string name, string version, uint256 chainId, address verifyingContract, bytes32 salt, uint256[] extensions)
func (_Contracts *ContractsCallerSession) Eip712Domain() (struct {
	Fields            [1]byte
	Name              string
	Version           string
	ChainId           *big.Int
	VerifyingContract common.Address
	Salt              [32]byte
	Extensions        []*big.Int
}, error) {
	return _Contracts.Contract.Eip712Domain(&_Contracts.CallOpts)
}

// GetVotes is a free data retrieval call binding the contract method 0xeb9019d4.
//
// Solidity: function getVotes(address account, uint256 timepoint) view returns(uint256)
func (_Contracts *ContractsCaller) GetVotes(opts *bind.CallOpts, account common.Address, timepoint *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "getVotes", account, timepoint)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetVotes is a free data retrieval call binding the contract method 0xeb9019d4.
//
// Solidity: function getVotes(address account, uint256 timepoint) view returns(uint256)
func (_Contracts *ContractsSession) GetVotes(account common.Address, timepoint *big.Int) (*big.Int, error) {
	return _Contracts.Contract.GetVotes(&_Contracts.CallOpts, account, timepoint)
}

// GetVotes is a free data retrieval call binding the contract method 0xeb9019d4.
//
// Solidity: function getVotes(address account, uint256 timepoint) view returns(uint256)
func (_Contracts *ContractsCallerSession) GetVotes(account common.Address, timepoint *big.Int) (*big.Int, error) {
	return _Contracts.Contract.GetVotes(&_Contracts.CallOpts, account, timepoint)
}

// GetVotesWithParams is a free data retrieval call binding the contract method 0x9a802a6d.
//
// Solidity: function getVotesWithParams(address account, uint256 timepoint, bytes params) view returns(uint256)
func (_Contracts *ContractsCaller) GetVotesWithParams(opts *bind.CallOpts, account common.Address, timepoint *big.Int, params []byte) (*big.Int, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "getVotesWithParams", account, timepoint, params)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetVotesWithParams is a free data retrieval call binding the contract method 0x9a802a6d.
//
// Solidity: function getVotesWithParams(address account, uint256 timepoint, bytes params) view returns(uint256)
func (_Contracts *ContractsSession) GetVotesWithParams(account common.Address, timepoint *big.Int, params []byte) (*big.Int, error) {
	return _Contracts.Contract.GetVotesWithParams(&_Contracts.CallOpts, account, timepoint, params)
}

// GetVotesWithParams is a free data retrieval call binding the contract method 0x9a802a6d.
//
// Solidity: function getVotesWithParams(address account, uint256 timepoint, bytes params) view returns(uint256)
func (_Contracts *ContractsCallerSession) GetVotesWithParams(account common.Address, timepoint *big.Int, params []byte) (*big.Int, error) {
	return _Contracts.Contract.GetVotesWithParams(&_Contracts.CallOpts, account, timepoint, params)
}

// HasVoted is a free data retrieval call binding the contract method 0x43859632.
//
// Solidity: function hasVoted(uint256 proposalId, address account) view returns(bool)
func (_Contracts *ContractsCaller) HasVoted(opts *bind.CallOpts, proposalId *big.Int, account common.Address) (bool, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "hasVoted", proposalId, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasVoted is a free data retrieval call binding the contract method 0x43859632.
//
// Solidity: function hasVoted(uint256 proposalId, address account) view returns(bool)
func (_Contracts *ContractsSession) HasVoted(proposalId *big.Int, account common.Address) (bool, error) {
	return _Contracts.Contract.HasVoted(&_Contracts.CallOpts, proposalId, account)
}

// HasVoted is a free data retrieval call binding the contract method 0x43859632.
//
// Solidity: function hasVoted(uint256 proposalId, address account) view returns(bool)
func (_Contracts *ContractsCallerSession) HasVoted(proposalId *big.Int, account common.Address) (bool, error) {
	return _Contracts.Contract.HasVoted(&_Contracts.CallOpts, proposalId, account)
}

// HashProposal is a free data retrieval call binding the contract method 0xc59057e4.
//
// Solidity: function hashProposal(address[] targets, uint256[] values, bytes[] calldatas, bytes32 descriptionHash) pure returns(uint256)
func (_Contracts *ContractsCaller) HashProposal(opts *bind.CallOpts, targets []common.Address, values []*big.Int, calldatas [][]byte, descriptionHash [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "hashProposal", targets, values, calldatas, descriptionHash)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// HashProposal is a free data retrieval call binding the contract method 0xc59057e4.
//
// Solidity: function hashProposal(address[] targets, uint256[] values, bytes[] calldatas, bytes32 descriptionHash) pure returns(uint256)
func (_Contracts *ContractsSession) HashProposal(targets []common.Address, values []*big.Int, calldatas [][]byte, descriptionHash [32]byte) (*big.Int, error) {
	return _Contracts.Contract.HashProposal(&_Contracts.CallOpts, targets, values, calldatas, descriptionHash)
}

// HashProposal is a free data retrieval call binding the contract method 0xc59057e4.
//
// Solidity: function hashProposal(address[] targets, uint256[] values, bytes[] calldatas, bytes32 descriptionHash) pure returns(uint256)
func (_Contracts *ContractsCallerSession) HashProposal(targets []common.Address, values []*big.Int, calldatas [][]byte, descriptionHash [32]byte) (*big.Int, error) {
	return _Contracts.Contract.HashProposal(&_Contracts.CallOpts, targets, values, calldatas, descriptionHash)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Contracts *ContractsCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Contracts *ContractsSession) Name() (string, error) {
	return _Contracts.Contract.Name(&_Contracts.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Contracts *ContractsCallerSession) Name() (string, error) {
	return _Contracts.Contract.Name(&_Contracts.CallOpts)
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address owner) view returns(uint256)
func (_Contracts *ContractsCaller) Nonces(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "nonces", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address owner) view returns(uint256)
func (_Contracts *ContractsSession) Nonces(owner common.Address) (*big.Int, error) {
	return _Contracts.Contract.Nonces(&_Contracts.CallOpts, owner)
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address owner) view returns(uint256)
func (_Contracts *ContractsCallerSession) Nonces(owner common.Address) (*big.Int, error) {
	return _Contracts.Contract.Nonces(&_Contracts.CallOpts, owner)
}

// ProposalDeadline is a free data retrieval call binding the contract method 0xc01f9e37.
//
// Solidity: function proposalDeadline(uint256 proposalId) view returns(uint256)
func (_Contracts *ContractsCaller) ProposalDeadline(opts *bind.CallOpts, proposalId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "proposalDeadline", proposalId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ProposalDeadline is a free data retrieval call binding the contract method 0xc01f9e37.
//
// Solidity: function proposalDeadline(uint256 proposalId) view returns(uint256)
func (_Contracts *ContractsSession) ProposalDeadline(proposalId *big.Int) (*big.Int, error) {
	return _Contracts.Contract.ProposalDeadline(&_Contracts.CallOpts, proposalId)
}

// ProposalDeadline is a free data retrieval call binding the contract method 0xc01f9e37.
//
// Solidity: function proposalDeadline(uint256 proposalId) view returns(uint256)
func (_Contracts *ContractsCallerSession) ProposalDeadline(proposalId *big.Int) (*big.Int, error) {
	return _Contracts.Contract.ProposalDeadline(&_Contracts.CallOpts, proposalId)
}

// ProposalEta is a free data retrieval call binding the contract method 0xab58fb8e.
//
// Solidity: function proposalEta(uint256 proposalId) view returns(uint256)
func (_Contracts *ContractsCaller) ProposalEta(opts *bind.CallOpts, proposalId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "proposalEta", proposalId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ProposalEta is a free data retrieval call binding the contract method 0xab58fb8e.
//
// Solidity: function proposalEta(uint256 proposalId) view returns(uint256)
func (_Contracts *ContractsSession) ProposalEta(proposalId *big.Int) (*big.Int, error) {
	return _Contracts.Contract.ProposalEta(&_Contracts.CallOpts, proposalId)
}

// ProposalEta is a free data retrieval call binding the contract method 0xab58fb8e.
//
// Solidity: function proposalEta(uint256 proposalId) view returns(uint256)
func (_Contracts *ContractsCallerSession) ProposalEta(proposalId *big.Int) (*big.Int, error) {
	return _Contracts.Contract.ProposalEta(&_Contracts.CallOpts, proposalId)
}

// ProposalNeedsQueuing is a free data retrieval call binding the contract method 0xa9a95294.
//
// Solidity: function proposalNeedsQueuing(uint256 ) view returns(bool)
func (_Contracts *ContractsCaller) ProposalNeedsQueuing(opts *bind.CallOpts, arg0 *big.Int) (bool, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "proposalNeedsQueuing", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// ProposalNeedsQueuing is a free data retrieval call binding the contract method 0xa9a95294.
//
// Solidity: function proposalNeedsQueuing(uint256 ) view returns(bool)
func (_Contracts *ContractsSession) ProposalNeedsQueuing(arg0 *big.Int) (bool, error) {
	return _Contracts.Contract.ProposalNeedsQueuing(&_Contracts.CallOpts, arg0)
}

// ProposalNeedsQueuing is a free data retrieval call binding the contract method 0xa9a95294.
//
// Solidity: function proposalNeedsQueuing(uint256 ) view returns(bool)
func (_Contracts *ContractsCallerSession) ProposalNeedsQueuing(arg0 *big.Int) (bool, error) {
	return _Contracts.Contract.ProposalNeedsQueuing(&_Contracts.CallOpts, arg0)
}

// ProposalProposer is a free data retrieval call binding the contract method 0x143489d0.
//
// Solidity: function proposalProposer(uint256 proposalId) view returns(address)
func (_Contracts *ContractsCaller) ProposalProposer(opts *bind.CallOpts, proposalId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "proposalProposer", proposalId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ProposalProposer is a free data retrieval call binding the contract method 0x143489d0.
//
// Solidity: function proposalProposer(uint256 proposalId) view returns(address)
func (_Contracts *ContractsSession) ProposalProposer(proposalId *big.Int) (common.Address, error) {
	return _Contracts.Contract.ProposalProposer(&_Contracts.CallOpts, proposalId)
}

// ProposalProposer is a free data retrieval call binding the contract method 0x143489d0.
//
// Solidity: function proposalProposer(uint256 proposalId) view returns(address)
func (_Contracts *ContractsCallerSession) ProposalProposer(proposalId *big.Int) (common.Address, error) {
	return _Contracts.Contract.ProposalProposer(&_Contracts.CallOpts, proposalId)
}

// ProposalSnapshot is a free data retrieval call binding the contract method 0x2d63f693.
//
// Solidity: function proposalSnapshot(uint256 proposalId) view returns(uint256)
func (_Contracts *ContractsCaller) ProposalSnapshot(opts *bind.CallOpts, proposalId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "proposalSnapshot", proposalId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ProposalSnapshot is a free data retrieval call binding the contract method 0x2d63f693.
//
// Solidity: function proposalSnapshot(uint256 proposalId) view returns(uint256)
func (_Contracts *ContractsSession) ProposalSnapshot(proposalId *big.Int) (*big.Int, error) {
	return _Contracts.Contract.ProposalSnapshot(&_Contracts.CallOpts, proposalId)
}

// ProposalSnapshot is a free data retrieval call binding the contract method 0x2d63f693.
//
// Solidity: function proposalSnapshot(uint256 proposalId) view returns(uint256)
func (_Contracts *ContractsCallerSession) ProposalSnapshot(proposalId *big.Int) (*big.Int, error) {
	return _Contracts.Contract.ProposalSnapshot(&_Contracts.CallOpts, proposalId)
}

// ProposalThreshold is a free data retrieval call binding the contract method 0xb58131b0.
//
// Solidity: function proposalThreshold() pure returns(uint256)
func (_Contracts *ContractsCaller) ProposalThreshold(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "proposalThreshold")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ProposalThreshold is a free data retrieval call binding the contract method 0xb58131b0.
//
// Solidity: function proposalThreshold() pure returns(uint256)
func (_Contracts *ContractsSession) ProposalThreshold() (*big.Int, error) {
	return _Contracts.Contract.ProposalThreshold(&_Contracts.CallOpts)
}

// ProposalThreshold is a free data retrieval call binding the contract method 0xb58131b0.
//
// Solidity: function proposalThreshold() pure returns(uint256)
func (_Contracts *ContractsCallerSession) ProposalThreshold() (*big.Int, error) {
	return _Contracts.Contract.ProposalThreshold(&_Contracts.CallOpts)
}

// ProposalVotes is a free data retrieval call binding the contract method 0x544ffc9c.
//
// Solidity: function proposalVotes(uint256 proposalId) view returns(uint256 againstVotes, uint256 forVotes, uint256 abstainVotes)
func (_Contracts *ContractsCaller) ProposalVotes(opts *bind.CallOpts, proposalId *big.Int) (struct {
	AgainstVotes *big.Int
	ForVotes     *big.Int
	AbstainVotes *big.Int
}, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "proposalVotes", proposalId)

	outstruct := new(struct {
		AgainstVotes *big.Int
		ForVotes     *big.Int
		AbstainVotes *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.AgainstVotes = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.ForVotes = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.AbstainVotes = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// ProposalVotes is a free data retrieval call binding the contract method 0x544ffc9c.
//
// Solidity: function proposalVotes(uint256 proposalId) view returns(uint256 againstVotes, uint256 forVotes, uint256 abstainVotes)
func (_Contracts *ContractsSession) ProposalVotes(proposalId *big.Int) (struct {
	AgainstVotes *big.Int
	ForVotes     *big.Int
	AbstainVotes *big.Int
}, error) {
	return _Contracts.Contract.ProposalVotes(&_Contracts.CallOpts, proposalId)
}

// ProposalVotes is a free data retrieval call binding the contract method 0x544ffc9c.
//
// Solidity: function proposalVotes(uint256 proposalId) view returns(uint256 againstVotes, uint256 forVotes, uint256 abstainVotes)
func (_Contracts *ContractsCallerSession) ProposalVotes(proposalId *big.Int) (struct {
	AgainstVotes *big.Int
	ForVotes     *big.Int
	AbstainVotes *big.Int
}, error) {
	return _Contracts.Contract.ProposalVotes(&_Contracts.CallOpts, proposalId)
}

// Quorum is a free data retrieval call binding the contract method 0xf8ce560a.
//
// Solidity: function quorum(uint256 timepoint) view returns(uint256)
func (_Contracts *ContractsCaller) Quorum(opts *bind.CallOpts, timepoint *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "quorum", timepoint)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Quorum is a free data retrieval call binding the contract method 0xf8ce560a.
//
// Solidity: function quorum(uint256 timepoint) view returns(uint256)
func (_Contracts *ContractsSession) Quorum(timepoint *big.Int) (*big.Int, error) {
	return _Contracts.Contract.Quorum(&_Contracts.CallOpts, timepoint)
}

// Quorum is a free data retrieval call binding the contract method 0xf8ce560a.
//
// Solidity: function quorum(uint256 timepoint) view returns(uint256)
func (_Contracts *ContractsCallerSession) Quorum(timepoint *big.Int) (*big.Int, error) {
	return _Contracts.Contract.Quorum(&_Contracts.CallOpts, timepoint)
}

// QuorumDenominator is a free data retrieval call binding the contract method 0x97c3d334.
//
// Solidity: function quorumDenominator() view returns(uint256)
func (_Contracts *ContractsCaller) QuorumDenominator(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "quorumDenominator")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// QuorumDenominator is a free data retrieval call binding the contract method 0x97c3d334.
//
// Solidity: function quorumDenominator() view returns(uint256)
func (_Contracts *ContractsSession) QuorumDenominator() (*big.Int, error) {
	return _Contracts.Contract.QuorumDenominator(&_Contracts.CallOpts)
}

// QuorumDenominator is a free data retrieval call binding the contract method 0x97c3d334.
//
// Solidity: function quorumDenominator() view returns(uint256)
func (_Contracts *ContractsCallerSession) QuorumDenominator() (*big.Int, error) {
	return _Contracts.Contract.QuorumDenominator(&_Contracts.CallOpts)
}

// QuorumNumerator is a free data retrieval call binding the contract method 0x60c4247f.
//
// Solidity: function quorumNumerator(uint256 timepoint) view returns(uint256)
func (_Contracts *ContractsCaller) QuorumNumerator(opts *bind.CallOpts, timepoint *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "quorumNumerator", timepoint)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// QuorumNumerator is a free data retrieval call binding the contract method 0x60c4247f.
//
// Solidity: function quorumNumerator(uint256 timepoint) view returns(uint256)
func (_Contracts *ContractsSession) QuorumNumerator(timepoint *big.Int) (*big.Int, error) {
	return _Contracts.Contract.QuorumNumerator(&_Contracts.CallOpts, timepoint)
}

// QuorumNumerator is a free data retrieval call binding the contract method 0x60c4247f.
//
// Solidity: function quorumNumerator(uint256 timepoint) view returns(uint256)
func (_Contracts *ContractsCallerSession) QuorumNumerator(timepoint *big.Int) (*big.Int, error) {
	return _Contracts.Contract.QuorumNumerator(&_Contracts.CallOpts, timepoint)
}

// QuorumNumerator0 is a free data retrieval call binding the contract method 0xa7713a70.
//
// Solidity: function quorumNumerator() view returns(uint256)
func (_Contracts *ContractsCaller) QuorumNumerator0(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "quorumNumerator0")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// QuorumNumerator0 is a free data retrieval call binding the contract method 0xa7713a70.
//
// Solidity: function quorumNumerator() view returns(uint256)
func (_Contracts *ContractsSession) QuorumNumerator0() (*big.Int, error) {
	return _Contracts.Contract.QuorumNumerator0(&_Contracts.CallOpts)
}

// QuorumNumerator0 is a free data retrieval call binding the contract method 0xa7713a70.
//
// Solidity: function quorumNumerator() view returns(uint256)
func (_Contracts *ContractsCallerSession) QuorumNumerator0() (*big.Int, error) {
	return _Contracts.Contract.QuorumNumerator0(&_Contracts.CallOpts)
}

// State is a free data retrieval call binding the contract method 0x3e4f49e6.
//
// Solidity: function state(uint256 proposalId) view returns(uint8)
func (_Contracts *ContractsCaller) State(opts *bind.CallOpts, proposalId *big.Int) (uint8, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "state", proposalId)

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// State is a free data retrieval call binding the contract method 0x3e4f49e6.
//
// Solidity: function state(uint256 proposalId) view returns(uint8)
func (_Contracts *ContractsSession) State(proposalId *big.Int) (uint8, error) {
	return _Contracts.Contract.State(&_Contracts.CallOpts, proposalId)
}

// State is a free data retrieval call binding the contract method 0x3e4f49e6.
//
// Solidity: function state(uint256 proposalId) view returns(uint8)
func (_Contracts *ContractsCallerSession) State(proposalId *big.Int) (uint8, error) {
	return _Contracts.Contract.State(&_Contracts.CallOpts, proposalId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Contracts *ContractsCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Contracts *ContractsSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Contracts.Contract.SupportsInterface(&_Contracts.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Contracts *ContractsCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Contracts.Contract.SupportsInterface(&_Contracts.CallOpts, interfaceId)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_Contracts *ContractsCaller) Token(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "token")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_Contracts *ContractsSession) Token() (common.Address, error) {
	return _Contracts.Contract.Token(&_Contracts.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_Contracts *ContractsCallerSession) Token() (common.Address, error) {
	return _Contracts.Contract.Token(&_Contracts.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_Contracts *ContractsCaller) Version(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "version")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_Contracts *ContractsSession) Version() (string, error) {
	return _Contracts.Contract.Version(&_Contracts.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_Contracts *ContractsCallerSession) Version() (string, error) {
	return _Contracts.Contract.Version(&_Contracts.CallOpts)
}

// VotingDelay is a free data retrieval call binding the contract method 0x3932abb1.
//
// Solidity: function votingDelay() pure returns(uint256)
func (_Contracts *ContractsCaller) VotingDelay(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "votingDelay")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// VotingDelay is a free data retrieval call binding the contract method 0x3932abb1.
//
// Solidity: function votingDelay() pure returns(uint256)
func (_Contracts *ContractsSession) VotingDelay() (*big.Int, error) {
	return _Contracts.Contract.VotingDelay(&_Contracts.CallOpts)
}

// VotingDelay is a free data retrieval call binding the contract method 0x3932abb1.
//
// Solidity: function votingDelay() pure returns(uint256)
func (_Contracts *ContractsCallerSession) VotingDelay() (*big.Int, error) {
	return _Contracts.Contract.VotingDelay(&_Contracts.CallOpts)
}

// VotingPeriod is a free data retrieval call binding the contract method 0x02a251a3.
//
// Solidity: function votingPeriod() pure returns(uint256)
func (_Contracts *ContractsCaller) VotingPeriod(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "votingPeriod")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// VotingPeriod is a free data retrieval call binding the contract method 0x02a251a3.
//
// Solidity: function votingPeriod() pure returns(uint256)
func (_Contracts *ContractsSession) VotingPeriod() (*big.Int, error) {
	return _Contracts.Contract.VotingPeriod(&_Contracts.CallOpts)
}

// VotingPeriod is a free data retrieval call binding the contract method 0x02a251a3.
//
// Solidity: function votingPeriod() pure returns(uint256)
func (_Contracts *ContractsCallerSession) VotingPeriod() (*big.Int, error) {
	return _Contracts.Contract.VotingPeriod(&_Contracts.CallOpts)
}

// Cancel is a paid mutator transaction binding the contract method 0x452115d6.
//
// Solidity: function cancel(address[] targets, uint256[] values, bytes[] calldatas, bytes32 descriptionHash) returns(uint256)
func (_Contracts *ContractsTransactor) Cancel(opts *bind.TransactOpts, targets []common.Address, values []*big.Int, calldatas [][]byte, descriptionHash [32]byte) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "cancel", targets, values, calldatas, descriptionHash)
}

// Cancel is a paid mutator transaction binding the contract method 0x452115d6.
//
// Solidity: function cancel(address[] targets, uint256[] values, bytes[] calldatas, bytes32 descriptionHash) returns(uint256)
func (_Contracts *ContractsSession) Cancel(targets []common.Address, values []*big.Int, calldatas [][]byte, descriptionHash [32]byte) (*types.Transaction, error) {
	return _Contracts.Contract.Cancel(&_Contracts.TransactOpts, targets, values, calldatas, descriptionHash)
}

// Cancel is a paid mutator transaction binding the contract method 0x452115d6.
//
// Solidity: function cancel(address[] targets, uint256[] values, bytes[] calldatas, bytes32 descriptionHash) returns(uint256)
func (_Contracts *ContractsTransactorSession) Cancel(targets []common.Address, values []*big.Int, calldatas [][]byte, descriptionHash [32]byte) (*types.Transaction, error) {
	return _Contracts.Contract.Cancel(&_Contracts.TransactOpts, targets, values, calldatas, descriptionHash)
}

// CastVote is a paid mutator transaction binding the contract method 0x56781388.
//
// Solidity: function castVote(uint256 proposalId, uint8 support) returns(uint256)
func (_Contracts *ContractsTransactor) CastVote(opts *bind.TransactOpts, proposalId *big.Int, support uint8) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "castVote", proposalId, support)
}

// CastVote is a paid mutator transaction binding the contract method 0x56781388.
//
// Solidity: function castVote(uint256 proposalId, uint8 support) returns(uint256)
func (_Contracts *ContractsSession) CastVote(proposalId *big.Int, support uint8) (*types.Transaction, error) {
	return _Contracts.Contract.CastVote(&_Contracts.TransactOpts, proposalId, support)
}

// CastVote is a paid mutator transaction binding the contract method 0x56781388.
//
// Solidity: function castVote(uint256 proposalId, uint8 support) returns(uint256)
func (_Contracts *ContractsTransactorSession) CastVote(proposalId *big.Int, support uint8) (*types.Transaction, error) {
	return _Contracts.Contract.CastVote(&_Contracts.TransactOpts, proposalId, support)
}

// CastVoteBySig is a paid mutator transaction binding the contract method 0x8ff262e3.
//
// Solidity: function castVoteBySig(uint256 proposalId, uint8 support, address voter, bytes signature) returns(uint256)
func (_Contracts *ContractsTransactor) CastVoteBySig(opts *bind.TransactOpts, proposalId *big.Int, support uint8, voter common.Address, signature []byte) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "castVoteBySig", proposalId, support, voter, signature)
}

// CastVoteBySig is a paid mutator transaction binding the contract method 0x8ff262e3.
//
// Solidity: function castVoteBySig(uint256 proposalId, uint8 support, address voter, bytes signature) returns(uint256)
func (_Contracts *ContractsSession) CastVoteBySig(proposalId *big.Int, support uint8, voter common.Address, signature []byte) (*types.Transaction, error) {
	return _Contracts.Contract.CastVoteBySig(&_Contracts.TransactOpts, proposalId, support, voter, signature)
}

// CastVoteBySig is a paid mutator transaction binding the contract method 0x8ff262e3.
//
// Solidity: function castVoteBySig(uint256 proposalId, uint8 support, address voter, bytes signature) returns(uint256)
func (_Contracts *ContractsTransactorSession) CastVoteBySig(proposalId *big.Int, support uint8, voter common.Address, signature []byte) (*types.Transaction, error) {
	return _Contracts.Contract.CastVoteBySig(&_Contracts.TransactOpts, proposalId, support, voter, signature)
}

// CastVoteWithReason is a paid mutator transaction binding the contract method 0x7b3c71d3.
//
// Solidity: function castVoteWithReason(uint256 proposalId, uint8 support, string reason) returns(uint256)
func (_Contracts *ContractsTransactor) CastVoteWithReason(opts *bind.TransactOpts, proposalId *big.Int, support uint8, reason string) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "castVoteWithReason", proposalId, support, reason)
}

// CastVoteWithReason is a paid mutator transaction binding the contract method 0x7b3c71d3.
//
// Solidity: function castVoteWithReason(uint256 proposalId, uint8 support, string reason) returns(uint256)
func (_Contracts *ContractsSession) CastVoteWithReason(proposalId *big.Int, support uint8, reason string) (*types.Transaction, error) {
	return _Contracts.Contract.CastVoteWithReason(&_Contracts.TransactOpts, proposalId, support, reason)
}

// CastVoteWithReason is a paid mutator transaction binding the contract method 0x7b3c71d3.
//
// Solidity: function castVoteWithReason(uint256 proposalId, uint8 support, string reason) returns(uint256)
func (_Contracts *ContractsTransactorSession) CastVoteWithReason(proposalId *big.Int, support uint8, reason string) (*types.Transaction, error) {
	return _Contracts.Contract.CastVoteWithReason(&_Contracts.TransactOpts, proposalId, support, reason)
}

// CastVoteWithReasonAndParams is a paid mutator transaction binding the contract method 0x5f398a14.
//
// Solidity: function castVoteWithReasonAndParams(uint256 proposalId, uint8 support, string reason, bytes params) returns(uint256)
func (_Contracts *ContractsTransactor) CastVoteWithReasonAndParams(opts *bind.TransactOpts, proposalId *big.Int, support uint8, reason string, params []byte) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "castVoteWithReasonAndParams", proposalId, support, reason, params)
}

// CastVoteWithReasonAndParams is a paid mutator transaction binding the contract method 0x5f398a14.
//
// Solidity: function castVoteWithReasonAndParams(uint256 proposalId, uint8 support, string reason, bytes params) returns(uint256)
func (_Contracts *ContractsSession) CastVoteWithReasonAndParams(proposalId *big.Int, support uint8, reason string, params []byte) (*types.Transaction, error) {
	return _Contracts.Contract.CastVoteWithReasonAndParams(&_Contracts.TransactOpts, proposalId, support, reason, params)
}

// CastVoteWithReasonAndParams is a paid mutator transaction binding the contract method 0x5f398a14.
//
// Solidity: function castVoteWithReasonAndParams(uint256 proposalId, uint8 support, string reason, bytes params) returns(uint256)
func (_Contracts *ContractsTransactorSession) CastVoteWithReasonAndParams(proposalId *big.Int, support uint8, reason string, params []byte) (*types.Transaction, error) {
	return _Contracts.Contract.CastVoteWithReasonAndParams(&_Contracts.TransactOpts, proposalId, support, reason, params)
}

// CastVoteWithReasonAndParamsBySig is a paid mutator transaction binding the contract method 0x5b8d0e0d.
//
// Solidity: function castVoteWithReasonAndParamsBySig(uint256 proposalId, uint8 support, address voter, string reason, bytes params, bytes signature) returns(uint256)
func (_Contracts *ContractsTransactor) CastVoteWithReasonAndParamsBySig(opts *bind.TransactOpts, proposalId *big.Int, support uint8, voter common.Address, reason string, params []byte, signature []byte) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "castVoteWithReasonAndParamsBySig", proposalId, support, voter, reason, params, signature)
}

// CastVoteWithReasonAndParamsBySig is a paid mutator transaction binding the contract method 0x5b8d0e0d.
//
// Solidity: function castVoteWithReasonAndParamsBySig(uint256 proposalId, uint8 support, address voter, string reason, bytes params, bytes signature) returns(uint256)
func (_Contracts *ContractsSession) CastVoteWithReasonAndParamsBySig(proposalId *big.Int, support uint8, voter common.Address, reason string, params []byte, signature []byte) (*types.Transaction, error) {
	return _Contracts.Contract.CastVoteWithReasonAndParamsBySig(&_Contracts.TransactOpts, proposalId, support, voter, reason, params, signature)
}

// CastVoteWithReasonAndParamsBySig is a paid mutator transaction binding the contract method 0x5b8d0e0d.
//
// Solidity: function castVoteWithReasonAndParamsBySig(uint256 proposalId, uint8 support, address voter, string reason, bytes params, bytes signature) returns(uint256)
func (_Contracts *ContractsTransactorSession) CastVoteWithReasonAndParamsBySig(proposalId *big.Int, support uint8, voter common.Address, reason string, params []byte, signature []byte) (*types.Transaction, error) {
	return _Contracts.Contract.CastVoteWithReasonAndParamsBySig(&_Contracts.TransactOpts, proposalId, support, voter, reason, params, signature)
}

// Execute is a paid mutator transaction binding the contract method 0x2656227d.
//
// Solidity: function execute(address[] targets, uint256[] values, bytes[] calldatas, bytes32 descriptionHash) payable returns(uint256)
func (_Contracts *ContractsTransactor) Execute(opts *bind.TransactOpts, targets []common.Address, values []*big.Int, calldatas [][]byte, descriptionHash [32]byte) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "execute", targets, values, calldatas, descriptionHash)
}

// Execute is a paid mutator transaction binding the contract method 0x2656227d.
//
// Solidity: function execute(address[] targets, uint256[] values, bytes[] calldatas, bytes32 descriptionHash) payable returns(uint256)
func (_Contracts *ContractsSession) Execute(targets []common.Address, values []*big.Int, calldatas [][]byte, descriptionHash [32]byte) (*types.Transaction, error) {
	return _Contracts.Contract.Execute(&_Contracts.TransactOpts, targets, values, calldatas, descriptionHash)
}

// Execute is a paid mutator transaction binding the contract method 0x2656227d.
//
// Solidity: function execute(address[] targets, uint256[] values, bytes[] calldatas, bytes32 descriptionHash) payable returns(uint256)
func (_Contracts *ContractsTransactorSession) Execute(targets []common.Address, values []*big.Int, calldatas [][]byte, descriptionHash [32]byte) (*types.Transaction, error) {
	return _Contracts.Contract.Execute(&_Contracts.TransactOpts, targets, values, calldatas, descriptionHash)
}

// OnERC1155BatchReceived is a paid mutator transaction binding the contract method 0xbc197c81.
//
// Solidity: function onERC1155BatchReceived(address , address , uint256[] , uint256[] , bytes ) returns(bytes4)
func (_Contracts *ContractsTransactor) OnERC1155BatchReceived(opts *bind.TransactOpts, arg0 common.Address, arg1 common.Address, arg2 []*big.Int, arg3 []*big.Int, arg4 []byte) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "onERC1155BatchReceived", arg0, arg1, arg2, arg3, arg4)
}

// OnERC1155BatchReceived is a paid mutator transaction binding the contract method 0xbc197c81.
//
// Solidity: function onERC1155BatchReceived(address , address , uint256[] , uint256[] , bytes ) returns(bytes4)
func (_Contracts *ContractsSession) OnERC1155BatchReceived(arg0 common.Address, arg1 common.Address, arg2 []*big.Int, arg3 []*big.Int, arg4 []byte) (*types.Transaction, error) {
	return _Contracts.Contract.OnERC1155BatchReceived(&_Contracts.TransactOpts, arg0, arg1, arg2, arg3, arg4)
}

// OnERC1155BatchReceived is a paid mutator transaction binding the contract method 0xbc197c81.
//
// Solidity: function onERC1155BatchReceived(address , address , uint256[] , uint256[] , bytes ) returns(bytes4)
func (_Contracts *ContractsTransactorSession) OnERC1155BatchReceived(arg0 common.Address, arg1 common.Address, arg2 []*big.Int, arg3 []*big.Int, arg4 []byte) (*types.Transaction, error) {
	return _Contracts.Contract.OnERC1155BatchReceived(&_Contracts.TransactOpts, arg0, arg1, arg2, arg3, arg4)
}

// OnERC1155Received is a paid mutator transaction binding the contract method 0xf23a6e61.
//
// Solidity: function onERC1155Received(address , address , uint256 , uint256 , bytes ) returns(bytes4)
func (_Contracts *ContractsTransactor) OnERC1155Received(opts *bind.TransactOpts, arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 *big.Int, arg4 []byte) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "onERC1155Received", arg0, arg1, arg2, arg3, arg4)
}

// OnERC1155Received is a paid mutator transaction binding the contract method 0xf23a6e61.
//
// Solidity: function onERC1155Received(address , address , uint256 , uint256 , bytes ) returns(bytes4)
func (_Contracts *ContractsSession) OnERC1155Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 *big.Int, arg4 []byte) (*types.Transaction, error) {
	return _Contracts.Contract.OnERC1155Received(&_Contracts.TransactOpts, arg0, arg1, arg2, arg3, arg4)
}

// OnERC1155Received is a paid mutator transaction binding the contract method 0xf23a6e61.
//
// Solidity: function onERC1155Received(address , address , uint256 , uint256 , bytes ) returns(bytes4)
func (_Contracts *ContractsTransactorSession) OnERC1155Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 *big.Int, arg4 []byte) (*types.Transaction, error) {
	return _Contracts.Contract.OnERC1155Received(&_Contracts.TransactOpts, arg0, arg1, arg2, arg3, arg4)
}

// OnERC721Received is a paid mutator transaction binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address , address , uint256 , bytes ) returns(bytes4)
func (_Contracts *ContractsTransactor) OnERC721Received(opts *bind.TransactOpts, arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 []byte) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "onERC721Received", arg0, arg1, arg2, arg3)
}

// OnERC721Received is a paid mutator transaction binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address , address , uint256 , bytes ) returns(bytes4)
func (_Contracts *ContractsSession) OnERC721Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 []byte) (*types.Transaction, error) {
	return _Contracts.Contract.OnERC721Received(&_Contracts.TransactOpts, arg0, arg1, arg2, arg3)
}

// OnERC721Received is a paid mutator transaction binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address , address , uint256 , bytes ) returns(bytes4)
func (_Contracts *ContractsTransactorSession) OnERC721Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 []byte) (*types.Transaction, error) {
	return _Contracts.Contract.OnERC721Received(&_Contracts.TransactOpts, arg0, arg1, arg2, arg3)
}

// Propose is a paid mutator transaction binding the contract method 0x7d5e81e2.
//
// Solidity: function propose(address[] targets, uint256[] values, bytes[] calldatas, string description) returns(uint256)
func (_Contracts *ContractsTransactor) Propose(opts *bind.TransactOpts, targets []common.Address, values []*big.Int, calldatas [][]byte, description string) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "propose", targets, values, calldatas, description)
}

// Propose is a paid mutator transaction binding the contract method 0x7d5e81e2.
//
// Solidity: function propose(address[] targets, uint256[] values, bytes[] calldatas, string description) returns(uint256)
func (_Contracts *ContractsSession) Propose(targets []common.Address, values []*big.Int, calldatas [][]byte, description string) (*types.Transaction, error) {
	return _Contracts.Contract.Propose(&_Contracts.TransactOpts, targets, values, calldatas, description)
}

// Propose is a paid mutator transaction binding the contract method 0x7d5e81e2.
//
// Solidity: function propose(address[] targets, uint256[] values, bytes[] calldatas, string description) returns(uint256)
func (_Contracts *ContractsTransactorSession) Propose(targets []common.Address, values []*big.Int, calldatas [][]byte, description string) (*types.Transaction, error) {
	return _Contracts.Contract.Propose(&_Contracts.TransactOpts, targets, values, calldatas, description)
}

// ProposeWithSnapshotBlockNumber is a paid mutator transaction binding the contract method 0x65218612.
//
// Solidity: function proposeWithSnapshotBlockNumber(address[] targets, uint256[] values, bytes[] calldatas, string description, uint48 _snapshotBlockNumber) returns(uint256)
func (_Contracts *ContractsTransactor) ProposeWithSnapshotBlockNumber(opts *bind.TransactOpts, targets []common.Address, values []*big.Int, calldatas [][]byte, description string, _snapshotBlockNumber *big.Int) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "proposeWithSnapshotBlockNumber", targets, values, calldatas, description, _snapshotBlockNumber)
}

// ProposeWithSnapshotBlockNumber is a paid mutator transaction binding the contract method 0x65218612.
//
// Solidity: function proposeWithSnapshotBlockNumber(address[] targets, uint256[] values, bytes[] calldatas, string description, uint48 _snapshotBlockNumber) returns(uint256)
func (_Contracts *ContractsSession) ProposeWithSnapshotBlockNumber(targets []common.Address, values []*big.Int, calldatas [][]byte, description string, _snapshotBlockNumber *big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.ProposeWithSnapshotBlockNumber(&_Contracts.TransactOpts, targets, values, calldatas, description, _snapshotBlockNumber)
}

// ProposeWithSnapshotBlockNumber is a paid mutator transaction binding the contract method 0x65218612.
//
// Solidity: function proposeWithSnapshotBlockNumber(address[] targets, uint256[] values, bytes[] calldatas, string description, uint48 _snapshotBlockNumber) returns(uint256)
func (_Contracts *ContractsTransactorSession) ProposeWithSnapshotBlockNumber(targets []common.Address, values []*big.Int, calldatas [][]byte, description string, _snapshotBlockNumber *big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.ProposeWithSnapshotBlockNumber(&_Contracts.TransactOpts, targets, values, calldatas, description, _snapshotBlockNumber)
}

// Queue is a paid mutator transaction binding the contract method 0x160cbed7.
//
// Solidity: function queue(address[] targets, uint256[] values, bytes[] calldatas, bytes32 descriptionHash) returns(uint256)
func (_Contracts *ContractsTransactor) Queue(opts *bind.TransactOpts, targets []common.Address, values []*big.Int, calldatas [][]byte, descriptionHash [32]byte) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "queue", targets, values, calldatas, descriptionHash)
}

// Queue is a paid mutator transaction binding the contract method 0x160cbed7.
//
// Solidity: function queue(address[] targets, uint256[] values, bytes[] calldatas, bytes32 descriptionHash) returns(uint256)
func (_Contracts *ContractsSession) Queue(targets []common.Address, values []*big.Int, calldatas [][]byte, descriptionHash [32]byte) (*types.Transaction, error) {
	return _Contracts.Contract.Queue(&_Contracts.TransactOpts, targets, values, calldatas, descriptionHash)
}

// Queue is a paid mutator transaction binding the contract method 0x160cbed7.
//
// Solidity: function queue(address[] targets, uint256[] values, bytes[] calldatas, bytes32 descriptionHash) returns(uint256)
func (_Contracts *ContractsTransactorSession) Queue(targets []common.Address, values []*big.Int, calldatas [][]byte, descriptionHash [32]byte) (*types.Transaction, error) {
	return _Contracts.Contract.Queue(&_Contracts.TransactOpts, targets, values, calldatas, descriptionHash)
}

// Relay is a paid mutator transaction binding the contract method 0xc28bc2fa.
//
// Solidity: function relay(address target, uint256 value, bytes data) payable returns()
func (_Contracts *ContractsTransactor) Relay(opts *bind.TransactOpts, target common.Address, value *big.Int, data []byte) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "relay", target, value, data)
}

// Relay is a paid mutator transaction binding the contract method 0xc28bc2fa.
//
// Solidity: function relay(address target, uint256 value, bytes data) payable returns()
func (_Contracts *ContractsSession) Relay(target common.Address, value *big.Int, data []byte) (*types.Transaction, error) {
	return _Contracts.Contract.Relay(&_Contracts.TransactOpts, target, value, data)
}

// Relay is a paid mutator transaction binding the contract method 0xc28bc2fa.
//
// Solidity: function relay(address target, uint256 value, bytes data) payable returns()
func (_Contracts *ContractsTransactorSession) Relay(target common.Address, value *big.Int, data []byte) (*types.Transaction, error) {
	return _Contracts.Contract.Relay(&_Contracts.TransactOpts, target, value, data)
}

// ResultFromSignal is a paid mutator transaction binding the contract method 0x9fbd2818.
//
// Solidity: function resultFromSignal(address app, bytes32 _messageHash, (uint64,address,uint64,uint64,uint256,uint256,uint256,uint256) message, bytes proof) returns()
func (_Contracts *ContractsTransactor) ResultFromSignal(opts *bind.TransactOpts, app common.Address, _messageHash [32]byte, message GovernorCountingFromSignalProposalVoteMessage, proof []byte) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "resultFromSignal", app, _messageHash, message, proof)
}

// ResultFromSignal is a paid mutator transaction binding the contract method 0x9fbd2818.
//
// Solidity: function resultFromSignal(address app, bytes32 _messageHash, (uint64,address,uint64,uint64,uint256,uint256,uint256,uint256) message, bytes proof) returns()
func (_Contracts *ContractsSession) ResultFromSignal(app common.Address, _messageHash [32]byte, message GovernorCountingFromSignalProposalVoteMessage, proof []byte) (*types.Transaction, error) {
	return _Contracts.Contract.ResultFromSignal(&_Contracts.TransactOpts, app, _messageHash, message, proof)
}

// ResultFromSignal is a paid mutator transaction binding the contract method 0x9fbd2818.
//
// Solidity: function resultFromSignal(address app, bytes32 _messageHash, (uint64,address,uint64,uint64,uint256,uint256,uint256,uint256) message, bytes proof) returns()
func (_Contracts *ContractsTransactorSession) ResultFromSignal(app common.Address, _messageHash [32]byte, message GovernorCountingFromSignalProposalVoteMessage, proof []byte) (*types.Transaction, error) {
	return _Contracts.Contract.ResultFromSignal(&_Contracts.TransactOpts, app, _messageHash, message, proof)
}

// UpdateQuorumNumerator is a paid mutator transaction binding the contract method 0x06f3f9e6.
//
// Solidity: function updateQuorumNumerator(uint256 newQuorumNumerator) returns()
func (_Contracts *ContractsTransactor) UpdateQuorumNumerator(opts *bind.TransactOpts, newQuorumNumerator *big.Int) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "updateQuorumNumerator", newQuorumNumerator)
}

// UpdateQuorumNumerator is a paid mutator transaction binding the contract method 0x06f3f9e6.
//
// Solidity: function updateQuorumNumerator(uint256 newQuorumNumerator) returns()
func (_Contracts *ContractsSession) UpdateQuorumNumerator(newQuorumNumerator *big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.UpdateQuorumNumerator(&_Contracts.TransactOpts, newQuorumNumerator)
}

// UpdateQuorumNumerator is a paid mutator transaction binding the contract method 0x06f3f9e6.
//
// Solidity: function updateQuorumNumerator(uint256 newQuorumNumerator) returns()
func (_Contracts *ContractsTransactorSession) UpdateQuorumNumerator(newQuorumNumerator *big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.UpdateQuorumNumerator(&_Contracts.TransactOpts, newQuorumNumerator)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Contracts *ContractsTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contracts.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Contracts *ContractsSession) Receive() (*types.Transaction, error) {
	return _Contracts.Contract.Receive(&_Contracts.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Contracts *ContractsTransactorSession) Receive() (*types.Transaction, error) {
	return _Contracts.Contract.Receive(&_Contracts.TransactOpts)
}

// ContractsEIP712DomainChangedIterator is returned from FilterEIP712DomainChanged and is used to iterate over the raw logs and unpacked data for EIP712DomainChanged events raised by the Contracts contract.
type ContractsEIP712DomainChangedIterator struct {
	Event *ContractsEIP712DomainChanged // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractsEIP712DomainChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractsEIP712DomainChanged)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractsEIP712DomainChanged)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractsEIP712DomainChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractsEIP712DomainChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractsEIP712DomainChanged represents a EIP712DomainChanged event raised by the Contracts contract.
type ContractsEIP712DomainChanged struct {
	Raw types.Log // Blockchain specific contextual infos
}

// String ContractsEIP712DomainChanged returns a readable string representing the event struct.
//
// Solidity: event EIP712DomainChanged()
func (e *ContractsEIP712DomainChanged) String() string {
	s := "Event: ContractsEIP712DomainChanged {\n"

	s += "}"

	return s
}

// FilterEIP712DomainChanged is a free log retrieval operation binding the contract event 0x0a6387c9ea3628b88a633bb4f3b151770f70085117a15f9bf3787cda53f13d31.
//
// Solidity: event EIP712DomainChanged()
func (_Contracts *ContractsFilterer) FilterEIP712DomainChanged(opts *bind.FilterOpts) (*ContractsEIP712DomainChangedIterator, error) {

	logs, sub, err := _Contracts.contract.FilterLogs(opts, "EIP712DomainChanged")
	if err != nil {
		return nil, err
	}
	return &ContractsEIP712DomainChangedIterator{contract: _Contracts.contract, event: "EIP712DomainChanged", logs: logs, sub: sub}, nil
}

// WatchEIP712DomainChanged is a free log subscription operation binding the contract event 0x0a6387c9ea3628b88a633bb4f3b151770f70085117a15f9bf3787cda53f13d31.
//
// Solidity: event EIP712DomainChanged()
func (_Contracts *ContractsFilterer) WatchEIP712DomainChanged(opts *bind.WatchOpts, sink chan<- *ContractsEIP712DomainChanged) (event.Subscription, error) {

	logs, sub, err := _Contracts.contract.WatchLogs(opts, "EIP712DomainChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractsEIP712DomainChanged)
				if err := _Contracts.contract.UnpackLog(event, "EIP712DomainChanged", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseEIP712DomainChanged is a log parse operation binding the contract event 0x0a6387c9ea3628b88a633bb4f3b151770f70085117a15f9bf3787cda53f13d31.
//
// Solidity: event EIP712DomainChanged()
func (_Contracts *ContractsFilterer) ParseEIP712DomainChanged(log types.Log) (*ContractsEIP712DomainChanged, error) {
	event := new(ContractsEIP712DomainChanged)
	if err := _Contracts.contract.UnpackLog(event, "EIP712DomainChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractsMessageSentIterator is returned from FilterMessageSent and is used to iterate over the raw logs and unpacked data for MessageSent events raised by the Contracts contract.
type ContractsMessageSentIterator struct {
	Event *ContractsMessageSent // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractsMessageSentIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractsMessageSent)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractsMessageSent)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractsMessageSentIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractsMessageSentIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractsMessageSent represents a MessageSent event raised by the Contracts contract.
type ContractsMessageSent struct {
	MessageHash [32]byte
	Message     OffchGovernorMessage
	Raw         types.Log // Blockchain specific contextual infos
}

// String ContractsMessageSent returns a readable string representing the event struct.
//
// Solidity: event MessageSent(bytes32 messageHash, (uint64,address,uint64,uint64,uint256,uint256,uint256,address) message)
func (e *ContractsMessageSent) String() string {
	s := "Event: ContractsMessageSent {\n"

	s += fmt.Sprintf("MessageHash: %v,\n", e.MessageHash)
	s += fmt.Sprintf("Message: %v,\n", e.Message)
	s += "}"

	return s
}

// FilterMessageSent is a free log retrieval operation binding the contract event 0x3d151231d49a27ffabff6c55740d7f90f6fa45e81fcfc2dec9ecf1637488e171.
//
// Solidity: event MessageSent(bytes32 messageHash, (uint64,address,uint64,uint64,uint256,uint256,uint256,address) message)
func (_Contracts *ContractsFilterer) FilterMessageSent(opts *bind.FilterOpts) (*ContractsMessageSentIterator, error) {

	logs, sub, err := _Contracts.contract.FilterLogs(opts, "MessageSent")
	if err != nil {
		return nil, err
	}
	return &ContractsMessageSentIterator{contract: _Contracts.contract, event: "MessageSent", logs: logs, sub: sub}, nil
}

// WatchMessageSent is a free log subscription operation binding the contract event 0x3d151231d49a27ffabff6c55740d7f90f6fa45e81fcfc2dec9ecf1637488e171.
//
// Solidity: event MessageSent(bytes32 messageHash, (uint64,address,uint64,uint64,uint256,uint256,uint256,address) message)
func (_Contracts *ContractsFilterer) WatchMessageSent(opts *bind.WatchOpts, sink chan<- *ContractsMessageSent) (event.Subscription, error) {

	logs, sub, err := _Contracts.contract.WatchLogs(opts, "MessageSent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractsMessageSent)
				if err := _Contracts.contract.UnpackLog(event, "MessageSent", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseMessageSent is a log parse operation binding the contract event 0x3d151231d49a27ffabff6c55740d7f90f6fa45e81fcfc2dec9ecf1637488e171.
//
// Solidity: event MessageSent(bytes32 messageHash, (uint64,address,uint64,uint64,uint256,uint256,uint256,address) message)
func (_Contracts *ContractsFilterer) ParseMessageSent(log types.Log) (*ContractsMessageSent, error) {
	event := new(ContractsMessageSent)
	if err := _Contracts.contract.UnpackLog(event, "MessageSent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractsProposalCanceledIterator is returned from FilterProposalCanceled and is used to iterate over the raw logs and unpacked data for ProposalCanceled events raised by the Contracts contract.
type ContractsProposalCanceledIterator struct {
	Event *ContractsProposalCanceled // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractsProposalCanceledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractsProposalCanceled)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractsProposalCanceled)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractsProposalCanceledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractsProposalCanceledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractsProposalCanceled represents a ProposalCanceled event raised by the Contracts contract.
type ContractsProposalCanceled struct {
	ProposalId *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// String ContractsProposalCanceled returns a readable string representing the event struct.
//
// Solidity: event ProposalCanceled(uint256 proposalId)
func (e *ContractsProposalCanceled) String() string {
	s := "Event: ContractsProposalCanceled {\n"

	s += fmt.Sprintf("ProposalId: %v,\n", e.ProposalId)
	s += "}"

	return s
}

// FilterProposalCanceled is a free log retrieval operation binding the contract event 0x789cf55be980739dad1d0699b93b58e806b51c9d96619bfa8fe0a28abaa7b30c.
//
// Solidity: event ProposalCanceled(uint256 proposalId)
func (_Contracts *ContractsFilterer) FilterProposalCanceled(opts *bind.FilterOpts) (*ContractsProposalCanceledIterator, error) {

	logs, sub, err := _Contracts.contract.FilterLogs(opts, "ProposalCanceled")
	if err != nil {
		return nil, err
	}
	return &ContractsProposalCanceledIterator{contract: _Contracts.contract, event: "ProposalCanceled", logs: logs, sub: sub}, nil
}

// WatchProposalCanceled is a free log subscription operation binding the contract event 0x789cf55be980739dad1d0699b93b58e806b51c9d96619bfa8fe0a28abaa7b30c.
//
// Solidity: event ProposalCanceled(uint256 proposalId)
func (_Contracts *ContractsFilterer) WatchProposalCanceled(opts *bind.WatchOpts, sink chan<- *ContractsProposalCanceled) (event.Subscription, error) {

	logs, sub, err := _Contracts.contract.WatchLogs(opts, "ProposalCanceled")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractsProposalCanceled)
				if err := _Contracts.contract.UnpackLog(event, "ProposalCanceled", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseProposalCanceled is a log parse operation binding the contract event 0x789cf55be980739dad1d0699b93b58e806b51c9d96619bfa8fe0a28abaa7b30c.
//
// Solidity: event ProposalCanceled(uint256 proposalId)
func (_Contracts *ContractsFilterer) ParseProposalCanceled(log types.Log) (*ContractsProposalCanceled, error) {
	event := new(ContractsProposalCanceled)
	if err := _Contracts.contract.UnpackLog(event, "ProposalCanceled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractsProposalCreatedIterator is returned from FilterProposalCreated and is used to iterate over the raw logs and unpacked data for ProposalCreated events raised by the Contracts contract.
type ContractsProposalCreatedIterator struct {
	Event *ContractsProposalCreated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractsProposalCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractsProposalCreated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractsProposalCreated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractsProposalCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractsProposalCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractsProposalCreated represents a ProposalCreated event raised by the Contracts contract.
type ContractsProposalCreated struct {
	ProposalId  *big.Int
	Proposer    common.Address
	Targets     []common.Address
	Values      []*big.Int
	Signatures  []string
	Calldatas   [][]byte
	VoteStart   *big.Int
	VoteEnd     *big.Int
	Description string
	Raw         types.Log // Blockchain specific contextual infos
}

// String ContractsProposalCreated returns a readable string representing the event struct.
//
// Solidity: event ProposalCreated(uint256 proposalId, address proposer, address[] targets, uint256[] values, string[] signatures, bytes[] calldatas, uint256 voteStart, uint256 voteEnd, string description)
func (e *ContractsProposalCreated) String() string {
	s := "Event: ContractsProposalCreated {\n"

	s += fmt.Sprintf("ProposalId: %v,\n", e.ProposalId)
	s += fmt.Sprintf("Proposer: %v,\n", e.Proposer)
	s += fmt.Sprintf("Targets: %v,\n", e.Targets)
	s += fmt.Sprintf("Values: %v,\n", e.Values)
	s += fmt.Sprintf("Signatures: %v,\n", e.Signatures)
	s += fmt.Sprintf("Calldatas: %v,\n", e.Calldatas)
	s += fmt.Sprintf("VoteStart: %v,\n", e.VoteStart)
	s += fmt.Sprintf("VoteEnd: %v,\n", e.VoteEnd)
	s += fmt.Sprintf("Description: %v,\n", e.Description)
	s += "}"

	return s
}

// FilterProposalCreated is a free log retrieval operation binding the contract event 0x7d84a6263ae0d98d3329bd7b46bb4e8d6f98cd35a7adb45c274c8b7fd5ebd5e0.
//
// Solidity: event ProposalCreated(uint256 proposalId, address proposer, address[] targets, uint256[] values, string[] signatures, bytes[] calldatas, uint256 voteStart, uint256 voteEnd, string description)
func (_Contracts *ContractsFilterer) FilterProposalCreated(opts *bind.FilterOpts) (*ContractsProposalCreatedIterator, error) {

	logs, sub, err := _Contracts.contract.FilterLogs(opts, "ProposalCreated")
	if err != nil {
		return nil, err
	}
	return &ContractsProposalCreatedIterator{contract: _Contracts.contract, event: "ProposalCreated", logs: logs, sub: sub}, nil
}

// WatchProposalCreated is a free log subscription operation binding the contract event 0x7d84a6263ae0d98d3329bd7b46bb4e8d6f98cd35a7adb45c274c8b7fd5ebd5e0.
//
// Solidity: event ProposalCreated(uint256 proposalId, address proposer, address[] targets, uint256[] values, string[] signatures, bytes[] calldatas, uint256 voteStart, uint256 voteEnd, string description)
func (_Contracts *ContractsFilterer) WatchProposalCreated(opts *bind.WatchOpts, sink chan<- *ContractsProposalCreated) (event.Subscription, error) {

	logs, sub, err := _Contracts.contract.WatchLogs(opts, "ProposalCreated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractsProposalCreated)
				if err := _Contracts.contract.UnpackLog(event, "ProposalCreated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseProposalCreated is a log parse operation binding the contract event 0x7d84a6263ae0d98d3329bd7b46bb4e8d6f98cd35a7adb45c274c8b7fd5ebd5e0.
//
// Solidity: event ProposalCreated(uint256 proposalId, address proposer, address[] targets, uint256[] values, string[] signatures, bytes[] calldatas, uint256 voteStart, uint256 voteEnd, string description)
func (_Contracts *ContractsFilterer) ParseProposalCreated(log types.Log) (*ContractsProposalCreated, error) {
	event := new(ContractsProposalCreated)
	if err := _Contracts.contract.UnpackLog(event, "ProposalCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractsProposalExecutedIterator is returned from FilterProposalExecuted and is used to iterate over the raw logs and unpacked data for ProposalExecuted events raised by the Contracts contract.
type ContractsProposalExecutedIterator struct {
	Event *ContractsProposalExecuted // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractsProposalExecutedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractsProposalExecuted)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractsProposalExecuted)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractsProposalExecutedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractsProposalExecutedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractsProposalExecuted represents a ProposalExecuted event raised by the Contracts contract.
type ContractsProposalExecuted struct {
	ProposalId *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// String ContractsProposalExecuted returns a readable string representing the event struct.
//
// Solidity: event ProposalExecuted(uint256 proposalId)
func (e *ContractsProposalExecuted) String() string {
	s := "Event: ContractsProposalExecuted {\n"

	s += fmt.Sprintf("ProposalId: %v,\n", e.ProposalId)
	s += "}"

	return s
}

// FilterProposalExecuted is a free log retrieval operation binding the contract event 0x712ae1383f79ac853f8d882153778e0260ef8f03b504e2866e0593e04d2b291f.
//
// Solidity: event ProposalExecuted(uint256 proposalId)
func (_Contracts *ContractsFilterer) FilterProposalExecuted(opts *bind.FilterOpts) (*ContractsProposalExecutedIterator, error) {

	logs, sub, err := _Contracts.contract.FilterLogs(opts, "ProposalExecuted")
	if err != nil {
		return nil, err
	}
	return &ContractsProposalExecutedIterator{contract: _Contracts.contract, event: "ProposalExecuted", logs: logs, sub: sub}, nil
}

// WatchProposalExecuted is a free log subscription operation binding the contract event 0x712ae1383f79ac853f8d882153778e0260ef8f03b504e2866e0593e04d2b291f.
//
// Solidity: event ProposalExecuted(uint256 proposalId)
func (_Contracts *ContractsFilterer) WatchProposalExecuted(opts *bind.WatchOpts, sink chan<- *ContractsProposalExecuted) (event.Subscription, error) {

	logs, sub, err := _Contracts.contract.WatchLogs(opts, "ProposalExecuted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractsProposalExecuted)
				if err := _Contracts.contract.UnpackLog(event, "ProposalExecuted", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseProposalExecuted is a log parse operation binding the contract event 0x712ae1383f79ac853f8d882153778e0260ef8f03b504e2866e0593e04d2b291f.
//
// Solidity: event ProposalExecuted(uint256 proposalId)
func (_Contracts *ContractsFilterer) ParseProposalExecuted(log types.Log) (*ContractsProposalExecuted, error) {
	event := new(ContractsProposalExecuted)
	if err := _Contracts.contract.UnpackLog(event, "ProposalExecuted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractsProposalQueuedIterator is returned from FilterProposalQueued and is used to iterate over the raw logs and unpacked data for ProposalQueued events raised by the Contracts contract.
type ContractsProposalQueuedIterator struct {
	Event *ContractsProposalQueued // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractsProposalQueuedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractsProposalQueued)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractsProposalQueued)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractsProposalQueuedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractsProposalQueuedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractsProposalQueued represents a ProposalQueued event raised by the Contracts contract.
type ContractsProposalQueued struct {
	ProposalId *big.Int
	EtaSeconds *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// String ContractsProposalQueued returns a readable string representing the event struct.
//
// Solidity: event ProposalQueued(uint256 proposalId, uint256 etaSeconds)
func (e *ContractsProposalQueued) String() string {
	s := "Event: ContractsProposalQueued {\n"

	s += fmt.Sprintf("ProposalId: %v,\n", e.ProposalId)
	s += fmt.Sprintf("EtaSeconds: %v,\n", e.EtaSeconds)
	s += "}"

	return s
}

// FilterProposalQueued is a free log retrieval operation binding the contract event 0x9a2e42fd6722813d69113e7d0079d3d940171428df7373df9c7f7617cfda2892.
//
// Solidity: event ProposalQueued(uint256 proposalId, uint256 etaSeconds)
func (_Contracts *ContractsFilterer) FilterProposalQueued(opts *bind.FilterOpts) (*ContractsProposalQueuedIterator, error) {

	logs, sub, err := _Contracts.contract.FilterLogs(opts, "ProposalQueued")
	if err != nil {
		return nil, err
	}
	return &ContractsProposalQueuedIterator{contract: _Contracts.contract, event: "ProposalQueued", logs: logs, sub: sub}, nil
}

// WatchProposalQueued is a free log subscription operation binding the contract event 0x9a2e42fd6722813d69113e7d0079d3d940171428df7373df9c7f7617cfda2892.
//
// Solidity: event ProposalQueued(uint256 proposalId, uint256 etaSeconds)
func (_Contracts *ContractsFilterer) WatchProposalQueued(opts *bind.WatchOpts, sink chan<- *ContractsProposalQueued) (event.Subscription, error) {

	logs, sub, err := _Contracts.contract.WatchLogs(opts, "ProposalQueued")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractsProposalQueued)
				if err := _Contracts.contract.UnpackLog(event, "ProposalQueued", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseProposalQueued is a log parse operation binding the contract event 0x9a2e42fd6722813d69113e7d0079d3d940171428df7373df9c7f7617cfda2892.
//
// Solidity: event ProposalQueued(uint256 proposalId, uint256 etaSeconds)
func (_Contracts *ContractsFilterer) ParseProposalQueued(log types.Log) (*ContractsProposalQueued, error) {
	event := new(ContractsProposalQueued)
	if err := _Contracts.contract.UnpackLog(event, "ProposalQueued", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractsQuorumNumeratorUpdatedIterator is returned from FilterQuorumNumeratorUpdated and is used to iterate over the raw logs and unpacked data for QuorumNumeratorUpdated events raised by the Contracts contract.
type ContractsQuorumNumeratorUpdatedIterator struct {
	Event *ContractsQuorumNumeratorUpdated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractsQuorumNumeratorUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractsQuorumNumeratorUpdated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractsQuorumNumeratorUpdated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractsQuorumNumeratorUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractsQuorumNumeratorUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractsQuorumNumeratorUpdated represents a QuorumNumeratorUpdated event raised by the Contracts contract.
type ContractsQuorumNumeratorUpdated struct {
	OldQuorumNumerator *big.Int
	NewQuorumNumerator *big.Int
	Raw                types.Log // Blockchain specific contextual infos
}

// String ContractsQuorumNumeratorUpdated returns a readable string representing the event struct.
//
// Solidity: event QuorumNumeratorUpdated(uint256 oldQuorumNumerator, uint256 newQuorumNumerator)
func (e *ContractsQuorumNumeratorUpdated) String() string {
	s := "Event: ContractsQuorumNumeratorUpdated {\n"

	s += fmt.Sprintf("OldQuorumNumerator: %v,\n", e.OldQuorumNumerator)
	s += fmt.Sprintf("NewQuorumNumerator: %v,\n", e.NewQuorumNumerator)
	s += "}"

	return s
}

// FilterQuorumNumeratorUpdated is a free log retrieval operation binding the contract event 0x0553476bf02ef2726e8ce5ced78d63e26e602e4a2257b1f559418e24b4633997.
//
// Solidity: event QuorumNumeratorUpdated(uint256 oldQuorumNumerator, uint256 newQuorumNumerator)
func (_Contracts *ContractsFilterer) FilterQuorumNumeratorUpdated(opts *bind.FilterOpts) (*ContractsQuorumNumeratorUpdatedIterator, error) {

	logs, sub, err := _Contracts.contract.FilterLogs(opts, "QuorumNumeratorUpdated")
	if err != nil {
		return nil, err
	}
	return &ContractsQuorumNumeratorUpdatedIterator{contract: _Contracts.contract, event: "QuorumNumeratorUpdated", logs: logs, sub: sub}, nil
}

// WatchQuorumNumeratorUpdated is a free log subscription operation binding the contract event 0x0553476bf02ef2726e8ce5ced78d63e26e602e4a2257b1f559418e24b4633997.
//
// Solidity: event QuorumNumeratorUpdated(uint256 oldQuorumNumerator, uint256 newQuorumNumerator)
func (_Contracts *ContractsFilterer) WatchQuorumNumeratorUpdated(opts *bind.WatchOpts, sink chan<- *ContractsQuorumNumeratorUpdated) (event.Subscription, error) {

	logs, sub, err := _Contracts.contract.WatchLogs(opts, "QuorumNumeratorUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractsQuorumNumeratorUpdated)
				if err := _Contracts.contract.UnpackLog(event, "QuorumNumeratorUpdated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseQuorumNumeratorUpdated is a log parse operation binding the contract event 0x0553476bf02ef2726e8ce5ced78d63e26e602e4a2257b1f559418e24b4633997.
//
// Solidity: event QuorumNumeratorUpdated(uint256 oldQuorumNumerator, uint256 newQuorumNumerator)
func (_Contracts *ContractsFilterer) ParseQuorumNumeratorUpdated(log types.Log) (*ContractsQuorumNumeratorUpdated, error) {
	event := new(ContractsQuorumNumeratorUpdated)
	if err := _Contracts.contract.UnpackLog(event, "QuorumNumeratorUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractsVoteCastIterator is returned from FilterVoteCast and is used to iterate over the raw logs and unpacked data for VoteCast events raised by the Contracts contract.
type ContractsVoteCastIterator struct {
	Event *ContractsVoteCast // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractsVoteCastIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractsVoteCast)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractsVoteCast)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractsVoteCastIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractsVoteCastIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractsVoteCast represents a VoteCast event raised by the Contracts contract.
type ContractsVoteCast struct {
	Voter      common.Address
	ProposalId *big.Int
	Support    uint8
	Weight     *big.Int
	Reason     string
	Raw        types.Log // Blockchain specific contextual infos
}

// String ContractsVoteCast returns a readable string representing the event struct.
//
// Solidity: event VoteCast(address indexed voter, uint256 proposalId, uint8 support, uint256 weight, string reason)
func (e *ContractsVoteCast) String() string {
	s := "Event: ContractsVoteCast {\n"

	s += fmt.Sprintf("Voter: %v,\n", e.Voter)
	s += fmt.Sprintf("ProposalId: %v,\n", e.ProposalId)
	s += fmt.Sprintf("Support: %v,\n", e.Support)
	s += fmt.Sprintf("Weight: %v,\n", e.Weight)
	s += fmt.Sprintf("Reason: %v,\n", e.Reason)
	s += "}"

	return s
}

// FilterVoteCast is a free log retrieval operation binding the contract event 0xb8e138887d0aa13bab447e82de9d5c1777041ecd21ca36ba824ff1e6c07ddda4.
//
// Solidity: event VoteCast(address indexed voter, uint256 proposalId, uint8 support, uint256 weight, string reason)
func (_Contracts *ContractsFilterer) FilterVoteCast(opts *bind.FilterOpts, voter []common.Address) (*ContractsVoteCastIterator, error) {

	var voterRule []interface{}
	for _, voterItem := range voter {
		voterRule = append(voterRule, voterItem)
	}

	logs, sub, err := _Contracts.contract.FilterLogs(opts, "VoteCast", voterRule)
	if err != nil {
		return nil, err
	}
	return &ContractsVoteCastIterator{contract: _Contracts.contract, event: "VoteCast", logs: logs, sub: sub}, nil
}

// WatchVoteCast is a free log subscription operation binding the contract event 0xb8e138887d0aa13bab447e82de9d5c1777041ecd21ca36ba824ff1e6c07ddda4.
//
// Solidity: event VoteCast(address indexed voter, uint256 proposalId, uint8 support, uint256 weight, string reason)
func (_Contracts *ContractsFilterer) WatchVoteCast(opts *bind.WatchOpts, sink chan<- *ContractsVoteCast, voter []common.Address) (event.Subscription, error) {

	var voterRule []interface{}
	for _, voterItem := range voter {
		voterRule = append(voterRule, voterItem)
	}

	logs, sub, err := _Contracts.contract.WatchLogs(opts, "VoteCast", voterRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractsVoteCast)
				if err := _Contracts.contract.UnpackLog(event, "VoteCast", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseVoteCast is a log parse operation binding the contract event 0xb8e138887d0aa13bab447e82de9d5c1777041ecd21ca36ba824ff1e6c07ddda4.
//
// Solidity: event VoteCast(address indexed voter, uint256 proposalId, uint8 support, uint256 weight, string reason)
func (_Contracts *ContractsFilterer) ParseVoteCast(log types.Log) (*ContractsVoteCast, error) {
	event := new(ContractsVoteCast)
	if err := _Contracts.contract.UnpackLog(event, "VoteCast", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractsVoteCastWithParamsIterator is returned from FilterVoteCastWithParams and is used to iterate over the raw logs and unpacked data for VoteCastWithParams events raised by the Contracts contract.
type ContractsVoteCastWithParamsIterator struct {
	Event *ContractsVoteCastWithParams // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractsVoteCastWithParamsIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractsVoteCastWithParams)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractsVoteCastWithParams)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractsVoteCastWithParamsIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractsVoteCastWithParamsIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractsVoteCastWithParams represents a VoteCastWithParams event raised by the Contracts contract.
type ContractsVoteCastWithParams struct {
	Voter      common.Address
	ProposalId *big.Int
	Support    uint8
	Weight     *big.Int
	Reason     string
	Params     []byte
	Raw        types.Log // Blockchain specific contextual infos
}

// String ContractsVoteCastWithParams returns a readable string representing the event struct.
//
// Solidity: event VoteCastWithParams(address indexed voter, uint256 proposalId, uint8 support, uint256 weight, string reason, bytes params)
func (e *ContractsVoteCastWithParams) String() string {
	s := "Event: ContractsVoteCastWithParams {\n"

	s += fmt.Sprintf("Voter: %v,\n", e.Voter)
	s += fmt.Sprintf("ProposalId: %v,\n", e.ProposalId)
	s += fmt.Sprintf("Support: %v,\n", e.Support)
	s += fmt.Sprintf("Weight: %v,\n", e.Weight)
	s += fmt.Sprintf("Reason: %v,\n", e.Reason)
	s += fmt.Sprintf("Params: %v,\n", e.Params)
	s += "}"

	return s
}

// FilterVoteCastWithParams is a free log retrieval operation binding the contract event 0xe2babfbac5889a709b63bb7f598b324e08bc5a4fb9ec647fb3cbc9ec07eb8712.
//
// Solidity: event VoteCastWithParams(address indexed voter, uint256 proposalId, uint8 support, uint256 weight, string reason, bytes params)
func (_Contracts *ContractsFilterer) FilterVoteCastWithParams(opts *bind.FilterOpts, voter []common.Address) (*ContractsVoteCastWithParamsIterator, error) {

	var voterRule []interface{}
	for _, voterItem := range voter {
		voterRule = append(voterRule, voterItem)
	}

	logs, sub, err := _Contracts.contract.FilterLogs(opts, "VoteCastWithParams", voterRule)
	if err != nil {
		return nil, err
	}
	return &ContractsVoteCastWithParamsIterator{contract: _Contracts.contract, event: "VoteCastWithParams", logs: logs, sub: sub}, nil
}

// WatchVoteCastWithParams is a free log subscription operation binding the contract event 0xe2babfbac5889a709b63bb7f598b324e08bc5a4fb9ec647fb3cbc9ec07eb8712.
//
// Solidity: event VoteCastWithParams(address indexed voter, uint256 proposalId, uint8 support, uint256 weight, string reason, bytes params)
func (_Contracts *ContractsFilterer) WatchVoteCastWithParams(opts *bind.WatchOpts, sink chan<- *ContractsVoteCastWithParams, voter []common.Address) (event.Subscription, error) {

	var voterRule []interface{}
	for _, voterItem := range voter {
		voterRule = append(voterRule, voterItem)
	}

	logs, sub, err := _Contracts.contract.WatchLogs(opts, "VoteCastWithParams", voterRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractsVoteCastWithParams)
				if err := _Contracts.contract.UnpackLog(event, "VoteCastWithParams", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseVoteCastWithParams is a log parse operation binding the contract event 0xe2babfbac5889a709b63bb7f598b324e08bc5a4fb9ec647fb3cbc9ec07eb8712.
//
// Solidity: event VoteCastWithParams(address indexed voter, uint256 proposalId, uint8 support, uint256 weight, string reason, bytes params)
func (_Contracts *ContractsFilterer) ParseVoteCastWithParams(log types.Log) (*ContractsVoteCastWithParams, error) {
	event := new(ContractsVoteCastWithParams)
	if err := _Contracts.contract.UnpackLog(event, "VoteCastWithParams", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
