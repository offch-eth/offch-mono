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

// ISignalServiceHop is an auto generated low-level Go binding around an user-defined struct.
type ISignalServiceHop struct {
	SignalRootRelay common.Address
	SignalRoot      [32]byte
	StorageProof    []byte
}

// String ISignalServiceHop returns a readable string representing the user-defined struct.
func (st *ISignalServiceHop) String() string {
	s := "User defined struct: ISignalServiceHop {\n"

	s += fmt.Sprintf("SignalRootRelay: %v,\n", st.SignalRootRelay)
	s += fmt.Sprintf("SignalRoot: %v,\n", st.SignalRoot)
	s += fmt.Sprintf("StorageProof: %v,\n", st.StorageProof)
	s += "}"

	return s
}

// ISignalServiceProof is an auto generated low-level Go binding around an user-defined struct.
type ISignalServiceProof struct {
	CrossChainSync common.Address
	Height         uint64
	StorageProof   []byte
	Hops           []ISignalServiceHop
}

// String ISignalServiceProof returns a readable string representing the user-defined struct.
func (st *ISignalServiceProof) String() string {
	s := "User defined struct: ISignalServiceProof {\n"

	s += fmt.Sprintf("CrossChainSync: %v,\n", st.CrossChainSync)
	s += fmt.Sprintf("Height: %v,\n", st.Height)
	s += fmt.Sprintf("StorageProof: %v,\n", st.StorageProof)
	s += fmt.Sprintf("Hops: %v,\n", st.Hops)
	s += "}"

	return s
}

// ContractsMetaData contains all meta data concerning the Contracts contract.
var ContractsMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"srcChainId\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"app\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"signal\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"crossChainSync\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"height\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"storageProof\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"signalRootRelay\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"signalRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"storageProof\",\"type\":\"bytes\"}],\"internalType\":\"structISignalService.Hop[]\",\"name\":\"hops\",\"type\":\"tuple[]\"}],\"internalType\":\"structISignalService.Proof\",\"name\":\"proof\",\"type\":\"tuple\"}],\"name\":\"_foo\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"chainId\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"app\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"signal\",\"type\":\"bytes32\"}],\"name\":\"getSignalSlot\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"signalSlot\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"app\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"signal\",\"type\":\"bytes32\"}],\"name\":\"isSignalSent\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"srcChainId\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"app\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"signal\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"proof\",\"type\":\"bytes\"}],\"name\":\"proveSignalReceived\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"signal\",\"type\":\"bytes32\"}],\"name\":\"sendSignal\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"storageSlot\",\"type\":\"bytes32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// ContractsABI is the input ABI used to generate the binding from.
// Deprecated: Use ContractsMetaData.ABI instead.
var ContractsABI = ContractsMetaData.ABI

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

// Foo is a free data retrieval call binding the contract method 0x5abaecd6.
//
// Solidity: function _foo(uint64 srcChainId, address app, bytes32 signal, (address,uint64,bytes,(address,bytes32,bytes)[]) proof) view returns(bool)
func (_Contracts *ContractsCaller) Foo(opts *bind.CallOpts, srcChainId uint64, app common.Address, signal [32]byte, proof ISignalServiceProof) (bool, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "_foo", srcChainId, app, signal, proof)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Foo is a free data retrieval call binding the contract method 0x5abaecd6.
//
// Solidity: function _foo(uint64 srcChainId, address app, bytes32 signal, (address,uint64,bytes,(address,bytes32,bytes)[]) proof) view returns(bool)
func (_Contracts *ContractsSession) Foo(srcChainId uint64, app common.Address, signal [32]byte, proof ISignalServiceProof) (bool, error) {
	return _Contracts.Contract.Foo(&_Contracts.CallOpts, srcChainId, app, signal, proof)
}

// Foo is a free data retrieval call binding the contract method 0x5abaecd6.
//
// Solidity: function _foo(uint64 srcChainId, address app, bytes32 signal, (address,uint64,bytes,(address,bytes32,bytes)[]) proof) view returns(bool)
func (_Contracts *ContractsCallerSession) Foo(srcChainId uint64, app common.Address, signal [32]byte, proof ISignalServiceProof) (bool, error) {
	return _Contracts.Contract.Foo(&_Contracts.CallOpts, srcChainId, app, signal, proof)
}

