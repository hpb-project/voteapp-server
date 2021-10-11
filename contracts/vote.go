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

// VoteABI is the input ABI used to generate the binding from.
const VoteABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"fetchAllVoteResult\",\"outputs\":[{\"name\":\"\",\"type\":\"address[]\"},{\"name\":\"\",\"type\":\"uint256[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"voterAddr\",\"type\":\"address\"}],\"name\":\"fetchVoteNumForVoter\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"voterAddr\",\"type\":\"address\"}],\"name\":\"fetchVoteInfoForVoter\",\"outputs\":[{\"name\":\"\",\"type\":\"address[]\"},{\"name\":\"\",\"type\":\"uint256[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"minLimit\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"deleteAdmin\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"gasLeftLimit\",\"type\":\"uint256\"}],\"name\":\"setGasLeftLimit\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getAdmins\",\"outputs\":[{\"name\":\"\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"kill\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"voterAddr\",\"type\":\"address\"},{\"name\":\"boeAddr\",\"type\":\"address\"}],\"name\":\"fetchVoteNumForVoterToCandidate\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"boeaddr\",\"type\":\"address\"},{\"name\":\"num\",\"type\":\"uint256\"}],\"name\":\"vote\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_minLimit\",\"type\":\"uint256\"}],\"name\":\"setMinLimit\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"boeaddr\",\"type\":\"address\"}],\"name\":\"cancelVoteForCandidate\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"addAdmin\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"boeaddr\",\"type\":\"address\"}],\"name\":\"fetchVoteNumForCandidate\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"refreshVoteForAll\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getOwner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"cancelVote\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getNodeContract\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"fetchAllVoters\",\"outputs\":[{\"name\":\"\",\"type\":\"address[]\"},{\"name\":\"\",\"type\":\"uint256[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"adminMap\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"boeaddrs\",\"type\":\"address[]\"},{\"name\":\"nums\",\"type\":\"uint256[]\"}],\"name\":\"batchvote\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"boeaddr\",\"type\":\"address\"}],\"name\":\"fetchVoteInfoForCandidate\",\"outputs\":[{\"name\":\"\",\"type\":\"address[]\"},{\"name\":\"\",\"type\":\"uint256[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"setNodeContract\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"constructor\"},{\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"approved\",\"type\":\"bool\"},{\"indexed\":true,\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"ApprovalFor\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"flag\",\"type\":\"uint256\"},{\"indexed\":true,\"name\":\"candidateAddr\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"voteAddr\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"num\",\"type\":\"uint256\"}],\"name\":\"DoVoted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"ReceivedHpb\",\"type\":\"event\"}]"

// Vote is an auto generated Go binding around an Ethereum contract.
type Vote struct {
	VoteCaller     // Read-only binding to the contract
	VoteTransactor // Write-only binding to the contract
	VoteFilterer   // Log filterer for contract events
}

// VoteCaller is an auto generated read-only Go binding around an Ethereum contract.
type VoteCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VoteTransactor is an auto generated write-only Go binding around an Ethereum contract.
type VoteTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VoteFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type VoteFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VoteSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type VoteSession struct {
	Contract     *Vote             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// VoteCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type VoteCallerSession struct {
	Contract *VoteCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// VoteTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type VoteTransactorSession struct {
	Contract     *VoteTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// VoteRaw is an auto generated low-level Go binding around an Ethereum contract.
type VoteRaw struct {
	Contract *Vote // Generic contract binding to access the raw methods on
}

// VoteCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type VoteCallerRaw struct {
	Contract *VoteCaller // Generic read-only contract binding to access the raw methods on
}

// VoteTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type VoteTransactorRaw struct {
	Contract *VoteTransactor // Generic write-only contract binding to access the raw methods on
}

