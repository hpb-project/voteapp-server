// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contracts

import (
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
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// LockABI is the input ABI used to generate the binding from.
const LockABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"nodeToLockAddr\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"nodeAddr\",\"type\":\"address\"}],\"name\":\"stake\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"deleteAdmin\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getAdmins\",\"outputs\":[{\"name\":\"\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"lockBal\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"kill\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"nodeAddr\",\"type\":\"address\"}],\"name\":\"withdraw\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"nodeAddr\",\"type\":\"address\"}],\"name\":\"fetchLockInfoByNodeAddr\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"},{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"addAdmin\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getOwner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getNodeContract\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"lockTime\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"adminMap\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"setNodeContract\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"constructor\"},{\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"ReceivedHpb\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"nodeAddr\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"addr\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"lockNum\",\"type\":\"uint256\"}],\"name\":\"LockBal\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"nodeAddr\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"addr\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"lockNum\",\"type\":\"uint256\"}],\"name\":\"withdrawBal\",\"type\":\"event\"}]"

// Lock is an auto generated Go binding around an Ethereum contract.
type Lock struct {
	LockCaller     // Read-only binding to the contract
	LockTransactor // Write-only binding to the contract
	LockFilterer   // Log filterer for contract events
}

// LockCaller is an auto generated read-only Go binding around an Ethereum contract.
type LockCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LockTransactor is an auto generated write-only Go binding around an Ethereum contract.
type LockTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LockFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type LockFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LockSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type LockSession struct {
	Contract     *Lock             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// LockCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type LockCallerSession struct {
	Contract *LockCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// LockTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type LockTransactorSession struct {
	Contract     *LockTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// LockRaw is an auto generated low-level Go binding around an Ethereum contract.
type LockRaw struct {
	Contract *Lock // Generic contract binding to access the raw methods on
}

// LockCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type LockCallerRaw struct {
	Contract *LockCaller // Generic read-only contract binding to access the raw methods on
}

// LockTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type LockTransactorRaw struct {
	Contract *LockTransactor // Generic write-only contract binding to access the raw methods on
}

