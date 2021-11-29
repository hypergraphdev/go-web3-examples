// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package avatar

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

// AvatarMetaData contains all meta data concerning the Avatar contract.
var AvatarMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"nounId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"AuctionBid\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"nounId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"startTime\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"endTime\",\"type\":\"uint256\"}],\"name\":\"AuctionCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"nounId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"winner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"AuctionSettled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"auction\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"nounId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"startTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endTime\",\"type\":\"uint256\"},{\"internalType\":\"addresspayable\",\"name\":\"bidder\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"settled\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"avatar\",\"outputs\":[{\"internalType\":\"contractIMetaToken\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_referId\",\"type\":\"uint256\"}],\"name\":\"createBid\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIMetaToken\",\"name\":\"_avatar\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_weth\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_reservePrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_referPrice\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_metaHolder\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"referPrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"reservePrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_reservePrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_referPrice\",\"type\":\"uint256\"}],\"name\":\"setPrice\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"weth\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// AvatarABI is the input ABI used to generate the binding from.
// Deprecated: Use AvatarMetaData.ABI instead.
var AvatarABI = AvatarMetaData.ABI

// Avatar is an auto generated Go binding around an Ethereum contract.
type Avatar struct {
	AvatarCaller     // Read-only binding to the contract
	AvatarTransactor // Write-only binding to the contract
	AvatarFilterer   // Log filterer for contract events
}

