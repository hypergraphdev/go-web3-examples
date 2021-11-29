// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package factory

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

// FactoryMetaData contains all meta data concerning the Factory contract.
var FactoryMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"nounId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"AuctionBid\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"nounId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"startTime\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"endTime\",\"type\":\"uint256\"}],\"name\":\"AuctionCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"nounId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"winner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"AuctionSettled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"auction\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"nounId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"startTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endTime\",\"type\":\"uint256\"},{\"internalType\":\"addresspayable\",\"name\":\"bidder\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"settled\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"avatar\",\"outputs\":[{\"internalType\":\"contractIMetaToken\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_referId\",\"type\":\"uint256\"}],\"name\":\"createBid\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIMetaToken\",\"name\":\"_avatar\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_weth\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_reservePrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_referPrice\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_metaHolder\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"referPrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"reservePrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_reservePrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_referPrice\",\"type\":\"uint256\"}],\"name\":\"setPrice\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"weth\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// FactoryABI is the input ABI used to generate the binding from.
// Deprecated: Use FactoryMetaData.ABI instead.
var FactoryABI = FactoryMetaData.ABI

// Factory is an auto generated Go binding around an Ethereum contract.
type Factory struct {
	FactoryCaller     // Read-only binding to the contract
	FactoryTransactor // Write-only binding to the contract
	FactoryFilterer   // Log filterer for contract events
}

// FactoryCaller is an auto generated read-only Go binding around an Ethereum contract.
type FactoryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FactoryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type FactoryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FactoryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type FactoryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FactorySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type FactorySession struct {
	Contract     *Factory          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// FactoryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type FactoryCallerSession struct {
	Contract *FactoryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// FactoryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type FactoryTransactorSession struct {
	Contract     *FactoryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// FactoryRaw is an auto generated low-level Go binding around an Ethereum contract.
type FactoryRaw struct {
	Contract *Factory // Generic contract binding to access the raw methods on
}

// FactoryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type FactoryCallerRaw struct {
	Contract *FactoryCaller // Generic read-only contract binding to access the raw methods on
}

// FactoryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type FactoryTransactorRaw struct {
	Contract *FactoryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewFactory creates a new instance of Factory, bound to a specific deployed contract.
