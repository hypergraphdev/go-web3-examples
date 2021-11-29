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

// IAvatarSeederSeed is an auto generated low-level Go binding around an user-defined struct.
type IAvatarSeederSeed struct {
	Background *big.Int
	Body       *big.Int
	Accessory  *big.Int
	Head       *big.Int
	Glasses    *big.Int
}

// AvatarMetaData contains all meta data concerning the Avatar contract.
var AvatarMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_minter\",\"type\":\"address\"},{\"internalType\":\"contractIAvatarDescriptor\",\"name\":\"_descriptor\",\"type\":\"address\"},{\"internalType\":\"contractIAvatarSeeder\",\"name\":\"_seeder\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"approved\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"delegator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"fromDelegate\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"toDelegate\",\"type\":\"address\"}],\"name\":\"DelegateChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"delegate\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"previousBalance\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newBalance\",\"type\":\"uint256\"}],\"name\":\"DelegateVotesChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"DescriptorLocked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"contractIAvatarDescriptor\",\"name\":\"descriptor\",\"type\":\"address\"}],\"name\":\"DescriptorUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"MinterLocked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"minter\",\"type\":\"address\"}],\"name\":\"MinterUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"NounBurned\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint48\",\"name\":\"background\",\"type\":\"uint48\"},{\"internalType\":\"uint48\",\"name\":\"body\",\"type\":\"uint48\"},{\"internalType\":\"uint48\",\"name\":\"accessory\",\"type\":\"uint48\"},{\"internalType\":\"uint48\",\"name\":\"head\",\"type\":\"uint48\"},{\"internalType\":\"uint48\",\"name\":\"glasses\",\"type\":\"uint48\"}],\"indexed\":false,\"internalType\":\"structIAvatarSeeder.Seed\",\"name\":\"seed\",\"type\":\"tuple\"}],\"name\":\"NounCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"preOwner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"NounExchanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"tokenURI\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"author\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"fhash\",\"type\":\"uint256\"}],\"name\":\"NounRecorded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"noundersDAO\",\"type\":\"address\"}],\"name\":\"NoundersDAOUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"SeederLocked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"contractIAvatarSeeder\",\"name\":\"seeder\",\"type\":\"address\"}],\"name\":\"SeederUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DELEGATION_TYPEHASH\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DOMAIN_TYPEHASH\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"_tokenURIMap\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"nounId\",\"type\":\"uint256\"}],\"name\":\"burn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"name\":\"checkpoints\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"fromBlock\",\"type\":\"uint32\"},{\"internalType\":\"uint96\",\"name\":\"votes\",\"type\":\"uint96\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"contractURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"dataURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"delegatee\",\"type\":\"address\"}],\"name\":\"delegate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"delegatee\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expiry\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"delegateBySig\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"delegator\",\"type\":\"address\"}],\"name\":\"delegates\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"descriptor\",\"outputs\":[{\"internalType\":\"contractIAvatarDescriptor\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getApproved\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"getCurrentVotes\",\"outputs\":[{\"internalType\":\"uint96\",\"name\":\"\",\"type\":\"uint96\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getNowSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"blockNumber\",\"type\":\"uint256\"}],\"name\":\"getPriorVotes\",\"outputs\":[{\"internalType\":\"uint96\",\"name\":\"\",\"type\":\"uint96\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isDescriptorLocked\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isMinterLocked\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isSeederLocked\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lockDescriptor\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lockMinter\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lockSeeder\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"mint\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"minter\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"nonces\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"numCheckpoints\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"seeder\",\"outputs\":[{\"internalType\":\"contractIAvatarSeeder\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"seeds\",\"outputs\":[{\"internalType\":\"uint48\",\"name\":\"background\",\"type\":\"uint48\"},{\"internalType\":\"uint48\",\"name\":\"body\",\"type\":\"uint48\"},{\"internalType\":\"uint48\",\"name\":\"accessory\",\"type\":\"uint48\"},{\"internalType\":\"uint48\",\"name\":\"head\",\"type\":\"uint48\"},{\"internalType\":\"uint48\",\"name\":\"glasses\",\"type\":\"uint48\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"newContractURIHash\",\"type\":\"string\"}],\"name\":\"setContractURIHash\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIAvatarDescriptor\",\"name\":\"_descriptor\",\"type\":\"address\"}],\"name\":\"setDescriptor\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_minter\",\"type\":\"address\"}],\"name\":\"setMinter\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIAvatarSeeder\",\"name\":\"_seeder\",\"type\":\"address\"}],\"name\":\"setSeeder\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"tokenByIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"tokenOfOwnerByIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"tokenURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"delegator\",\"type\":\"address\"}],\"name\":\"votesToDelegate\",\"outputs\":[{\"internalType\":\"uint96\",\"name\":\"\",\"type\":\"uint96\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
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

