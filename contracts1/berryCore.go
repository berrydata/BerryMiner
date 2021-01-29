// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contracts1

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

// SafeMathABI is the input ABI used to generate the binding from.
const SafeMathABI = "[]"

// SafeMathBin is the compiled bytecode used for deploying new contracts.
const SafeMathBin = `0x60556023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea265627a7a72305820b0b25139b75e676b66895a7d1fb0a1f523d8c20a178651d70497b0640c4406c364736f6c634300050a0032`

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

// BerryABI is the input ABI used to generate the binding from.
const BerryABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"_spender\",\"type\":\"address\"},{\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"depositStake\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_from\",\"type\":\"address\"},{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_propNewBerryAddress\",\"type\":\"address\"}],\"name\":\"proposeFork\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"requestStakingWithdraw\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_c_sapi\",\"type\":\"string\"},{\"name\":\"_c_symbol\",\"type\":\"string\"},{\"name\":\"_granularity\",\"type\":\"uint256\"},{\"name\":\"_tip\",\"type\":\"uint256\"}],\"name\":\"requestData\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_disputeId\",\"type\":\"uint256\"}],\"name\":\"tallyVotes\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_nonce\",\"type\":\"string\"},{\"name\":\"_requestId\",\"type\":\"uint256\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"submitMiningSolution\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_requestId\",\"type\":\"uint256\"},{\"name\":\"_tip\",\"type\":\"uint256\"}],\"name\":\"addTip\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_requestId\",\"type\":\"uint256\"},{\"name\":\"_timestamp\",\"type\":\"uint256\"},{\"name\":\"_minerIndex\",\"type\":\"uint256\"}],\"name\":\"beginDispute\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_address\",\"type\":\"address\"},{\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"theLazyCoon\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"withdrawStake\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_disputeId\",\"type\":\"uint256\"},{\"name\":\"_supportsDispute\",\"type\":\"bool\"}],\"name\":\"vote\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// BerryBin is the compiled bytecode used for deploying new contracts.
const BerryBin = `0x608060405234801561001057600080fd5b50610c0b806100206000396000f3fe608060405234801561001057600080fd5b50600436106100f55760003560e01c806368c180d511610097578063b079f64a11610066578063b079f64a1461037b578063bed9d861146103a7578063c9d27afe146103af578063f2fde38b146103d4576100f5565b806368c180d51461028d578063752d49a1146103035780638581af1914610326578063a9059cbb1461034f576100f5565b806326b7d9f6116100d357806326b7d9f61461017a57806328449c3a146101a05780633fff2816146101a85780634d318b0e14610270576100f5565b8063095ea7b3146100fa5780630d2d76a21461013a57806323b872dd14610144575b600080fd5b6101266004803603604081101561011057600080fd5b506001600160a01b0381351690602001356103fa565b604080519115158252519081900360200190f35b610142610497565b005b6101266004803603606081101561015a57600080fd5b506001600160a01b03813581169160208101359091169060400135610501565b6101426004803603602081101561019057600080fd5b50356001600160a01b03166105a7565b610142610621565b610142600480360360808110156101be57600080fd5b8101906020810181356401000000008111156101d957600080fd5b8201836020820111156101eb57600080fd5b8035906020019184600183028401116401000000008311171561020d57600080fd5b91939092909160208101903564010000000081111561022b57600080fd5b82018360208201111561023d57600080fd5b8035906020019184600183028401116401000000008311171561025f57600080fd5b919350915080359060200135610671565b6101426004803603602081101561028657600080fd5b5035610753565b610142600480360360608110156102a357600080fd5b8101906020810181356401000000008111156102be57600080fd5b8201836020820111156102d057600080fd5b803590602001918460018302840111640100000000831117156102f257600080fd5b9193509150803590602001356107aa565b6101426004803603604081101561031957600080fd5b508035906020013561085e565b6101426004803603606081101561033c57600080fd5b50803590602081013590604001356108d8565b6101266004803603604081101561036557600080fd5b506001600160a01b03813516906020013561095a565b6101426004803603604081101561039157600080fd5b506001600160a01b0381351690602001356109c4565b610142610a2a565b610142600480360360408110156103c557600080fd5b50803590602001351515610a7a565b610142600480360360208110156103ea57600080fd5b50356001600160a01b0316610ad9565b60408051634286e61960e11b81526000600482018190526001600160a01b038516602483015260448201849052915173__$ff8037bb7c49d17e1e79a1c3f1e9fdeb7a$__9163850dcc32916064808301926020929190829003018186803b15801561046457600080fd5b505af4158015610478573d6000803e3d6000fd5b505050506040513d602081101561048e57600080fd5b50519392505050565b6040805163410516b360e11b8152600060048201819052915173__$799602413129f49037f52758954cf5aa52$__9263820a2d669260248082019391829003018186803b1580156104e757600080fd5b505af41580156104fb573d6000803e3d6000fd5b50505050565b6040805163ca50189960e01b81526000600482018190526001600160a01b0380871660248401528516604483015260648201849052915173__$ff8037bb7c49d17e1e79a1c3f1e9fdeb7a$__9163ca501899916084808301926020929190829003018186803b15801561057357600080fd5b505af4158015610587573d6000803e3d6000fd5b505050506040513d602081101561059d57600080fd5b5051949350505050565b6040805163694bf49f60e01b81526000600482018190526001600160a01b0384166024830152915173__$5c50ea773c3f1223822c80c8edbce14ea7$__9263694bf49f9260448082019391829003018186803b15801561060657600080fd5b505af415801561061a573d6000803e3d6000fd5b5050505050565b60408051633273d79360e21b8152600060048201819052915173__$799602413129f49037f52758954cf5aa52$__9263c9cf5e4c9260248082019391829003018186803b1580156104e757600080fd5b600073__$d2345629a19a8ba674464680be5d957238$__6385997bf690918888888888886040518863ffffffff1660e01b81526004018088815260200180602001806020018581526020018481526020018381038352898982818152602001925080828437600083820152601f01601f191690910184810383528781526020019050878780828437600081840152601f19601f820116905080830192505050995050505050505050505060006040518083038186803b15801561073357600080fd5b505af4158015610747573d6000803e3d6000fd5b50505050505050505050565b6040805163def6fac760e01b815260006004820181905260248201849052915173__$5c50ea773c3f1223822c80c8edbce14ea7$__9263def6fac79260448082019391829003018186803b15801561060657600080fd5b600073__$d2345629a19a8ba674464680be5d957238$__63a098b5b49091868686866040518663ffffffff1660e01b815260040180868152602001806020018481526020018381526020018281038252868682818152602001925080828437600081840152601f19601f820116905080830192505050965050505050505060006040518083038186803b15801561084057600080fd5b505af4158015610854573d6000803e3d6000fd5b5050505050505050565b604080516302e8f21b60e01b81526000600482018190526024820185905260448201849052915173__$d2345629a19a8ba674464680be5d957238$__926302e8f21b9260648082019391829003018186803b1580156108bc57600080fd5b505af41580156108d0573d6000803e3d6000fd5b505050505050565b6040805163ca9a4ea560e01b8152600060048201819052602482018690526044820185905260648201849052915173__$5c50ea773c3f1223822c80c8edbce14ea7$__9263ca9a4ea59260848082019391829003018186803b15801561093d57600080fd5b505af4158015610951573d6000803e3d6000fd5b50505050505050565b6040805163c84b96f560e01b81526000600482018190526001600160a01b038516602483015260448201849052915173__$ff8037bb7c49d17e1e79a1c3f1e9fdeb7a$__9163c84b96f5916064808301926020929190829003018186803b15801561046457600080fd5b60408051632432b5d360e11b81526000600482018190526001600160a01b038516602483015260448201849052915173__$d2345629a19a8ba674464680be5d957238$__926348656ba69260648082019391829003018186803b1580156108bc57600080fd5b604080516344bacc4b60e01b8152600060048201819052915173__$799602413129f49037f52758954cf5aa52$__926344bacc4b9260248082019391829003018186803b1580156104e757600080fd5b604080516316d0383760e11b8152600060048201819052602482018590528315156044830152915173__$5c50ea773c3f1223822c80c8edbce14ea7$__92632da0706e9260648082019391829003018186803b1580156108bc57600080fd5b610aea60008263ffffffff610aed16565b50565b60408051652fb7bbb732b960d11b815281519081900360060190206000908152603f84016020522054336001600160a01b0390911614610b2c57600080fd5b60408051652fb7bbb732b960d11b815281519081900360060181206000908152603f8501602052918220546001600160a01b03808516939116917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09190a360408051652fb7bbb732b960d11b815281519081900360060190206000908152603f90930160205290912080546001600160a01b039092166001600160a01b031990921691909117905556fea265627a7a723058204ed69f8834274910d3d9493dcb33431ec6bbda4897487775254418b8ea1efcc464736f6c634300050a0032`