func NewFactory(address common.Address, backend bind.ContractBackend) (*Factory, error) {
	contract, err := bindFactory(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Factory{FactoryCaller: FactoryCaller{contract: contract}, FactoryTransactor: FactoryTransactor{contract: contract}, FactoryFilterer: FactoryFilterer{contract: contract}}, nil
}

// NewFactoryCaller creates a new read-only instance of Factory, bound to a specific deployed contract.
func NewFactoryCaller(address common.Address, caller bind.ContractCaller) (*FactoryCaller, error) {
	contract, err := bindFactory(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &FactoryCaller{contract: contract}, nil
}

// NewFactoryTransactor creates a new write-only instance of Factory, bound to a specific deployed contract.
func NewFactoryTransactor(address common.Address, transactor bind.ContractTransactor) (*FactoryTransactor, error) {
	contract, err := bindFactory(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &FactoryTransactor{contract: contract}, nil
}

// NewFactoryFilterer creates a new log filterer instance of Factory, bound to a specific deployed contract.
func NewFactoryFilterer(address common.Address, filterer bind.ContractFilterer) (*FactoryFilterer, error) {
	contract, err := bindFactory(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &FactoryFilterer{contract: contract}, nil
}

// bindFactory binds a generic wrapper to an already deployed contract.
func bindFactory(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(FactoryABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Factory *FactoryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Factory.Contract.FactoryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Factory *FactoryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Factory.Contract.FactoryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Factory *FactoryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Factory.Contract.FactoryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Factory *FactoryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Factory.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Factory *FactoryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Factory.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Factory *FactoryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Factory.Contract.contract.Transact(opts, method, params...)
}

// Auction is a free data retrieval call binding the contract method 0x7d9f6db5.
//
// Solidity: function auction() view returns(uint256 nounId, uint256 amount, uint256 startTime, uint256 endTime, address bidder, bool settled)
func (_Factory *FactoryCaller) Auction(opts *bind.CallOpts) (struct {
	NounId    *big.Int
	Amount    *big.Int
	StartTime *big.Int
	EndTime   *big.Int
	Bidder    common.Address
	Settled   bool
}, error) {
	var out []interface{}
	err := _Factory.contract.Call(opts, &out, "auction")

	outstruct := new(struct {
		NounId    *big.Int
		Amount    *big.Int
		StartTime *big.Int
		EndTime   *big.Int
		Bidder    common.Address
		Settled   bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.NounId = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Amount = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.StartTime = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.EndTime = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.Bidder = *abi.ConvertType(out[4], new(common.Address)).(*common.Address)
	outstruct.Settled = *abi.ConvertType(out[5], new(bool)).(*bool)

	return *outstruct, err

}

// Auction is a free data retrieval call binding the contract method 0x7d9f6db5.
//
// Solidity: function auction() view returns(uint256 nounId, uint256 amount, uint256 startTime, uint256 endTime, address bidder, bool settled)
func (_Factory *FactorySession) Auction() (struct {
	NounId    *big.Int
	Amount    *big.Int
	StartTime *big.Int
	EndTime   *big.Int
	Bidder    common.Address
	Settled   bool
}, error) {
	return _Factory.Contract.Auction(&_Factory.CallOpts)
}

// Auction is a free data retrieval call binding the contract method 0x7d9f6db5.
//
// Solidity: function auction() view returns(uint256 nounId, uint256 amount, uint256 startTime, uint256 endTime, address bidder, bool settled)
func (_Factory *FactoryCallerSession) Auction() (struct {
	NounId    *big.Int
	Amount    *big.Int
	StartTime *big.Int
	EndTime   *big.Int
	Bidder    common.Address
	Settled   bool
}, error) {
	return _Factory.Contract.Auction(&_Factory.CallOpts)
}

// Avatar is a free data retrieval call binding the contract method 0x5aef7de6.
//
// Solidity: function avatar() view returns(address)
func (_Factory *FactoryCaller) Avatar(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Factory.contract.Call(opts, &out, "avatar")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Avatar is a free data retrieval call binding the contract method 0x5aef7de6.
//
// Solidity: function avatar() view returns(address)
func (_Factory *FactorySession) Avatar() (common.Address, error) {
	return _Factory.Contract.Avatar(&_Factory.CallOpts)
}

// Avatar is a free data retrieval call binding the contract method 0x5aef7de6.
//
// Solidity: function avatar() view returns(address)
func (_Factory *FactoryCallerSession) Avatar() (common.Address, error) {
	return _Factory.Contract.Avatar(&_Factory.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Factory *FactoryCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Factory.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Factory *FactorySession) Owner() (common.Address, error) {
	return _Factory.Contract.Owner(&_Factory.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Factory *FactoryCallerSession) Owner() (common.Address, error) {
	return _Factory.Contract.Owner(&_Factory.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Factory *FactoryCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Factory.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Factory *FactorySession) Paused() (bool, error) {
	return _Factory.Contract.Paused(&_Factory.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Factory *FactoryCallerSession) Paused() (bool, error) {
	return _Factory.Contract.Paused(&_Factory.CallOpts)
}

// ReferPrice is a free data retrieval call binding the contract method 0x6f7b0bab.
//
// Solidity: function referPrice() view returns(uint256)
func (_Factory *FactoryCaller) ReferPrice(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Factory.contract.Call(opts, &out, "referPrice")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ReferPrice is a free data retrieval call binding the contract method 0x6f7b0bab.
//
// Solidity: function referPrice() view returns(uint256)
func (_Factory *FactorySession) ReferPrice() (*big.Int, error) {
	return _Factory.Contract.ReferPrice(&_Factory.CallOpts)
}

// ReferPrice is a free data retrieval call binding the contract method 0x6f7b0bab.
//
// Solidity: function referPrice() view returns(uint256)
func (_Factory *FactoryCallerSession) ReferPrice() (*big.Int, error) {
	return _Factory.Contract.ReferPrice(&_Factory.CallOpts)
}

// ReservePrice is a free data retrieval call binding the contract method 0xdb2e1eed.
//
// Solidity: function reservePrice() view returns(uint256)
func (_Factory *FactoryCaller) ReservePrice(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Factory.contract.Call(opts, &out, "reservePrice")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ReservePrice is a free data retrieval call binding the contract method 0xdb2e1eed.
//
// Solidity: function reservePrice() view returns(uint256)
func (_Factory *FactorySession) ReservePrice() (*big.Int, error) {
	return _Factory.Contract.ReservePrice(&_Factory.CallOpts)
}

// ReservePrice is a free data retrieval call binding the contract method 0xdb2e1eed.
//
// Solidity: function reservePrice() view returns(uint256)
func (_Factory *FactoryCallerSession) ReservePrice() (*big.Int, error) {
	return _Factory.Contract.ReservePrice(&_Factory.CallOpts)
}

// Weth is a free data retrieval call binding the contract method 0x3fc8cef3.
//
// Solidity: function weth() view returns(address)
func (_Factory *FactoryCaller) Weth(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Factory.contract.Call(opts, &out, "weth")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Weth is a free data retrieval call binding the contract method 0x3fc8cef3.
//
// Solidity: function weth() view returns(address)
func (_Factory *FactorySession) Weth() (common.Address, error) {
	return _Factory.Contract.Weth(&_Factory.CallOpts)
}

// Weth is a free data retrieval call binding the contract method 0x3fc8cef3.
//
// Solidity: function weth() view returns(address)
func (_Factory *FactoryCallerSession) Weth() (common.Address, error) {
	return _Factory.Contract.Weth(&_Factory.CallOpts)
}

// CreateBid is a paid mutator transaction binding the contract method 0x659dd2b4.
//
// Solidity: function createBid(uint256 _referId) payable returns()
func (_Factory *FactoryTransactor) CreateBid(opts *bind.TransactOpts, _referId *big.Int) (*types.Transaction, error) {
	return _Factory.contract.Transact(opts, "createBid", _referId)
}

// CreateBid is a paid mutator transaction binding the contract method 0x659dd2b4.
//
// Solidity: function createBid(uint256 _referId) payable returns()
func (_Factory *FactorySession) CreateBid(_referId *big.Int) (*types.Transaction, error) {
	return _Factory.Contract.CreateBid(&_Factory.TransactOpts, _referId)
}

// CreateBid is a paid mutator transaction binding the contract method 0x659dd2b4.
//
// Solidity: function createBid(uint256 _referId) payable returns()
func (_Factory *FactoryTransactorSession) CreateBid(_referId *big.Int) (*types.Transaction, error) {
	return _Factory.Contract.CreateBid(&_Factory.TransactOpts, _referId)
}

// Initialize is a paid mutator transaction binding the contract method 0x57fb25cc.
//
// Solidity: function initialize(address _avatar, address _weth, uint256 _reservePrice, uint256 _referPrice, address _metaHolder) returns()
func (_Factory *FactoryTransactor) Initialize(opts *bind.TransactOpts, _avatar common.Address, _weth common.Address, _reservePrice *big.Int, _referPrice *big.Int, _metaHolder common.Address) (*types.Transaction, error) {
	return _Factory.contract.Transact(opts, "initialize", _avatar, _weth, _reservePrice, _referPrice, _metaHolder)
}

// Initialize is a paid mutator transaction binding the contract method 0x57fb25cc.
//
// Solidity: function initialize(address _avatar, address _weth, uint256 _reservePrice, uint256 _referPrice, address _metaHolder) returns()
func (_Factory *FactorySession) Initialize(_avatar common.Address, _weth common.Address, _reservePrice *big.Int, _referPrice *big.Int, _metaHolder common.Address) (*types.Transaction, error) {
	return _Factory.Contract.Initialize(&_Factory.TransactOpts, _avatar, _weth, _reservePrice, _referPrice, _metaHolder)
}

// Initialize is a paid mutator transaction binding the contract method 0x57fb25cc.
//
// Solidity: function initialize(address _avatar, address _weth, uint256 _reservePrice, uint256 _referPrice, address _metaHolder) returns()
func (_Factory *FactoryTransactorSession) Initialize(_avatar common.Address, _weth common.Address, _reservePrice *big.Int, _referPrice *big.Int, _metaHolder common.Address) (*types.Transaction, error) {
	return _Factory.Contract.Initialize(&_Factory.TransactOpts, _avatar, _weth, _reservePrice, _referPrice, _metaHolder)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Factory *FactoryTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Factory.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Factory *FactorySession) Pause() (*types.Transaction, error) {
	return _Factory.Contract.Pause(&_Factory.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Factory *FactoryTransactorSession) Pause() (*types.Transaction, error) {
	return _Factory.Contract.Pause(&_Factory.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Factory *FactoryTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Factory.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Factory *FactorySession) RenounceOwnership() (*types.Transaction, error) {
	return _Factory.Contract.RenounceOwnership(&_Factory.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Factory *FactoryTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Factory.Contract.RenounceOwnership(&_Factory.TransactOpts)
}

// SetPrice is a paid mutator transaction binding the contract method 0xf7d97577.
//
// Solidity: function setPrice(uint256 _reservePrice, uint256 _referPrice) returns()
func (_Factory *FactoryTransactor) SetPrice(opts *bind.TransactOpts, _reservePrice *big.Int, _referPrice *big.Int) (*types.Transaction, error) {
	return _Factory.contract.Transact(opts, "setPrice", _reservePrice, _referPrice)
}

// SetPrice is a paid mutator transaction binding the contract method 0xf7d97577.
//
// Solidity: function setPrice(uint256 _reservePrice, uint256 _referPrice) returns()
func (_Factory *FactorySession) SetPrice(_reservePrice *big.Int, _referPrice *big.Int) (*types.Transaction, error) {
	return _Factory.Contract.SetPrice(&_Factory.TransactOpts, _reservePrice, _referPrice)
}

// SetPrice is a paid mutator transaction binding the contract method 0xf7d97577.
//
// Solidity: function setPrice(uint256 _reservePrice, uint256 _referPrice) returns()
func (_Factory *FactoryTransactorSession) SetPrice(_reservePrice *big.Int, _referPrice *big.Int) (*types.Transaction, error) {
	return _Factory.Contract.SetPrice(&_Factory.TransactOpts, _reservePrice, _referPrice)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Factory *FactoryTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Factory.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Factory *FactorySession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Factory.Contract.TransferOwnership(&_Factory.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Factory *FactoryTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Factory.Contract.TransferOwnership(&_Factory.TransactOpts, newOwner)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Factory *FactoryTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Factory.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Factory *FactorySession) Unpause() (*types.Transaction, error) {
	return _Factory.Contract.Unpause(&_Factory.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Factory *FactoryTransactorSession) Unpause() (*types.Transaction, error) {
	return _Factory.Contract.Unpause(&_Factory.TransactOpts)
}

// FactoryAuctionBidIterator is returned from FilterAuctionBid and is used to iterate over the raw logs and unpacked data for AuctionBid events raised by the Factory contract.
type FactoryAuctionBidIterator struct {
	Event *FactoryAuctionBid // Event containing the contract specifics and raw log

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
func (it *FactoryAuctionBidIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FactoryAuctionBid)
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
		it.Event = new(FactoryAuctionBid)
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
func (it *FactoryAuctionBidIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FactoryAuctionBidIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FactoryAuctionBid represents a AuctionBid event raised by the Factory contract.
type FactoryAuctionBid struct {
	NounId *big.Int
	Sender common.Address
	Value  *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterAuctionBid is a free log retrieval operation binding the contract event 0xb8d756f2d1da4663767eb4d559780ace84f2f65a421a60ddcb47e8b2e5d2fd26.
//
// Solidity: event AuctionBid(uint256 indexed nounId, address sender, uint256 value)
func (_Factory *FactoryFilterer) FilterAuctionBid(opts *bind.FilterOpts, nounId []*big.Int) (*FactoryAuctionBidIterator, error) {

	var nounIdRule []interface{}
	for _, nounIdItem := range nounId {
		nounIdRule = append(nounIdRule, nounIdItem)
	}

	logs, sub, err := _Factory.contract.FilterLogs(opts, "AuctionBid", nounIdRule)
	if err != nil {
		return nil, err
	}
	return &FactoryAuctionBidIterator{contract: _Factory.contract, event: "AuctionBid", logs: logs, sub: sub}, nil
}

// WatchAuctionBid is a free log subscription operation binding the contract event 0xb8d756f2d1da4663767eb4d559780ace84f2f65a421a60ddcb47e8b2e5d2fd26.
//
// Solidity: event AuctionBid(uint256 indexed nounId, address sender, uint256 value)
func (_Factory *FactoryFilterer) WatchAuctionBid(opts *bind.WatchOpts, sink chan<- *FactoryAuctionBid, nounId []*big.Int) (event.Subscription, error) {

	var nounIdRule []interface{}
	for _, nounIdItem := range nounId {
		nounIdRule = append(nounIdRule, nounIdItem)
	}

	logs, sub, err := _Factory.contract.WatchLogs(opts, "AuctionBid", nounIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FactoryAuctionBid)
				if err := _Factory.contract.UnpackLog(event, "AuctionBid", log); err != nil {
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

// ParseAuctionBid is a log parse operation binding the contract event 0xb8d756f2d1da4663767eb4d559780ace84f2f65a421a60ddcb47e8b2e5d2fd26.
//
// Solidity: event AuctionBid(uint256 indexed nounId, address sender, uint256 value)
func (_Factory *FactoryFilterer) ParseAuctionBid(log types.Log) (*FactoryAuctionBid, error) {
	event := new(FactoryAuctionBid)
	if err := _Factory.contract.UnpackLog(event, "AuctionBid", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FactoryAuctionCreatedIterator is returned from FilterAuctionCreated and is used to iterate over the raw logs and unpacked data for AuctionCreated events raised by the Factory contract.
type FactoryAuctionCreatedIterator struct {
	Event *FactoryAuctionCreated // Event containing the contract specifics and raw log

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
func (it *FactoryAuctionCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FactoryAuctionCreated)
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
		it.Event = new(FactoryAuctionCreated)
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
func (it *FactoryAuctionCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FactoryAuctionCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FactoryAuctionCreated represents a AuctionCreated event raised by the Factory contract.
type FactoryAuctionCreated struct {
	NounId    *big.Int
	StartTime *big.Int
	EndTime   *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterAuctionCreated is a free log retrieval operation binding the contract event 0xd6eddd1118d71820909c1197aa966dbc15ed6f508554252169cc3d5ccac756ca.
//
// Solidity: event AuctionCreated(uint256 indexed nounId, uint256 startTime, uint256 endTime)
func (_Factory *FactoryFilterer) FilterAuctionCreated(opts *bind.FilterOpts, nounId []*big.Int) (*FactoryAuctionCreatedIterator, error) {

	var nounIdRule []interface{}
	for _, nounIdItem := range nounId {
		nounIdRule = append(nounIdRule, nounIdItem)
	}

	logs, sub, err := _Factory.contract.FilterLogs(opts, "AuctionCreated", nounIdRule)
	if err != nil {
		return nil, err
	}
	return &FactoryAuctionCreatedIterator{contract: _Factory.contract, event: "AuctionCreated", logs: logs, sub: sub}, nil
}

// WatchAuctionCreated is a free log subscription operation binding the contract event 0xd6eddd1118d71820909c1197aa966dbc15ed6f508554252169cc3d5ccac756ca.
//
// Solidity: event AuctionCreated(uint256 indexed nounId, uint256 startTime, uint256 endTime)
func (_Factory *FactoryFilterer) WatchAuctionCreated(opts *bind.WatchOpts, sink chan<- *FactoryAuctionCreated, nounId []*big.Int) (event.Subscription, error) {

	var nounIdRule []interface{}
	for _, nounIdItem := range nounId {
		nounIdRule = append(nounIdRule, nounIdItem)
	}

	logs, sub, err := _Factory.contract.WatchLogs(opts, "AuctionCreated", nounIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FactoryAuctionCreated)
				if err := _Factory.contract.UnpackLog(event, "AuctionCreated", log); err != nil {
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

// ParseAuctionCreated is a log parse operation binding the contract event 0xd6eddd1118d71820909c1197aa966dbc15ed6f508554252169cc3d5ccac756ca.
//
// Solidity: event AuctionCreated(uint256 indexed nounId, uint256 startTime, uint256 endTime)
func (_Factory *FactoryFilterer) ParseAuctionCreated(log types.Log) (*FactoryAuctionCreated, error) {
	event := new(FactoryAuctionCreated)
	if err := _Factory.contract.UnpackLog(event, "AuctionCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FactoryAuctionSettledIterator is returned from FilterAuctionSettled and is used to iterate over the raw logs and unpacked data for AuctionSettled events raised by the Factory contract.
type FactoryAuctionSettledIterator struct {
	Event *FactoryAuctionSettled // Event containing the contract specifics and raw log

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
func (it *FactoryAuctionSettledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FactoryAuctionSettled)
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
		it.Event = new(FactoryAuctionSettled)
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
func (it *FactoryAuctionSettledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FactoryAuctionSettledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FactoryAuctionSettled represents a AuctionSettled event raised by the Factory contract.
type FactoryAuctionSettled struct {
	NounId *big.Int
	Winner common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterAuctionSettled is a free log retrieval operation binding the contract event 0xc9f72b276a388619c6d185d146697036241880c36654b1a3ffdad07c24038d99.
//
// Solidity: event AuctionSettled(uint256 indexed nounId, address winner, uint256 amount)
func (_Factory *FactoryFilterer) FilterAuctionSettled(opts *bind.FilterOpts, nounId []*big.Int) (*FactoryAuctionSettledIterator, error) {

	var nounIdRule []interface{}
	for _, nounIdItem := range nounId {
		nounIdRule = append(nounIdRule, nounIdItem)
	}

	logs, sub, err := _Factory.contract.FilterLogs(opts, "AuctionSettled", nounIdRule)
	if err != nil {
		return nil, err
	}
	return &FactoryAuctionSettledIterator{contract: _Factory.contract, event: "AuctionSettled", logs: logs, sub: sub}, nil
}

// WatchAuctionSettled is a free log subscription operation binding the contract event 0xc9f72b276a388619c6d185d146697036241880c36654b1a3ffdad07c24038d99.
//
// Solidity: event AuctionSettled(uint256 indexed nounId, address winner, uint256 amount)
func (_Factory *FactoryFilterer) WatchAuctionSettled(opts *bind.WatchOpts, sink chan<- *FactoryAuctionSettled, nounId []*big.Int) (event.Subscription, error) {

	var nounIdRule []interface{}
	for _, nounIdItem := range nounId {
		nounIdRule = append(nounIdRule, nounIdItem)
	}

	logs, sub, err := _Factory.contract.WatchLogs(opts, "AuctionSettled", nounIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FactoryAuctionSettled)
				if err := _Factory.contract.UnpackLog(event, "AuctionSettled", log); err != nil {
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

// ParseAuctionSettled is a log parse operation binding the contract event 0xc9f72b276a388619c6d185d146697036241880c36654b1a3ffdad07c24038d99.
//
// Solidity: event AuctionSettled(uint256 indexed nounId, address winner, uint256 amount)
func (_Factory *FactoryFilterer) ParseAuctionSettled(log types.Log) (*FactoryAuctionSettled, error) {
	event := new(FactoryAuctionSettled)
	if err := _Factory.contract.UnpackLog(event, "AuctionSettled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FactoryOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Factory contract.
type FactoryOwnershipTransferredIterator struct {
	Event *FactoryOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *FactoryOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FactoryOwnershipTransferred)
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
		it.Event = new(FactoryOwnershipTransferred)
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
func (it *FactoryOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FactoryOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FactoryOwnershipTransferred represents a OwnershipTransferred event raised by the Factory contract.
type FactoryOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Factory *FactoryFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*FactoryOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Factory.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &FactoryOwnershipTransferredIterator{contract: _Factory.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Factory *FactoryFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *FactoryOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Factory.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FactoryOwnershipTransferred)
				if err := _Factory.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Factory *FactoryFilterer) ParseOwnershipTransferred(log types.Log) (*FactoryOwnershipTransferred, error) {
	event := new(FactoryOwnershipTransferred)
	if err := _Factory.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FactoryPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the Factory contract.
type FactoryPausedIterator struct {
	Event *FactoryPaused // Event containing the contract specifics and raw log

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
func (it *FactoryPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FactoryPaused)
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
		it.Event = new(FactoryPaused)
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
func (it *FactoryPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FactoryPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FactoryPaused represents a Paused event raised by the Factory contract.
type FactoryPaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_Factory *FactoryFilterer) FilterPaused(opts *bind.FilterOpts) (*FactoryPausedIterator, error) {

	logs, sub, err := _Factory.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &FactoryPausedIterator{contract: _Factory.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_Factory *FactoryFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *FactoryPaused) (event.Subscription, error) {

	logs, sub, err := _Factory.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FactoryPaused)
				if err := _Factory.contract.UnpackLog(event, "Paused", log); err != nil {
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

// ParsePaused is a log parse operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_Factory *FactoryFilterer) ParsePaused(log types.Log) (*FactoryPaused, error) {
	event := new(FactoryPaused)
	if err := _Factory.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FactoryUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the Factory contract.
type FactoryUnpausedIterator struct {
	Event *FactoryUnpaused // Event containing the contract specifics and raw log

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
func (it *FactoryUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FactoryUnpaused)
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
		it.Event = new(FactoryUnpaused)
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
func (it *FactoryUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FactoryUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FactoryUnpaused represents a Unpaused event raised by the Factory contract.
type FactoryUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_Factory *FactoryFilterer) FilterUnpaused(opts *bind.FilterOpts) (*FactoryUnpausedIterator, error) {

	logs, sub, err := _Factory.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &FactoryUnpausedIterator{contract: _Factory.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_Factory *FactoryFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *FactoryUnpaused) (event.Subscription, error) {

	logs, sub, err := _Factory.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FactoryUnpaused)
				if err := _Factory.contract.UnpackLog(event, "Unpaused", log); err != nil {
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

// ParseUnpaused is a log parse operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_Factory *FactoryFilterer) ParseUnpaused(log types.Log) (*FactoryUnpaused, error) {
	event := new(FactoryUnpaused)
	if err := _Factory.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