// NewVote creates a new instance of Vote, bound to a specific deployed contract.
func NewVote(address common.Address, backend bind.ContractBackend) (*Vote, error) {
	contract, err := bindVote(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Vote{VoteCaller: VoteCaller{contract: contract}, VoteTransactor: VoteTransactor{contract: contract}, VoteFilterer: VoteFilterer{contract: contract}}, nil
}

// NewVoteCaller creates a new read-only instance of Vote, bound to a specific deployed contract.
func NewVoteCaller(address common.Address, caller bind.ContractCaller) (*VoteCaller, error) {
	contract, err := bindVote(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &VoteCaller{contract: contract}, nil
}

// NewVoteTransactor creates a new write-only instance of Vote, bound to a specific deployed contract.
func NewVoteTransactor(address common.Address, transactor bind.ContractTransactor) (*VoteTransactor, error) {
	contract, err := bindVote(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &VoteTransactor{contract: contract}, nil
}

// NewVoteFilterer creates a new log filterer instance of Vote, bound to a specific deployed contract.
func NewVoteFilterer(address common.Address, filterer bind.ContractFilterer) (*VoteFilterer, error) {
	contract, err := bindVote(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &VoteFilterer{contract: contract}, nil
}

// bindVote binds a generic wrapper to an already deployed contract.
func bindVote(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(VoteABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Vote *VoteRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Vote.Contract.VoteCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Vote *VoteRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Vote.Contract.VoteTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Vote *VoteRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Vote.Contract.VoteTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Vote *VoteCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Vote.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Vote *VoteTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Vote.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Vote *VoteTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Vote.Contract.contract.Transact(opts, method, params...)
}

// AdminMap is a free data retrieval call binding the contract method 0xdbbc830b.
//
// Solidity: function adminMap(address ) view returns(address)
func (_Vote *VoteCaller) AdminMap(opts *bind.CallOpts, arg0 common.Address) (common.Address, error) {
	var out []interface{}
	err := _Vote.contract.Call(opts, &out, "adminMap", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AdminMap is a free data retrieval call binding the contract method 0xdbbc830b.
//
// Solidity: function adminMap(address ) view returns(address)
func (_Vote *VoteSession) AdminMap(arg0 common.Address) (common.Address, error) {
	return _Vote.Contract.AdminMap(&_Vote.CallOpts, arg0)
}

// AdminMap is a free data retrieval call binding the contract method 0xdbbc830b.
//
// Solidity: function adminMap(address ) view returns(address)
func (_Vote *VoteCallerSession) AdminMap(arg0 common.Address) (common.Address, error) {
	return _Vote.Contract.AdminMap(&_Vote.CallOpts, arg0)
}

// FetchAllVoteResult is a free data retrieval call binding the contract method 0x0304363a.
//
// Solidity: function fetchAllVoteResult() view returns(address[], uint256[])
func (_Vote *VoteCaller) FetchAllVoteResult(opts *bind.CallOpts) ([]common.Address, []*big.Int, error) {
	var out []interface{}
	err := _Vote.contract.Call(opts, &out, "fetchAllVoteResult")

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
func (_Vote *VoteSession) FetchAllVoteResult() ([]common.Address, []*big.Int, error) {
	return _Vote.Contract.FetchAllVoteResult(&_Vote.CallOpts)
}

// FetchAllVoteResult is a free data retrieval call binding the contract method 0x0304363a.
//
// Solidity: function fetchAllVoteResult() view returns(address[], uint256[])
func (_Vote *VoteCallerSession) FetchAllVoteResult() ([]common.Address, []*big.Int, error) {
	return _Vote.Contract.FetchAllVoteResult(&_Vote.CallOpts)
}

// FetchAllVoters is a free data retrieval call binding the contract method 0xd2d3d7fb.
//
// Solidity: function fetchAllVoters() view returns(address[], uint256[])
func (_Vote *VoteCaller) FetchAllVoters(opts *bind.CallOpts) ([]common.Address, []*big.Int, error) {
	var out []interface{}
	err := _Vote.contract.Call(opts, &out, "fetchAllVoters")

	if err != nil {
		return *new([]common.Address), *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)
	out1 := *abi.ConvertType(out[1], new([]*big.Int)).(*[]*big.Int)

	return out0, out1, err

}

// FetchAllVoters is a free data retrieval call binding the contract method 0xd2d3d7fb.
//
// Solidity: function fetchAllVoters() view returns(address[], uint256[])
func (_Vote *VoteSession) FetchAllVoters() ([]common.Address, []*big.Int, error) {
	return _Vote.Contract.FetchAllVoters(&_Vote.CallOpts)
}

// FetchAllVoters is a free data retrieval call binding the contract method 0xd2d3d7fb.
//
// Solidity: function fetchAllVoters() view returns(address[], uint256[])
func (_Vote *VoteCallerSession) FetchAllVoters() ([]common.Address, []*big.Int, error) {
	return _Vote.Contract.FetchAllVoters(&_Vote.CallOpts)
}

// FetchVoteInfoForCandidate is a free data retrieval call binding the contract method 0xf25d2422.
//
// Solidity: function fetchVoteInfoForCandidate(address boeaddr) view returns(address[], uint256[])
func (_Vote *VoteCaller) FetchVoteInfoForCandidate(opts *bind.CallOpts, boeaddr common.Address) ([]common.Address, []*big.Int, error) {
	var out []interface{}
	err := _Vote.contract.Call(opts, &out, "fetchVoteInfoForCandidate", boeaddr)

	if err != nil {
		return *new([]common.Address), *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)
	out1 := *abi.ConvertType(out[1], new([]*big.Int)).(*[]*big.Int)

	return out0, out1, err

}

// FetchVoteInfoForCandidate is a free data retrieval call binding the contract method 0xf25d2422.
//
// Solidity: function fetchVoteInfoForCandidate(address boeaddr) view returns(address[], uint256[])
func (_Vote *VoteSession) FetchVoteInfoForCandidate(boeaddr common.Address) ([]common.Address, []*big.Int, error) {
	return _Vote.Contract.FetchVoteInfoForCandidate(&_Vote.CallOpts, boeaddr)
}

// FetchVoteInfoForCandidate is a free data retrieval call binding the contract method 0xf25d2422.
//
// Solidity: function fetchVoteInfoForCandidate(address boeaddr) view returns(address[], uint256[])
func (_Vote *VoteCallerSession) FetchVoteInfoForCandidate(boeaddr common.Address) ([]common.Address, []*big.Int, error) {
	return _Vote.Contract.FetchVoteInfoForCandidate(&_Vote.CallOpts, boeaddr)
}

// FetchVoteInfoForVoter is a free data retrieval call binding the contract method 0x1391ff43.
//
// Solidity: function fetchVoteInfoForVoter(address voterAddr) view returns(address[], uint256[])
func (_Vote *VoteCaller) FetchVoteInfoForVoter(opts *bind.CallOpts, voterAddr common.Address) ([]common.Address, []*big.Int, error) {
	var out []interface{}
	err := _Vote.contract.Call(opts, &out, "fetchVoteInfoForVoter", voterAddr)

	if err != nil {
		return *new([]common.Address), *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)
	out1 := *abi.ConvertType(out[1], new([]*big.Int)).(*[]*big.Int)

	return out0, out1, err

}

// FetchVoteInfoForVoter is a free data retrieval call binding the contract method 0x1391ff43.
//
// Solidity: function fetchVoteInfoForVoter(address voterAddr) view returns(address[], uint256[])
func (_Vote *VoteSession) FetchVoteInfoForVoter(voterAddr common.Address) ([]common.Address, []*big.Int, error) {
	return _Vote.Contract.FetchVoteInfoForVoter(&_Vote.CallOpts, voterAddr)
}

// FetchVoteInfoForVoter is a free data retrieval call binding the contract method 0x1391ff43.
//
// Solidity: function fetchVoteInfoForVoter(address voterAddr) view returns(address[], uint256[])
func (_Vote *VoteCallerSession) FetchVoteInfoForVoter(voterAddr common.Address) ([]common.Address, []*big.Int, error) {
	return _Vote.Contract.FetchVoteInfoForVoter(&_Vote.CallOpts, voterAddr)
}

// FetchVoteNumForCandidate is a free data retrieval call binding the contract method 0x753cb4eb.
//
// Solidity: function fetchVoteNumForCandidate(address boeaddr) view returns(uint256)
func (_Vote *VoteCaller) FetchVoteNumForCandidate(opts *bind.CallOpts, boeaddr common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Vote.contract.Call(opts, &out, "fetchVoteNumForCandidate", boeaddr)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FetchVoteNumForCandidate is a free data retrieval call binding the contract method 0x753cb4eb.
//
// Solidity: function fetchVoteNumForCandidate(address boeaddr) view returns(uint256)
func (_Vote *VoteSession) FetchVoteNumForCandidate(boeaddr common.Address) (*big.Int, error) {
	return _Vote.Contract.FetchVoteNumForCandidate(&_Vote.CallOpts, boeaddr)
}

// FetchVoteNumForCandidate is a free data retrieval call binding the contract method 0x753cb4eb.
//
// Solidity: function fetchVoteNumForCandidate(address boeaddr) view returns(uint256)
func (_Vote *VoteCallerSession) FetchVoteNumForCandidate(boeaddr common.Address) (*big.Int, error) {
	return _Vote.Contract.FetchVoteNumForCandidate(&_Vote.CallOpts, boeaddr)
}

// FetchVoteNumForVoter is a free data retrieval call binding the contract method 0x062ff957.
//
// Solidity: function fetchVoteNumForVoter(address voterAddr) view returns(uint256)
func (_Vote *VoteCaller) FetchVoteNumForVoter(opts *bind.CallOpts, voterAddr common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Vote.contract.Call(opts, &out, "fetchVoteNumForVoter", voterAddr)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FetchVoteNumForVoter is a free data retrieval call binding the contract method 0x062ff957.
//
// Solidity: function fetchVoteNumForVoter(address voterAddr) view returns(uint256)
func (_Vote *VoteSession) FetchVoteNumForVoter(voterAddr common.Address) (*big.Int, error) {
	return _Vote.Contract.FetchVoteNumForVoter(&_Vote.CallOpts, voterAddr)
}

// FetchVoteNumForVoter is a free data retrieval call binding the contract method 0x062ff957.
//
// Solidity: function fetchVoteNumForVoter(address voterAddr) view returns(uint256)
func (_Vote *VoteCallerSession) FetchVoteNumForVoter(voterAddr common.Address) (*big.Int, error) {
	return _Vote.Contract.FetchVoteNumForVoter(&_Vote.CallOpts, voterAddr)
}

// FetchVoteNumForVoterToCandidate is a free data retrieval call binding the contract method 0x5c857333.
//
// Solidity: function fetchVoteNumForVoterToCandidate(address voterAddr, address boeAddr) view returns(uint256)
func (_Vote *VoteCaller) FetchVoteNumForVoterToCandidate(opts *bind.CallOpts, voterAddr common.Address, boeAddr common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Vote.contract.Call(opts, &out, "fetchVoteNumForVoterToCandidate", voterAddr, boeAddr)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FetchVoteNumForVoterToCandidate is a free data retrieval call binding the contract method 0x5c857333.
//
// Solidity: function fetchVoteNumForVoterToCandidate(address voterAddr, address boeAddr) view returns(uint256)
func (_Vote *VoteSession) FetchVoteNumForVoterToCandidate(voterAddr common.Address, boeAddr common.Address) (*big.Int, error) {
	return _Vote.Contract.FetchVoteNumForVoterToCandidate(&_Vote.CallOpts, voterAddr, boeAddr)
}

// FetchVoteNumForVoterToCandidate is a free data retrieval call binding the contract method 0x5c857333.
//
// Solidity: function fetchVoteNumForVoterToCandidate(address voterAddr, address boeAddr) view returns(uint256)
func (_Vote *VoteCallerSession) FetchVoteNumForVoterToCandidate(voterAddr common.Address, boeAddr common.Address) (*big.Int, error) {
	return _Vote.Contract.FetchVoteNumForVoterToCandidate(&_Vote.CallOpts, voterAddr, boeAddr)
}

// GetAdmins is a free data retrieval call binding the contract method 0x31ae450b.
//
// Solidity: function getAdmins() view returns(address[])
func (_Vote *VoteCaller) GetAdmins(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _Vote.contract.Call(opts, &out, "getAdmins")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetAdmins is a free data retrieval call binding the contract method 0x31ae450b.
//
// Solidity: function getAdmins() view returns(address[])
func (_Vote *VoteSession) GetAdmins() ([]common.Address, error) {
	return _Vote.Contract.GetAdmins(&_Vote.CallOpts)
}

// GetAdmins is a free data retrieval call binding the contract method 0x31ae450b.
//
// Solidity: function getAdmins() view returns(address[])
func (_Vote *VoteCallerSession) GetAdmins() ([]common.Address, error) {
	return _Vote.Contract.GetAdmins(&_Vote.CallOpts)
}

// GetNodeContract is a free data retrieval call binding the contract method 0x9fd2e0be.
//
// Solidity: function getNodeContract() view returns(address)
func (_Vote *VoteCaller) GetNodeContract(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Vote.contract.Call(opts, &out, "getNodeContract")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetNodeContract is a free data retrieval call binding the contract method 0x9fd2e0be.
//
// Solidity: function getNodeContract() view returns(address)
func (_Vote *VoteSession) GetNodeContract() (common.Address, error) {
	return _Vote.Contract.GetNodeContract(&_Vote.CallOpts)
}

// GetNodeContract is a free data retrieval call binding the contract method 0x9fd2e0be.
//
// Solidity: function getNodeContract() view returns(address)
func (_Vote *VoteCallerSession) GetNodeContract() (common.Address, error) {
	return _Vote.Contract.GetNodeContract(&_Vote.CallOpts)
}

// GetOwner is a free data retrieval call binding the contract method 0x893d20e8.
//
// Solidity: function getOwner() view returns(address)
func (_Vote *VoteCaller) GetOwner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Vote.contract.Call(opts, &out, "getOwner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetOwner is a free data retrieval call binding the contract method 0x893d20e8.
//
// Solidity: function getOwner() view returns(address)
func (_Vote *VoteSession) GetOwner() (common.Address, error) {
	return _Vote.Contract.GetOwner(&_Vote.CallOpts)
}

// GetOwner is a free data retrieval call binding the contract method 0x893d20e8.
//
// Solidity: function getOwner() view returns(address)
func (_Vote *VoteCallerSession) GetOwner() (common.Address, error) {
	return _Vote.Contract.GetOwner(&_Vote.CallOpts)
}

// MinLimit is a free data retrieval call binding the contract method 0x1fd8088d.
//
// Solidity: function minLimit() view returns(uint256)
func (_Vote *VoteCaller) MinLimit(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Vote.contract.Call(opts, &out, "minLimit")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinLimit is a free data retrieval call binding the contract method 0x1fd8088d.
//
// Solidity: function minLimit() view returns(uint256)
func (_Vote *VoteSession) MinLimit() (*big.Int, error) {
	return _Vote.Contract.MinLimit(&_Vote.CallOpts)
}

// MinLimit is a free data retrieval call binding the contract method 0x1fd8088d.
//
// Solidity: function minLimit() view returns(uint256)
func (_Vote *VoteCallerSession) MinLimit() (*big.Int, error) {
	return _Vote.Contract.MinLimit(&_Vote.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Vote *VoteCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Vote.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Vote *VoteSession) Owner() (common.Address, error) {
	return _Vote.Contract.Owner(&_Vote.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Vote *VoteCallerSession) Owner() (common.Address, error) {
	return _Vote.Contract.Owner(&_Vote.CallOpts)
}

// AddAdmin is a paid mutator transaction binding the contract method 0x70480275.
//
// Solidity: function addAdmin(address addr) returns(bool)
func (_Vote *VoteTransactor) AddAdmin(opts *bind.TransactOpts, addr common.Address) (*types.Transaction, error) {
	return _Vote.contract.Transact(opts, "addAdmin", addr)
}

// AddAdmin is a paid mutator transaction binding the contract method 0x70480275.
//
// Solidity: function addAdmin(address addr) returns(bool)
func (_Vote *VoteSession) AddAdmin(addr common.Address) (*types.Transaction, error) {
	return _Vote.Contract.AddAdmin(&_Vote.TransactOpts, addr)
}

// AddAdmin is a paid mutator transaction binding the contract method 0x70480275.
//
// Solidity: function addAdmin(address addr) returns(bool)
func (_Vote *VoteTransactorSession) AddAdmin(addr common.Address) (*types.Transaction, error) {
	return _Vote.Contract.AddAdmin(&_Vote.TransactOpts, addr)
}

// Batchvote is a paid mutator transaction binding the contract method 0xe1e325c1.
//
// Solidity: function batchvote(address[] boeaddrs, uint256[] nums) returns()
func (_Vote *VoteTransactor) Batchvote(opts *bind.TransactOpts, boeaddrs []common.Address, nums []*big.Int) (*types.Transaction, error) {
	return _Vote.contract.Transact(opts, "batchvote", boeaddrs, nums)
}

// Batchvote is a paid mutator transaction binding the contract method 0xe1e325c1.
//
// Solidity: function batchvote(address[] boeaddrs, uint256[] nums) returns()
func (_Vote *VoteSession) Batchvote(boeaddrs []common.Address, nums []*big.Int) (*types.Transaction, error) {
	return _Vote.Contract.Batchvote(&_Vote.TransactOpts, boeaddrs, nums)
}

// Batchvote is a paid mutator transaction binding the contract method 0xe1e325c1.
//
// Solidity: function batchvote(address[] boeaddrs, uint256[] nums) returns()
func (_Vote *VoteTransactorSession) Batchvote(boeaddrs []common.Address, nums []*big.Int) (*types.Transaction, error) {
	return _Vote.Contract.Batchvote(&_Vote.TransactOpts, boeaddrs, nums)
}

// CancelVote is a paid mutator transaction binding the contract method 0x95c96554.
//
// Solidity: function cancelVote() returns()
func (_Vote *VoteTransactor) CancelVote(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Vote.contract.Transact(opts, "cancelVote")
}

// CancelVote is a paid mutator transaction binding the contract method 0x95c96554.
//
// Solidity: function cancelVote() returns()
func (_Vote *VoteSession) CancelVote() (*types.Transaction, error) {
	return _Vote.Contract.CancelVote(&_Vote.TransactOpts)
}

// CancelVote is a paid mutator transaction binding the contract method 0x95c96554.
//
// Solidity: function cancelVote() returns()
func (_Vote *VoteTransactorSession) CancelVote() (*types.Transaction, error) {
	return _Vote.Contract.CancelVote(&_Vote.TransactOpts)
}

// CancelVoteForCandidate is a paid mutator transaction binding the contract method 0x6eefbd7d.
//
// Solidity: function cancelVoteForCandidate(address boeaddr) returns(bool)
func (_Vote *VoteTransactor) CancelVoteForCandidate(opts *bind.TransactOpts, boeaddr common.Address) (*types.Transaction, error) {
	return _Vote.contract.Transact(opts, "cancelVoteForCandidate", boeaddr)
}

// CancelVoteForCandidate is a paid mutator transaction binding the contract method 0x6eefbd7d.
//
// Solidity: function cancelVoteForCandidate(address boeaddr) returns(bool)
func (_Vote *VoteSession) CancelVoteForCandidate(boeaddr common.Address) (*types.Transaction, error) {
	return _Vote.Contract.CancelVoteForCandidate(&_Vote.TransactOpts, boeaddr)
}

// CancelVoteForCandidate is a paid mutator transaction binding the contract method 0x6eefbd7d.
//
// Solidity: function cancelVoteForCandidate(address boeaddr) returns(bool)
func (_Vote *VoteTransactorSession) CancelVoteForCandidate(boeaddr common.Address) (*types.Transaction, error) {
	return _Vote.Contract.CancelVoteForCandidate(&_Vote.TransactOpts, boeaddr)
}

// DeleteAdmin is a paid mutator transaction binding the contract method 0x27e1f7df.
//
// Solidity: function deleteAdmin(address addr) returns(bool)
func (_Vote *VoteTransactor) DeleteAdmin(opts *bind.TransactOpts, addr common.Address) (*types.Transaction, error) {
	return _Vote.contract.Transact(opts, "deleteAdmin", addr)
}

// DeleteAdmin is a paid mutator transaction binding the contract method 0x27e1f7df.
//
// Solidity: function deleteAdmin(address addr) returns(bool)
func (_Vote *VoteSession) DeleteAdmin(addr common.Address) (*types.Transaction, error) {
	return _Vote.Contract.DeleteAdmin(&_Vote.TransactOpts, addr)
}

// DeleteAdmin is a paid mutator transaction binding the contract method 0x27e1f7df.
//
// Solidity: function deleteAdmin(address addr) returns(bool)
func (_Vote *VoteTransactorSession) DeleteAdmin(addr common.Address) (*types.Transaction, error) {
	return _Vote.Contract.DeleteAdmin(&_Vote.TransactOpts, addr)
}

// Kill is a paid mutator transaction binding the contract method 0x41c0e1b5.
//
// Solidity: function kill() returns(bool)
func (_Vote *VoteTransactor) Kill(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Vote.contract.Transact(opts, "kill")
}

// Kill is a paid mutator transaction binding the contract method 0x41c0e1b5.
//
// Solidity: function kill() returns(bool)
func (_Vote *VoteSession) Kill() (*types.Transaction, error) {
	return _Vote.Contract.Kill(&_Vote.TransactOpts)
}

// Kill is a paid mutator transaction binding the contract method 0x41c0e1b5.
//
// Solidity: function kill() returns(bool)
func (_Vote *VoteTransactorSession) Kill() (*types.Transaction, error) {
	return _Vote.Contract.Kill(&_Vote.TransactOpts)
}

// RefreshVoteForAll is a paid mutator transaction binding the contract method 0x816194f8.
//
// Solidity: function refreshVoteForAll() returns()
func (_Vote *VoteTransactor) RefreshVoteForAll(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Vote.contract.Transact(opts, "refreshVoteForAll")
}

// RefreshVoteForAll is a paid mutator transaction binding the contract method 0x816194f8.
//
// Solidity: function refreshVoteForAll() returns()
func (_Vote *VoteSession) RefreshVoteForAll() (*types.Transaction, error) {
	return _Vote.Contract.RefreshVoteForAll(&_Vote.TransactOpts)
}

// RefreshVoteForAll is a paid mutator transaction binding the contract method 0x816194f8.
//
// Solidity: function refreshVoteForAll() returns()
func (_Vote *VoteTransactorSession) RefreshVoteForAll() (*types.Transaction, error) {
	return _Vote.Contract.RefreshVoteForAll(&_Vote.TransactOpts)
}

// SetGasLeftLimit is a paid mutator transaction binding the contract method 0x306419a6.
//
// Solidity: function setGasLeftLimit(uint256 gasLeftLimit) returns(bool)
func (_Vote *VoteTransactor) SetGasLeftLimit(opts *bind.TransactOpts, gasLeftLimit *big.Int) (*types.Transaction, error) {
	return _Vote.contract.Transact(opts, "setGasLeftLimit", gasLeftLimit)
}

// SetGasLeftLimit is a paid mutator transaction binding the contract method 0x306419a6.
//
// Solidity: function setGasLeftLimit(uint256 gasLeftLimit) returns(bool)
func (_Vote *VoteSession) SetGasLeftLimit(gasLeftLimit *big.Int) (*types.Transaction, error) {
	return _Vote.Contract.SetGasLeftLimit(&_Vote.TransactOpts, gasLeftLimit)
}

// SetGasLeftLimit is a paid mutator transaction binding the contract method 0x306419a6.
//
// Solidity: function setGasLeftLimit(uint256 gasLeftLimit) returns(bool)
func (_Vote *VoteTransactorSession) SetGasLeftLimit(gasLeftLimit *big.Int) (*types.Transaction, error) {
	return _Vote.Contract.SetGasLeftLimit(&_Vote.TransactOpts, gasLeftLimit)
}

// SetMinLimit is a paid mutator transaction binding the contract method 0x6ec6d4a6.
//
// Solidity: function setMinLimit(uint256 _minLimit) returns(bool)
func (_Vote *VoteTransactor) SetMinLimit(opts *bind.TransactOpts, _minLimit *big.Int) (*types.Transaction, error) {
	return _Vote.contract.Transact(opts, "setMinLimit", _minLimit)
}

// SetMinLimit is a paid mutator transaction binding the contract method 0x6ec6d4a6.
//
// Solidity: function setMinLimit(uint256 _minLimit) returns(bool)
func (_Vote *VoteSession) SetMinLimit(_minLimit *big.Int) (*types.Transaction, error) {
	return _Vote.Contract.SetMinLimit(&_Vote.TransactOpts, _minLimit)
}

// SetMinLimit is a paid mutator transaction binding the contract method 0x6ec6d4a6.
//
// Solidity: function setMinLimit(uint256 _minLimit) returns(bool)
func (_Vote *VoteTransactorSession) SetMinLimit(_minLimit *big.Int) (*types.Transaction, error) {
	return _Vote.Contract.SetMinLimit(&_Vote.TransactOpts, _minLimit)
}

// SetNodeContract is a paid mutator transaction binding the contract method 0xf50fd926.
//
// Solidity: function setNodeContract(address addr) returns()
func (_Vote *VoteTransactor) SetNodeContract(opts *bind.TransactOpts, addr common.Address) (*types.Transaction, error) {
	return _Vote.contract.Transact(opts, "setNodeContract", addr)
}

// SetNodeContract is a paid mutator transaction binding the contract method 0xf50fd926.
//
// Solidity: function setNodeContract(address addr) returns()
func (_Vote *VoteSession) SetNodeContract(addr common.Address) (*types.Transaction, error) {
	return _Vote.Contract.SetNodeContract(&_Vote.TransactOpts, addr)
}

// SetNodeContract is a paid mutator transaction binding the contract method 0xf50fd926.
//
// Solidity: function setNodeContract(address addr) returns()
func (_Vote *VoteTransactorSession) SetNodeContract(addr common.Address) (*types.Transaction, error) {
	return _Vote.Contract.SetNodeContract(&_Vote.TransactOpts, addr)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns(bool)
func (_Vote *VoteTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Vote.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns(bool)
func (_Vote *VoteSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Vote.Contract.TransferOwnership(&_Vote.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns(bool)
func (_Vote *VoteTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Vote.Contract.TransferOwnership(&_Vote.TransactOpts, newOwner)
}

// Vote is a paid mutator transaction binding the contract method 0x5f74bbde.
//
// Solidity: function vote(address boeaddr, uint256 num) returns()
func (_Vote *VoteTransactor) Vote(opts *bind.TransactOpts, boeaddr common.Address, num *big.Int) (*types.Transaction, error) {
	return _Vote.contract.Transact(opts, "vote", boeaddr, num)
}

// Vote is a paid mutator transaction binding the contract method 0x5f74bbde.
//
// Solidity: function vote(address boeaddr, uint256 num) returns()
func (_Vote *VoteSession) Vote(boeaddr common.Address, num *big.Int) (*types.Transaction, error) {
	return _Vote.Contract.Vote(&_Vote.TransactOpts, boeaddr, num)
}

// Vote is a paid mutator transaction binding the contract method 0x5f74bbde.
//
// Solidity: function vote(address boeaddr, uint256 num) returns()
func (_Vote *VoteTransactorSession) Vote(boeaddr common.Address, num *big.Int) (*types.Transaction, error) {
	return _Vote.Contract.Vote(&_Vote.TransactOpts, boeaddr, num)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_Vote *VoteTransactor) Fallback(opts *bind.TransactOpts, calldata []byte) (*types.Transaction, error) {
	return _Vote.contract.RawTransact(opts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_Vote *VoteSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _Vote.Contract.Fallback(&_Vote.TransactOpts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_Vote *VoteTransactorSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _Vote.Contract.Fallback(&_Vote.TransactOpts, calldata)
}

// VoteApprovalForIterator is returned from FilterApprovalFor and is used to iterate over the raw logs and unpacked data for ApprovalFor events raised by the Vote contract.
type VoteApprovalForIterator struct {
	Event *VoteApprovalFor // Event containing the contract specifics and raw log

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
func (it *VoteApprovalForIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VoteApprovalFor)
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
		it.Event = new(VoteApprovalFor)
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
func (it *VoteApprovalForIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VoteApprovalForIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VoteApprovalFor represents a ApprovalFor event raised by the Vote contract.
type VoteApprovalFor struct {
	Approved bool
	Operator common.Address
	Owner    common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalFor is a free log retrieval operation binding the contract event 0xd9652388f1a6dc7a64879d6be6aa07bd7d3126379b2ce2f3210c217dabdbd442.
//
// Solidity: event ApprovalFor(bool indexed approved, address indexed operator, address indexed owner)
func (_Vote *VoteFilterer) FilterApprovalFor(opts *bind.FilterOpts, approved []bool, operator []common.Address, owner []common.Address) (*VoteApprovalForIterator, error) {

	var approvedRule []interface{}
	for _, approvedItem := range approved {
		approvedRule = append(approvedRule, approvedItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _Vote.contract.FilterLogs(opts, "ApprovalFor", approvedRule, operatorRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return &VoteApprovalForIterator{contract: _Vote.contract, event: "ApprovalFor", logs: logs, sub: sub}, nil
}

// WatchApprovalFor is a free log subscription operation binding the contract event 0xd9652388f1a6dc7a64879d6be6aa07bd7d3126379b2ce2f3210c217dabdbd442.
//
// Solidity: event ApprovalFor(bool indexed approved, address indexed operator, address indexed owner)
func (_Vote *VoteFilterer) WatchApprovalFor(opts *bind.WatchOpts, sink chan<- *VoteApprovalFor, approved []bool, operator []common.Address, owner []common.Address) (event.Subscription, error) {

	var approvedRule []interface{}
	for _, approvedItem := range approved {
		approvedRule = append(approvedRule, approvedItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _Vote.contract.WatchLogs(opts, "ApprovalFor", approvedRule, operatorRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VoteApprovalFor)
				if err := _Vote.contract.UnpackLog(event, "ApprovalFor", log); err != nil {
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

// ParseApprovalFor is a log parse operation binding the contract event 0xd9652388f1a6dc7a64879d6be6aa07bd7d3126379b2ce2f3210c217dabdbd442.
//
// Solidity: event ApprovalFor(bool indexed approved, address indexed operator, address indexed owner)
func (_Vote *VoteFilterer) ParseApprovalFor(log types.Log) (*VoteApprovalFor, error) {
	event := new(VoteApprovalFor)
	if err := _Vote.contract.UnpackLog(event, "ApprovalFor", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VoteDoVotedIterator is returned from FilterDoVoted and is used to iterate over the raw logs and unpacked data for DoVoted events raised by the Vote contract.
type VoteDoVotedIterator struct {
	Event *VoteDoVoted // Event containing the contract specifics and raw log

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
func (it *VoteDoVotedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VoteDoVoted)
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
		it.Event = new(VoteDoVoted)
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
func (it *VoteDoVotedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VoteDoVotedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VoteDoVoted represents a DoVoted event raised by the Vote contract.
type VoteDoVoted struct {
	Flag          *big.Int
	CandidateAddr common.Address
	VoteAddr      common.Address
	Num           *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterDoVoted is a free log retrieval operation binding the contract event 0xea9260397205cf3e0b87085d1160d8d78a5f9d560ddcf5f52570abe1184ba46a.
//
// Solidity: event DoVoted(uint256 indexed flag, address indexed candidateAddr, address indexed voteAddr, uint256 num)
func (_Vote *VoteFilterer) FilterDoVoted(opts *bind.FilterOpts, flag []*big.Int, candidateAddr []common.Address, voteAddr []common.Address) (*VoteDoVotedIterator, error) {

	var flagRule []interface{}
	for _, flagItem := range flag {
		flagRule = append(flagRule, flagItem)
	}
	var candidateAddrRule []interface{}
	for _, candidateAddrItem := range candidateAddr {
		candidateAddrRule = append(candidateAddrRule, candidateAddrItem)
	}
	var voteAddrRule []interface{}
	for _, voteAddrItem := range voteAddr {
		voteAddrRule = append(voteAddrRule, voteAddrItem)
	}

	logs, sub, err := _Vote.contract.FilterLogs(opts, "DoVoted", flagRule, candidateAddrRule, voteAddrRule)
	if err != nil {
		return nil, err
	}
	return &VoteDoVotedIterator{contract: _Vote.contract, event: "DoVoted", logs: logs, sub: sub}, nil
}

// WatchDoVoted is a free log subscription operation binding the contract event 0xea9260397205cf3e0b87085d1160d8d78a5f9d560ddcf5f52570abe1184ba46a.
//
// Solidity: event DoVoted(uint256 indexed flag, address indexed candidateAddr, address indexed voteAddr, uint256 num)
func (_Vote *VoteFilterer) WatchDoVoted(opts *bind.WatchOpts, sink chan<- *VoteDoVoted, flag []*big.Int, candidateAddr []common.Address, voteAddr []common.Address) (event.Subscription, error) {

	var flagRule []interface{}
	for _, flagItem := range flag {
		flagRule = append(flagRule, flagItem)
	}
	var candidateAddrRule []interface{}
	for _, candidateAddrItem := range candidateAddr {
		candidateAddrRule = append(candidateAddrRule, candidateAddrItem)
	}
	var voteAddrRule []interface{}
	for _, voteAddrItem := range voteAddr {
		voteAddrRule = append(voteAddrRule, voteAddrItem)
	}

	logs, sub, err := _Vote.contract.WatchLogs(opts, "DoVoted", flagRule, candidateAddrRule, voteAddrRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VoteDoVoted)
				if err := _Vote.contract.UnpackLog(event, "DoVoted", log); err != nil {
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

// ParseDoVoted is a log parse operation binding the contract event 0xea9260397205cf3e0b87085d1160d8d78a5f9d560ddcf5f52570abe1184ba46a.
//
// Solidity: event DoVoted(uint256 indexed flag, address indexed candidateAddr, address indexed voteAddr, uint256 num)
func (_Vote *VoteFilterer) ParseDoVoted(log types.Log) (*VoteDoVoted, error) {
	event := new(VoteDoVoted)
	if err := _Vote.contract.UnpackLog(event, "DoVoted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VoteReceivedHpbIterator is returned from FilterReceivedHpb and is used to iterate over the raw logs and unpacked data for ReceivedHpb events raised by the Vote contract.
type VoteReceivedHpbIterator struct {
	Event *VoteReceivedHpb // Event containing the contract specifics and raw log

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
func (it *VoteReceivedHpbIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VoteReceivedHpb)
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
		it.Event = new(VoteReceivedHpb)
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
func (it *VoteReceivedHpbIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VoteReceivedHpbIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VoteReceivedHpb represents a ReceivedHpb event raised by the Vote contract.
type VoteReceivedHpb struct {
	Sender common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterReceivedHpb is a free log retrieval operation binding the contract event 0x7129701436f0cdc265d1e2cda298e8a1ccd6ed5fce7f69343e16530b07a2e06e.
//
// Solidity: event ReceivedHpb(address indexed sender, uint256 amount)
func (_Vote *VoteFilterer) FilterReceivedHpb(opts *bind.FilterOpts, sender []common.Address) (*VoteReceivedHpbIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Vote.contract.FilterLogs(opts, "ReceivedHpb", senderRule)
	if err != nil {
		return nil, err
	}
	return &VoteReceivedHpbIterator{contract: _Vote.contract, event: "ReceivedHpb", logs: logs, sub: sub}, nil
}

// WatchReceivedHpb is a free log subscription operation binding the contract event 0x7129701436f0cdc265d1e2cda298e8a1ccd6ed5fce7f69343e16530b07a2e06e.
//
// Solidity: event ReceivedHpb(address indexed sender, uint256 amount)
func (_Vote *VoteFilterer) WatchReceivedHpb(opts *bind.WatchOpts, sink chan<- *VoteReceivedHpb, sender []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Vote.contract.WatchLogs(opts, "ReceivedHpb", senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VoteReceivedHpb)
				if err := _Vote.contract.UnpackLog(event, "ReceivedHpb", log); err != nil {
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
func (_Vote *VoteFilterer) ParseReceivedHpb(log types.Log) (*VoteReceivedHpb, error) {
	event := new(VoteReceivedHpb)
	if err := _Vote.contract.UnpackLog(event, "ReceivedHpb", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
