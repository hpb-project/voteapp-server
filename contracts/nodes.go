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

// NodesABI is the input ABI used to generate the binding from.
const NodesABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"boeaddr\",\"type\":\"address\"}],\"name\":\"setHolderBoeAddr\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"holderAddr\",\"type\":\"address\"}],\"name\":\"setHolderAddr\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"isLockNode\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"nodeAddr\",\"type\":\"address\"}],\"name\":\"stake\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"deleteAdmin\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getAdmins\",\"outputs\":[{\"name\":\"\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"kill\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getAllBoesAddrs\",\"outputs\":[{\"name\":\"\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"nodeAddr\",\"type\":\"address\"}],\"name\":\"withdraw\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"coinbase\",\"type\":\"address\"}],\"name\":\"deleteBoeNode\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"addAdmin\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getAlllockNode\",\"outputs\":[{\"name\":\"\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"isBoeNode\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"coinbase\",\"type\":\"address\"},{\"name\":\"cid1\",\"type\":\"bytes32\"},{\"name\":\"cid2\",\"type\":\"bytes32\"},{\"name\":\"hid\",\"type\":\"bytes32\"}],\"name\":\"addBoeNode\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getOwner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"coinbases\",\"type\":\"address[]\"},{\"name\":\"cid1s\",\"type\":\"bytes32[]\"},{\"name\":\"cid2s\",\"type\":\"bytes32[]\"},{\"name\":\"hids\",\"type\":\"bytes32[]\"}],\"name\":\"updateBoeNodeBatch\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"fetchAllHolderAddrs\",\"outputs\":[{\"name\":\"\",\"type\":\"address[]\"},{\"name\":\"\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"setLockContract\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"coinbases\",\"type\":\"address[]\"}],\"name\":\"deleteBoeNodeBatch\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getLockContract\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"boeaddr\",\"type\":\"address\"}],\"name\":\"getHolderAddr\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getAllBoes\",\"outputs\":[{\"name\":\"\",\"type\":\"address[]\"},{\"name\":\"\",\"type\":\"bytes32[]\"},{\"name\":\"\",\"type\":\"bytes32[]\"},{\"name\":\"\",\"type\":\"bytes32[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"adminMap\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getAllHpbNodes\",\"outputs\":[{\"name\":\"\",\"type\":\"address[]\"},{\"name\":\"\",\"type\":\"bytes32[]\"},{\"name\":\"\",\"type\":\"bytes32[]\"},{\"name\":\"\",\"type\":\"bytes32[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"coinbase\",\"type\":\"address\"},{\"name\":\"cid1\",\"type\":\"bytes32\"},{\"name\":\"cid2\",\"type\":\"bytes32\"},{\"name\":\"hid\",\"type\":\"bytes32\"}],\"name\":\"updateBoeNode\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"coinbases\",\"type\":\"address[]\"},{\"name\":\"cid1s\",\"type\":\"bytes32[]\"},{\"name\":\"cid2s\",\"type\":\"bytes32[]\"},{\"name\":\"hids\",\"type\":\"bytes32[]\"}],\"name\":\"addBoeNodeBatch\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"constructor\"},{\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"ReceivedHpb\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"coinbase\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"cid1\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"cid2\",\"type\":\"bytes32\"},{\"indexed\":true,\"name\":\"hid\",\"type\":\"bytes32\"}],\"name\":\"AddBoeNode\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"coinbase\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"cid1\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"cid2\",\"type\":\"bytes32\"},{\"indexed\":true,\"name\":\"hid\",\"type\":\"bytes32\"}],\"name\":\"UpdateBoeNode\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"coinbase\",\"type\":\"address\"}],\"name\":\"DeleteBoeNode\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"coinBase\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"holderAddr\",\"type\":\"address\"}],\"name\":\"SetHolderAddr\",\"type\":\"event\"}]"

// Nodes is an auto generated Go binding around an Ethereum contract.
type Nodes struct {
	NodesCaller     // Read-only binding to the contract
	NodesTransactor // Write-only binding to the contract
	NodesFilterer   // Log filterer for contract events
}

// NodesCaller is an auto generated read-only Go binding around an Ethereum contract.
type NodesCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NodesTransactor is an auto generated write-only Go binding around an Ethereum contract.
type NodesTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NodesFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type NodesFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NodesSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type NodesSession struct {
	Contract     *Nodes            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// NodesCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type NodesCallerSession struct {
	Contract *NodesCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// NodesTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type NodesTransactorSession struct {
	Contract     *NodesTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// NodesRaw is an auto generated low-level Go binding around an Ethereum contract.
type NodesRaw struct {
	Contract *Nodes // Generic contract binding to access the raw methods on
}

// NodesCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type NodesCallerRaw struct {
	Contract *NodesCaller // Generic read-only contract binding to access the raw methods on
}

// NodesTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type NodesTransactorRaw struct {
	Contract *NodesTransactor // Generic write-only contract binding to access the raw methods on
}