// DeployBerry deploys a new Ethereum contract, binding an instance of Berry to it.
func DeployBerry(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Berry, error) {
	parsed, err := abi.JSON(strings.NewReader(BerryABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(BerryBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Berry{BerryCaller: BerryCaller{contract: contract}, BerryTransactor: BerryTransactor{contract: contract}, BerryFilterer: BerryFilterer{contract: contract}}, nil
}

// Berry is an auto generated Go binding around an Ethereum contract.
type Berry struct {
	BerryCaller     // Read-only binding to the contract
	BerryTransactor // Write-only binding to the contract
	BerryFilterer   // Log filterer for contract events
}

// BerryCaller is an auto generated read-only Go binding around an Ethereum contract.
type BerryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BerryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BerryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BerryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BerryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BerrySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BerrySession struct {
	Contract     *Berry           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BerryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BerryCallerSession struct {
	Contract *BerryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// BerryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BerryTransactorSession struct {
	Contract     *BerryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BerryRaw is an auto generated low-level Go binding around an Ethereum contract.
type BerryRaw struct {
	Contract *Berry // Generic contract binding to access the raw methods on
}

// BerryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BerryCallerRaw struct {
	Contract *BerryCaller // Generic read-only contract binding to access the raw methods on
}

// BerryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BerryTransactorRaw struct {
	Contract *BerryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBerry creates a new instance of Berry, bound to a specific deployed contract.
func NewBerry(address common.Address, backend bind.ContractBackend) (*Berry, error) {
	contract, err := bindBerry(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Berry{BerryCaller: BerryCaller{contract: contract}, BerryTransactor: BerryTransactor{contract: contract}, BerryFilterer: BerryFilterer{contract: contract}}, nil
}

// NewBerryCaller creates a new read-only instance of Berry, bound to a specific deployed contract.
func NewBerryCaller(address common.Address, caller bind.ContractCaller) (*BerryCaller, error) {
	contract, err := bindBerry(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BerryCaller{contract: contract}, nil
}

// NewBerryTransactor creates a new write-only instance of Berry, bound to a specific deployed contract.
func NewBerryTransactor(address common.Address, transactor bind.ContractTransactor) (*BerryTransactor, error) {
	contract, err := bindBerry(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BerryTransactor{contract: contract}, nil
}

// NewBerryFilterer creates a new log filterer instance of Berry, bound to a specific deployed contract.
func NewBerryFilterer(address common.Address, filterer bind.ContractFilterer) (*BerryFilterer, error) {
	contract, err := bindBerry(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BerryFilterer{contract: contract}, nil
}

// bindBerry binds a generic wrapper to an already deployed contract.
func bindBerry(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BerryABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Berry *BerryRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Berry.Contract.BerryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Berry *BerryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Berry.Contract.BerryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Berry *BerryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Berry.Contract.BerryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Berry *BerryCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Berry.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Berry *BerryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Berry.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Berry *BerryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Berry.Contract.contract.Transact(opts, method, params...)
}

// AddTip is a paid mutator transaction binding the contract method 0x752d49a1.
//
// Solidity: function addTip(_requestId uint256, _tip uint256) returns()
func (_Berry *BerryTransactor) AddTip(opts *bind.TransactOpts, _requestId *big.Int, _tip *big.Int) (*types.Transaction, error) {
	return _Berry.contract.Transact(opts, "addTip", _requestId, _tip)
}

// AddTip is a paid mutator transaction binding the contract method 0x752d49a1.
//
// Solidity: function addTip(_requestId uint256, _tip uint256) returns()
func (_Berry *BerrySession) AddTip(_requestId *big.Int, _tip *big.Int) (*types.Transaction, error) {
	return _Berry.Contract.AddTip(&_Berry.TransactOpts, _requestId, _tip)
}

// AddTip is a paid mutator transaction binding the contract method 0x752d49a1.
//
// Solidity: function addTip(_requestId uint256, _tip uint256) returns()
func (_Berry *BerryTransactorSession) AddTip(_requestId *big.Int, _tip *big.Int) (*types.Transaction, error) {
	return _Berry.Contract.AddTip(&_Berry.TransactOpts, _requestId, _tip)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(_spender address, _amount uint256) returns(bool)
func (_Berry *BerryTransactor) Approve(opts *bind.TransactOpts, _spender common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Berry.contract.Transact(opts, "approve", _spender, _amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(_spender address, _amount uint256) returns(bool)
func (_Berry *BerrySession) Approve(_spender common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Berry.Contract.Approve(&_Berry.TransactOpts, _spender, _amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(_spender address, _amount uint256) returns(bool)
func (_Berry *BerryTransactorSession) Approve(_spender common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Berry.Contract.Approve(&_Berry.TransactOpts, _spender, _amount)
}

// BeginDispute is a paid mutator transaction binding the contract method 0x8581af19.
//
// Solidity: function beginDispute(_requestId uint256, _timestamp uint256, _minerIndex uint256) returns()
func (_Berry *BerryTransactor) BeginDispute(opts *bind.TransactOpts, _requestId *big.Int, _timestamp *big.Int, _minerIndex *big.Int) (*types.Transaction, error) {
	return _Berry.contract.Transact(opts, "beginDispute", _requestId, _timestamp, _minerIndex)
}

// BeginDispute is a paid mutator transaction binding the contract method 0x8581af19.
//
// Solidity: function beginDispute(_requestId uint256, _timestamp uint256, _minerIndex uint256) returns()
func (_Berry *BerrySession) BeginDispute(_requestId *big.Int, _timestamp *big.Int, _minerIndex *big.Int) (*types.Transaction, error) {
	return _Berry.Contract.BeginDispute(&_Berry.TransactOpts, _requestId, _timestamp, _minerIndex)
}

// BeginDispute is a paid mutator transaction binding the contract method 0x8581af19.
//
// Solidity: function beginDispute(_requestId uint256, _timestamp uint256, _minerIndex uint256) returns()
func (_Berry *BerryTransactorSession) BeginDispute(_requestId *big.Int, _timestamp *big.Int, _minerIndex *big.Int) (*types.Transaction, error) {
	return _Berry.Contract.BeginDispute(&_Berry.TransactOpts, _requestId, _timestamp, _minerIndex)
}

// DepositStake is a paid mutator transaction binding the contract method 0x0d2d76a2.
//
// Solidity: function depositStake() returns()
func (_Berry *BerryTransactor) DepositStake(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Berry.contract.Transact(opts, "depositStake")
}

// DepositStake is a paid mutator transaction binding the contract method 0x0d2d76a2.
//
// Solidity: function depositStake() returns()
func (_Berry *BerrySession) DepositStake() (*types.Transaction, error) {
	return _Berry.Contract.DepositStake(&_Berry.TransactOpts)
}

// DepositStake is a paid mutator transaction binding the contract method 0x0d2d76a2.
//
// Solidity: function depositStake() returns()
func (_Berry *BerryTransactorSession) DepositStake() (*types.Transaction, error) {
	return _Berry.Contract.DepositStake(&_Berry.TransactOpts)
}

// ProposeFork is a paid mutator transaction binding the contract method 0x26b7d9f6.
//
// Solidity: function proposeFork(_propNewBerryAddress address) returns()
func (_Berry *BerryTransactor) ProposeFork(opts *bind.TransactOpts, _propNewBerryAddress common.Address) (*types.Transaction, error) {
	return _Berry.contract.Transact(opts, "proposeFork", _propNewBerryAddress)
}

// ProposeFork is a paid mutator transaction binding the contract method 0x26b7d9f6.
//
// Solidity: function proposeFork(_propNewBerryAddress address) returns()
func (_Berry *BerrySession) ProposeFork(_propNewBerryAddress common.Address) (*types.Transaction, error) {
	return _Berry.Contract.ProposeFork(&_Berry.TransactOpts, _propNewBerryAddress)
}

// ProposeFork is a paid mutator transaction binding the contract method 0x26b7d9f6.
//
// Solidity: function proposeFork(_propNewBerryAddress address) returns()
func (_Berry *BerryTransactorSession) ProposeFork(_propNewBerryAddress common.Address) (*types.Transaction, error) {
	return _Berry.Contract.ProposeFork(&_Berry.TransactOpts, _propNewBerryAddress)
}

// RequestData is a paid mutator transaction binding the contract method 0x3fff2816.
//
// Solidity: function requestData(_c_sapi string, _c_symbol string, _granularity uint256, _tip uint256) returns()
func (_Berry *BerryTransactor) RequestData(opts *bind.TransactOpts, _c_sapi string, _c_symbol string, _granularity *big.Int, _tip *big.Int) (*types.Transaction, error) {
	return _Berry.contract.Transact(opts, "requestData", _c_sapi, _c_symbol, _granularity, _tip)
}

// RequestData is a paid mutator transaction binding the contract method 0x3fff2816.
//
// Solidity: function requestData(_c_sapi string, _c_symbol string, _granularity uint256, _tip uint256) returns()
func (_Berry *BerrySession) RequestData(_c_sapi string, _c_symbol string, _granularity *big.Int, _tip *big.Int) (*types.Transaction, error) {
	return _Berry.Contract.RequestData(&_Berry.TransactOpts, _c_sapi, _c_symbol, _granularity, _tip)
}

// RequestData is a paid mutator transaction binding the contract method 0x3fff2816.
//
// Solidity: function requestData(_c_sapi string, _c_symbol string, _granularity uint256, _tip uint256) returns()
func (_Berry *BerryTransactorSession) RequestData(_c_sapi string, _c_symbol string, _granularity *big.Int, _tip *big.Int) (*types.Transaction, error) {
	return _Berry.Contract.RequestData(&_Berry.TransactOpts, _c_sapi, _c_symbol, _granularity, _tip)
}

// RequestStakingWithdraw is a paid mutator transaction binding the contract method 0x28449c3a.
//
// Solidity: function requestStakingWithdraw() returns()
func (_Berry *BerryTransactor) RequestStakingWithdraw(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Berry.contract.Transact(opts, "requestStakingWithdraw")
}

// RequestStakingWithdraw is a paid mutator transaction binding the contract method 0x28449c3a.
//
// Solidity: function requestStakingWithdraw() returns()
func (_Berry *BerrySession) RequestStakingWithdraw() (*types.Transaction, error) {
	return _Berry.Contract.RequestStakingWithdraw(&_Berry.TransactOpts)
}

// RequestStakingWithdraw is a paid mutator transaction binding the contract method 0x28449c3a.
//
// Solidity: function requestStakingWithdraw() returns()
func (_Berry *BerryTransactorSession) RequestStakingWithdraw() (*types.Transaction, error) {
	return _Berry.Contract.RequestStakingWithdraw(&_Berry.TransactOpts)
}

// SubmitMiningSolution is a paid mutator transaction binding the contract method 0x68c180d5.
//
// Solidity: function submitMiningSolution(_nonce string, _requestId uint256, _value uint256) returns()
func (_Berry *BerryTransactor) SubmitMiningSolution(opts *bind.TransactOpts, _nonce string, _requestId *big.Int, _value *big.Int) (*types.Transaction, error) {
	return _Berry.contract.Transact(opts, "submitMiningSolution", _nonce, _requestId, _value)
}

// SubmitMiningSolution is a paid mutator transaction binding the contract method 0x68c180d5.
//
// Solidity: function submitMiningSolution(_nonce string, _requestId uint256, _value uint256) returns()
func (_Berry *BerrySession) SubmitMiningSolution(_nonce string, _requestId *big.Int, _value *big.Int) (*types.Transaction, error) {
	return _Berry.Contract.SubmitMiningSolution(&_Berry.TransactOpts, _nonce, _requestId, _value)
}

// SubmitMiningSolution is a paid mutator transaction binding the contract method 0x68c180d5.
//
// Solidity: function submitMiningSolution(_nonce string, _requestId uint256, _value uint256) returns()
func (_Berry *BerryTransactorSession) SubmitMiningSolution(_nonce string, _requestId *big.Int, _value *big.Int) (*types.Transaction, error) {
	return _Berry.Contract.SubmitMiningSolution(&_Berry.TransactOpts, _nonce, _requestId, _value)
}

// TallyVotes is a paid mutator transaction binding the contract method 0x4d318b0e.
//
// Solidity: function tallyVotes(_disputeId uint256) returns()
func (_Berry *BerryTransactor) TallyVotes(opts *bind.TransactOpts, _disputeId *big.Int) (*types.Transaction, error) {
	return _Berry.contract.Transact(opts, "tallyVotes", _disputeId)
}

// TallyVotes is a paid mutator transaction binding the contract method 0x4d318b0e.
//
// Solidity: function tallyVotes(_disputeId uint256) returns()
func (_Berry *BerrySession) TallyVotes(_disputeId *big.Int) (*types.Transaction, error) {
	return _Berry.Contract.TallyVotes(&_Berry.TransactOpts, _disputeId)
}

// TallyVotes is a paid mutator transaction binding the contract method 0x4d318b0e.
//
// Solidity: function tallyVotes(_disputeId uint256) returns()
func (_Berry *BerryTransactorSession) TallyVotes(_disputeId *big.Int) (*types.Transaction, error) {
	return _Berry.Contract.TallyVotes(&_Berry.TransactOpts, _disputeId)
}

// TheLazyCoon is a paid mutator transaction binding the contract method 0xb079f64a.
//
// Solidity: function theLazyCoon(_address address, _amount uint256) returns()
func (_Berry *BerryTransactor) TheLazyCoon(opts *bind.TransactOpts, _address common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Berry.contract.Transact(opts, "theLazyCoon", _address, _amount)
}

// TheLazyCoon is a paid mutator transaction binding the contract method 0xb079f64a.
//
// Solidity: function theLazyCoon(_address address, _amount uint256) returns()
func (_Berry *BerrySession) TheLazyCoon(_address common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Berry.Contract.TheLazyCoon(&_Berry.TransactOpts, _address, _amount)
}

// TheLazyCoon is a paid mutator transaction binding the contract method 0xb079f64a.
//
// Solidity: function theLazyCoon(_address address, _amount uint256) returns()
func (_Berry *BerryTransactorSession) TheLazyCoon(_address common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Berry.Contract.TheLazyCoon(&_Berry.TransactOpts, _address, _amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(_to address, _amount uint256) returns(bool)
func (_Berry *BerryTransactor) Transfer(opts *bind.TransactOpts, _to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Berry.contract.Transact(opts, "transfer", _to, _amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(_to address, _amount uint256) returns(bool)
func (_Berry *BerrySession) Transfer(_to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Berry.Contract.Transfer(&_Berry.TransactOpts, _to, _amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(_to address, _amount uint256) returns(bool)
func (_Berry *BerryTransactorSession) Transfer(_to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Berry.Contract.Transfer(&_Berry.TransactOpts, _to, _amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(_from address, _to address, _amount uint256) returns(bool)
func (_Berry *BerryTransactor) TransferFrom(opts *bind.TransactOpts, _from common.Address, _to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Berry.contract.Transact(opts, "transferFrom", _from, _to, _amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(_from address, _to address, _amount uint256) returns(bool)
func (_Berry *BerrySession) TransferFrom(_from common.Address, _to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Berry.Contract.TransferFrom(&_Berry.TransactOpts, _from, _to, _amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(_from address, _to address, _amount uint256) returns(bool)
func (_Berry *BerryTransactorSession) TransferFrom(_from common.Address, _to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Berry.Contract.TransferFrom(&_Berry.TransactOpts, _from, _to, _amount)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(_newOwner address) returns()
func (_Berry *BerryTransactor) TransferOwnership(opts *bind.TransactOpts, _newOwner common.Address) (*types.Transaction, error) {
	return _Berry.contract.Transact(opts, "transferOwnership", _newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(_newOwner address) returns()
func (_Berry *BerrySession) TransferOwnership(_newOwner common.Address) (*types.Transaction, error) {
	return _Berry.Contract.TransferOwnership(&_Berry.TransactOpts, _newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(_newOwner address) returns()
func (_Berry *BerryTransactorSession) TransferOwnership(_newOwner common.Address) (*types.Transaction, error) {
	return _Berry.Contract.TransferOwnership(&_Berry.TransactOpts, _newOwner)
}

// Vote is a paid mutator transaction binding the contract method 0xc9d27afe.
//
// Solidity: function vote(_disputeId uint256, _supportsDispute bool) returns()
func (_Berry *BerryTransactor) Vote(opts *bind.TransactOpts, _disputeId *big.Int, _supportsDispute bool) (*types.Transaction, error) {
	return _Berry.contract.Transact(opts, "vote", _disputeId, _supportsDispute)
}

// Vote is a paid mutator transaction binding the contract method 0xc9d27afe.
//
// Solidity: function vote(_disputeId uint256, _supportsDispute bool) returns()
func (_Berry *BerrySession) Vote(_disputeId *big.Int, _supportsDispute bool) (*types.Transaction, error) {
	return _Berry.Contract.Vote(&_Berry.TransactOpts, _disputeId, _supportsDispute)
}

// Vote is a paid mutator transaction binding the contract method 0xc9d27afe.
//
// Solidity: function vote(_disputeId uint256, _supportsDispute bool) returns()
func (_Berry *BerryTransactorSession) Vote(_disputeId *big.Int, _supportsDispute bool) (*types.Transaction, error) {
	return _Berry.Contract.Vote(&_Berry.TransactOpts, _disputeId, _supportsDispute)
}

// WithdrawStake is a paid mutator transaction binding the contract method 0xbed9d861.
//
// Solidity: function withdrawStake() returns()
func (_Berry *BerryTransactor) WithdrawStake(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Berry.contract.Transact(opts, "withdrawStake")
}

// WithdrawStake is a paid mutator transaction binding the contract method 0xbed9d861.
//
// Solidity: function withdrawStake() returns()
func (_Berry *BerrySession) WithdrawStake() (*types.Transaction, error) {
	return _Berry.Contract.WithdrawStake(&_Berry.TransactOpts)
}

// WithdrawStake is a paid mutator transaction binding the contract method 0xbed9d861.
//
// Solidity: function withdrawStake() returns()
func (_Berry *BerryTransactorSession) WithdrawStake() (*types.Transaction, error) {
	return _Berry.Contract.WithdrawStake(&_Berry.TransactOpts)
}

// BerryDisputeABI is the input ABI used to generate the binding from.
const BerryDisputeABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_disputeId\",\"type\":\"uint256\"},{\"indexed\":true,\"name\":\"_requestId\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"_timestamp\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"_miner\",\"type\":\"address\"}],\"name\":\"NewDispute\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_disputeID\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"_position\",\"type\":\"bool\"},{\"indexed\":true,\"name\":\"_voter\",\"type\":\"address\"}],\"name\":\"Voted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_disputeID\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"_result\",\"type\":\"int256\"},{\"indexed\":true,\"name\":\"_reportedMiner\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_reportingParty\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_active\",\"type\":\"bool\"}],\"name\":\"DisputeVoteTallied\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_newBerry\",\"type\":\"address\"}],\"name\":\"NewBerryAddress\",\"type\":\"event\"}]"

// BerryDisputeBin is the compiled bytecode used for deploying new contracts.
const BerryDisputeBin = `0x6115f2610026600b82828239805160001a60731461001957fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600436106100615760003560e01c80632da0706e14610066578063694bf49f146100a0578063ca9a4ea5146100d9578063def6fac714610115578063e15f6f7014610145575b600080fd5b81801561007257600080fd5b5061009e6004803603606081101561008957600080fd5b5080359060208101359060400135151561016f565b005b8180156100ac57600080fd5b5061009e600480360360408110156100c357600080fd5b50803590602001356001600160a01b0316610362565b8180156100e557600080fd5b5061009e600480360360808110156100fc57600080fd5b508035906020810135906040810135906060013561079d565b81801561012157600080fd5b5061009e6004803603604081101561013857600080fd5b5080359060200135610dda565b81801561015157600080fd5b5061009e6004803603602081101561016857600080fd5b50356113e7565b600082815260448085016020908152604080842081516a313637b1b5a73ab6b132b960a91b8152825190819003600b018120865260058201845282862054633f48b1ff60e01b8252600482018a905233602483015294810194909452905190939273__$ff8037bb7c49d17e1e79a1c3f1e9fdeb7a$__92633f48b1ff92606480840193829003018186803b15801561020657600080fd5b505af415801561021a573d6000803e3d6000fd5b505050506040513d602081101561023057600080fd5b505133600090815260068401602052604090205490915060ff1615156001141561025957600080fd5b6000811161026657600080fd5b3360009081526047860160205260409020546003141561028557600080fd5b336000908152600680840160209081526040808420805460ff1916600190811790915581516c6e756d6265724f66566f74657360981b8152825190819003600d01812086526005880180855283872080549093019092556571756f72756d60d01b81528251908190039094019093208452919052902080548201905582156103165760018201805482019055610322565b60018201805482900390555b6040805184151581529051339186917f86abfce99b7dd908bec0169288797f85049ec73cbe046ed9de818fab3a497ae09181900360200190a35050505050565b604080516bffffffffffffffffffffffff19606084901b1660208083019190915282518083036014018152603490920183528151918101919091206000818152604a860190925291902054156103b757600080fd5b60408051696469737075746546656560b01b8152815190819003600a01812060009081528286016020528281205463c7bb46ad60e01b8352600483018790523360248401523060448401526064830152915173__$ff8037bb7c49d17e1e79a1c3f1e9fdeb7a$__9263c7bb46ad9260848082019391829003018186803b15801561044057600080fd5b505af4158015610454573d6000803e3d6000fd5b5050505082604001600060405180806b191a5cdc1d5d1950dbdd5b9d60a21b815250600c0190506040518091039020815260200190815260200160002060008154809291906001019190505550600083604001600060405180806b191a5cdc1d5d1950dbdd5b9d60a21b815250600c019050604051809103902081526020019081526020016000205490508084604a0160008481526020019081526020016000208190555060405180610100016040528083815260200160008152602001600015158152602001600015158152602001600115158152602001336001600160a01b03168152602001336001600160a01b03168152602001846001600160a01b0316815250846044016000838152602001908152602001600020600082015181600001556020820151816001015560408201518160020160006101000a81548160ff02191690831515021790555060608201518160020160016101000a81548160ff02191690831515021790555060808201518160020160026101000a81548160ff02191690831515021790555060a08201518160020160036101000a8154816001600160a01b0302191690836001600160a01b0316021790555060c08201518160030160006101000a8154816001600160a01b0302191690836001600160a01b0316021790555060e08201518160040160006101000a8154816001600160a01b0302191690836001600160a01b0316021790555090505043846044016000838152602001908152602001600020600501600060405180806a313637b1b5a73ab6b132b960a91b815250600b01905060405180910390208152602001908152602001600020819055508360400160006040518080696469737075746546656560b01b815250600a0190506040518091039020815260200190815260200160002054846044016000838152602001908152602001600020600501600060405180806266656560e81b815250600301905060405180910390208152602001908152602001600020819055504262093a8001846044016000838152602001908152602001600020600501600060405180806f6d696e457865637574696f6e4461746560801b8152506010019050604051809103902081526020019081526020016000208190555050505050565b6000838152604885016020908152604080832085845260058101909252909120546090439190910311156107d057600080fd5b60008381526005820160205260409020546107ea57600080fd5b600582106107f757600080fd5b60008381526008820160205260408120836005811061081257fe5b0154604080516bffffffffffffffffffffffff19606084901b1660208083019190915260348201899052605480830189905283518084039091018152607490920183528151918101919091206000818152604a8b01909252919020546001600160a01b039092169250901561088657600080fd5b60408051696469737075746546656560b01b8152815190819003600a0181206000908152828a016020528281205463c7bb46ad60e01b8352600483018b90523360248401523060448401526064830152915173__$ff8037bb7c49d17e1e79a1c3f1e9fdeb7a$__9263c7bb46ad9260848082019391829003018186803b15801561090f57600080fd5b505af4158015610923573d6000803e3d6000fd5b5050505086604001600060405180806b191a5cdc1d5d1950dbdd5b9d60a21b815250600c019050604051809103902081526020019081526020016000205460010187604001600060405180806b191a5cdc1d5d1950dbdd5b9d60a21b815250600c0190506040518091039020815260200190815260200160002081905550600087604001600060405180806b191a5cdc1d5d1950dbdd5b9d60a21b815250600c019050604051809103902081526020019081526020016000205490508088604a0160008481526020019081526020016000208190555060405180610100016040528083815260200160008152602001600015158152602001600015158152602001600015158152602001846001600160a01b03168152602001336001600160a01b0316815260200160006001600160a01b0316815250886044016000838152602001908152602001600020600082015181600001556020820151816001015560408201518160020160006101000a81548160ff02191690831515021790555060608201518160020160016101000a81548160ff02191690831515021790555060808201518160020160026101000a81548160ff02191690831515021790555060a08201518160020160036101000a8154816001600160a01b0302191690836001600160a01b0316021790555060c08201518160030160006101000a8154816001600160a01b0302191690836001600160a01b0316021790555060e08201518160040160006101000a8154816001600160a01b0302191690836001600160a01b031602179055509050508688604401600083815260200190815260200160002060050160006040518080681c995c5d595cdd125960ba1b81525060090190506040518091039020815260200190815260200160002081905550858860440160008381526020019081526020016000206005016000604051808068074696d657374616d760bc1b815250600901905060405180910390208152602001908152602001600020819055508360090160008781526020019081526020016000208560058110610c2a57fe5b0154600082815260448a016020818152604080842081516476616c756560d81b8152825190819003600590810182208752909101808452828620969096558685528383526f6d696e457865637574696f6e4461746560801b81528151908190036010018120855285835281852062093a80420190558685528383526a313637b1b5a73ab6b132b960a91b8152815190819003600b0181208552858352818520439055868552838352681b5a5b995c94db1bdd60ba1b8152815190819003600901812085528583528185208b9055696469737075746546656560b01b8152815190819003600a0181208552818e018352818520548786529383526266656560e81b815281519081900360030190208452939052919020556002851415610d7357600087815260488901602090815260408083208984526007019091529020805460ff191660011790555b6001600160a01b038316600081815260478a016020908152604091829020600390558151898152908101929092528051899284927feceec1aebf67772b2440120c4b4dc913a1fe1b865509219f9456785c23b9da6492918290030190a35050505050505050565b600081815260448301602090815260408083208151681c995c5d595cdd125960ba1b81528251908190036009019020845260058101835281842054845260488601909252909120600282015460ff1615610e3357600080fd5b604080516f6d696e457865637574696f6e4461746560801b8152815190819003601001902060009081526005840160205220544211610e7157600080fd5b600282015462010000900460ff16611258576002820154630100000090046001600160a01b03166000908152604785016020526040812060018401549091121561110d576000815562015180420642036001820155604080516a1cdd185ad95c90dbdd5b9d60aa1b8152815190819003600b0190206000908152818701602052208054600019019055610f03856113e7565b60028301546003840154604080516a1cdd185ad9505b5bdd5b9d60aa1b8152815190819003600b0181206000908152828a016020528281205463c7bb46ad60e01b8352600483018b90526001600160a01b0363010000009096048616602484015293909416604482015260648101929092525173__$ff8037bb7c49d17e1e79a1c3f1e9fdeb7a$__9263c7bb46ad9260848082019391829003018186803b158015610fad57600080fd5b505af4158015610fc1573d6000803e3d6000fd5b505050600380850154604080516266656560e81b815281519081900390930183206000908152600588016020528181205463c7bb46ad60e01b8552600485018b90523060248601526001600160a01b03909316604485015260648401929092525173__$ff8037bb7c49d17e1e79a1c3f1e9fdeb7a$__935063c7bb46ad926084808201939291829003018186803b15801561105b57600080fd5b505af415801561106f573d6000803e3d6000fd5b50505060028401805461ff001916610100179055506040805168074696d657374616d760bc1b815281519081900360090190206000908152600585016020908152828220548252600785019052205460ff16151560011415611108576040805168074696d657374616d760bc1b815281519081900360090190206000908152600585016020908152828220548252600685019052908120555b611252565b600181556002830154604080516266656560e81b815281519081900360030181206000908152600587016020528281205463c7bb46ad60e01b8352600483018a90523060248401526001600160a01b0363010000009095049490941660448301526064820193909352905173__$ff8037bb7c49d17e1e79a1c3f1e9fdeb7a$__9263c7bb46ad9260848082019391829003018186803b1580156111af57600080fd5b505af41580156111c3573d6000803e3d6000fd5b50506040805168074696d657374616d760bc1b815281519081900360090190206000908152600587016020908152828220548252600787019052205460ff1615156001141591506112529050576040805168074696d657374616d760bc1b81528151908190036009019020600090815260058501602090815282822054825260078501905220805460ff191690555b50611369565b60008260010154131561136957604080516b746f74616c5f737570706c7960a01b8152815190819003600c01902060009081528186016020522054606490601402604080516571756f72756d60d01b815281519081900360060190206000908152600586016020522054919004106112cf57600080fd5b600482018054604080516d1d195b1b1bdc90dbdb9d1c9858dd60921b8152815190819003600e0181206000908152603f890160209081529083902080546001600160a01b0319166001600160a01b0395861617905560028701805461ff00191661010017905593549092168252517fc2d1449eb0b6547aa426e09d9942a77fa4fc8cd3296305b3163e22452e0bcb8d929181900390910190a15b60028201805460ff19166001908117918290558301546003840154604080519283526001600160a01b039182166020840152610100840460ff16151583820152516301000000909304169185917f21459c2f5447ebcf83a7f0a238c32c71076faef0d12295e771c0cb1e10434739919081900360600190a350505050565b604080516b7461726765744d696e65727360a01b8152815190819003600c01812060009081528284016020818152848320546a1cdd185ad95c90dbdd5b9d60aa1b8552855194859003600b0190942083525291909120546103e8919082028161144c57fe5b04101561154657604080516b7461726765744d696e65727360a01b8152815190819003600c01812060009081528284016020818152848320546a1cdd185ad95c90dbdd5b9d60aa1b8552855194859003600b0190942083525291909120546115159167d02ab486cedc0000916103e891611508918302816114c957fe5b604080516a1cdd185ad9505b5bdd5b9d60aa1b8152815190819003600b019020600090815281890160205220549190046103e80363ffffffff61158016565b8161150f57fe5b046115a7565b60408051696469737075746546656560b01b8152815190819003600a0190206000908152818401602052205561157d565b60408051696469737075746546656560b01b8152815190819003600a01902060009081528183016020522067d02ab486cedc000090555b50565b600082820283158061159a57508284828161159757fe5b04145b6115a057fe5b9392505050565b60008183116115b657816115a0565b509091905056fea265627a7a72305820dadb2b8460cb9a7746c8f349d0279e13a2fc595a14a9d398772ad34d573d450064736f6c634300050a0032`

// DeployBerryDispute deploys a new Ethereum contract, binding an instance of BerryDispute to it.
func DeployBerryDispute(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *BerryDispute, error) {
	parsed, err := abi.JSON(strings.NewReader(BerryDisputeABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(BerryDisputeBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &BerryDispute{BerryDisputeCaller: BerryDisputeCaller{contract: contract}, BerryDisputeTransactor: BerryDisputeTransactor{contract: contract}, BerryDisputeFilterer: BerryDisputeFilterer{contract: contract}}, nil
}

// BerryDispute is an auto generated Go binding around an Ethereum contract.
type BerryDispute struct {
	BerryDisputeCaller     // Read-only binding to the contract
	BerryDisputeTransactor // Write-only binding to the contract
	BerryDisputeFilterer   // Log filterer for contract events
}

// BerryDisputeCaller is an auto generated read-only Go binding around an Ethereum contract.
type BerryDisputeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BerryDisputeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BerryDisputeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BerryDisputeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BerryDisputeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BerryDisputeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BerryDisputeSession struct {
	Contract     *BerryDispute    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BerryDisputeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BerryDisputeCallerSession struct {
	Contract *BerryDisputeCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// BerryDisputeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BerryDisputeTransactorSession struct {
	Contract     *BerryDisputeTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// BerryDisputeRaw is an auto generated low-level Go binding around an Ethereum contract.
type BerryDisputeRaw struct {
	Contract *BerryDispute // Generic contract binding to access the raw methods on
}

// BerryDisputeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BerryDisputeCallerRaw struct {
	Contract *BerryDisputeCaller // Generic read-only contract binding to access the raw methods on
}

// BerryDisputeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BerryDisputeTransactorRaw struct {
	Contract *BerryDisputeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBerryDispute creates a new instance of BerryDispute, bound to a specific deployed contract.
func NewBerryDispute(address common.Address, backend bind.ContractBackend) (*BerryDispute, error) {
	contract, err := bindBerryDispute(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BerryDispute{BerryDisputeCaller: BerryDisputeCaller{contract: contract}, BerryDisputeTransactor: BerryDisputeTransactor{contract: contract}, BerryDisputeFilterer: BerryDisputeFilterer{contract: contract}}, nil
}

// NewBerryDisputeCaller creates a new read-only instance of BerryDispute, bound to a specific deployed contract.
func NewBerryDisputeCaller(address common.Address, caller bind.ContractCaller) (*BerryDisputeCaller, error) {
	contract, err := bindBerryDispute(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BerryDisputeCaller{contract: contract}, nil
}

// NewBerryDisputeTransactor creates a new write-only instance of BerryDispute, bound to a specific deployed contract.
func NewBerryDisputeTransactor(address common.Address, transactor bind.ContractTransactor) (*BerryDisputeTransactor, error) {
	contract, err := bindBerryDispute(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BerryDisputeTransactor{contract: contract}, nil
}

// NewBerryDisputeFilterer creates a new log filterer instance of BerryDispute, bound to a specific deployed contract.
func NewBerryDisputeFilterer(address common.Address, filterer bind.ContractFilterer) (*BerryDisputeFilterer, error) {
	contract, err := bindBerryDispute(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BerryDisputeFilterer{contract: contract}, nil
}

// bindBerryDispute binds a generic wrapper to an already deployed contract.
func bindBerryDispute(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BerryDisputeABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BerryDispute *BerryDisputeRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _BerryDispute.Contract.BerryDisputeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BerryDispute *BerryDisputeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BerryDispute.Contract.BerryDisputeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BerryDispute *BerryDisputeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BerryDispute.Contract.BerryDisputeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BerryDispute *BerryDisputeCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _BerryDispute.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BerryDispute *BerryDisputeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BerryDispute.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BerryDispute *BerryDisputeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BerryDispute.Contract.contract.Transact(opts, method, params...)
}

// BerryDisputeDisputeVoteTalliedIterator is returned from FilterDisputeVoteTallied and is used to iterate over the raw logs and unpacked data for DisputeVoteTallied events raised by the BerryDispute contract.
type BerryDisputeDisputeVoteTalliedIterator struct {
	Event *BerryDisputeDisputeVoteTallied // Event containing the contract specifics and raw log

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
func (it *BerryDisputeDisputeVoteTalliedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BerryDisputeDisputeVoteTallied)
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
		it.Event = new(BerryDisputeDisputeVoteTallied)
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
func (it *BerryDisputeDisputeVoteTalliedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BerryDisputeDisputeVoteTalliedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BerryDisputeDisputeVoteTallied represents a DisputeVoteTallied event raised by the BerryDispute contract.
type BerryDisputeDisputeVoteTallied struct {
	DisputeID      *big.Int
	Result         *big.Int
	ReportedMiner  common.Address
	ReportingParty common.Address
	Active         bool
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterDisputeVoteTallied is a free log retrieval operation binding the contract event 0x21459c2f5447ebcf83a7f0a238c32c71076faef0d12295e771c0cb1e10434739.
//
// Solidity: e DisputeVoteTallied(_disputeID indexed uint256, _result int256, _reportedMiner indexed address, _reportingParty address, _active bool)
func (_BerryDispute *BerryDisputeFilterer) FilterDisputeVoteTallied(opts *bind.FilterOpts, _disputeID []*big.Int, _reportedMiner []common.Address) (*BerryDisputeDisputeVoteTalliedIterator, error) {

	var _disputeIDRule []interface{}
	for _, _disputeIDItem := range _disputeID {
		_disputeIDRule = append(_disputeIDRule, _disputeIDItem)
	}

	var _reportedMinerRule []interface{}
	for _, _reportedMinerItem := range _reportedMiner {
		_reportedMinerRule = append(_reportedMinerRule, _reportedMinerItem)
	}

	logs, sub, err := _BerryDispute.contract.FilterLogs(opts, "DisputeVoteTallied", _disputeIDRule, _reportedMinerRule)
	if err != nil {
		return nil, err
	}
	return &BerryDisputeDisputeVoteTalliedIterator{contract: _BerryDispute.contract, event: "DisputeVoteTallied", logs: logs, sub: sub}, nil
}

// WatchDisputeVoteTallied is a free log subscription operation binding the contract event 0x21459c2f5447ebcf83a7f0a238c32c71076faef0d12295e771c0cb1e10434739.
//
// Solidity: e DisputeVoteTallied(_disputeID indexed uint256, _result int256, _reportedMiner indexed address, _reportingParty address, _active bool)
func (_BerryDispute *BerryDisputeFilterer) WatchDisputeVoteTallied(opts *bind.WatchOpts, sink chan<- *BerryDisputeDisputeVoteTallied, _disputeID []*big.Int, _reportedMiner []common.Address) (event.Subscription, error) {

	var _disputeIDRule []interface{}
	for _, _disputeIDItem := range _disputeID {
		_disputeIDRule = append(_disputeIDRule, _disputeIDItem)
	}

	var _reportedMinerRule []interface{}
	for _, _reportedMinerItem := range _reportedMiner {
		_reportedMinerRule = append(_reportedMinerRule, _reportedMinerItem)
	}

	logs, sub, err := _BerryDispute.contract.WatchLogs(opts, "DisputeVoteTallied", _disputeIDRule, _reportedMinerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BerryDisputeDisputeVoteTallied)
				if err := _BerryDispute.contract.UnpackLog(event, "DisputeVoteTallied", log); err != nil {
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

// BerryDisputeNewDisputeIterator is returned from FilterNewDispute and is used to iterate over the raw logs and unpacked data for NewDispute events raised by the BerryDispute contract.
type BerryDisputeNewDisputeIterator struct {
	Event *BerryDisputeNewDispute // Event containing the contract specifics and raw log

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
func (it *BerryDisputeNewDisputeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BerryDisputeNewDispute)
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
		it.Event = new(BerryDisputeNewDispute)
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
func (it *BerryDisputeNewDisputeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BerryDisputeNewDisputeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BerryDisputeNewDispute represents a NewDispute event raised by the BerryDispute contract.
type BerryDisputeNewDispute struct {
	DisputeId *big.Int
	RequestId *big.Int
	Timestamp *big.Int
	Miner     common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterNewDispute is a free log retrieval operation binding the contract event 0xeceec1aebf67772b2440120c4b4dc913a1fe1b865509219f9456785c23b9da64.
//
// Solidity: e NewDispute(_disputeId indexed uint256, _requestId indexed uint256, _timestamp uint256, _miner address)
func (_BerryDispute *BerryDisputeFilterer) FilterNewDispute(opts *bind.FilterOpts, _disputeId []*big.Int, _requestId []*big.Int) (*BerryDisputeNewDisputeIterator, error) {

	var _disputeIdRule []interface{}
	for _, _disputeIdItem := range _disputeId {
		_disputeIdRule = append(_disputeIdRule, _disputeIdItem)
	}
	var _requestIdRule []interface{}
	for _, _requestIdItem := range _requestId {
		_requestIdRule = append(_requestIdRule, _requestIdItem)
	}

	logs, sub, err := _BerryDispute.contract.FilterLogs(opts, "NewDispute", _disputeIdRule, _requestIdRule)
	if err != nil {
		return nil, err
	}
	return &BerryDisputeNewDisputeIterator{contract: _BerryDispute.contract, event: "NewDispute", logs: logs, sub: sub}, nil
}

// WatchNewDispute is a free log subscription operation binding the contract event 0xeceec1aebf67772b2440120c4b4dc913a1fe1b865509219f9456785c23b9da64.
//
// Solidity: e NewDispute(_disputeId indexed uint256, _requestId indexed uint256, _timestamp uint256, _miner address)
func (_BerryDispute *BerryDisputeFilterer) WatchNewDispute(opts *bind.WatchOpts, sink chan<- *BerryDisputeNewDispute, _disputeId []*big.Int, _requestId []*big.Int) (event.Subscription, error) {

	var _disputeIdRule []interface{}
	for _, _disputeIdItem := range _disputeId {
		_disputeIdRule = append(_disputeIdRule, _disputeIdItem)
	}
	var _requestIdRule []interface{}
	for _, _requestIdItem := range _requestId {
		_requestIdRule = append(_requestIdRule, _requestIdItem)
	}

	logs, sub, err := _BerryDispute.contract.WatchLogs(opts, "NewDispute", _disputeIdRule, _requestIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BerryDisputeNewDispute)
				if err := _BerryDispute.contract.UnpackLog(event, "NewDispute", log); err != nil {
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

// BerryDisputeNewBerryAddressIterator is returned from FilterNewBerryAddress and is used to iterate over the raw logs and unpacked data for NewBerryAddress events raised by the BerryDispute contract.
type BerryDisputeNewBerryAddressIterator struct {
	Event *BerryDisputeNewBerryAddress // Event containing the contract specifics and raw log

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
func (it *BerryDisputeNewBerryAddressIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BerryDisputeNewBerryAddress)
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
		it.Event = new(BerryDisputeNewBerryAddress)
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
func (it *BerryDisputeNewBerryAddressIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BerryDisputeNewBerryAddressIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BerryDisputeNewBerryAddress represents a NewBerryAddress event raised by the BerryDispute contract.
type BerryDisputeNewBerryAddress struct {
	NewBerry common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterNewBerryAddress is a free log retrieval operation binding the contract event 0xc2d1449eb0b6547aa426e09d9942a77fa4fc8cd3296305b3163e22452e0bcb8d.
//
// Solidity: e NewBerryAddress(_newBerry address)
func (_BerryDispute *BerryDisputeFilterer) FilterNewBerryAddress(opts *bind.FilterOpts) (*BerryDisputeNewBerryAddressIterator, error) {

	logs, sub, err := _BerryDispute.contract.FilterLogs(opts, "NewBerryAddress")
	if err != nil {
		return nil, err
	}
	return &BerryDisputeNewBerryAddressIterator{contract: _BerryDispute.contract, event: "NewBerryAddress", logs: logs, sub: sub}, nil
}

// WatchNewBerryAddress is a free log subscription operation binding the contract event 0xc2d1449eb0b6547aa426e09d9942a77fa4fc8cd3296305b3163e22452e0bcb8d.
//
// Solidity: e NewBerryAddress(_newBerry address)
func (_BerryDispute *BerryDisputeFilterer) WatchNewBerryAddress(opts *bind.WatchOpts, sink chan<- *BerryDisputeNewBerryAddress) (event.Subscription, error) {

	logs, sub, err := _BerryDispute.contract.WatchLogs(opts, "NewBerryAddress")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BerryDisputeNewBerryAddress)
				if err := _BerryDispute.contract.UnpackLog(event, "NewBerryAddress", log); err != nil {
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

// BerryDisputeVotedIterator is returned from FilterVoted and is used to iterate over the raw logs and unpacked data for Voted events raised by the BerryDispute contract.
type BerryDisputeVotedIterator struct {
	Event *BerryDisputeVoted // Event containing the contract specifics and raw log

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
func (it *BerryDisputeVotedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BerryDisputeVoted)
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
		it.Event = new(BerryDisputeVoted)
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
func (it *BerryDisputeVotedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BerryDisputeVotedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BerryDisputeVoted represents a Voted event raised by the BerryDispute contract.
type BerryDisputeVoted struct {
	DisputeID *big.Int
	Position  bool
	Voter     common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterVoted is a free log retrieval operation binding the contract event 0x86abfce99b7dd908bec0169288797f85049ec73cbe046ed9de818fab3a497ae0.
//
// Solidity: e Voted(_disputeID indexed uint256, _position bool, _voter indexed address)
func (_BerryDispute *BerryDisputeFilterer) FilterVoted(opts *bind.FilterOpts, _disputeID []*big.Int, _voter []common.Address) (*BerryDisputeVotedIterator, error) {

	var _disputeIDRule []interface{}
	for _, _disputeIDItem := range _disputeID {
		_disputeIDRule = append(_disputeIDRule, _disputeIDItem)
	}

	var _voterRule []interface{}
	for _, _voterItem := range _voter {
		_voterRule = append(_voterRule, _voterItem)
	}

	logs, sub, err := _BerryDispute.contract.FilterLogs(opts, "Voted", _disputeIDRule, _voterRule)
	if err != nil {
		return nil, err
	}
	return &BerryDisputeVotedIterator{contract: _BerryDispute.contract, event: "Voted", logs: logs, sub: sub}, nil
}

// WatchVoted is a free log subscription operation binding the contract event 0x86abfce99b7dd908bec0169288797f85049ec73cbe046ed9de818fab3a497ae0.
//
// Solidity: e Voted(_disputeID indexed uint256, _position bool, _voter indexed address)
func (_BerryDispute *BerryDisputeFilterer) WatchVoted(opts *bind.WatchOpts, sink chan<- *BerryDisputeVoted, _disputeID []*big.Int, _voter []common.Address) (event.Subscription, error) {

	var _disputeIDRule []interface{}
	for _, _disputeIDItem := range _disputeID {
		_disputeIDRule = append(_disputeIDRule, _disputeIDItem)
	}

	var _voterRule []interface{}
	for _, _voterItem := range _voter {
		_voterRule = append(_voterRule, _voterItem)
	}

	logs, sub, err := _BerryDispute.contract.WatchLogs(opts, "Voted", _disputeIDRule, _voterRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BerryDisputeVoted)
				if err := _BerryDispute.contract.UnpackLog(event, "Voted", log); err != nil {
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

// BerryGettersLibraryABI is the input ABI used to generate the binding from.
const BerryGettersLibraryABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_newBerry\",\"type\":\"address\"}],\"name\":\"NewBerryAddress\",\"type\":\"event\"}]"

// BerryGettersLibraryBin is the compiled bytecode used for deploying new contracts.
const BerryGettersLibraryBin = `0x60556023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea265627a7a72305820da6ebe9dfbd29ba4e44451f27443534bf73520041279e6b4bb3ce66caf83e1e364736f6c634300050a0032`

// DeployBerryGettersLibrary deploys a new Ethereum contract, binding an instance of BerryGettersLibrary to it.
func DeployBerryGettersLibrary(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *BerryGettersLibrary, error) {
	parsed, err := abi.JSON(strings.NewReader(BerryGettersLibraryABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(BerryGettersLibraryBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &BerryGettersLibrary{BerryGettersLibraryCaller: BerryGettersLibraryCaller{contract: contract}, BerryGettersLibraryTransactor: BerryGettersLibraryTransactor{contract: contract}, BerryGettersLibraryFilterer: BerryGettersLibraryFilterer{contract: contract}}, nil
}

// BerryGettersLibrary is an auto generated Go binding around an Ethereum contract.
type BerryGettersLibrary struct {
	BerryGettersLibraryCaller     // Read-only binding to the contract
	BerryGettersLibraryTransactor // Write-only binding to the contract
	BerryGettersLibraryFilterer   // Log filterer for contract events
}

// BerryGettersLibraryCaller is an auto generated read-only Go binding around an Ethereum contract.
type BerryGettersLibraryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BerryGettersLibraryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BerryGettersLibraryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BerryGettersLibraryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BerryGettersLibraryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BerryGettersLibrarySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BerryGettersLibrarySession struct {
	Contract     *BerryGettersLibrary // Generic contract binding to set the session for
	CallOpts     bind.CallOpts         // Call options to use throughout this session
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// BerryGettersLibraryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BerryGettersLibraryCallerSession struct {
	Contract *BerryGettersLibraryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts               // Call options to use throughout this session
}

// BerryGettersLibraryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BerryGettersLibraryTransactorSession struct {
	Contract     *BerryGettersLibraryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts               // Transaction auth options to use throughout this session
}

// BerryGettersLibraryRaw is an auto generated low-level Go binding around an Ethereum contract.
type BerryGettersLibraryRaw struct {
	Contract *BerryGettersLibrary // Generic contract binding to access the raw methods on
}

// BerryGettersLibraryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BerryGettersLibraryCallerRaw struct {
	Contract *BerryGettersLibraryCaller // Generic read-only contract binding to access the raw methods on
}

// BerryGettersLibraryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BerryGettersLibraryTransactorRaw struct {
	Contract *BerryGettersLibraryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBerryGettersLibrary creates a new instance of BerryGettersLibrary, bound to a specific deployed contract.
func NewBerryGettersLibrary(address common.Address, backend bind.ContractBackend) (*BerryGettersLibrary, error) {
	contract, err := bindBerryGettersLibrary(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BerryGettersLibrary{BerryGettersLibraryCaller: BerryGettersLibraryCaller{contract: contract}, BerryGettersLibraryTransactor: BerryGettersLibraryTransactor{contract: contract}, BerryGettersLibraryFilterer: BerryGettersLibraryFilterer{contract: contract}}, nil
}

// NewBerryGettersLibraryCaller creates a new read-only instance of BerryGettersLibrary, bound to a specific deployed contract.
func NewBerryGettersLibraryCaller(address common.Address, caller bind.ContractCaller) (*BerryGettersLibraryCaller, error) {
	contract, err := bindBerryGettersLibrary(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BerryGettersLibraryCaller{contract: contract}, nil
}

// NewBerryGettersLibraryTransactor creates a new write-only instance of BerryGettersLibrary, bound to a specific deployed contract.
func NewBerryGettersLibraryTransactor(address common.Address, transactor bind.ContractTransactor) (*BerryGettersLibraryTransactor, error) {
	contract, err := bindBerryGettersLibrary(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BerryGettersLibraryTransactor{contract: contract}, nil
}

// NewBerryGettersLibraryFilterer creates a new log filterer instance of BerryGettersLibrary, bound to a specific deployed contract.
func NewBerryGettersLibraryFilterer(address common.Address, filterer bind.ContractFilterer) (*BerryGettersLibraryFilterer, error) {
	contract, err := bindBerryGettersLibrary(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BerryGettersLibraryFilterer{contract: contract}, nil
}

// bindBerryGettersLibrary binds a generic wrapper to an already deployed contract.
func bindBerryGettersLibrary(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BerryGettersLibraryABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BerryGettersLibrary *BerryGettersLibraryRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _BerryGettersLibrary.Contract.BerryGettersLibraryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BerryGettersLibrary *BerryGettersLibraryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BerryGettersLibrary.Contract.BerryGettersLibraryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BerryGettersLibrary *BerryGettersLibraryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BerryGettersLibrary.Contract.BerryGettersLibraryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BerryGettersLibrary *BerryGettersLibraryCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _BerryGettersLibrary.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BerryGettersLibrary *BerryGettersLibraryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BerryGettersLibrary.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BerryGettersLibrary *BerryGettersLibraryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BerryGettersLibrary.Contract.contract.Transact(opts, method, params...)
}

// BerryGettersLibraryNewBerryAddressIterator is returned from FilterNewBerryAddress and is used to iterate over the raw logs and unpacked data for NewBerryAddress events raised by the BerryGettersLibrary contract.
type BerryGettersLibraryNewBerryAddressIterator struct {
	Event *BerryGettersLibraryNewBerryAddress // Event containing the contract specifics and raw log

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
func (it *BerryGettersLibraryNewBerryAddressIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BerryGettersLibraryNewBerryAddress)
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
		it.Event = new(BerryGettersLibraryNewBerryAddress)
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
func (it *BerryGettersLibraryNewBerryAddressIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BerryGettersLibraryNewBerryAddressIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BerryGettersLibraryNewBerryAddress represents a NewBerryAddress event raised by the BerryGettersLibrary contract.
type BerryGettersLibraryNewBerryAddress struct {
	NewBerry common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterNewBerryAddress is a free log retrieval operation binding the contract event 0xc2d1449eb0b6547aa426e09d9942a77fa4fc8cd3296305b3163e22452e0bcb8d.
//
// Solidity: e NewBerryAddress(_newBerry address)
func (_BerryGettersLibrary *BerryGettersLibraryFilterer) FilterNewBerryAddress(opts *bind.FilterOpts) (*BerryGettersLibraryNewBerryAddressIterator, error) {

	logs, sub, err := _BerryGettersLibrary.contract.FilterLogs(opts, "NewBerryAddress")
	if err != nil {
		return nil, err
	}
	return &BerryGettersLibraryNewBerryAddressIterator{contract: _BerryGettersLibrary.contract, event: "NewBerryAddress", logs: logs, sub: sub}, nil
}

// WatchNewBerryAddress is a free log subscription operation binding the contract event 0xc2d1449eb0b6547aa426e09d9942a77fa4fc8cd3296305b3163e22452e0bcb8d.
//
// Solidity: e NewBerryAddress(_newBerry address)
func (_BerryGettersLibrary *BerryGettersLibraryFilterer) WatchNewBerryAddress(opts *bind.WatchOpts, sink chan<- *BerryGettersLibraryNewBerryAddress) (event.Subscription, error) {

	logs, sub, err := _BerryGettersLibrary.contract.WatchLogs(opts, "NewBerryAddress")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BerryGettersLibraryNewBerryAddress)
				if err := _BerryGettersLibrary.contract.UnpackLog(event, "NewBerryAddress", log); err != nil {
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

// BerryLibraryABI is the input ABI used to generate the binding from.
const BerryLibraryABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_sender\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"_requestId\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"_tip\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"_totalTips\",\"type\":\"uint256\"}],\"name\":\"TipAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_sender\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_query\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"_querySymbol\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"_granularity\",\"type\":\"uint256\"},{\"indexed\":true,\"name\":\"_requestId\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"_totalTips\",\"type\":\"uint256\"}],\"name\":\"DataRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_currentChallenge\",\"type\":\"bytes32\"},{\"indexed\":true,\"name\":\"_currentRequestId\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"_difficulty\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"_multiplier\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"_query\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"_totalTips\",\"type\":\"uint256\"}],\"name\":\"NewChallenge\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_requestId\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"_query\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"_onDeckQueryHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"_onDeckTotalTips\",\"type\":\"uint256\"}],\"name\":\"NewRequestOnDeck\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_requestId\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"_time\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"_value\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"_totalTips\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"_currentChallenge\",\"type\":\"bytes32\"}],\"name\":\"NewValue\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_miner\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_nonce\",\"type\":\"string\"},{\"indexed\":true,\"name\":\"_requestId\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"_value\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"_currentChallenge\",\"type\":\"bytes32\"}],\"name\":\"NonceSubmitted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"_newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"}]"

// BerryLibraryBin is the compiled bytecode used for deploying new contracts.
const BerryLibraryBin = `0x6126ec610026600b82828239805160001a60731461001957fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600436106100565760003560e01c806302e8f21b1461005b57806348656ba61461009357806385997bf6146100d2578063a098b5b414610218575b600080fd5b81801561006757600080fd5b506100916004803603606081101561007e57600080fd5b50803590602081013590604001356102d7565b005b81801561009f57600080fd5b50610091600480360360608110156100b657600080fd5b508035906001600160a01b0360208201351690604001356103ed565b8180156100de57600080fd5b50610091600480360360a08110156100f557600080fd5b8135919081019060408101602082013564010000000081111561011757600080fd5b82018360208201111561012957600080fd5b8035906020019184600183028401116401000000008311171561014b57600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929594936020810193503591505064010000000081111561019e57600080fd5b8201836020820111156101b057600080fd5b803590602001918460018302840111640100000000831117156101d257600080fd5b91908080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525092955050823593505050602001356104a5565b81801561022457600080fd5b506100916004803603608081101561023b57600080fd5b8135919081019060408101602082013564010000000081111561025d57600080fd5b82018360208201111561026f57600080fd5b8035906020019184600183028401116401000000008311171561029157600080fd5b91908080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152509295505082359350505060200135610901565b600082116102e457600080fd5b8015610368576040805163c7bb46ad60e01b81526004810185905233602482015230604482015260648101839052905173__$ff8037bb7c49d17e1e79a1c3f1e9fdeb7a$__9163c7bb46ad916084808301926000929190829003018186803b15801561034f57600080fd5b505af4158015610363573d6000803e3d6000fd5b505050505b610373838383610dc2565b600082815260488401602090815260408083208151670746f74616c5469760c41b815282519081900360080181208552600490910183529281902054848452918301919091528051849233927fd32134405b68f6f7220f9c38ae310df1b648d16188006768d45be2f7c24e882092918290030190a3505050565b604080516b746f74616c5f737570706c7960a01b8152815190819003600c018120600090815282860160209081528382208054860190556001600160a01b0386168252604587019052828120631d6f7b8160e31b8352600483015260248201849052915173__$ff8037bb7c49d17e1e79a1c3f1e9fdeb7a$__9263eb7bdc089260448082019391829003018186803b15801561048857600080fd5b505af415801561049c573d6000803e3d6000fd5b50505050505050565b600082116104b257600080fd5b670de0b6b3a76400008211156104c757600080fd5b8351849084906104d657600080fd5b60408151106104e457600080fd5b600082856040516020018083805190602001908083835b6020831061051a5780518252601f1990920191602091820191016104fb565b51815160209384036101000a60001901801990921691161790529201938452506040805180850381529382018152835193820193909320600081815260498e019092529290205491935050151590506108da57604080516b1c995c5d595cdd10dbdd5b9d60a21b8082528251600c928190038301812060009081528c8501602081815286832080546001019055938352855192839003909401822081529282528383205460808201855287825281830187905281850186905284518481528084018652606083015280845260488d01835293909220825180519192610604928492909101906124f8565b50602082810151805161061d92600185019201906124f8565b506040820151600282015560608201518051610643916003840191602090910190612576565b505050600081815260488a016020818152604080842081516a6772616e756c617269747960a81b8152825190819003600b018120865260049091018084528286208c90558686528484526f3932b8bab2b9ba28a837b9b4ba34b7b760811b825282519182900360100182208652808452828620869055868652938352670746f74616c5469760c41b81528151908190036008019020845291815281832083905584835260498c01905290208190558415610775576040805163c7bb46ad60e01b8152600481018b905233602482015230604482015260648101879052905173__$ff8037bb7c49d17e1e79a1c3f1e9fdeb7a$__9163c7bb46ad916084808301926000929190829003018186803b15801561075c57600080fd5b505af4158015610770573d6000803e3d6000fd5b505050505b610780898287610dc2565b600081815260488a016020908152604091829020825192830189905260608301889052608080845281546002600180831615610100026000190190921604918501829052859433947f6d7f869757848b19c8c2456e95cd98885bc6bed062fda4077bb07985e2f76cc99493928401928d928d929091829182019060a08301908890801561084e5780601f106108235761010080835404028352916020019161084e565b820191906000526020600020905b81548152906001019060200180831161083157829003601f168201915b50508381038252865460026000196101006001841615020190911604808252602090910190879080156108c25780601f10610897576101008083540402835291602001916108c2565b820191906000526020600020905b8154815290600101906020018083116108a557829003601f168201915b5050965050505050505060405180910390a3506108f7565b60008181526049890160205260409020546108f7908990866102d7565b5050505050505050565b33600090815260478501602052604090205460011461091f57600080fd5b604080516f18dd5c9c995b9d14995c5d595cdd125960821b8152815190819003601001902060009081528186016020522054821461095c57600080fd5b836040016000604051808069646966666963756c747960b01b815250600a0190506040518091039020815260200190815260200160002054600260038660000154338760405160200180848152602001836001600160a01b03166001600160a01b031660601b815260140182805190602001908083835b602083106109f25780518252601f1990920191602091820191016109d3565b6001836020036101000a038019825116818451168082178552505050505050905001935050505060405160208183030381529060405280519060200120604051602001808281526020019150506040516020818303038152906040526040518082805190602001908083835b60208310610a7d5780518252601f199092019160209182019101610a5e565b51815160209384036101000a60001901801990921691161790526040519190930194509192505080830381855afa158015610abc573d6000803e3d6000fd5b5050506040515160601b60405160200180826bffffffffffffffffffffffff19166bffffffffffffffffffffffff191681526014019150506040516020818303038152906040526040518082805190602001908083835b60208310610b325780518252601f199092019160209182019101610b13565b51815160209384036101000a60001901801990921691161790526040519190930194509192505080830381855afa158015610b71573d6000803e3d6000fd5b5050506040513d6020811015610b8657600080fd5b505181610b8f57fe5b0615610b9a57600080fd5b83546000908152604185016020908152604080832033845290915290205460ff1615610bc557600080fd5b604080516b736c6f7450726f677265737360a01b8152815190819003600c019020600090815281860160205220548190603586019060058110610c0457fe5b6002020155604080516b736c6f7450726f677265737360a01b8152815190819003600c019020600090815281860160205220543390603586019060058110610c4857fe5b60020201600190810180546001600160a01b0319166001600160a01b039390931692909217909155604080516b736c6f7450726f677265737360a01b81528151600c91819003919091018120600090815287830160209081528382208054860190558854825260418901815283822033808452908252848320805460ff19169096179095558854818401879052938301849052606080845288519084015287518795947fe6d63a2aee0aaed2ab49702313ce54114f2145af219d7db30d6624af9e6dffc4948a9489949293919283926080840192918801918190849084905b83811015610d3f578181015183820152602001610d27565b50505050905090810190601f168015610d6c5780820380516001836020036101000a031916815260200191505b5094505050505060405180910390a3604080516b736c6f7450726f677265737360a01b8152815190819003600c0190206000908152818601602052205460051415610dbc57610dbc8484846114ad565b50505050565b6000828152604884016020526040812090610ddc85612421565b90508215610e4b5760408051670746f74616c5469760c41b815281519081900360080190206000908152600484016020522054610e1f908463ffffffff61247d16565b60408051670746f74616c5469760c41b8152815190819003600801902060009081526004850160205220555b60408051670746f74616c5469760c41b815281519081900360080181206000908152600485016020908152838220546f18dd5c9c995b9d14995c5d595cdd125960821b84528451938490036010019093208252888401905291909120546112235760008360040160006040518080670746f74616c5469760c41b815250600801905060405180910390208152602001908152602001600020819055508486604001600060405180806f18dd5c9c995b9d14995c5d595cdd125960821b815250601001905060405180910390208152602001908152602001600020819055508086604001600060405180806f63757272656e74546f74616c5469707360801b81525060100190506040518091039020815260200190815260200160002081905550808660000154600143034060405160200180848152602001838152602001828152602001935050505060405160208183030381529060405280519060200120866000018190555085604001600060405180806f18dd5c9c995b9d14995c5d595cdd125960821b815250601001905060405180910390208152602001908152602001600020547f9e959738f09af22daede9ff9563f8edf3f5cec8e17ce60287c911c95fa767f428760000154886040016000604051808069646966666963756c747960b01b815250600a01905060405180910390208152602001908152602001600020548960480160008b604001600060405180806f18dd5c9c995b9d14995c5d595cdd125960821b815250601001905060405180910390208152602001908152602001600020548152602001908152602001600020600401600060405180806a6772616e756c617269747960a81b815250600b01905060405180910390208152602001908152602001600020548a60480160008c604001600060405180806f18dd5c9c995b9d14995c5d595cdd125960821b8152506010019050604051809103902081526020019081526020016000205481526020019081526020016000206000018b604001600060405180806f63757272656e74546f74616c5469707360801b81525060100190506040518091039020815260200190815260200160002054604051808681526020018581526020018481526020018060200183815260200182810382528481815460018160011615610100020316600290048152602001915080546001816001161561010002031660029004801561120c5780601f106111e15761010080835404028352916020019161120c565b820191906000526020600020905b8154815290600101906020018083116111ef57829003601f168201915b5050965050505050505060405180910390a26114a5565b600082815260488701602090815260408083208151670746f74616c5469760c41b815282519081900360080190208452600401909152902054811180611267575081155b156113345760028084015460408051602081018390529081018490526060808252865460001961010060018316150201169390930492810183905287927f5ab383fe3e28263c839564bc7545520a7dd34b78cbd7ab21b2da5d2fd027bf6c928792909186919081906080820190869080156113235780601f106112f857610100808354040283529160200191611323565b820191906000526020600020905b81548152906001019060200180831161130657829003601f168201915b505094505050505060405180910390a25b604080516f3932b8bab2b9ba28a837b9b4ba34b7b760811b815281519081900360100190206000908152600485016020522054611452576040805161066081019182905260009182916113a99160018b019060339082845b81548152602001906001019080831161138c575050505050612493565b9092509050818311806113ba575081155b1561144b57828860010182603381106113cf57fe5b0155600081815260438901602081815260408084208054855260488d01835281852082516f3932b8bab2b9ba28a837b9b4ba34b7b760811b80825284519182900360109081018320895260049384018752858920899055898952968652928e905591825282519182900390940190208452918801905290208190555b50506114a5565b83156114a557604080516f3932b8bab2b9ba28a837b9b4ba34b7b760811b815281519081900360100190206000908152600485016020522054849060018801906033811061149c57fe5b01805490910190555b505050505050565b6000818152604884016020908152604080832081517174696d654f664c6173744e657756616c756560701b81528251601291819003919091018120855282880180855283862054691d1a5b5955185c99d95d60b21b83528451600a938190038401812088528287528588205469646966666963756c747960b01b808352875192839003860183208a52848952878a20549083528751928390039095019091208852919095529285205491946064429590950390930302929092059091019081136115a5576040805169646966666963756c747960b01b8152815190819003600a019020600090815281870160205220600190556115d4565b6040805169646966666963756c747960b01b8152815190819003600a0190206000908152818701602052208190555b603c42604080517174696d654f664c6173744e657756616c756560701b81528151908190036012019020600090815281890160205220919006420390556116196125b0565b6040805160a081019091526035870160056000835b8282101561166f5760408051808201909152600283028501805482526001908101546001600160a01b0316602080840191909152918352909201910161162e565b509293506001925050505b600581101561179957600082826005811061169157fe5b602002015151905060008383600581106116a757fe5b602002015160200151905060008390505b6000811180156116db57508460018203600581106116d257fe5b60200201515183105b1561174d578460018203600581106116ef57fe5b60200201515185826005811061170157fe5b6020020151528460001982016005811061171757fe5b60200201516020015185826005811061172c57fe5b602090810291909101516001600160a01b03909216910152600019016116b8565b8381101561178e578285826005811061176257fe5b6020020151528185826005811061177557fe5b602090810291909101516001600160a01b039092169101525b50505060010161167a565b5060005b60058110156118ba5773__$ff8037bb7c49d17e1e79a1c3f1e9fdeb7a$__63c7bb46ad88308585600581106117ce57fe5b60200201516020015160058c604001600060405180806f63757272656e74546f74616c5469707360801b815250601001905060405180910390208152602001908152602001600020548161181e57fe5b04674563918244f40000016040518563ffffffff1660e01b815260040180858152602001846001600160a01b03166001600160a01b03168152602001836001600160a01b03166001600160a01b0316815260200182815260200194505050505060006040518083038186803b15801561189657600080fd5b505af41580156118aa573d6000803e3d6000fd5b50506001909201915061179d9050565b604080517174696d654f664c6173744e657756616c756560701b81528151908190036012018120600090815289830160208181528483205487860151516f63757272656e74546f74616c5469707360801b80875287519687900360109081018820875285855288872054918852885197889003018720865293835293869020548d549186529185019390935260059091069003828401526060820152905186917fc509d04e0782579ee59912139246ccbdf6c36c43abd90591d912017f3467dd16916080918190039190910190a2604080516b746f74616c5f737570706c7960a01b8152815190819003600c0181206000908152828a016020908152838220805468017da3a04c7b3e0000019055652fb7bbb732b960d11b835283519283900360060183208252603f8b0190528281205463c7bb46ad60e01b8352600483018b90523060248401526001600160a01b031660448301526722b1c8c1227a00006064830152915173__$ff8037bb7c49d17e1e79a1c3f1e9fdeb7a$__9263c7bb46ad9260848082019391829003018186803b158015611a5757600080fd5b505af4158015611a6b573d6000803e3d6000fd5b5050505081600260058110611a7c57fe5b6020908102919091015151604080517174696d654f664c6173744e657756616c756560701b80825282516012928190038301812060009081528d850180885285822054825260068c018852858220969096558282528451918290038401822081528587528481205460038c018054600181018255908352888320015560a08201855288518701516001600160a01b03908116835289880151880151811683890152898601518801518116838701526060808b01518901518216908401526080808b0151890151909116908301528451928352845192839003909301909120825292845281812054815260088801909352909120611b7a9160056125de565b506040805160a081018252835151815260208085015151818301528483015151828401526060808601515190830152608080860151519083015282517174696d654f664c6173744e657756616c756560701b8152835190819003601201902060009081528a840182528381205481526009880190915291909120611bff916005612632565b50604080517174696d654f664c6173744e657756616c756560701b808252825191829003601290810183206000908152848c01602081815286832054835260058b01815286832043905584865286519586900384018620835281815286832054835260428e0181528683208c9055938552855194859003909201842081528183528481205460348d01805460018101825590835284832001556b736c6f7450726f677265737360a01b8452845193849003600c019093208352905290812055611cc787612421565b604080516f18dd5c9c995b9d14995c5d595cdd125960821b80825282516010928190038301812060009081528c850160208181528683209790975592825284519182900390930190208252909252902054156123e057604080516f18dd5c9c995b9d14995c5d595cdd125960821b808252825191829003601090810183206000908152848c01602081815286832054835260488e01808252878420670746f74616c5469760c41b88528851978890036008018820855260049081018352888520546f63757272656e74546f74616c5469707360801b8952895198899003870189208652848452898620559587528751968790038501872084529181528683205483529081528582206f3932b8bab2b9ba28a837b9b4ba34b7b760811b86528651958690039093019094208152910190915290812054600189019060338110611e0b57fe5b018190555060008760430160008960480160008b604001600060405180806f18dd5c9c995b9d14995c5d595cdd125960821b815250601001905060405180910390208152602001908152602001600020548152602001908152602001600020600401600060405180806f3932b8bab2b9ba28a837b9b4ba34b7b760811b81525060100190506040518091039020815260200190815260200160002054815260200190815260200160002081905550600087604801600089604001600060405180806f18dd5c9c995b9d14995c5d595cdd125960821b815250601001905060405180910390208152602001908152602001600020548152602001908152602001600020600401600060405180806f3932b8bab2b9ba28a837b9b4ba34b7b760811b81525060100190506040518091039020815260200190815260200160002081905550600087604801600089604001600060405180806f18dd5c9c995b9d14995c5d595cdd125960821b81525060100190506040518091039020815260200190815260200160002054815260200190815260200160002060040160006040518080670746f74616c5469760c41b815250600801905060405180910390208152602001908152602001600020819055506000611fe488612421565b905086886000015460014303406040516020018084805190602001908083835b602083106120235780518252601f199092019160209182019101612004565b6001836020036101000a038019825116818451168082178552505050505050905001838152602001828152602001935050505060405160208183030381529060405280519060200120886000018190555087604001600060405180806f18dd5c9c995b9d14995c5d595cdd125960821b815250601001905060405180910390208152602001908152602001600020547f9e959738f09af22daede9ff9563f8edf3f5cec8e17ce60287c911c95fa767f4289600001548a6040016000604051808069646966666963756c747960b01b815250600a01905060405180910390208152602001908152602001600020548b60480160008d604001600060405180806f18dd5c9c995b9d14995c5d595cdd125960821b815250601001905060405180910390208152602001908152602001600020548152602001908152602001600020600401600060405180806a6772616e756c617269747960a81b815250600b01905060405180910390208152602001908152602001600020548c60480160008e604001600060405180806f18dd5c9c995b9d14995c5d595cdd125960821b8152506010019050604051809103902081526020019081526020016000205481526020019081526020016000206000018d604001600060405180806f63757272656e74546f74616c5469707360801b8152506010019050604051809103902081526020019081526020016000205460405180868152602001858152602001848152602001806020018381526020018281038252848181546001816001161561010002031660029004815260200191508054600181600116156101000203166002900480156122ce5780601f106122a3576101008083540402835291602001916122ce565b820191906000526020600020905b8154815290600101906020018083116122b157829003601f168201915b5050965050505050505060405180910390a2600081815260488901602090815260408083206002808201548351670746f74616c5469760c41b81528451908190036008018120875260048401865295849020549486018190529285018490526060808652825460001961010060018316150201169190910490850181905285947f5ab383fe3e28263c839564bc7545520a7dd34b78cbd7ab21b2da5d2fd027bf6c949293929181906080820190869080156123ca5780601f1061239f576101008083540402835291602001916123ca565b820191906000526020600020905b8154815290600101906020018083116123ad57829003601f168201915b505094505050505060405180910390a25061049c565b604080516f63757272656e74546f74616c5469707360801b815281519081900360100190206000908152818901602052908120819055875550505050505050565b60408051610660810191829052600091829182916124629190600187019060339082845b8154815260200190600101908083116124455750505050506124ca565b60009081526043909501602052505060409092205492915050565b60008282018381101561248c57fe5b9392505050565b610640810151603260315b80156124c45760208102840151838110156124ba578093508192505b506000190161249e565b50915091565b60008060005b60338110156124c45760208102840151808410156124ef578093508192505b506001016124d0565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f1061253957805160ff1916838001178555612566565b82800160010185558215612566579182015b8281111561256657825182559160200191906001019061254b565b5061257292915061265f565b5090565b828054828255906000526020600020908101928215612566579160200282018281111561256657825182559160200191906001019061254b565b6040518061014001604052806005905b6125c861267c565b8152602001906001900390816125c05790505090565b8260058101928215612626579160200282015b8281111561262657825182546001600160a01b0319166001600160a01b039091161782556020909201916001909101906125f1565b50612572929150612693565b8260058101928215612566579160200282018281111561256657825182559160200191906001019061254b565b61267991905b808211156125725760008155600101612665565b90565b604080518082019091526000808252602082015290565b61267991905b808211156125725780546001600160a01b031916815560010161269956fea265627a7a72305820cd24d65286aac042a0fcdf1c16229df2a2d28a01f35a638603abb9b2969c0fd164736f6c634300050a0032`

// DeployBerryLibrary deploys a new Ethereum contract, binding an instance of BerryLibrary to it.
func DeployBerryLibrary(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *BerryLibrary, error) {
	parsed, err := abi.JSON(strings.NewReader(BerryLibraryABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(BerryLibraryBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &BerryLibrary{BerryLibraryCaller: BerryLibraryCaller{contract: contract}, BerryLibraryTransactor: BerryLibraryTransactor{contract: contract}, BerryLibraryFilterer: BerryLibraryFilterer{contract: contract}}, nil
}

// BerryLibrary is an auto generated Go binding around an Ethereum contract.
type BerryLibrary struct {
	BerryLibraryCaller     // Read-only binding to the contract
	BerryLibraryTransactor // Write-only binding to the contract
	BerryLibraryFilterer   // Log filterer for contract events
}

// BerryLibraryCaller is an auto generated read-only Go binding around an Ethereum contract.
type BerryLibraryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BerryLibraryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BerryLibraryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BerryLibraryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BerryLibraryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BerryLibrarySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BerryLibrarySession struct {
	Contract     *BerryLibrary    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BerryLibraryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BerryLibraryCallerSession struct {
	Contract *BerryLibraryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// BerryLibraryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BerryLibraryTransactorSession struct {
	Contract     *BerryLibraryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// BerryLibraryRaw is an auto generated low-level Go binding around an Ethereum contract.
type BerryLibraryRaw struct {
	Contract *BerryLibrary // Generic contract binding to access the raw methods on
}

// BerryLibraryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BerryLibraryCallerRaw struct {
	Contract *BerryLibraryCaller // Generic read-only contract binding to access the raw methods on
}

// BerryLibraryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BerryLibraryTransactorRaw struct {
	Contract *BerryLibraryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBerryLibrary creates a new instance of BerryLibrary, bound to a specific deployed contract.
func NewBerryLibrary(address common.Address, backend bind.ContractBackend) (*BerryLibrary, error) {
	contract, err := bindBerryLibrary(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BerryLibrary{BerryLibraryCaller: BerryLibraryCaller{contract: contract}, BerryLibraryTransactor: BerryLibraryTransactor{contract: contract}, BerryLibraryFilterer: BerryLibraryFilterer{contract: contract}}, nil
}

// NewBerryLibraryCaller creates a new read-only instance of BerryLibrary, bound to a specific deployed contract.
func NewBerryLibraryCaller(address common.Address, caller bind.ContractCaller) (*BerryLibraryCaller, error) {
	contract, err := bindBerryLibrary(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BerryLibraryCaller{contract: contract}, nil
}

// NewBerryLibraryTransactor creates a new write-only instance of BerryLibrary, bound to a specific deployed contract.
func NewBerryLibraryTransactor(address common.Address, transactor bind.ContractTransactor) (*BerryLibraryTransactor, error) {
	contract, err := bindBerryLibrary(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BerryLibraryTransactor{contract: contract}, nil
}

// NewBerryLibraryFilterer creates a new log filterer instance of BerryLibrary, bound to a specific deployed contract.
func NewBerryLibraryFilterer(address common.Address, filterer bind.ContractFilterer) (*BerryLibraryFilterer, error) {
	contract, err := bindBerryLibrary(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BerryLibraryFilterer{contract: contract}, nil
}

// bindBerryLibrary binds a generic wrapper to an already deployed contract.
func bindBerryLibrary(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BerryLibraryABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BerryLibrary *BerryLibraryRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _BerryLibrary.Contract.BerryLibraryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BerryLibrary *BerryLibraryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BerryLibrary.Contract.BerryLibraryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BerryLibrary *BerryLibraryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BerryLibrary.Contract.BerryLibraryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BerryLibrary *BerryLibraryCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _BerryLibrary.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BerryLibrary *BerryLibraryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BerryLibrary.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BerryLibrary *BerryLibraryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BerryLibrary.Contract.contract.Transact(opts, method, params...)
}

// BerryLibraryDataRequestedIterator is returned from FilterDataRequested and is used to iterate over the raw logs and unpacked data for DataRequested events raised by the BerryLibrary contract.
type BerryLibraryDataRequestedIterator struct {
	Event *BerryLibraryDataRequested // Event containing the contract specifics and raw log

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
func (it *BerryLibraryDataRequestedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BerryLibraryDataRequested)
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
		it.Event = new(BerryLibraryDataRequested)
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
func (it *BerryLibraryDataRequestedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BerryLibraryDataRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BerryLibraryDataRequested represents a DataRequested event raised by the BerryLibrary contract.
type BerryLibraryDataRequested struct {
	Sender      common.Address
	Query       string
	QuerySymbol string
	Granularity *big.Int
	RequestId   *big.Int
	TotalTips   *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterDataRequested is a free log retrieval operation binding the contract event 0x6d7f869757848b19c8c2456e95cd98885bc6bed062fda4077bb07985e2f76cc9.
//
// Solidity: e DataRequested(_sender indexed address, _query string, _querySymbol string, _granularity uint256, _requestId indexed uint256, _totalTips uint256)
func (_BerryLibrary *BerryLibraryFilterer) FilterDataRequested(opts *bind.FilterOpts, _sender []common.Address, _requestId []*big.Int) (*BerryLibraryDataRequestedIterator, error) {

	var _senderRule []interface{}
	for _, _senderItem := range _sender {
		_senderRule = append(_senderRule, _senderItem)
	}

	var _requestIdRule []interface{}
	for _, _requestIdItem := range _requestId {
		_requestIdRule = append(_requestIdRule, _requestIdItem)
	}

	logs, sub, err := _BerryLibrary.contract.FilterLogs(opts, "DataRequested", _senderRule, _requestIdRule)
	if err != nil {
		return nil, err
	}
	return &BerryLibraryDataRequestedIterator{contract: _BerryLibrary.contract, event: "DataRequested", logs: logs, sub: sub}, nil
}

// WatchDataRequested is a free log subscription operation binding the contract event 0x6d7f869757848b19c8c2456e95cd98885bc6bed062fda4077bb07985e2f76cc9.
//
// Solidity: e DataRequested(_sender indexed address, _query string, _querySymbol string, _granularity uint256, _requestId indexed uint256, _totalTips uint256)
func (_BerryLibrary *BerryLibraryFilterer) WatchDataRequested(opts *bind.WatchOpts, sink chan<- *BerryLibraryDataRequested, _sender []common.Address, _requestId []*big.Int) (event.Subscription, error) {

	var _senderRule []interface{}
	for _, _senderItem := range _sender {
		_senderRule = append(_senderRule, _senderItem)
	}

	var _requestIdRule []interface{}
	for _, _requestIdItem := range _requestId {
		_requestIdRule = append(_requestIdRule, _requestIdItem)
	}

	logs, sub, err := _BerryLibrary.contract.WatchLogs(opts, "DataRequested", _senderRule, _requestIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BerryLibraryDataRequested)
				if err := _BerryLibrary.contract.UnpackLog(event, "DataRequested", log); err != nil {
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

// BerryLibraryNewChallengeIterator is returned from FilterNewChallenge and is used to iterate over the raw logs and unpacked data for NewChallenge events raised by the BerryLibrary contract.
type BerryLibraryNewChallengeIterator struct {
	Event *BerryLibraryNewChallenge // Event containing the contract specifics and raw log

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
func (it *BerryLibraryNewChallengeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BerryLibraryNewChallenge)
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
		it.Event = new(BerryLibraryNewChallenge)
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
func (it *BerryLibraryNewChallengeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BerryLibraryNewChallengeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BerryLibraryNewChallenge represents a NewChallenge event raised by the BerryLibrary contract.
type BerryLibraryNewChallenge struct {
	CurrentChallenge [32]byte
	CurrentRequestId *big.Int
	Difficulty       *big.Int
	Multiplier       *big.Int
	Query            string
	TotalTips        *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterNewChallenge is a free log retrieval operation binding the contract event 0x9e959738f09af22daede9ff9563f8edf3f5cec8e17ce60287c911c95fa767f42.
//
// Solidity: e NewChallenge(_currentChallenge bytes32, _currentRequestId indexed uint256, _difficulty uint256, _multiplier uint256, _query string, _totalTips uint256)
func (_BerryLibrary *BerryLibraryFilterer) FilterNewChallenge(opts *bind.FilterOpts, _currentRequestId []*big.Int) (*BerryLibraryNewChallengeIterator, error) {

	var _currentRequestIdRule []interface{}
	for _, _currentRequestIdItem := range _currentRequestId {
		_currentRequestIdRule = append(_currentRequestIdRule, _currentRequestIdItem)
	}

	logs, sub, err := _BerryLibrary.contract.FilterLogs(opts, "NewChallenge", _currentRequestIdRule)
	if err != nil {
		return nil, err
	}
	return &BerryLibraryNewChallengeIterator{contract: _BerryLibrary.contract, event: "NewChallenge", logs: logs, sub: sub}, nil
}

// WatchNewChallenge is a free log subscription operation binding the contract event 0x9e959738f09af22daede9ff9563f8edf3f5cec8e17ce60287c911c95fa767f42.
//
// Solidity: e NewChallenge(_currentChallenge bytes32, _currentRequestId indexed uint256, _difficulty uint256, _multiplier uint256, _query string, _totalTips uint256)
func (_BerryLibrary *BerryLibraryFilterer) WatchNewChallenge(opts *bind.WatchOpts, sink chan<- *BerryLibraryNewChallenge, _currentRequestId []*big.Int) (event.Subscription, error) {

	var _currentRequestIdRule []interface{}
	for _, _currentRequestIdItem := range _currentRequestId {
		_currentRequestIdRule = append(_currentRequestIdRule, _currentRequestIdItem)
	}

	logs, sub, err := _BerryLibrary.contract.WatchLogs(opts, "NewChallenge", _currentRequestIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BerryLibraryNewChallenge)
				if err := _BerryLibrary.contract.UnpackLog(event, "NewChallenge", log); err != nil {
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

// BerryLibraryNewRequestOnDeckIterator is returned from FilterNewRequestOnDeck and is used to iterate over the raw logs and unpacked data for NewRequestOnDeck events raised by the BerryLibrary contract.
type BerryLibraryNewRequestOnDeckIterator struct {
	Event *BerryLibraryNewRequestOnDeck // Event containing the contract specifics and raw log

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
func (it *BerryLibraryNewRequestOnDeckIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BerryLibraryNewRequestOnDeck)
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
		it.Event = new(BerryLibraryNewRequestOnDeck)
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
func (it *BerryLibraryNewRequestOnDeckIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BerryLibraryNewRequestOnDeckIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BerryLibraryNewRequestOnDeck represents a NewRequestOnDeck event raised by the BerryLibrary contract.
type BerryLibraryNewRequestOnDeck struct {
	RequestId       *big.Int
	Query           string
	OnDeckQueryHash [32]byte
	OnDeckTotalTips *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterNewRequestOnDeck is a free log retrieval operation binding the contract event 0x5ab383fe3e28263c839564bc7545520a7dd34b78cbd7ab21b2da5d2fd027bf6c.
//
// Solidity: e NewRequestOnDeck(_requestId indexed uint256, _query string, _onDeckQueryHash bytes32, _onDeckTotalTips uint256)
func (_BerryLibrary *BerryLibraryFilterer) FilterNewRequestOnDeck(opts *bind.FilterOpts, _requestId []*big.Int) (*BerryLibraryNewRequestOnDeckIterator, error) {

	var _requestIdRule []interface{}
	for _, _requestIdItem := range _requestId {
		_requestIdRule = append(_requestIdRule, _requestIdItem)
	}

	logs, sub, err := _BerryLibrary.contract.FilterLogs(opts, "NewRequestOnDeck", _requestIdRule)
	if err != nil {
		return nil, err
	}
	return &BerryLibraryNewRequestOnDeckIterator{contract: _BerryLibrary.contract, event: "NewRequestOnDeck", logs: logs, sub: sub}, nil
}

// WatchNewRequestOnDeck is a free log subscription operation binding the contract event 0x5ab383fe3e28263c839564bc7545520a7dd34b78cbd7ab21b2da5d2fd027bf6c.
//
// Solidity: e NewRequestOnDeck(_requestId indexed uint256, _query string, _onDeckQueryHash bytes32, _onDeckTotalTips uint256)
func (_BerryLibrary *BerryLibraryFilterer) WatchNewRequestOnDeck(opts *bind.WatchOpts, sink chan<- *BerryLibraryNewRequestOnDeck, _requestId []*big.Int) (event.Subscription, error) {

	var _requestIdRule []interface{}
	for _, _requestIdItem := range _requestId {
		_requestIdRule = append(_requestIdRule, _requestIdItem)
	}

	logs, sub, err := _BerryLibrary.contract.WatchLogs(opts, "NewRequestOnDeck", _requestIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BerryLibraryNewRequestOnDeck)
				if err := _BerryLibrary.contract.UnpackLog(event, "NewRequestOnDeck", log); err != nil {
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

// BerryLibraryNewValueIterator is returned from FilterNewValue and is used to iterate over the raw logs and unpacked data for NewValue events raised by the BerryLibrary contract.
type BerryLibraryNewValueIterator struct {
	Event *BerryLibraryNewValue // Event containing the contract specifics and raw log

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
func (it *BerryLibraryNewValueIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BerryLibraryNewValue)
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
		it.Event = new(BerryLibraryNewValue)
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
func (it *BerryLibraryNewValueIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BerryLibraryNewValueIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BerryLibraryNewValue represents a NewValue event raised by the BerryLibrary contract.
type BerryLibraryNewValue struct {
	RequestId        *big.Int
	Time             *big.Int
	Value            *big.Int
	TotalTips        *big.Int
	CurrentChallenge [32]byte
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterNewValue is a free log retrieval operation binding the contract event 0xc509d04e0782579ee59912139246ccbdf6c36c43abd90591d912017f3467dd16.
//
// Solidity: e NewValue(_requestId indexed uint256, _time uint256, _value uint256, _totalTips uint256, _currentChallenge bytes32)
func (_BerryLibrary *BerryLibraryFilterer) FilterNewValue(opts *bind.FilterOpts, _requestId []*big.Int) (*BerryLibraryNewValueIterator, error) {

	var _requestIdRule []interface{}
	for _, _requestIdItem := range _requestId {
		_requestIdRule = append(_requestIdRule, _requestIdItem)
	}

	logs, sub, err := _BerryLibrary.contract.FilterLogs(opts, "NewValue", _requestIdRule)
	if err != nil {
		return nil, err
	}
	return &BerryLibraryNewValueIterator{contract: _BerryLibrary.contract, event: "NewValue", logs: logs, sub: sub}, nil
}

// WatchNewValue is a free log subscription operation binding the contract event 0xc509d04e0782579ee59912139246ccbdf6c36c43abd90591d912017f3467dd16.
//
// Solidity: e NewValue(_requestId indexed uint256, _time uint256, _value uint256, _totalTips uint256, _currentChallenge bytes32)
func (_BerryLibrary *BerryLibraryFilterer) WatchNewValue(opts *bind.WatchOpts, sink chan<- *BerryLibraryNewValue, _requestId []*big.Int) (event.Subscription, error) {

	var _requestIdRule []interface{}
	for _, _requestIdItem := range _requestId {
		_requestIdRule = append(_requestIdRule, _requestIdItem)
	}

	logs, sub, err := _BerryLibrary.contract.WatchLogs(opts, "NewValue", _requestIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BerryLibraryNewValue)
				if err := _BerryLibrary.contract.UnpackLog(event, "NewValue", log); err != nil {
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

// BerryLibraryNonceSubmittedIterator is returned from FilterNonceSubmitted and is used to iterate over the raw logs and unpacked data for NonceSubmitted events raised by the BerryLibrary contract.
type BerryLibraryNonceSubmittedIterator struct {
	Event *BerryLibraryNonceSubmitted // Event containing the contract specifics and raw log

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
func (it *BerryLibraryNonceSubmittedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BerryLibraryNonceSubmitted)
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
		it.Event = new(BerryLibraryNonceSubmitted)
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
func (it *BerryLibraryNonceSubmittedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BerryLibraryNonceSubmittedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BerryLibraryNonceSubmitted represents a NonceSubmitted event raised by the BerryLibrary contract.
type BerryLibraryNonceSubmitted struct {
	Miner            common.Address
	Nonce            string
	RequestId        *big.Int
	Value            *big.Int
	CurrentChallenge [32]byte
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterNonceSubmitted is a free log retrieval operation binding the contract event 0xe6d63a2aee0aaed2ab49702313ce54114f2145af219d7db30d6624af9e6dffc4.
//
// Solidity: e NonceSubmitted(_miner indexed address, _nonce string, _requestId indexed uint256, _value uint256, _currentChallenge bytes32)
func (_BerryLibrary *BerryLibraryFilterer) FilterNonceSubmitted(opts *bind.FilterOpts, _miner []common.Address, _requestId []*big.Int) (*BerryLibraryNonceSubmittedIterator, error) {

	var _minerRule []interface{}
	for _, _minerItem := range _miner {
		_minerRule = append(_minerRule, _minerItem)
	}

	var _requestIdRule []interface{}
	for _, _requestIdItem := range _requestId {
		_requestIdRule = append(_requestIdRule, _requestIdItem)
	}

	logs, sub, err := _BerryLibrary.contract.FilterLogs(opts, "NonceSubmitted", _minerRule, _requestIdRule)
	if err != nil {
		return nil, err
	}
	return &BerryLibraryNonceSubmittedIterator{contract: _BerryLibrary.contract, event: "NonceSubmitted", logs: logs, sub: sub}, nil
}

// WatchNonceSubmitted is a free log subscription operation binding the contract event 0xe6d63a2aee0aaed2ab49702313ce54114f2145af219d7db30d6624af9e6dffc4.
//
// Solidity: e NonceSubmitted(_miner indexed address, _nonce string, _requestId indexed uint256, _value uint256, _currentChallenge bytes32)
func (_BerryLibrary *BerryLibraryFilterer) WatchNonceSubmitted(opts *bind.WatchOpts, sink chan<- *BerryLibraryNonceSubmitted, _miner []common.Address, _requestId []*big.Int) (event.Subscription, error) {

	var _minerRule []interface{}
	for _, _minerItem := range _miner {
		_minerRule = append(_minerRule, _minerItem)
	}

	var _requestIdRule []interface{}
	for _, _requestIdItem := range _requestId {
		_requestIdRule = append(_requestIdRule, _requestIdItem)
	}

	logs, sub, err := _BerryLibrary.contract.WatchLogs(opts, "NonceSubmitted", _minerRule, _requestIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BerryLibraryNonceSubmitted)
				if err := _BerryLibrary.contract.UnpackLog(event, "NonceSubmitted", log); err != nil {
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

// BerryLibraryOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the BerryLibrary contract.
type BerryLibraryOwnershipTransferredIterator struct {
	Event *BerryLibraryOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *BerryLibraryOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BerryLibraryOwnershipTransferred)
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
		it.Event = new(BerryLibraryOwnershipTransferred)
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
func (it *BerryLibraryOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BerryLibraryOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BerryLibraryOwnershipTransferred represents a OwnershipTransferred event raised by the BerryLibrary contract.
type BerryLibraryOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: e OwnershipTransferred(_previousOwner indexed address, _newOwner indexed address)
func (_BerryLibrary *BerryLibraryFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, _previousOwner []common.Address, _newOwner []common.Address) (*BerryLibraryOwnershipTransferredIterator, error) {

	var _previousOwnerRule []interface{}
	for _, _previousOwnerItem := range _previousOwner {
		_previousOwnerRule = append(_previousOwnerRule, _previousOwnerItem)
	}
	var _newOwnerRule []interface{}
	for _, _newOwnerItem := range _newOwner {
		_newOwnerRule = append(_newOwnerRule, _newOwnerItem)
	}

	logs, sub, err := _BerryLibrary.contract.FilterLogs(opts, "OwnershipTransferred", _previousOwnerRule, _newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &BerryLibraryOwnershipTransferredIterator{contract: _BerryLibrary.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: e OwnershipTransferred(_previousOwner indexed address, _newOwner indexed address)
func (_BerryLibrary *BerryLibraryFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *BerryLibraryOwnershipTransferred, _previousOwner []common.Address, _newOwner []common.Address) (event.Subscription, error) {

	var _previousOwnerRule []interface{}
	for _, _previousOwnerItem := range _previousOwner {
		_previousOwnerRule = append(_previousOwnerRule, _previousOwnerItem)
	}
	var _newOwnerRule []interface{}
	for _, _newOwnerItem := range _newOwner {
		_newOwnerRule = append(_newOwnerRule, _newOwnerItem)
	}

	logs, sub, err := _BerryLibrary.contract.WatchLogs(opts, "OwnershipTransferred", _previousOwnerRule, _newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BerryLibraryOwnershipTransferred)
				if err := _BerryLibrary.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// BerryLibraryTipAddedIterator is returned from FilterTipAdded and is used to iterate over the raw logs and unpacked data for TipAdded events raised by the BerryLibrary contract.
type BerryLibraryTipAddedIterator struct {
	Event *BerryLibraryTipAdded // Event containing the contract specifics and raw log

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
func (it *BerryLibraryTipAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BerryLibraryTipAdded)
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
		it.Event = new(BerryLibraryTipAdded)
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
func (it *BerryLibraryTipAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BerryLibraryTipAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BerryLibraryTipAdded represents a TipAdded event raised by the BerryLibrary contract.
type BerryLibraryTipAdded struct {
	Sender    common.Address
	RequestId *big.Int
	Tip       *big.Int
	TotalTips *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterTipAdded is a free log retrieval operation binding the contract event 0xd32134405b68f6f7220f9c38ae310df1b648d16188006768d45be2f7c24e8820.
//
// Solidity: e TipAdded(_sender indexed address, _requestId indexed uint256, _tip uint256, _totalTips uint256)
func (_BerryLibrary *BerryLibraryFilterer) FilterTipAdded(opts *bind.FilterOpts, _sender []common.Address, _requestId []*big.Int) (*BerryLibraryTipAddedIterator, error) {

	var _senderRule []interface{}
	for _, _senderItem := range _sender {
		_senderRule = append(_senderRule, _senderItem)
	}
	var _requestIdRule []interface{}
	for _, _requestIdItem := range _requestId {
		_requestIdRule = append(_requestIdRule, _requestIdItem)
	}

	logs, sub, err := _BerryLibrary.contract.FilterLogs(opts, "TipAdded", _senderRule, _requestIdRule)
	if err != nil {
		return nil, err
	}
	return &BerryLibraryTipAddedIterator{contract: _BerryLibrary.contract, event: "TipAdded", logs: logs, sub: sub}, nil
}

// WatchTipAdded is a free log subscription operation binding the contract event 0xd32134405b68f6f7220f9c38ae310df1b648d16188006768d45be2f7c24e8820.
//
// Solidity: e TipAdded(_sender indexed address, _requestId indexed uint256, _tip uint256, _totalTips uint256)
func (_BerryLibrary *BerryLibraryFilterer) WatchTipAdded(opts *bind.WatchOpts, sink chan<- *BerryLibraryTipAdded, _sender []common.Address, _requestId []*big.Int) (event.Subscription, error) {

	var _senderRule []interface{}
	for _, _senderItem := range _sender {
		_senderRule = append(_senderRule, _senderItem)
	}
	var _requestIdRule []interface{}
	for _, _requestIdItem := range _requestId {
		_requestIdRule = append(_requestIdRule, _requestIdItem)
	}

	logs, sub, err := _BerryLibrary.contract.WatchLogs(opts, "TipAdded", _senderRule, _requestIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BerryLibraryTipAdded)
				if err := _BerryLibrary.contract.UnpackLog(event, "TipAdded", log); err != nil {
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

// BerryStakeABI is the input ABI used to generate the binding from.
const BerryStakeABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_sender\",\"type\":\"address\"}],\"name\":\"NewStake\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_sender\",\"type\":\"address\"}],\"name\":\"StakeWithdrawn\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_sender\",\"type\":\"address\"}],\"name\":\"StakeWithdrawRequested\",\"type\":\"event\"}]"

// BerryStakeBin is the compiled bytecode used for deploying new contracts.
const BerryStakeBin = `0x6108c3610026600b82828239805160001a60731461001957fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600436106100565760003560e01c806344bacc4b1461005b5780634601f1cd14610087578063820a2d66146100b1578063c9cf5e4c146100db575b600080fd5b81801561006757600080fd5b506100856004803603602081101561007e57600080fd5b5035610105565b005b81801561009357600080fd5b50610085600480360360208110156100aa57600080fd5b5035610173565b8180156100bd57600080fd5b50610085600480360360208110156100d457600080fd5b5035610559565b8180156100e757600080fd5b50610085600480360360208110156100fe57600080fd5b50356105cd565b3360009081526047820160205260409020600181015462093a8090620151804206420303101561013457600080fd5b805460021461014257600080fd5b600080825560405133917f4a7934670bd8304e7da22378be1368f7c4fef17c5aee81804beda8638fe428ec91a25050565b6040805167646563696d616c7360c01b8152815190819003600801902060009081528183016020522054156101a757600080fd5b3060009081526045820160205260408082208151631d6f7b8160e31b8152600481019190915269014542ba12a337c00000196024820152905173__$ff8037bb7c49d17e1e79a1c3f1e9fdeb7a$__9263eb7bdc089260448082019391829003018186803b15801561021757600080fd5b505af415801561022b573d6000803e3d6000fd5b50505050610237610870565b506040805160c08101825273e037ec8ec9ec423826750853899394de7f024fee815273cdd8fa31af8475574b8909f135d510579a8087d3602082015273b9dd5afd86547df817da2d0fb89334a6f8edd8919181019190915273230570cd052f40e14c14a81038c6f3aa685d712b6060820152733233afa02644ccd048587f8ba6e99b3c00a34dcc608082015273e010ac6e0248790e08f42d5f697160dedf97e02460a082015260005b60068110156103b95773__$ff8037bb7c49d17e1e79a1c3f1e9fdeb7a$__63eb7bdc0884604501600085856006811061031557fe5b60200201516001600160a01b03166001600160a01b03168152602001908152602001600020683635c9adc5dea000006040518363ffffffff1660e01b8152600401808381526020018281526020019250505060006040518083038186803b15801561037f57600080fd5b505af4158015610393573d6000803e3d6000fd5b505050506103b1838383600681106103a757fe5b60200201516106c4565b6001016102e0565b50604080516b746f74616c5f737570706c7960a01b8152815190819003600c908101822060009081528386016020818152858320805469014542ba12a337c0000001905567646563696d616c7360c01b855285519485900360080185208352818152858320601290556b7461726765744d696e65727360a01b85528551948590039093018420825280835284822060c890556a1cdd185ad9505b5bdd5b9d60aa1b8452845193849003600b0184208252808352848220683635c9adc5dea000009055696469737075746546656560b01b8452845193849003600a908101852083528184528583206834957444b840e800009055691d1a5b5955185c99d95d60b21b80865286519586900382018620845282855286842061025890558552855194859003019093208152919052205442816104ef57fe5b604080517174696d654f664c6173744e657756616c756560701b815281519081900360120181206000908152958201602081815283882095909406420390945569646966666963756c747960b01b8152815190819003600a01902085529190529091206001905550565b61056381336106c4565b73__$5c50ea773c3f1223822c80c8edbce14ea7$__63e15f6f70826040518263ffffffff1660e01b81526004018082815260200191505060006040518083038186803b1580156105b257600080fd5b505af41580156105c6573d6000803e3d6000fd5b5050505050565b336000908152604782016020526040902080546001146105ec57600080fd5b6002815562015180420642036001820155604080516a1cdd185ad95c90dbdd5b9d60aa1b8152815190819003600b01812060009081528285016020528281208054600019019055630e15f6f760e41b825260048201859052915173__$5c50ea773c3f1223822c80c8edbce14ea7$__9263e15f6f709260248082019391829003018186803b15801561067d57600080fd5b505af4158015610691573d6000803e3d6000fd5b50506040513392507f453865710d0cb4b14ad25de371c860da196368895daa9662e5087711d14daecf9150600090a25050565b604080516a1cdd185ad9505b5bdd5b9d60aa1b8152815190819003600b0181206000908152828501602090815290839020546393b182b360e01b8352600483018690526001600160a01b0385166024840152925173__$ff8037bb7c49d17e1e79a1c3f1e9fdeb7a$__926393b182b3926044808301939192829003018186803b15801561075057600080fd5b505af4158015610764573d6000803e3d6000fd5b505050506040513d602081101561077a57600080fd5b5051101561078757600080fd5b6001600160a01b038116600090815260478301602052604090205415806107c857506001600160a01b03811660009081526047830160205260409020546002145b6107d157600080fd5b604080516a1cdd185ad95c90dbdd5b9d60aa1b8152815190819003600b01812060009081528483016020908152838220805460019081019091558385018552808452620151804290810690038285019081526001600160a01b038716808552604789019093528584209451855551930192909255915190917f46d8ab1385f70e5a3673e97c23c764f7600f7ed7a09b6687deae7131d51752e291a25050565b6040518060c00160405280600690602082028038833950919291505056fea265627a7a723058203cc3787f4bf5992f9b9f3d798c1bfdb3f004f46517aa8f3cf7a95e43f392268564736f6c634300050a0032`

// DeployBerryStake deploys a new Ethereum contract, binding an instance of BerryStake to it.
func DeployBerryStake(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *BerryStake, error) {
	parsed, err := abi.JSON(strings.NewReader(BerryStakeABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(BerryStakeBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &BerryStake{BerryStakeCaller: BerryStakeCaller{contract: contract}, BerryStakeTransactor: BerryStakeTransactor{contract: contract}, BerryStakeFilterer: BerryStakeFilterer{contract: contract}}, nil
}

// BerryStake is an auto generated Go binding around an Ethereum contract.
type BerryStake struct {
	BerryStakeCaller     // Read-only binding to the contract
	BerryStakeTransactor // Write-only binding to the contract
	BerryStakeFilterer   // Log filterer for contract events
}

// BerryStakeCaller is an auto generated read-only Go binding around an Ethereum contract.
type BerryStakeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BerryStakeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BerryStakeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BerryStakeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BerryStakeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BerryStakeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BerryStakeSession struct {
	Contract     *BerryStake      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BerryStakeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BerryStakeCallerSession struct {
	Contract *BerryStakeCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// BerryStakeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BerryStakeTransactorSession struct {
	Contract     *BerryStakeTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// BerryStakeRaw is an auto generated low-level Go binding around an Ethereum contract.
type BerryStakeRaw struct {
	Contract *BerryStake // Generic contract binding to access the raw methods on
}

// BerryStakeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BerryStakeCallerRaw struct {
	Contract *BerryStakeCaller // Generic read-only contract binding to access the raw methods on
}

// BerryStakeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BerryStakeTransactorRaw struct {
	Contract *BerryStakeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBerryStake creates a new instance of BerryStake, bound to a specific deployed contract.
func NewBerryStake(address common.Address, backend bind.ContractBackend) (*BerryStake, error) {
	contract, err := bindBerryStake(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BerryStake{BerryStakeCaller: BerryStakeCaller{contract: contract}, BerryStakeTransactor: BerryStakeTransactor{contract: contract}, BerryStakeFilterer: BerryStakeFilterer{contract: contract}}, nil
}

// NewBerryStakeCaller creates a new read-only instance of BerryStake, bound to a specific deployed contract.
func NewBerryStakeCaller(address common.Address, caller bind.ContractCaller) (*BerryStakeCaller, error) {
	contract, err := bindBerryStake(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BerryStakeCaller{contract: contract}, nil
}

// NewBerryStakeTransactor creates a new write-only instance of BerryStake, bound to a specific deployed contract.
func NewBerryStakeTransactor(address common.Address, transactor bind.ContractTransactor) (*BerryStakeTransactor, error) {
	contract, err := bindBerryStake(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BerryStakeTransactor{contract: contract}, nil
}

// NewBerryStakeFilterer creates a new log filterer instance of BerryStake, bound to a specific deployed contract.
func NewBerryStakeFilterer(address common.Address, filterer bind.ContractFilterer) (*BerryStakeFilterer, error) {
	contract, err := bindBerryStake(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BerryStakeFilterer{contract: contract}, nil
}

// bindBerryStake binds a generic wrapper to an already deployed contract.
func bindBerryStake(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BerryStakeABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BerryStake *BerryStakeRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _BerryStake.Contract.BerryStakeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BerryStake *BerryStakeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BerryStake.Contract.BerryStakeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BerryStake *BerryStakeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BerryStake.Contract.BerryStakeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BerryStake *BerryStakeCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _BerryStake.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BerryStake *BerryStakeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BerryStake.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BerryStake *BerryStakeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BerryStake.Contract.contract.Transact(opts, method, params...)
}

// BerryStakeNewStakeIterator is returned from FilterNewStake and is used to iterate over the raw logs and unpacked data for NewStake events raised by the BerryStake contract.
type BerryStakeNewStakeIterator struct {
	Event *BerryStakeNewStake // Event containing the contract specifics and raw log

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
func (it *BerryStakeNewStakeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BerryStakeNewStake)
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
		it.Event = new(BerryStakeNewStake)
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
func (it *BerryStakeNewStakeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BerryStakeNewStakeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BerryStakeNewStake represents a NewStake event raised by the BerryStake contract.
type BerryStakeNewStake struct {
	Sender common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterNewStake is a free log retrieval operation binding the contract event 0x46d8ab1385f70e5a3673e97c23c764f7600f7ed7a09b6687deae7131d51752e2.
//
// Solidity: e NewStake(_sender indexed address)
func (_BerryStake *BerryStakeFilterer) FilterNewStake(opts *bind.FilterOpts, _sender []common.Address) (*BerryStakeNewStakeIterator, error) {

	var _senderRule []interface{}
	for _, _senderItem := range _sender {
		_senderRule = append(_senderRule, _senderItem)
	}

	logs, sub, err := _BerryStake.contract.FilterLogs(opts, "NewStake", _senderRule)
	if err != nil {
		return nil, err
	}
	return &BerryStakeNewStakeIterator{contract: _BerryStake.contract, event: "NewStake", logs: logs, sub: sub}, nil
}

// WatchNewStake is a free log subscription operation binding the contract event 0x46d8ab1385f70e5a3673e97c23c764f7600f7ed7a09b6687deae7131d51752e2.
//
// Solidity: e NewStake(_sender indexed address)
func (_BerryStake *BerryStakeFilterer) WatchNewStake(opts *bind.WatchOpts, sink chan<- *BerryStakeNewStake, _sender []common.Address) (event.Subscription, error) {

	var _senderRule []interface{}
	for _, _senderItem := range _sender {
		_senderRule = append(_senderRule, _senderItem)
	}

	logs, sub, err := _BerryStake.contract.WatchLogs(opts, "NewStake", _senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BerryStakeNewStake)
				if err := _BerryStake.contract.UnpackLog(event, "NewStake", log); err != nil {
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

// BerryStakeStakeWithdrawRequestedIterator is returned from FilterStakeWithdrawRequested and is used to iterate over the raw logs and unpacked data for StakeWithdrawRequested events raised by the BerryStake contract.
type BerryStakeStakeWithdrawRequestedIterator struct {
	Event *BerryStakeStakeWithdrawRequested // Event containing the contract specifics and raw log

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
func (it *BerryStakeStakeWithdrawRequestedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BerryStakeStakeWithdrawRequested)
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
		it.Event = new(BerryStakeStakeWithdrawRequested)
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
func (it *BerryStakeStakeWithdrawRequestedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BerryStakeStakeWithdrawRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BerryStakeStakeWithdrawRequested represents a StakeWithdrawRequested event raised by the BerryStake contract.
type BerryStakeStakeWithdrawRequested struct {
	Sender common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterStakeWithdrawRequested is a free log retrieval operation binding the contract event 0x453865710d0cb4b14ad25de371c860da196368895daa9662e5087711d14daecf.
//
// Solidity: e StakeWithdrawRequested(_sender indexed address)
func (_BerryStake *BerryStakeFilterer) FilterStakeWithdrawRequested(opts *bind.FilterOpts, _sender []common.Address) (*BerryStakeStakeWithdrawRequestedIterator, error) {

	var _senderRule []interface{}
	for _, _senderItem := range _sender {
		_senderRule = append(_senderRule, _senderItem)
	}

	logs, sub, err := _BerryStake.contract.FilterLogs(opts, "StakeWithdrawRequested", _senderRule)
	if err != nil {
		return nil, err
	}
	return &BerryStakeStakeWithdrawRequestedIterator{contract: _BerryStake.contract, event: "StakeWithdrawRequested", logs: logs, sub: sub}, nil
}

// WatchStakeWithdrawRequested is a free log subscription operation binding the contract event 0x453865710d0cb4b14ad25de371c860da196368895daa9662e5087711d14daecf.
//
// Solidity: e StakeWithdrawRequested(_sender indexed address)
func (_BerryStake *BerryStakeFilterer) WatchStakeWithdrawRequested(opts *bind.WatchOpts, sink chan<- *BerryStakeStakeWithdrawRequested, _sender []common.Address) (event.Subscription, error) {

	var _senderRule []interface{}
	for _, _senderItem := range _sender {
		_senderRule = append(_senderRule, _senderItem)
	}

	logs, sub, err := _BerryStake.contract.WatchLogs(opts, "StakeWithdrawRequested", _senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BerryStakeStakeWithdrawRequested)
				if err := _BerryStake.contract.UnpackLog(event, "StakeWithdrawRequested", log); err != nil {
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

// BerryStakeStakeWithdrawnIterator is returned from FilterStakeWithdrawn and is used to iterate over the raw logs and unpacked data for StakeWithdrawn events raised by the BerryStake contract.
type BerryStakeStakeWithdrawnIterator struct {
	Event *BerryStakeStakeWithdrawn // Event containing the contract specifics and raw log

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
func (it *BerryStakeStakeWithdrawnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BerryStakeStakeWithdrawn)
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
		it.Event = new(BerryStakeStakeWithdrawn)
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
func (it *BerryStakeStakeWithdrawnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BerryStakeStakeWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BerryStakeStakeWithdrawn represents a StakeWithdrawn event raised by the BerryStake contract.
type BerryStakeStakeWithdrawn struct {
	Sender common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterStakeWithdrawn is a free log retrieval operation binding the contract event 0x4a7934670bd8304e7da22378be1368f7c4fef17c5aee81804beda8638fe428ec.
//
// Solidity: e StakeWithdrawn(_sender indexed address)
func (_BerryStake *BerryStakeFilterer) FilterStakeWithdrawn(opts *bind.FilterOpts, _sender []common.Address) (*BerryStakeStakeWithdrawnIterator, error) {

	var _senderRule []interface{}
	for _, _senderItem := range _sender {
		_senderRule = append(_senderRule, _senderItem)
	}

	logs, sub, err := _BerryStake.contract.FilterLogs(opts, "StakeWithdrawn", _senderRule)
	if err != nil {
		return nil, err
	}
	return &BerryStakeStakeWithdrawnIterator{contract: _BerryStake.contract, event: "StakeWithdrawn", logs: logs, sub: sub}, nil
}

// WatchStakeWithdrawn is a free log subscription operation binding the contract event 0x4a7934670bd8304e7da22378be1368f7c4fef17c5aee81804beda8638fe428ec.
//
// Solidity: e StakeWithdrawn(_sender indexed address)
func (_BerryStake *BerryStakeFilterer) WatchStakeWithdrawn(opts *bind.WatchOpts, sink chan<- *BerryStakeStakeWithdrawn, _sender []common.Address) (event.Subscription, error) {

	var _senderRule []interface{}
	for _, _senderItem := range _sender {
		_senderRule = append(_senderRule, _senderItem)
	}

	logs, sub, err := _BerryStake.contract.WatchLogs(opts, "StakeWithdrawn", _senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BerryStakeStakeWithdrawn)
				if err := _BerryStake.contract.UnpackLog(event, "StakeWithdrawn", log); err != nil {
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

// BerryStorageABI is the input ABI used to generate the binding from.
const BerryStorageABI = "[]"

// BerryStorageBin is the compiled bytecode used for deploying new contracts.
const BerryStorageBin = `0x60556023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea265627a7a7230582058ec2f862ab3051bed82bfb7dfd8444570e7f6447a286cf14fc693a0a697f7c664736f6c634300050a0032`

// DeployBerryStorage deploys a new Ethereum contract, binding an instance of BerryStorage to it.
func DeployBerryStorage(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *BerryStorage, error) {
	parsed, err := abi.JSON(strings.NewReader(BerryStorageABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(BerryStorageBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &BerryStorage{BerryStorageCaller: BerryStorageCaller{contract: contract}, BerryStorageTransactor: BerryStorageTransactor{contract: contract}, BerryStorageFilterer: BerryStorageFilterer{contract: contract}}, nil
}

// BerryStorage is an auto generated Go binding around an Ethereum contract.
type BerryStorage struct {
	BerryStorageCaller     // Read-only binding to the contract
	BerryStorageTransactor // Write-only binding to the contract
	BerryStorageFilterer   // Log filterer for contract events
}

// BerryStorageCaller is an auto generated read-only Go binding around an Ethereum contract.
type BerryStorageCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BerryStorageTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BerryStorageTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BerryStorageFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BerryStorageFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BerryStorageSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BerryStorageSession struct {
	Contract     *BerryStorage    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BerryStorageCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BerryStorageCallerSession struct {
	Contract *BerryStorageCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// BerryStorageTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BerryStorageTransactorSession struct {
	Contract     *BerryStorageTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// BerryStorageRaw is an auto generated low-level Go binding around an Ethereum contract.
type BerryStorageRaw struct {
	Contract *BerryStorage // Generic contract binding to access the raw methods on
}

// BerryStorageCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BerryStorageCallerRaw struct {
	Contract *BerryStorageCaller // Generic read-only contract binding to access the raw methods on
}

// BerryStorageTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BerryStorageTransactorRaw struct {
	Contract *BerryStorageTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBerryStorage creates a new instance of BerryStorage, bound to a specific deployed contract.
func NewBerryStorage(address common.Address, backend bind.ContractBackend) (*BerryStorage, error) {
	contract, err := bindBerryStorage(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BerryStorage{BerryStorageCaller: BerryStorageCaller{contract: contract}, BerryStorageTransactor: BerryStorageTransactor{contract: contract}, BerryStorageFilterer: BerryStorageFilterer{contract: contract}}, nil
}

// NewBerryStorageCaller creates a new read-only instance of BerryStorage, bound to a specific deployed contract.
func NewBerryStorageCaller(address common.Address, caller bind.ContractCaller) (*BerryStorageCaller, error) {
	contract, err := bindBerryStorage(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BerryStorageCaller{contract: contract}, nil
}

// NewBerryStorageTransactor creates a new write-only instance of BerryStorage, bound to a specific deployed contract.
func NewBerryStorageTransactor(address common.Address, transactor bind.ContractTransactor) (*BerryStorageTransactor, error) {
	contract, err := bindBerryStorage(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BerryStorageTransactor{contract: contract}, nil
}

// NewBerryStorageFilterer creates a new log filterer instance of BerryStorage, bound to a specific deployed contract.
func NewBerryStorageFilterer(address common.Address, filterer bind.ContractFilterer) (*BerryStorageFilterer, error) {
	contract, err := bindBerryStorage(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BerryStorageFilterer{contract: contract}, nil
}

// bindBerryStorage binds a generic wrapper to an already deployed contract.
func bindBerryStorage(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BerryStorageABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BerryStorage *BerryStorageRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _BerryStorage.Contract.BerryStorageCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BerryStorage *BerryStorageRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BerryStorage.Contract.BerryStorageTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BerryStorage *BerryStorageRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BerryStorage.Contract.BerryStorageTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BerryStorage *BerryStorageCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _BerryStorage.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BerryStorage *BerryStorageTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BerryStorage.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BerryStorage *BerryStorageTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BerryStorage.Contract.contract.Transact(opts, method, params...)
}

// BerryTransferABI is the input ABI used to generate the binding from.
const BerryTransferABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_owner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"_spender\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"_to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"}]"

// BerryTransferBin is the compiled bytecode used for deploying new contracts.
const BerryTransferBin = `0x610936610026600b82828239805160001a60731461001957fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600436106100a85760003560e01c8063bf32006c11610070578063bf32006c146101c5578063c7bb46ad146101f9578063c84b96f514610244578063ca50189914610283578063eb7bdc08146102cc576100a8565b80633f48b1ff146100ad578063850dcc32146100f157806393b182b3146101445780639be5647f14610170578063acaab9e214610193575b600080fd5b6100df600480360360608110156100c357600080fd5b508035906001600160a01b0360208201351690604001356102fc565b60408051918252519081900360200190f35b8180156100fd57600080fd5b506101306004803603606081101561011457600080fd5b508035906001600160a01b036020820135169060400135610395565b604080519115158252519081900360200190f35b6100df6004803603604081101561015a57600080fd5b50803590602001356001600160a01b0316610428565b6100df6004803603604081101561018657600080fd5b508035906020013561043e565b610130600480360360608110156101a957600080fd5b508035906001600160a01b03602082013516906040013561056e565b6100df600480360360608110156101db57600080fd5b508035906001600160a01b0360208201358116916040013516610614565b81801561020557600080fd5b506102426004803603608081101561021c57600080fd5b508035906001600160a01b03602082013581169160408101359091169060600135610641565b005b81801561025057600080fd5b506101306004803603606081101561026757600080fd5b508035906001600160a01b03602082013516906040013561073f565b81801561028f57600080fd5b50610130600480360360808110156102a657600080fd5b508035906001600160a01b03602082013581169160408101359091169060600135610757565b8180156102d857600080fd5b50610242600480360360408110156102ef57600080fd5b50803590602001356107cc565b6001600160a01b0382166000908152604584016020526040812054158061035a57506001600160a01b03831660009081526045850160205260408120805484929061034357fe5b6000918252602090912001546001600160801b0316115b156103675750600061038e565b6001600160a01b0383166000908152604585016020526040902061038b908361043e565b90505b9392505050565b60006103a284338461056e565b6103ab57600080fd5b6001600160a01b0383166103be57600080fd5b33600081815260468601602090815260408083206001600160a01b03881680855290835292819020869055805186815290519293927f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925929181900390910190a35060019392505050565b60006104358383436102fc565b90505b92915050565b815460009061044f57506000610438565b82548390600019810190811061046157fe5b6000918252602090912001546001600160801b031682106104b15782548390600019810190811061048e57fe5b600091825260209091200154600160801b90046001600160801b03169050610438565b826000815481106104be57fe5b6000918252602090912001546001600160801b03168210156104e257506000610438565b8254600090600019015b8181111561053d57600060026001838501010490508486828154811061050e57fe5b6000918252602090912001546001600160801b03161161053057809250610537565b6001810391505b506104ec565b84828154811061054957fe5b600091825260209091200154600160801b90046001600160801b031695945050505050565b6001600160a01b0382166000908152604784016020526040812054156105ed57604080516a1cdd185ad9505b5bdd5b9d60aa1b8152815190819003600b0190206000908152818601602052908120546105db9084906105cf90818989610428565b9063ffffffff6108a516565b106105e85750600161038e565b61060a565b60006105fd836105cf8787610428565b1061060a5750600161038e565b5060009392505050565b6001600160a01b039182166000908152604693909301602090815260408085209290931684525290205490565b6000811161064e57600080fd5b6001600160a01b03821661066157600080fd5b61066c84848361056e565b61067557600080fd5b60006106828585436102fc565b6001600160a01b038516600090815260458701602052604090209091506106ab908383036107cc565b6106b68584436102fc565b90508082820110156106c757600080fd5b6001600160a01b038316600090815260458601602052604090206106ed908284016107cc565b826001600160a01b0316846001600160a01b03167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef846040518082815260200191505060405180910390a35050505050565b600061074d84338585610641565b5060019392505050565b6001600160a01b0383166000908152604685016020908152604080832033845290915281205482111561078957600080fd5b6001600160a01b038416600090815260468601602090815260408083203384529091529020805483900390556107c185858585610641565b506001949350505050565b81541580610800575081544390839060001981019081106107e957fe5b6000918252602090912001546001600160801b0316105b15610867578154600090839061081982600183016108b7565b8154811061082357fe5b600091825260209091200180546001600160801b03848116600160801b024382166fffffffffffffffffffffffffffffffff199093169290921716179055506108a1565b81546000908390600019810190811061087c57fe5b600091825260209091200180546001600160801b03808516600160801b029116179055505b5050565b6000828211156108b157fe5b50900390565b8154818355818111156108db576000838152602090206108db9181019083016108e0565b505050565b6108fe91905b808211156108fa57600081556001016108e6565b5090565b9056fea265627a7a72305820782f81105098c84c9eb011f38e6970ea78434680be8cc6b5d842e4b5772271ea64736f6c634300050a0032`

// DeployBerryTransfer deploys a new Ethereum contract, binding an instance of BerryTransfer to it.
func DeployBerryTransfer(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *BerryTransfer, error) {
	parsed, err := abi.JSON(strings.NewReader(BerryTransferABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(BerryTransferBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &BerryTransfer{BerryTransferCaller: BerryTransferCaller{contract: contract}, BerryTransferTransactor: BerryTransferTransactor{contract: contract}, BerryTransferFilterer: BerryTransferFilterer{contract: contract}}, nil
}

// BerryTransfer is an auto generated Go binding around an Ethereum contract.
type BerryTransfer struct {
	BerryTransferCaller     // Read-only binding to the contract
	BerryTransferTransactor // Write-only binding to the contract
	BerryTransferFilterer   // Log filterer for contract events
}

// BerryTransferCaller is an auto generated read-only Go binding around an Ethereum contract.
type BerryTransferCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BerryTransferTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BerryTransferTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BerryTransferFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BerryTransferFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BerryTransferSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BerryTransferSession struct {
	Contract     *BerryTransfer   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BerryTransferCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BerryTransferCallerSession struct {
	Contract *BerryTransferCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// BerryTransferTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BerryTransferTransactorSession struct {
	Contract     *BerryTransferTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// BerryTransferRaw is an auto generated low-level Go binding around an Ethereum contract.
type BerryTransferRaw struct {
	Contract *BerryTransfer // Generic contract binding to access the raw methods on
}

// BerryTransferCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BerryTransferCallerRaw struct {
	Contract *BerryTransferCaller // Generic read-only contract binding to access the raw methods on
}

// BerryTransferTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BerryTransferTransactorRaw struct {
	Contract *BerryTransferTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBerryTransfer creates a new instance of BerryTransfer, bound to a specific deployed contract.
func NewBerryTransfer(address common.Address, backend bind.ContractBackend) (*BerryTransfer, error) {
	contract, err := bindBerryTransfer(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BerryTransfer{BerryTransferCaller: BerryTransferCaller{contract: contract}, BerryTransferTransactor: BerryTransferTransactor{contract: contract}, BerryTransferFilterer: BerryTransferFilterer{contract: contract}}, nil
}

// NewBerryTransferCaller creates a new read-only instance of BerryTransfer, bound to a specific deployed contract.
func NewBerryTransferCaller(address common.Address, caller bind.ContractCaller) (*BerryTransferCaller, error) {
	contract, err := bindBerryTransfer(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BerryTransferCaller{contract: contract}, nil
}

// NewBerryTransferTransactor creates a new write-only instance of BerryTransfer, bound to a specific deployed contract.
func NewBerryTransferTransactor(address common.Address, transactor bind.ContractTransactor) (*BerryTransferTransactor, error) {
	contract, err := bindBerryTransfer(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BerryTransferTransactor{contract: contract}, nil
}

// NewBerryTransferFilterer creates a new log filterer instance of BerryTransfer, bound to a specific deployed contract.
func NewBerryTransferFilterer(address common.Address, filterer bind.ContractFilterer) (*BerryTransferFilterer, error) {
	contract, err := bindBerryTransfer(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BerryTransferFilterer{contract: contract}, nil
}

// bindBerryTransfer binds a generic wrapper to an already deployed contract.
func bindBerryTransfer(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BerryTransferABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BerryTransfer *BerryTransferRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _BerryTransfer.Contract.BerryTransferCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BerryTransfer *BerryTransferRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BerryTransfer.Contract.BerryTransferTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BerryTransfer *BerryTransferRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BerryTransfer.Contract.BerryTransferTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BerryTransfer *BerryTransferCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _BerryTransfer.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BerryTransfer *BerryTransferTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BerryTransfer.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BerryTransfer *BerryTransferTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BerryTransfer.Contract.contract.Transact(opts, method, params...)
}

// BerryTransferApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the BerryTransfer contract.
type BerryTransferApprovalIterator struct {
	Event *BerryTransferApproval // Event containing the contract specifics and raw log

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
func (it *BerryTransferApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BerryTransferApproval)
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
		it.Event = new(BerryTransferApproval)
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
func (it *BerryTransferApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BerryTransferApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BerryTransferApproval represents a Approval event raised by the BerryTransfer contract.
type BerryTransferApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: e Approval(_owner indexed address, _spender indexed address, _value uint256)
func (_BerryTransfer *BerryTransferFilterer) FilterApproval(opts *bind.FilterOpts, _owner []common.Address, _spender []common.Address) (*BerryTransferApprovalIterator, error) {

	var _ownerRule []interface{}
	for _, _ownerItem := range _owner {
		_ownerRule = append(_ownerRule, _ownerItem)
	}
	var _spenderRule []interface{}
	for _, _spenderItem := range _spender {
		_spenderRule = append(_spenderRule, _spenderItem)
	}

	logs, sub, err := _BerryTransfer.contract.FilterLogs(opts, "Approval", _ownerRule, _spenderRule)
	if err != nil {
		return nil, err
	}
	return &BerryTransferApprovalIterator{contract: _BerryTransfer.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: e Approval(_owner indexed address, _spender indexed address, _value uint256)
func (_BerryTransfer *BerryTransferFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *BerryTransferApproval, _owner []common.Address, _spender []common.Address) (event.Subscription, error) {

	var _ownerRule []interface{}
	for _, _ownerItem := range _owner {
		_ownerRule = append(_ownerRule, _ownerItem)
	}
	var _spenderRule []interface{}
	for _, _spenderItem := range _spender {
		_spenderRule = append(_spenderRule, _spenderItem)
	}

	logs, sub, err := _BerryTransfer.contract.WatchLogs(opts, "Approval", _ownerRule, _spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BerryTransferApproval)
				if err := _BerryTransfer.contract.UnpackLog(event, "Approval", log); err != nil {
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

// BerryTransferTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the BerryTransfer contract.
type BerryTransferTransferIterator struct {
	Event *BerryTransferTransfer // Event containing the contract specifics and raw log

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
func (it *BerryTransferTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BerryTransferTransfer)
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
		it.Event = new(BerryTransferTransfer)
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
func (it *BerryTransferTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BerryTransferTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BerryTransferTransfer represents a Transfer event raised by the BerryTransfer contract.
type BerryTransferTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: e Transfer(_from indexed address, _to indexed address, _value uint256)
func (_BerryTransfer *BerryTransferFilterer) FilterTransfer(opts *bind.FilterOpts, _from []common.Address, _to []common.Address) (*BerryTransferTransferIterator, error) {

	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}
	var _toRule []interface{}
	for _, _toItem := range _to {
		_toRule = append(_toRule, _toItem)
	}

	logs, sub, err := _BerryTransfer.contract.FilterLogs(opts, "Transfer", _fromRule, _toRule)
	if err != nil {
		return nil, err
	}
	return &BerryTransferTransferIterator{contract: _BerryTransfer.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: e Transfer(_from indexed address, _to indexed address, _value uint256)
func (_BerryTransfer *BerryTransferFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *BerryTransferTransfer, _from []common.Address, _to []common.Address) (event.Subscription, error) {

	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}
	var _toRule []interface{}
	for _, _toItem := range _to {
		_toRule = append(_toRule, _toItem)
	}

	logs, sub, err := _BerryTransfer.contract.WatchLogs(opts, "Transfer", _fromRule, _toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BerryTransferTransfer)
				if err := _BerryTransfer.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// UtilitiesABI is the input ABI used to generate the binding from.
const UtilitiesABI = "[]"

// UtilitiesBin is the compiled bytecode used for deploying new contracts.
const UtilitiesBin = `0x60556023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea265627a7a723058203a5f66bda6f7a12c4d2010ffb4ddea6b3e3fea0390b238de6b2ad2897120313b64736f6c634300050a0032`

// DeployUtilities deploys a new Ethereum contract, binding an instance of Utilities to it.
func DeployUtilities(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Utilities, error) {
	parsed, err := abi.JSON(strings.NewReader(UtilitiesABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(UtilitiesBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Utilities{UtilitiesCaller: UtilitiesCaller{contract: contract}, UtilitiesTransactor: UtilitiesTransactor{contract: contract}, UtilitiesFilterer: UtilitiesFilterer{contract: contract}}, nil
}

// Utilities is an auto generated Go binding around an Ethereum contract.
type Utilities struct {
	UtilitiesCaller     // Read-only binding to the contract
	UtilitiesTransactor // Write-only binding to the contract
	UtilitiesFilterer   // Log filterer for contract events
}

// UtilitiesCaller is an auto generated read-only Go binding around an Ethereum contract.
type UtilitiesCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UtilitiesTransactor is an auto generated write-only Go binding around an Ethereum contract.
type UtilitiesTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UtilitiesFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type UtilitiesFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UtilitiesSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type UtilitiesSession struct {
	Contract     *Utilities        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// UtilitiesCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type UtilitiesCallerSession struct {
	Contract *UtilitiesCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// UtilitiesTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type UtilitiesTransactorSession struct {
	Contract     *UtilitiesTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// UtilitiesRaw is an auto generated low-level Go binding around an Ethereum contract.
type UtilitiesRaw struct {
	Contract *Utilities // Generic contract binding to access the raw methods on
}

// UtilitiesCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type UtilitiesCallerRaw struct {
	Contract *UtilitiesCaller // Generic read-only contract binding to access the raw methods on
}

// UtilitiesTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type UtilitiesTransactorRaw struct {
	Contract *UtilitiesTransactor // Generic write-only contract binding to access the raw methods on
}

// NewUtilities creates a new instance of Utilities, bound to a specific deployed contract.
func NewUtilities(address common.Address, backend bind.ContractBackend) (*Utilities, error) {
	contract, err := bindUtilities(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Utilities{UtilitiesCaller: UtilitiesCaller{contract: contract}, UtilitiesTransactor: UtilitiesTransactor{contract: contract}, UtilitiesFilterer: UtilitiesFilterer{contract: contract}}, nil
}

// NewUtilitiesCaller creates a new read-only instance of Utilities, bound to a specific deployed contract.
func NewUtilitiesCaller(address common.Address, caller bind.ContractCaller) (*UtilitiesCaller, error) {
	contract, err := bindUtilities(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &UtilitiesCaller{contract: contract}, nil
}

// NewUtilitiesTransactor creates a new write-only instance of Utilities, bound to a specific deployed contract.
func NewUtilitiesTransactor(address common.Address, transactor bind.ContractTransactor) (*UtilitiesTransactor, error) {
	contract, err := bindUtilities(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &UtilitiesTransactor{contract: contract}, nil
}

// NewUtilitiesFilterer creates a new log filterer instance of Utilities, bound to a specific deployed contract.
func NewUtilitiesFilterer(address common.Address, filterer bind.ContractFilterer) (*UtilitiesFilterer, error) {
	contract, err := bindUtilities(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &UtilitiesFilterer{contract: contract}, nil
}

// bindUtilities binds a generic wrapper to an already deployed contract.
func bindUtilities(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(UtilitiesABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Utilities *UtilitiesRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Utilities.Contract.UtilitiesCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Utilities *UtilitiesRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Utilities.Contract.UtilitiesTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Utilities *UtilitiesRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Utilities.Contract.UtilitiesTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Utilities *UtilitiesCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Utilities.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Utilities *UtilitiesTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Utilities.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Utilities *UtilitiesTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Utilities.Contract.contract.Transact(opts, method, params...)
}
