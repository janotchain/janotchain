// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contract

import (
	"math/big"
	"strings"

	ethereum "github.com/janetachain/janetachain"
	"github.com/janetachain/janetachain/accounts/abi"
	"github.com/janetachain/janetachain/accounts/abi/bind"
	"github.com/janetachain/janetachain/common"
	"github.com/janetachain/janetachain/core/types"
	"github.com/janetachain/janetachain/event"
)

// SafeMathABI is the input ABI used to generate the binding from.
const SafeMathABI = "[]"

// SafeMathBin is the compiled bytecode used for deploying new contracts.
const SafeMathBin = `0x604c602c600b82828239805160001a60731460008114601c57601e565bfe5b5030600052607381538281f30073000000000000000000000000000000000000000030146060604052600080fd00a165627a7a72305820b9407d48ebc7efee5c9f08b3b3a957df2939281f5913225e8c1291f069b900490029`

// DeploySafeMath deploys a new Ethereum contract, binding an instance of SafeMath to it.
func DeploySafeMath(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *SafeMath, error) {
	parsed, err := abi.JSON(strings.NewReader(SafeMathABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(SafeMathBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &SafeMath{SafeMathCaller: SafeMathCaller{contract: contract}, SafeMathTransactor: SafeMathTransactor{contract: contract}, SafeMathFilterer: SafeMathFilterer{contract: contract}}, nil
}

// SafeMath is an auto generated Go binding around an Ethereum contract.
type SafeMath struct {
	SafeMathCaller     // Read-only binding to the contract
	SafeMathTransactor // Write-only binding to the contract
	SafeMathFilterer   // Log filterer for contract events
}

// SafeMathCaller is an auto generated read-only Go binding around an Ethereum contract.
type SafeMathCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeMathTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SafeMathTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeMathFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SafeMathFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeMathSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SafeMathSession struct {
	Contract     *SafeMath         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SafeMathCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SafeMathCallerSession struct {
	Contract *SafeMathCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// SafeMathTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SafeMathTransactorSession struct {
	Contract     *SafeMathTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// SafeMathRaw is an auto generated low-level Go binding around an Ethereum contract.
type SafeMathRaw struct {
	Contract *SafeMath // Generic contract binding to access the raw methods on
}

// SafeMathCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SafeMathCallerRaw struct {
	Contract *SafeMathCaller // Generic read-only contract binding to access the raw methods on
}

// SafeMathTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SafeMathTransactorRaw struct {
	Contract *SafeMathTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSafeMath creates a new instance of SafeMath, bound to a specific deployed contract.
func NewSafeMath(address common.Address, backend bind.ContractBackend) (*SafeMath, error) {
	contract, err := bindSafeMath(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SafeMath{SafeMathCaller: SafeMathCaller{contract: contract}, SafeMathTransactor: SafeMathTransactor{contract: contract}, SafeMathFilterer: SafeMathFilterer{contract: contract}}, nil
}

// NewSafeMathCaller creates a new read-only instance of SafeMath, bound to a specific deployed contract.
func NewSafeMathCaller(address common.Address, caller bind.ContractCaller) (*SafeMathCaller, error) {
	contract, err := bindSafeMath(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SafeMathCaller{contract: contract}, nil
}

// NewSafeMathTransactor creates a new write-only instance of SafeMath, bound to a specific deployed contract.
func NewSafeMathTransactor(address common.Address, transactor bind.ContractTransactor) (*SafeMathTransactor, error) {
	contract, err := bindSafeMath(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SafeMathTransactor{contract: contract}, nil
}

// NewSafeMathFilterer creates a new log filterer instance of SafeMath, bound to a specific deployed contract.
func NewSafeMathFilterer(address common.Address, filterer bind.ContractFilterer) (*SafeMathFilterer, error) {
	contract, err := bindSafeMath(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SafeMathFilterer{contract: contract}, nil
}

// bindSafeMath binds a generic wrapper to an already deployed contract.
func bindSafeMath(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SafeMathABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SafeMath *SafeMathRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _SafeMath.Contract.SafeMathCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SafeMath *SafeMathRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SafeMath.Contract.SafeMathTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SafeMath *SafeMathRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SafeMath.Contract.SafeMathTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SafeMath *SafeMathCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _SafeMath.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SafeMath *SafeMathTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SafeMath.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SafeMath *SafeMathTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SafeMath.Contract.contract.Transact(opts, method, params...)
}

// JanetaValidatorABI is the input ABI used to generate the binding from.
const JanetaValidatorABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"_candidate\",\"type\":\"address\"}],\"name\":\"propose\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_candidate\",\"type\":\"address\"},{\"name\":\"_cap\",\"type\":\"uint256\"}],\"name\":\"unvote\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getCandidates\",\"outputs\":[{\"name\":\"\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_blockNumber\",\"type\":\"uint256\"}],\"name\":\"getWithdrawCap\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_candidate\",\"type\":\"address\"}],\"name\":\"getVoters\",\"outputs\":[{\"name\":\"\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getWithdrawBlockNumbers\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_candidate\",\"type\":\"address\"},{\"name\":\"_voter\",\"type\":\"address\"}],\"name\":\"getVoterCap\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"candidates\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_blockNumber\",\"type\":\"uint256\"},{\"name\":\"_index\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_candidate\",\"type\":\"address\"}],\"name\":\"getCandidateCap\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_candidate\",\"type\":\"address\"}],\"name\":\"vote\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"candidateCount\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"voterWithdrawDelay\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_candidate\",\"type\":\"address\"}],\"name\":\"resign\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_candidate\",\"type\":\"address\"}],\"name\":\"getCandidateOwner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"maxValidatorNumber\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"candidateWithdrawDelay\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_candidate\",\"type\":\"address\"}],\"name\":\"isCandidate\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"minCandidateCap\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"minVoterCap\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_candidates\",\"type\":\"address[]\"},{\"name\":\"_caps\",\"type\":\"uint256[]\"},{\"name\":\"_firstOwner\",\"type\":\"address\"},{\"name\":\"_minCandidateCap\",\"type\":\"uint256\"},{\"name\":\"_minVoterCap\",\"type\":\"uint256\"},{\"name\":\"_maxValidatorNumber\",\"type\":\"uint256\"},{\"name\":\"_candidateWithdrawDelay\",\"type\":\"uint256\"},{\"name\":\"_voterWithdrawDelay\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_voter\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_candidate\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_cap\",\"type\":\"uint256\"}],\"name\":\"Vote\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_voter\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_candidate\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_cap\",\"type\":\"uint256\"}],\"name\":\"Unvote\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_owner\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_candidate\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_cap\",\"type\":\"uint256\"}],\"name\":\"Propose\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_owner\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_candidate\",\"type\":\"address\"}],\"name\":\"Resign\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_owner\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_blockNumber\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"_cap\",\"type\":\"uint256\"}],\"name\":\"Withdraw\",\"type\":\"event\"}]"

// JanetaValidatorBin is the compiled bytecode used for deploying new contracts.
const JanetaValidatorBin = `0x6060604052600060045534156200001557600080fd5b60405162001415380380620014158339810160405280805182019190602001805182019190602001805191906020018051919060200180519190602001805191906020018051919060200180516005879055600686905560078590556008849055600981905591506000905088516004555060005b88518110156200027c576003805460018101620000a883826200028b565b916000526020600020900160008b8481518110620000c257fe5b90602001906020020151909190916101000a815481600160a060020a030219169083600160a060020a031602179055505060606040519081016040908152600160a060020a03891682526001602083015281018983815181106200012257fe5b906020019060200201519052600160008b84815181106200013f57fe5b90602001906020020151600160a060020a03168152602081019190915260400160002081518154600160a060020a031916600160a060020a039190911617815560208201518154901515740100000000000000000000000000000000000000000260a060020a60ff0219909116178155604082015160019091015550600260008a8381518110620001cc57fe5b90602001906020020151600160a060020a0316815260208101919091526040016000208054600181016200020183826200028b565b50600091825260208220018054600160a060020a031916600160a060020a038a16179055600554906001908b84815181106200023957fe5b90602001906020020151600160a060020a03908116825260208083019390935260409182016000908120918c16815260029091019092529020556001016200008a565b505050505050505050620002db565b815481835581811511620002b257600083815260209020620002b2918101908301620002b7565b505050565b620002d891905b80821115620002d45760008155600101620002be565b5090565b90565b61112a80620002eb6000396000f3006060604052600436106101115763ffffffff7c010000000000000000000000000000000000000000000000000000000060003504166301267951811461011657806302aa9be21461012c57806306a49fce1461014e57806315febd68146101b45780632d15cc04146101dc5780632f9c4bba146101fb578063302b68721461020e5780633477ee2e14610233578063441a3e701461026557806358e7525f1461027e5780636dd7d8ea1461029d578063a9a981a3146102b1578063a9ff959e146102c4578063ae6e43f5146102d7578063b642facd146102f6578063d09f1ab414610315578063d161c76714610328578063d51b9e931461033b578063d55b7dff1461036e578063f8ac9dd514610381575b600080fd5b61012a600160a060020a0360043516610394565b005b341561013757600080fd5b61012a600160a060020a0360043516602435610616565b341561015957600080fd5b610161610849565b60405160208082528190810183818151815260200191508051906020019060200280838360005b838110156101a0578082015183820152602001610188565b505050509050019250505060405180910390f35b34156101bf57600080fd5b6101ca6004356108b2565b60405190815260200160405180910390f35b34156101e757600080fd5b610161600160a060020a03600435166108d6565b341561020657600080fd5b610161610963565b341561021957600080fd5b6101ca600160a060020a03600435811690602435166109e5565b341561023e57600080fd5b610249600435610a14565b604051600160a060020a03909116815260200160405180910390f35b341561027057600080fd5b61012a600435602435610a3c565b341561028957600080fd5b6101ca600160a060020a0360043516610ba3565b61012a600160a060020a0360043516610bc2565b34156102bc57600080fd5b6101ca610d7f565b34156102cf57600080fd5b6101ca610d85565b34156102e257600080fd5b61012a600160a060020a0360043516610d8b565b341561030157600080fd5b610249600160a060020a0360043516611022565b341561032057600080fd5b6101ca611040565b341561033357600080fd5b6101ca611046565b341561034657600080fd5b61035a600160a060020a036004351661104c565b604051901515815260200160405180910390f35b341561037957600080fd5b6101ca611071565b341561038c57600080fd5b6101ca611077565b6005546000903410156103a657600080fd5b600160a060020a038216600090815260016020526040902054829060a060020a900460ff16156103d557600080fd5b600160a060020a03831660009081526001602081905260409091200154610402903463ffffffff61107d16565b91506003805480600101828161041891906110a5565b506000918252602090912001805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a03851617905560606040519081016040908152600160a060020a0333811683526001602080850182905283850187905291871660009081529152208151815473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a03919091161781556020820151815490151560a060020a0274ff0000000000000000000000000000000000000000199091161781556040820151600191820155600160a060020a03808616600090815260209283526040808220339093168252600290920190925290205461051d91503463ffffffff61107d16565b600160a060020a038085166000908152600160208181526040808420339095168452600290940190529190209190915560045461055f9163ffffffff61107d16565b600455600160a060020a038316600090815260026020526040902080546001810161058a83826110a5565b506000918252602090912001805473ffffffffffffffffffffffffffffffffffffffff191633600160a060020a038116919091179091557f7635f1d87b47fba9f2b09e56eb4be75cca030e0cb179c1602ac9261d39a8f5c1908434604051600160a060020a039384168152919092166020820152604080820192909252606001905180910390a1505050565b600160a060020a0380831660009081526001602090815260408083203390941683526002909301905290812054839083908190101561065457600080fd5b600160a060020a03828116600090815260016020526040902054338216911614156106c257600554600160a060020a0380841660009081526001602090815260408083203390941683526002909301905220546106b7908363ffffffff61109316565b10156106c257600080fd5b600160a060020a038516600090815260016020819052604090912001546106ef908563ffffffff61109316565b600160a060020a038087166000908152600160208181526040808420928301959095553390931682526002019091522054610730908563ffffffff61109316565b600160a060020a03808716600090815260016020908152604080832033909416835260029093019052205560095461076e904363ffffffff61107d16565b600160a060020a0333166000908152602081815260408083208484529091529020549093506107a3908563ffffffff61107d16565b600160a060020a03331660008181526020818152604080832088845280835290832094909455918152905260019081018054909181016107e383826110a5565b5060009182526020909120018390557faa0e554f781c3c3b2be110a0557f260f11af9a8aa2c64bc1e7a31dbb21e32fa2338686604051600160a060020a039384168152919092166020820152604080820192909252606001905180910390a15050505050565b6108516110ce565b60038054806020026020016040519081016040528092919081815260200182805480156108a757602002820191906000526020600020905b8154600160a060020a03168152600190910190602001808311610889575b505050505090505b90565b33600160a060020a0316600090815260208181526040808320938352929052205490565b6108de6110ce565b6002600083600160a060020a0316600160a060020a0316815260200190815260200160002080548060200260200160405190810160405280929190818152602001828054801561095757602002820191906000526020600020905b8154600160a060020a03168152600190910190602001808311610939575b50505050509050919050565b61096b6110ce565b60008033600160a060020a0316600160a060020a031681526020019081526020016000206001018054806020026020016040519081016040528092919081815260200182805480156108a757602002820191906000526020600020905b8154815260200190600101908083116109c8575050505050905090565b600160a060020a0391821660009081526001602090815260408083209390941682526002909201909152205490565b6003805482908110610a2257fe5b600091825260209091200154600160a060020a0316905081565b60008282828211610a4c57600080fd5b4382901015610a5a57600080fd5b600160a060020a03331660009081526020818152604080832085845290915281205411610a8657600080fd5b600160a060020a0333166000908152602081905260409020600101805483919083908110610ab057fe5b60009182526020909120015414610ac657600080fd5b600160a060020a03331660008181526020818152604080832089845280835290832080549084905593835291905260010180549194509085908110610b0757fe5b6000918252602082200155600160a060020a03331683156108fc0284604051600060405180830381858888f193505050501515610b4357600080fd5b7ff279e6a1f5e320cca91135676d9cb6e44ca8a08c0b88342bcdb1144f6511b5683386856040518084600160a060020a0316600160a060020a03168152602001838152602001828152602001935050505060405180910390a15050505050565b600160a060020a03166000908152600160208190526040909120015490565b600654341015610bd157600080fd5b600160a060020a038116600090815260016020526040902054819060a060020a900460ff161515610c0157600080fd5b600160a060020a03821660009081526001602081905260409091200154610c2e903463ffffffff61107d16565b600160a060020a0380841660009081526001602081815260408084209283019590955533909316825260020190915220541515610cc057600160a060020a0382166000908152600260205260409020805460018101610c8d83826110a5565b506000918252602090912001805473ffffffffffffffffffffffffffffffffffffffff191633600160a060020a03161790555b600160a060020a038083166000908152600160209081526040808320339094168352600290930190522054610cfb903463ffffffff61107d16565b600160a060020a03808416600090815260016020908152604080832033948516845260020190915290819020929092557f66a9138482c99e9baf08860110ef332cc0c23b4a199a53593d8db0fc8f96fbfc918490349051600160a060020a039384168152919092166020820152604080820192909252606001905180910390a15050565b60045481565b60095481565b600160a060020a038181166000908152600160205260408120549091829182918591338216911614610dbc57600080fd5b600160a060020a038516600090815260016020526040902054859060a060020a900460ff161515610dec57600080fd5b600160a060020a0386166000908152600160208190526040909120805474ff000000000000000000000000000000000000000019169055600454610e359163ffffffff61109316565b600455600094505b600354851015610ebf5785600160a060020a0316600386815481101515610e6057fe5b600091825260209091200154600160a060020a03161415610eb4576003805486908110610e8957fe5b6000918252602090912001805473ffffffffffffffffffffffffffffffffffffffff19169055610ebf565b600190940193610e3d565b600160a060020a03808716600081815260016020818152604080842033909616845260028601825283205493909252908190529190910154909450610f0a908563ffffffff61109316565b600160a060020a0380881660009081526001602081815260408084209283019590955533909316825260020190915290812055600854610f50904363ffffffff61107d16565b600160a060020a033316600090815260208181526040808320848452909152902054909350610f85908563ffffffff61107d16565b600160a060020a0333166000818152602081815260408083208884528083529083209490945591815290526001908101805490918101610fc583826110a5565b5060009182526020909120018390557f4edf3e325d0063213a39f9085522994a1c44bea5f39e7d63ef61260a1e58c6d33387604051600160a060020a039283168152911660208201526040908101905180910390a1505050505050565b600160a060020a039081166000908152600160205260409020541690565b60075481565b60085481565b600160a060020a031660009081526001602052604090205460a060020a900460ff1690565b60055481565b60065481565b60008282018381101561108c57fe5b9392505050565b60008282111561109f57fe5b50900390565b8154818355818115116110c9576000838152602090206110c99181019083016110e0565b505050565b60206040519081016040526000815290565b6108af91905b808211156110fa57600081556001016110e6565b50905600a165627a7a72305820555de7c5131842a4fccb258fccd95ae1539019bb744b4253893b37fed1b3d8e90029`

// DeployJanetaValidator deploys a new Ethereum contract, binding an instance of JanetaValidator to it.
func DeployJanetaValidator(auth *bind.TransactOpts, backend bind.ContractBackend, _candidates []common.Address, _caps []*big.Int, _firstOwner common.Address, _minCandidateCap *big.Int, _minVoterCap *big.Int, _maxValidatorNumber *big.Int, _candidateWithdrawDelay *big.Int, _voterWithdrawDelay *big.Int) (common.Address, *types.Transaction, *JanetaValidator, error) {
	parsed, err := abi.JSON(strings.NewReader(JanetaValidatorABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(JanetaValidatorBin), backend, _candidates, _caps, _firstOwner, _minCandidateCap, _minVoterCap, _maxValidatorNumber, _candidateWithdrawDelay, _voterWithdrawDelay)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &JanetaValidator{JanetaValidatorCaller: JanetaValidatorCaller{contract: contract}, JanetaValidatorTransactor: JanetaValidatorTransactor{contract: contract}, JanetaValidatorFilterer: JanetaValidatorFilterer{contract: contract}}, nil
}

// JanetaValidator is an auto generated Go binding around an Ethereum contract.
type JanetaValidator struct {
	JanetaValidatorCaller     // Read-only binding to the contract
	JanetaValidatorTransactor // Write-only binding to the contract
	JanetaValidatorFilterer   // Log filterer for contract events
}

// JanetaValidatorCaller is an auto generated read-only Go binding around an Ethereum contract.
type JanetaValidatorCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// JanetaValidatorTransactor is an auto generated write-only Go binding around an Ethereum contract.
type JanetaValidatorTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// JanetaValidatorFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type JanetaValidatorFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// JanetaValidatorSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type JanetaValidatorSession struct {
	Contract     *JanetaValidator    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// JanetaValidatorCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type JanetaValidatorCallerSession struct {
	Contract *JanetaValidatorCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// JanetaValidatorTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type JanetaValidatorTransactorSession struct {
	Contract     *JanetaValidatorTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// JanetaValidatorRaw is an auto generated low-level Go binding around an Ethereum contract.
type JanetaValidatorRaw struct {
	Contract *JanetaValidator // Generic contract binding to access the raw methods on
}

// JanetaValidatorCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type JanetaValidatorCallerRaw struct {
	Contract *JanetaValidatorCaller // Generic read-only contract binding to access the raw methods on
}

// JanetaValidatorTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type JanetaValidatorTransactorRaw struct {
	Contract *JanetaValidatorTransactor // Generic write-only contract binding to access the raw methods on
}

// NewJanetaValidator creates a new instance of JanetaValidator, bound to a specific deployed contract.
func NewJanetaValidator(address common.Address, backend bind.ContractBackend) (*JanetaValidator, error) {
	contract, err := bindJanetaValidator(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &JanetaValidator{JanetaValidatorCaller: JanetaValidatorCaller{contract: contract}, JanetaValidatorTransactor: JanetaValidatorTransactor{contract: contract}, JanetaValidatorFilterer: JanetaValidatorFilterer{contract: contract}}, nil
}

// NewJanetaValidatorCaller creates a new read-only instance of JanetaValidator, bound to a specific deployed contract.
func NewJanetaValidatorCaller(address common.Address, caller bind.ContractCaller) (*JanetaValidatorCaller, error) {
	contract, err := bindJanetaValidator(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &JanetaValidatorCaller{contract: contract}, nil
}

// NewJanetaValidatorTransactor creates a new write-only instance of JanetaValidator, bound to a specific deployed contract.
func NewJanetaValidatorTransactor(address common.Address, transactor bind.ContractTransactor) (*JanetaValidatorTransactor, error) {
	contract, err := bindJanetaValidator(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &JanetaValidatorTransactor{contract: contract}, nil
}

// NewJanetaValidatorFilterer creates a new log filterer instance of JanetaValidator, bound to a specific deployed contract.
func NewJanetaValidatorFilterer(address common.Address, filterer bind.ContractFilterer) (*JanetaValidatorFilterer, error) {
	contract, err := bindJanetaValidator(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &JanetaValidatorFilterer{contract: contract}, nil
}

// bindJanetaValidator binds a generic wrapper to an already deployed contract.
func bindJanetaValidator(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(JanetaValidatorABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_JanetaValidator *JanetaValidatorRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _JanetaValidator.Contract.JanetaValidatorCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_JanetaValidator *JanetaValidatorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _JanetaValidator.Contract.JanetaValidatorTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_JanetaValidator *JanetaValidatorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _JanetaValidator.Contract.JanetaValidatorTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_JanetaValidator *JanetaValidatorCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _JanetaValidator.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_JanetaValidator *JanetaValidatorTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _JanetaValidator.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_JanetaValidator *JanetaValidatorTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _JanetaValidator.Contract.contract.Transact(opts, method, params...)
}

// CandidateCount is a free data retrieval call binding the contract method 0xa9a981a3.
//
// Solidity: function candidateCount() constant returns(uint256)
func (_JanetaValidator *JanetaValidatorCaller) CandidateCount(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _JanetaValidator.contract.Call(opts, out, "candidateCount")
	return *ret0, err
}

// CandidateCount is a free data retrieval call binding the contract method 0xa9a981a3.
//
// Solidity: function candidateCount() constant returns(uint256)
func (_JanetaValidator *JanetaValidatorSession) CandidateCount() (*big.Int, error) {
	return _JanetaValidator.Contract.CandidateCount(&_JanetaValidator.CallOpts)
}

// CandidateCount is a free data retrieval call binding the contract method 0xa9a981a3.
//
// Solidity: function candidateCount() constant returns(uint256)
func (_JanetaValidator *JanetaValidatorCallerSession) CandidateCount() (*big.Int, error) {
	return _JanetaValidator.Contract.CandidateCount(&_JanetaValidator.CallOpts)
}

// CandidateWithdrawDelay is a free data retrieval call binding the contract method 0xd161c767.
//
// Solidity: function candidateWithdrawDelay() constant returns(uint256)
func (_JanetaValidator *JanetaValidatorCaller) CandidateWithdrawDelay(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _JanetaValidator.contract.Call(opts, out, "candidateWithdrawDelay")
	return *ret0, err
}

// CandidateWithdrawDelay is a free data retrieval call binding the contract method 0xd161c767.
//
// Solidity: function candidateWithdrawDelay() constant returns(uint256)
func (_JanetaValidator *JanetaValidatorSession) CandidateWithdrawDelay() (*big.Int, error) {
	return _JanetaValidator.Contract.CandidateWithdrawDelay(&_JanetaValidator.CallOpts)
}

// CandidateWithdrawDelay is a free data retrieval call binding the contract method 0xd161c767.
//
// Solidity: function candidateWithdrawDelay() constant returns(uint256)
func (_JanetaValidator *JanetaValidatorCallerSession) CandidateWithdrawDelay() (*big.Int, error) {
	return _JanetaValidator.Contract.CandidateWithdrawDelay(&_JanetaValidator.CallOpts)
}

// Candidates is a free data retrieval call binding the contract method 0x3477ee2e.
//
// Solidity: function candidates( uint256) constant returns(address)
func (_JanetaValidator *JanetaValidatorCaller) Candidates(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _JanetaValidator.contract.Call(opts, out, "candidates", arg0)
	return *ret0, err
}

// Candidates is a free data retrieval call binding the contract method 0x3477ee2e.
//
// Solidity: function candidates( uint256) constant returns(address)
func (_JanetaValidator *JanetaValidatorSession) Candidates(arg0 *big.Int) (common.Address, error) {
	return _JanetaValidator.Contract.Candidates(&_JanetaValidator.CallOpts, arg0)
}

// Candidates is a free data retrieval call binding the contract method 0x3477ee2e.
//
// Solidity: function candidates( uint256) constant returns(address)
func (_JanetaValidator *JanetaValidatorCallerSession) Candidates(arg0 *big.Int) (common.Address, error) {
	return _JanetaValidator.Contract.Candidates(&_JanetaValidator.CallOpts, arg0)
}

// GetCandidateCap is a free data retrieval call binding the contract method 0x58e7525f.
//
// Solidity: function getCandidateCap(_candidate address) constant returns(uint256)
func (_JanetaValidator *JanetaValidatorCaller) GetCandidateCap(opts *bind.CallOpts, _candidate common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _JanetaValidator.contract.Call(opts, out, "getCandidateCap", _candidate)
	return *ret0, err
}

// GetCandidateCap is a free data retrieval call binding the contract method 0x58e7525f.
//
// Solidity: function getCandidateCap(_candidate address) constant returns(uint256)
func (_JanetaValidator *JanetaValidatorSession) GetCandidateCap(_candidate common.Address) (*big.Int, error) {
	return _JanetaValidator.Contract.GetCandidateCap(&_JanetaValidator.CallOpts, _candidate)
}

// GetCandidateCap is a free data retrieval call binding the contract method 0x58e7525f.
//
// Solidity: function getCandidateCap(_candidate address) constant returns(uint256)
func (_JanetaValidator *JanetaValidatorCallerSession) GetCandidateCap(_candidate common.Address) (*big.Int, error) {
	return _JanetaValidator.Contract.GetCandidateCap(&_JanetaValidator.CallOpts, _candidate)
}

// GetCandidateOwner is a free data retrieval call binding the contract method 0xb642facd.
//
// Solidity: function getCandidateOwner(_candidate address) constant returns(address)
func (_JanetaValidator *JanetaValidatorCaller) GetCandidateOwner(opts *bind.CallOpts, _candidate common.Address) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _JanetaValidator.contract.Call(opts, out, "getCandidateOwner", _candidate)
	return *ret0, err
}

// GetCandidateOwner is a free data retrieval call binding the contract method 0xb642facd.
//
// Solidity: function getCandidateOwner(_candidate address) constant returns(address)
func (_JanetaValidator *JanetaValidatorSession) GetCandidateOwner(_candidate common.Address) (common.Address, error) {
	return _JanetaValidator.Contract.GetCandidateOwner(&_JanetaValidator.CallOpts, _candidate)
}

// GetCandidateOwner is a free data retrieval call binding the contract method 0xb642facd.
//
// Solidity: function getCandidateOwner(_candidate address) constant returns(address)
func (_JanetaValidator *JanetaValidatorCallerSession) GetCandidateOwner(_candidate common.Address) (common.Address, error) {
	return _JanetaValidator.Contract.GetCandidateOwner(&_JanetaValidator.CallOpts, _candidate)
}

// GetCandidates is a free data retrieval call binding the contract method 0x06a49fce.
//
// Solidity: function getCandidates() constant returns(address[])
func (_JanetaValidator *JanetaValidatorCaller) GetCandidates(opts *bind.CallOpts) ([]common.Address, error) {
	var (
		ret0 = new([]common.Address)
	)
	out := ret0
	err := _JanetaValidator.contract.Call(opts, out, "getCandidates")
	return *ret0, err
}

// GetCandidates is a free data retrieval call binding the contract method 0x06a49fce.
//
// Solidity: function getCandidates() constant returns(address[])
func (_JanetaValidator *JanetaValidatorSession) GetCandidates() ([]common.Address, error) {
	return _JanetaValidator.Contract.GetCandidates(&_JanetaValidator.CallOpts)
}

// GetCandidates is a free data retrieval call binding the contract method 0x06a49fce.
//
// Solidity: function getCandidates() constant returns(address[])
func (_JanetaValidator *JanetaValidatorCallerSession) GetCandidates() ([]common.Address, error) {
	return _JanetaValidator.Contract.GetCandidates(&_JanetaValidator.CallOpts)
}

// GetVoterCap is a free data retrieval call binding the contract method 0x302b6872.
//
// Solidity: function getVoterCap(_candidate address, _voter address) constant returns(uint256)
func (_JanetaValidator *JanetaValidatorCaller) GetVoterCap(opts *bind.CallOpts, _candidate common.Address, _voter common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _JanetaValidator.contract.Call(opts, out, "getVoterCap", _candidate, _voter)
	return *ret0, err
}

// GetVoterCap is a free data retrieval call binding the contract method 0x302b6872.
//
// Solidity: function getVoterCap(_candidate address, _voter address) constant returns(uint256)
func (_JanetaValidator *JanetaValidatorSession) GetVoterCap(_candidate common.Address, _voter common.Address) (*big.Int, error) {
	return _JanetaValidator.Contract.GetVoterCap(&_JanetaValidator.CallOpts, _candidate, _voter)
}

// GetVoterCap is a free data retrieval call binding the contract method 0x302b6872.
//
// Solidity: function getVoterCap(_candidate address, _voter address) constant returns(uint256)
func (_JanetaValidator *JanetaValidatorCallerSession) GetVoterCap(_candidate common.Address, _voter common.Address) (*big.Int, error) {
	return _JanetaValidator.Contract.GetVoterCap(&_JanetaValidator.CallOpts, _candidate, _voter)
}

// GetVoters is a free data retrieval call binding the contract method 0x2d15cc04.
//
// Solidity: function getVoters(_candidate address) constant returns(address[])
func (_JanetaValidator *JanetaValidatorCaller) GetVoters(opts *bind.CallOpts, _candidate common.Address) ([]common.Address, error) {
	var (
		ret0 = new([]common.Address)
	)
	out := ret0
	err := _JanetaValidator.contract.Call(opts, out, "getVoters", _candidate)
	return *ret0, err
}

// GetVoters is a free data retrieval call binding the contract method 0x2d15cc04.
//
// Solidity: function getVoters(_candidate address) constant returns(address[])
func (_JanetaValidator *JanetaValidatorSession) GetVoters(_candidate common.Address) ([]common.Address, error) {
	return _JanetaValidator.Contract.GetVoters(&_JanetaValidator.CallOpts, _candidate)
}

// GetVoters is a free data retrieval call binding the contract method 0x2d15cc04.
//
// Solidity: function getVoters(_candidate address) constant returns(address[])
func (_JanetaValidator *JanetaValidatorCallerSession) GetVoters(_candidate common.Address) ([]common.Address, error) {
	return _JanetaValidator.Contract.GetVoters(&_JanetaValidator.CallOpts, _candidate)
}

// GetWithdrawBlockNumbers is a free data retrieval call binding the contract method 0x2f9c4bba.
//
// Solidity: function getWithdrawBlockNumbers() constant returns(uint256[])
func (_JanetaValidator *JanetaValidatorCaller) GetWithdrawBlockNumbers(opts *bind.CallOpts) ([]*big.Int, error) {
	var (
		ret0 = new([]*big.Int)
	)
	out := ret0
	err := _JanetaValidator.contract.Call(opts, out, "getWithdrawBlockNumbers")
	return *ret0, err
}

// GetWithdrawBlockNumbers is a free data retrieval call binding the contract method 0x2f9c4bba.
//
// Solidity: function getWithdrawBlockNumbers() constant returns(uint256[])
func (_JanetaValidator *JanetaValidatorSession) GetWithdrawBlockNumbers() ([]*big.Int, error) {
	return _JanetaValidator.Contract.GetWithdrawBlockNumbers(&_JanetaValidator.CallOpts)
}

// GetWithdrawBlockNumbers is a free data retrieval call binding the contract method 0x2f9c4bba.
//
// Solidity: function getWithdrawBlockNumbers() constant returns(uint256[])
func (_JanetaValidator *JanetaValidatorCallerSession) GetWithdrawBlockNumbers() ([]*big.Int, error) {
	return _JanetaValidator.Contract.GetWithdrawBlockNumbers(&_JanetaValidator.CallOpts)
}

// GetWithdrawCap is a free data retrieval call binding the contract method 0x15febd68.
//
// Solidity: function getWithdrawCap(_blockNumber uint256) constant returns(uint256)
func (_JanetaValidator *JanetaValidatorCaller) GetWithdrawCap(opts *bind.CallOpts, _blockNumber *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _JanetaValidator.contract.Call(opts, out, "getWithdrawCap", _blockNumber)
	return *ret0, err
}

// GetWithdrawCap is a free data retrieval call binding the contract method 0x15febd68.
//
// Solidity: function getWithdrawCap(_blockNumber uint256) constant returns(uint256)
func (_JanetaValidator *JanetaValidatorSession) GetWithdrawCap(_blockNumber *big.Int) (*big.Int, error) {
	return _JanetaValidator.Contract.GetWithdrawCap(&_JanetaValidator.CallOpts, _blockNumber)
}

// GetWithdrawCap is a free data retrieval call binding the contract method 0x15febd68.
//
// Solidity: function getWithdrawCap(_blockNumber uint256) constant returns(uint256)
func (_JanetaValidator *JanetaValidatorCallerSession) GetWithdrawCap(_blockNumber *big.Int) (*big.Int, error) {
	return _JanetaValidator.Contract.GetWithdrawCap(&_JanetaValidator.CallOpts, _blockNumber)
}

// IsCandidate is a free data retrieval call binding the contract method 0xd51b9e93.
//
// Solidity: function isCandidate(_candidate address) constant returns(bool)
func (_JanetaValidator *JanetaValidatorCaller) IsCandidate(opts *bind.CallOpts, _candidate common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _JanetaValidator.contract.Call(opts, out, "isCandidate", _candidate)
	return *ret0, err
}

// IsCandidate is a free data retrieval call binding the contract method 0xd51b9e93.
//
// Solidity: function isCandidate(_candidate address) constant returns(bool)
func (_JanetaValidator *JanetaValidatorSession) IsCandidate(_candidate common.Address) (bool, error) {
	return _JanetaValidator.Contract.IsCandidate(&_JanetaValidator.CallOpts, _candidate)
}

// IsCandidate is a free data retrieval call binding the contract method 0xd51b9e93.
//
// Solidity: function isCandidate(_candidate address) constant returns(bool)
func (_JanetaValidator *JanetaValidatorCallerSession) IsCandidate(_candidate common.Address) (bool, error) {
	return _JanetaValidator.Contract.IsCandidate(&_JanetaValidator.CallOpts, _candidate)
}

// MaxValidatorNumber is a free data retrieval call binding the contract method 0xd09f1ab4.
//
// Solidity: function maxValidatorNumber() constant returns(uint256)
func (_JanetaValidator *JanetaValidatorCaller) MaxValidatorNumber(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _JanetaValidator.contract.Call(opts, out, "maxValidatorNumber")
	return *ret0, err
}

// MaxValidatorNumber is a free data retrieval call binding the contract method 0xd09f1ab4.
//
// Solidity: function maxValidatorNumber() constant returns(uint256)
func (_JanetaValidator *JanetaValidatorSession) MaxValidatorNumber() (*big.Int, error) {
	return _JanetaValidator.Contract.MaxValidatorNumber(&_JanetaValidator.CallOpts)
}

// MaxValidatorNumber is a free data retrieval call binding the contract method 0xd09f1ab4.
//
// Solidity: function maxValidatorNumber() constant returns(uint256)
func (_JanetaValidator *JanetaValidatorCallerSession) MaxValidatorNumber() (*big.Int, error) {
	return _JanetaValidator.Contract.MaxValidatorNumber(&_JanetaValidator.CallOpts)
}

// MinCandidateCap is a free data retrieval call binding the contract method 0xd55b7dff.
//
// Solidity: function minCandidateCap() constant returns(uint256)
func (_JanetaValidator *JanetaValidatorCaller) MinCandidateCap(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _JanetaValidator.contract.Call(opts, out, "minCandidateCap")
	return *ret0, err
}

// MinCandidateCap is a free data retrieval call binding the contract method 0xd55b7dff.
//
// Solidity: function minCandidateCap() constant returns(uint256)
func (_JanetaValidator *JanetaValidatorSession) MinCandidateCap() (*big.Int, error) {
	return _JanetaValidator.Contract.MinCandidateCap(&_JanetaValidator.CallOpts)
}

// MinCandidateCap is a free data retrieval call binding the contract method 0xd55b7dff.
//
// Solidity: function minCandidateCap() constant returns(uint256)
func (_JanetaValidator *JanetaValidatorCallerSession) MinCandidateCap() (*big.Int, error) {
	return _JanetaValidator.Contract.MinCandidateCap(&_JanetaValidator.CallOpts)
}

// MinVoterCap is a free data retrieval call binding the contract method 0xf8ac9dd5.
//
// Solidity: function minVoterCap() constant returns(uint256)
func (_JanetaValidator *JanetaValidatorCaller) MinVoterCap(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _JanetaValidator.contract.Call(opts, out, "minVoterCap")
	return *ret0, err
}

// MinVoterCap is a free data retrieval call binding the contract method 0xf8ac9dd5.
//
// Solidity: function minVoterCap() constant returns(uint256)
func (_JanetaValidator *JanetaValidatorSession) MinVoterCap() (*big.Int, error) {
	return _JanetaValidator.Contract.MinVoterCap(&_JanetaValidator.CallOpts)
}

// MinVoterCap is a free data retrieval call binding the contract method 0xf8ac9dd5.
//
// Solidity: function minVoterCap() constant returns(uint256)
func (_JanetaValidator *JanetaValidatorCallerSession) MinVoterCap() (*big.Int, error) {
	return _JanetaValidator.Contract.MinVoterCap(&_JanetaValidator.CallOpts)
}

// VoterWithdrawDelay is a free data retrieval call binding the contract method 0xa9ff959e.
//
// Solidity: function voterWithdrawDelay() constant returns(uint256)
func (_JanetaValidator *JanetaValidatorCaller) VoterWithdrawDelay(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _JanetaValidator.contract.Call(opts, out, "voterWithdrawDelay")
	return *ret0, err
}

// VoterWithdrawDelay is a free data retrieval call binding the contract method 0xa9ff959e.
//
// Solidity: function voterWithdrawDelay() constant returns(uint256)
func (_JanetaValidator *JanetaValidatorSession) VoterWithdrawDelay() (*big.Int, error) {
	return _JanetaValidator.Contract.VoterWithdrawDelay(&_JanetaValidator.CallOpts)
}

// VoterWithdrawDelay is a free data retrieval call binding the contract method 0xa9ff959e.
//
// Solidity: function voterWithdrawDelay() constant returns(uint256)
func (_JanetaValidator *JanetaValidatorCallerSession) VoterWithdrawDelay() (*big.Int, error) {
	return _JanetaValidator.Contract.VoterWithdrawDelay(&_JanetaValidator.CallOpts)
}

// Propose is a paid mutator transaction binding the contract method 0x01267951.
//
// Solidity: function propose(_candidate address) returns()
func (_JanetaValidator *JanetaValidatorTransactor) Propose(opts *bind.TransactOpts, _candidate common.Address) (*types.Transaction, error) {
	return _JanetaValidator.contract.Transact(opts, "propose", _candidate)
}

// Propose is a paid mutator transaction binding the contract method 0x01267951.
//
// Solidity: function propose(_candidate address) returns()
func (_JanetaValidator *JanetaValidatorSession) Propose(_candidate common.Address) (*types.Transaction, error) {
	return _JanetaValidator.Contract.Propose(&_JanetaValidator.TransactOpts, _candidate)
}

// Propose is a paid mutator transaction binding the contract method 0x01267951.
//
// Solidity: function propose(_candidate address) returns()
func (_JanetaValidator *JanetaValidatorTransactorSession) Propose(_candidate common.Address) (*types.Transaction, error) {
	return _JanetaValidator.Contract.Propose(&_JanetaValidator.TransactOpts, _candidate)
}

// Resign is a paid mutator transaction binding the contract method 0xae6e43f5.
//
// Solidity: function resign(_candidate address) returns()
func (_JanetaValidator *JanetaValidatorTransactor) Resign(opts *bind.TransactOpts, _candidate common.Address) (*types.Transaction, error) {
	return _JanetaValidator.contract.Transact(opts, "resign", _candidate)
}

// Resign is a paid mutator transaction binding the contract method 0xae6e43f5.
//
// Solidity: function resign(_candidate address) returns()
func (_JanetaValidator *JanetaValidatorSession) Resign(_candidate common.Address) (*types.Transaction, error) {
	return _JanetaValidator.Contract.Resign(&_JanetaValidator.TransactOpts, _candidate)
}

// Resign is a paid mutator transaction binding the contract method 0xae6e43f5.
//
// Solidity: function resign(_candidate address) returns()
func (_JanetaValidator *JanetaValidatorTransactorSession) Resign(_candidate common.Address) (*types.Transaction, error) {
	return _JanetaValidator.Contract.Resign(&_JanetaValidator.TransactOpts, _candidate)
}

// Unvote is a paid mutator transaction binding the contract method 0x02aa9be2.
//
// Solidity: function unvote(_candidate address, _cap uint256) returns()
func (_JanetaValidator *JanetaValidatorTransactor) Unvote(opts *bind.TransactOpts, _candidate common.Address, _cap *big.Int) (*types.Transaction, error) {
	return _JanetaValidator.contract.Transact(opts, "unvote", _candidate, _cap)
}

// Unvote is a paid mutator transaction binding the contract method 0x02aa9be2.
//
// Solidity: function unvote(_candidate address, _cap uint256) returns()
func (_JanetaValidator *JanetaValidatorSession) Unvote(_candidate common.Address, _cap *big.Int) (*types.Transaction, error) {
	return _JanetaValidator.Contract.Unvote(&_JanetaValidator.TransactOpts, _candidate, _cap)
}

// Unvote is a paid mutator transaction binding the contract method 0x02aa9be2.
//
// Solidity: function unvote(_candidate address, _cap uint256) returns()
func (_JanetaValidator *JanetaValidatorTransactorSession) Unvote(_candidate common.Address, _cap *big.Int) (*types.Transaction, error) {
	return _JanetaValidator.Contract.Unvote(&_JanetaValidator.TransactOpts, _candidate, _cap)
}

// Vote is a paid mutator transaction binding the contract method 0x6dd7d8ea.
//
// Solidity: function vote(_candidate address) returns()
func (_JanetaValidator *JanetaValidatorTransactor) Vote(opts *bind.TransactOpts, _candidate common.Address) (*types.Transaction, error) {
	return _JanetaValidator.contract.Transact(opts, "vote", _candidate)
}

// Vote is a paid mutator transaction binding the contract method 0x6dd7d8ea.
//
// Solidity: function vote(_candidate address) returns()
func (_JanetaValidator *JanetaValidatorSession) Vote(_candidate common.Address) (*types.Transaction, error) {
	return _JanetaValidator.Contract.Vote(&_JanetaValidator.TransactOpts, _candidate)
}

// Vote is a paid mutator transaction binding the contract method 0x6dd7d8ea.
//
// Solidity: function vote(_candidate address) returns()
func (_JanetaValidator *JanetaValidatorTransactorSession) Vote(_candidate common.Address) (*types.Transaction, error) {
	return _JanetaValidator.Contract.Vote(&_JanetaValidator.TransactOpts, _candidate)
}

// Withdraw is a paid mutator transaction binding the contract method 0x441a3e70.
//
// Solidity: function withdraw(_blockNumber uint256, _index uint256) returns()
func (_JanetaValidator *JanetaValidatorTransactor) Withdraw(opts *bind.TransactOpts, _blockNumber *big.Int, _index *big.Int) (*types.Transaction, error) {
	return _JanetaValidator.contract.Transact(opts, "withdraw", _blockNumber, _index)
}

// Withdraw is a paid mutator transaction binding the contract method 0x441a3e70.
//
// Solidity: function withdraw(_blockNumber uint256, _index uint256) returns()
func (_JanetaValidator *JanetaValidatorSession) Withdraw(_blockNumber *big.Int, _index *big.Int) (*types.Transaction, error) {
	return _JanetaValidator.Contract.Withdraw(&_JanetaValidator.TransactOpts, _blockNumber, _index)
}

// Withdraw is a paid mutator transaction binding the contract method 0x441a3e70.
//
// Solidity: function withdraw(_blockNumber uint256, _index uint256) returns()
func (_JanetaValidator *JanetaValidatorTransactorSession) Withdraw(_blockNumber *big.Int, _index *big.Int) (*types.Transaction, error) {
	return _JanetaValidator.Contract.Withdraw(&_JanetaValidator.TransactOpts, _blockNumber, _index)
}

// JanetaValidatorProposeIterator is returned from FilterPropose and is used to iterate over the raw logs and unpacked data for Propose events raised by the JanetaValidator contract.
type JanetaValidatorProposeIterator struct {
	Event *JanetaValidatorPropose // Event containing the contract specifics and raw log

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
func (it *JanetaValidatorProposeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(JanetaValidatorPropose)
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
		it.Event = new(JanetaValidatorPropose)
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
func (it *JanetaValidatorProposeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *JanetaValidatorProposeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// JanetaValidatorPropose represents a Propose event raised by the JanetaValidator contract.
type JanetaValidatorPropose struct {
	Owner     common.Address
	Candidate common.Address
	Cap       *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterPropose is a free log retrieval operation binding the contract event 0x7635f1d87b47fba9f2b09e56eb4be75cca030e0cb179c1602ac9261d39a8f5c1.
//
// Solidity: event Propose(_owner address, _candidate address, _cap uint256)
func (_JanetaValidator *JanetaValidatorFilterer) FilterPropose(opts *bind.FilterOpts) (*JanetaValidatorProposeIterator, error) {

	logs, sub, err := _JanetaValidator.contract.FilterLogs(opts, "Propose")
	if err != nil {
		return nil, err
	}
	return &JanetaValidatorProposeIterator{contract: _JanetaValidator.contract, event: "Propose", logs: logs, sub: sub}, nil
}

// WatchPropose is a free log subscription operation binding the contract event 0x7635f1d87b47fba9f2b09e56eb4be75cca030e0cb179c1602ac9261d39a8f5c1.
//
// Solidity: event Propose(_owner address, _candidate address, _cap uint256)
func (_JanetaValidator *JanetaValidatorFilterer) WatchPropose(opts *bind.WatchOpts, sink chan<- *JanetaValidatorPropose) (event.Subscription, error) {

	logs, sub, err := _JanetaValidator.contract.WatchLogs(opts, "Propose")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(JanetaValidatorPropose)
				if err := _JanetaValidator.contract.UnpackLog(event, "Propose", log); err != nil {
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

// JanetaValidatorResignIterator is returned from FilterResign and is used to iterate over the raw logs and unpacked data for Resign events raised by the JanetaValidator contract.
type JanetaValidatorResignIterator struct {
	Event *JanetaValidatorResign // Event containing the contract specifics and raw log

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
func (it *JanetaValidatorResignIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(JanetaValidatorResign)
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
		it.Event = new(JanetaValidatorResign)
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
func (it *JanetaValidatorResignIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *JanetaValidatorResignIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// JanetaValidatorResign represents a Resign event raised by the JanetaValidator contract.
type JanetaValidatorResign struct {
	Owner     common.Address
	Candidate common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterResign is a free log retrieval operation binding the contract event 0x4edf3e325d0063213a39f9085522994a1c44bea5f39e7d63ef61260a1e58c6d3.
//
// Solidity: event Resign(_owner address, _candidate address)
func (_JanetaValidator *JanetaValidatorFilterer) FilterResign(opts *bind.FilterOpts) (*JanetaValidatorResignIterator, error) {

	logs, sub, err := _JanetaValidator.contract.FilterLogs(opts, "Resign")
	if err != nil {
		return nil, err
	}
	return &JanetaValidatorResignIterator{contract: _JanetaValidator.contract, event: "Resign", logs: logs, sub: sub}, nil
}

// WatchResign is a free log subscription operation binding the contract event 0x4edf3e325d0063213a39f9085522994a1c44bea5f39e7d63ef61260a1e58c6d3.
//
// Solidity: event Resign(_owner address, _candidate address)
func (_JanetaValidator *JanetaValidatorFilterer) WatchResign(opts *bind.WatchOpts, sink chan<- *JanetaValidatorResign) (event.Subscription, error) {

	logs, sub, err := _JanetaValidator.contract.WatchLogs(opts, "Resign")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(JanetaValidatorResign)
				if err := _JanetaValidator.contract.UnpackLog(event, "Resign", log); err != nil {
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

// JanetaValidatorUnvoteIterator is returned from FilterUnvote and is used to iterate over the raw logs and unpacked data for Unvote events raised by the JanetaValidator contract.
type JanetaValidatorUnvoteIterator struct {
	Event *JanetaValidatorUnvote // Event containing the contract specifics and raw log

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
func (it *JanetaValidatorUnvoteIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(JanetaValidatorUnvote)
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
		it.Event = new(JanetaValidatorUnvote)
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
func (it *JanetaValidatorUnvoteIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *JanetaValidatorUnvoteIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// JanetaValidatorUnvote represents a Unvote event raised by the JanetaValidator contract.
type JanetaValidatorUnvote struct {
	Voter     common.Address
	Candidate common.Address
	Cap       *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterUnvote is a free log retrieval operation binding the contract event 0xaa0e554f781c3c3b2be110a0557f260f11af9a8aa2c64bc1e7a31dbb21e32fa2.
//
// Solidity: event Unvote(_voter address, _candidate address, _cap uint256)
func (_JanetaValidator *JanetaValidatorFilterer) FilterUnvote(opts *bind.FilterOpts) (*JanetaValidatorUnvoteIterator, error) {

	logs, sub, err := _JanetaValidator.contract.FilterLogs(opts, "Unvote")
	if err != nil {
		return nil, err
	}
	return &JanetaValidatorUnvoteIterator{contract: _JanetaValidator.contract, event: "Unvote", logs: logs, sub: sub}, nil
}

// WatchUnvote is a free log subscription operation binding the contract event 0xaa0e554f781c3c3b2be110a0557f260f11af9a8aa2c64bc1e7a31dbb21e32fa2.
//
// Solidity: event Unvote(_voter address, _candidate address, _cap uint256)
func (_JanetaValidator *JanetaValidatorFilterer) WatchUnvote(opts *bind.WatchOpts, sink chan<- *JanetaValidatorUnvote) (event.Subscription, error) {

	logs, sub, err := _JanetaValidator.contract.WatchLogs(opts, "Unvote")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(JanetaValidatorUnvote)
				if err := _JanetaValidator.contract.UnpackLog(event, "Unvote", log); err != nil {
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

// JanetaValidatorVoteIterator is returned from FilterVote and is used to iterate over the raw logs and unpacked data for Vote events raised by the JanetaValidator contract.
type JanetaValidatorVoteIterator struct {
	Event *JanetaValidatorVote // Event containing the contract specifics and raw log

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
func (it *JanetaValidatorVoteIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(JanetaValidatorVote)
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
		it.Event = new(JanetaValidatorVote)
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
func (it *JanetaValidatorVoteIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *JanetaValidatorVoteIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// JanetaValidatorVote represents a Vote event raised by the JanetaValidator contract.
type JanetaValidatorVote struct {
	Voter     common.Address
	Candidate common.Address
	Cap       *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterVote is a free log retrieval operation binding the contract event 0x66a9138482c99e9baf08860110ef332cc0c23b4a199a53593d8db0fc8f96fbfc.
//
// Solidity: event Vote(_voter address, _candidate address, _cap uint256)
func (_JanetaValidator *JanetaValidatorFilterer) FilterVote(opts *bind.FilterOpts) (*JanetaValidatorVoteIterator, error) {

	logs, sub, err := _JanetaValidator.contract.FilterLogs(opts, "Vote")
	if err != nil {
		return nil, err
	}
	return &JanetaValidatorVoteIterator{contract: _JanetaValidator.contract, event: "Vote", logs: logs, sub: sub}, nil
}

// WatchVote is a free log subscription operation binding the contract event 0x66a9138482c99e9baf08860110ef332cc0c23b4a199a53593d8db0fc8f96fbfc.
//
// Solidity: event Vote(_voter address, _candidate address, _cap uint256)
func (_JanetaValidator *JanetaValidatorFilterer) WatchVote(opts *bind.WatchOpts, sink chan<- *JanetaValidatorVote) (event.Subscription, error) {

	logs, sub, err := _JanetaValidator.contract.WatchLogs(opts, "Vote")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(JanetaValidatorVote)
				if err := _JanetaValidator.contract.UnpackLog(event, "Vote", log); err != nil {
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

// JanetaValidatorWithdrawIterator is returned from FilterWithdraw and is used to iterate over the raw logs and unpacked data for Withdraw events raised by the JanetaValidator contract.
type JanetaValidatorWithdrawIterator struct {
	Event *JanetaValidatorWithdraw // Event containing the contract specifics and raw log

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
func (it *JanetaValidatorWithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(JanetaValidatorWithdraw)
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
		it.Event = new(JanetaValidatorWithdraw)
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
func (it *JanetaValidatorWithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *JanetaValidatorWithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// JanetaValidatorWithdraw represents a Withdraw event raised by the JanetaValidator contract.
type JanetaValidatorWithdraw struct {
	Owner       common.Address
	BlockNumber *big.Int
	Cap         *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterWithdraw is a free log retrieval operation binding the contract event 0xf279e6a1f5e320cca91135676d9cb6e44ca8a08c0b88342bcdb1144f6511b568.
//
// Solidity: event Withdraw(_owner address, _blockNumber uint256, _cap uint256)
func (_JanetaValidator *JanetaValidatorFilterer) FilterWithdraw(opts *bind.FilterOpts) (*JanetaValidatorWithdrawIterator, error) {

	logs, sub, err := _JanetaValidator.contract.FilterLogs(opts, "Withdraw")
	if err != nil {
		return nil, err
	}
	return &JanetaValidatorWithdrawIterator{contract: _JanetaValidator.contract, event: "Withdraw", logs: logs, sub: sub}, nil
}

// WatchWithdraw is a free log subscription operation binding the contract event 0xf279e6a1f5e320cca91135676d9cb6e44ca8a08c0b88342bcdb1144f6511b568.
//
// Solidity: event Withdraw(_owner address, _blockNumber uint256, _cap uint256)
func (_JanetaValidator *JanetaValidatorFilterer) WatchWithdraw(opts *bind.WatchOpts, sink chan<- *JanetaValidatorWithdraw) (event.Subscription, error) {

	logs, sub, err := _JanetaValidator.contract.WatchLogs(opts, "Withdraw")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(JanetaValidatorWithdraw)
				if err := _JanetaValidator.contract.UnpackLog(event, "Withdraw", log); err != nil {
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
