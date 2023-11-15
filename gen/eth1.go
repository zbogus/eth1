// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package eth1

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
	_ = abi.ConvertType
)

// Eth1MetaData contains all meta data concerning the Eth1 contract.
var Eth1MetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"getBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferEth\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
	Bin: "0x608060405234801561000f575f80fd5b5061028a8061001d5f395ff3fe60806040526004361061002c575f3560e01c8063e9bb84c214610037578063f8b2cb4f1461005357610033565b3661003357005b5f80fd5b610051600480360381019061004c9190610188565b61008f565b005b34801561005e575f80fd5b5061007960048036038101906100749190610201565b6100d7565b604051610086919061023b565b60405180910390f35b8173ffffffffffffffffffffffffffffffffffffffff166108fc8290811502906040515f60405180830381858888f193505050501580156100d2573d5f803e3d5ffd5b505050565b5f8173ffffffffffffffffffffffffffffffffffffffff16319050919050565b5f80fd5b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f610124826100fb565b9050919050565b6101348161011a565b811461013e575f80fd5b50565b5f8135905061014f8161012b565b92915050565b5f819050919050565b61016781610155565b8114610171575f80fd5b50565b5f813590506101828161015e565b92915050565b5f806040838503121561019e5761019d6100f7565b5b5f6101ab85828601610141565b92505060206101bc85828601610174565b9150509250929050565b5f6101d0826100fb565b9050919050565b6101e0816101c6565b81146101ea575f80fd5b50565b5f813590506101fb816101d7565b92915050565b5f60208284031215610216576102156100f7565b5b5f610223848285016101ed565b91505092915050565b61023581610155565b82525050565b5f60208201905061024e5f83018461022c565b9291505056fea26469706673582212203e1f61df5f88306432a6d5d3a5ff8453bc762664dde726fd7efdf674a427b2bd64736f6c63430008150033",
}

// Eth1ABI is the input ABI used to generate the binding from.
// Deprecated: Use Eth1MetaData.ABI instead.
var Eth1ABI = Eth1MetaData.ABI

// Eth1Bin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use Eth1MetaData.Bin instead.
var Eth1Bin = Eth1MetaData.Bin

// DeployEth1 deploys a new Ethereum contract, binding an instance of Eth1 to it.
func DeployEth1(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Eth1, error) {
	parsed, err := Eth1MetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(Eth1Bin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Eth1{Eth1Caller: Eth1Caller{contract: contract}, Eth1Transactor: Eth1Transactor{contract: contract}, Eth1Filterer: Eth1Filterer{contract: contract}}, nil
}

// Eth1 is an auto generated Go binding around an Ethereum contract.
type Eth1 struct {
	Eth1Caller     // Read-only binding to the contract
	Eth1Transactor // Write-only binding to the contract
	Eth1Filterer   // Log filterer for contract events
}

// Eth1Caller is an auto generated read-only Go binding around an Ethereum contract.
type Eth1Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Eth1Transactor is an auto generated write-only Go binding around an Ethereum contract.
type Eth1Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Eth1Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type Eth1Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Eth1Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type Eth1Session struct {
	Contract     *Eth1             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// Eth1CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type Eth1CallerSession struct {
	Contract *Eth1Caller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// Eth1TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type Eth1TransactorSession struct {
	Contract     *Eth1Transactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// Eth1Raw is an auto generated low-level Go binding around an Ethereum contract.
type Eth1Raw struct {
	Contract *Eth1 // Generic contract binding to access the raw methods on
}

// Eth1CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type Eth1CallerRaw struct {
	Contract *Eth1Caller // Generic read-only contract binding to access the raw methods on
}

// Eth1TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type Eth1TransactorRaw struct {
	Contract *Eth1Transactor // Generic write-only contract binding to access the raw methods on
}

