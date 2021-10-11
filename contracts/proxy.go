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

// ProxyABI is the input ABI used to generate the binding from.
const ProxyABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"fetchAllVoteResult\",\"outputs\":[{\"name\":\"\",\"type\":\"address[]\"},{\"name\":\"\",\"type\":\"uint256[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"deleteAdmin\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getAdmins\",\"outputs\":[{\"name\":\"\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"addAdmin\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"lockaddr\",\"type\":\"address\"}],\"name\":\"setlockcontract\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getOwner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"voteaddr\",\"type\":\"address\"}],\"name\":\"setvotecontract\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"fetchAllHolderAddrs\",\"outputs\":[{\"name\":\"\",\"type\":\"address[]\"},{\"name\":\"\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"nodeaddr\",\"type\":\"address\"}],\"name\":\"setnodecontract\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"adminMap\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getAllHpbNodes\",\"outputs\":[{\"name\":\"\",\"type\":\"address[]\"},{\"name\":\"\",\"type\":\"bytes32[]\"},{\"name\":\"\",\"type\":\"bytes32[]\"},{\"name\":\"\",\"type\":\"bytes32[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getcontract\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"},{\"name\":\"\",\"type\":\"address\"},{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"constructor\"}]"

// Proxy is an auto generated Go binding around an Ethereum contract.
type Proxy struct {
	ProxyCaller     // Read-only binding to the contract
	ProxyTransactor // Write-only binding to the contract
	ProxyFilterer   // Log filterer for contract events
}

// ProxyCaller is an auto generated read-only Go binding around an Ethereum contract.
type ProxyCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ProxyTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ProxyTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ProxyFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ProxyFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ProxySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ProxySession struct {
	Contract     *Proxy            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ProxyCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ProxyCallerSession struct {
	Contract *ProxyCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// ProxyTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ProxyTransactorSession struct {
	Contract     *ProxyTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ProxyRaw is an auto generated low-level Go binding around an Ethereum contract.
type ProxyRaw struct {
	Contract *Proxy // Generic contract binding to access the raw methods on
}

// ProxyCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ProxyCallerRaw struct {
	Contract *ProxyCaller // Generic read-only contract binding to access the raw methods on
}

// ProxyTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ProxyTransactorRaw struct {
	Contract *ProxyTransactor // Generic write-only contract binding to access the raw methods on
}

// NewProxy creates a new instance of Proxy, bound to a specific deployed contract.
func NewProxy(address common.Address, backend bind.ContractBackend) (*Proxy, error) {
	contract, err := bindProxy(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Proxy{ProxyCaller: ProxyCaller{contract: contract}, ProxyTransactor: ProxyTransactor{contract: contract}, ProxyFilterer: ProxyFilterer{contract: contract}}, nil
}

// NewProxyCaller creates a new read-only instance of Proxy, bound to a specific deployed contract.
func NewProxyCaller(address common.Address, caller bind.ContractCaller) (*ProxyCaller, error) {
	contract, err := bindProxy(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ProxyCaller{contract: contract}, nil
}

// NewProxyTransactor creates a new write-only instance of Proxy, bound to a specific deployed contract.
func NewProxyTransactor(address common.Address, transactor bind.ContractTransactor) (*ProxyTransactor, error) {
	contract, err := bindProxy(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ProxyTransactor{contract: contract}, nil
}

// NewProxyFilterer creates a new log filterer instance of Proxy, bound to a specific deployed contract.
func NewProxyFilterer(address common.Address, filterer bind.ContractFilterer) (*ProxyFilterer, error) {
	contract, err := bindProxy(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ProxyFilterer{contract: contract}, nil
}

// bindProxy binds a generic wrapper to an already deployed contract.
func bindProxy(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ProxyABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Proxy *ProxyRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Proxy.Contract.ProxyCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Proxy *ProxyRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Proxy.Contract.ProxyTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Proxy *ProxyRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Proxy.Contract.ProxyTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Proxy *ProxyCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Proxy.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Proxy *ProxyTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Proxy.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Proxy *ProxyTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Proxy.Contract.contract.Transact(opts, method, params...)
}

// AdminMap is a free data retrieval call binding the contract method 0xdbbc830b.
//
// Solidity: function adminMap(address ) view returns(address)
func (_Proxy *ProxyCaller) AdminMap(opts *bind.CallOpts, arg0 common.Address) (common.Address, error) {
	var out []interface{}
	err := _Proxy.contract.Call(opts, &out, "adminMap", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AdminMap is a free data retrieval call binding the contract method 0xdbbc830b.
//
// Solidity: function adminMap(address ) view returns(address)
func (_Proxy *ProxySession) AdminMap(arg0 common.Address) (common.Address, error) {
	return _Proxy.Contract.AdminMap(&_Proxy.CallOpts, arg0)
}

// AdminMap is a free data retrieval call binding the contract method 0xdbbc830b.
//
// Solidity: function adminMap(address ) view returns(address)
func (_Proxy *ProxyCallerSession) AdminMap(arg0 common.Address) (common.Address, error) {
	return _Proxy.Contract.AdminMap(&_Proxy.CallOpts, arg0)
}

// FetchAllHolderAddrs is a free data retrieval call binding the contract method 0xab2d2c0b.
//
// Solidity: function fetchAllHolderAddrs() view returns(address[], address[])
func (_Proxy *ProxyCaller) FetchAllHolderAddrs(opts *bind.CallOpts) ([]common.Address, []common.Address, error) {
	var out []interface{}
	err := _Proxy.contract.Call(opts, &out, "fetchAllHolderAddrs")

	if err != nil {
		return *new([]common.Address), *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)
	out1 := *abi.ConvertType(out[1], new([]common.Address)).(*[]common.Address)

	return out0, out1, err

}

// FetchAllHolderAddrs is a free data retrieval call binding the contract method 0xab2d2c0b.
//
// Solidity: function fetchAllHolderAddrs() view returns(address[], address[])
func (_Proxy *ProxySession) FetchAllHolderAddrs() ([]common.Address, []common.Address, error) {
	return _Proxy.Contract.FetchAllHolderAddrs(&_Proxy.CallOpts)
}

// FetchAllHolderAddrs is a free data retrieval call binding the contract method 0xab2d2c0b.
//
// Solidity: function fetchAllHolderAddrs() view returns(address[], address[])
func (_Proxy *ProxyCallerSession) FetchAllHolderAddrs() ([]common.Address, []common.Address, error) {
	return _Proxy.Contract.FetchAllHolderAddrs(&_Proxy.CallOpts)
}

// FetchAllVoteResult is a free data retrieval call binding the contract method 0x0304363a.
//
// Solidity: function fetchAllVoteResult() view returns(address[], uint256[])
func (_Proxy *ProxyCaller) FetchAllVoteResult(opts *bind.CallOpts) ([]common.Address, []*big.Int, error) {
	var out []interface{}
	err := _Proxy.contract.Call(opts, &out, "fetchAllVoteResult")

	if err != nil {
		return *new([]common.Address), *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)
	out1 := *abi.ConvertType(out[1], new([]*big.Int)).(*[]*big.Int)

	return out0, out1, err

}

// FetchAllVoteResult is a free data retrieval call binding the contract method 0x0304363a.
//
// Solidity: function fetchAllVoteResult() view returns(address[], uint256[])
func (_Proxy *ProxySession) FetchAllVoteResult() ([]common.Address, []*big.Int, error) {
	return _Proxy.Contract.FetchAllVoteResult(&_Proxy.CallOpts)
}

// FetchAllVoteResult is a free data retrieval call binding the contract method 0x0304363a.
//
// Solidity: function fetchAllVoteResult() view returns(address[], uint256[])
func (_Proxy *ProxyCallerSession) FetchAllVoteResult() ([]common.Address, []*big.Int, error) {
	return _Proxy.Contract.FetchAllVoteResult(&_Proxy.CallOpts)
}

// GetAdmins is a free data retrieval call binding the contract method 0x31ae450b.
//
// Solidity: function getAdmins() view returns(address[])
func (_Proxy *ProxyCaller) GetAdmins(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _Proxy.contract.Call(opts, &out, "getAdmins")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetAdmins is a free data retrieval call binding the contract method 0x31ae450b.
//
// Solidity: function getAdmins() view returns(address[])
func (_Proxy *ProxySession) GetAdmins() ([]common.Address, error) {
	return _Proxy.Contract.GetAdmins(&_Proxy.CallOpts)
}

// GetAdmins is a free data retrieval call binding the contract method 0x31ae450b.
//
// Solidity: function getAdmins() view returns(address[])
func (_Proxy *ProxyCallerSession) GetAdmins() ([]common.Address, error) {
	return _Proxy.Contract.GetAdmins(&_Proxy.CallOpts)
}

// GetAllHpbNodes is a free data retrieval call binding the contract method 0xdbe01790.
//
// Solidity: function getAllHpbNodes() view returns(address[], bytes32[], bytes32[], bytes32[])
func (_Proxy *ProxyCaller) GetAllHpbNodes(opts *bind.CallOpts) ([]common.Address, [][32]byte, [][32]byte, [][32]byte, error) {
	var out []interface{}
	err := _Proxy.contract.Call(opts, &out, "getAllHpbNodes")

	if err != nil {
		return *new([]common.Address), *new([][32]byte), *new([][32]byte), *new([][32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)
	out1 := *abi.ConvertType(out[1], new([][32]byte)).(*[][32]byte)
	out2 := *abi.ConvertType(out[2], new([][32]byte)).(*[][32]byte)
	out3 := *abi.ConvertType(out[3], new([][32]byte)).(*[][32]byte)

	return out0, out1, out2, out3, err

}

// GetAllHpbNodes is a free data retrieval call binding the contract method 0xdbe01790.
//
// Solidity: function getAllHpbNodes() view returns(address[], bytes32[], bytes32[], bytes32[])
func (_Proxy *ProxySession) GetAllHpbNodes() ([]common.Address, [][32]byte, [][32]byte, [][32]byte, error) {
	return _Proxy.Contract.GetAllHpbNodes(&_Proxy.CallOpts)
}

// GetAllHpbNodes is a free data retrieval call binding the contract method 0xdbe01790.
//
// Solidity: function getAllHpbNodes() view returns(address[], bytes32[], bytes32[], bytes32[])
func (_Proxy *ProxyCallerSession) GetAllHpbNodes() ([]common.Address, [][32]byte, [][32]byte, [][32]byte, error) {
	return _Proxy.Contract.GetAllHpbNodes(&_Proxy.CallOpts)
}

// GetOwner is a free data retrieval call binding the contract method 0x893d20e8.
//
// Solidity: function getOwner() view returns(address)
func (_Proxy *ProxyCaller) GetOwner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Proxy.contract.Call(opts, &out, "getOwner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetOwner is a free data retrieval call binding the contract method 0x893d20e8.
//
// Solidity: function getOwner() view returns(address)
func (_Proxy *ProxySession) GetOwner() (common.Address, error) {
	return _Proxy.Contract.GetOwner(&_Proxy.CallOpts)
}

// GetOwner is a free data retrieval call binding the contract method 0x893d20e8.
//
// Solidity: function getOwner() view returns(address)
func (_Proxy *ProxyCallerSession) GetOwner() (common.Address, error) {
	return _Proxy.Contract.GetOwner(&_Proxy.CallOpts)
}

// Getcontract is a free data retrieval call binding the contract method 0xe8940888.
//
// Solidity: function getcontract() view returns(address, address, address)
func (_Proxy *ProxyCaller) Getcontract(opts *bind.CallOpts) (common.Address, common.Address, common.Address, error) {
	var out []interface{}
	err := _Proxy.contract.Call(opts, &out, "getcontract")

	if err != nil {
		return *new(common.Address), *new(common.Address), *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	out1 := *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	out2 := *abi.ConvertType(out[2], new(common.Address)).(*common.Address)

	return out0, out1, out2, err

}

// Getcontract is a free data retrieval call binding the contract method 0xe8940888.
//
// Solidity: function getcontract() view returns(address, address, address)
func (_Proxy *ProxySession) Getcontract() (common.Address, common.Address, common.Address, error) {
	return _Proxy.Contract.Getcontract(&_Proxy.CallOpts)
}

// Getcontract is a free data retrieval call binding the contract method 0xe8940888.
//
// Solidity: function getcontract() view returns(address, address, address)
func (_Proxy *ProxyCallerSession) Getcontract() (common.Address, common.Address, common.Address, error) {
	return _Proxy.Contract.Getcontract(&_Proxy.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Proxy *ProxyCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Proxy.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Proxy *ProxySession) Owner() (common.Address, error) {
	return _Proxy.Contract.Owner(&_Proxy.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Proxy *ProxyCallerSession) Owner() (common.Address, error) {
	return _Proxy.Contract.Owner(&_Proxy.CallOpts)
}

// AddAdmin is a paid mutator transaction binding the contract method 0x70480275.
//
// Solidity: function addAdmin(address addr) returns(bool)
func (_Proxy *ProxyTransactor) AddAdmin(opts *bind.TransactOpts, addr common.Address) (*types.Transaction, error) {
	return _Proxy.contract.Transact(opts, "addAdmin", addr)
}

// AddAdmin is a paid mutator transaction binding the contract method 0x70480275.
//
// Solidity: function addAdmin(address addr) returns(bool)
func (_Proxy *ProxySession) AddAdmin(addr common.Address) (*types.Transaction, error) {
	return _Proxy.Contract.AddAdmin(&_Proxy.TransactOpts, addr)
}

// AddAdmin is a paid mutator transaction binding the contract method 0x70480275.
//
// Solidity: function addAdmin(address addr) returns(bool)
func (_Proxy *ProxyTransactorSession) AddAdmin(addr common.Address) (*types.Transaction, error) {
	return _Proxy.Contract.AddAdmin(&_Proxy.TransactOpts, addr)
}

// DeleteAdmin is a paid mutator transaction binding the contract method 0x27e1f7df.
//
// Solidity: function deleteAdmin(address addr) returns(bool)
func (_Proxy *ProxyTransactor) DeleteAdmin(opts *bind.TransactOpts, addr common.Address) (*types.Transaction, error) {
	return _Proxy.contract.Transact(opts, "deleteAdmin", addr)
}

// DeleteAdmin is a paid mutator transaction binding the contract method 0x27e1f7df.
//
// Solidity: function deleteAdmin(address addr) returns(bool)
func (_Proxy *ProxySession) DeleteAdmin(addr common.Address) (*types.Transaction, error) {
	return _Proxy.Contract.DeleteAdmin(&_Proxy.TransactOpts, addr)
}

// DeleteAdmin is a paid mutator transaction binding the contract method 0x27e1f7df.
//
// Solidity: function deleteAdmin(address addr) returns(bool)
func (_Proxy *ProxyTransactorSession) DeleteAdmin(addr common.Address) (*types.Transaction, error) {
	return _Proxy.Contract.DeleteAdmin(&_Proxy.TransactOpts, addr)
}

// Setlockcontract is a paid mutator transaction binding the contract method 0x758e8a7d.
//
// Solidity: function setlockcontract(address lockaddr) returns()
func (_Proxy *ProxyTransactor) Setlockcontract(opts *bind.TransactOpts, lockaddr common.Address) (*types.Transaction, error) {
	return _Proxy.contract.Transact(opts, "setlockcontract", lockaddr)
}

// Setlockcontract is a paid mutator transaction binding the contract method 0x758e8a7d.
//
// Solidity: function setlockcontract(address lockaddr) returns()
func (_Proxy *ProxySession) Setlockcontract(lockaddr common.Address) (*types.Transaction, error) {
	return _Proxy.Contract.Setlockcontract(&_Proxy.TransactOpts, lockaddr)
}

// Setlockcontract is a paid mutator transaction binding the contract method 0x758e8a7d.
//
// Solidity: function setlockcontract(address lockaddr) returns()
func (_Proxy *ProxyTransactorSession) Setlockcontract(lockaddr common.Address) (*types.Transaction, error) {
	return _Proxy.Contract.Setlockcontract(&_Proxy.TransactOpts, lockaddr)
}

// Setnodecontract is a paid mutator transaction binding the contract method 0xbad6e0af.
//
// Solidity: function setnodecontract(address nodeaddr) returns()
func (_Proxy *ProxyTransactor) Setnodecontract(opts *bind.TransactOpts, nodeaddr common.Address) (*types.Transaction, error) {
	return _Proxy.contract.Transact(opts, "setnodecontract", nodeaddr)
}

// Setnodecontract is a paid mutator transaction binding the contract method 0xbad6e0af.
//
// Solidity: function setnodecontract(address nodeaddr) returns()
func (_Proxy *ProxySession) Setnodecontract(nodeaddr common.Address) (*types.Transaction, error) {
	return _Proxy.Contract.Setnodecontract(&_Proxy.TransactOpts, nodeaddr)
}

// Setnodecontract is a paid mutator transaction binding the contract method 0xbad6e0af.
//
// Solidity: function setnodecontract(address nodeaddr) returns()
func (_Proxy *ProxyTransactorSession) Setnodecontract(nodeaddr common.Address) (*types.Transaction, error) {
	return _Proxy.Contract.Setnodecontract(&_Proxy.TransactOpts, nodeaddr)
}

// Setvotecontract is a paid mutator transaction binding the contract method 0x99b8a92f.
//
// Solidity: function setvotecontract(address voteaddr) returns()
func (_Proxy *ProxyTransactor) Setvotecontract(opts *bind.TransactOpts, voteaddr common.Address) (*types.Transaction, error) {
	return _Proxy.contract.Transact(opts, "setvotecontract", voteaddr)
}

// Setvotecontract is a paid mutator transaction binding the contract method 0x99b8a92f.
//
// Solidity: function setvotecontract(address voteaddr) returns()
func (_Proxy *ProxySession) Setvotecontract(voteaddr common.Address) (*types.Transaction, error) {
	return _Proxy.Contract.Setvotecontract(&_Proxy.TransactOpts, voteaddr)
}

// Setvotecontract is a paid mutator transaction binding the contract method 0x99b8a92f.
//
// Solidity: function setvotecontract(address voteaddr) returns()
func (_Proxy *ProxyTransactorSession) Setvotecontract(voteaddr common.Address) (*types.Transaction, error) {
	return _Proxy.Contract.Setvotecontract(&_Proxy.TransactOpts, voteaddr)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns(bool)
func (_Proxy *ProxyTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Proxy.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns(bool)
func (_Proxy *ProxySession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Proxy.Contract.TransferOwnership(&_Proxy.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns(bool)
func (_Proxy *ProxyTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Proxy.Contract.TransferOwnership(&_Proxy.TransactOpts, newOwner)
}
