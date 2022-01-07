// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package multicallabi

import (
	"errors"
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
)

// Struct0 is an auto generated low-level Go binding around an user-defined struct.
type Struct0 struct {
	Target   common.Address
	CallData []byte
}

// MulticallabiMetaData contains all meta data concerning the Multicallabi contract.
var MulticallabiMetaData = &bind.MetaData{
	ABI: "[{\"constant\":true,\"inputs\":[],\"name\":\"getCurrentBlockTimestamp\",\"outputs\":[{\"name\":\"timestamp\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"components\":[{\"name\":\"target\",\"type\":\"address\"},{\"name\":\"callData\",\"type\":\"bytes\"}],\"name\":\"calls\",\"type\":\"tuple[]\"}],\"name\":\"aggregate\",\"outputs\":[{\"name\":\"blockNumber\",\"type\":\"uint256\"},{\"name\":\"returnData\",\"type\":\"bytes[]\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getLastBlockHash\",\"outputs\":[{\"name\":\"blockHash\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"getEthBalance\",\"outputs\":[{\"name\":\"balance\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getCurrentBlockDifficulty\",\"outputs\":[{\"name\":\"difficulty\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getCurrentBlockGasLimit\",\"outputs\":[{\"name\":\"gaslimit\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getCurrentBlockCoinbase\",\"outputs\":[{\"name\":\"coinbase\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"blockNumber\",\"type\":\"uint256\"}],\"name\":\"getBlockHash\",\"outputs\":[{\"name\":\"blockHash\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// MulticallabiABI is the input ABI used to generate the binding from.
// Deprecated: Use MulticallabiMetaData.ABI instead.
var MulticallabiABI = MulticallabiMetaData.ABI

// Multicallabi is an auto generated Go binding around an Ethereum contract.
type Multicallabi struct {
	MulticallabiCaller     // Read-only binding to the contract
	MulticallabiTransactor // Write-only binding to the contract
	MulticallabiFilterer   // Log filterer for contract events
}

// MulticallabiCaller is an auto generated read-only Go binding around an Ethereum contract.
type MulticallabiCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MulticallabiTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MulticallabiTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MulticallabiFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MulticallabiFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MulticallabiSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MulticallabiSession struct {
	Contract     *Multicallabi     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MulticallabiCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MulticallabiCallerSession struct {
	Contract *MulticallabiCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// MulticallabiTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MulticallabiTransactorSession struct {
	Contract     *MulticallabiTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// MulticallabiRaw is an auto generated low-level Go binding around an Ethereum contract.
type MulticallabiRaw struct {
	Contract *Multicallabi // Generic contract binding to access the raw methods on
}

// MulticallabiCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MulticallabiCallerRaw struct {
	Contract *MulticallabiCaller // Generic read-only contract binding to access the raw methods on
}

// MulticallabiTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MulticallabiTransactorRaw struct {
	Contract *MulticallabiTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMulticallabi creates a new instance of Multicallabi, bound to a specific deployed contract.
func NewMulticallabi(address common.Address, backend bind.ContractBackend) (*Multicallabi, error) {
	contract, err := bindMulticallabi(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Multicallabi{MulticallabiCaller: MulticallabiCaller{contract: contract}, MulticallabiTransactor: MulticallabiTransactor{contract: contract}, MulticallabiFilterer: MulticallabiFilterer{contract: contract}}, nil
}

// NewMulticallabiCaller creates a new read-only instance of Multicallabi, bound to a specific deployed contract.
func NewMulticallabiCaller(address common.Address, caller bind.ContractCaller) (*MulticallabiCaller, error) {
	contract, err := bindMulticallabi(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MulticallabiCaller{contract: contract}, nil
}

// NewMulticallabiTransactor creates a new write-only instance of Multicallabi, bound to a specific deployed contract.
func NewMulticallabiTransactor(address common.Address, transactor bind.ContractTransactor) (*MulticallabiTransactor, error) {
	contract, err := bindMulticallabi(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MulticallabiTransactor{contract: contract}, nil
}

// NewMulticallabiFilterer creates a new log filterer instance of Multicallabi, bound to a specific deployed contract.
func NewMulticallabiFilterer(address common.Address, filterer bind.ContractFilterer) (*MulticallabiFilterer, error) {
	contract, err := bindMulticallabi(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MulticallabiFilterer{contract: contract}, nil
}

// bindMulticallabi binds a generic wrapper to an already deployed contract.
func bindMulticallabi(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(MulticallabiABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Multicallabi *MulticallabiRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Multicallabi.Contract.MulticallabiCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Multicallabi *MulticallabiRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Multicallabi.Contract.MulticallabiTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Multicallabi *MulticallabiRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Multicallabi.Contract.MulticallabiTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Multicallabi *MulticallabiCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Multicallabi.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Multicallabi *MulticallabiTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Multicallabi.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Multicallabi *MulticallabiTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Multicallabi.Contract.contract.Transact(opts, method, params...)
}

// GetBlockHash is a free data retrieval call binding the contract method 0xee82ac5e.
//
// Solidity: function getBlockHash(uint256 blockNumber) view returns(bytes32 blockHash)
func (_Multicallabi *MulticallabiCaller) GetBlockHash(opts *bind.CallOpts, blockNumber *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _Multicallabi.contract.Call(opts, &out, "getBlockHash", blockNumber)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetBlockHash is a free data retrieval call binding the contract method 0xee82ac5e.
//
// Solidity: function getBlockHash(uint256 blockNumber) view returns(bytes32 blockHash)
func (_Multicallabi *MulticallabiSession) GetBlockHash(blockNumber *big.Int) ([32]byte, error) {
	return _Multicallabi.Contract.GetBlockHash(&_Multicallabi.CallOpts, blockNumber)
}

// GetBlockHash is a free data retrieval call binding the contract method 0xee82ac5e.
//
// Solidity: function getBlockHash(uint256 blockNumber) view returns(bytes32 blockHash)
func (_Multicallabi *MulticallabiCallerSession) GetBlockHash(blockNumber *big.Int) ([32]byte, error) {
	return _Multicallabi.Contract.GetBlockHash(&_Multicallabi.CallOpts, blockNumber)
}

// GetCurrentBlockCoinbase is a free data retrieval call binding the contract method 0xa8b0574e.
//
// Solidity: function getCurrentBlockCoinbase() view returns(address coinbase)
func (_Multicallabi *MulticallabiCaller) GetCurrentBlockCoinbase(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Multicallabi.contract.Call(opts, &out, "getCurrentBlockCoinbase")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetCurrentBlockCoinbase is a free data retrieval call binding the contract method 0xa8b0574e.
//
// Solidity: function getCurrentBlockCoinbase() view returns(address coinbase)
func (_Multicallabi *MulticallabiSession) GetCurrentBlockCoinbase() (common.Address, error) {
	return _Multicallabi.Contract.GetCurrentBlockCoinbase(&_Multicallabi.CallOpts)
}

// GetCurrentBlockCoinbase is a free data retrieval call binding the contract method 0xa8b0574e.
//
// Solidity: function getCurrentBlockCoinbase() view returns(address coinbase)
func (_Multicallabi *MulticallabiCallerSession) GetCurrentBlockCoinbase() (common.Address, error) {
	return _Multicallabi.Contract.GetCurrentBlockCoinbase(&_Multicallabi.CallOpts)
}

// GetCurrentBlockDifficulty is a free data retrieval call binding the contract method 0x72425d9d.
//
// Solidity: function getCurrentBlockDifficulty() view returns(uint256 difficulty)
func (_Multicallabi *MulticallabiCaller) GetCurrentBlockDifficulty(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Multicallabi.contract.Call(opts, &out, "getCurrentBlockDifficulty")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetCurrentBlockDifficulty is a free data retrieval call binding the contract method 0x72425d9d.
//
// Solidity: function getCurrentBlockDifficulty() view returns(uint256 difficulty)
func (_Multicallabi *MulticallabiSession) GetCurrentBlockDifficulty() (*big.Int, error) {
	return _Multicallabi.Contract.GetCurrentBlockDifficulty(&_Multicallabi.CallOpts)
}

// GetCurrentBlockDifficulty is a free data retrieval call binding the contract method 0x72425d9d.
//
// Solidity: function getCurrentBlockDifficulty() view returns(uint256 difficulty)
func (_Multicallabi *MulticallabiCallerSession) GetCurrentBlockDifficulty() (*big.Int, error) {
	return _Multicallabi.Contract.GetCurrentBlockDifficulty(&_Multicallabi.CallOpts)
}

// GetCurrentBlockGasLimit is a free data retrieval call binding the contract method 0x86d516e8.
//
// Solidity: function getCurrentBlockGasLimit() view returns(uint256 gaslimit)
func (_Multicallabi *MulticallabiCaller) GetCurrentBlockGasLimit(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Multicallabi.contract.Call(opts, &out, "getCurrentBlockGasLimit")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetCurrentBlockGasLimit is a free data retrieval call binding the contract method 0x86d516e8.
//
// Solidity: function getCurrentBlockGasLimit() view returns(uint256 gaslimit)
func (_Multicallabi *MulticallabiSession) GetCurrentBlockGasLimit() (*big.Int, error) {
	return _Multicallabi.Contract.GetCurrentBlockGasLimit(&_Multicallabi.CallOpts)
}

// GetCurrentBlockGasLimit is a free data retrieval call binding the contract method 0x86d516e8.
//
// Solidity: function getCurrentBlockGasLimit() view returns(uint256 gaslimit)
func (_Multicallabi *MulticallabiCallerSession) GetCurrentBlockGasLimit() (*big.Int, error) {
	return _Multicallabi.Contract.GetCurrentBlockGasLimit(&_Multicallabi.CallOpts)
}

// GetCurrentBlockTimestamp is a free data retrieval call binding the contract method 0x0f28c97d.
//
// Solidity: function getCurrentBlockTimestamp() view returns(uint256 timestamp)
func (_Multicallabi *MulticallabiCaller) GetCurrentBlockTimestamp(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Multicallabi.contract.Call(opts, &out, "getCurrentBlockTimestamp")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetCurrentBlockTimestamp is a free data retrieval call binding the contract method 0x0f28c97d.
//
// Solidity: function getCurrentBlockTimestamp() view returns(uint256 timestamp)
func (_Multicallabi *MulticallabiSession) GetCurrentBlockTimestamp() (*big.Int, error) {
	return _Multicallabi.Contract.GetCurrentBlockTimestamp(&_Multicallabi.CallOpts)
}

// GetCurrentBlockTimestamp is a free data retrieval call binding the contract method 0x0f28c97d.
//
// Solidity: function getCurrentBlockTimestamp() view returns(uint256 timestamp)
func (_Multicallabi *MulticallabiCallerSession) GetCurrentBlockTimestamp() (*big.Int, error) {
	return _Multicallabi.Contract.GetCurrentBlockTimestamp(&_Multicallabi.CallOpts)
}

// GetEthBalance is a free data retrieval call binding the contract method 0x4d2301cc.
//
// Solidity: function getEthBalance(address addr) view returns(uint256 balance)
func (_Multicallabi *MulticallabiCaller) GetEthBalance(opts *bind.CallOpts, addr common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Multicallabi.contract.Call(opts, &out, "getEthBalance", addr)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetEthBalance is a free data retrieval call binding the contract method 0x4d2301cc.
//
// Solidity: function getEthBalance(address addr) view returns(uint256 balance)
func (_Multicallabi *MulticallabiSession) GetEthBalance(addr common.Address) (*big.Int, error) {
	return _Multicallabi.Contract.GetEthBalance(&_Multicallabi.CallOpts, addr)
}

// GetEthBalance is a free data retrieval call binding the contract method 0x4d2301cc.
//
// Solidity: function getEthBalance(address addr) view returns(uint256 balance)
func (_Multicallabi *MulticallabiCallerSession) GetEthBalance(addr common.Address) (*big.Int, error) {
	return _Multicallabi.Contract.GetEthBalance(&_Multicallabi.CallOpts, addr)
}

// GetLastBlockHash is a free data retrieval call binding the contract method 0x27e86d6e.
//
// Solidity: function getLastBlockHash() view returns(bytes32 blockHash)
func (_Multicallabi *MulticallabiCaller) GetLastBlockHash(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Multicallabi.contract.Call(opts, &out, "getLastBlockHash")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetLastBlockHash is a free data retrieval call binding the contract method 0x27e86d6e.
//
// Solidity: function getLastBlockHash() view returns(bytes32 blockHash)
func (_Multicallabi *MulticallabiSession) GetLastBlockHash() ([32]byte, error) {
	return _Multicallabi.Contract.GetLastBlockHash(&_Multicallabi.CallOpts)
}

// GetLastBlockHash is a free data retrieval call binding the contract method 0x27e86d6e.
//
// Solidity: function getLastBlockHash() view returns(bytes32 blockHash)
func (_Multicallabi *MulticallabiCallerSession) GetLastBlockHash() ([32]byte, error) {
	return _Multicallabi.Contract.GetLastBlockHash(&_Multicallabi.CallOpts)
}

// Aggregate is a paid mutator transaction binding the contract method 0x252dba42.
//
// Solidity: function aggregate((address,bytes)[] calls) returns(uint256 blockNumber, bytes[] returnData)
func (_Multicallabi *MulticallabiTransactor) Aggregate(opts *bind.TransactOpts, calls []Struct0) (*types.Transaction, error) {
	return _Multicallabi.contract.Transact(opts, "aggregate", calls)
}

// Aggregate is a paid mutator transaction binding the contract method 0x252dba42.
//
// Solidity: function aggregate((address,bytes)[] calls) returns(uint256 blockNumber, bytes[] returnData)
func (_Multicallabi *MulticallabiSession) Aggregate(calls []Struct0) (*types.Transaction, error) {
	return _Multicallabi.Contract.Aggregate(&_Multicallabi.TransactOpts, calls)
}

// Aggregate is a paid mutator transaction binding the contract method 0x252dba42.
//
// Solidity: function aggregate((address,bytes)[] calls) returns(uint256 blockNumber, bytes[] returnData)
func (_Multicallabi *MulticallabiTransactorSession) Aggregate(calls []Struct0) (*types.Transaction, error) {
	return _Multicallabi.Contract.Aggregate(&_Multicallabi.TransactOpts, calls)
}