// NewEth1 creates a new instance of Eth1, bound to a specific deployed contract.
func NewEth1(address common.Address, backend bind.ContractBackend) (*Eth1, error) {
	contract, err := bindEth1(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Eth1{Eth1Caller: Eth1Caller{contract: contract}, Eth1Transactor: Eth1Transactor{contract: contract}, Eth1Filterer: Eth1Filterer{contract: contract}}, nil
}

// NewEth1Caller creates a new read-only instance of Eth1, bound to a specific deployed contract.
func NewEth1Caller(address common.Address, caller bind.ContractCaller) (*Eth1Caller, error) {
	contract, err := bindEth1(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &Eth1Caller{contract: contract}, nil
}

// NewEth1Transactor creates a new write-only instance of Eth1, bound to a specific deployed contract.
func NewEth1Transactor(address common.Address, transactor bind.ContractTransactor) (*Eth1Transactor, error) {
	contract, err := bindEth1(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &Eth1Transactor{contract: contract}, nil
}

// NewEth1Filterer creates a new log filterer instance of Eth1, bound to a specific deployed contract.
func NewEth1Filterer(address common.Address, filterer bind.ContractFilterer) (*Eth1Filterer, error) {
	contract, err := bindEth1(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &Eth1Filterer{contract: contract}, nil
}

// bindEth1 binds a generic wrapper to an already deployed contract.
func bindEth1(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := Eth1MetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Eth1 *Eth1Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Eth1.Contract.Eth1Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Eth1 *Eth1Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Eth1.Contract.Eth1Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Eth1 *Eth1Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Eth1.Contract.Eth1Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Eth1 *Eth1CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Eth1.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Eth1 *Eth1TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Eth1.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Eth1 *Eth1TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Eth1.Contract.contract.Transact(opts, method, params...)
}

// GetBalance is a free data retrieval call binding the contract method 0xf8b2cb4f.
//
// Solidity: function getBalance(address account) view returns(uint256)
func (_Eth1 *Eth1Caller) GetBalance(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Eth1.contract.Call(opts, &out, "getBalance", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetBalance is a free data retrieval call binding the contract method 0xf8b2cb4f.
//
// Solidity: function getBalance(address account) view returns(uint256)
func (_Eth1 *Eth1Session) GetBalance(account common.Address) (*big.Int, error) {
	return _Eth1.Contract.GetBalance(&_Eth1.CallOpts, account)
}

// GetBalance is a free data retrieval call binding the contract method 0xf8b2cb4f.
//
// Solidity: function getBalance(address account) view returns(uint256)
func (_Eth1 *Eth1CallerSession) GetBalance(account common.Address) (*big.Int, error) {
	return _Eth1.Contract.GetBalance(&_Eth1.CallOpts, account)
}

// TransferEth is a paid mutator transaction binding the contract method 0xe9bb84c2.
//
// Solidity: function transferEth(address to, uint256 amount) payable returns()
func (_Eth1 *Eth1Transactor) TransferEth(opts *bind.TransactOpts, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Eth1.contract.Transact(opts, "transferEth", to, amount)
}

// TransferEth is a paid mutator transaction binding the contract method 0xe9bb84c2.
//
// Solidity: function transferEth(address to, uint256 amount) payable returns()
func (_Eth1 *Eth1Session) TransferEth(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Eth1.Contract.TransferEth(&_Eth1.TransactOpts, to, amount)
}

// TransferEth is a paid mutator transaction binding the contract method 0xe9bb84c2.
//
// Solidity: function transferEth(address to, uint256 amount) payable returns()
func (_Eth1 *Eth1TransactorSession) TransferEth(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Eth1.Contract.TransferEth(&_Eth1.TransactOpts, to, amount)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Eth1 *Eth1Transactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Eth1.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Eth1 *Eth1Session) Receive() (*types.Transaction, error) {
	return _Eth1.Contract.Receive(&_Eth1.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Eth1 *Eth1TransactorSession) Receive() (*types.Transaction, error) {
	return _Eth1.Contract.Receive(&_Eth1.TransactOpts)
}