// NewLock creates a new instance of Lock, bound to a specific deployed contract.
func NewLock(address common.Address, backend bind.ContractBackend) (*Lock, error) {
	contract, err := bindLock(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Lock{LockCaller: LockCaller{contract: contract}, LockTransactor: LockTransactor{contract: contract}, LockFilterer: LockFilterer{contract: contract}}, nil
}

// NewLockCaller creates a new read-only instance of Lock, bound to a specific deployed contract.
func NewLockCaller(address common.Address, caller bind.ContractCaller) (*LockCaller, error) {
	contract, err := bindLock(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &LockCaller{contract: contract}, nil
}

// NewLockTransactor creates a new write-only instance of Lock, bound to a specific deployed contract.
func NewLockTransactor(address common.Address, transactor bind.ContractTransactor) (*LockTransactor, error) {
	contract, err := bindLock(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &LockTransactor{contract: contract}, nil
}

// NewLockFilterer creates a new log filterer instance of Lock, bound to a specific deployed contract.
func NewLockFilterer(address common.Address, filterer bind.ContractFilterer) (*LockFilterer, error) {
	contract, err := bindLock(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &LockFilterer{contract: contract}, nil
}

// bindLock binds a generic wrapper to an already deployed contract.
func bindLock(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(LockABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Lock *LockRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Lock.Contract.LockCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Lock *LockRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Lock.Contract.LockTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Lock *LockRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Lock.Contract.LockTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Lock *LockCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Lock.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Lock *LockTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Lock.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Lock *LockTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Lock.Contract.contract.Transact(opts, method, params...)
}

// AdminMap is a free data retrieval call binding the contract method 0xdbbc830b.
//
// Solidity: function adminMap(address ) view returns(address)
func (_Lock *LockCaller) AdminMap(opts *bind.CallOpts, arg0 common.Address) (common.Address, error) {
	var out []interface{}
	err := _Lock.contract.Call(opts, &out, "adminMap", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AdminMap is a free data retrieval call binding the contract method 0xdbbc830b.
//
// Solidity: function adminMap(address ) view returns(address)
func (_Lock *LockSession) AdminMap(arg0 common.Address) (common.Address, error) {
	return _Lock.Contract.AdminMap(&_Lock.CallOpts, arg0)
}

// AdminMap is a free data retrieval call binding the contract method 0xdbbc830b.
//
// Solidity: function adminMap(address ) view returns(address)
func (_Lock *LockCallerSession) AdminMap(arg0 common.Address) (common.Address, error) {
	return _Lock.Contract.AdminMap(&_Lock.CallOpts, arg0)
}

// FetchLockInfoByNodeAddr is a free data retrieval call binding the contract method 0x695c6269.
//
// Solidity: function fetchLockInfoByNodeAddr(address nodeAddr) view returns(address, uint256)
func (_Lock *LockCaller) FetchLockInfoByNodeAddr(opts *bind.CallOpts, nodeAddr common.Address) (common.Address, *big.Int, error) {
	var out []interface{}
	err := _Lock.contract.Call(opts, &out, "fetchLockInfoByNodeAddr", nodeAddr)

	if err != nil {
		return *new(common.Address), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return out0, out1, err

}

// FetchLockInfoByNodeAddr is a free data retrieval call binding the contract method 0x695c6269.
//
// Solidity: function fetchLockInfoByNodeAddr(address nodeAddr) view returns(address, uint256)
func (_Lock *LockSession) FetchLockInfoByNodeAddr(nodeAddr common.Address) (common.Address, *big.Int, error) {
	return _Lock.Contract.FetchLockInfoByNodeAddr(&_Lock.CallOpts, nodeAddr)
}

// FetchLockInfoByNodeAddr is a free data retrieval call binding the contract method 0x695c6269.
//
// Solidity: function fetchLockInfoByNodeAddr(address nodeAddr) view returns(address, uint256)
func (_Lock *LockCallerSession) FetchLockInfoByNodeAddr(nodeAddr common.Address) (common.Address, *big.Int, error) {
	return _Lock.Contract.FetchLockInfoByNodeAddr(&_Lock.CallOpts, nodeAddr)
}

// GetAdmins is a free data retrieval call binding the contract method 0x31ae450b.
//
// Solidity: function getAdmins() view returns(address[])
func (_Lock *LockCaller) GetAdmins(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _Lock.contract.Call(opts, &out, "getAdmins")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetAdmins is a free data retrieval call binding the contract method 0x31ae450b.
//
// Solidity: function getAdmins() view returns(address[])
func (_Lock *LockSession) GetAdmins() ([]common.Address, error) {
	return _Lock.Contract.GetAdmins(&_Lock.CallOpts)
}

// GetAdmins is a free data retrieval call binding the contract method 0x31ae450b.
//
// Solidity: function getAdmins() view returns(address[])
func (_Lock *LockCallerSession) GetAdmins() ([]common.Address, error) {
	return _Lock.Contract.GetAdmins(&_Lock.CallOpts)
}

// GetNodeContract is a free data retrieval call binding the contract method 0x9fd2e0be.
//
// Solidity: function getNodeContract() view returns(address)
func (_Lock *LockCaller) GetNodeContract(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Lock.contract.Call(opts, &out, "getNodeContract")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetNodeContract is a free data retrieval call binding the contract method 0x9fd2e0be.
//
// Solidity: function getNodeContract() view returns(address)
func (_Lock *LockSession) GetNodeContract() (common.Address, error) {
	return _Lock.Contract.GetNodeContract(&_Lock.CallOpts)
}

// GetNodeContract is a free data retrieval call binding the contract method 0x9fd2e0be.
//
// Solidity: function getNodeContract() view returns(address)
func (_Lock *LockCallerSession) GetNodeContract() (common.Address, error) {
	return _Lock.Contract.GetNodeContract(&_Lock.CallOpts)
}

// GetOwner is a free data retrieval call binding the contract method 0x893d20e8.
//
// Solidity: function getOwner() view returns(address)
func (_Lock *LockCaller) GetOwner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Lock.contract.Call(opts, &out, "getOwner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetOwner is a free data retrieval call binding the contract method 0x893d20e8.
//
// Solidity: function getOwner() view returns(address)
func (_Lock *LockSession) GetOwner() (common.Address, error) {
	return _Lock.Contract.GetOwner(&_Lock.CallOpts)
}

// GetOwner is a free data retrieval call binding the contract method 0x893d20e8.
//
// Solidity: function getOwner() view returns(address)
func (_Lock *LockCallerSession) GetOwner() (common.Address, error) {
	return _Lock.Contract.GetOwner(&_Lock.CallOpts)
}

// LockBal is a free data retrieval call binding the contract method 0x31dddf75.
//
// Solidity: function lockBal(address ) view returns(uint256)
func (_Lock *LockCaller) LockBal(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Lock.contract.Call(opts, &out, "lockBal", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LockBal is a free data retrieval call binding the contract method 0x31dddf75.
//
// Solidity: function lockBal(address ) view returns(uint256)
func (_Lock *LockSession) LockBal(arg0 common.Address) (*big.Int, error) {
	return _Lock.Contract.LockBal(&_Lock.CallOpts, arg0)
}

// LockBal is a free data retrieval call binding the contract method 0x31dddf75.
//
// Solidity: function lockBal(address ) view returns(uint256)
func (_Lock *LockCallerSession) LockBal(arg0 common.Address) (*big.Int, error) {
	return _Lock.Contract.LockBal(&_Lock.CallOpts, arg0)
}

// LockTime is a free data retrieval call binding the contract method 0xa4beda63.
//
// Solidity: function lockTime(address ) view returns(uint256)
func (_Lock *LockCaller) LockTime(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Lock.contract.Call(opts, &out, "lockTime", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LockTime is a free data retrieval call binding the contract method 0xa4beda63.
//
// Solidity: function lockTime(address ) view returns(uint256)
func (_Lock *LockSession) LockTime(arg0 common.Address) (*big.Int, error) {
	return _Lock.Contract.LockTime(&_Lock.CallOpts, arg0)
}

// LockTime is a free data retrieval call binding the contract method 0xa4beda63.
//
// Solidity: function lockTime(address ) view returns(uint256)
func (_Lock *LockCallerSession) LockTime(arg0 common.Address) (*big.Int, error) {
	return _Lock.Contract.LockTime(&_Lock.CallOpts, arg0)
}

// NodeToLockAddr is a free data retrieval call binding the contract method 0x022277b1.
//
// Solidity: function nodeToLockAddr(address ) view returns(address)
func (_Lock *LockCaller) NodeToLockAddr(opts *bind.CallOpts, arg0 common.Address) (common.Address, error) {
	var out []interface{}
	err := _Lock.contract.Call(opts, &out, "nodeToLockAddr", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// NodeToLockAddr is a free data retrieval call binding the contract method 0x022277b1.
//
// Solidity: function nodeToLockAddr(address ) view returns(address)
func (_Lock *LockSession) NodeToLockAddr(arg0 common.Address) (common.Address, error) {
	return _Lock.Contract.NodeToLockAddr(&_Lock.CallOpts, arg0)
}

// NodeToLockAddr is a free data retrieval call binding the contract method 0x022277b1.
//
// Solidity: function nodeToLockAddr(address ) view returns(address)
func (_Lock *LockCallerSession) NodeToLockAddr(arg0 common.Address) (common.Address, error) {
	return _Lock.Contract.NodeToLockAddr(&_Lock.CallOpts, arg0)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Lock *LockCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Lock.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Lock *LockSession) Owner() (common.Address, error) {
	return _Lock.Contract.Owner(&_Lock.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Lock *LockCallerSession) Owner() (common.Address, error) {
	return _Lock.Contract.Owner(&_Lock.CallOpts)
}

// AddAdmin is a paid mutator transaction binding the contract method 0x70480275.
//
// Solidity: function addAdmin(address addr) returns(bool)
func (_Lock *LockTransactor) AddAdmin(opts *bind.TransactOpts, addr common.Address) (*types.Transaction, error) {
	return _Lock.contract.Transact(opts, "addAdmin", addr)
}

// AddAdmin is a paid mutator transaction binding the contract method 0x70480275.
//
// Solidity: function addAdmin(address addr) returns(bool)
func (_Lock *LockSession) AddAdmin(addr common.Address) (*types.Transaction, error) {
	return _Lock.Contract.AddAdmin(&_Lock.TransactOpts, addr)
}

// AddAdmin is a paid mutator transaction binding the contract method 0x70480275.
//
// Solidity: function addAdmin(address addr) returns(bool)
func (_Lock *LockTransactorSession) AddAdmin(addr common.Address) (*types.Transaction, error) {
	return _Lock.Contract.AddAdmin(&_Lock.TransactOpts, addr)
}

// DeleteAdmin is a paid mutator transaction binding the contract method 0x27e1f7df.
//
// Solidity: function deleteAdmin(address addr) returns(bool)
func (_Lock *LockTransactor) DeleteAdmin(opts *bind.TransactOpts, addr common.Address) (*types.Transaction, error) {
	return _Lock.contract.Transact(opts, "deleteAdmin", addr)
}

// DeleteAdmin is a paid mutator transaction binding the contract method 0x27e1f7df.
//
// Solidity: function deleteAdmin(address addr) returns(bool)
func (_Lock *LockSession) DeleteAdmin(addr common.Address) (*types.Transaction, error) {
	return _Lock.Contract.DeleteAdmin(&_Lock.TransactOpts, addr)
}

// DeleteAdmin is a paid mutator transaction binding the contract method 0x27e1f7df.
//
// Solidity: function deleteAdmin(address addr) returns(bool)
func (_Lock *LockTransactorSession) DeleteAdmin(addr common.Address) (*types.Transaction, error) {
	return _Lock.Contract.DeleteAdmin(&_Lock.TransactOpts, addr)
}

// Kill is a paid mutator transaction binding the contract method 0x41c0e1b5.
//
// Solidity: function kill() payable returns(bool)
func (_Lock *LockTransactor) Kill(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Lock.contract.Transact(opts, "kill")
}

// Kill is a paid mutator transaction binding the contract method 0x41c0e1b5.
//
// Solidity: function kill() payable returns(bool)
func (_Lock *LockSession) Kill() (*types.Transaction, error) {
	return _Lock.Contract.Kill(&_Lock.TransactOpts)
}

// Kill is a paid mutator transaction binding the contract method 0x41c0e1b5.
//
// Solidity: function kill() payable returns(bool)
func (_Lock *LockTransactorSession) Kill() (*types.Transaction, error) {
	return _Lock.Contract.Kill(&_Lock.TransactOpts)
}

// SetNodeContract is a paid mutator transaction binding the contract method 0xf50fd926.
//
// Solidity: function setNodeContract(address addr) returns()
func (_Lock *LockTransactor) SetNodeContract(opts *bind.TransactOpts, addr common.Address) (*types.Transaction, error) {
	return _Lock.contract.Transact(opts, "setNodeContract", addr)
}

// SetNodeContract is a paid mutator transaction binding the contract method 0xf50fd926.
//
// Solidity: function setNodeContract(address addr) returns()
func (_Lock *LockSession) SetNodeContract(addr common.Address) (*types.Transaction, error) {
	return _Lock.Contract.SetNodeContract(&_Lock.TransactOpts, addr)
}

// SetNodeContract is a paid mutator transaction binding the contract method 0xf50fd926.
//
// Solidity: function setNodeContract(address addr) returns()
func (_Lock *LockTransactorSession) SetNodeContract(addr common.Address) (*types.Transaction, error) {
	return _Lock.Contract.SetNodeContract(&_Lock.TransactOpts, addr)
}

// Stake is a paid mutator transaction binding the contract method 0x26476204.
//
// Solidity: function stake(address nodeAddr) payable returns(bool)
func (_Lock *LockTransactor) Stake(opts *bind.TransactOpts, nodeAddr common.Address) (*types.Transaction, error) {
	return _Lock.contract.Transact(opts, "stake", nodeAddr)
}

// Stake is a paid mutator transaction binding the contract method 0x26476204.
//
// Solidity: function stake(address nodeAddr) payable returns(bool)
func (_Lock *LockSession) Stake(nodeAddr common.Address) (*types.Transaction, error) {
	return _Lock.Contract.Stake(&_Lock.TransactOpts, nodeAddr)
}

// Stake is a paid mutator transaction binding the contract method 0x26476204.
//
// Solidity: function stake(address nodeAddr) payable returns(bool)
func (_Lock *LockTransactorSession) Stake(nodeAddr common.Address) (*types.Transaction, error) {
	return _Lock.Contract.Stake(&_Lock.TransactOpts, nodeAddr)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns(bool)
func (_Lock *LockTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Lock.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns(bool)
func (_Lock *LockSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Lock.Contract.TransferOwnership(&_Lock.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns(bool)
func (_Lock *LockTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Lock.Contract.TransferOwnership(&_Lock.TransactOpts, newOwner)
}

// Withdraw is a paid mutator transaction binding the contract method 0x51cff8d9.
//
// Solidity: function withdraw(address nodeAddr) payable returns(bool)
func (_Lock *LockTransactor) Withdraw(opts *bind.TransactOpts, nodeAddr common.Address) (*types.Transaction, error) {
	return _Lock.contract.Transact(opts, "withdraw", nodeAddr)
}

// Withdraw is a paid mutator transaction binding the contract method 0x51cff8d9.
//
// Solidity: function withdraw(address nodeAddr) payable returns(bool)
func (_Lock *LockSession) Withdraw(nodeAddr common.Address) (*types.Transaction, error) {
	return _Lock.Contract.Withdraw(&_Lock.TransactOpts, nodeAddr)
}

// Withdraw is a paid mutator transaction binding the contract method 0x51cff8d9.
//
// Solidity: function withdraw(address nodeAddr) payable returns(bool)
func (_Lock *LockTransactorSession) Withdraw(nodeAddr common.Address) (*types.Transaction, error) {
	return _Lock.Contract.Withdraw(&_Lock.TransactOpts, nodeAddr)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_Lock *LockTransactor) Fallback(opts *bind.TransactOpts, calldata []byte) (*types.Transaction, error) {
	return _Lock.contract.RawTransact(opts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_Lock *LockSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _Lock.Contract.Fallback(&_Lock.TransactOpts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_Lock *LockTransactorSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _Lock.Contract.Fallback(&_Lock.TransactOpts, calldata)
}

// LockLockBalIterator is returned from FilterLockBal and is used to iterate over the raw logs and unpacked data for LockBal events raised by the Lock contract.
type LockLockBalIterator struct {
	Event *LockLockBal // Event containing the contract specifics and raw log

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
func (it *LockLockBalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LockLockBal)
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
		it.Event = new(LockLockBal)
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
func (it *LockLockBalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LockLockBalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LockLockBal represents a LockBal event raised by the Lock contract.
type LockLockBal struct {
	NodeAddr common.Address
	Addr     common.Address
	LockNum  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterLockBal is a free log retrieval operation binding the contract event 0xb87d4f14e7ba38c29817fe73e71e3e3a24821eb336c4ea274b891072dbacdf34.
//
// Solidity: event LockBal(address indexed nodeAddr, address indexed addr, uint256 indexed lockNum)
func (_Lock *LockFilterer) FilterLockBal(opts *bind.FilterOpts, nodeAddr []common.Address, addr []common.Address, lockNum []*big.Int) (*LockLockBalIterator, error) {

	var nodeAddrRule []interface{}
	for _, nodeAddrItem := range nodeAddr {
		nodeAddrRule = append(nodeAddrRule, nodeAddrItem)
	}
	var addrRule []interface{}
	for _, addrItem := range addr {
		addrRule = append(addrRule, addrItem)
	}
	var lockNumRule []interface{}
	for _, lockNumItem := range lockNum {
		lockNumRule = append(lockNumRule, lockNumItem)
	}

	logs, sub, err := _Lock.contract.FilterLogs(opts, "LockBal", nodeAddrRule, addrRule, lockNumRule)
	if err != nil {
		return nil, err
	}
	return &LockLockBalIterator{contract: _Lock.contract, event: "LockBal", logs: logs, sub: sub}, nil
}

// WatchLockBal is a free log subscription operation binding the contract event 0xb87d4f14e7ba38c29817fe73e71e3e3a24821eb336c4ea274b891072dbacdf34.
//
// Solidity: event LockBal(address indexed nodeAddr, address indexed addr, uint256 indexed lockNum)
func (_Lock *LockFilterer) WatchLockBal(opts *bind.WatchOpts, sink chan<- *LockLockBal, nodeAddr []common.Address, addr []common.Address, lockNum []*big.Int) (event.Subscription, error) {

	var nodeAddrRule []interface{}
	for _, nodeAddrItem := range nodeAddr {
		nodeAddrRule = append(nodeAddrRule, nodeAddrItem)
	}
	var addrRule []interface{}
	for _, addrItem := range addr {
		addrRule = append(addrRule, addrItem)
	}
	var lockNumRule []interface{}
	for _, lockNumItem := range lockNum {
		lockNumRule = append(lockNumRule, lockNumItem)
	}

	logs, sub, err := _Lock.contract.WatchLogs(opts, "LockBal", nodeAddrRule, addrRule, lockNumRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LockLockBal)
				if err := _Lock.contract.UnpackLog(event, "LockBal", log); err != nil {
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

// ParseLockBal is a log parse operation binding the contract event 0xb87d4f14e7ba38c29817fe73e71e3e3a24821eb336c4ea274b891072dbacdf34.
//
// Solidity: event LockBal(address indexed nodeAddr, address indexed addr, uint256 indexed lockNum)
func (_Lock *LockFilterer) ParseLockBal(log types.Log) (*LockLockBal, error) {
	event := new(LockLockBal)
	if err := _Lock.contract.UnpackLog(event, "LockBal", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LockReceivedHpbIterator is returned from FilterReceivedHpb and is used to iterate over the raw logs and unpacked data for ReceivedHpb events raised by the Lock contract.
type LockReceivedHpbIterator struct {
	Event *LockReceivedHpb // Event containing the contract specifics and raw log

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
func (it *LockReceivedHpbIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LockReceivedHpb)
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
		it.Event = new(LockReceivedHpb)
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
func (it *LockReceivedHpbIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LockReceivedHpbIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LockReceivedHpb represents a ReceivedHpb event raised by the Lock contract.
type LockReceivedHpb struct {
	Sender common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterReceivedHpb is a free log retrieval operation binding the contract event 0x7129701436f0cdc265d1e2cda298e8a1ccd6ed5fce7f69343e16530b07a2e06e.
//
// Solidity: event ReceivedHpb(address indexed sender, uint256 amount)
func (_Lock *LockFilterer) FilterReceivedHpb(opts *bind.FilterOpts, sender []common.Address) (*LockReceivedHpbIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Lock.contract.FilterLogs(opts, "ReceivedHpb", senderRule)
	if err != nil {
		return nil, err
	}
	return &LockReceivedHpbIterator{contract: _Lock.contract, event: "ReceivedHpb", logs: logs, sub: sub}, nil
}

// WatchReceivedHpb is a free log subscription operation binding the contract event 0x7129701436f0cdc265d1e2cda298e8a1ccd6ed5fce7f69343e16530b07a2e06e.
//
// Solidity: event ReceivedHpb(address indexed sender, uint256 amount)
func (_Lock *LockFilterer) WatchReceivedHpb(opts *bind.WatchOpts, sink chan<- *LockReceivedHpb, sender []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Lock.contract.WatchLogs(opts, "ReceivedHpb", senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LockReceivedHpb)
				if err := _Lock.contract.UnpackLog(event, "ReceivedHpb", log); err != nil {
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

// ParseReceivedHpb is a log parse operation binding the contract event 0x7129701436f0cdc265d1e2cda298e8a1ccd6ed5fce7f69343e16530b07a2e06e.
//
// Solidity: event ReceivedHpb(address indexed sender, uint256 amount)
func (_Lock *LockFilterer) ParseReceivedHpb(log types.Log) (*LockReceivedHpb, error) {
	event := new(LockReceivedHpb)
	if err := _Lock.contract.UnpackLog(event, "ReceivedHpb", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LockWithdrawBalIterator is returned from FilterWithdrawBal and is used to iterate over the raw logs and unpacked data for WithdrawBal events raised by the Lock contract.
type LockWithdrawBalIterator struct {
	Event *LockWithdrawBal // Event containing the contract specifics and raw log

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
func (it *LockWithdrawBalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LockWithdrawBal)
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
		it.Event = new(LockWithdrawBal)
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
func (it *LockWithdrawBalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LockWithdrawBalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LockWithdrawBal represents a WithdrawBal event raised by the Lock contract.
type LockWithdrawBal struct {
	NodeAddr common.Address
	Addr     common.Address
	LockNum  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterWithdrawBal is a free log retrieval operation binding the contract event 0x8f2663a76951a6b5fefd39a998d0ff0d6fc9abc3ea3b14431557da631aea1df4.
//
// Solidity: event withdrawBal(address indexed nodeAddr, address indexed addr, uint256 indexed lockNum)
func (_Lock *LockFilterer) FilterWithdrawBal(opts *bind.FilterOpts, nodeAddr []common.Address, addr []common.Address, lockNum []*big.Int) (*LockWithdrawBalIterator, error) {

	var nodeAddrRule []interface{}
	for _, nodeAddrItem := range nodeAddr {
		nodeAddrRule = append(nodeAddrRule, nodeAddrItem)
	}
	var addrRule []interface{}
	for _, addrItem := range addr {
		addrRule = append(addrRule, addrItem)
	}
	var lockNumRule []interface{}
	for _, lockNumItem := range lockNum {
		lockNumRule = append(lockNumRule, lockNumItem)
	}

	logs, sub, err := _Lock.contract.FilterLogs(opts, "withdrawBal", nodeAddrRule, addrRule, lockNumRule)
	if err != nil {
		return nil, err
	}
	return &LockWithdrawBalIterator{contract: _Lock.contract, event: "withdrawBal", logs: logs, sub: sub}, nil
}

// WatchWithdrawBal is a free log subscription operation binding the contract event 0x8f2663a76951a6b5fefd39a998d0ff0d6fc9abc3ea3b14431557da631aea1df4.
//
// Solidity: event withdrawBal(address indexed nodeAddr, address indexed addr, uint256 indexed lockNum)
func (_Lock *LockFilterer) WatchWithdrawBal(opts *bind.WatchOpts, sink chan<- *LockWithdrawBal, nodeAddr []common.Address, addr []common.Address, lockNum []*big.Int) (event.Subscription, error) {

	var nodeAddrRule []interface{}
	for _, nodeAddrItem := range nodeAddr {
		nodeAddrRule = append(nodeAddrRule, nodeAddrItem)
	}
	var addrRule []interface{}
	for _, addrItem := range addr {
		addrRule = append(addrRule, addrItem)
	}
	var lockNumRule []interface{}
	for _, lockNumItem := range lockNum {
		lockNumRule = append(lockNumRule, lockNumItem)
	}

	logs, sub, err := _Lock.contract.WatchLogs(opts, "withdrawBal", nodeAddrRule, addrRule, lockNumRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LockWithdrawBal)
				if err := _Lock.contract.UnpackLog(event, "withdrawBal", log); err != nil {
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

// ParseWithdrawBal is a log parse operation binding the contract event 0x8f2663a76951a6b5fefd39a998d0ff0d6fc9abc3ea3b14431557da631aea1df4.
//
// Solidity: event withdrawBal(address indexed nodeAddr, address indexed addr, uint256 indexed lockNum)
func (_Lock *LockFilterer) ParseWithdrawBal(log types.Log) (*LockWithdrawBal, error) {
	event := new(LockWithdrawBal)
	if err := _Lock.contract.UnpackLog(event, "withdrawBal", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