// GetSignalSlot is a free data retrieval call binding the contract method 0x91f3f74b.
//
// Solidity: function getSignalSlot(uint64 chainId, address app, bytes32 signal) pure returns(bytes32 signalSlot)
func (_Contracts *ContractsCaller) GetSignalSlot(opts *bind.CallOpts, chainId uint64, app common.Address, signal [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "getSignalSlot", chainId, app, signal)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetSignalSlot is a free data retrieval call binding the contract method 0x91f3f74b.
//
// Solidity: function getSignalSlot(uint64 chainId, address app, bytes32 signal) pure returns(bytes32 signalSlot)
func (_Contracts *ContractsSession) GetSignalSlot(chainId uint64, app common.Address, signal [32]byte) ([32]byte, error) {
	return _Contracts.Contract.GetSignalSlot(&_Contracts.CallOpts, chainId, app, signal)
}

// GetSignalSlot is a free data retrieval call binding the contract method 0x91f3f74b.
//
// Solidity: function getSignalSlot(uint64 chainId, address app, bytes32 signal) pure returns(bytes32 signalSlot)
func (_Contracts *ContractsCallerSession) GetSignalSlot(chainId uint64, app common.Address, signal [32]byte) ([32]byte, error) {
	return _Contracts.Contract.GetSignalSlot(&_Contracts.CallOpts, chainId, app, signal)
}

// IsSignalSent is a free data retrieval call binding the contract method 0x32676bc6.
//
// Solidity: function isSignalSent(address app, bytes32 signal) view returns(bool)
func (_Contracts *ContractsCaller) IsSignalSent(opts *bind.CallOpts, app common.Address, signal [32]byte) (bool, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "isSignalSent", app, signal)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsSignalSent is a free data retrieval call binding the contract method 0x32676bc6.
//
// Solidity: function isSignalSent(address app, bytes32 signal) view returns(bool)
func (_Contracts *ContractsSession) IsSignalSent(app common.Address, signal [32]byte) (bool, error) {
	return _Contracts.Contract.IsSignalSent(&_Contracts.CallOpts, app, signal)
}

// IsSignalSent is a free data retrieval call binding the contract method 0x32676bc6.
//
// Solidity: function isSignalSent(address app, bytes32 signal) view returns(bool)
func (_Contracts *ContractsCallerSession) IsSignalSent(app common.Address, signal [32]byte) (bool, error) {
	return _Contracts.Contract.IsSignalSent(&_Contracts.CallOpts, app, signal)
}

// ProveSignalReceived is a free data retrieval call binding the contract method 0x910af6ed.
//
// Solidity: function proveSignalReceived(uint64 srcChainId, address app, bytes32 signal, bytes proof) view returns(bool)
func (_Contracts *ContractsCaller) ProveSignalReceived(opts *bind.CallOpts, srcChainId uint64, app common.Address, signal [32]byte, proof []byte) (bool, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "proveSignalReceived", srcChainId, app, signal, proof)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// ProveSignalReceived is a free data retrieval call binding the contract method 0x910af6ed.
//
// Solidity: function proveSignalReceived(uint64 srcChainId, address app, bytes32 signal, bytes proof) view returns(bool)
func (_Contracts *ContractsSession) ProveSignalReceived(srcChainId uint64, app common.Address, signal [32]byte, proof []byte) (bool, error) {
	return _Contracts.Contract.ProveSignalReceived(&_Contracts.CallOpts, srcChainId, app, signal, proof)
}

// ProveSignalReceived is a free data retrieval call binding the contract method 0x910af6ed.
//
// Solidity: function proveSignalReceived(uint64 srcChainId, address app, bytes32 signal, bytes proof) view returns(bool)
func (_Contracts *ContractsCallerSession) ProveSignalReceived(srcChainId uint64, app common.Address, signal [32]byte, proof []byte) (bool, error) {
	return _Contracts.Contract.ProveSignalReceived(&_Contracts.CallOpts, srcChainId, app, signal, proof)
}

// SendSignal is a paid mutator transaction binding the contract method 0x66ca2bc0.
//
// Solidity: function sendSignal(bytes32 signal) returns(bytes32 storageSlot)
func (_Contracts *ContractsTransactor) SendSignal(opts *bind.TransactOpts, signal [32]byte) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "sendSignal", signal)
}

// SendSignal is a paid mutator transaction binding the contract method 0x66ca2bc0.
//
// Solidity: function sendSignal(bytes32 signal) returns(bytes32 storageSlot)
func (_Contracts *ContractsSession) SendSignal(signal [32]byte) (*types.Transaction, error) {
	return _Contracts.Contract.SendSignal(&_Contracts.TransactOpts, signal)
}

// SendSignal is a paid mutator transaction binding the contract method 0x66ca2bc0.
//
// Solidity: function sendSignal(bytes32 signal) returns(bytes32 storageSlot)
func (_Contracts *ContractsTransactorSession) SendSignal(signal [32]byte) (*types.Transaction, error) {
	return _Contracts.Contract.SendSignal(&_Contracts.TransactOpts, signal)
}