// DELEGATIONTYPEHASH is a free data retrieval call binding the contract method 0xe7a324dc.
//
// Solidity: function DELEGATION_TYPEHASH() view returns(bytes32)
func (_Avatar *AvatarCaller) DELEGATIONTYPEHASH(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Avatar.contract.Call(opts, &out, "DELEGATION_TYPEHASH")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DELEGATIONTYPEHASH is a free data retrieval call binding the contract method 0xe7a324dc.
//
// Solidity: function DELEGATION_TYPEHASH() view returns(bytes32)
func (_Avatar *AvatarSession) DELEGATIONTYPEHASH() ([32]byte, error) {
	return _Avatar.Contract.DELEGATIONTYPEHASH(&_Avatar.CallOpts)
}

// DELEGATIONTYPEHASH is a free data retrieval call binding the contract method 0xe7a324dc.
//
// Solidity: function DELEGATION_TYPEHASH() view returns(bytes32)
func (_Avatar *AvatarCallerSession) DELEGATIONTYPEHASH() ([32]byte, error) {
	return _Avatar.Contract.DELEGATIONTYPEHASH(&_Avatar.CallOpts)
}

// DOMAINTYPEHASH is a free data retrieval call binding the contract method 0x20606b70.
//
// Solidity: function DOMAIN_TYPEHASH() view returns(bytes32)
func (_Avatar *AvatarCaller) DOMAINTYPEHASH(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Avatar.contract.Call(opts, &out, "DOMAIN_TYPEHASH")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DOMAINTYPEHASH is a free data retrieval call binding the contract method 0x20606b70.
//
// Solidity: function DOMAIN_TYPEHASH() view returns(bytes32)
func (_Avatar *AvatarSession) DOMAINTYPEHASH() ([32]byte, error) {
	return _Avatar.Contract.DOMAINTYPEHASH(&_Avatar.CallOpts)
}

// DOMAINTYPEHASH is a free data retrieval call binding the contract method 0x20606b70.
//
// Solidity: function DOMAIN_TYPEHASH() view returns(bytes32)
func (_Avatar *AvatarCallerSession) DOMAINTYPEHASH() ([32]byte, error) {
	return _Avatar.Contract.DOMAINTYPEHASH(&_Avatar.CallOpts)
}

// TokenURIMap is a free data retrieval call binding the contract method 0x3589cac4.
//
// Solidity: function _tokenURIMap(uint256 ) view returns(string)
func (_Avatar *AvatarCaller) TokenURIMap(opts *bind.CallOpts, arg0 *big.Int) (string, error) {
	var out []interface{}
	err := _Avatar.contract.Call(opts, &out, "_tokenURIMap", arg0)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// TokenURIMap is a free data retrieval call binding the contract method 0x3589cac4.
//
// Solidity: function _tokenURIMap(uint256 ) view returns(string)
func (_Avatar *AvatarSession) TokenURIMap(arg0 *big.Int) (string, error) {
	return _Avatar.Contract.TokenURIMap(&_Avatar.CallOpts, arg0)
}

// TokenURIMap is a free data retrieval call binding the contract method 0x3589cac4.
//
// Solidity: function _tokenURIMap(uint256 ) view returns(string)
func (_Avatar *AvatarCallerSession) TokenURIMap(arg0 *big.Int) (string, error) {
	return _Avatar.Contract.TokenURIMap(&_Avatar.CallOpts, arg0)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Avatar *AvatarCaller) BalanceOf(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Avatar.contract.Call(opts, &out, "balanceOf", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Avatar *AvatarSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _Avatar.Contract.BalanceOf(&_Avatar.CallOpts, owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Avatar *AvatarCallerSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _Avatar.Contract.BalanceOf(&_Avatar.CallOpts, owner)
}

// Checkpoints is a free data retrieval call binding the contract method 0xf1127ed8.
//
// Solidity: function checkpoints(address , uint32 ) view returns(uint32 fromBlock, uint96 votes)
func (_Avatar *AvatarCaller) Checkpoints(opts *bind.CallOpts, arg0 common.Address, arg1 uint32) (struct {
	FromBlock uint32
	Votes     *big.Int
}, error) {
	var out []interface{}
	err := _Avatar.contract.Call(opts, &out, "checkpoints", arg0, arg1)

	outstruct := new(struct {
		FromBlock uint32
		Votes     *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.FromBlock = *abi.ConvertType(out[0], new(uint32)).(*uint32)
	outstruct.Votes = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Checkpoints is a free data retrieval call binding the contract method 0xf1127ed8.
//
// Solidity: function checkpoints(address , uint32 ) view returns(uint32 fromBlock, uint96 votes)
func (_Avatar *AvatarSession) Checkpoints(arg0 common.Address, arg1 uint32) (struct {
	FromBlock uint32
	Votes     *big.Int
}, error) {
	return _Avatar.Contract.Checkpoints(&_Avatar.CallOpts, arg0, arg1)
}

// Checkpoints is a free data retrieval call binding the contract method 0xf1127ed8.
//
// Solidity: function checkpoints(address , uint32 ) view returns(uint32 fromBlock, uint96 votes)
func (_Avatar *AvatarCallerSession) Checkpoints(arg0 common.Address, arg1 uint32) (struct {
	FromBlock uint32
	Votes     *big.Int
}, error) {
	return _Avatar.Contract.Checkpoints(&_Avatar.CallOpts, arg0, arg1)
}

// ContractURI is a free data retrieval call binding the contract method 0xe8a3d485.
//
// Solidity: function contractURI() view returns(string)
func (_Avatar *AvatarCaller) ContractURI(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Avatar.contract.Call(opts, &out, "contractURI")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// ContractURI is a free data retrieval call binding the contract method 0xe8a3d485.
//
// Solidity: function contractURI() view returns(string)
func (_Avatar *AvatarSession) ContractURI() (string, error) {
	return _Avatar.Contract.ContractURI(&_Avatar.CallOpts)
}

// ContractURI is a free data retrieval call binding the contract method 0xe8a3d485.
//
// Solidity: function contractURI() view returns(string)
func (_Avatar *AvatarCallerSession) ContractURI() (string, error) {
	return _Avatar.Contract.ContractURI(&_Avatar.CallOpts)
}

// DataURI is a free data retrieval call binding the contract method 0x5ac1e3bb.
//
// Solidity: function dataURI(uint256 tokenId) view returns(string)
func (_Avatar *AvatarCaller) DataURI(opts *bind.CallOpts, tokenId *big.Int) (string, error) {
	var out []interface{}
	err := _Avatar.contract.Call(opts, &out, "dataURI", tokenId)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// DataURI is a free data retrieval call binding the contract method 0x5ac1e3bb.
//
// Solidity: function dataURI(uint256 tokenId) view returns(string)
func (_Avatar *AvatarSession) DataURI(tokenId *big.Int) (string, error) {
	return _Avatar.Contract.DataURI(&_Avatar.CallOpts, tokenId)
}

// DataURI is a free data retrieval call binding the contract method 0x5ac1e3bb.
//
// Solidity: function dataURI(uint256 tokenId) view returns(string)
func (_Avatar *AvatarCallerSession) DataURI(tokenId *big.Int) (string, error) {
	return _Avatar.Contract.DataURI(&_Avatar.CallOpts, tokenId)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Avatar *AvatarCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Avatar.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Avatar *AvatarSession) Decimals() (uint8, error) {
	return _Avatar.Contract.Decimals(&_Avatar.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Avatar *AvatarCallerSession) Decimals() (uint8, error) {
	return _Avatar.Contract.Decimals(&_Avatar.CallOpts)
}

// Delegates is a free data retrieval call binding the contract method 0x587cde1e.
//
// Solidity: function delegates(address delegator) view returns(address)
func (_Avatar *AvatarCaller) Delegates(opts *bind.CallOpts, delegator common.Address) (common.Address, error) {
	var out []interface{}
	err := _Avatar.contract.Call(opts, &out, "delegates", delegator)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Delegates is a free data retrieval call binding the contract method 0x587cde1e.
//
// Solidity: function delegates(address delegator) view returns(address)
func (_Avatar *AvatarSession) Delegates(delegator common.Address) (common.Address, error) {
	return _Avatar.Contract.Delegates(&_Avatar.CallOpts, delegator)
}

// Delegates is a free data retrieval call binding the contract method 0x587cde1e.
//
// Solidity: function delegates(address delegator) view returns(address)
func (_Avatar *AvatarCallerSession) Delegates(delegator common.Address) (common.Address, error) {
	return _Avatar.Contract.Delegates(&_Avatar.CallOpts, delegator)
}

// Descriptor is a free data retrieval call binding the contract method 0x303e74df.
//
// Solidity: function descriptor() view returns(address)
func (_Avatar *AvatarCaller) Descriptor(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Avatar.contract.Call(opts, &out, "descriptor")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Descriptor is a free data retrieval call binding the contract method 0x303e74df.
//
// Solidity: function descriptor() view returns(address)
func (_Avatar *AvatarSession) Descriptor() (common.Address, error) {
	return _Avatar.Contract.Descriptor(&_Avatar.CallOpts)
}

// Descriptor is a free data retrieval call binding the contract method 0x303e74df.
//
// Solidity: function descriptor() view returns(address)
func (_Avatar *AvatarCallerSession) Descriptor() (common.Address, error) {
	return _Avatar.Contract.Descriptor(&_Avatar.CallOpts)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_Avatar *AvatarCaller) GetApproved(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Avatar.contract.Call(opts, &out, "getApproved", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_Avatar *AvatarSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _Avatar.Contract.GetApproved(&_Avatar.CallOpts, tokenId)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_Avatar *AvatarCallerSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _Avatar.Contract.GetApproved(&_Avatar.CallOpts, tokenId)
}

// GetCurrentVotes is a free data retrieval call binding the contract method 0xb4b5ea57.
//
// Solidity: function getCurrentVotes(address account) view returns(uint96)
func (_Avatar *AvatarCaller) GetCurrentVotes(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Avatar.contract.Call(opts, &out, "getCurrentVotes", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetCurrentVotes is a free data retrieval call binding the contract method 0xb4b5ea57.
//
// Solidity: function getCurrentVotes(address account) view returns(uint96)
func (_Avatar *AvatarSession) GetCurrentVotes(account common.Address) (*big.Int, error) {
	return _Avatar.Contract.GetCurrentVotes(&_Avatar.CallOpts, account)
}

// GetCurrentVotes is a free data retrieval call binding the contract method 0xb4b5ea57.
//
// Solidity: function getCurrentVotes(address account) view returns(uint96)
func (_Avatar *AvatarCallerSession) GetCurrentVotes(account common.Address) (*big.Int, error) {
	return _Avatar.Contract.GetCurrentVotes(&_Avatar.CallOpts, account)
}

// GetNowSupply is a free data retrieval call binding the contract method 0x7bd7ecef.
//
// Solidity: function getNowSupply() view returns(uint256)
func (_Avatar *AvatarCaller) GetNowSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Avatar.contract.Call(opts, &out, "getNowSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetNowSupply is a free data retrieval call binding the contract method 0x7bd7ecef.
//
// Solidity: function getNowSupply() view returns(uint256)
func (_Avatar *AvatarSession) GetNowSupply() (*big.Int, error) {
	return _Avatar.Contract.GetNowSupply(&_Avatar.CallOpts)
}

// GetNowSupply is a free data retrieval call binding the contract method 0x7bd7ecef.
//
// Solidity: function getNowSupply() view returns(uint256)
func (_Avatar *AvatarCallerSession) GetNowSupply() (*big.Int, error) {
	return _Avatar.Contract.GetNowSupply(&_Avatar.CallOpts)
}

// GetPriorVotes is a free data retrieval call binding the contract method 0x782d6fe1.
//
// Solidity: function getPriorVotes(address account, uint256 blockNumber) view returns(uint96)
func (_Avatar *AvatarCaller) GetPriorVotes(opts *bind.CallOpts, account common.Address, blockNumber *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Avatar.contract.Call(opts, &out, "getPriorVotes", account, blockNumber)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetPriorVotes is a free data retrieval call binding the contract method 0x782d6fe1.
//
// Solidity: function getPriorVotes(address account, uint256 blockNumber) view returns(uint96)
func (_Avatar *AvatarSession) GetPriorVotes(account common.Address, blockNumber *big.Int) (*big.Int, error) {
	return _Avatar.Contract.GetPriorVotes(&_Avatar.CallOpts, account, blockNumber)
}

// GetPriorVotes is a free data retrieval call binding the contract method 0x782d6fe1.
//
// Solidity: function getPriorVotes(address account, uint256 blockNumber) view returns(uint96)
func (_Avatar *AvatarCallerSession) GetPriorVotes(account common.Address, blockNumber *big.Int) (*big.Int, error) {
	return _Avatar.Contract.GetPriorVotes(&_Avatar.CallOpts, account, blockNumber)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_Avatar *AvatarCaller) IsApprovedForAll(opts *bind.CallOpts, owner common.Address, operator common.Address) (bool, error) {
	var out []interface{}
	err := _Avatar.contract.Call(opts, &out, "isApprovedForAll", owner, operator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_Avatar *AvatarSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _Avatar.Contract.IsApprovedForAll(&_Avatar.CallOpts, owner, operator)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_Avatar *AvatarCallerSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _Avatar.Contract.IsApprovedForAll(&_Avatar.CallOpts, owner, operator)
}

// IsDescriptorLocked is a free data retrieval call binding the contract method 0xc1b8e4e1.
//
// Solidity: function isDescriptorLocked() view returns(bool)
func (_Avatar *AvatarCaller) IsDescriptorLocked(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Avatar.contract.Call(opts, &out, "isDescriptorLocked")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsDescriptorLocked is a free data retrieval call binding the contract method 0xc1b8e4e1.
//
// Solidity: function isDescriptorLocked() view returns(bool)
func (_Avatar *AvatarSession) IsDescriptorLocked() (bool, error) {
	return _Avatar.Contract.IsDescriptorLocked(&_Avatar.CallOpts)
}

// IsDescriptorLocked is a free data retrieval call binding the contract method 0xc1b8e4e1.
//
// Solidity: function isDescriptorLocked() view returns(bool)
func (_Avatar *AvatarCallerSession) IsDescriptorLocked() (bool, error) {
	return _Avatar.Contract.IsDescriptorLocked(&_Avatar.CallOpts)
}

// IsMinterLocked is a free data retrieval call binding the contract method 0x1e688e10.
//
// Solidity: function isMinterLocked() view returns(bool)
func (_Avatar *AvatarCaller) IsMinterLocked(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Avatar.contract.Call(opts, &out, "isMinterLocked")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsMinterLocked is a free data retrieval call binding the contract method 0x1e688e10.
//
// Solidity: function isMinterLocked() view returns(bool)
func (_Avatar *AvatarSession) IsMinterLocked() (bool, error) {
	return _Avatar.Contract.IsMinterLocked(&_Avatar.CallOpts)
}

// IsMinterLocked is a free data retrieval call binding the contract method 0x1e688e10.
//
// Solidity: function isMinterLocked() view returns(bool)
func (_Avatar *AvatarCallerSession) IsMinterLocked() (bool, error) {
	return _Avatar.Contract.IsMinterLocked(&_Avatar.CallOpts)
}

// IsSeederLocked is a free data retrieval call binding the contract method 0xc8fc0c23.
//
// Solidity: function isSeederLocked() view returns(bool)
func (_Avatar *AvatarCaller) IsSeederLocked(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Avatar.contract.Call(opts, &out, "isSeederLocked")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsSeederLocked is a free data retrieval call binding the contract method 0xc8fc0c23.
//
// Solidity: function isSeederLocked() view returns(bool)
func (_Avatar *AvatarSession) IsSeederLocked() (bool, error) {
	return _Avatar.Contract.IsSeederLocked(&_Avatar.CallOpts)
}

// IsSeederLocked is a free data retrieval call binding the contract method 0xc8fc0c23.
//
// Solidity: function isSeederLocked() view returns(bool)
func (_Avatar *AvatarCallerSession) IsSeederLocked() (bool, error) {
	return _Avatar.Contract.IsSeederLocked(&_Avatar.CallOpts)
}

// Minter is a free data retrieval call binding the contract method 0x07546172.
//
// Solidity: function minter() view returns(address)
func (_Avatar *AvatarCaller) Minter(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Avatar.contract.Call(opts, &out, "minter")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Minter is a free data retrieval call binding the contract method 0x07546172.
//
// Solidity: function minter() view returns(address)
func (_Avatar *AvatarSession) Minter() (common.Address, error) {
	return _Avatar.Contract.Minter(&_Avatar.CallOpts)
}

// Minter is a free data retrieval call binding the contract method 0x07546172.
//
// Solidity: function minter() view returns(address)
func (_Avatar *AvatarCallerSession) Minter() (common.Address, error) {
	return _Avatar.Contract.Minter(&_Avatar.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Avatar *AvatarCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Avatar.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Avatar *AvatarSession) Name() (string, error) {
	return _Avatar.Contract.Name(&_Avatar.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Avatar *AvatarCallerSession) Name() (string, error) {
	return _Avatar.Contract.Name(&_Avatar.CallOpts)
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address ) view returns(uint256)
func (_Avatar *AvatarCaller) Nonces(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Avatar.contract.Call(opts, &out, "nonces", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address ) view returns(uint256)
func (_Avatar *AvatarSession) Nonces(arg0 common.Address) (*big.Int, error) {
	return _Avatar.Contract.Nonces(&_Avatar.CallOpts, arg0)
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address ) view returns(uint256)
func (_Avatar *AvatarCallerSession) Nonces(arg0 common.Address) (*big.Int, error) {
	return _Avatar.Contract.Nonces(&_Avatar.CallOpts, arg0)
}

// NumCheckpoints is a free data retrieval call binding the contract method 0x6fcfff45.
//
// Solidity: function numCheckpoints(address ) view returns(uint32)
func (_Avatar *AvatarCaller) NumCheckpoints(opts *bind.CallOpts, arg0 common.Address) (uint32, error) {
	var out []interface{}
	err := _Avatar.contract.Call(opts, &out, "numCheckpoints", arg0)

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// NumCheckpoints is a free data retrieval call binding the contract method 0x6fcfff45.
//
// Solidity: function numCheckpoints(address ) view returns(uint32)
func (_Avatar *AvatarSession) NumCheckpoints(arg0 common.Address) (uint32, error) {
	return _Avatar.Contract.NumCheckpoints(&_Avatar.CallOpts, arg0)
}

// NumCheckpoints is a free data retrieval call binding the contract method 0x6fcfff45.
//
// Solidity: function numCheckpoints(address ) view returns(uint32)
func (_Avatar *AvatarCallerSession) NumCheckpoints(arg0 common.Address) (uint32, error) {
	return _Avatar.Contract.NumCheckpoints(&_Avatar.CallOpts, arg0)
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

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_Avatar *AvatarCaller) OwnerOf(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Avatar.contract.Call(opts, &out, "ownerOf", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_Avatar *AvatarSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _Avatar.Contract.OwnerOf(&_Avatar.CallOpts, tokenId)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_Avatar *AvatarCallerSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _Avatar.Contract.OwnerOf(&_Avatar.CallOpts, tokenId)
}

// Seeder is a free data retrieval call binding the contract method 0x684931ed.
//
// Solidity: function seeder() view returns(address)
func (_Avatar *AvatarCaller) Seeder(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Avatar.contract.Call(opts, &out, "seeder")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Seeder is a free data retrieval call binding the contract method 0x684931ed.
//
// Solidity: function seeder() view returns(address)
func (_Avatar *AvatarSession) Seeder() (common.Address, error) {
	return _Avatar.Contract.Seeder(&_Avatar.CallOpts)
}

// Seeder is a free data retrieval call binding the contract method 0x684931ed.
//
// Solidity: function seeder() view returns(address)
func (_Avatar *AvatarCallerSession) Seeder() (common.Address, error) {
	return _Avatar.Contract.Seeder(&_Avatar.CallOpts)
}

// Seeds is a free data retrieval call binding the contract method 0xf0503e80.
//
// Solidity: function seeds(uint256 ) view returns(uint48 background, uint48 body, uint48 accessory, uint48 head, uint48 glasses)
func (_Avatar *AvatarCaller) Seeds(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Background *big.Int
	Body       *big.Int
	Accessory  *big.Int
	Head       *big.Int
	Glasses    *big.Int
}, error) {
	var out []interface{}
	err := _Avatar.contract.Call(opts, &out, "seeds", arg0)

	outstruct := new(struct {
		Background *big.Int
		Body       *big.Int
		Accessory  *big.Int
		Head       *big.Int
		Glasses    *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Background = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Body = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.Accessory = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.Head = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.Glasses = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Seeds is a free data retrieval call binding the contract method 0xf0503e80.
//
// Solidity: function seeds(uint256 ) view returns(uint48 background, uint48 body, uint48 accessory, uint48 head, uint48 glasses)
func (_Avatar *AvatarSession) Seeds(arg0 *big.Int) (struct {
	Background *big.Int
	Body       *big.Int
	Accessory  *big.Int
	Head       *big.Int
	Glasses    *big.Int
}, error) {
	return _Avatar.Contract.Seeds(&_Avatar.CallOpts, arg0)
}

// Seeds is a free data retrieval call binding the contract method 0xf0503e80.
//
// Solidity: function seeds(uint256 ) view returns(uint48 background, uint48 body, uint48 accessory, uint48 head, uint48 glasses)
func (_Avatar *AvatarCallerSession) Seeds(arg0 *big.Int) (struct {
	Background *big.Int
	Body       *big.Int
	Accessory  *big.Int
	Head       *big.Int
	Glasses    *big.Int
}, error) {
	return _Avatar.Contract.Seeds(&_Avatar.CallOpts, arg0)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Avatar *AvatarCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _Avatar.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Avatar *AvatarSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Avatar.Contract.SupportsInterface(&_Avatar.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Avatar *AvatarCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Avatar.Contract.SupportsInterface(&_Avatar.CallOpts, interfaceId)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Avatar *AvatarCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Avatar.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Avatar *AvatarSession) Symbol() (string, error) {
	return _Avatar.Contract.Symbol(&_Avatar.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Avatar *AvatarCallerSession) Symbol() (string, error) {
	return _Avatar.Contract.Symbol(&_Avatar.CallOpts)
}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(uint256 index) view returns(uint256)
func (_Avatar *AvatarCaller) TokenByIndex(opts *bind.CallOpts, index *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Avatar.contract.Call(opts, &out, "tokenByIndex", index)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(uint256 index) view returns(uint256)
func (_Avatar *AvatarSession) TokenByIndex(index *big.Int) (*big.Int, error) {
	return _Avatar.Contract.TokenByIndex(&_Avatar.CallOpts, index)
}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(uint256 index) view returns(uint256)
func (_Avatar *AvatarCallerSession) TokenByIndex(index *big.Int) (*big.Int, error) {
	return _Avatar.Contract.TokenByIndex(&_Avatar.CallOpts, index)
}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address owner, uint256 index) view returns(uint256)
func (_Avatar *AvatarCaller) TokenOfOwnerByIndex(opts *bind.CallOpts, owner common.Address, index *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Avatar.contract.Call(opts, &out, "tokenOfOwnerByIndex", owner, index)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address owner, uint256 index) view returns(uint256)
func (_Avatar *AvatarSession) TokenOfOwnerByIndex(owner common.Address, index *big.Int) (*big.Int, error) {
	return _Avatar.Contract.TokenOfOwnerByIndex(&_Avatar.CallOpts, owner, index)
}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address owner, uint256 index) view returns(uint256)
func (_Avatar *AvatarCallerSession) TokenOfOwnerByIndex(owner common.Address, index *big.Int) (*big.Int, error) {
	return _Avatar.Contract.TokenOfOwnerByIndex(&_Avatar.CallOpts, owner, index)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_Avatar *AvatarCaller) TokenURI(opts *bind.CallOpts, tokenId *big.Int) (string, error) {
	var out []interface{}
	err := _Avatar.contract.Call(opts, &out, "tokenURI", tokenId)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_Avatar *AvatarSession) TokenURI(tokenId *big.Int) (string, error) {
	return _Avatar.Contract.TokenURI(&_Avatar.CallOpts, tokenId)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_Avatar *AvatarCallerSession) TokenURI(tokenId *big.Int) (string, error) {
	return _Avatar.Contract.TokenURI(&_Avatar.CallOpts, tokenId)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Avatar *AvatarCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Avatar.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Avatar *AvatarSession) TotalSupply() (*big.Int, error) {
	return _Avatar.Contract.TotalSupply(&_Avatar.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Avatar *AvatarCallerSession) TotalSupply() (*big.Int, error) {
	return _Avatar.Contract.TotalSupply(&_Avatar.CallOpts)
}

// VotesToDelegate is a free data retrieval call binding the contract method 0xe9580e91.
//
// Solidity: function votesToDelegate(address delegator) view returns(uint96)
func (_Avatar *AvatarCaller) VotesToDelegate(opts *bind.CallOpts, delegator common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Avatar.contract.Call(opts, &out, "votesToDelegate", delegator)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// VotesToDelegate is a free data retrieval call binding the contract method 0xe9580e91.
//
// Solidity: function votesToDelegate(address delegator) view returns(uint96)
func (_Avatar *AvatarSession) VotesToDelegate(delegator common.Address) (*big.Int, error) {
	return _Avatar.Contract.VotesToDelegate(&_Avatar.CallOpts, delegator)
}

// VotesToDelegate is a free data retrieval call binding the contract method 0xe9580e91.
//
// Solidity: function votesToDelegate(address delegator) view returns(uint96)
func (_Avatar *AvatarCallerSession) VotesToDelegate(delegator common.Address) (*big.Int, error) {
	return _Avatar.Contract.VotesToDelegate(&_Avatar.CallOpts, delegator)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_Avatar *AvatarTransactor) Approve(opts *bind.TransactOpts, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Avatar.contract.Transact(opts, "approve", to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_Avatar *AvatarSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Avatar.Contract.Approve(&_Avatar.TransactOpts, to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_Avatar *AvatarTransactorSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Avatar.Contract.Approve(&_Avatar.TransactOpts, to, tokenId)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 nounId) returns()
func (_Avatar *AvatarTransactor) Burn(opts *bind.TransactOpts, nounId *big.Int) (*types.Transaction, error) {
	return _Avatar.contract.Transact(opts, "burn", nounId)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 nounId) returns()
func (_Avatar *AvatarSession) Burn(nounId *big.Int) (*types.Transaction, error) {
	return _Avatar.Contract.Burn(&_Avatar.TransactOpts, nounId)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 nounId) returns()
func (_Avatar *AvatarTransactorSession) Burn(nounId *big.Int) (*types.Transaction, error) {
	return _Avatar.Contract.Burn(&_Avatar.TransactOpts, nounId)
}

// Delegate is a paid mutator transaction binding the contract method 0x5c19a95c.
//
// Solidity: function delegate(address delegatee) returns()
func (_Avatar *AvatarTransactor) Delegate(opts *bind.TransactOpts, delegatee common.Address) (*types.Transaction, error) {
	return _Avatar.contract.Transact(opts, "delegate", delegatee)
}

// Delegate is a paid mutator transaction binding the contract method 0x5c19a95c.
//
// Solidity: function delegate(address delegatee) returns()
func (_Avatar *AvatarSession) Delegate(delegatee common.Address) (*types.Transaction, error) {
	return _Avatar.Contract.Delegate(&_Avatar.TransactOpts, delegatee)
}

// Delegate is a paid mutator transaction binding the contract method 0x5c19a95c.
//
// Solidity: function delegate(address delegatee) returns()
func (_Avatar *AvatarTransactorSession) Delegate(delegatee common.Address) (*types.Transaction, error) {
	return _Avatar.Contract.Delegate(&_Avatar.TransactOpts, delegatee)
}

// DelegateBySig is a paid mutator transaction binding the contract method 0xc3cda520.
//
// Solidity: function delegateBySig(address delegatee, uint256 nonce, uint256 expiry, uint8 v, bytes32 r, bytes32 s) returns()
func (_Avatar *AvatarTransactor) DelegateBySig(opts *bind.TransactOpts, delegatee common.Address, nonce *big.Int, expiry *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _Avatar.contract.Transact(opts, "delegateBySig", delegatee, nonce, expiry, v, r, s)
}

// DelegateBySig is a paid mutator transaction binding the contract method 0xc3cda520.
//
// Solidity: function delegateBySig(address delegatee, uint256 nonce, uint256 expiry, uint8 v, bytes32 r, bytes32 s) returns()
func (_Avatar *AvatarSession) DelegateBySig(delegatee common.Address, nonce *big.Int, expiry *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _Avatar.Contract.DelegateBySig(&_Avatar.TransactOpts, delegatee, nonce, expiry, v, r, s)
}

// DelegateBySig is a paid mutator transaction binding the contract method 0xc3cda520.
//
// Solidity: function delegateBySig(address delegatee, uint256 nonce, uint256 expiry, uint8 v, bytes32 r, bytes32 s) returns()
func (_Avatar *AvatarTransactorSession) DelegateBySig(delegatee common.Address, nonce *big.Int, expiry *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _Avatar.Contract.DelegateBySig(&_Avatar.TransactOpts, delegatee, nonce, expiry, v, r, s)
}

// LockDescriptor is a paid mutator transaction binding the contract method 0x41b5d0de.
//
// Solidity: function lockDescriptor() returns()
func (_Avatar *AvatarTransactor) LockDescriptor(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Avatar.contract.Transact(opts, "lockDescriptor")
}

// LockDescriptor is a paid mutator transaction binding the contract method 0x41b5d0de.
//
// Solidity: function lockDescriptor() returns()
func (_Avatar *AvatarSession) LockDescriptor() (*types.Transaction, error) {
	return _Avatar.Contract.LockDescriptor(&_Avatar.TransactOpts)
}

// LockDescriptor is a paid mutator transaction binding the contract method 0x41b5d0de.
//
// Solidity: function lockDescriptor() returns()
func (_Avatar *AvatarTransactorSession) LockDescriptor() (*types.Transaction, error) {
	return _Avatar.Contract.LockDescriptor(&_Avatar.TransactOpts)
}

// LockMinter is a paid mutator transaction binding the contract method 0x76daebe1.
//
// Solidity: function lockMinter() returns()
func (_Avatar *AvatarTransactor) LockMinter(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Avatar.contract.Transact(opts, "lockMinter")
}

// LockMinter is a paid mutator transaction binding the contract method 0x76daebe1.
//
// Solidity: function lockMinter() returns()
func (_Avatar *AvatarSession) LockMinter() (*types.Transaction, error) {
	return _Avatar.Contract.LockMinter(&_Avatar.TransactOpts)
}

// LockMinter is a paid mutator transaction binding the contract method 0x76daebe1.
//
// Solidity: function lockMinter() returns()
func (_Avatar *AvatarTransactorSession) LockMinter() (*types.Transaction, error) {
	return _Avatar.Contract.LockMinter(&_Avatar.TransactOpts)
}

// LockSeeder is a paid mutator transaction binding the contract method 0x5f295a67.
//
// Solidity: function lockSeeder() returns()
func (_Avatar *AvatarTransactor) LockSeeder(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Avatar.contract.Transact(opts, "lockSeeder")
}

// LockSeeder is a paid mutator transaction binding the contract method 0x5f295a67.
//
// Solidity: function lockSeeder() returns()
func (_Avatar *AvatarSession) LockSeeder() (*types.Transaction, error) {
	return _Avatar.Contract.LockSeeder(&_Avatar.TransactOpts)
}

// LockSeeder is a paid mutator transaction binding the contract method 0x5f295a67.
//
// Solidity: function lockSeeder() returns()
func (_Avatar *AvatarTransactorSession) LockSeeder() (*types.Transaction, error) {
	return _Avatar.Contract.LockSeeder(&_Avatar.TransactOpts)
}

// Mint is a paid mutator transaction binding the contract method 0x1249c58b.
//
// Solidity: function mint() returns(uint256)
func (_Avatar *AvatarTransactor) Mint(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Avatar.contract.Transact(opts, "mint")
}

// Mint is a paid mutator transaction binding the contract method 0x1249c58b.
//
// Solidity: function mint() returns(uint256)
func (_Avatar *AvatarSession) Mint() (*types.Transaction, error) {
	return _Avatar.Contract.Mint(&_Avatar.TransactOpts)
}

// Mint is a paid mutator transaction binding the contract method 0x1249c58b.
//
// Solidity: function mint() returns(uint256)
func (_Avatar *AvatarTransactorSession) Mint() (*types.Transaction, error) {
	return _Avatar.Contract.Mint(&_Avatar.TransactOpts)
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

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_Avatar *AvatarTransactor) SafeTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Avatar.contract.Transact(opts, "safeTransferFrom", from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_Avatar *AvatarSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Avatar.Contract.SafeTransferFrom(&_Avatar.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_Avatar *AvatarTransactorSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Avatar.Contract.SafeTransferFrom(&_Avatar.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes _data) returns()
func (_Avatar *AvatarTransactor) SafeTransferFrom0(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _Avatar.contract.Transact(opts, "safeTransferFrom0", from, to, tokenId, _data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes _data) returns()
func (_Avatar *AvatarSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _Avatar.Contract.SafeTransferFrom0(&_Avatar.TransactOpts, from, to, tokenId, _data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes _data) returns()
func (_Avatar *AvatarTransactorSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _Avatar.Contract.SafeTransferFrom0(&_Avatar.TransactOpts, from, to, tokenId, _data)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Avatar *AvatarTransactor) SetApprovalForAll(opts *bind.TransactOpts, operator common.Address, approved bool) (*types.Transaction, error) {
	return _Avatar.contract.Transact(opts, "setApprovalForAll", operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Avatar *AvatarSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _Avatar.Contract.SetApprovalForAll(&_Avatar.TransactOpts, operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Avatar *AvatarTransactorSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _Avatar.Contract.SetApprovalForAll(&_Avatar.TransactOpts, operator, approved)
}

// SetContractURIHash is a paid mutator transaction binding the contract method 0xbaedc1c4.
//
// Solidity: function setContractURIHash(string newContractURIHash) returns()
func (_Avatar *AvatarTransactor) SetContractURIHash(opts *bind.TransactOpts, newContractURIHash string) (*types.Transaction, error) {
	return _Avatar.contract.Transact(opts, "setContractURIHash", newContractURIHash)
}

// SetContractURIHash is a paid mutator transaction binding the contract method 0xbaedc1c4.
//
// Solidity: function setContractURIHash(string newContractURIHash) returns()
func (_Avatar *AvatarSession) SetContractURIHash(newContractURIHash string) (*types.Transaction, error) {
	return _Avatar.Contract.SetContractURIHash(&_Avatar.TransactOpts, newContractURIHash)
}

// SetContractURIHash is a paid mutator transaction binding the contract method 0xbaedc1c4.
//
// Solidity: function setContractURIHash(string newContractURIHash) returns()
func (_Avatar *AvatarTransactorSession) SetContractURIHash(newContractURIHash string) (*types.Transaction, error) {
	return _Avatar.Contract.SetContractURIHash(&_Avatar.TransactOpts, newContractURIHash)
}

// SetDescriptor is a paid mutator transaction binding the contract method 0x01b9a397.
//
// Solidity: function setDescriptor(address _descriptor) returns()
func (_Avatar *AvatarTransactor) SetDescriptor(opts *bind.TransactOpts, _descriptor common.Address) (*types.Transaction, error) {
	return _Avatar.contract.Transact(opts, "setDescriptor", _descriptor)
}

// SetDescriptor is a paid mutator transaction binding the contract method 0x01b9a397.
//
// Solidity: function setDescriptor(address _descriptor) returns()
func (_Avatar *AvatarSession) SetDescriptor(_descriptor common.Address) (*types.Transaction, error) {
	return _Avatar.Contract.SetDescriptor(&_Avatar.TransactOpts, _descriptor)
}

// SetDescriptor is a paid mutator transaction binding the contract method 0x01b9a397.
//
// Solidity: function setDescriptor(address _descriptor) returns()
func (_Avatar *AvatarTransactorSession) SetDescriptor(_descriptor common.Address) (*types.Transaction, error) {
	return _Avatar.Contract.SetDescriptor(&_Avatar.TransactOpts, _descriptor)
}

// SetMinter is a paid mutator transaction binding the contract method 0xfca3b5aa.
//
// Solidity: function setMinter(address _minter) returns()
func (_Avatar *AvatarTransactor) SetMinter(opts *bind.TransactOpts, _minter common.Address) (*types.Transaction, error) {
	return _Avatar.contract.Transact(opts, "setMinter", _minter)
}

// SetMinter is a paid mutator transaction binding the contract method 0xfca3b5aa.
//
// Solidity: function setMinter(address _minter) returns()
func (_Avatar *AvatarSession) SetMinter(_minter common.Address) (*types.Transaction, error) {
	return _Avatar.Contract.SetMinter(&_Avatar.TransactOpts, _minter)
}

// SetMinter is a paid mutator transaction binding the contract method 0xfca3b5aa.
//
// Solidity: function setMinter(address _minter) returns()
func (_Avatar *AvatarTransactorSession) SetMinter(_minter common.Address) (*types.Transaction, error) {
	return _Avatar.Contract.SetMinter(&_Avatar.TransactOpts, _minter)
}

// SetSeeder is a paid mutator transaction binding the contract method 0xd50b31eb.
//
// Solidity: function setSeeder(address _seeder) returns()
func (_Avatar *AvatarTransactor) SetSeeder(opts *bind.TransactOpts, _seeder common.Address) (*types.Transaction, error) {
	return _Avatar.contract.Transact(opts, "setSeeder", _seeder)
}

// SetSeeder is a paid mutator transaction binding the contract method 0xd50b31eb.
//
// Solidity: function setSeeder(address _seeder) returns()
func (_Avatar *AvatarSession) SetSeeder(_seeder common.Address) (*types.Transaction, error) {
	return _Avatar.Contract.SetSeeder(&_Avatar.TransactOpts, _seeder)
}

// SetSeeder is a paid mutator transaction binding the contract method 0xd50b31eb.
//
// Solidity: function setSeeder(address _seeder) returns()
func (_Avatar *AvatarTransactorSession) SetSeeder(_seeder common.Address) (*types.Transaction, error) {
	return _Avatar.Contract.SetSeeder(&_Avatar.TransactOpts, _seeder)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_Avatar *AvatarTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Avatar.contract.Transact(opts, "transferFrom", from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_Avatar *AvatarSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Avatar.Contract.TransferFrom(&_Avatar.TransactOpts, from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_Avatar *AvatarTransactorSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Avatar.Contract.TransferFrom(&_Avatar.TransactOpts, from, to, tokenId)
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

// AvatarApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the Avatar contract.
type AvatarApprovalIterator struct {
	Event *AvatarApproval // Event containing the contract specifics and raw log

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
func (it *AvatarApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AvatarApproval)
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
		it.Event = new(AvatarApproval)
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
func (it *AvatarApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AvatarApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AvatarApproval represents a Approval event raised by the Avatar contract.
type AvatarApproval struct {
	Owner    common.Address
	Approved common.Address
	TokenId  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_Avatar *AvatarFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, approved []common.Address, tokenId []*big.Int) (*AvatarApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var approvedRule []interface{}
	for _, approvedItem := range approved {
		approvedRule = append(approvedRule, approvedItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Avatar.contract.FilterLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &AvatarApprovalIterator{contract: _Avatar.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_Avatar *AvatarFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *AvatarApproval, owner []common.Address, approved []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var approvedRule []interface{}
	for _, approvedItem := range approved {
		approvedRule = append(approvedRule, approvedItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Avatar.contract.WatchLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AvatarApproval)
				if err := _Avatar.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_Avatar *AvatarFilterer) ParseApproval(log types.Log) (*AvatarApproval, error) {
	event := new(AvatarApproval)
	if err := _Avatar.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AvatarApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the Avatar contract.
type AvatarApprovalForAllIterator struct {
	Event *AvatarApprovalForAll // Event containing the contract specifics and raw log

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
func (it *AvatarApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AvatarApprovalForAll)
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
		it.Event = new(AvatarApprovalForAll)
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
func (it *AvatarApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AvatarApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AvatarApprovalForAll represents a ApprovalForAll event raised by the Avatar contract.
type AvatarApprovalForAll struct {
	Owner    common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_Avatar *AvatarFilterer) FilterApprovalForAll(opts *bind.FilterOpts, owner []common.Address, operator []common.Address) (*AvatarApprovalForAllIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _Avatar.contract.FilterLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &AvatarApprovalForAllIterator{contract: _Avatar.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_Avatar *AvatarFilterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *AvatarApprovalForAll, owner []common.Address, operator []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _Avatar.contract.WatchLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AvatarApprovalForAll)
				if err := _Avatar.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
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

// ParseApprovalForAll is a log parse operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_Avatar *AvatarFilterer) ParseApprovalForAll(log types.Log) (*AvatarApprovalForAll, error) {
	event := new(AvatarApprovalForAll)
	if err := _Avatar.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AvatarDelegateChangedIterator is returned from FilterDelegateChanged and is used to iterate over the raw logs and unpacked data for DelegateChanged events raised by the Avatar contract.
type AvatarDelegateChangedIterator struct {
	Event *AvatarDelegateChanged // Event containing the contract specifics and raw log

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
func (it *AvatarDelegateChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AvatarDelegateChanged)
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
		it.Event = new(AvatarDelegateChanged)
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
func (it *AvatarDelegateChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AvatarDelegateChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AvatarDelegateChanged represents a DelegateChanged event raised by the Avatar contract.
type AvatarDelegateChanged struct {
	Delegator    common.Address
	FromDelegate common.Address
	ToDelegate   common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterDelegateChanged is a free log retrieval operation binding the contract event 0x3134e8a2e6d97e929a7e54011ea5485d7d196dd5f0ba4d4ef95803e8e3fc257f.
//
// Solidity: event DelegateChanged(address indexed delegator, address indexed fromDelegate, address indexed toDelegate)
func (_Avatar *AvatarFilterer) FilterDelegateChanged(opts *bind.FilterOpts, delegator []common.Address, fromDelegate []common.Address, toDelegate []common.Address) (*AvatarDelegateChangedIterator, error) {

	var delegatorRule []interface{}
	for _, delegatorItem := range delegator {
		delegatorRule = append(delegatorRule, delegatorItem)
	}
	var fromDelegateRule []interface{}
	for _, fromDelegateItem := range fromDelegate {
		fromDelegateRule = append(fromDelegateRule, fromDelegateItem)
	}
	var toDelegateRule []interface{}
	for _, toDelegateItem := range toDelegate {
		toDelegateRule = append(toDelegateRule, toDelegateItem)
	}

	logs, sub, err := _Avatar.contract.FilterLogs(opts, "DelegateChanged", delegatorRule, fromDelegateRule, toDelegateRule)
	if err != nil {
		return nil, err
	}
	return &AvatarDelegateChangedIterator{contract: _Avatar.contract, event: "DelegateChanged", logs: logs, sub: sub}, nil
}

// WatchDelegateChanged is a free log subscription operation binding the contract event 0x3134e8a2e6d97e929a7e54011ea5485d7d196dd5f0ba4d4ef95803e8e3fc257f.
//
// Solidity: event DelegateChanged(address indexed delegator, address indexed fromDelegate, address indexed toDelegate)
func (_Avatar *AvatarFilterer) WatchDelegateChanged(opts *bind.WatchOpts, sink chan<- *AvatarDelegateChanged, delegator []common.Address, fromDelegate []common.Address, toDelegate []common.Address) (event.Subscription, error) {

	var delegatorRule []interface{}
	for _, delegatorItem := range delegator {
		delegatorRule = append(delegatorRule, delegatorItem)
	}
	var fromDelegateRule []interface{}
	for _, fromDelegateItem := range fromDelegate {
		fromDelegateRule = append(fromDelegateRule, fromDelegateItem)
	}
	var toDelegateRule []interface{}
	for _, toDelegateItem := range toDelegate {
		toDelegateRule = append(toDelegateRule, toDelegateItem)
	}

	logs, sub, err := _Avatar.contract.WatchLogs(opts, "DelegateChanged", delegatorRule, fromDelegateRule, toDelegateRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AvatarDelegateChanged)
				if err := _Avatar.contract.UnpackLog(event, "DelegateChanged", log); err != nil {
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

// ParseDelegateChanged is a log parse operation binding the contract event 0x3134e8a2e6d97e929a7e54011ea5485d7d196dd5f0ba4d4ef95803e8e3fc257f.
//
// Solidity: event DelegateChanged(address indexed delegator, address indexed fromDelegate, address indexed toDelegate)
func (_Avatar *AvatarFilterer) ParseDelegateChanged(log types.Log) (*AvatarDelegateChanged, error) {
	event := new(AvatarDelegateChanged)
	if err := _Avatar.contract.UnpackLog(event, "DelegateChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AvatarDelegateVotesChangedIterator is returned from FilterDelegateVotesChanged and is used to iterate over the raw logs and unpacked data for DelegateVotesChanged events raised by the Avatar contract.
type AvatarDelegateVotesChangedIterator struct {
	Event *AvatarDelegateVotesChanged // Event containing the contract specifics and raw log

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
func (it *AvatarDelegateVotesChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AvatarDelegateVotesChanged)
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
		it.Event = new(AvatarDelegateVotesChanged)
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
func (it *AvatarDelegateVotesChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AvatarDelegateVotesChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AvatarDelegateVotesChanged represents a DelegateVotesChanged event raised by the Avatar contract.
type AvatarDelegateVotesChanged struct {
	Delegate        common.Address
	PreviousBalance *big.Int
	NewBalance      *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterDelegateVotesChanged is a free log retrieval operation binding the contract event 0xdec2bacdd2f05b59de34da9b523dff8be42e5e38e818c82fdb0bae774387a724.
//
// Solidity: event DelegateVotesChanged(address indexed delegate, uint256 previousBalance, uint256 newBalance)
func (_Avatar *AvatarFilterer) FilterDelegateVotesChanged(opts *bind.FilterOpts, delegate []common.Address) (*AvatarDelegateVotesChangedIterator, error) {

	var delegateRule []interface{}
	for _, delegateItem := range delegate {
		delegateRule = append(delegateRule, delegateItem)
	}

	logs, sub, err := _Avatar.contract.FilterLogs(opts, "DelegateVotesChanged", delegateRule)
	if err != nil {
		return nil, err
	}
	return &AvatarDelegateVotesChangedIterator{contract: _Avatar.contract, event: "DelegateVotesChanged", logs: logs, sub: sub}, nil
}

// WatchDelegateVotesChanged is a free log subscription operation binding the contract event 0xdec2bacdd2f05b59de34da9b523dff8be42e5e38e818c82fdb0bae774387a724.
//
// Solidity: event DelegateVotesChanged(address indexed delegate, uint256 previousBalance, uint256 newBalance)
func (_Avatar *AvatarFilterer) WatchDelegateVotesChanged(opts *bind.WatchOpts, sink chan<- *AvatarDelegateVotesChanged, delegate []common.Address) (event.Subscription, error) {

	var delegateRule []interface{}
	for _, delegateItem := range delegate {
		delegateRule = append(delegateRule, delegateItem)
	}

	logs, sub, err := _Avatar.contract.WatchLogs(opts, "DelegateVotesChanged", delegateRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AvatarDelegateVotesChanged)
				if err := _Avatar.contract.UnpackLog(event, "DelegateVotesChanged", log); err != nil {
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

// ParseDelegateVotesChanged is a log parse operation binding the contract event 0xdec2bacdd2f05b59de34da9b523dff8be42e5e38e818c82fdb0bae774387a724.
//
// Solidity: event DelegateVotesChanged(address indexed delegate, uint256 previousBalance, uint256 newBalance)
func (_Avatar *AvatarFilterer) ParseDelegateVotesChanged(log types.Log) (*AvatarDelegateVotesChanged, error) {
	event := new(AvatarDelegateVotesChanged)
	if err := _Avatar.contract.UnpackLog(event, "DelegateVotesChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AvatarDescriptorLockedIterator is returned from FilterDescriptorLocked and is used to iterate over the raw logs and unpacked data for DescriptorLocked events raised by the Avatar contract.
type AvatarDescriptorLockedIterator struct {
	Event *AvatarDescriptorLocked // Event containing the contract specifics and raw log

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
func (it *AvatarDescriptorLockedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AvatarDescriptorLocked)
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
		it.Event = new(AvatarDescriptorLocked)
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
func (it *AvatarDescriptorLockedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AvatarDescriptorLockedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AvatarDescriptorLocked represents a DescriptorLocked event raised by the Avatar contract.
type AvatarDescriptorLocked struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterDescriptorLocked is a free log retrieval operation binding the contract event 0x593e31e306c198bef259d839f7c6dc4ff7fc10c07f76fab193a210b03704105f.
//
// Solidity: event DescriptorLocked()
func (_Avatar *AvatarFilterer) FilterDescriptorLocked(opts *bind.FilterOpts) (*AvatarDescriptorLockedIterator, error) {

	logs, sub, err := _Avatar.contract.FilterLogs(opts, "DescriptorLocked")
	if err != nil {
		return nil, err
	}
	return &AvatarDescriptorLockedIterator{contract: _Avatar.contract, event: "DescriptorLocked", logs: logs, sub: sub}, nil
}

// WatchDescriptorLocked is a free log subscription operation binding the contract event 0x593e31e306c198bef259d839f7c6dc4ff7fc10c07f76fab193a210b03704105f.
//
// Solidity: event DescriptorLocked()
func (_Avatar *AvatarFilterer) WatchDescriptorLocked(opts *bind.WatchOpts, sink chan<- *AvatarDescriptorLocked) (event.Subscription, error) {

	logs, sub, err := _Avatar.contract.WatchLogs(opts, "DescriptorLocked")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AvatarDescriptorLocked)
				if err := _Avatar.contract.UnpackLog(event, "DescriptorLocked", log); err != nil {
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

// ParseDescriptorLocked is a log parse operation binding the contract event 0x593e31e306c198bef259d839f7c6dc4ff7fc10c07f76fab193a210b03704105f.
//
// Solidity: event DescriptorLocked()
func (_Avatar *AvatarFilterer) ParseDescriptorLocked(log types.Log) (*AvatarDescriptorLocked, error) {
	event := new(AvatarDescriptorLocked)
	if err := _Avatar.contract.UnpackLog(event, "DescriptorLocked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AvatarDescriptorUpdatedIterator is returned from FilterDescriptorUpdated and is used to iterate over the raw logs and unpacked data for DescriptorUpdated events raised by the Avatar contract.
type AvatarDescriptorUpdatedIterator struct {
	Event *AvatarDescriptorUpdated // Event containing the contract specifics and raw log

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
func (it *AvatarDescriptorUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AvatarDescriptorUpdated)
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
		it.Event = new(AvatarDescriptorUpdated)
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
func (it *AvatarDescriptorUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AvatarDescriptorUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AvatarDescriptorUpdated represents a DescriptorUpdated event raised by the Avatar contract.
type AvatarDescriptorUpdated struct {
	Descriptor common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterDescriptorUpdated is a free log retrieval operation binding the contract event 0x6e66ab22238a5471005895947c8f57db923c2a9c9c73180eff80864c21295c1b.
//
// Solidity: event DescriptorUpdated(address descriptor)
func (_Avatar *AvatarFilterer) FilterDescriptorUpdated(opts *bind.FilterOpts) (*AvatarDescriptorUpdatedIterator, error) {

	logs, sub, err := _Avatar.contract.FilterLogs(opts, "DescriptorUpdated")
	if err != nil {
		return nil, err
	}
	return &AvatarDescriptorUpdatedIterator{contract: _Avatar.contract, event: "DescriptorUpdated", logs: logs, sub: sub}, nil
}

// WatchDescriptorUpdated is a free log subscription operation binding the contract event 0x6e66ab22238a5471005895947c8f57db923c2a9c9c73180eff80864c21295c1b.
//
// Solidity: event DescriptorUpdated(address descriptor)
func (_Avatar *AvatarFilterer) WatchDescriptorUpdated(opts *bind.WatchOpts, sink chan<- *AvatarDescriptorUpdated) (event.Subscription, error) {

	logs, sub, err := _Avatar.contract.WatchLogs(opts, "DescriptorUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AvatarDescriptorUpdated)
				if err := _Avatar.contract.UnpackLog(event, "DescriptorUpdated", log); err != nil {
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

// ParseDescriptorUpdated is a log parse operation binding the contract event 0x6e66ab22238a5471005895947c8f57db923c2a9c9c73180eff80864c21295c1b.
//
// Solidity: event DescriptorUpdated(address descriptor)
func (_Avatar *AvatarFilterer) ParseDescriptorUpdated(log types.Log) (*AvatarDescriptorUpdated, error) {
	event := new(AvatarDescriptorUpdated)
	if err := _Avatar.contract.UnpackLog(event, "DescriptorUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AvatarMinterLockedIterator is returned from FilterMinterLocked and is used to iterate over the raw logs and unpacked data for MinterLocked events raised by the Avatar contract.
type AvatarMinterLockedIterator struct {
	Event *AvatarMinterLocked // Event containing the contract specifics and raw log

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
func (it *AvatarMinterLockedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AvatarMinterLocked)
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
		it.Event = new(AvatarMinterLocked)
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
func (it *AvatarMinterLockedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AvatarMinterLockedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AvatarMinterLocked represents a MinterLocked event raised by the Avatar contract.
type AvatarMinterLocked struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterMinterLocked is a free log retrieval operation binding the contract event 0x192417b3f16b1ce69e0c59b0376549666650245ffc05e4b2569089dda8589b66.
//
// Solidity: event MinterLocked()
func (_Avatar *AvatarFilterer) FilterMinterLocked(opts *bind.FilterOpts) (*AvatarMinterLockedIterator, error) {

	logs, sub, err := _Avatar.contract.FilterLogs(opts, "MinterLocked")
	if err != nil {
		return nil, err
	}
	return &AvatarMinterLockedIterator{contract: _Avatar.contract, event: "MinterLocked", logs: logs, sub: sub}, nil
}

// WatchMinterLocked is a free log subscription operation binding the contract event 0x192417b3f16b1ce69e0c59b0376549666650245ffc05e4b2569089dda8589b66.
//
// Solidity: event MinterLocked()
func (_Avatar *AvatarFilterer) WatchMinterLocked(opts *bind.WatchOpts, sink chan<- *AvatarMinterLocked) (event.Subscription, error) {

	logs, sub, err := _Avatar.contract.WatchLogs(opts, "MinterLocked")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AvatarMinterLocked)
				if err := _Avatar.contract.UnpackLog(event, "MinterLocked", log); err != nil {
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

// ParseMinterLocked is a log parse operation binding the contract event 0x192417b3f16b1ce69e0c59b0376549666650245ffc05e4b2569089dda8589b66.
//
// Solidity: event MinterLocked()
func (_Avatar *AvatarFilterer) ParseMinterLocked(log types.Log) (*AvatarMinterLocked, error) {
	event := new(AvatarMinterLocked)
	if err := _Avatar.contract.UnpackLog(event, "MinterLocked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AvatarMinterUpdatedIterator is returned from FilterMinterUpdated and is used to iterate over the raw logs and unpacked data for MinterUpdated events raised by the Avatar contract.
type AvatarMinterUpdatedIterator struct {
	Event *AvatarMinterUpdated // Event containing the contract specifics and raw log

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
func (it *AvatarMinterUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AvatarMinterUpdated)
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
		it.Event = new(AvatarMinterUpdated)
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
func (it *AvatarMinterUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AvatarMinterUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AvatarMinterUpdated represents a MinterUpdated event raised by the Avatar contract.
type AvatarMinterUpdated struct {
	Minter common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterMinterUpdated is a free log retrieval operation binding the contract event 0xad0f299ec81a386c98df0ac27dae11dd020ed1b56963c53a7292e7a3a314539a.
//
// Solidity: event MinterUpdated(address minter)
func (_Avatar *AvatarFilterer) FilterMinterUpdated(opts *bind.FilterOpts) (*AvatarMinterUpdatedIterator, error) {

	logs, sub, err := _Avatar.contract.FilterLogs(opts, "MinterUpdated")
	if err != nil {
		return nil, err
	}
	return &AvatarMinterUpdatedIterator{contract: _Avatar.contract, event: "MinterUpdated", logs: logs, sub: sub}, nil
}

// WatchMinterUpdated is a free log subscription operation binding the contract event 0xad0f299ec81a386c98df0ac27dae11dd020ed1b56963c53a7292e7a3a314539a.
//
// Solidity: event MinterUpdated(address minter)
func (_Avatar *AvatarFilterer) WatchMinterUpdated(opts *bind.WatchOpts, sink chan<- *AvatarMinterUpdated) (event.Subscription, error) {

	logs, sub, err := _Avatar.contract.WatchLogs(opts, "MinterUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AvatarMinterUpdated)
				if err := _Avatar.contract.UnpackLog(event, "MinterUpdated", log); err != nil {
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

// ParseMinterUpdated is a log parse operation binding the contract event 0xad0f299ec81a386c98df0ac27dae11dd020ed1b56963c53a7292e7a3a314539a.
//
// Solidity: event MinterUpdated(address minter)
func (_Avatar *AvatarFilterer) ParseMinterUpdated(log types.Log) (*AvatarMinterUpdated, error) {
	event := new(AvatarMinterUpdated)
	if err := _Avatar.contract.UnpackLog(event, "MinterUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AvatarNounBurnedIterator is returned from FilterNounBurned and is used to iterate over the raw logs and unpacked data for NounBurned events raised by the Avatar contract.
type AvatarNounBurnedIterator struct {
	Event *AvatarNounBurned // Event containing the contract specifics and raw log

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
func (it *AvatarNounBurnedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AvatarNounBurned)
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
		it.Event = new(AvatarNounBurned)
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
func (it *AvatarNounBurnedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AvatarNounBurnedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AvatarNounBurned represents a NounBurned event raised by the Avatar contract.
type AvatarNounBurned struct {
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterNounBurned is a free log retrieval operation binding the contract event 0x806be94a2ac8b92d74e99aa8add5a8e54528a01ec914a9e00d201a6480ed9863.
//
// Solidity: event NounBurned(uint256 indexed tokenId)
func (_Avatar *AvatarFilterer) FilterNounBurned(opts *bind.FilterOpts, tokenId []*big.Int) (*AvatarNounBurnedIterator, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Avatar.contract.FilterLogs(opts, "NounBurned", tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &AvatarNounBurnedIterator{contract: _Avatar.contract, event: "NounBurned", logs: logs, sub: sub}, nil
}

// WatchNounBurned is a free log subscription operation binding the contract event 0x806be94a2ac8b92d74e99aa8add5a8e54528a01ec914a9e00d201a6480ed9863.
//
// Solidity: event NounBurned(uint256 indexed tokenId)
func (_Avatar *AvatarFilterer) WatchNounBurned(opts *bind.WatchOpts, sink chan<- *AvatarNounBurned, tokenId []*big.Int) (event.Subscription, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Avatar.contract.WatchLogs(opts, "NounBurned", tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AvatarNounBurned)
				if err := _Avatar.contract.UnpackLog(event, "NounBurned", log); err != nil {
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

// ParseNounBurned is a log parse operation binding the contract event 0x806be94a2ac8b92d74e99aa8add5a8e54528a01ec914a9e00d201a6480ed9863.
//
// Solidity: event NounBurned(uint256 indexed tokenId)
func (_Avatar *AvatarFilterer) ParseNounBurned(log types.Log) (*AvatarNounBurned, error) {
	event := new(AvatarNounBurned)
	if err := _Avatar.contract.UnpackLog(event, "NounBurned", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AvatarNounCreatedIterator is returned from FilterNounCreated and is used to iterate over the raw logs and unpacked data for NounCreated events raised by the Avatar contract.
type AvatarNounCreatedIterator struct {
	Event *AvatarNounCreated // Event containing the contract specifics and raw log

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
func (it *AvatarNounCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AvatarNounCreated)
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
		it.Event = new(AvatarNounCreated)
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
func (it *AvatarNounCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AvatarNounCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AvatarNounCreated represents a NounCreated event raised by the Avatar contract.
type AvatarNounCreated struct {
	TokenId *big.Int
	Seed    IAvatarSeederSeed
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterNounCreated is a free log retrieval operation binding the contract event 0x1106ee9d020bfbb5ee34cf5535a5fbf024a011bd130078088cbf124ab3092478.
//
// Solidity: event NounCreated(uint256 indexed tokenId, (uint48,uint48,uint48,uint48,uint48) seed)
func (_Avatar *AvatarFilterer) FilterNounCreated(opts *bind.FilterOpts, tokenId []*big.Int) (*AvatarNounCreatedIterator, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Avatar.contract.FilterLogs(opts, "NounCreated", tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &AvatarNounCreatedIterator{contract: _Avatar.contract, event: "NounCreated", logs: logs, sub: sub}, nil
}

// WatchNounCreated is a free log subscription operation binding the contract event 0x1106ee9d020bfbb5ee34cf5535a5fbf024a011bd130078088cbf124ab3092478.
//
// Solidity: event NounCreated(uint256 indexed tokenId, (uint48,uint48,uint48,uint48,uint48) seed)
func (_Avatar *AvatarFilterer) WatchNounCreated(opts *bind.WatchOpts, sink chan<- *AvatarNounCreated, tokenId []*big.Int) (event.Subscription, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Avatar.contract.WatchLogs(opts, "NounCreated", tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AvatarNounCreated)
				if err := _Avatar.contract.UnpackLog(event, "NounCreated", log); err != nil {
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

// ParseNounCreated is a log parse operation binding the contract event 0x1106ee9d020bfbb5ee34cf5535a5fbf024a011bd130078088cbf124ab3092478.
//
// Solidity: event NounCreated(uint256 indexed tokenId, (uint48,uint48,uint48,uint48,uint48) seed)
func (_Avatar *AvatarFilterer) ParseNounCreated(log types.Log) (*AvatarNounCreated, error) {
	event := new(AvatarNounCreated)
	if err := _Avatar.contract.UnpackLog(event, "NounCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AvatarNounExchangedIterator is returned from FilterNounExchanged and is used to iterate over the raw logs and unpacked data for NounExchanged events raised by the Avatar contract.
type AvatarNounExchangedIterator struct {
	Event *AvatarNounExchanged // Event containing the contract specifics and raw log

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
func (it *AvatarNounExchangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AvatarNounExchanged)
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
		it.Event = new(AvatarNounExchanged)
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
func (it *AvatarNounExchangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AvatarNounExchangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AvatarNounExchanged represents a NounExchanged event raised by the Avatar contract.
type AvatarNounExchanged struct {
	TokenId  *big.Int
	PreOwner common.Address
	NewOwner common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterNounExchanged is a free log retrieval operation binding the contract event 0xc965bb22dd0227059a403fdf88fe10a4afe3f9b440348a374c19f41462fa0606.
//
// Solidity: event NounExchanged(uint256 tokenId, address preOwner, address newOwner)
func (_Avatar *AvatarFilterer) FilterNounExchanged(opts *bind.FilterOpts) (*AvatarNounExchangedIterator, error) {

	logs, sub, err := _Avatar.contract.FilterLogs(opts, "NounExchanged")
	if err != nil {
		return nil, err
	}
	return &AvatarNounExchangedIterator{contract: _Avatar.contract, event: "NounExchanged", logs: logs, sub: sub}, nil
}

// WatchNounExchanged is a free log subscription operation binding the contract event 0xc965bb22dd0227059a403fdf88fe10a4afe3f9b440348a374c19f41462fa0606.
//
// Solidity: event NounExchanged(uint256 tokenId, address preOwner, address newOwner)
func (_Avatar *AvatarFilterer) WatchNounExchanged(opts *bind.WatchOpts, sink chan<- *AvatarNounExchanged) (event.Subscription, error) {

	logs, sub, err := _Avatar.contract.WatchLogs(opts, "NounExchanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AvatarNounExchanged)
				if err := _Avatar.contract.UnpackLog(event, "NounExchanged", log); err != nil {
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

// ParseNounExchanged is a log parse operation binding the contract event 0xc965bb22dd0227059a403fdf88fe10a4afe3f9b440348a374c19f41462fa0606.
//
// Solidity: event NounExchanged(uint256 tokenId, address preOwner, address newOwner)
func (_Avatar *AvatarFilterer) ParseNounExchanged(log types.Log) (*AvatarNounExchanged, error) {
	event := new(AvatarNounExchanged)
	if err := _Avatar.contract.UnpackLog(event, "NounExchanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AvatarNounRecordedIterator is returned from FilterNounRecorded and is used to iterate over the raw logs and unpacked data for NounRecorded events raised by the Avatar contract.
type AvatarNounRecordedIterator struct {
	Event *AvatarNounRecorded // Event containing the contract specifics and raw log

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
func (it *AvatarNounRecordedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AvatarNounRecorded)
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
		it.Event = new(AvatarNounRecorded)
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
func (it *AvatarNounRecordedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AvatarNounRecordedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AvatarNounRecorded represents a NounRecorded event raised by the Avatar contract.
type AvatarNounRecorded struct {
	TokenURI string
	Author   common.Address
	Fhash    *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterNounRecorded is a free log retrieval operation binding the contract event 0xaeba371f7859a7ac3c00919908a62675633bf0f922723561dba1c9af458f1492.
//
// Solidity: event NounRecorded(string tokenURI, address author, uint256 fhash)
func (_Avatar *AvatarFilterer) FilterNounRecorded(opts *bind.FilterOpts) (*AvatarNounRecordedIterator, error) {

	logs, sub, err := _Avatar.contract.FilterLogs(opts, "NounRecorded")
	if err != nil {
		return nil, err
	}
	return &AvatarNounRecordedIterator{contract: _Avatar.contract, event: "NounRecorded", logs: logs, sub: sub}, nil
}

// WatchNounRecorded is a free log subscription operation binding the contract event 0xaeba371f7859a7ac3c00919908a62675633bf0f922723561dba1c9af458f1492.
//
// Solidity: event NounRecorded(string tokenURI, address author, uint256 fhash)
func (_Avatar *AvatarFilterer) WatchNounRecorded(opts *bind.WatchOpts, sink chan<- *AvatarNounRecorded) (event.Subscription, error) {

	logs, sub, err := _Avatar.contract.WatchLogs(opts, "NounRecorded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AvatarNounRecorded)
				if err := _Avatar.contract.UnpackLog(event, "NounRecorded", log); err != nil {
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

// ParseNounRecorded is a log parse operation binding the contract event 0xaeba371f7859a7ac3c00919908a62675633bf0f922723561dba1c9af458f1492.
//
// Solidity: event NounRecorded(string tokenURI, address author, uint256 fhash)
func (_Avatar *AvatarFilterer) ParseNounRecorded(log types.Log) (*AvatarNounRecorded, error) {
	event := new(AvatarNounRecorded)
	if err := _Avatar.contract.UnpackLog(event, "NounRecorded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AvatarNoundersDAOUpdatedIterator is returned from FilterNoundersDAOUpdated and is used to iterate over the raw logs and unpacked data for NoundersDAOUpdated events raised by the Avatar contract.
type AvatarNoundersDAOUpdatedIterator struct {
	Event *AvatarNoundersDAOUpdated // Event containing the contract specifics and raw log

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
func (it *AvatarNoundersDAOUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AvatarNoundersDAOUpdated)
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
		it.Event = new(AvatarNoundersDAOUpdated)
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
func (it *AvatarNoundersDAOUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AvatarNoundersDAOUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AvatarNoundersDAOUpdated represents a NoundersDAOUpdated event raised by the Avatar contract.
type AvatarNoundersDAOUpdated struct {
	NoundersDAO common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterNoundersDAOUpdated is a free log retrieval operation binding the contract event 0x3a0b923617f180781f3530e464cb4a8b9393e69f47607e4eb28d61cd87ce968c.
//
// Solidity: event NoundersDAOUpdated(address noundersDAO)
func (_Avatar *AvatarFilterer) FilterNoundersDAOUpdated(opts *bind.FilterOpts) (*AvatarNoundersDAOUpdatedIterator, error) {

	logs, sub, err := _Avatar.contract.FilterLogs(opts, "NoundersDAOUpdated")
	if err != nil {
		return nil, err
	}
	return &AvatarNoundersDAOUpdatedIterator{contract: _Avatar.contract, event: "NoundersDAOUpdated", logs: logs, sub: sub}, nil
}

// WatchNoundersDAOUpdated is a free log subscription operation binding the contract event 0x3a0b923617f180781f3530e464cb4a8b9393e69f47607e4eb28d61cd87ce968c.
//
// Solidity: event NoundersDAOUpdated(address noundersDAO)
func (_Avatar *AvatarFilterer) WatchNoundersDAOUpdated(opts *bind.WatchOpts, sink chan<- *AvatarNoundersDAOUpdated) (event.Subscription, error) {

	logs, sub, err := _Avatar.contract.WatchLogs(opts, "NoundersDAOUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AvatarNoundersDAOUpdated)
				if err := _Avatar.contract.UnpackLog(event, "NoundersDAOUpdated", log); err != nil {
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

// ParseNoundersDAOUpdated is a log parse operation binding the contract event 0x3a0b923617f180781f3530e464cb4a8b9393e69f47607e4eb28d61cd87ce968c.
//
// Solidity: event NoundersDAOUpdated(address noundersDAO)
func (_Avatar *AvatarFilterer) ParseNoundersDAOUpdated(log types.Log) (*AvatarNoundersDAOUpdated, error) {
	event := new(AvatarNoundersDAOUpdated)
	if err := _Avatar.contract.UnpackLog(event, "NoundersDAOUpdated", log); err != nil {
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

// AvatarSeederLockedIterator is returned from FilterSeederLocked and is used to iterate over the raw logs and unpacked data for SeederLocked events raised by the Avatar contract.
type AvatarSeederLockedIterator struct {
	Event *AvatarSeederLocked // Event containing the contract specifics and raw log

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
func (it *AvatarSeederLockedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AvatarSeederLocked)
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
		it.Event = new(AvatarSeederLocked)
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
func (it *AvatarSeederLockedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AvatarSeederLockedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AvatarSeederLocked represents a SeederLocked event raised by the Avatar contract.
type AvatarSeederLocked struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterSeederLocked is a free log retrieval operation binding the contract event 0xf59561f22794afcfb1e6be5c4733f5449fd167252a96b74bb06d341fb0dac7ed.
//
// Solidity: event SeederLocked()
func (_Avatar *AvatarFilterer) FilterSeederLocked(opts *bind.FilterOpts) (*AvatarSeederLockedIterator, error) {

	logs, sub, err := _Avatar.contract.FilterLogs(opts, "SeederLocked")
	if err != nil {
		return nil, err
	}
	return &AvatarSeederLockedIterator{contract: _Avatar.contract, event: "SeederLocked", logs: logs, sub: sub}, nil
}

// WatchSeederLocked is a free log subscription operation binding the contract event 0xf59561f22794afcfb1e6be5c4733f5449fd167252a96b74bb06d341fb0dac7ed.
//
// Solidity: event SeederLocked()
func (_Avatar *AvatarFilterer) WatchSeederLocked(opts *bind.WatchOpts, sink chan<- *AvatarSeederLocked) (event.Subscription, error) {

	logs, sub, err := _Avatar.contract.WatchLogs(opts, "SeederLocked")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AvatarSeederLocked)
				if err := _Avatar.contract.UnpackLog(event, "SeederLocked", log); err != nil {
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

// ParseSeederLocked is a log parse operation binding the contract event 0xf59561f22794afcfb1e6be5c4733f5449fd167252a96b74bb06d341fb0dac7ed.
//
// Solidity: event SeederLocked()
func (_Avatar *AvatarFilterer) ParseSeederLocked(log types.Log) (*AvatarSeederLocked, error) {
	event := new(AvatarSeederLocked)
	if err := _Avatar.contract.UnpackLog(event, "SeederLocked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AvatarSeederUpdatedIterator is returned from FilterSeederUpdated and is used to iterate over the raw logs and unpacked data for SeederUpdated events raised by the Avatar contract.
type AvatarSeederUpdatedIterator struct {
	Event *AvatarSeederUpdated // Event containing the contract specifics and raw log

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
func (it *AvatarSeederUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AvatarSeederUpdated)
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
		it.Event = new(AvatarSeederUpdated)
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
func (it *AvatarSeederUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AvatarSeederUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AvatarSeederUpdated represents a SeederUpdated event raised by the Avatar contract.
type AvatarSeederUpdated struct {
	Seeder common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterSeederUpdated is a free log retrieval operation binding the contract event 0xb3025222d01ce9a26c7f9d52bc3bfd0352366bd90a793c273fbfe1c81e0e288e.
//
// Solidity: event SeederUpdated(address seeder)
func (_Avatar *AvatarFilterer) FilterSeederUpdated(opts *bind.FilterOpts) (*AvatarSeederUpdatedIterator, error) {

	logs, sub, err := _Avatar.contract.FilterLogs(opts, "SeederUpdated")
	if err != nil {
		return nil, err
	}
	return &AvatarSeederUpdatedIterator{contract: _Avatar.contract, event: "SeederUpdated", logs: logs, sub: sub}, nil
}

// WatchSeederUpdated is a free log subscription operation binding the contract event 0xb3025222d01ce9a26c7f9d52bc3bfd0352366bd90a793c273fbfe1c81e0e288e.
//
// Solidity: event SeederUpdated(address seeder)
func (_Avatar *AvatarFilterer) WatchSeederUpdated(opts *bind.WatchOpts, sink chan<- *AvatarSeederUpdated) (event.Subscription, error) {

	logs, sub, err := _Avatar.contract.WatchLogs(opts, "SeederUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AvatarSeederUpdated)
				if err := _Avatar.contract.UnpackLog(event, "SeederUpdated", log); err != nil {
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

// ParseSeederUpdated is a log parse operation binding the contract event 0xb3025222d01ce9a26c7f9d52bc3bfd0352366bd90a793c273fbfe1c81e0e288e.
//
// Solidity: event SeederUpdated(address seeder)
func (_Avatar *AvatarFilterer) ParseSeederUpdated(log types.Log) (*AvatarSeederUpdated, error) {
	event := new(AvatarSeederUpdated)
	if err := _Avatar.contract.UnpackLog(event, "SeederUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AvatarTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Avatar contract.
type AvatarTransferIterator struct {
	Event *AvatarTransfer // Event containing the contract specifics and raw log

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
func (it *AvatarTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AvatarTransfer)
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
		it.Event = new(AvatarTransfer)
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
func (it *AvatarTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AvatarTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AvatarTransfer represents a Transfer event raised by the Avatar contract.
type AvatarTransfer struct {
	From    common.Address
	To      common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_Avatar *AvatarFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address, tokenId []*big.Int) (*AvatarTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Avatar.contract.FilterLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &AvatarTransferIterator{contract: _Avatar.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_Avatar *AvatarFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *AvatarTransfer, from []common.Address, to []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Avatar.contract.WatchLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AvatarTransfer)
				if err := _Avatar.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_Avatar *AvatarFilterer) ParseTransfer(log types.Log) (*AvatarTransfer, error) {
	event := new(AvatarTransfer)
	if err := _Avatar.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