// NewNodes creates a new instance of Nodes, bound to a specific deployed contract.
func NewNodes(address common.Address, backend bind.ContractBackend) (*Nodes, error) {
	contract, err := bindNodes(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Nodes{NodesCaller: NodesCaller{contract: contract}, NodesTransactor: NodesTransactor{contract: contract}, NodesFilterer: NodesFilterer{contract: contract}}, nil
}

// NewNodesCaller creates a new read-only instance of Nodes, bound to a specific deployed contract.
func NewNodesCaller(address common.Address, caller bind.ContractCaller) (*NodesCaller, error) {
	contract, err := bindNodes(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &NodesCaller{contract: contract}, nil
}

// NewNodesTransactor creates a new write-only instance of Nodes, bound to a specific deployed contract.
func NewNodesTransactor(address common.Address, transactor bind.ContractTransactor) (*NodesTransactor, error) {
	contract, err := bindNodes(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &NodesTransactor{contract: contract}, nil
}

// NewNodesFilterer creates a new log filterer instance of Nodes, bound to a specific deployed contract.
func NewNodesFilterer(address common.Address, filterer bind.ContractFilterer) (*NodesFilterer, error) {
	contract, err := bindNodes(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &NodesFilterer{contract: contract}, nil
}

// bindNodes binds a generic wrapper to an already deployed contract.
func bindNodes(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(NodesABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Nodes *NodesRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Nodes.Contract.NodesCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Nodes *NodesRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Nodes.Contract.NodesTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Nodes *NodesRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Nodes.Contract.NodesTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Nodes *NodesCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Nodes.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Nodes *NodesTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Nodes.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Nodes *NodesTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Nodes.Contract.contract.Transact(opts, method, params...)
}

// AdminMap is a free data retrieval call binding the contract method 0xdbbc830b.
//
// Solidity: function adminMap(address ) view returns(address)
func (_Nodes *NodesCaller) AdminMap(opts *bind.CallOpts, arg0 common.Address) (common.Address, error) {
	var out []interface{}
	err := _Nodes.contract.Call(opts, &out, "adminMap", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AdminMap is a free data retrieval call binding the contract method 0xdbbc830b.
//
// Solidity: function adminMap(address ) view returns(address)
func (_Nodes *NodesSession) AdminMap(arg0 common.Address) (common.Address, error) {
	return _Nodes.Contract.AdminMap(&_Nodes.CallOpts, arg0)
}

// AdminMap is a free data retrieval call binding the contract method 0xdbbc830b.
//
// Solidity: function adminMap(address ) view returns(address)
func (_Nodes *NodesCallerSession) AdminMap(arg0 common.Address) (common.Address, error) {
	return _Nodes.Contract.AdminMap(&_Nodes.CallOpts, arg0)
}

// FetchAllHolderAddrs is a free data retrieval call binding the contract method 0xab2d2c0b.
//
// Solidity: function fetchAllHolderAddrs() view returns(address[], address[])
func (_Nodes *NodesCaller) FetchAllHolderAddrs(opts *bind.CallOpts) ([]common.Address, []common.Address, error) {
	var out []interface{}
	err := _Nodes.contract.Call(opts, &out, "fetchAllHolderAddrs")

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
func (_Nodes *NodesSession) FetchAllHolderAddrs() ([]common.Address, []common.Address, error) {
	return _Nodes.Contract.FetchAllHolderAddrs(&_Nodes.CallOpts)
}

// FetchAllHolderAddrs is a free data retrieval call binding the contract method 0xab2d2c0b.
//
// Solidity: function fetchAllHolderAddrs() view returns(address[], address[])
func (_Nodes *NodesCallerSession) FetchAllHolderAddrs() ([]common.Address, []common.Address, error) {
	return _Nodes.Contract.FetchAllHolderAddrs(&_Nodes.CallOpts)
}

// GetAdmins is a free data retrieval call binding the contract method 0x31ae450b.
//
// Solidity: function getAdmins() view returns(address[])
func (_Nodes *NodesCaller) GetAdmins(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _Nodes.contract.Call(opts, &out, "getAdmins")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetAdmins is a free data retrieval call binding the contract method 0x31ae450b.
//
// Solidity: function getAdmins() view returns(address[])
func (_Nodes *NodesSession) GetAdmins() ([]common.Address, error) {
	return _Nodes.Contract.GetAdmins(&_Nodes.CallOpts)
}

// GetAdmins is a free data retrieval call binding the contract method 0x31ae450b.
//
// Solidity: function getAdmins() view returns(address[])
func (_Nodes *NodesCallerSession) GetAdmins() ([]common.Address, error) {
	return _Nodes.Contract.GetAdmins(&_Nodes.CallOpts)
}

// GetAllBoes is a free data retrieval call binding the contract method 0xd68e9cf4.
//
// Solidity: function getAllBoes() view returns(address[], bytes32[], bytes32[], bytes32[])
func (_Nodes *NodesCaller) GetAllBoes(opts *bind.CallOpts) ([]common.Address, [][32]byte, [][32]byte, [][32]byte, error) {
	var out []interface{}
	err := _Nodes.contract.Call(opts, &out, "getAllBoes")

	if err != nil {
		return *new([]common.Address), *new([][32]byte), *new([][32]byte), *new([][32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)
	out1 := *abi.ConvertType(out[1], new([][32]byte)).(*[][32]byte)
	out2 := *abi.ConvertType(out[2], new([][32]byte)).(*[][32]byte)
	out3 := *abi.ConvertType(out[3], new([][32]byte)).(*[][32]byte)

	return out0, out1, out2, out3, err

}

// GetAllBoes is a free data retrieval call binding the contract method 0xd68e9cf4.
//
// Solidity: function getAllBoes() view returns(address[], bytes32[], bytes32[], bytes32[])
func (_Nodes *NodesSession) GetAllBoes() ([]common.Address, [][32]byte, [][32]byte, [][32]byte, error) {
	return _Nodes.Contract.GetAllBoes(&_Nodes.CallOpts)
}

// GetAllBoes is a free data retrieval call binding the contract method 0xd68e9cf4.
//
// Solidity: function getAllBoes() view returns(address[], bytes32[], bytes32[], bytes32[])
func (_Nodes *NodesCallerSession) GetAllBoes() ([]common.Address, [][32]byte, [][32]byte, [][32]byte, error) {
	return _Nodes.Contract.GetAllBoes(&_Nodes.CallOpts)
}

// GetAllBoesAddrs is a free data retrieval call binding the contract method 0x4d92591b.
//
// Solidity: function getAllBoesAddrs() view returns(address[])
func (_Nodes *NodesCaller) GetAllBoesAddrs(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _Nodes.contract.Call(opts, &out, "getAllBoesAddrs")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetAllBoesAddrs is a free data retrieval call binding the contract method 0x4d92591b.
//
// Solidity: function getAllBoesAddrs() view returns(address[])
func (_Nodes *NodesSession) GetAllBoesAddrs() ([]common.Address, error) {
	return _Nodes.Contract.GetAllBoesAddrs(&_Nodes.CallOpts)
}

// GetAllBoesAddrs is a free data retrieval call binding the contract method 0x4d92591b.
//
// Solidity: function getAllBoesAddrs() view returns(address[])
func (_Nodes *NodesCallerSession) GetAllBoesAddrs() ([]common.Address, error) {
	return _Nodes.Contract.GetAllBoesAddrs(&_Nodes.CallOpts)
}

// GetAllHpbNodes is a free data retrieval call binding the contract method 0xdbe01790.
//
// Solidity: function getAllHpbNodes() view returns(address[], bytes32[], bytes32[], bytes32[])
func (_Nodes *NodesCaller) GetAllHpbNodes(opts *bind.CallOpts) ([]common.Address, [][32]byte, [][32]byte, [][32]byte, error) {
	var out []interface{}
	err := _Nodes.contract.Call(opts, &out, "getAllHpbNodes")

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
func (_Nodes *NodesSession) GetAllHpbNodes() ([]common.Address, [][32]byte, [][32]byte, [][32]byte, error) {
	return _Nodes.Contract.GetAllHpbNodes(&_Nodes.CallOpts)
}

// GetAllHpbNodes is a free data retrieval call binding the contract method 0xdbe01790.
//
// Solidity: function getAllHpbNodes() view returns(address[], bytes32[], bytes32[], bytes32[])
func (_Nodes *NodesCallerSession) GetAllHpbNodes() ([]common.Address, [][32]byte, [][32]byte, [][32]byte, error) {
	return _Nodes.Contract.GetAllHpbNodes(&_Nodes.CallOpts)
}

// GetAlllockNode is a free data retrieval call binding the contract method 0x794cef01.
//
// Solidity: function getAlllockNode() view returns(address[])
func (_Nodes *NodesCaller) GetAlllockNode(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _Nodes.contract.Call(opts, &out, "getAlllockNode")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetAlllockNode is a free data retrieval call binding the contract method 0x794cef01.
//
// Solidity: function getAlllockNode() view returns(address[])
func (_Nodes *NodesSession) GetAlllockNode() ([]common.Address, error) {
	return _Nodes.Contract.GetAlllockNode(&_Nodes.CallOpts)
}

// GetAlllockNode is a free data retrieval call binding the contract method 0x794cef01.
//
// Solidity: function getAlllockNode() view returns(address[])
func (_Nodes *NodesCallerSession) GetAlllockNode() ([]common.Address, error) {
	return _Nodes.Contract.GetAlllockNode(&_Nodes.CallOpts)
}

// GetHolderAddr is a free data retrieval call binding the contract method 0xd3c932c2.
//
// Solidity: function getHolderAddr(address boeaddr) view returns(address)
func (_Nodes *NodesCaller) GetHolderAddr(opts *bind.CallOpts, boeaddr common.Address) (common.Address, error) {
	var out []interface{}
	err := _Nodes.contract.Call(opts, &out, "getHolderAddr", boeaddr)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetHolderAddr is a free data retrieval call binding the contract method 0xd3c932c2.
//
// Solidity: function getHolderAddr(address boeaddr) view returns(address)
func (_Nodes *NodesSession) GetHolderAddr(boeaddr common.Address) (common.Address, error) {
	return _Nodes.Contract.GetHolderAddr(&_Nodes.CallOpts, boeaddr)
}

// GetHolderAddr is a free data retrieval call binding the contract method 0xd3c932c2.
//
// Solidity: function getHolderAddr(address boeaddr) view returns(address)
func (_Nodes *NodesCallerSession) GetHolderAddr(boeaddr common.Address) (common.Address, error) {
	return _Nodes.Contract.GetHolderAddr(&_Nodes.CallOpts, boeaddr)
}

// GetLockContract is a free data retrieval call binding the contract method 0xd0c05a92.
//
// Solidity: function getLockContract() view returns(address)
func (_Nodes *NodesCaller) GetLockContract(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Nodes.contract.Call(opts, &out, "getLockContract")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetLockContract is a free data retrieval call binding the contract method 0xd0c05a92.
//
// Solidity: function getLockContract() view returns(address)
func (_Nodes *NodesSession) GetLockContract() (common.Address, error) {
	return _Nodes.Contract.GetLockContract(&_Nodes.CallOpts)
}

// GetLockContract is a free data retrieval call binding the contract method 0xd0c05a92.
//
// Solidity: function getLockContract() view returns(address)
func (_Nodes *NodesCallerSession) GetLockContract() (common.Address, error) {
	return _Nodes.Contract.GetLockContract(&_Nodes.CallOpts)
}

// GetOwner is a free data retrieval call binding the contract method 0x893d20e8.
//
// Solidity: function getOwner() view returns(address)
func (_Nodes *NodesCaller) GetOwner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Nodes.contract.Call(opts, &out, "getOwner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetOwner is a free data retrieval call binding the contract method 0x893d20e8.
//
// Solidity: function getOwner() view returns(address)
func (_Nodes *NodesSession) GetOwner() (common.Address, error) {
	return _Nodes.Contract.GetOwner(&_Nodes.CallOpts)
}

// GetOwner is a free data retrieval call binding the contract method 0x893d20e8.
//
// Solidity: function getOwner() view returns(address)
func (_Nodes *NodesCallerSession) GetOwner() (common.Address, error) {
	return _Nodes.Contract.GetOwner(&_Nodes.CallOpts)
}

// IsBoeNode is a free data retrieval call binding the contract method 0x7e533751.
//
// Solidity: function isBoeNode(address addr) view returns(bool)
func (_Nodes *NodesCaller) IsBoeNode(opts *bind.CallOpts, addr common.Address) (bool, error) {
	var out []interface{}
	err := _Nodes.contract.Call(opts, &out, "isBoeNode", addr)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsBoeNode is a free data retrieval call binding the contract method 0x7e533751.
//
// Solidity: function isBoeNode(address addr) view returns(bool)
func (_Nodes *NodesSession) IsBoeNode(addr common.Address) (bool, error) {
	return _Nodes.Contract.IsBoeNode(&_Nodes.CallOpts, addr)
}

// IsBoeNode is a free data retrieval call binding the contract method 0x7e533751.
//
// Solidity: function isBoeNode(address addr) view returns(bool)
func (_Nodes *NodesCallerSession) IsBoeNode(addr common.Address) (bool, error) {
	return _Nodes.Contract.IsBoeNode(&_Nodes.CallOpts, addr)
}

// IsLockNode is a free data retrieval call binding the contract method 0x228101b6.
//
// Solidity: function isLockNode(address addr) view returns(bool)
func (_Nodes *NodesCaller) IsLockNode(opts *bind.CallOpts, addr common.Address) (bool, error) {
	var out []interface{}
	err := _Nodes.contract.Call(opts, &out, "isLockNode", addr)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsLockNode is a free data retrieval call binding the contract method 0x228101b6.
//
// Solidity: function isLockNode(address addr) view returns(bool)
func (_Nodes *NodesSession) IsLockNode(addr common.Address) (bool, error) {
	return _Nodes.Contract.IsLockNode(&_Nodes.CallOpts, addr)
}

// IsLockNode is a free data retrieval call binding the contract method 0x228101b6.
//
// Solidity: function isLockNode(address addr) view returns(bool)
func (_Nodes *NodesCallerSession) IsLockNode(addr common.Address) (bool, error) {
	return _Nodes.Contract.IsLockNode(&_Nodes.CallOpts, addr)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Nodes *NodesCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Nodes.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Nodes *NodesSession) Name() (string, error) {
	return _Nodes.Contract.Name(&_Nodes.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Nodes *NodesCallerSession) Name() (string, error) {
	return _Nodes.Contract.Name(&_Nodes.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Nodes *NodesCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Nodes.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Nodes *NodesSession) Owner() (common.Address, error) {
	return _Nodes.Contract.Owner(&_Nodes.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Nodes *NodesCallerSession) Owner() (common.Address, error) {
	return _Nodes.Contract.Owner(&_Nodes.CallOpts)
}

// AddAdmin is a paid mutator transaction binding the contract method 0x70480275.
//
// Solidity: function addAdmin(address addr) returns(bool)
func (_Nodes *NodesTransactor) AddAdmin(opts *bind.TransactOpts, addr common.Address) (*types.Transaction, error) {
	return _Nodes.contract.Transact(opts, "addAdmin", addr)
}

// AddAdmin is a paid mutator transaction binding the contract method 0x70480275.
//
// Solidity: function addAdmin(address addr) returns(bool)
func (_Nodes *NodesSession) AddAdmin(addr common.Address) (*types.Transaction, error) {
	return _Nodes.Contract.AddAdmin(&_Nodes.TransactOpts, addr)
}

// AddAdmin is a paid mutator transaction binding the contract method 0x70480275.
//
// Solidity: function addAdmin(address addr) returns(bool)
func (_Nodes *NodesTransactorSession) AddAdmin(addr common.Address) (*types.Transaction, error) {
	return _Nodes.Contract.AddAdmin(&_Nodes.TransactOpts, addr)
}

// AddBoeNode is a paid mutator transaction binding the contract method 0x84547875.
//
// Solidity: function addBoeNode(address coinbase, bytes32 cid1, bytes32 cid2, bytes32 hid) returns(bool)
func (_Nodes *NodesTransactor) AddBoeNode(opts *bind.TransactOpts, coinbase common.Address, cid1 [32]byte, cid2 [32]byte, hid [32]byte) (*types.Transaction, error) {
	return _Nodes.contract.Transact(opts, "addBoeNode", coinbase, cid1, cid2, hid)
}

// AddBoeNode is a paid mutator transaction binding the contract method 0x84547875.
//
// Solidity: function addBoeNode(address coinbase, bytes32 cid1, bytes32 cid2, bytes32 hid) returns(bool)
func (_Nodes *NodesSession) AddBoeNode(coinbase common.Address, cid1 [32]byte, cid2 [32]byte, hid [32]byte) (*types.Transaction, error) {
	return _Nodes.Contract.AddBoeNode(&_Nodes.TransactOpts, coinbase, cid1, cid2, hid)
}

// AddBoeNode is a paid mutator transaction binding the contract method 0x84547875.
//
// Solidity: function addBoeNode(address coinbase, bytes32 cid1, bytes32 cid2, bytes32 hid) returns(bool)
func (_Nodes *NodesTransactorSession) AddBoeNode(coinbase common.Address, cid1 [32]byte, cid2 [32]byte, hid [32]byte) (*types.Transaction, error) {
	return _Nodes.Contract.AddBoeNode(&_Nodes.TransactOpts, coinbase, cid1, cid2, hid)
}

// AddBoeNodeBatch is a paid mutator transaction binding the contract method 0xf65b5378.
//
// Solidity: function addBoeNodeBatch(address[] coinbases, bytes32[] cid1s, bytes32[] cid2s, bytes32[] hids) returns(bool)
func (_Nodes *NodesTransactor) AddBoeNodeBatch(opts *bind.TransactOpts, coinbases []common.Address, cid1s [][32]byte, cid2s [][32]byte, hids [][32]byte) (*types.Transaction, error) {
	return _Nodes.contract.Transact(opts, "addBoeNodeBatch", coinbases, cid1s, cid2s, hids)
}

// AddBoeNodeBatch is a paid mutator transaction binding the contract method 0xf65b5378.
//
// Solidity: function addBoeNodeBatch(address[] coinbases, bytes32[] cid1s, bytes32[] cid2s, bytes32[] hids) returns(bool)
func (_Nodes *NodesSession) AddBoeNodeBatch(coinbases []common.Address, cid1s [][32]byte, cid2s [][32]byte, hids [][32]byte) (*types.Transaction, error) {
	return _Nodes.Contract.AddBoeNodeBatch(&_Nodes.TransactOpts, coinbases, cid1s, cid2s, hids)
}

// AddBoeNodeBatch is a paid mutator transaction binding the contract method 0xf65b5378.
//
// Solidity: function addBoeNodeBatch(address[] coinbases, bytes32[] cid1s, bytes32[] cid2s, bytes32[] hids) returns(bool)
func (_Nodes *NodesTransactorSession) AddBoeNodeBatch(coinbases []common.Address, cid1s [][32]byte, cid2s [][32]byte, hids [][32]byte) (*types.Transaction, error) {
	return _Nodes.Contract.AddBoeNodeBatch(&_Nodes.TransactOpts, coinbases, cid1s, cid2s, hids)
}

// DeleteAdmin is a paid mutator transaction binding the contract method 0x27e1f7df.
//
// Solidity: function deleteAdmin(address addr) returns(bool)
func (_Nodes *NodesTransactor) DeleteAdmin(opts *bind.TransactOpts, addr common.Address) (*types.Transaction, error) {
	return _Nodes.contract.Transact(opts, "deleteAdmin", addr)
}

// DeleteAdmin is a paid mutator transaction binding the contract method 0x27e1f7df.
//
// Solidity: function deleteAdmin(address addr) returns(bool)
func (_Nodes *NodesSession) DeleteAdmin(addr common.Address) (*types.Transaction, error) {
	return _Nodes.Contract.DeleteAdmin(&_Nodes.TransactOpts, addr)
}

// DeleteAdmin is a paid mutator transaction binding the contract method 0x27e1f7df.
//
// Solidity: function deleteAdmin(address addr) returns(bool)
func (_Nodes *NodesTransactorSession) DeleteAdmin(addr common.Address) (*types.Transaction, error) {
	return _Nodes.Contract.DeleteAdmin(&_Nodes.TransactOpts, addr)
}

// DeleteBoeNode is a paid mutator transaction binding the contract method 0x595e2d90.
//
// Solidity: function deleteBoeNode(address coinbase) returns(bool)
func (_Nodes *NodesTransactor) DeleteBoeNode(opts *bind.TransactOpts, coinbase common.Address) (*types.Transaction, error) {
	return _Nodes.contract.Transact(opts, "deleteBoeNode", coinbase)
}

// DeleteBoeNode is a paid mutator transaction binding the contract method 0x595e2d90.
//
// Solidity: function deleteBoeNode(address coinbase) returns(bool)
func (_Nodes *NodesSession) DeleteBoeNode(coinbase common.Address) (*types.Transaction, error) {
	return _Nodes.Contract.DeleteBoeNode(&_Nodes.TransactOpts, coinbase)
}

// DeleteBoeNode is a paid mutator transaction binding the contract method 0x595e2d90.
//
// Solidity: function deleteBoeNode(address coinbase) returns(bool)
func (_Nodes *NodesTransactorSession) DeleteBoeNode(coinbase common.Address) (*types.Transaction, error) {
	return _Nodes.Contract.DeleteBoeNode(&_Nodes.TransactOpts, coinbase)
}

// DeleteBoeNodeBatch is a paid mutator transaction binding the contract method 0xced99931.
//
// Solidity: function deleteBoeNodeBatch(address[] coinbases) returns(bool)
func (_Nodes *NodesTransactor) DeleteBoeNodeBatch(opts *bind.TransactOpts, coinbases []common.Address) (*types.Transaction, error) {
	return _Nodes.contract.Transact(opts, "deleteBoeNodeBatch", coinbases)
}

// DeleteBoeNodeBatch is a paid mutator transaction binding the contract method 0xced99931.
//
// Solidity: function deleteBoeNodeBatch(address[] coinbases) returns(bool)
func (_Nodes *NodesSession) DeleteBoeNodeBatch(coinbases []common.Address) (*types.Transaction, error) {
	return _Nodes.Contract.DeleteBoeNodeBatch(&_Nodes.TransactOpts, coinbases)
}

// DeleteBoeNodeBatch is a paid mutator transaction binding the contract method 0xced99931.
//
// Solidity: function deleteBoeNodeBatch(address[] coinbases) returns(bool)
func (_Nodes *NodesTransactorSession) DeleteBoeNodeBatch(coinbases []common.Address) (*types.Transaction, error) {
	return _Nodes.Contract.DeleteBoeNodeBatch(&_Nodes.TransactOpts, coinbases)
}

// Kill is a paid mutator transaction binding the contract method 0x41c0e1b5.
//
// Solidity: function kill() payable returns(bool)
func (_Nodes *NodesTransactor) Kill(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Nodes.contract.Transact(opts, "kill")
}

// Kill is a paid mutator transaction binding the contract method 0x41c0e1b5.
//
// Solidity: function kill() payable returns(bool)
func (_Nodes *NodesSession) Kill() (*types.Transaction, error) {
	return _Nodes.Contract.Kill(&_Nodes.TransactOpts)
}

// Kill is a paid mutator transaction binding the contract method 0x41c0e1b5.
//
// Solidity: function kill() payable returns(bool)
func (_Nodes *NodesTransactorSession) Kill() (*types.Transaction, error) {
	return _Nodes.Contract.Kill(&_Nodes.TransactOpts)
}

// SetHolderAddr is a paid mutator transaction binding the contract method 0x20ad7c7e.
//
// Solidity: function setHolderAddr(address holderAddr) returns()
func (_Nodes *NodesTransactor) SetHolderAddr(opts *bind.TransactOpts, holderAddr common.Address) (*types.Transaction, error) {
	return _Nodes.contract.Transact(opts, "setHolderAddr", holderAddr)
}

// SetHolderAddr is a paid mutator transaction binding the contract method 0x20ad7c7e.
//
// Solidity: function setHolderAddr(address holderAddr) returns()
func (_Nodes *NodesSession) SetHolderAddr(holderAddr common.Address) (*types.Transaction, error) {
	return _Nodes.Contract.SetHolderAddr(&_Nodes.TransactOpts, holderAddr)
}

// SetHolderAddr is a paid mutator transaction binding the contract method 0x20ad7c7e.
//
// Solidity: function setHolderAddr(address holderAddr) returns()
func (_Nodes *NodesTransactorSession) SetHolderAddr(holderAddr common.Address) (*types.Transaction, error) {
	return _Nodes.Contract.SetHolderAddr(&_Nodes.TransactOpts, holderAddr)
}

// SetHolderBoeAddr is a paid mutator transaction binding the contract method 0x05230630.
//
// Solidity: function setHolderBoeAddr(address boeaddr) returns()
func (_Nodes *NodesTransactor) SetHolderBoeAddr(opts *bind.TransactOpts, boeaddr common.Address) (*types.Transaction, error) {
	return _Nodes.contract.Transact(opts, "setHolderBoeAddr", boeaddr)
}

// SetHolderBoeAddr is a paid mutator transaction binding the contract method 0x05230630.
//
// Solidity: function setHolderBoeAddr(address boeaddr) returns()
func (_Nodes *NodesSession) SetHolderBoeAddr(boeaddr common.Address) (*types.Transaction, error) {
	return _Nodes.Contract.SetHolderBoeAddr(&_Nodes.TransactOpts, boeaddr)
}

// SetHolderBoeAddr is a paid mutator transaction binding the contract method 0x05230630.
//
// Solidity: function setHolderBoeAddr(address boeaddr) returns()
func (_Nodes *NodesTransactorSession) SetHolderBoeAddr(boeaddr common.Address) (*types.Transaction, error) {
	return _Nodes.Contract.SetHolderBoeAddr(&_Nodes.TransactOpts, boeaddr)
}

// SetLockContract is a paid mutator transaction binding the contract method 0xacd487c2.
//
// Solidity: function setLockContract(address addr) returns(bool)
func (_Nodes *NodesTransactor) SetLockContract(opts *bind.TransactOpts, addr common.Address) (*types.Transaction, error) {
	return _Nodes.contract.Transact(opts, "setLockContract", addr)
}

// SetLockContract is a paid mutator transaction binding the contract method 0xacd487c2.
//
// Solidity: function setLockContract(address addr) returns(bool)
func (_Nodes *NodesSession) SetLockContract(addr common.Address) (*types.Transaction, error) {
	return _Nodes.Contract.SetLockContract(&_Nodes.TransactOpts, addr)
}

// SetLockContract is a paid mutator transaction binding the contract method 0xacd487c2.
//
// Solidity: function setLockContract(address addr) returns(bool)
func (_Nodes *NodesTransactorSession) SetLockContract(addr common.Address) (*types.Transaction, error) {
	return _Nodes.Contract.SetLockContract(&_Nodes.TransactOpts, addr)
}

// Stake is a paid mutator transaction binding the contract method 0x26476204.
//
// Solidity: function stake(address nodeAddr) returns()
func (_Nodes *NodesTransactor) Stake(opts *bind.TransactOpts, nodeAddr common.Address) (*types.Transaction, error) {
	return _Nodes.contract.Transact(opts, "stake", nodeAddr)
}

// Stake is a paid mutator transaction binding the contract method 0x26476204.
//
// Solidity: function stake(address nodeAddr) returns()
func (_Nodes *NodesSession) Stake(nodeAddr common.Address) (*types.Transaction, error) {
	return _Nodes.Contract.Stake(&_Nodes.TransactOpts, nodeAddr)
}

// Stake is a paid mutator transaction binding the contract method 0x26476204.
//
// Solidity: function stake(address nodeAddr) returns()
func (_Nodes *NodesTransactorSession) Stake(nodeAddr common.Address) (*types.Transaction, error) {
	return _Nodes.Contract.Stake(&_Nodes.TransactOpts, nodeAddr)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns(bool)
func (_Nodes *NodesTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Nodes.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns(bool)
func (_Nodes *NodesSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Nodes.Contract.TransferOwnership(&_Nodes.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns(bool)
func (_Nodes *NodesTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Nodes.Contract.TransferOwnership(&_Nodes.TransactOpts, newOwner)
}

// UpdateBoeNode is a paid mutator transaction binding the contract method 0xf16f531c.
//
// Solidity: function updateBoeNode(address coinbase, bytes32 cid1, bytes32 cid2, bytes32 hid) returns(bool)
func (_Nodes *NodesTransactor) UpdateBoeNode(opts *bind.TransactOpts, coinbase common.Address, cid1 [32]byte, cid2 [32]byte, hid [32]byte) (*types.Transaction, error) {
	return _Nodes.contract.Transact(opts, "updateBoeNode", coinbase, cid1, cid2, hid)
}

// UpdateBoeNode is a paid mutator transaction binding the contract method 0xf16f531c.
//
// Solidity: function updateBoeNode(address coinbase, bytes32 cid1, bytes32 cid2, bytes32 hid) returns(bool)
func (_Nodes *NodesSession) UpdateBoeNode(coinbase common.Address, cid1 [32]byte, cid2 [32]byte, hid [32]byte) (*types.Transaction, error) {
	return _Nodes.Contract.UpdateBoeNode(&_Nodes.TransactOpts, coinbase, cid1, cid2, hid)
}

// UpdateBoeNode is a paid mutator transaction binding the contract method 0xf16f531c.
//
// Solidity: function updateBoeNode(address coinbase, bytes32 cid1, bytes32 cid2, bytes32 hid) returns(bool)
func (_Nodes *NodesTransactorSession) UpdateBoeNode(coinbase common.Address, cid1 [32]byte, cid2 [32]byte, hid [32]byte) (*types.Transaction, error) {
	return _Nodes.Contract.UpdateBoeNode(&_Nodes.TransactOpts, coinbase, cid1, cid2, hid)
}

// UpdateBoeNodeBatch is a paid mutator transaction binding the contract method 0x8fc532ce.
//
// Solidity: function updateBoeNodeBatch(address[] coinbases, bytes32[] cid1s, bytes32[] cid2s, bytes32[] hids) returns(bool)
func (_Nodes *NodesTransactor) UpdateBoeNodeBatch(opts *bind.TransactOpts, coinbases []common.Address, cid1s [][32]byte, cid2s [][32]byte, hids [][32]byte) (*types.Transaction, error) {
	return _Nodes.contract.Transact(opts, "updateBoeNodeBatch", coinbases, cid1s, cid2s, hids)
}

// UpdateBoeNodeBatch is a paid mutator transaction binding the contract method 0x8fc532ce.
//
// Solidity: function updateBoeNodeBatch(address[] coinbases, bytes32[] cid1s, bytes32[] cid2s, bytes32[] hids) returns(bool)
func (_Nodes *NodesSession) UpdateBoeNodeBatch(coinbases []common.Address, cid1s [][32]byte, cid2s [][32]byte, hids [][32]byte) (*types.Transaction, error) {
	return _Nodes.Contract.UpdateBoeNodeBatch(&_Nodes.TransactOpts, coinbases, cid1s, cid2s, hids)
}

// UpdateBoeNodeBatch is a paid mutator transaction binding the contract method 0x8fc532ce.
//
// Solidity: function updateBoeNodeBatch(address[] coinbases, bytes32[] cid1s, bytes32[] cid2s, bytes32[] hids) returns(bool)
func (_Nodes *NodesTransactorSession) UpdateBoeNodeBatch(coinbases []common.Address, cid1s [][32]byte, cid2s [][32]byte, hids [][32]byte) (*types.Transaction, error) {
	return _Nodes.Contract.UpdateBoeNodeBatch(&_Nodes.TransactOpts, coinbases, cid1s, cid2s, hids)
}

// Withdraw is a paid mutator transaction binding the contract method 0x51cff8d9.
//
// Solidity: function withdraw(address nodeAddr) returns()
func (_Nodes *NodesTransactor) Withdraw(opts *bind.TransactOpts, nodeAddr common.Address) (*types.Transaction, error) {
	return _Nodes.contract.Transact(opts, "withdraw", nodeAddr)
}

// Withdraw is a paid mutator transaction binding the contract method 0x51cff8d9.
//
// Solidity: function withdraw(address nodeAddr) returns()
func (_Nodes *NodesSession) Withdraw(nodeAddr common.Address) (*types.Transaction, error) {
	return _Nodes.Contract.Withdraw(&_Nodes.TransactOpts, nodeAddr)
}

// Withdraw is a paid mutator transaction binding the contract method 0x51cff8d9.
//
// Solidity: function withdraw(address nodeAddr) returns()
func (_Nodes *NodesTransactorSession) Withdraw(nodeAddr common.Address) (*types.Transaction, error) {
	return _Nodes.Contract.Withdraw(&_Nodes.TransactOpts, nodeAddr)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_Nodes *NodesTransactor) Fallback(opts *bind.TransactOpts, calldata []byte) (*types.Transaction, error) {
	return _Nodes.contract.RawTransact(opts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_Nodes *NodesSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _Nodes.Contract.Fallback(&_Nodes.TransactOpts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_Nodes *NodesTransactorSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _Nodes.Contract.Fallback(&_Nodes.TransactOpts, calldata)
}

// NodesAddBoeNodeIterator is returned from FilterAddBoeNode and is used to iterate over the raw logs and unpacked data for AddBoeNode events raised by the Nodes contract.
type NodesAddBoeNodeIterator struct {
	Event *NodesAddBoeNode // Event containing the contract specifics and raw log

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
func (it *NodesAddBoeNodeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NodesAddBoeNode)
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
		it.Event = new(NodesAddBoeNode)
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
func (it *NodesAddBoeNodeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NodesAddBoeNodeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NodesAddBoeNode represents a AddBoeNode event raised by the Nodes contract.
type NodesAddBoeNode struct {
	Coinbase common.Address
	Cid1     [32]byte
	Cid2     [32]byte
	Hid      [32]byte
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterAddBoeNode is a free log retrieval operation binding the contract event 0x1f0e0433d82137325a5487c23f2f9a932e6c911fcd18807cd83facdf84b57f8b.
//
// Solidity: event AddBoeNode(address indexed coinbase, bytes32 cid1, bytes32 cid2, bytes32 indexed hid)
func (_Nodes *NodesFilterer) FilterAddBoeNode(opts *bind.FilterOpts, coinbase []common.Address, hid [][32]byte) (*NodesAddBoeNodeIterator, error) {

	var coinbaseRule []interface{}
	for _, coinbaseItem := range coinbase {
		coinbaseRule = append(coinbaseRule, coinbaseItem)
	}

	var hidRule []interface{}
	for _, hidItem := range hid {
		hidRule = append(hidRule, hidItem)
	}

	logs, sub, err := _Nodes.contract.FilterLogs(opts, "AddBoeNode", coinbaseRule, hidRule)
	if err != nil {
		return nil, err
	}
	return &NodesAddBoeNodeIterator{contract: _Nodes.contract, event: "AddBoeNode", logs: logs, sub: sub}, nil
}

// WatchAddBoeNode is a free log subscription operation binding the contract event 0x1f0e0433d82137325a5487c23f2f9a932e6c911fcd18807cd83facdf84b57f8b.
//
// Solidity: event AddBoeNode(address indexed coinbase, bytes32 cid1, bytes32 cid2, bytes32 indexed hid)
func (_Nodes *NodesFilterer) WatchAddBoeNode(opts *bind.WatchOpts, sink chan<- *NodesAddBoeNode, coinbase []common.Address, hid [][32]byte) (event.Subscription, error) {

	var coinbaseRule []interface{}
	for _, coinbaseItem := range coinbase {
		coinbaseRule = append(coinbaseRule, coinbaseItem)
	}

	var hidRule []interface{}
	for _, hidItem := range hid {
		hidRule = append(hidRule, hidItem)
	}

	logs, sub, err := _Nodes.contract.WatchLogs(opts, "AddBoeNode", coinbaseRule, hidRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NodesAddBoeNode)
				if err := _Nodes.contract.UnpackLog(event, "AddBoeNode", log); err != nil {
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

// ParseAddBoeNode is a log parse operation binding the contract event 0x1f0e0433d82137325a5487c23f2f9a932e6c911fcd18807cd83facdf84b57f8b.
//
// Solidity: event AddBoeNode(address indexed coinbase, bytes32 cid1, bytes32 cid2, bytes32 indexed hid)
func (_Nodes *NodesFilterer) ParseAddBoeNode(log types.Log) (*NodesAddBoeNode, error) {
	event := new(NodesAddBoeNode)
	if err := _Nodes.contract.UnpackLog(event, "AddBoeNode", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NodesDeleteBoeNodeIterator is returned from FilterDeleteBoeNode and is used to iterate over the raw logs and unpacked data for DeleteBoeNode events raised by the Nodes contract.
type NodesDeleteBoeNodeIterator struct {
	Event *NodesDeleteBoeNode // Event containing the contract specifics and raw log

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
func (it *NodesDeleteBoeNodeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NodesDeleteBoeNode)
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
		it.Event = new(NodesDeleteBoeNode)
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
func (it *NodesDeleteBoeNodeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NodesDeleteBoeNodeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NodesDeleteBoeNode represents a DeleteBoeNode event raised by the Nodes contract.
type NodesDeleteBoeNode struct {
	Coinbase common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterDeleteBoeNode is a free log retrieval operation binding the contract event 0x468305133c010ade14f30fceab8c26460cc84b5efb05d75e7c1b3e290ef8fefa.
//
// Solidity: event DeleteBoeNode(address indexed coinbase)
func (_Nodes *NodesFilterer) FilterDeleteBoeNode(opts *bind.FilterOpts, coinbase []common.Address) (*NodesDeleteBoeNodeIterator, error) {

	var coinbaseRule []interface{}
	for _, coinbaseItem := range coinbase {
		coinbaseRule = append(coinbaseRule, coinbaseItem)
	}

	logs, sub, err := _Nodes.contract.FilterLogs(opts, "DeleteBoeNode", coinbaseRule)
	if err != nil {
		return nil, err
	}
	return &NodesDeleteBoeNodeIterator{contract: _Nodes.contract, event: "DeleteBoeNode", logs: logs, sub: sub}, nil
}

// WatchDeleteBoeNode is a free log subscription operation binding the contract event 0x468305133c010ade14f30fceab8c26460cc84b5efb05d75e7c1b3e290ef8fefa.
//
// Solidity: event DeleteBoeNode(address indexed coinbase)
func (_Nodes *NodesFilterer) WatchDeleteBoeNode(opts *bind.WatchOpts, sink chan<- *NodesDeleteBoeNode, coinbase []common.Address) (event.Subscription, error) {

	var coinbaseRule []interface{}
	for _, coinbaseItem := range coinbase {
		coinbaseRule = append(coinbaseRule, coinbaseItem)
	}

	logs, sub, err := _Nodes.contract.WatchLogs(opts, "DeleteBoeNode", coinbaseRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NodesDeleteBoeNode)
				if err := _Nodes.contract.UnpackLog(event, "DeleteBoeNode", log); err != nil {
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

// ParseDeleteBoeNode is a log parse operation binding the contract event 0x468305133c010ade14f30fceab8c26460cc84b5efb05d75e7c1b3e290ef8fefa.
//
// Solidity: event DeleteBoeNode(address indexed coinbase)
func (_Nodes *NodesFilterer) ParseDeleteBoeNode(log types.Log) (*NodesDeleteBoeNode, error) {
	event := new(NodesDeleteBoeNode)
	if err := _Nodes.contract.UnpackLog(event, "DeleteBoeNode", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NodesReceivedHpbIterator is returned from FilterReceivedHpb and is used to iterate over the raw logs and unpacked data for ReceivedHpb events raised by the Nodes contract.
type NodesReceivedHpbIterator struct {
	Event *NodesReceivedHpb // Event containing the contract specifics and raw log

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
func (it *NodesReceivedHpbIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NodesReceivedHpb)
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
		it.Event = new(NodesReceivedHpb)
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
func (it *NodesReceivedHpbIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NodesReceivedHpbIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NodesReceivedHpb represents a ReceivedHpb event raised by the Nodes contract.
type NodesReceivedHpb struct {
	Sender common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterReceivedHpb is a free log retrieval operation binding the contract event 0x7129701436f0cdc265d1e2cda298e8a1ccd6ed5fce7f69343e16530b07a2e06e.
//
// Solidity: event ReceivedHpb(address indexed sender, uint256 amount)
func (_Nodes *NodesFilterer) FilterReceivedHpb(opts *bind.FilterOpts, sender []common.Address) (*NodesReceivedHpbIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Nodes.contract.FilterLogs(opts, "ReceivedHpb", senderRule)
	if err != nil {
		return nil, err
	}
	return &NodesReceivedHpbIterator{contract: _Nodes.contract, event: "ReceivedHpb", logs: logs, sub: sub}, nil
}

// WatchReceivedHpb is a free log subscription operation binding the contract event 0x7129701436f0cdc265d1e2cda298e8a1ccd6ed5fce7f69343e16530b07a2e06e.
//
// Solidity: event ReceivedHpb(address indexed sender, uint256 amount)
func (_Nodes *NodesFilterer) WatchReceivedHpb(opts *bind.WatchOpts, sink chan<- *NodesReceivedHpb, sender []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Nodes.contract.WatchLogs(opts, "ReceivedHpb", senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NodesReceivedHpb)
				if err := _Nodes.contract.UnpackLog(event, "ReceivedHpb", log); err != nil {
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
func (_Nodes *NodesFilterer) ParseReceivedHpb(log types.Log) (*NodesReceivedHpb, error) {
	event := new(NodesReceivedHpb)
	if err := _Nodes.contract.UnpackLog(event, "ReceivedHpb", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NodesSetHolderAddrIterator is returned from FilterSetHolderAddr and is used to iterate over the raw logs and unpacked data for SetHolderAddr events raised by the Nodes contract.
type NodesSetHolderAddrIterator struct {
	Event *NodesSetHolderAddr // Event containing the contract specifics and raw log

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
func (it *NodesSetHolderAddrIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NodesSetHolderAddr)
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
		it.Event = new(NodesSetHolderAddr)
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
func (it *NodesSetHolderAddrIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NodesSetHolderAddrIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NodesSetHolderAddr represents a SetHolderAddr event raised by the Nodes contract.
type NodesSetHolderAddr struct {
	CoinBase   common.Address
	HolderAddr common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterSetHolderAddr is a free log retrieval operation binding the contract event 0xca5ebc9b84eae8eaaa14db3d55b033e7bfe33877548ac26baaa0b5d9847a2fdf.
//
// Solidity: event SetHolderAddr(address indexed coinBase, address indexed holderAddr)
func (_Nodes *NodesFilterer) FilterSetHolderAddr(opts *bind.FilterOpts, coinBase []common.Address, holderAddr []common.Address) (*NodesSetHolderAddrIterator, error) {

	var coinBaseRule []interface{}
	for _, coinBaseItem := range coinBase {
		coinBaseRule = append(coinBaseRule, coinBaseItem)
	}
	var holderAddrRule []interface{}
	for _, holderAddrItem := range holderAddr {
		holderAddrRule = append(holderAddrRule, holderAddrItem)
	}

	logs, sub, err := _Nodes.contract.FilterLogs(opts, "SetHolderAddr", coinBaseRule, holderAddrRule)
	if err != nil {
		return nil, err
	}
	return &NodesSetHolderAddrIterator{contract: _Nodes.contract, event: "SetHolderAddr", logs: logs, sub: sub}, nil
}

// WatchSetHolderAddr is a free log subscription operation binding the contract event 0xca5ebc9b84eae8eaaa14db3d55b033e7bfe33877548ac26baaa0b5d9847a2fdf.
//
// Solidity: event SetHolderAddr(address indexed coinBase, address indexed holderAddr)
func (_Nodes *NodesFilterer) WatchSetHolderAddr(opts *bind.WatchOpts, sink chan<- *NodesSetHolderAddr, coinBase []common.Address, holderAddr []common.Address) (event.Subscription, error) {

	var coinBaseRule []interface{}
	for _, coinBaseItem := range coinBase {
		coinBaseRule = append(coinBaseRule, coinBaseItem)
	}
	var holderAddrRule []interface{}
	for _, holderAddrItem := range holderAddr {
		holderAddrRule = append(holderAddrRule, holderAddrItem)
	}

	logs, sub, err := _Nodes.contract.WatchLogs(opts, "SetHolderAddr", coinBaseRule, holderAddrRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NodesSetHolderAddr)
				if err := _Nodes.contract.UnpackLog(event, "SetHolderAddr", log); err != nil {
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

// ParseSetHolderAddr is a log parse operation binding the contract event 0xca5ebc9b84eae8eaaa14db3d55b033e7bfe33877548ac26baaa0b5d9847a2fdf.
//
// Solidity: event SetHolderAddr(address indexed coinBase, address indexed holderAddr)
func (_Nodes *NodesFilterer) ParseSetHolderAddr(log types.Log) (*NodesSetHolderAddr, error) {
	event := new(NodesSetHolderAddr)
	if err := _Nodes.contract.UnpackLog(event, "SetHolderAddr", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NodesUpdateBoeNodeIterator is returned from FilterUpdateBoeNode and is used to iterate over the raw logs and unpacked data for UpdateBoeNode events raised by the Nodes contract.
type NodesUpdateBoeNodeIterator struct {
	Event *NodesUpdateBoeNode // Event containing the contract specifics and raw log

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
func (it *NodesUpdateBoeNodeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NodesUpdateBoeNode)
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
		it.Event = new(NodesUpdateBoeNode)
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
func (it *NodesUpdateBoeNodeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NodesUpdateBoeNodeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NodesUpdateBoeNode represents a UpdateBoeNode event raised by the Nodes contract.
type NodesUpdateBoeNode struct {
	Coinbase common.Address
	Cid1     [32]byte
	Cid2     [32]byte
	Hid      [32]byte
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterUpdateBoeNode is a free log retrieval operation binding the contract event 0x5d6bae5ddcb39615e06b1e2e5926890ad3f2b1dbe5d43bc7350c83d40e891426.
//
// Solidity: event UpdateBoeNode(address indexed coinbase, bytes32 cid1, bytes32 cid2, bytes32 indexed hid)
func (_Nodes *NodesFilterer) FilterUpdateBoeNode(opts *bind.FilterOpts, coinbase []common.Address, hid [][32]byte) (*NodesUpdateBoeNodeIterator, error) {

	var coinbaseRule []interface{}
	for _, coinbaseItem := range coinbase {
		coinbaseRule = append(coinbaseRule, coinbaseItem)
	}

	var hidRule []interface{}
	for _, hidItem := range hid {
		hidRule = append(hidRule, hidItem)
	}

	logs, sub, err := _Nodes.contract.FilterLogs(opts, "UpdateBoeNode", coinbaseRule, hidRule)
	if err != nil {
		return nil, err
	}
	return &NodesUpdateBoeNodeIterator{contract: _Nodes.contract, event: "UpdateBoeNode", logs: logs, sub: sub}, nil
}

// WatchUpdateBoeNode is a free log subscription operation binding the contract event 0x5d6bae5ddcb39615e06b1e2e5926890ad3f2b1dbe5d43bc7350c83d40e891426.
//
// Solidity: event UpdateBoeNode(address indexed coinbase, bytes32 cid1, bytes32 cid2, bytes32 indexed hid)
func (_Nodes *NodesFilterer) WatchUpdateBoeNode(opts *bind.WatchOpts, sink chan<- *NodesUpdateBoeNode, coinbase []common.Address, hid [][32]byte) (event.Subscription, error) {

	var coinbaseRule []interface{}
	for _, coinbaseItem := range coinbase {
		coinbaseRule = append(coinbaseRule, coinbaseItem)
	}

	var hidRule []interface{}
	for _, hidItem := range hid {
		hidRule = append(hidRule, hidItem)
	}

	logs, sub, err := _Nodes.contract.WatchLogs(opts, "UpdateBoeNode", coinbaseRule, hidRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NodesUpdateBoeNode)
				if err := _Nodes.contract.UnpackLog(event, "UpdateBoeNode", log); err != nil {
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

// ParseUpdateBoeNode is a log parse operation binding the contract event 0x5d6bae5ddcb39615e06b1e2e5926890ad3f2b1dbe5d43bc7350c83d40e891426.
//
// Solidity: event UpdateBoeNode(address indexed coinbase, bytes32 cid1, bytes32 cid2, bytes32 indexed hid)
func (_Nodes *NodesFilterer) ParseUpdateBoeNode(log types.Log) (*NodesUpdateBoeNode, error) {
	event := new(NodesUpdateBoeNode)
	if err := _Nodes.contract.UnpackLog(event, "UpdateBoeNode", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