// AvatarCaller is an auto generated read-only Go binding around an Ethereum contract.
type AvatarCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AvatarTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AvatarTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AvatarFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AvatarFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AvatarSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AvatarSession struct {
	Contract     *Avatar           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AvatarCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AvatarCallerSession struct {
	Contract *AvatarCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// AvatarTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AvatarTransactorSession struct {
	Contract     *AvatarTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AvatarRaw is an auto generated low-level Go binding around an Ethereum contract.
type AvatarRaw struct {
	Contract *Avatar // Generic contract binding to access the raw methods on
}

// AvatarCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AvatarCallerRaw struct {
	Contract *AvatarCaller // Generic read-only contract binding to access the raw methods on
}

// AvatarTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AvatarTransactorRaw struct {
	Contract *AvatarTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAvatar creates a new instance of Avatar, bound to a specific deployed contract.
func NewAvatar(address common.Address, backend bind.ContractBackend) (*Avatar, error) {
	contract, err := bindAvatar(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Avatar{AvatarCaller: AvatarCaller{contract: contract}, AvatarTransactor: AvatarTransactor{contract: contract}, AvatarFilterer: AvatarFilterer{contract: contract}}, nil
}

// NewAvatarCaller creates a new read-only instance of Avatar, bound to a specific deployed contract.
func NewAvatarCaller(address common.Address, caller bind.ContractCaller) (*AvatarCaller, error) {
	contract, err := bindAvatar(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AvatarCaller{contract: contract}, nil
}

// NewAvatarTransactor creates a new write-only instance of Avatar, bound to a specific deployed contract.
func NewAvatarTransactor(address common.Address, transactor bind.ContractTransactor) (*AvatarTransactor, error) {
	contract, err := bindAvatar(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AvatarTransactor{contract: contract}, nil
}

// NewAvatarFilterer creates a new log filterer instance of Avatar, bound to a specific deployed contract.
func NewAvatarFilterer(address common.Address, filterer bind.ContractFilterer) (*AvatarFilterer, error) {
	contract, err := bindAvatar(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AvatarFilterer{contract: contract}, nil
}

// bindAvatar binds a generic wrapper to an already deployed contract.
func bindAvatar(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(AvatarABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Avatar *AvatarRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Avatar.Contract.AvatarCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Avatar *AvatarRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Avatar.Contract.AvatarTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Avatar *AvatarRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Avatar.Contract.AvatarTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Avatar *AvatarCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Avatar.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Avatar *AvatarTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Avatar.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Avatar *AvatarTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Avatar.Contract.contract.Transact(opts, method, params...)
}

// Auction is a free data retrieval call binding the contract method 0x7d9f6db5.
//
// Solidity: function auction() view returns(uint256 nounId, uint256 amount, uint256 startTime, uint256 endTime, address bidder, bool settled)
func (_Avatar *AvatarCaller) Auction(opts *bind.CallOpts) (struct {
	NounId    *big.Int
	Amount    *big.Int
	StartTime *big.Int
	EndTime   *big.Int
	Bidder    common.Address
	Settled   bool
}, error) {
	var out []interface{}
	err := _Avatar.contract.Call(opts, &out, "auction")

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
func (_Avatar *AvatarSession) Auction() (struct {
	NounId    *big.Int
	Amount    *big.Int
	StartTime *big.Int
	EndTime   *big.Int
	Bidder    common.Address
	Settled   bool
}, error) {
	return _Avatar.Contract.Auction(&_Avatar.CallOpts)
}

// Auction is a free data retrieval call binding the contract method 0x7d9f6db5.
//
// Solidity: function auction() view returns(uint256 nounId, uint256 amount, uint256 startTime, uint256 endTime, address bidder, bool settled)
func (_Avatar *AvatarCallerSession) Auction() (struct {
	NounId    *big.Int
	Amount    *big.Int
	StartTime *big.Int
	EndTime   *big.Int
	Bidder    common.Address
	Settled   bool
}, error) {
	return _Avatar.Contract.Auction(&_Avatar.CallOpts)
}

// Avatar is a free data retrieval call binding the contract method 0x5aef7de6.
//
// Solidity: function avatar() view returns(address)
func (_Avatar *AvatarCaller) Avatar(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Avatar.contract.Call(opts, &out, "avatar")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Avatar is a free data retrieval call binding the contract method 0x5aef7de6.
//
// Solidity: function avatar() view returns(address)
func (_Avatar *AvatarSession) Avatar() (common.Address, error) {
	return _Avatar.Contract.Avatar(&_Avatar.CallOpts)
}

// Avatar is a free data retrieval call binding the contract method 0x5aef7de6.
//
// Solidity: function avatar() view returns(address)
func (_Avatar *AvatarCallerSession) Avatar() (common.Address, error) {
	return _Avatar.Contract.Avatar(&_Avatar.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Avatar *AvatarCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Avatar.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Avatar *AvatarSession) Owner() (common.Address, error) {
	return _Avatar.Contract.Owner(&_Avatar.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Avatar *AvatarCallerSession) Owner() (common.Address, error) {
	return _Avatar.Contract.Owner(&_Avatar.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Avatar *AvatarCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Avatar.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Avatar *AvatarSession) Paused() (bool, error) {
	return _Avatar.Contract.Paused(&_Avatar.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Avatar *AvatarCallerSession) Paused() (bool, error) {
	return _Avatar.Contract.Paused(&_Avatar.CallOpts)
}

// ReferPrice is a free data retrieval call binding the contract method 0x6f7b0bab.
//
// Solidity: function referPrice() view returns(uint256)
func (_Avatar *AvatarCaller) ReferPrice(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Avatar.contract.Call(opts, &out, "referPrice")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ReferPrice is a free data retrieval call binding the contract method 0x6f7b0bab.
//
// Solidity: function referPrice() view returns(uint256)
func (_Avatar *AvatarSession) ReferPrice() (*big.Int, error) {
	return _Avatar.Contract.ReferPrice(&_Avatar.CallOpts)
}

// ReferPrice is a free data retrieval call binding the contract method 0x6f7b0bab.
//
// Solidity: function referPrice() view returns(uint256)
func (_Avatar *AvatarCallerSession) ReferPrice() (*big.Int, error) {
	return _Avatar.Contract.ReferPrice(&_Avatar.CallOpts)
}

// ReservePrice is a free data retrieval call binding the contract method 0xdb2e1eed.
//
// Solidity: function reservePrice() view returns(uint256)
func (_Avatar *AvatarCaller) ReservePrice(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Avatar.contract.Call(opts, &out, "reservePrice")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ReservePrice is a free data retrieval call binding the contract method 0xdb2e1eed.
//
// Solidity: function reservePrice() view returns(uint256)
func (_Avatar *AvatarSession) ReservePrice() (*big.Int, error) {
	return _Avatar.Contract.ReservePrice(&_Avatar.CallOpts)
}

// ReservePrice is a free data retrieval call binding the contract method 0xdb2e1eed.
//
// Solidity: function reservePrice() view returns(uint256)
func (_Avatar *AvatarCallerSession) ReservePrice() (*big.Int, error) {
	return _Avatar.Contract.ReservePrice(&_Avatar.CallOpts)
}

// Weth is a free data retrieval call binding the contract method 0x3fc8cef3.
//
// Solidity: function weth() view returns(address)
func (_Avatar *AvatarCaller) Weth(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Avatar.contract.Call(opts, &out, "weth")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Weth is a free data retrieval call binding the contract method 0x3fc8cef3.
//
// Solidity: function weth() view returns(address)
func (_Avatar *AvatarSession) Weth() (common.Address, error) {
	return _Avatar.Contract.Weth(&_Avatar.CallOpts)
}

// Weth is a free data retrieval call binding the contract method 0x3fc8cef3.
//
// Solidity: function weth() view returns(address)
func (_Avatar *AvatarCallerSession) Weth() (common.Address, error) {
	return _Avatar.Contract.Weth(&_Avatar.CallOpts)
}

// CreateBid is a paid mutator transaction binding the contract method 0x659dd2b4.
//
// Solidity: function createBid(uint256 _referId) payable returns()
func (_Avatar *AvatarTransactor) CreateBid(opts *bind.TransactOpts, _referId *big.Int) (*types.Transaction, error) {
	return _Avatar.contract.Transact(opts, "createBid", _referId)
}

// CreateBid is a paid mutator transaction binding the contract method 0x659dd2b4.
//
// Solidity: function createBid(uint256 _referId) payable returns()
func (_Avatar *AvatarSession) CreateBid(_referId *big.Int) (*types.Transaction, error) {
	return _Avatar.Contract.CreateBid(&_Avatar.TransactOpts, _referId)
}

// CreateBid is a paid mutator transaction binding the contract method 0x659dd2b4.
//
// Solidity: function createBid(uint256 _referId) payable returns()
func (_Avatar *AvatarTransactorSession) CreateBid(_referId *big.Int) (*types.Transaction, error) {
	return _Avatar.Contract.CreateBid(&_Avatar.TransactOpts, _referId)
}

// Initialize is a paid mutator transaction binding the contract method 0x57fb25cc.
//
// Solidity: function initialize(address _avatar, address _weth, uint256 _reservePrice, uint256 _referPrice, address _metaHolder) returns()
func (_Avatar *AvatarTransactor) Initialize(opts *bind.TransactOpts, _avatar common.Address, _weth common.Address, _reservePrice *big.Int, _referPrice *big.Int, _metaHolder common.Address) (*types.Transaction, error) {
	return _Avatar.contract.Transact(opts, "initialize", _avatar, _weth, _reservePrice, _referPrice, _metaHolder)
}

// Initialize is a paid mutator transaction binding the contract method 0x57fb25cc.
//
// Solidity: function initialize(address _avatar, address _weth, uint256 _reservePrice, uint256 _referPrice, address _metaHolder) returns()
func (_Avatar *AvatarSession) Initialize(_avatar common.Address, _weth common.Address, _reservePrice *big.Int, _referPrice *big.Int, _metaHolder common.Address) (*types.Transaction, error) {
	return _Avatar.Contract.Initialize(&_Avatar.TransactOpts, _avatar, _weth, _reservePrice, _referPrice, _metaHolder)
}

// Initialize is a paid mutator transaction binding the contract method 0x57fb25cc.
//
// Solidity: function initialize(address _avatar, address _weth, uint256 _reservePrice, uint256 _referPrice, address _metaHolder) returns()
func (_Avatar *AvatarTransactorSession) Initialize(_avatar common.Address, _weth common.Address, _reservePrice *big.Int, _referPrice *big.Int, _metaHolder common.Address) (*types.Transaction, error) {
	return _Avatar.Contract.Initialize(&_Avatar.TransactOpts, _avatar, _weth, _reservePrice, _referPrice, _metaHolder)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Avatar *AvatarTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Avatar.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Avatar *AvatarSession) Pause() (*types.Transaction, error) {
	return _Avatar.Contract.Pause(&_Avatar.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Avatar *AvatarTransactorSession) Pause() (*types.Transaction, error) {
	return _Avatar.Contract.Pause(&_Avatar.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Avatar *AvatarTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Avatar.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Avatar *AvatarSession) RenounceOwnership() (*types.Transaction, error) {
	return _Avatar.Contract.RenounceOwnership(&_Avatar.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Avatar *AvatarTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Avatar.Contract.RenounceOwnership(&_Avatar.TransactOpts)
}

// SetPrice is a paid mutator transaction binding the contract method 0xf7d97577.
//
// Solidity: function setPrice(uint256 _reservePrice, uint256 _referPrice) returns()
func (_Avatar *AvatarTransactor) SetPrice(opts *bind.TransactOpts, _reservePrice *big.Int, _referPrice *big.Int) (*types.Transaction, error) {
	return _Avatar.contract.Transact(opts, "setPrice", _reservePrice, _referPrice)
}

// SetPrice is a paid mutator transaction binding the contract method 0xf7d97577.
//
// Solidity: function setPrice(uint256 _reservePrice, uint256 _referPrice) returns()
func (_Avatar *AvatarSession) SetPrice(_reservePrice *big.Int, _referPrice *big.Int) (*types.Transaction, error) {
	return _Avatar.Contract.SetPrice(&_Avatar.TransactOpts, _reservePrice, _referPrice)
}

// SetPrice is a paid mutator transaction binding the contract method 0xf7d97577.
//
// Solidity: function setPrice(uint256 _reservePrice, uint256 _referPrice) returns()
func (_Avatar *AvatarTransactorSession) SetPrice(_reservePrice *big.Int, _referPrice *big.Int) (*types.Transaction, error) {
	return _Avatar.Contract.SetPrice(&_Avatar.TransactOpts, _reservePrice, _referPrice)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Avatar *AvatarTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Avatar.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Avatar *AvatarSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Avatar.Contract.TransferOwnership(&_Avatar.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Avatar *AvatarTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Avatar.Contract.TransferOwnership(&_Avatar.TransactOpts, newOwner)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Avatar *AvatarTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Avatar.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Avatar *AvatarSession) Unpause() (*types.Transaction, error) {
	return _Avatar.Contract.Unpause(&_Avatar.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Avatar *AvatarTransactorSession) Unpause() (*types.Transaction, error) {
	return _Avatar.Contract.Unpause(&_Avatar.TransactOpts)
}

// AvatarAuctionBidIterator is returned from FilterAuctionBid and is used to iterate over the raw logs and unpacked data for AuctionBid events raised by the Avatar contract.
type AvatarAuctionBidIterator struct {
	Event *AvatarAuctionBid // Event containing the contract specifics and raw log

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
func (it *AvatarAuctionBidIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AvatarAuctionBid)
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
		it.Event = new(AvatarAuctionBid)
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
func (it *AvatarAuctionBidIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AvatarAuctionBidIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AvatarAuctionBid represents a AuctionBid event raised by the Avatar contract.
type AvatarAuctionBid struct {
	NounId *big.Int
	Sender common.Address
	Value  *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterAuctionBid is a free log retrieval operation binding the contract event 0xb8d756f2d1da4663767eb4d559780ace84f2f65a421a60ddcb47e8b2e5d2fd26.
//
// Solidity: event AuctionBid(uint256 indexed nounId, address sender, uint256 value)
func (_Avatar *AvatarFilterer) FilterAuctionBid(opts *bind.FilterOpts, nounId []*big.Int) (*AvatarAuctionBidIterator, error) {

	var nounIdRule []interface{}
	for _, nounIdItem := range nounId {
		nounIdRule = append(nounIdRule, nounIdItem)
	}

	logs, sub, err := _Avatar.contract.FilterLogs(opts, "AuctionBid", nounIdRule)
	if err != nil {
		return nil, err
	}
	return &AvatarAuctionBidIterator{contract: _Avatar.contract, event: "AuctionBid", logs: logs, sub: sub}, nil
}

// WatchAuctionBid is a free log subscription operation binding the contract event 0xb8d756f2d1da4663767eb4d559780ace84f2f65a421a60ddcb47e8b2e5d2fd26.
//
// Solidity: event AuctionBid(uint256 indexed nounId, address sender, uint256 value)
func (_Avatar *AvatarFilterer) WatchAuctionBid(opts *bind.WatchOpts, sink chan<- *AvatarAuctionBid, nounId []*big.Int) (event.Subscription, error) {

	var nounIdRule []interface{}
	for _, nounIdItem := range nounId {
		nounIdRule = append(nounIdRule, nounIdItem)
	}

	logs, sub, err := _Avatar.contract.WatchLogs(opts, "AuctionBid", nounIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AvatarAuctionBid)
				if err := _Avatar.contract.UnpackLog(event, "AuctionBid", log); err != nil {
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
func (_Avatar *AvatarFilterer) ParseAuctionBid(log types.Log) (*AvatarAuctionBid, error) {
	event := new(AvatarAuctionBid)
	if err := _Avatar.contract.UnpackLog(event, "AuctionBid", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AvatarAuctionCreatedIterator is returned from FilterAuctionCreated and is used to iterate over the raw logs and unpacked data for AuctionCreated events raised by the Avatar contract.
type AvatarAuctionCreatedIterator struct {
	Event *AvatarAuctionCreated // Event containing the contract specifics and raw log

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
func (it *AvatarAuctionCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AvatarAuctionCreated)
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
		it.Event = new(AvatarAuctionCreated)
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
func (it *AvatarAuctionCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AvatarAuctionCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AvatarAuctionCreated represents a AuctionCreated event raised by the Avatar contract.
type AvatarAuctionCreated struct {
	NounId    *big.Int
	StartTime *big.Int
	EndTime   *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterAuctionCreated is a free log retrieval operation binding the contract event 0xd6eddd1118d71820909c1197aa966dbc15ed6f508554252169cc3d5ccac756ca.
//
// Solidity: event AuctionCreated(uint256 indexed nounId, uint256 startTime, uint256 endTime)
func (_Avatar *AvatarFilterer) FilterAuctionCreated(opts *bind.FilterOpts, nounId []*big.Int) (*AvatarAuctionCreatedIterator, error) {

	var nounIdRule []interface{}
	for _, nounIdItem := range nounId {
		nounIdRule = append(nounIdRule, nounIdItem)
	}

	logs, sub, err := _Avatar.contract.FilterLogs(opts, "AuctionCreated", nounIdRule)
	if err != nil {
		return nil, err
	}
	return &AvatarAuctionCreatedIterator{contract: _Avatar.contract, event: "AuctionCreated", logs: logs, sub: sub}, nil
}

// WatchAuctionCreated is a free log subscription operation binding the contract event 0xd6eddd1118d71820909c1197aa966dbc15ed6f508554252169cc3d5ccac756ca.
//
// Solidity: event AuctionCreated(uint256 indexed nounId, uint256 startTime, uint256 endTime)
func (_Avatar *AvatarFilterer) WatchAuctionCreated(opts *bind.WatchOpts, sink chan<- *AvatarAuctionCreated, nounId []*big.Int) (event.Subscription, error) {

	var nounIdRule []interface{}
	for _, nounIdItem := range nounId {
		nounIdRule = append(nounIdRule, nounIdItem)
	}

	logs, sub, err := _Avatar.contract.WatchLogs(opts, "AuctionCreated", nounIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AvatarAuctionCreated)
				if err := _Avatar.contract.UnpackLog(event, "AuctionCreated", log); err != nil {
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
func (_Avatar *AvatarFilterer) ParseAuctionCreated(log types.Log) (*AvatarAuctionCreated, error) {
	event := new(AvatarAuctionCreated)
	if err := _Avatar.contract.UnpackLog(event, "AuctionCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AvatarAuctionSettledIterator is returned from FilterAuctionSettled and is used to iterate over the raw logs and unpacked data for AuctionSettled events raised by the Avatar contract.
type AvatarAuctionSettledIterator struct {
	Event *AvatarAuctionSettled // Event containing the contract specifics and raw log

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
func (it *AvatarAuctionSettledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AvatarAuctionSettled)
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
		it.Event = new(AvatarAuctionSettled)
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
func (it *AvatarAuctionSettledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AvatarAuctionSettledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AvatarAuctionSettled represents a AuctionSettled event raised by the Avatar contract.
type AvatarAuctionSettled struct {
	NounId *big.Int
	Winner common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterAuctionSettled is a free log retrieval operation binding the contract event 0xc9f72b276a388619c6d185d146697036241880c36654b1a3ffdad07c24038d99.
//
// Solidity: event AuctionSettled(uint256 indexed nounId, address winner, uint256 amount)
func (_Avatar *AvatarFilterer) FilterAuctionSettled(opts *bind.FilterOpts, nounId []*big.Int) (*AvatarAuctionSettledIterator, error) {

	var nounIdRule []interface{}
	for _, nounIdItem := range nounId {
		nounIdRule = append(nounIdRule, nounIdItem)
	}

	logs, sub, err := _Avatar.contract.FilterLogs(opts, "AuctionSettled", nounIdRule)
	if err != nil {
		return nil, err
	}
	return &AvatarAuctionSettledIterator{contract: _Avatar.contract, event: "AuctionSettled", logs: logs, sub: sub}, nil
}

// WatchAuctionSettled is a free log subscription operation binding the contract event 0xc9f72b276a388619c6d185d146697036241880c36654b1a3ffdad07c24038d99.
//
// Solidity: event AuctionSettled(uint256 indexed nounId, address winner, uint256 amount)
func (_Avatar *AvatarFilterer) WatchAuctionSettled(opts *bind.WatchOpts, sink chan<- *AvatarAuctionSettled, nounId []*big.Int) (event.Subscription, error) {

	var nounIdRule []interface{}
	for _, nounIdItem := range nounId {
		nounIdRule = append(nounIdRule, nounIdItem)
	}

	logs, sub, err := _Avatar.contract.WatchLogs(opts, "AuctionSettled", nounIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AvatarAuctionSettled)
				if err := _Avatar.contract.UnpackLog(event, "AuctionSettled", log); err != nil {
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
func (_Avatar *AvatarFilterer) ParseAuctionSettled(log types.Log) (*AvatarAuctionSettled, error) {
	event := new(AvatarAuctionSettled)
	if err := _Avatar.contract.UnpackLog(event, "AuctionSettled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AvatarOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Avatar contract.
type AvatarOwnershipTransferredIterator struct {
	Event *AvatarOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *AvatarOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AvatarOwnershipTransferred)
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
		it.Event = new(AvatarOwnershipTransferred)
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
func (it *AvatarOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AvatarOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AvatarOwnershipTransferred represents a OwnershipTransferred event raised by the Avatar contract.
type AvatarOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Avatar *AvatarFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*AvatarOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Avatar.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &AvatarOwnershipTransferredIterator{contract: _Avatar.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Avatar *AvatarFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *AvatarOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Avatar.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AvatarOwnershipTransferred)
				if err := _Avatar.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Avatar *AvatarFilterer) ParseOwnershipTransferred(log types.Log) (*AvatarOwnershipTransferred, error) {
	event := new(AvatarOwnershipTransferred)
	if err := _Avatar.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AvatarPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the Avatar contract.
type AvatarPausedIterator struct {
	Event *AvatarPaused // Event containing the contract specifics and raw log

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
func (it *AvatarPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AvatarPaused)
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
		it.Event = new(AvatarPaused)
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
func (it *AvatarPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AvatarPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AvatarPaused represents a Paused event raised by the Avatar contract.
type AvatarPaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_Avatar *AvatarFilterer) FilterPaused(opts *bind.FilterOpts) (*AvatarPausedIterator, error) {

	logs, sub, err := _Avatar.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &AvatarPausedIterator{contract: _Avatar.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_Avatar *AvatarFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *AvatarPaused) (event.Subscription, error) {

	logs, sub, err := _Avatar.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AvatarPaused)
				if err := _Avatar.contract.UnpackLog(event, "Paused", log); err != nil {
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
func (_Avatar *AvatarFilterer) ParsePaused(log types.Log) (*AvatarPaused, error) {
	event := new(AvatarPaused)
	if err := _Avatar.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AvatarUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the Avatar contract.
type AvatarUnpausedIterator struct {
	Event *AvatarUnpaused // Event containing the contract specifics and raw log

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
func (it *AvatarUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AvatarUnpaused)
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
		it.Event = new(AvatarUnpaused)
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
func (it *AvatarUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AvatarUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AvatarUnpaused represents a Unpaused event raised by the Avatar contract.
type AvatarUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_Avatar *AvatarFilterer) FilterUnpaused(opts *bind.FilterOpts) (*AvatarUnpausedIterator, error) {

	logs, sub, err := _Avatar.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &AvatarUnpausedIterator{contract: _Avatar.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_Avatar *AvatarFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *AvatarUnpaused) (event.Subscription, error) {

	logs, sub, err := _Avatar.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AvatarUnpaused)
				if err := _Avatar.contract.UnpackLog(event, "Unpaused", log); err != nil {
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
func (_Avatar *AvatarFilterer) ParseUnpaused(log types.Log) (*AvatarUnpaused, error) {
	event := new(AvatarUnpaused)
	if err := _Avatar.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
