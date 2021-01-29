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

// BerryGettersABI is the input ABI used to generate the binding from.
const BerryGettersABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"_timestamp\",\"type\":\"uint256\"}],\"name\":\"getRequestIdByTimestamp\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_requestId\",\"type\":\"uint256\"},{\"name\":\"_timestamp\",\"type\":\"uint256\"}],\"name\":\"getSubmissionsByTimestamp\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256[5]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_data\",\"type\":\"bytes32\"}],\"name\":\"getAddressVars\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getSymbol\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getName\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getVariablesOnDeck\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_request\",\"type\":\"bytes32\"}],\"name\":\"getRequestIdByQueryHash\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_requestId\",\"type\":\"uint256\"}],\"name\":\"getLastNewValueById\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_requestId\",\"type\":\"uint256\"},{\"name\":\"_timestamp\",\"type\":\"uint256\"}],\"name\":\"isInDispute\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_requestId\",\"type\":\"uint256\"}],\"name\":\"getNewValueCountbyRequestId\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_user\",\"type\":\"address\"},{\"name\":\"_blockNumber\",\"type\":\"uint256\"}],\"name\":\"balanceOfAt\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_data\",\"type\":\"bytes32\"}],\"name\":\"getUintVar\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_index\",\"type\":\"uint256\"}],\"name\":\"getRequestIdByRequestQIndex\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_challenge\",\"type\":\"bytes32\"},{\"name\":\"_miner\",\"type\":\"address\"}],\"name\":\"didMine\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_requestId\",\"type\":\"uint256\"},{\"name\":\"_timestamp\",\"type\":\"uint256\"}],\"name\":\"getMinersByRequestIdAndTimestamp\",\"outputs\":[{\"name\":\"\",\"type\":\"address[5]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_user\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_staker\",\"type\":\"address\"}],\"name\":\"getStakerInfo\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_requestID\",\"type\":\"uint256\"},{\"name\":\"_index\",\"type\":\"uint256\"}],\"name\":\"getTimestampbyRequestIDandIndex\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_disputeId\",\"type\":\"uint256\"},{\"name\":\"_data\",\"type\":\"bytes32\"}],\"name\":\"getDisputeUintVars\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_requestId\",\"type\":\"uint256\"},{\"name\":\"_timestamp\",\"type\":\"uint256\"}],\"name\":\"retrieveData\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_user\",\"type\":\"address\"},{\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"allowedToTrade\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getCurrentVariables\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"},{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"string\"},{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_disputeId\",\"type\":\"uint256\"},{\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"didVote\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_disputeId\",\"type\":\"uint256\"}],\"name\":\"getAllDisputeVars\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"},{\"name\":\"\",\"type\":\"bool\"},{\"name\":\"\",\"type\":\"bool\"},{\"name\":\"\",\"type\":\"bool\"},{\"name\":\"\",\"type\":\"address\"},{\"name\":\"\",\"type\":\"address\"},{\"name\":\"\",\"type\":\"address\"},{\"name\":\"\",\"type\":\"uint256[9]\"},{\"name\":\"\",\"type\":\"int256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getRequestQ\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256[51]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_requestId\",\"type\":\"uint256\"},{\"name\":\"_timestamp\",\"type\":\"uint256\"}],\"name\":\"getMinedBlockNum\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_hash\",\"type\":\"bytes32\"}],\"name\":\"getDisputeIdByDisputeHash\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_user\",\"type\":\"address\"},{\"name\":\"_spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_requestId\",\"type\":\"uint256\"},{\"name\":\"_data\",\"type\":\"bytes32\"}],\"name\":\"getRequestUintVars\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_requestId\",\"type\":\"uint256\"}],\"name\":\"getRequestVars\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"},{\"name\":\"\",\"type\":\"string\"},{\"name\":\"\",\"type\":\"bytes32\"},{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getLastNewValue\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// BerryGettersBin is the compiled bytecode used for deploying new contracts.
const BerryGettersBin = `0x608060405234801561001057600080fd5b50611a2a806100206000396000f3fe608060405234801561001057600080fd5b50600436106101f05760003560e01c806370a082311161010f578063af0b1327116100a2578063dd62ed3e11610071578063dd62ed3e146107fa578063e0ae93c114610828578063e1eee6d61461084b578063fc7cf0a014610962576101f0565b8063af0b1327146106f8578063b54130291461079c578063c775b542146107ba578063da379941146107dd576101f0565b806393fa4915116100de57806393fa4915146105da578063999cf26c146105fd578063a22e407a14610629578063a7c438bc146106cc576101f0565b806370a082311461052f578063733bdef01461055557806377fbb663146105945780637f6fd5d9146105b7576101f0565b80633180f8df11610187578063612c8f7f11610156578063612c8f7f146104a65780636173c0b8146104c357806363bb82ad146104e057806369026d631461050c576101f0565b80633180f8df146103f05780633df0777b1461042657806346eee1c41461045d5780634ee2cd7e1461047a576101f0565b806317d7de7c116101c357806317d7de7c1461033557806318160ddd1461033d57806319e8e03b146103455780631db842f0146103d3576101f0565b80630f0b424d146101f557806311c9851214610224578063133bee5e1461027f57806315070401146102b8575b600080fd5b6102126004803603602081101561020b57600080fd5b503561096a565b60408051918252519081900360200190f35b6102476004803603604081101561023a57600080fd5b5080359060200135610982565b604051808260a080838360005b8381101561026c578181015183820152602001610254565b5050505090500191505060405180910390f35b61029c6004803603602081101561029557600080fd5b50356109a3565b604080516001600160a01b039092168252519081900360200190f35b6102c06109b5565b6040805160208082528351818301528351919283929083019185019080838360005b838110156102fa5781810151838201526020016102e2565b50505050905090810190601f1680156103275780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b6102c06109c6565b6102126109d2565b61034d6109de565b6040518084815260200183815260200180602001828103825283818151815260200191508051906020019080838360005b8381101561039657818101518382015260200161037e565b50505050905090810190601f1680156103c35780820380516001836020036101000a031916815260200191505b5094505050505060405180910390f35b610212600480360360208110156103e957600080fd5b50356109f8565b61040d6004803603602081101561040657600080fd5b5035610a0a565b6040805192835290151560208301528051918290030190f35b6104496004803603604081101561043c57600080fd5b5080359060200135610a26565b604080519115158252519081900360200190f35b6102126004803603602081101561047357600080fd5b5035610a39565b6102126004803603604081101561049057600080fd5b506001600160a01b038135169060200135610a4b565b610212600480360360208110156104bc57600080fd5b5035610ae8565b610212600480360360208110156104d957600080fd5b5035610afa565b610449600480360360408110156104f657600080fd5b50803590602001356001600160a01b0316610b0c565b6102476004803603604081101561052257600080fd5b5080359060200135610b1f565b6102126004803603602081101561054557600080fd5b50356001600160a01b0316610b39565b61057b6004803603602081101561056b57600080fd5b50356001600160a01b0316610bce565b6040805192835260208301919091528051918290030190f35b610212600480360360408110156105aa57600080fd5b5080359060200135610be1565b610212600480360360408110156105cd57600080fd5b5080359060200135610bf4565b610212600480360360408110156105f057600080fd5b5080359060200135610c07565b6104496004803603604081101561061357600080fd5b506001600160a01b038135169060200135610c1a565b610631610c84565b6040518087815260200186815260200185815260200180602001848152602001838152602001828103825285818151815260200191508051906020019080838360005b8381101561068c578181015183820152602001610674565b50505050905090810190601f1680156106b95780820380516001836020036101000a031916815260200191505b5097505050505050505060405180910390f35b610449600480360360408110156106e257600080fd5b50803590602001356001600160a01b0316610cab565b6107156004803603602081101561070e57600080fd5b5035610cbe565b604080518a815289151560208201528815159181019190915286151560608201526001600160a01b03808716608083015285811660a0830152841660c082015260e081018361012080838360005b8381101561077b578181015183820152602001610763565b50505050905001828152602001995050505050505050505060405180910390f35b6107a4610d02565b6040518151815280826106608083836020610254565b610212600480360360408110156107d057600080fd5b5080359060200135610d14565b610212600480360360208110156107f357600080fd5b5035610d27565b6102126004803603604081101561081057600080fd5b506001600160a01b0381358116916020013516610d39565b6102126004803603604081101561083e57600080fd5b5080359060200135610da4565b6108686004803603602081101561086157600080fd5b5035610db7565b604051808060200180602001878152602001868152602001858152602001848152602001838103835289818151815260200191508051906020019080838360005b838110156108c15781810151838201526020016108a9565b50505050905090810190601f1680156108ee5780820380516001836020036101000a031916815260200191505b5083810382528851815288516020918201918a019080838360005b83811015610921578181015183820152602001610909565b50505050905090810190601f16801561094e5780820380516001836020036101000a031916815260200191505b509850505050505050505060405180910390f35b61040d610de3565b600061097c818363ffffffff610df816565b92915050565b61098a611999565b61099c6000848463ffffffff610e0e16565b9392505050565b600061097c818363ffffffff610e6816565b60606109c16000610e87565b905090565b60606109c16000610ea4565b60006109c16000610ece565b60008060606109ed6000610f01565b925092509250909192565b600061097c818363ffffffff610fed16565b600080610a1d818463ffffffff61100316565b91509150915091565b600061099c81848463ffffffff61106916565b600061097c818363ffffffff61109016565b60408051633f48b1ff60e01b81526000600482018190526001600160a01b038516602483015260448201849052915173__$ff8037bb7c49d17e1e79a1c3f1e9fdeb7a$__91633f48b1ff916064808301926020929190829003018186803b158015610ab557600080fd5b505af4158015610ac9573d6000803e3d6000fd5b505050506040513d6020811015610adf57600080fd5b50519392505050565b600061097c818363ffffffff6110a916565b600061097c818363ffffffff6110bb16565b600061099c81848463ffffffff6110e216565b610b27611999565b61099c6000848463ffffffff61110f16565b604080516393b182b360e01b81526000600482018190526001600160a01b0384166024830152915173__$ff8037bb7c49d17e1e79a1c3f1e9fdeb7a$__916393b182b3916044808301926020929190829003018186803b158015610b9c57600080fd5b505af4158015610bb0573d6000803e3d6000fd5b505050506040513d6020811015610bc657600080fd5b505192915050565b600080610a1d818463ffffffff61117316565b600061099c81848463ffffffff61119a16565b600061099c81848463ffffffff6111cd16565b600061099c81848463ffffffff6111f116565b604080516356555cf160e11b81526000600482018190526001600160a01b038516602483015260448201849052915173__$ff8037bb7c49d17e1e79a1c3f1e9fdeb7a$__9163acaab9e2916064808301926020929190829003018186803b158015610ab557600080fd5b60008060006060600080610c986000611215565b949b939a50919850965094509092509050565b600061099c81848463ffffffff6113d316565b6000806000806000806000610cd16119b7565b6000610ce3818b63ffffffff61140516565b9850985098509850985098509850985098509193959799909294969850565b610d0a6119d6565b6109c1600061161a565b600061099c81848463ffffffff61165a16565b600061097c818363ffffffff61167e16565b60408051632fcc801b60e21b81526000600482018190526001600160a01b03808616602484015284166044830152915173__$ff8037bb7c49d17e1e79a1c3f1e9fdeb7a$__9163bf32006c916064808301926020929190829003018186803b158015610ab557600080fd5b600061099c81848463ffffffff61169416565b6060806000808080610dcf818863ffffffff6116b816565b949c939b5091995097509550909350915050565b600080610df06000611893565b915091509091565b6000908152604291909101602052604090205490565b610e16611999565b6000838152604885016020908152604080832085845260090190915290819020815160a08101928390529160059082845b815481526020019060010190808311610e4757505050505090509392505050565b6000908152603f9190910160205260409020546001600160a01b031690565b50604080518082019091526002815261151560f21b602082015290565b5060408051808201909152600f81526e54656c6c6f7220547269627574657360881b602082015290565b604080516b746f74616c5f737570706c7960a01b8152815190819003600c01902060009081528183016020522054919050565b60008060606000610f1185611909565b600081815260488701602081815260408084208151670746f74616c5469760c41b8152825160089181900391909101812086526004820184528286205495879052938352805460026001821615610100026000190190911604601f810184900484028501840190925281845294955085949291839190830182828015610fd85780601f10610fad57610100808354040283529160200191610fd8565b820191906000526020600020905b815481529060010190602001808311610fbb57829003601f168201915b50505050509050935093509350509193909250565b6000908152604991909101602052604090205490565b600081815260488301602052604081206003810154829190156110595760038101805461104d918791879190600019810190811061103d57fe5b90600052602060002001546111f1565b60019250925050611062565b50600091508190505b9250929050565b60009182526048929092016020908152604080832093835260079093019052205460ff1690565b6000908152604891909101602052604090206003015490565b60009081526040918201602052205490565b600060328211156110cb57600080fd5b506000908152604391909101602052604090205490565b6000918252604192909201602090815260408083206001600160a01b039094168352929052205460ff1690565b611117611999565b6000838152604885016020908152604080832085845260080190915290819020815160a08101928390529160059082845b81546001600160a01b0316815260019091019060200180831161114857505050505090509392505050565b6001600160a01b031660009081526047919091016020526040902080546001909101549091565b600082815260488401602052604081206003018054839081106111b957fe5b906000526020600020015490509392505050565b60009182526044929092016020908152604080832093835260059093019052205490565b60009182526048929092016020908152604080832093835260069093019052205490565b8054604080516f18dd5c9c995b9d14995c5d595cdd125960821b80825282519182900360109081018320600090815284870160208181528683205469646966666963756c747960b01b8752875196879003600a01872084528282528784205486885288519788900386018820855283835288852054855260488b01808452898620888a528a51998a900388018a2087528585528a87205487528185528a87206a6772616e756c617269747960a81b8b528b519a8b9003600b018b208852600490810186528b882054998b528b519a8b90039098018a2087529484528986205486528352888520670746f74616c5469760c41b8952895198899003600801892086529095018252878420548354601f600260001961010060018516150201909216919091049081018490048402890184019099528888529398899889986060988a98899894979596909594919391929185918301828280156113b75780601f1061138c576101008083540402835291602001916113b7565b820191906000526020600020905b81548152906001019060200180831161139a57829003601f168201915b5050505050925095509550955095509550955091939550919395565b600082815260448401602090815260408083206001600160a01b038516845260060190915290205460ff169392505050565b60008060008060008060006114186119b7565b5050506000868152604488016020908152604080832080546002820154600380840154600485015486516101208101808952681c995c5d595cdd125960ba1b905287518082036101290190208a526005808801808b52898c20548352895168074696d657374616d760bc1b81528a519081900360099081019091208d52818c528a8d2054848d01528a516476616c756560d81b81528b51908190039093019092208c52808b52898c2054838b015289516f6d696e457865637574696f6e4461746560801b81528a519081900360100190208c52808b52898c2054606084015289516c6e756d6265724f66566f74657360981b81528a5190819003600d0190208c52808b52898c2054608084015289516a313637b1b5a73ab6b132b960a91b81528a5190819003600b0190208c52808b52898c205460a08401528951681b5a5b995c94db1bdd60ba1b81528a51908190039092019091208b52808a52888b205460c083015288516571756f72756d60d01b815289519081900360060190208b52808a52888b205460e083015288516266656560e81b81528951908190039095019094208a5292909752949096205461010087810191909152600190930154919a5060ff8082169a509281048316985062010000810490921696506001600160a01b0363010000009092048216955091811693921691909295985092959850929598565b6116226119d6565b6040805161066081019182905290600184019060339082845b81548152602001906001019080831161163b5750505050509050919050565b60009182526048929092016020908152604080832093835260059093019052205490565b6000908152604a91909101602052604090205490565b60009182526048929092016020908152604080832093835260049093019052205490565b6000818152604883016020908152604080832060028082015483516a6772616e756c617269747960a81b8152845190819003600b018120875260048401808752858820546f3932b8bab2b9ba28a837b9b4ba34b7b760811b83528651928390036010018320895281885286892054670746f74616c5469760c41b845287519384900360080184208a52918852868920548654601f6000196101006001848116159190910291909101909216979097049687018a90048a0285018a019098528584526060998a999098899889988998919788979388019692959294909188918301828280156117e75780601f106117bc576101008083540402835291602001916117e7565b820191906000526020600020905b8154815290600101906020018083116117ca57829003601f168201915b5050885460408051602060026001851615610100026000190190941693909304601f8101849004840282018401909252818152959b508a9450925084019050828280156118755780601f1061184a57610100808354040283529160200191611875565b820191906000526020600020905b81548152906001019060200180831161185857829003601f168201915b50505050509450965096509650965096509650509295509295509295565b604080517174696d654f664c6173744e657756616c756560701b80825282516012928190038301812060009081528585016020818152868320548352604288018152868320549484528651938490039095019092208152925291812054909182916118ff9185916111f1565b9360019350915050565b604080516106608101918290526000918291829161194a9190600187019060339082845b81548152602001906001019080831161192d575050505050611965565b60009081526043909501602052505060409092205492915050565b60008060005b603381101561199357602081028401518084101561198a578093508192505b5060010161196b565b50915091565b6040518060a001604052806005906020820280388339509192915050565b6040518061012001604052806009906020820280388339509192915050565b604051806106600160405280603390602082028038833950919291505056fea265627a7a72305820cc7cdcaa2fea4ffcef4610ccbe2f219baa0c95dc886d7d996fc8219f373ffd7364736f6c634300050a0032`

// DeployBerryGetters deploys a new Ethereum contract, binding an instance of BerryGetters to it.
func DeployBerryGetters(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *BerryGetters, error) {
	parsed, err := abi.JSON(strings.NewReader(BerryGettersABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(BerryGettersBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &BerryGetters{BerryGettersCaller: BerryGettersCaller{contract: contract}, BerryGettersTransactor: BerryGettersTransactor{contract: contract}, BerryGettersFilterer: BerryGettersFilterer{contract: contract}}, nil
}

// BerryGetters is an auto generated Go binding around an Ethereum contract.
type BerryGetters struct {
	BerryGettersCaller     // Read-only binding to the contract
	BerryGettersTransactor // Write-only binding to the contract
	BerryGettersFilterer   // Log filterer for contract events
}

// BerryGettersCaller is an auto generated read-only Go binding around an Ethereum contract.
type BerryGettersCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BerryGettersTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BerryGettersTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BerryGettersFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BerryGettersFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BerryGettersSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BerryGettersSession struct {
	Contract     *BerryGetters    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BerryGettersCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BerryGettersCallerSession struct {
	Contract *BerryGettersCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// BerryGettersTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BerryGettersTransactorSession struct {
	Contract     *BerryGettersTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// BerryGettersRaw is an auto generated low-level Go binding around an Ethereum contract.
type BerryGettersRaw struct {
	Contract *BerryGetters // Generic contract binding to access the raw methods on
}

// BerryGettersCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BerryGettersCallerRaw struct {
	Contract *BerryGettersCaller // Generic read-only contract binding to access the raw methods on
}

// BerryGettersTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BerryGettersTransactorRaw struct {
	Contract *BerryGettersTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBerryGetters creates a new instance of BerryGetters, bound to a specific deployed contract.
func NewBerryGetters(address common.Address, backend bind.ContractBackend) (*BerryGetters, error) {
	contract, err := bindBerryGetters(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BerryGetters{BerryGettersCaller: BerryGettersCaller{contract: contract}, BerryGettersTransactor: BerryGettersTransactor{contract: contract}, BerryGettersFilterer: BerryGettersFilterer{contract: contract}}, nil
}

// NewBerryGettersCaller creates a new read-only instance of BerryGetters, bound to a specific deployed contract.
func NewBerryGettersCaller(address common.Address, caller bind.ContractCaller) (*BerryGettersCaller, error) {
	contract, err := bindBerryGetters(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BerryGettersCaller{contract: contract}, nil
}

// NewBerryGettersTransactor creates a new write-only instance of BerryGetters, bound to a specific deployed contract.
func NewBerryGettersTransactor(address common.Address, transactor bind.ContractTransactor) (*BerryGettersTransactor, error) {
	contract, err := bindBerryGetters(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BerryGettersTransactor{contract: contract}, nil
}

// NewBerryGettersFilterer creates a new log filterer instance of BerryGetters, bound to a specific deployed contract.
func NewBerryGettersFilterer(address common.Address, filterer bind.ContractFilterer) (*BerryGettersFilterer, error) {
	contract, err := bindBerryGetters(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BerryGettersFilterer{contract: contract}, nil
}

// bindBerryGetters binds a generic wrapper to an already deployed contract.
func bindBerryGetters(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BerryGettersABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BerryGetters *BerryGettersRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _BerryGetters.Contract.BerryGettersCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BerryGetters *BerryGettersRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BerryGetters.Contract.BerryGettersTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BerryGetters *BerryGettersRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BerryGetters.Contract.BerryGettersTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BerryGetters *BerryGettersCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _BerryGetters.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BerryGetters *BerryGettersTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BerryGetters.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BerryGetters *BerryGettersTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BerryGetters.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(_user address, _spender address) constant returns(uint256)
func (_BerryGetters *BerryGettersCaller) Allowance(opts *bind.CallOpts, _user common.Address, _spender common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BerryGetters.contract.Call(opts, out, "allowance", _user, _spender)
	return *ret0, err
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(_user address, _spender address) constant returns(uint256)
func (_BerryGetters *BerryGettersSession) Allowance(_user common.Address, _spender common.Address) (*big.Int, error) {
	return _BerryGetters.Contract.Allowance(&_BerryGetters.CallOpts, _user, _spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(_user address, _spender address) constant returns(uint256)
func (_BerryGetters *BerryGettersCallerSession) Allowance(_user common.Address, _spender common.Address) (*big.Int, error) {
	return _BerryGetters.Contract.Allowance(&_BerryGetters.CallOpts, _user, _spender)
}

// AllowedToTrade is a free data retrieval call binding the contract method 0x999cf26c.
//
// Solidity: function allowedToTrade(_user address, _amount uint256) constant returns(bool)
func (_BerryGetters *BerryGettersCaller) AllowedToTrade(opts *bind.CallOpts, _user common.Address, _amount *big.Int) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _BerryGetters.contract.Call(opts, out, "allowedToTrade", _user, _amount)
	return *ret0, err
}

// AllowedToTrade is a free data retrieval call binding the contract method 0x999cf26c.
//
// Solidity: function allowedToTrade(_user address, _amount uint256) constant returns(bool)
func (_BerryGetters *BerryGettersSession) AllowedToTrade(_user common.Address, _amount *big.Int) (bool, error) {
	return _BerryGetters.Contract.AllowedToTrade(&_BerryGetters.CallOpts, _user, _amount)
}

// AllowedToTrade is a free data retrieval call binding the contract method 0x999cf26c.
//
// Solidity: function allowedToTrade(_user address, _amount uint256) constant returns(bool)
func (_BerryGetters *BerryGettersCallerSession) AllowedToTrade(_user common.Address, _amount *big.Int) (bool, error) {
	return _BerryGetters.Contract.AllowedToTrade(&_BerryGetters.CallOpts, _user, _amount)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_user address) constant returns(uint256)
func (_BerryGetters *BerryGettersCaller) BalanceOf(opts *bind.CallOpts, _user common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BerryGetters.contract.Call(opts, out, "balanceOf", _user)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_user address) constant returns(uint256)
func (_BerryGetters *BerryGettersSession) BalanceOf(_user common.Address) (*big.Int, error) {
	return _BerryGetters.Contract.BalanceOf(&_BerryGetters.CallOpts, _user)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_user address) constant returns(uint256)
func (_BerryGetters *BerryGettersCallerSession) BalanceOf(_user common.Address) (*big.Int, error) {
	return _BerryGetters.Contract.BalanceOf(&_BerryGetters.CallOpts, _user)
}

// BalanceOfAt is a free data retrieval call binding the contract method 0x4ee2cd7e.
//
// Solidity: function balanceOfAt(_user address, _blockNumber uint256) constant returns(uint256)
func (_BerryGetters *BerryGettersCaller) BalanceOfAt(opts *bind.CallOpts, _user common.Address, _blockNumber *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BerryGetters.contract.Call(opts, out, "balanceOfAt", _user, _blockNumber)
	return *ret0, err
}

// BalanceOfAt is a free data retrieval call binding the contract method 0x4ee2cd7e.
//
// Solidity: function balanceOfAt(_user address, _blockNumber uint256) constant returns(uint256)
func (_BerryGetters *BerryGettersSession) BalanceOfAt(_user common.Address, _blockNumber *big.Int) (*big.Int, error) {
	return _BerryGetters.Contract.BalanceOfAt(&_BerryGetters.CallOpts, _user, _blockNumber)
}

// BalanceOfAt is a free data retrieval call binding the contract method 0x4ee2cd7e.
//
// Solidity: function balanceOfAt(_user address, _blockNumber uint256) constant returns(uint256)
func (_BerryGetters *BerryGettersCallerSession) BalanceOfAt(_user common.Address, _blockNumber *big.Int) (*big.Int, error) {
	return _BerryGetters.Contract.BalanceOfAt(&_BerryGetters.CallOpts, _user, _blockNumber)
}

// DidMine is a free data retrieval call binding the contract method 0x63bb82ad.
//
// Solidity: function didMine(_challenge bytes32, _miner address) constant returns(bool)
func (_BerryGetters *BerryGettersCaller) DidMine(opts *bind.CallOpts, _challenge [32]byte, _miner common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _BerryGetters.contract.Call(opts, out, "didMine", _challenge, _miner)
	return *ret0, err
}

// DidMine is a free data retrieval call binding the contract method 0x63bb82ad.
//
// Solidity: function didMine(_challenge bytes32, _miner address) constant returns(bool)
func (_BerryGetters *BerryGettersSession) DidMine(_challenge [32]byte, _miner common.Address) (bool, error) {
	return _BerryGetters.Contract.DidMine(&_BerryGetters.CallOpts, _challenge, _miner)
}

// DidMine is a free data retrieval call binding the contract method 0x63bb82ad.
//
// Solidity: function didMine(_challenge bytes32, _miner address) constant returns(bool)
func (_BerryGetters *BerryGettersCallerSession) DidMine(_challenge [32]byte, _miner common.Address) (bool, error) {
	return _BerryGetters.Contract.DidMine(&_BerryGetters.CallOpts, _challenge, _miner)
}

// DidVote is a free data retrieval call binding the contract method 0xa7c438bc.
//
// Solidity: function didVote(_disputeId uint256, _address address) constant returns(bool)
func (_BerryGetters *BerryGettersCaller) DidVote(opts *bind.CallOpts, _disputeId *big.Int, _address common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _BerryGetters.contract.Call(opts, out, "didVote", _disputeId, _address)
	return *ret0, err
}

// DidVote is a free data retrieval call binding the contract method 0xa7c438bc.
//
// Solidity: function didVote(_disputeId uint256, _address address) constant returns(bool)
func (_BerryGetters *BerryGettersSession) DidVote(_disputeId *big.Int, _address common.Address) (bool, error) {
	return _BerryGetters.Contract.DidVote(&_BerryGetters.CallOpts, _disputeId, _address)
}

// DidVote is a free data retrieval call binding the contract method 0xa7c438bc.
//
// Solidity: function didVote(_disputeId uint256, _address address) constant returns(bool)
func (_BerryGetters *BerryGettersCallerSession) DidVote(_disputeId *big.Int, _address common.Address) (bool, error) {
	return _BerryGetters.Contract.DidVote(&_BerryGetters.CallOpts, _disputeId, _address)
}

// GetAddressVars is a free data retrieval call binding the contract method 0x133bee5e.
//
// Solidity: function getAddressVars(_data bytes32) constant returns(address)
func (_BerryGetters *BerryGettersCaller) GetAddressVars(opts *bind.CallOpts, _data [32]byte) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _BerryGetters.contract.Call(opts, out, "getAddressVars", _data)
	return *ret0, err
}

// GetAddressVars is a free data retrieval call binding the contract method 0x133bee5e.
//
// Solidity: function getAddressVars(_data bytes32) constant returns(address)
func (_BerryGetters *BerryGettersSession) GetAddressVars(_data [32]byte) (common.Address, error) {
	return _BerryGetters.Contract.GetAddressVars(&_BerryGetters.CallOpts, _data)
}

// GetAddressVars is a free data retrieval call binding the contract method 0x133bee5e.
//
// Solidity: function getAddressVars(_data bytes32) constant returns(address)
func (_BerryGetters *BerryGettersCallerSession) GetAddressVars(_data [32]byte) (common.Address, error) {
	return _BerryGetters.Contract.GetAddressVars(&_BerryGetters.CallOpts, _data)
}

// GetAllDisputeVars is a free data retrieval call binding the contract method 0xaf0b1327.
//
// Solidity: function getAllDisputeVars(_disputeId uint256) constant returns(bytes32, bool, bool, bool, address, address, address, uint256[9], int256)
func (_BerryGetters *BerryGettersCaller) GetAllDisputeVars(opts *bind.CallOpts, _disputeId *big.Int) ([32]byte, bool, bool, bool, common.Address, common.Address, common.Address, [9]*big.Int, *big.Int, error) {
	var (
		ret0 = new([32]byte)
		ret1 = new(bool)
		ret2 = new(bool)
		ret3 = new(bool)
		ret4 = new(common.Address)
		ret5 = new(common.Address)
		ret6 = new(common.Address)
		ret7 = new([9]*big.Int)
		ret8 = new(*big.Int)
	)
	out := &[]interface{}{
		ret0,
		ret1,
		ret2,
		ret3,
		ret4,
		ret5,
		ret6,
		ret7,
		ret8,
	}
	err := _BerryGetters.contract.Call(opts, out, "getAllDisputeVars", _disputeId)
	return *ret0, *ret1, *ret2, *ret3, *ret4, *ret5, *ret6, *ret7, *ret8, err
}

// GetAllDisputeVars is a free data retrieval call binding the contract method 0xaf0b1327.
//
// Solidity: function getAllDisputeVars(_disputeId uint256) constant returns(bytes32, bool, bool, bool, address, address, address, uint256[9], int256)
func (_BerryGetters *BerryGettersSession) GetAllDisputeVars(_disputeId *big.Int) ([32]byte, bool, bool, bool, common.Address, common.Address, common.Address, [9]*big.Int, *big.Int, error) {
	return _BerryGetters.Contract.GetAllDisputeVars(&_BerryGetters.CallOpts, _disputeId)
}

// GetAllDisputeVars is a free data retrieval call binding the contract method 0xaf0b1327.
//
// Solidity: function getAllDisputeVars(_disputeId uint256) constant returns(bytes32, bool, bool, bool, address, address, address, uint256[9], int256)
func (_BerryGetters *BerryGettersCallerSession) GetAllDisputeVars(_disputeId *big.Int) ([32]byte, bool, bool, bool, common.Address, common.Address, common.Address, [9]*big.Int, *big.Int, error) {
	return _BerryGetters.Contract.GetAllDisputeVars(&_BerryGetters.CallOpts, _disputeId)
}

// GetCurrentVariables is a free data retrieval call binding the contract method 0xa22e407a.
//
// Solidity: function getCurrentVariables() constant returns(bytes32, uint256, uint256, string, uint256, uint256)
func (_BerryGetters *BerryGettersCaller) GetCurrentVariables(opts *bind.CallOpts) ([32]byte, *big.Int, *big.Int, string, *big.Int, *big.Int, error) {
	var (
		ret0 = new([32]byte)
		ret1 = new(*big.Int)
		ret2 = new(*big.Int)
		ret3 = new(string)
		ret4 = new(*big.Int)
		ret5 = new(*big.Int)
	)
	out := &[]interface{}{
		ret0,
		ret1,
		ret2,
		ret3,
		ret4,
		ret5,
	}
	err := _BerryGetters.contract.Call(opts, out, "getCurrentVariables")
	return *ret0, *ret1, *ret2, *ret3, *ret4, *ret5, err
}

// GetCurrentVariables is a free data retrieval call binding the contract method 0xa22e407a.
//
// Solidity: function getCurrentVariables() constant returns(bytes32, uint256, uint256, string, uint256, uint256)
func (_BerryGetters *BerryGettersSession) GetCurrentVariables() ([32]byte, *big.Int, *big.Int, string, *big.Int, *big.Int, error) {
	return _BerryGetters.Contract.GetCurrentVariables(&_BerryGetters.CallOpts)
}

// GetCurrentVariables is a free data retrieval call binding the contract method 0xa22e407a.
//
// Solidity: function getCurrentVariables() constant returns(bytes32, uint256, uint256, string, uint256, uint256)
func (_BerryGetters *BerryGettersCallerSession) GetCurrentVariables() ([32]byte, *big.Int, *big.Int, string, *big.Int, *big.Int, error) {
	return _BerryGetters.Contract.GetCurrentVariables(&_BerryGetters.CallOpts)
}

// GetDisputeIdByDisputeHash is a free data retrieval call binding the contract method 0xda379941.
//
// Solidity: function getDisputeIdByDisputeHash(_hash bytes32) constant returns(uint256)
func (_BerryGetters *BerryGettersCaller) GetDisputeIdByDisputeHash(opts *bind.CallOpts, _hash [32]byte) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BerryGetters.contract.Call(opts, out, "getDisputeIdByDisputeHash", _hash)
	return *ret0, err
}

// GetDisputeIdByDisputeHash is a free data retrieval call binding the contract method 0xda379941.
//
// Solidity: function getDisputeIdByDisputeHash(_hash bytes32) constant returns(uint256)
func (_BerryGetters *BerryGettersSession) GetDisputeIdByDisputeHash(_hash [32]byte) (*big.Int, error) {
	return _BerryGetters.Contract.GetDisputeIdByDisputeHash(&_BerryGetters.CallOpts, _hash)
}

// GetDisputeIdByDisputeHash is a free data retrieval call binding the contract method 0xda379941.
//
// Solidity: function getDisputeIdByDisputeHash(_hash bytes32) constant returns(uint256)
func (_BerryGetters *BerryGettersCallerSession) GetDisputeIdByDisputeHash(_hash [32]byte) (*big.Int, error) {
	return _BerryGetters.Contract.GetDisputeIdByDisputeHash(&_BerryGetters.CallOpts, _hash)
}

// GetDisputeUintVars is a free data retrieval call binding the contract method 0x7f6fd5d9.
//
// Solidity: function getDisputeUintVars(_disputeId uint256, _data bytes32) constant returns(uint256)
func (_BerryGetters *BerryGettersCaller) GetDisputeUintVars(opts *bind.CallOpts, _disputeId *big.Int, _data [32]byte) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BerryGetters.contract.Call(opts, out, "getDisputeUintVars", _disputeId, _data)
	return *ret0, err
}

// GetDisputeUintVars is a free data retrieval call binding the contract method 0x7f6fd5d9.
//
// Solidity: function getDisputeUintVars(_disputeId uint256, _data bytes32) constant returns(uint256)
func (_BerryGetters *BerryGettersSession) GetDisputeUintVars(_disputeId *big.Int, _data [32]byte) (*big.Int, error) {
	return _BerryGetters.Contract.GetDisputeUintVars(&_BerryGetters.CallOpts, _disputeId, _data)
}

// GetDisputeUintVars is a free data retrieval call binding the contract method 0x7f6fd5d9.
//
// Solidity: function getDisputeUintVars(_disputeId uint256, _data bytes32) constant returns(uint256)
func (_BerryGetters *BerryGettersCallerSession) GetDisputeUintVars(_disputeId *big.Int, _data [32]byte) (*big.Int, error) {
	return _BerryGetters.Contract.GetDisputeUintVars(&_BerryGetters.CallOpts, _disputeId, _data)
}

// GetLastNewValue is a free data retrieval call binding the contract method 0xfc7cf0a0.
//
// Solidity: function getLastNewValue() constant returns(uint256, bool)
func (_BerryGetters *BerryGettersCaller) GetLastNewValue(opts *bind.CallOpts) (*big.Int, bool, error) {
	var (
		ret0 = new(*big.Int)
		ret1 = new(bool)
	)
	out := &[]interface{}{
		ret0,
		ret1,
	}
	err := _BerryGetters.contract.Call(opts, out, "getLastNewValue")
	return *ret0, *ret1, err
}

// GetLastNewValue is a free data retrieval call binding the contract method 0xfc7cf0a0.
//
// Solidity: function getLastNewValue() constant returns(uint256, bool)
func (_BerryGetters *BerryGettersSession) GetLastNewValue() (*big.Int, bool, error) {
	return _BerryGetters.Contract.GetLastNewValue(&_BerryGetters.CallOpts)
}

// GetLastNewValue is a free data retrieval call binding the contract method 0xfc7cf0a0.
//
// Solidity: function getLastNewValue() constant returns(uint256, bool)
func (_BerryGetters *BerryGettersCallerSession) GetLastNewValue() (*big.Int, bool, error) {
	return _BerryGetters.Contract.GetLastNewValue(&_BerryGetters.CallOpts)
}

// GetLastNewValueById is a free data retrieval call binding the contract method 0x3180f8df.
//
// Solidity: function getLastNewValueById(_requestId uint256) constant returns(uint256, bool)
func (_BerryGetters *BerryGettersCaller) GetLastNewValueById(opts *bind.CallOpts, _requestId *big.Int) (*big.Int, bool, error) {
	var (
		ret0 = new(*big.Int)
		ret1 = new(bool)
	)
	out := &[]interface{}{
		ret0,
		ret1,
	}
	err := _BerryGetters.contract.Call(opts, out, "getLastNewValueById", _requestId)
	return *ret0, *ret1, err
}

// GetLastNewValueById is a free data retrieval call binding the contract method 0x3180f8df.
//
// Solidity: function getLastNewValueById(_requestId uint256) constant returns(uint256, bool)
func (_BerryGetters *BerryGettersSession) GetLastNewValueById(_requestId *big.Int) (*big.Int, bool, error) {
	return _BerryGetters.Contract.GetLastNewValueById(&_BerryGetters.CallOpts, _requestId)
}

// GetLastNewValueById is a free data retrieval call binding the contract method 0x3180f8df.
//
// Solidity: function getLastNewValueById(_requestId uint256) constant returns(uint256, bool)
func (_BerryGetters *BerryGettersCallerSession) GetLastNewValueById(_requestId *big.Int) (*big.Int, bool, error) {
	return _BerryGetters.Contract.GetLastNewValueById(&_BerryGetters.CallOpts, _requestId)
}

// GetMinedBlockNum is a free data retrieval call binding the contract method 0xc775b542.
//
// Solidity: function getMinedBlockNum(_requestId uint256, _timestamp uint256) constant returns(uint256)
func (_BerryGetters *BerryGettersCaller) GetMinedBlockNum(opts *bind.CallOpts, _requestId *big.Int, _timestamp *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BerryGetters.contract.Call(opts, out, "getMinedBlockNum", _requestId, _timestamp)
	return *ret0, err
}

// GetMinedBlockNum is a free data retrieval call binding the contract method 0xc775b542.
//
// Solidity: function getMinedBlockNum(_requestId uint256, _timestamp uint256) constant returns(uint256)
func (_BerryGetters *BerryGettersSession) GetMinedBlockNum(_requestId *big.Int, _timestamp *big.Int) (*big.Int, error) {
	return _BerryGetters.Contract.GetMinedBlockNum(&_BerryGetters.CallOpts, _requestId, _timestamp)
}

// GetMinedBlockNum is a free data retrieval call binding the contract method 0xc775b542.
//
// Solidity: function getMinedBlockNum(_requestId uint256, _timestamp uint256) constant returns(uint256)
func (_BerryGetters *BerryGettersCallerSession) GetMinedBlockNum(_requestId *big.Int, _timestamp *big.Int) (*big.Int, error) {
	return _BerryGetters.Contract.GetMinedBlockNum(&_BerryGetters.CallOpts, _requestId, _timestamp)
}

// GetMinersByRequestIdAndTimestamp is a free data retrieval call binding the contract method 0x69026d63.
//
// Solidity: function getMinersByRequestIdAndTimestamp(_requestId uint256, _timestamp uint256) constant returns(address[5])
func (_BerryGetters *BerryGettersCaller) GetMinersByRequestIdAndTimestamp(opts *bind.CallOpts, _requestId *big.Int, _timestamp *big.Int) ([5]common.Address, error) {
	var (
		ret0 = new([5]common.Address)
	)
	out := ret0
	err := _BerryGetters.contract.Call(opts, out, "getMinersByRequestIdAndTimestamp", _requestId, _timestamp)
	return *ret0, err
}

// GetMinersByRequestIdAndTimestamp is a free data retrieval call binding the contract method 0x69026d63.
//
// Solidity: function getMinersByRequestIdAndTimestamp(_requestId uint256, _timestamp uint256) constant returns(address[5])
func (_BerryGetters *BerryGettersSession) GetMinersByRequestIdAndTimestamp(_requestId *big.Int, _timestamp *big.Int) ([5]common.Address, error) {
	return _BerryGetters.Contract.GetMinersByRequestIdAndTimestamp(&_BerryGetters.CallOpts, _requestId, _timestamp)
}

// GetMinersByRequestIdAndTimestamp is a free data retrieval call binding the contract method 0x69026d63.
//
// Solidity: function getMinersByRequestIdAndTimestamp(_requestId uint256, _timestamp uint256) constant returns(address[5])
func (_BerryGetters *BerryGettersCallerSession) GetMinersByRequestIdAndTimestamp(_requestId *big.Int, _timestamp *big.Int) ([5]common.Address, error) {
	return _BerryGetters.Contract.GetMinersByRequestIdAndTimestamp(&_BerryGetters.CallOpts, _requestId, _timestamp)
}

// GetName is a free data retrieval call binding the contract method 0x17d7de7c.
//
// Solidity: function getName() constant returns(string)
func (_BerryGetters *BerryGettersCaller) GetName(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _BerryGetters.contract.Call(opts, out, "getName")
	return *ret0, err
}

// GetName is a free data retrieval call binding the contract method 0x17d7de7c.
//
// Solidity: function getName() constant returns(string)
func (_BerryGetters *BerryGettersSession) GetName() (string, error) {
	return _BerryGetters.Contract.GetName(&_BerryGetters.CallOpts)
}

// GetName is a free data retrieval call binding the contract method 0x17d7de7c.
//
// Solidity: function getName() constant returns(string)
func (_BerryGetters *BerryGettersCallerSession) GetName() (string, error) {
	return _BerryGetters.Contract.GetName(&_BerryGetters.CallOpts)
}

// GetNewValueCountbyRequestId is a free data retrieval call binding the contract method 0x46eee1c4.
//
// Solidity: function getNewValueCountbyRequestId(_requestId uint256) constant returns(uint256)
func (_BerryGetters *BerryGettersCaller) GetNewValueCountbyRequestId(opts *bind.CallOpts, _requestId *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BerryGetters.contract.Call(opts, out, "getNewValueCountbyRequestId", _requestId)
	return *ret0, err
}

// GetNewValueCountbyRequestId is a free data retrieval call binding the contract method 0x46eee1c4.
//
// Solidity: function getNewValueCountbyRequestId(_requestId uint256) constant returns(uint256)
func (_BerryGetters *BerryGettersSession) GetNewValueCountbyRequestId(_requestId *big.Int) (*big.Int, error) {
	return _BerryGetters.Contract.GetNewValueCountbyRequestId(&_BerryGetters.CallOpts, _requestId)
}

// GetNewValueCountbyRequestId is a free data retrieval call binding the contract method 0x46eee1c4.
//
// Solidity: function getNewValueCountbyRequestId(_requestId uint256) constant returns(uint256)
func (_BerryGetters *BerryGettersCallerSession) GetNewValueCountbyRequestId(_requestId *big.Int) (*big.Int, error) {
	return _BerryGetters.Contract.GetNewValueCountbyRequestId(&_BerryGetters.CallOpts, _requestId)
}

// GetRequestIdByQueryHash is a free data retrieval call binding the contract method 0x1db842f0.
//
// Solidity: function getRequestIdByQueryHash(_request bytes32) constant returns(uint256)
func (_BerryGetters *BerryGettersCaller) GetRequestIdByQueryHash(opts *bind.CallOpts, _request [32]byte) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BerryGetters.contract.Call(opts, out, "getRequestIdByQueryHash", _request)
	return *ret0, err
}

// GetRequestIdByQueryHash is a free data retrieval call binding the contract method 0x1db842f0.
//
// Solidity: function getRequestIdByQueryHash(_request bytes32) constant returns(uint256)
func (_BerryGetters *BerryGettersSession) GetRequestIdByQueryHash(_request [32]byte) (*big.Int, error) {
	return _BerryGetters.Contract.GetRequestIdByQueryHash(&_BerryGetters.CallOpts, _request)
}

// GetRequestIdByQueryHash is a free data retrieval call binding the contract method 0x1db842f0.
//
// Solidity: function getRequestIdByQueryHash(_request bytes32) constant returns(uint256)
func (_BerryGetters *BerryGettersCallerSession) GetRequestIdByQueryHash(_request [32]byte) (*big.Int, error) {
	return _BerryGetters.Contract.GetRequestIdByQueryHash(&_BerryGetters.CallOpts, _request)
}

// GetRequestIdByRequestQIndex is a free data retrieval call binding the contract method 0x6173c0b8.
//
// Solidity: function getRequestIdByRequestQIndex(_index uint256) constant returns(uint256)
func (_BerryGetters *BerryGettersCaller) GetRequestIdByRequestQIndex(opts *bind.CallOpts, _index *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BerryGetters.contract.Call(opts, out, "getRequestIdByRequestQIndex", _index)
	return *ret0, err
}

// GetRequestIdByRequestQIndex is a free data retrieval call binding the contract method 0x6173c0b8.
//
// Solidity: function getRequestIdByRequestQIndex(_index uint256) constant returns(uint256)
func (_BerryGetters *BerryGettersSession) GetRequestIdByRequestQIndex(_index *big.Int) (*big.Int, error) {
	return _BerryGetters.Contract.GetRequestIdByRequestQIndex(&_BerryGetters.CallOpts, _index)
}

// GetRequestIdByRequestQIndex is a free data retrieval call binding the contract method 0x6173c0b8.
//
// Solidity: function getRequestIdByRequestQIndex(_index uint256) constant returns(uint256)
func (_BerryGetters *BerryGettersCallerSession) GetRequestIdByRequestQIndex(_index *big.Int) (*big.Int, error) {
	return _BerryGetters.Contract.GetRequestIdByRequestQIndex(&_BerryGetters.CallOpts, _index)
}

// GetRequestIdByTimestamp is a free data retrieval call binding the contract method 0x0f0b424d.
//
// Solidity: function getRequestIdByTimestamp(_timestamp uint256) constant returns(uint256)
func (_BerryGetters *BerryGettersCaller) GetRequestIdByTimestamp(opts *bind.CallOpts, _timestamp *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BerryGetters.contract.Call(opts, out, "getRequestIdByTimestamp", _timestamp)
	return *ret0, err
}

// GetRequestIdByTimestamp is a free data retrieval call binding the contract method 0x0f0b424d.
//
// Solidity: function getRequestIdByTimestamp(_timestamp uint256) constant returns(uint256)
func (_BerryGetters *BerryGettersSession) GetRequestIdByTimestamp(_timestamp *big.Int) (*big.Int, error) {
	return _BerryGetters.Contract.GetRequestIdByTimestamp(&_BerryGetters.CallOpts, _timestamp)
}

// GetRequestIdByTimestamp is a free data retrieval call binding the contract method 0x0f0b424d.
//
// Solidity: function getRequestIdByTimestamp(_timestamp uint256) constant returns(uint256)
func (_BerryGetters *BerryGettersCallerSession) GetRequestIdByTimestamp(_timestamp *big.Int) (*big.Int, error) {
	return _BerryGetters.Contract.GetRequestIdByTimestamp(&_BerryGetters.CallOpts, _timestamp)
}

// GetRequestQ is a free data retrieval call binding the contract method 0xb5413029.
//
// Solidity: function getRequestQ() constant returns(uint256[51])
func (_BerryGetters *BerryGettersCaller) GetRequestQ(opts *bind.CallOpts) ([51]*big.Int, error) {
	var (
		ret0 = new([51]*big.Int)
	)
	out := ret0
	err := _BerryGetters.contract.Call(opts, out, "getRequestQ")
	return *ret0, err
}

// GetRequestQ is a free data retrieval call binding the contract method 0xb5413029.
//
// Solidity: function getRequestQ() constant returns(uint256[51])
func (_BerryGetters *BerryGettersSession) GetRequestQ() ([51]*big.Int, error) {
	return _BerryGetters.Contract.GetRequestQ(&_BerryGetters.CallOpts)
}

// GetRequestQ is a free data retrieval call binding the contract method 0xb5413029.
//
// Solidity: function getRequestQ() constant returns(uint256[51])
func (_BerryGetters *BerryGettersCallerSession) GetRequestQ() ([51]*big.Int, error) {
	return _BerryGetters.Contract.GetRequestQ(&_BerryGetters.CallOpts)
}

// GetRequestUintVars is a free data retrieval call binding the contract method 0xe0ae93c1.
//
// Solidity: function getRequestUintVars(_requestId uint256, _data bytes32) constant returns(uint256)
func (_BerryGetters *BerryGettersCaller) GetRequestUintVars(opts *bind.CallOpts, _requestId *big.Int, _data [32]byte) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BerryGetters.contract.Call(opts, out, "getRequestUintVars", _requestId, _data)
	return *ret0, err
}

// GetRequestUintVars is a free data retrieval call binding the contract method 0xe0ae93c1.
//
// Solidity: function getRequestUintVars(_requestId uint256, _data bytes32) constant returns(uint256)
func (_BerryGetters *BerryGettersSession) GetRequestUintVars(_requestId *big.Int, _data [32]byte) (*big.Int, error) {
	return _BerryGetters.Contract.GetRequestUintVars(&_BerryGetters.CallOpts, _requestId, _data)
}

// GetRequestUintVars is a free data retrieval call binding the contract method 0xe0ae93c1.
//
// Solidity: function getRequestUintVars(_requestId uint256, _data bytes32) constant returns(uint256)
func (_BerryGetters *BerryGettersCallerSession) GetRequestUintVars(_requestId *big.Int, _data [32]byte) (*big.Int, error) {
	return _BerryGetters.Contract.GetRequestUintVars(&_BerryGetters.CallOpts, _requestId, _data)
}

// GetRequestVars is a free data retrieval call binding the contract method 0xe1eee6d6.
//
// Solidity: function getRequestVars(_requestId uint256) constant returns(string, string, bytes32, uint256, uint256, uint256)
func (_BerryGetters *BerryGettersCaller) GetRequestVars(opts *bind.CallOpts, _requestId *big.Int) (string, string, [32]byte, *big.Int, *big.Int, *big.Int, error) {
	var (
		ret0 = new(string)
		ret1 = new(string)
		ret2 = new([32]byte)
		ret3 = new(*big.Int)
		ret4 = new(*big.Int)
		ret5 = new(*big.Int)
	)
	out := &[]interface{}{
		ret0,
		ret1,
		ret2,
		ret3,
		ret4,
		ret5,
	}
	err := _BerryGetters.contract.Call(opts, out, "getRequestVars", _requestId)
	return *ret0, *ret1, *ret2, *ret3, *ret4, *ret5, err
}

// GetRequestVars is a free data retrieval call binding the contract method 0xe1eee6d6.
//
// Solidity: function getRequestVars(_requestId uint256) constant returns(string, string, bytes32, uint256, uint256, uint256)
func (_BerryGetters *BerryGettersSession) GetRequestVars(_requestId *big.Int) (string, string, [32]byte, *big.Int, *big.Int, *big.Int, error) {
	return _BerryGetters.Contract.GetRequestVars(&_BerryGetters.CallOpts, _requestId)
}

// GetRequestVars is a free data retrieval call binding the contract method 0xe1eee6d6.
//
// Solidity: function getRequestVars(_requestId uint256) constant returns(string, string, bytes32, uint256, uint256, uint256)
func (_BerryGetters *BerryGettersCallerSession) GetRequestVars(_requestId *big.Int) (string, string, [32]byte, *big.Int, *big.Int, *big.Int, error) {
	return _BerryGetters.Contract.GetRequestVars(&_BerryGetters.CallOpts, _requestId)
}

// GetStakerInfo is a free data retrieval call binding the contract method 0x733bdef0.
//
// Solidity: function getStakerInfo(_staker address) constant returns(uint256, uint256)
func (_BerryGetters *BerryGettersCaller) GetStakerInfo(opts *bind.CallOpts, _staker common.Address) (*big.Int, *big.Int, error) {
	var (
		ret0 = new(*big.Int)
		ret1 = new(*big.Int)
	)
	out := &[]interface{}{
		ret0,
		ret1,
	}
	err := _BerryGetters.contract.Call(opts, out, "getStakerInfo", _staker)
	return *ret0, *ret1, err
}

// GetStakerInfo is a free data retrieval call binding the contract method 0x733bdef0.
//
// Solidity: function getStakerInfo(_staker address) constant returns(uint256, uint256)
func (_BerryGetters *BerryGettersSession) GetStakerInfo(_staker common.Address) (*big.Int, *big.Int, error) {
	return _BerryGetters.Contract.GetStakerInfo(&_BerryGetters.CallOpts, _staker)
}

// GetStakerInfo is a free data retrieval call binding the contract method 0x733bdef0.
//
// Solidity: function getStakerInfo(_staker address) constant returns(uint256, uint256)
func (_BerryGetters *BerryGettersCallerSession) GetStakerInfo(_staker common.Address) (*big.Int, *big.Int, error) {
	return _BerryGetters.Contract.GetStakerInfo(&_BerryGetters.CallOpts, _staker)
}

// GetSubmissionsByTimestamp is a free data retrieval call binding the contract method 0x11c98512.
//
// Solidity: function getSubmissionsByTimestamp(_requestId uint256, _timestamp uint256) constant returns(uint256[5])
func (_BerryGetters *BerryGettersCaller) GetSubmissionsByTimestamp(opts *bind.CallOpts, _requestId *big.Int, _timestamp *big.Int) ([5]*big.Int, error) {
	var (
		ret0 = new([5]*big.Int)
	)
	out := ret0
	err := _BerryGetters.contract.Call(opts, out, "getSubmissionsByTimestamp", _requestId, _timestamp)
	return *ret0, err
}

// GetSubmissionsByTimestamp is a free data retrieval call binding the contract method 0x11c98512.
//
// Solidity: function getSubmissionsByTimestamp(_requestId uint256, _timestamp uint256) constant returns(uint256[5])
func (_BerryGetters *BerryGettersSession) GetSubmissionsByTimestamp(_requestId *big.Int, _timestamp *big.Int) ([5]*big.Int, error) {
	return _BerryGetters.Contract.GetSubmissionsByTimestamp(&_BerryGetters.CallOpts, _requestId, _timestamp)
}

// GetSubmissionsByTimestamp is a free data retrieval call binding the contract method 0x11c98512.
//
// Solidity: function getSubmissionsByTimestamp(_requestId uint256, _timestamp uint256) constant returns(uint256[5])
func (_BerryGetters *BerryGettersCallerSession) GetSubmissionsByTimestamp(_requestId *big.Int, _timestamp *big.Int) ([5]*big.Int, error) {
	return _BerryGetters.Contract.GetSubmissionsByTimestamp(&_BerryGetters.CallOpts, _requestId, _timestamp)
}

// GetSymbol is a free data retrieval call binding the contract method 0x15070401.
//
// Solidity: function getSymbol() constant returns(string)
func (_BerryGetters *BerryGettersCaller) GetSymbol(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _BerryGetters.contract.Call(opts, out, "getSymbol")
	return *ret0, err
}

// GetSymbol is a free data retrieval call binding the contract method 0x15070401.
//
// Solidity: function getSymbol() constant returns(string)
func (_BerryGetters *BerryGettersSession) GetSymbol() (string, error) {
	return _BerryGetters.Contract.GetSymbol(&_BerryGetters.CallOpts)
}

// GetSymbol is a free data retrieval call binding the contract method 0x15070401.
//
// Solidity: function getSymbol() constant returns(string)
func (_BerryGetters *BerryGettersCallerSession) GetSymbol() (string, error) {
	return _BerryGetters.Contract.GetSymbol(&_BerryGetters.CallOpts)
}

// GetTimestampbyRequestIDandIndex is a free data retrieval call binding the contract method 0x77fbb663.
//
// Solidity: function getTimestampbyRequestIDandIndex(_requestID uint256, _index uint256) constant returns(uint256)
func (_BerryGetters *BerryGettersCaller) GetTimestampbyRequestIDandIndex(opts *bind.CallOpts, _requestID *big.Int, _index *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BerryGetters.contract.Call(opts, out, "getTimestampbyRequestIDandIndex", _requestID, _index)
	return *ret0, err
}

// GetTimestampbyRequestIDandIndex is a free data retrieval call binding the contract method 0x77fbb663.
//
// Solidity: function getTimestampbyRequestIDandIndex(_requestID uint256, _index uint256) constant returns(uint256)
func (_BerryGetters *BerryGettersSession) GetTimestampbyRequestIDandIndex(_requestID *big.Int, _index *big.Int) (*big.Int, error) {
	return _BerryGetters.Contract.GetTimestampbyRequestIDandIndex(&_BerryGetters.CallOpts, _requestID, _index)
}

// GetTimestampbyRequestIDandIndex is a free data retrieval call binding the contract method 0x77fbb663.
//
// Solidity: function getTimestampbyRequestIDandIndex(_requestID uint256, _index uint256) constant returns(uint256)
func (_BerryGetters *BerryGettersCallerSession) GetTimestampbyRequestIDandIndex(_requestID *big.Int, _index *big.Int) (*big.Int, error) {
	return _BerryGetters.Contract.GetTimestampbyRequestIDandIndex(&_BerryGetters.CallOpts, _requestID, _index)
}

// GetUintVar is a free data retrieval call binding the contract method 0x612c8f7f.
//
// Solidity: function getUintVar(_data bytes32) constant returns(uint256)
func (_BerryGetters *BerryGettersCaller) GetUintVar(opts *bind.CallOpts, _data [32]byte) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BerryGetters.contract.Call(opts, out, "getUintVar", _data)
	return *ret0, err
}

// GetUintVar is a free data retrieval call binding the contract method 0x612c8f7f.
//
// Solidity: function getUintVar(_data bytes32) constant returns(uint256)
func (_BerryGetters *BerryGettersSession) GetUintVar(_data [32]byte) (*big.Int, error) {
	return _BerryGetters.Contract.GetUintVar(&_BerryGetters.CallOpts, _data)
}

// GetUintVar is a free data retrieval call binding the contract method 0x612c8f7f.
//
// Solidity: function getUintVar(_data bytes32) constant returns(uint256)
func (_BerryGetters *BerryGettersCallerSession) GetUintVar(_data [32]byte) (*big.Int, error) {
	return _BerryGetters.Contract.GetUintVar(&_BerryGetters.CallOpts, _data)
}

// GetVariablesOnDeck is a free data retrieval call binding the contract method 0x19e8e03b.
//
// Solidity: function getVariablesOnDeck() constant returns(uint256, uint256, string)
func (_BerryGetters *BerryGettersCaller) GetVariablesOnDeck(opts *bind.CallOpts) (*big.Int, *big.Int, string, error) {
	var (
		ret0 = new(*big.Int)
		ret1 = new(*big.Int)
		ret2 = new(string)
	)
	out := &[]interface{}{
		ret0,
		ret1,
		ret2,
	}
	err := _BerryGetters.contract.Call(opts, out, "getVariablesOnDeck")
	return *ret0, *ret1, *ret2, err
}

// GetVariablesOnDeck is a free data retrieval call binding the contract method 0x19e8e03b.
//
// Solidity: function getVariablesOnDeck() constant returns(uint256, uint256, string)
func (_BerryGetters *BerryGettersSession) GetVariablesOnDeck() (*big.Int, *big.Int, string, error) {
	return _BerryGetters.Contract.GetVariablesOnDeck(&_BerryGetters.CallOpts)
}

// GetVariablesOnDeck is a free data retrieval call binding the contract method 0x19e8e03b.
//
// Solidity: function getVariablesOnDeck() constant returns(uint256, uint256, string)
func (_BerryGetters *BerryGettersCallerSession) GetVariablesOnDeck() (*big.Int, *big.Int, string, error) {
	return _BerryGetters.Contract.GetVariablesOnDeck(&_BerryGetters.CallOpts)
}

// IsInDispute is a free data retrieval call binding the contract method 0x3df0777b.
//
// Solidity: function isInDispute(_requestId uint256, _timestamp uint256) constant returns(bool)
func (_BerryGetters *BerryGettersCaller) IsInDispute(opts *bind.CallOpts, _requestId *big.Int, _timestamp *big.Int) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _BerryGetters.contract.Call(opts, out, "isInDispute", _requestId, _timestamp)
	return *ret0, err
}

// IsInDispute is a free data retrieval call binding the contract method 0x3df0777b.
//
// Solidity: function isInDispute(_requestId uint256, _timestamp uint256) constant returns(bool)
func (_BerryGetters *BerryGettersSession) IsInDispute(_requestId *big.Int, _timestamp *big.Int) (bool, error) {
	return _BerryGetters.Contract.IsInDispute(&_BerryGetters.CallOpts, _requestId, _timestamp)
}

// IsInDispute is a free data retrieval call binding the contract method 0x3df0777b.
//
// Solidity: function isInDispute(_requestId uint256, _timestamp uint256) constant returns(bool)
func (_BerryGetters *BerryGettersCallerSession) IsInDispute(_requestId *big.Int, _timestamp *big.Int) (bool, error) {
	return _BerryGetters.Contract.IsInDispute(&_BerryGetters.CallOpts, _requestId, _timestamp)
}

// RetrieveData is a free data retrieval call binding the contract method 0x93fa4915.
//
// Solidity: function retrieveData(_requestId uint256, _timestamp uint256) constant returns(uint256)
func (_BerryGetters *BerryGettersCaller) RetrieveData(opts *bind.CallOpts, _requestId *big.Int, _timestamp *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BerryGetters.contract.Call(opts, out, "retrieveData", _requestId, _timestamp)
	return *ret0, err
}

// RetrieveData is a free data retrieval call binding the contract method 0x93fa4915.
//
// Solidity: function retrieveData(_requestId uint256, _timestamp uint256) constant returns(uint256)
func (_BerryGetters *BerryGettersSession) RetrieveData(_requestId *big.Int, _timestamp *big.Int) (*big.Int, error) {
	return _BerryGetters.Contract.RetrieveData(&_BerryGetters.CallOpts, _requestId, _timestamp)
}

// RetrieveData is a free data retrieval call binding the contract method 0x93fa4915.
//
// Solidity: function retrieveData(_requestId uint256, _timestamp uint256) constant returns(uint256)
func (_BerryGetters *BerryGettersCallerSession) RetrieveData(_requestId *big.Int, _timestamp *big.Int) (*big.Int, error) {
	return _BerryGetters.Contract.RetrieveData(&_BerryGetters.CallOpts, _requestId, _timestamp)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_BerryGetters *BerryGettersCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BerryGetters.contract.Call(opts, out, "totalSupply")
	return *ret0, err
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_BerryGetters *BerryGettersSession) TotalSupply() (*big.Int, error) {
	return _BerryGetters.Contract.TotalSupply(&_BerryGetters.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_BerryGetters *BerryGettersCallerSession) TotalSupply() (*big.Int, error) {
	return _BerryGetters.Contract.TotalSupply(&_BerryGetters.CallOpts)
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

// BerryMasterABI is the input ABI used to generate the binding from.
const BerryMasterABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"_timestamp\",\"type\":\"uint256\"}],\"name\":\"getRequestIdByTimestamp\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_requestId\",\"type\":\"uint256\"},{\"name\":\"_timestamp\",\"type\":\"uint256\"}],\"name\":\"getSubmissionsByTimestamp\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256[5]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_data\",\"type\":\"bytes32\"}],\"name\":\"getAddressVars\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getSymbol\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getName\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getVariablesOnDeck\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_request\",\"type\":\"bytes32\"}],\"name\":\"getRequestIdByQueryHash\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_requestId\",\"type\":\"uint256\"}],\"name\":\"getLastNewValueById\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_requestId\",\"type\":\"uint256\"},{\"name\":\"_timestamp\",\"type\":\"uint256\"}],\"name\":\"isInDispute\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_requestId\",\"type\":\"uint256\"}],\"name\":\"getNewValueCountbyRequestId\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_newDeity\",\"type\":\"address\"}],\"name\":\"changeDeity\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_user\",\"type\":\"address\"},{\"name\":\"_blockNumber\",\"type\":\"uint256\"}],\"name\":\"balanceOfAt\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_data\",\"type\":\"bytes32\"}],\"name\":\"getUintVar\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_index\",\"type\":\"uint256\"}],\"name\":\"getRequestIdByRequestQIndex\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_challenge\",\"type\":\"bytes32\"},{\"name\":\"_miner\",\"type\":\"address\"}],\"name\":\"didMine\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_requestId\",\"type\":\"uint256\"},{\"name\":\"_timestamp\",\"type\":\"uint256\"}],\"name\":\"getMinersByRequestIdAndTimestamp\",\"outputs\":[{\"name\":\"\",\"type\":\"address[5]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_user\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_staker\",\"type\":\"address\"}],\"name\":\"getStakerInfo\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_requestID\",\"type\":\"uint256\"},{\"name\":\"_index\",\"type\":\"uint256\"}],\"name\":\"getTimestampbyRequestIDandIndex\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_disputeId\",\"type\":\"uint256\"},{\"name\":\"_data\",\"type\":\"bytes32\"}],\"name\":\"getDisputeUintVars\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_requestId\",\"type\":\"uint256\"},{\"name\":\"_timestamp\",\"type\":\"uint256\"}],\"name\":\"retrieveData\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_user\",\"type\":\"address\"},{\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"allowedToTrade\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getCurrentVariables\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"},{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"string\"},{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_disputeId\",\"type\":\"uint256\"},{\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"didVote\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_berryContract\",\"type\":\"address\"}],\"name\":\"changeBerryContract\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_disputeId\",\"type\":\"uint256\"}],\"name\":\"getAllDisputeVars\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"},{\"name\":\"\",\"type\":\"bool\"},{\"name\":\"\",\"type\":\"bool\"},{\"name\":\"\",\"type\":\"bool\"},{\"name\":\"\",\"type\":\"address\"},{\"name\":\"\",\"type\":\"address\"},{\"name\":\"\",\"type\":\"address\"},{\"name\":\"\",\"type\":\"uint256[9]\"},{\"name\":\"\",\"type\":\"int256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getRequestQ\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256[51]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_requestId\",\"type\":\"uint256\"},{\"name\":\"_timestamp\",\"type\":\"uint256\"}],\"name\":\"getMinedBlockNum\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_hash\",\"type\":\"bytes32\"}],\"name\":\"getDisputeIdByDisputeHash\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_user\",\"type\":\"address\"},{\"name\":\"_spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_requestId\",\"type\":\"uint256\"},{\"name\":\"_data\",\"type\":\"bytes32\"}],\"name\":\"getRequestUintVars\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_requestId\",\"type\":\"uint256\"}],\"name\":\"getRequestVars\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"},{\"name\":\"\",\"type\":\"string\"},{\"name\":\"\",\"type\":\"bytes32\"},{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getLastNewValue\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_berryContract\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_newBerry\",\"type\":\"address\"}],\"name\":\"NewBerryAddress\",\"type\":\"event\"}]"

// BerryMasterBin is the compiled bytecode used for deploying new contracts.
const BerryMasterBin = `0x608060405234801561001057600080fd5b506040516120123803806120128339818101604052602081101561003357600080fd5b5051604080517f4601f1cd000000000000000000000000000000000000000000000000000000008152600060048201819052915173__$799602413129f49037f52758954cf5aa52$__92634601f1cd9260248082019391829003018186803b15801561009e57600080fd5b505af41580156100b2573d6000803e3d6000fd5b5050604080517f5f6f776e657200000000000000000000000000000000000000000000000000008152815190819003600690810182206000908152603f602081815285832080546001600160a01b0319908116339081179092557f5f646569747900000000000000000000000000000000000000000000000000008752875196879003909501862084528282528684208054861690911790557f74656c6c6f72436f6e74726163740000000000000000000000000000000000008552855194859003600e01852083529081529084902080546001600160a01b03891693168317905590825291517fc2d1449eb0b6547aa426e09d9942a77fa4fc8cd3296305b3163e22452e0bcb8d9450908190039091019150a150611e3c806101d66000396000f3fe6080604052600436106101f95760003560e01c806370a082311161010d578063ae0a8279116100a0578063da3799411161006f578063da37994114610a3e578063dd62ed3e14610a68578063e0ae93c114610aa3578063e1eee6d614610ad3578063fc7cf0a014610bf7576101f9565b8063ae0a8279146108ff578063af0b132714610932578063b5413029146109e3578063c775b54214610a0e576101f9565b806393fa4915116100dc57806393fa4915146107ad578063999cf26c146107dd578063a22e407a14610816578063a7c438bc146108c6576101f9565b806370a08231146106ce578063733bdef01461070157806377fbb6631461074d5780637f6fd5d91461077d576101f9565b80633180f8df116101905780634ee2cd7e1161015f5780634ee2cd7e146105d8578063612c8f7f146106115780636173c0b81461063b57806363bb82ad1461066557806369026d631461069e576101f9565b80633180f8df146104f25780633df0777b1461053557806346eee1c41461057957806347abd7f1146105a3576101f9565b806317d7de7c116101cc57806317d7de7c1461040357806318160ddd1461041857806319e8e03b1461042d5780631db842f0146104c8576101f9565b80630f0b424d1461028f57806311c98512146102cb578063133bee5e146103335780631507040114610379575b604080516d1d195b1b1bdc90dbdb9d1c9858dd60921b8152815190819003600e0181206000908152603f602090815283822054601f369081018390048302850183019095528484526001600160a01b03169360609392918190840183828082843760009201829052508451949550938493509150506020840185600019f43d604051816000823e82801561028b578282f35b8282fd5b34801561029b57600080fd5b506102b9600480360360208110156102b257600080fd5b5035610c0c565b60408051918252519081900360200190f35b3480156102d757600080fd5b506102fb600480360360408110156102ee57600080fd5b5080359060200135610c24565b604051808260a080838360005b83811015610320578181015183820152602001610308565b5050505090500191505060405180910390f35b34801561033f57600080fd5b5061035d6004803603602081101561035657600080fd5b5035610c45565b604080516001600160a01b039092168252519081900360200190f35b34801561038557600080fd5b5061038e610c57565b6040805160208082528351818301528351919283929083019185019080838360005b838110156103c85781810151838201526020016103b0565b50505050905090810190601f1680156103f55780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b34801561040f57600080fd5b5061038e610c68565b34801561042457600080fd5b506102b9610c74565b34801561043957600080fd5b50610442610c80565b6040518084815260200183815260200180602001828103825283818151815260200191508051906020019080838360005b8381101561048b578181015183820152602001610473565b50505050905090810190601f1680156104b85780820380516001836020036101000a031916815260200191505b5094505050505060405180910390f35b3480156104d457600080fd5b506102b9600480360360208110156104eb57600080fd5b5035610c9a565b3480156104fe57600080fd5b5061051c6004803603602081101561051557600080fd5b5035610cac565b6040805192835290151560208301528051918290030190f35b34801561054157600080fd5b506105656004803603604081101561055857600080fd5b5080359060200135610cc8565b604080519115158252519081900360200190f35b34801561058557600080fd5b506102b96004803603602081101561059c57600080fd5b5035610cdb565b3480156105af57600080fd5b506105d6600480360360208110156105c657600080fd5b50356001600160a01b0316610ced565b005b3480156105e457600080fd5b506102b9600480360360408110156105fb57600080fd5b506001600160a01b038135169060200135610d01565b34801561061d57600080fd5b506102b96004803603602081101561063457600080fd5b5035610d9e565b34801561064757600080fd5b506102b96004803603602081101561065e57600080fd5b5035610db0565b34801561067157600080fd5b506105656004803603604081101561068857600080fd5b50803590602001356001600160a01b0316610dc2565b3480156106aa57600080fd5b506102fb600480360360408110156106c157600080fd5b5080359060200135610dd5565b3480156106da57600080fd5b506102b9600480360360208110156106f157600080fd5b50356001600160a01b0316610def565b34801561070d57600080fd5b506107346004803603602081101561072457600080fd5b50356001600160a01b0316610e84565b6040805192835260208301919091528051918290030190f35b34801561075957600080fd5b506102b96004803603604081101561077057600080fd5b5080359060200135610e97565b34801561078957600080fd5b506102b9600480360360408110156107a057600080fd5b5080359060200135610eaa565b3480156107b957600080fd5b506102b9600480360360408110156107d057600080fd5b5080359060200135610ebd565b3480156107e957600080fd5b506105656004803603604081101561080057600080fd5b506001600160a01b038135169060200135610ed0565b34801561082257600080fd5b5061082b610f3a565b6040518087815260200186815260200185815260200180602001848152602001838152602001828103825285818151815260200191508051906020019080838360005b8381101561088657818101518382015260200161086e565b50505050905090810190601f1680156108b35780820380516001836020036101000a031916815260200191505b5097505050505050505060405180910390f35b3480156108d257600080fd5b50610565600480360360408110156108e957600080fd5b50803590602001356001600160a01b0316610f61565b34801561090b57600080fd5b506105d66004803603602081101561092257600080fd5b50356001600160a01b0316610f74565b34801561093e57600080fd5b5061095c6004803603602081101561095557600080fd5b5035610f85565b604080518a815289151560208201528815159181019190915286151560608201526001600160a01b03808716608083015285811660a0830152841660c082015260e081018361012080838360005b838110156109c25781810151838201526020016109aa565b50505050905001828152602001995050505050505050505060405180910390f35b3480156109ef57600080fd5b506109f8610fc9565b6040518151815280826106608083836020610308565b348015610a1a57600080fd5b506102b960048036036040811015610a3157600080fd5b5080359060200135610fdb565b348015610a4a57600080fd5b506102b960048036036020811015610a6157600080fd5b5035610fee565b348015610a7457600080fd5b506102b960048036036040811015610a8b57600080fd5b506001600160a01b0381358116916020013516611000565b348015610aaf57600080fd5b506102b960048036036040811015610ac657600080fd5b508035906020013561106b565b348015610adf57600080fd5b50610afd60048036036020811015610af657600080fd5b503561107e565b604051808060200180602001878152602001868152602001858152602001848152602001838103835289818151815260200191508051906020019080838360005b83811015610b56578181015183820152602001610b3e565b50505050905090810190601f168015610b835780820380516001836020036101000a031916815260200191505b5083810382528851815288516020918201918a019080838360005b83811015610bb6578181015183820152602001610b9e565b50505050905090810190601f168015610be35780820380516001836020036101000a031916815260200191505b509850505050505050505060405180910390f35b348015610c0357600080fd5b5061051c6110aa565b6000610c1e818363ffffffff6110bf16565b92915050565b610c2c611dab565b610c3e6000848463ffffffff6110d516565b9392505050565b6000610c1e818363ffffffff61112f16565b6060610c63600061114e565b905090565b6060610c63600061116b565b6000610c636000611195565b6000806060610c8f60006111c8565b925092509250909192565b6000610c1e818363ffffffff6112b416565b600080610cbf818463ffffffff6112ca16565b91509150915091565b6000610c3e81848463ffffffff61133016565b6000610c1e818363ffffffff61135716565b610cfe60008263ffffffff61137016565b50565b60408051633f48b1ff60e01b81526000600482018190526001600160a01b038516602483015260448201849052915173__$ff8037bb7c49d17e1e79a1c3f1e9fdeb7a$__91633f48b1ff916064808301926020929190829003018186803b158015610d6b57600080fd5b505af4158015610d7f573d6000803e3d6000fd5b505050506040513d6020811015610d9557600080fd5b50519392505050565b6000610c1e818363ffffffff6113f916565b6000610c1e818363ffffffff61140b16565b6000610c3e81848463ffffffff61143216565b610ddd611dab565b610c3e6000848463ffffffff61145f16565b604080516393b182b360e01b81526000600482018190526001600160a01b0384166024830152915173__$ff8037bb7c49d17e1e79a1c3f1e9fdeb7a$__916393b182b3916044808301926020929190829003018186803b158015610e5257600080fd5b505af4158015610e66573d6000803e3d6000fd5b505050506040513d6020811015610e7c57600080fd5b505192915050565b600080610cbf818463ffffffff6114c316565b6000610c3e81848463ffffffff6114ea16565b6000610c3e81848463ffffffff61151d16565b6000610c3e81848463ffffffff61154116565b604080516356555cf160e11b81526000600482018190526001600160a01b038516602483015260448201849052915173__$ff8037bb7c49d17e1e79a1c3f1e9fdeb7a$__9163acaab9e2916064808301926020929190829003018186803b158015610d6b57600080fd5b60008060006060600080610f4e6000611565565b949b939a50919850965094509092509050565b6000610c3e81848463ffffffff61172316565b610cfe60008263ffffffff61175516565b6000806000806000806000610f98611dc9565b6000610faa818b63ffffffff61181716565b9850985098509850985098509850985098509193959799909294969850565b610fd1611de8565b610c636000611a2c565b6000610c3e81848463ffffffff611a6c16565b6000610c1e818363ffffffff611a9016565b60408051632fcc801b60e21b81526000600482018190526001600160a01b03808616602484015284166044830152915173__$ff8037bb7c49d17e1e79a1c3f1e9fdeb7a$__9163bf32006c916064808301926020929190829003018186803b158015610d6b57600080fd5b6000610c3e81848463ffffffff611aa616565b6060806000808080611096818863ffffffff611aca16565b949c939b5091995097509550909350915050565b6000806110b76000611ca5565b915091509091565b6000908152604291909101602052604090205490565b6110dd611dab565b6000838152604885016020908152604080832085845260090190915290819020815160a08101928390529160059082845b81548152602001906001019080831161110e57505050505090509392505050565b6000908152603f9190910160205260409020546001600160a01b031690565b50604080518082019091526002815261151560f21b602082015290565b5060408051808201909152600f81526e54656c6c6f7220547269627574657360881b602082015290565b604080516b746f74616c5f737570706c7960a01b8152815190819003600c01902060009081528183016020522054919050565b600080606060006111d885611d1b565b600081815260488701602081815260408084208151670746f74616c5469760c41b8152825160089181900391909101812086526004820184528286205495879052938352805460026001821615610100026000190190911604601f81018490048402850184019092528184529495508594929183919083018282801561129f5780601f106112745761010080835404028352916020019161129f565b820191906000526020600020905b81548152906001019060200180831161128257829003601f168201915b50505050509050935093509350509193909250565b6000908152604991909101602052604090205490565b6000818152604883016020526040812060038101548291901561132057600381018054611314918791879190600019810190811061130457fe5b9060005260206000200154611541565b60019250925050611329565b50600091508190505b9250929050565b60009182526048929092016020908152604080832093835260079093019052205460ff1690565b6000908152604891909101602052604090206003015490565b60408051655f646569747960d01b815281519081900360060190206000908152603f840160205220546001600160a01b031633146113ad57600080fd5b60408051655f646569747960d01b815281519081900360060190206000908152603f90930160205290912080546001600160a01b039092166001600160a01b0319909216919091179055565b60009081526040918201602052205490565b6000603282111561141b57600080fd5b506000908152604391909101602052604090205490565b6000918252604192909201602090815260408083206001600160a01b039094168352929052205460ff1690565b611467611dab565b6000838152604885016020908152604080832085845260080190915290819020815160a08101928390529160059082845b81546001600160a01b0316815260019091019060200180831161149857505050505090509392505050565b6001600160a01b031660009081526047919091016020526040902080546001909101549091565b6000828152604884016020526040812060030180548390811061150957fe5b906000526020600020015490509392505050565b60009182526044929092016020908152604080832093835260059093019052205490565b60009182526048929092016020908152604080832093835260069093019052205490565b8054604080516f18dd5c9c995b9d14995c5d595cdd125960821b80825282519182900360109081018320600090815284870160208181528683205469646966666963756c747960b01b8752875196879003600a01872084528282528784205486885288519788900386018820855283835288852054855260488b01808452898620888a528a51998a900388018a2087528585528a87205487528185528a87206a6772616e756c617269747960a81b8b528b519a8b9003600b018b208852600490810186528b882054998b528b519a8b90039098018a2087529484528986205486528352888520670746f74616c5469760c41b8952895198899003600801892086529095018252878420548354601f600260001961010060018516150201909216919091049081018490048402890184019099528888529398899889986060988a98899894979596909594919391929185918301828280156117075780601f106116dc57610100808354040283529160200191611707565b820191906000526020600020905b8154815290600101906020018083116116ea57829003601f168201915b5050505050925095509550955095509550955091939550919395565b600082815260448401602090815260408083206001600160a01b038516845260060190915290205460ff169392505050565b60408051655f646569747960d01b815281519081900360060190206000908152603f840160205220546001600160a01b0316331461179257600080fd5b604080516d1d195b1b1bdc90dbdb9d1c9858dd60921b8152815190819003600e0181206000908152603f850160209081529083902080546001600160a01b0386166001600160a01b03199091168117909155825291517fc2d1449eb0b6547aa426e09d9942a77fa4fc8cd3296305b3163e22452e0bcb8d929181900390910190a15050565b600080600080600080600061182a611dc9565b5050506000868152604488016020908152604080832080546002820154600380840154600485015486516101208101808952681c995c5d595cdd125960ba1b905287518082036101290190208a526005808801808b52898c20548352895168074696d657374616d760bc1b81528a519081900360099081019091208d52818c528a8d2054848d01528a516476616c756560d81b81528b51908190039093019092208c52808b52898c2054838b015289516f6d696e457865637574696f6e4461746560801b81528a519081900360100190208c52808b52898c2054606084015289516c6e756d6265724f66566f74657360981b81528a5190819003600d0190208c52808b52898c2054608084015289516a313637b1b5a73ab6b132b960a91b81528a5190819003600b0190208c52808b52898c205460a08401528951681b5a5b995c94db1bdd60ba1b81528a51908190039092019091208b52808a52888b205460c083015288516571756f72756d60d01b815289519081900360060190208b52808a52888b205460e083015288516266656560e81b81528951908190039095019094208a5292909752949096205461010087810191909152600190930154919a5060ff8082169a509281048316985062010000810490921696506001600160a01b0363010000009092048216955091811693921691909295985092959850929598565b611a34611de8565b6040805161066081019182905290600184019060339082845b815481526020019060010190808311611a4d5750505050509050919050565b60009182526048929092016020908152604080832093835260059093019052205490565b6000908152604a91909101602052604090205490565b60009182526048929092016020908152604080832093835260049093019052205490565b6000818152604883016020908152604080832060028082015483516a6772616e756c617269747960a81b8152845190819003600b018120875260048401808752858820546f3932b8bab2b9ba28a837b9b4ba34b7b760811b83528651928390036010018320895281885286892054670746f74616c5469760c41b845287519384900360080184208a52918852868920548654601f6000196101006001848116159190910291909101909216979097049687018a90048a0285018a019098528584526060998a99909889988998899891978897938801969295929490918891830182828015611bf95780601f10611bce57610100808354040283529160200191611bf9565b820191906000526020600020905b815481529060010190602001808311611bdc57829003601f168201915b5050885460408051602060026001851615610100026000190190941693909304601f8101849004840282018401909252818152959b508a945092508401905082828015611c875780601f10611c5c57610100808354040283529160200191611c87565b820191906000526020600020905b815481529060010190602001808311611c6a57829003601f168201915b50505050509450965096509650965096509650509295509295509295565b604080517174696d654f664c6173744e657756616c756560701b8082528251601292819003830181206000908152858501602081815286832054835260428801815286832054948452865193849003909501909220815292529181205490918291611d11918591611541565b9360019350915050565b6040805161066081019182905260009182918291611d5c9190600187019060339082845b815481526020019060010190808311611d3f575050505050611d77565b60009081526043909501602052505060409092205492915050565b60008060005b6033811015611da5576020810284015180841015611d9c578093508192505b50600101611d7d565b50915091565b6040518060a001604052806005906020820280388339509192915050565b6040518061012001604052806009906020820280388339509192915050565b604051806106600160405280603390602082028038833950919291505056fea265627a7a723058208f825b37f40114c9f8689b2b8792c2aa6e313b233bb3f67a270a1197636999e964736f6c634300050a0032`

// DeployBerryMaster deploys a new Ethereum contract, binding an instance of BerryMaster to it.
func DeployBerryMaster(auth *bind.TransactOpts, backend bind.ContractBackend, _berryContract common.Address) (common.Address, *types.Transaction, *BerryMaster, error) {
	parsed, err := abi.JSON(strings.NewReader(BerryMasterABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(BerryMasterBin), backend, _berryContract)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &BerryMaster{BerryMasterCaller: BerryMasterCaller{contract: contract}, BerryMasterTransactor: BerryMasterTransactor{contract: contract}, BerryMasterFilterer: BerryMasterFilterer{contract: contract}}, nil
}

// BerryMaster is an auto generated Go binding around an Ethereum contract.
type BerryMaster struct {
	BerryMasterCaller     // Read-only binding to the contract
	BerryMasterTransactor // Write-only binding to the contract
	BerryMasterFilterer   // Log filterer for contract events
}

// BerryMasterCaller is an auto generated read-only Go binding around an Ethereum contract.
type BerryMasterCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BerryMasterTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BerryMasterTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BerryMasterFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BerryMasterFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BerryMasterSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BerryMasterSession struct {
	Contract     *BerryMaster     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BerryMasterCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BerryMasterCallerSession struct {
	Contract *BerryMasterCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// BerryMasterTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BerryMasterTransactorSession struct {
	Contract     *BerryMasterTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// BerryMasterRaw is an auto generated low-level Go binding around an Ethereum contract.
type BerryMasterRaw struct {
	Contract *BerryMaster // Generic contract binding to access the raw methods on
}

// BerryMasterCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BerryMasterCallerRaw struct {
	Contract *BerryMasterCaller // Generic read-only contract binding to access the raw methods on
}

// BerryMasterTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BerryMasterTransactorRaw struct {
	Contract *BerryMasterTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBerryMaster creates a new instance of BerryMaster, bound to a specific deployed contract.
func NewBerryMaster(address common.Address, backend bind.ContractBackend) (*BerryMaster, error) {
	contract, err := bindBerryMaster(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BerryMaster{BerryMasterCaller: BerryMasterCaller{contract: contract}, BerryMasterTransactor: BerryMasterTransactor{contract: contract}, BerryMasterFilterer: BerryMasterFilterer{contract: contract}}, nil
}

// NewBerryMasterCaller creates a new read-only instance of BerryMaster, bound to a specific deployed contract.
func NewBerryMasterCaller(address common.Address, caller bind.ContractCaller) (*BerryMasterCaller, error) {
	contract, err := bindBerryMaster(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BerryMasterCaller{contract: contract}, nil
}

// NewBerryMasterTransactor creates a new write-only instance of BerryMaster, bound to a specific deployed contract.
func NewBerryMasterTransactor(address common.Address, transactor bind.ContractTransactor) (*BerryMasterTransactor, error) {
	contract, err := bindBerryMaster(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BerryMasterTransactor{contract: contract}, nil
}

// NewBerryMasterFilterer creates a new log filterer instance of BerryMaster, bound to a specific deployed contract.
func NewBerryMasterFilterer(address common.Address, filterer bind.ContractFilterer) (*BerryMasterFilterer, error) {
	contract, err := bindBerryMaster(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BerryMasterFilterer{contract: contract}, nil
}

// bindBerryMaster binds a generic wrapper to an already deployed contract.
func bindBerryMaster(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BerryMasterABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BerryMaster *BerryMasterRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _BerryMaster.Contract.BerryMasterCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BerryMaster *BerryMasterRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BerryMaster.Contract.BerryMasterTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BerryMaster *BerryMasterRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BerryMaster.Contract.BerryMasterTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BerryMaster *BerryMasterCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _BerryMaster.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BerryMaster *BerryMasterTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BerryMaster.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BerryMaster *BerryMasterTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BerryMaster.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(_user address, _spender address) constant returns(uint256)
func (_BerryMaster *BerryMasterCaller) Allowance(opts *bind.CallOpts, _user common.Address, _spender common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BerryMaster.contract.Call(opts, out, "allowance", _user, _spender)
	return *ret0, err
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(_user address, _spender address) constant returns(uint256)
func (_BerryMaster *BerryMasterSession) Allowance(_user common.Address, _spender common.Address) (*big.Int, error) {
	return _BerryMaster.Contract.Allowance(&_BerryMaster.CallOpts, _user, _spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(_user address, _spender address) constant returns(uint256)
func (_BerryMaster *BerryMasterCallerSession) Allowance(_user common.Address, _spender common.Address) (*big.Int, error) {
	return _BerryMaster.Contract.Allowance(&_BerryMaster.CallOpts, _user, _spender)
}

// AllowedToTrade is a free data retrieval call binding the contract method 0x999cf26c.
//
// Solidity: function allowedToTrade(_user address, _amount uint256) constant returns(bool)
func (_BerryMaster *BerryMasterCaller) AllowedToTrade(opts *bind.CallOpts, _user common.Address, _amount *big.Int) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _BerryMaster.contract.Call(opts, out, "allowedToTrade", _user, _amount)
	return *ret0, err
}

// AllowedToTrade is a free data retrieval call binding the contract method 0x999cf26c.
//
// Solidity: function allowedToTrade(_user address, _amount uint256) constant returns(bool)
func (_BerryMaster *BerryMasterSession) AllowedToTrade(_user common.Address, _amount *big.Int) (bool, error) {
	return _BerryMaster.Contract.AllowedToTrade(&_BerryMaster.CallOpts, _user, _amount)
}

// AllowedToTrade is a free data retrieval call binding the contract method 0x999cf26c.
//
// Solidity: function allowedToTrade(_user address, _amount uint256) constant returns(bool)
func (_BerryMaster *BerryMasterCallerSession) AllowedToTrade(_user common.Address, _amount *big.Int) (bool, error) {
	return _BerryMaster.Contract.AllowedToTrade(&_BerryMaster.CallOpts, _user, _amount)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_user address) constant returns(uint256)
func (_BerryMaster *BerryMasterCaller) BalanceOf(opts *bind.CallOpts, _user common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BerryMaster.contract.Call(opts, out, "balanceOf", _user)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_user address) constant returns(uint256)
func (_BerryMaster *BerryMasterSession) BalanceOf(_user common.Address) (*big.Int, error) {
	return _BerryMaster.Contract.BalanceOf(&_BerryMaster.CallOpts, _user)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_user address) constant returns(uint256)
func (_BerryMaster *BerryMasterCallerSession) BalanceOf(_user common.Address) (*big.Int, error) {
	return _BerryMaster.Contract.BalanceOf(&_BerryMaster.CallOpts, _user)
}

// BalanceOfAt is a free data retrieval call binding the contract method 0x4ee2cd7e.
//
// Solidity: function balanceOfAt(_user address, _blockNumber uint256) constant returns(uint256)
func (_BerryMaster *BerryMasterCaller) BalanceOfAt(opts *bind.CallOpts, _user common.Address, _blockNumber *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BerryMaster.contract.Call(opts, out, "balanceOfAt", _user, _blockNumber)
	return *ret0, err
}

// BalanceOfAt is a free data retrieval call binding the contract method 0x4ee2cd7e.
//
// Solidity: function balanceOfAt(_user address, _blockNumber uint256) constant returns(uint256)
func (_BerryMaster *BerryMasterSession) BalanceOfAt(_user common.Address, _blockNumber *big.Int) (*big.Int, error) {
	return _BerryMaster.Contract.BalanceOfAt(&_BerryMaster.CallOpts, _user, _blockNumber)
}

// BalanceOfAt is a free data retrieval call binding the contract method 0x4ee2cd7e.
//
// Solidity: function balanceOfAt(_user address, _blockNumber uint256) constant returns(uint256)
func (_BerryMaster *BerryMasterCallerSession) BalanceOfAt(_user common.Address, _blockNumber *big.Int) (*big.Int, error) {
	return _BerryMaster.Contract.BalanceOfAt(&_BerryMaster.CallOpts, _user, _blockNumber)
}

// DidMine is a free data retrieval call binding the contract method 0x63bb82ad.
//
// Solidity: function didMine(_challenge bytes32, _miner address) constant returns(bool)
func (_BerryMaster *BerryMasterCaller) DidMine(opts *bind.CallOpts, _challenge [32]byte, _miner common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _BerryMaster.contract.Call(opts, out, "didMine", _challenge, _miner)
	return *ret0, err
}

// DidMine is a free data retrieval call binding the contract method 0x63bb82ad.
//
// Solidity: function didMine(_challenge bytes32, _miner address) constant returns(bool)
func (_BerryMaster *BerryMasterSession) DidMine(_challenge [32]byte, _miner common.Address) (bool, error) {
	return _BerryMaster.Contract.DidMine(&_BerryMaster.CallOpts, _challenge, _miner)
}

// DidMine is a free data retrieval call binding the contract method 0x63bb82ad.
//
// Solidity: function didMine(_challenge bytes32, _miner address) constant returns(bool)
func (_BerryMaster *BerryMasterCallerSession) DidMine(_challenge [32]byte, _miner common.Address) (bool, error) {
	return _BerryMaster.Contract.DidMine(&_BerryMaster.CallOpts, _challenge, _miner)
}

// DidVote is a free data retrieval call binding the contract method 0xa7c438bc.
//
// Solidity: function didVote(_disputeId uint256, _address address) constant returns(bool)
func (_BerryMaster *BerryMasterCaller) DidVote(opts *bind.CallOpts, _disputeId *big.Int, _address common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _BerryMaster.contract.Call(opts, out, "didVote", _disputeId, _address)
	return *ret0, err
}

// DidVote is a free data retrieval call binding the contract method 0xa7c438bc.
//
// Solidity: function didVote(_disputeId uint256, _address address) constant returns(bool)
func (_BerryMaster *BerryMasterSession) DidVote(_disputeId *big.Int, _address common.Address) (bool, error) {
	return _BerryMaster.Contract.DidVote(&_BerryMaster.CallOpts, _disputeId, _address)
}

// DidVote is a free data retrieval call binding the contract method 0xa7c438bc.
//
// Solidity: function didVote(_disputeId uint256, _address address) constant returns(bool)
func (_BerryMaster *BerryMasterCallerSession) DidVote(_disputeId *big.Int, _address common.Address) (bool, error) {
	return _BerryMaster.Contract.DidVote(&_BerryMaster.CallOpts, _disputeId, _address)
}

// GetAddressVars is a free data retrieval call binding the contract method 0x133bee5e.
//
// Solidity: function getAddressVars(_data bytes32) constant returns(address)
func (_BerryMaster *BerryMasterCaller) GetAddressVars(opts *bind.CallOpts, _data [32]byte) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _BerryMaster.contract.Call(opts, out, "getAddressVars", _data)
	return *ret0, err
}

// GetAddressVars is a free data retrieval call binding the contract method 0x133bee5e.
//
// Solidity: function getAddressVars(_data bytes32) constant returns(address)
func (_BerryMaster *BerryMasterSession) GetAddressVars(_data [32]byte) (common.Address, error) {
	return _BerryMaster.Contract.GetAddressVars(&_BerryMaster.CallOpts, _data)
}

// GetAddressVars is a free data retrieval call binding the contract method 0x133bee5e.
//
// Solidity: function getAddressVars(_data bytes32) constant returns(address)
func (_BerryMaster *BerryMasterCallerSession) GetAddressVars(_data [32]byte) (common.Address, error) {
	return _BerryMaster.Contract.GetAddressVars(&_BerryMaster.CallOpts, _data)
}

// GetAllDisputeVars is a free data retrieval call binding the contract method 0xaf0b1327.
//
// Solidity: function getAllDisputeVars(_disputeId uint256) constant returns(bytes32, bool, bool, bool, address, address, address, uint256[9], int256)
func (_BerryMaster *BerryMasterCaller) GetAllDisputeVars(opts *bind.CallOpts, _disputeId *big.Int) ([32]byte, bool, bool, bool, common.Address, common.Address, common.Address, [9]*big.Int, *big.Int, error) {
	var (
		ret0 = new([32]byte)
		ret1 = new(bool)
		ret2 = new(bool)
		ret3 = new(bool)
		ret4 = new(common.Address)
		ret5 = new(common.Address)
		ret6 = new(common.Address)
		ret7 = new([9]*big.Int)
		ret8 = new(*big.Int)
	)
	out := &[]interface{}{
		ret0,
		ret1,
		ret2,
		ret3,
		ret4,
		ret5,
		ret6,
		ret7,
		ret8,
	}
	err := _BerryMaster.contract.Call(opts, out, "getAllDisputeVars", _disputeId)
	return *ret0, *ret1, *ret2, *ret3, *ret4, *ret5, *ret6, *ret7, *ret8, err
}

// GetAllDisputeVars is a free data retrieval call binding the contract method 0xaf0b1327.
//
// Solidity: function getAllDisputeVars(_disputeId uint256) constant returns(bytes32, bool, bool, bool, address, address, address, uint256[9], int256)
func (_BerryMaster *BerryMasterSession) GetAllDisputeVars(_disputeId *big.Int) ([32]byte, bool, bool, bool, common.Address, common.Address, common.Address, [9]*big.Int, *big.Int, error) {
	return _BerryMaster.Contract.GetAllDisputeVars(&_BerryMaster.CallOpts, _disputeId)
}

// GetAllDisputeVars is a free data retrieval call binding the contract method 0xaf0b1327.
//
// Solidity: function getAllDisputeVars(_disputeId uint256) constant returns(bytes32, bool, bool, bool, address, address, address, uint256[9], int256)
func (_BerryMaster *BerryMasterCallerSession) GetAllDisputeVars(_disputeId *big.Int) ([32]byte, bool, bool, bool, common.Address, common.Address, common.Address, [9]*big.Int, *big.Int, error) {
	return _BerryMaster.Contract.GetAllDisputeVars(&_BerryMaster.CallOpts, _disputeId)
}

// GetCurrentVariables is a free data retrieval call binding the contract method 0xa22e407a.
//
// Solidity: function getCurrentVariables() constant returns(bytes32, uint256, uint256, string, uint256, uint256)
func (_BerryMaster *BerryMasterCaller) GetCurrentVariables(opts *bind.CallOpts) ([32]byte, *big.Int, *big.Int, string, *big.Int, *big.Int, error) {
	var (
		ret0 = new([32]byte)
		ret1 = new(*big.Int)
		ret2 = new(*big.Int)
		ret3 = new(string)
		ret4 = new(*big.Int)
		ret5 = new(*big.Int)
	)
	out := &[]interface{}{
		ret0,
		ret1,
		ret2,
		ret3,
		ret4,
		ret5,
	}
	err := _BerryMaster.contract.Call(opts, out, "getCurrentVariables")
	return *ret0, *ret1, *ret2, *ret3, *ret4, *ret5, err
}

// GetCurrentVariables is a free data retrieval call binding the contract method 0xa22e407a.
//
// Solidity: function getCurrentVariables() constant returns(bytes32, uint256, uint256, string, uint256, uint256)
func (_BerryMaster *BerryMasterSession) GetCurrentVariables() ([32]byte, *big.Int, *big.Int, string, *big.Int, *big.Int, error) {
	return _BerryMaster.Contract.GetCurrentVariables(&_BerryMaster.CallOpts)
}

// GetCurrentVariables is a free data retrieval call binding the contract method 0xa22e407a.
//
// Solidity: function getCurrentVariables() constant returns(bytes32, uint256, uint256, string, uint256, uint256)
func (_BerryMaster *BerryMasterCallerSession) GetCurrentVariables() ([32]byte, *big.Int, *big.Int, string, *big.Int, *big.Int, error) {
	return _BerryMaster.Contract.GetCurrentVariables(&_BerryMaster.CallOpts)
}

// GetDisputeIdByDisputeHash is a free data retrieval call binding the contract method 0xda379941.
//
// Solidity: function getDisputeIdByDisputeHash(_hash bytes32) constant returns(uint256)
func (_BerryMaster *BerryMasterCaller) GetDisputeIdByDisputeHash(opts *bind.CallOpts, _hash [32]byte) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BerryMaster.contract.Call(opts, out, "getDisputeIdByDisputeHash", _hash)
	return *ret0, err
}

// GetDisputeIdByDisputeHash is a free data retrieval call binding the contract method 0xda379941.
//
// Solidity: function getDisputeIdByDisputeHash(_hash bytes32) constant returns(uint256)
func (_BerryMaster *BerryMasterSession) GetDisputeIdByDisputeHash(_hash [32]byte) (*big.Int, error) {
	return _BerryMaster.Contract.GetDisputeIdByDisputeHash(&_BerryMaster.CallOpts, _hash)
}

// GetDisputeIdByDisputeHash is a free data retrieval call binding the contract method 0xda379941.
//
// Solidity: function getDisputeIdByDisputeHash(_hash bytes32) constant returns(uint256)
func (_BerryMaster *BerryMasterCallerSession) GetDisputeIdByDisputeHash(_hash [32]byte) (*big.Int, error) {
	return _BerryMaster.Contract.GetDisputeIdByDisputeHash(&_BerryMaster.CallOpts, _hash)
}

// GetDisputeUintVars is a free data retrieval call binding the contract method 0x7f6fd5d9.
//
// Solidity: function getDisputeUintVars(_disputeId uint256, _data bytes32) constant returns(uint256)
func (_BerryMaster *BerryMasterCaller) GetDisputeUintVars(opts *bind.CallOpts, _disputeId *big.Int, _data [32]byte) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BerryMaster.contract.Call(opts, out, "getDisputeUintVars", _disputeId, _data)
	return *ret0, err
}

// GetDisputeUintVars is a free data retrieval call binding the contract method 0x7f6fd5d9.
//
// Solidity: function getDisputeUintVars(_disputeId uint256, _data bytes32) constant returns(uint256)
func (_BerryMaster *BerryMasterSession) GetDisputeUintVars(_disputeId *big.Int, _data [32]byte) (*big.Int, error) {
	return _BerryMaster.Contract.GetDisputeUintVars(&_BerryMaster.CallOpts, _disputeId, _data)
}

// GetDisputeUintVars is a free data retrieval call binding the contract method 0x7f6fd5d9.
//
// Solidity: function getDisputeUintVars(_disputeId uint256, _data bytes32) constant returns(uint256)
func (_BerryMaster *BerryMasterCallerSession) GetDisputeUintVars(_disputeId *big.Int, _data [32]byte) (*big.Int, error) {
	return _BerryMaster.Contract.GetDisputeUintVars(&_BerryMaster.CallOpts, _disputeId, _data)
}

// GetLastNewValue is a free data retrieval call binding the contract method 0xfc7cf0a0.
//
// Solidity: function getLastNewValue() constant returns(uint256, bool)
func (_BerryMaster *BerryMasterCaller) GetLastNewValue(opts *bind.CallOpts) (*big.Int, bool, error) {
	var (
		ret0 = new(*big.Int)
		ret1 = new(bool)
	)
	out := &[]interface{}{
		ret0,
		ret1,
	}
	err := _BerryMaster.contract.Call(opts, out, "getLastNewValue")
	return *ret0, *ret1, err
}

// GetLastNewValue is a free data retrieval call binding the contract method 0xfc7cf0a0.
//
// Solidity: function getLastNewValue() constant returns(uint256, bool)
func (_BerryMaster *BerryMasterSession) GetLastNewValue() (*big.Int, bool, error) {
	return _BerryMaster.Contract.GetLastNewValue(&_BerryMaster.CallOpts)
}

// GetLastNewValue is a free data retrieval call binding the contract method 0xfc7cf0a0.
//
// Solidity: function getLastNewValue() constant returns(uint256, bool)
func (_BerryMaster *BerryMasterCallerSession) GetLastNewValue() (*big.Int, bool, error) {
	return _BerryMaster.Contract.GetLastNewValue(&_BerryMaster.CallOpts)
}

// GetLastNewValueById is a free data retrieval call binding the contract method 0x3180f8df.
//
// Solidity: function getLastNewValueById(_requestId uint256) constant returns(uint256, bool)
func (_BerryMaster *BerryMasterCaller) GetLastNewValueById(opts *bind.CallOpts, _requestId *big.Int) (*big.Int, bool, error) {
	var (
		ret0 = new(*big.Int)
		ret1 = new(bool)
	)
	out := &[]interface{}{
		ret0,
		ret1,
	}
	err := _BerryMaster.contract.Call(opts, out, "getLastNewValueById", _requestId)
	return *ret0, *ret1, err
}

// GetLastNewValueById is a free data retrieval call binding the contract method 0x3180f8df.
//
// Solidity: function getLastNewValueById(_requestId uint256) constant returns(uint256, bool)
func (_BerryMaster *BerryMasterSession) GetLastNewValueById(_requestId *big.Int) (*big.Int, bool, error) {
	return _BerryMaster.Contract.GetLastNewValueById(&_BerryMaster.CallOpts, _requestId)
}

// GetLastNewValueById is a free data retrieval call binding the contract method 0x3180f8df.
//
// Solidity: function getLastNewValueById(_requestId uint256) constant returns(uint256, bool)
func (_BerryMaster *BerryMasterCallerSession) GetLastNewValueById(_requestId *big.Int) (*big.Int, bool, error) {
	return _BerryMaster.Contract.GetLastNewValueById(&_BerryMaster.CallOpts, _requestId)
}

// GetMinedBlockNum is a free data retrieval call binding the contract method 0xc775b542.
//
// Solidity: function getMinedBlockNum(_requestId uint256, _timestamp uint256) constant returns(uint256)
func (_BerryMaster *BerryMasterCaller) GetMinedBlockNum(opts *bind.CallOpts, _requestId *big.Int, _timestamp *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BerryMaster.contract.Call(opts, out, "getMinedBlockNum", _requestId, _timestamp)
	return *ret0, err
}

// GetMinedBlockNum is a free data retrieval call binding the contract method 0xc775b542.
//
// Solidity: function getMinedBlockNum(_requestId uint256, _timestamp uint256) constant returns(uint256)
func (_BerryMaster *BerryMasterSession) GetMinedBlockNum(_requestId *big.Int, _timestamp *big.Int) (*big.Int, error) {
	return _BerryMaster.Contract.GetMinedBlockNum(&_BerryMaster.CallOpts, _requestId, _timestamp)
}

// GetMinedBlockNum is a free data retrieval call binding the contract method 0xc775b542.
//
// Solidity: function getMinedBlockNum(_requestId uint256, _timestamp uint256) constant returns(uint256)
func (_BerryMaster *BerryMasterCallerSession) GetMinedBlockNum(_requestId *big.Int, _timestamp *big.Int) (*big.Int, error) {
	return _BerryMaster.Contract.GetMinedBlockNum(&_BerryMaster.CallOpts, _requestId, _timestamp)
}

// GetMinersByRequestIdAndTimestamp is a free data retrieval call binding the contract method 0x69026d63.
//
// Solidity: function getMinersByRequestIdAndTimestamp(_requestId uint256, _timestamp uint256) constant returns(address[5])
func (_BerryMaster *BerryMasterCaller) GetMinersByRequestIdAndTimestamp(opts *bind.CallOpts, _requestId *big.Int, _timestamp *big.Int) ([5]common.Address, error) {
	var (
		ret0 = new([5]common.Address)
	)
	out := ret0
	err := _BerryMaster.contract.Call(opts, out, "getMinersByRequestIdAndTimestamp", _requestId, _timestamp)
	return *ret0, err
}

// GetMinersByRequestIdAndTimestamp is a free data retrieval call binding the contract method 0x69026d63.
//
// Solidity: function getMinersByRequestIdAndTimestamp(_requestId uint256, _timestamp uint256) constant returns(address[5])
func (_BerryMaster *BerryMasterSession) GetMinersByRequestIdAndTimestamp(_requestId *big.Int, _timestamp *big.Int) ([5]common.Address, error) {
	return _BerryMaster.Contract.GetMinersByRequestIdAndTimestamp(&_BerryMaster.CallOpts, _requestId, _timestamp)
}

// GetMinersByRequestIdAndTimestamp is a free data retrieval call binding the contract method 0x69026d63.
//
// Solidity: function getMinersByRequestIdAndTimestamp(_requestId uint256, _timestamp uint256) constant returns(address[5])
func (_BerryMaster *BerryMasterCallerSession) GetMinersByRequestIdAndTimestamp(_requestId *big.Int, _timestamp *big.Int) ([5]common.Address, error) {
	return _BerryMaster.Contract.GetMinersByRequestIdAndTimestamp(&_BerryMaster.CallOpts, _requestId, _timestamp)
}

// GetName is a free data retrieval call binding the contract method 0x17d7de7c.
//
// Solidity: function getName() constant returns(string)
func (_BerryMaster *BerryMasterCaller) GetName(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _BerryMaster.contract.Call(opts, out, "getName")
	return *ret0, err
}

// GetName is a free data retrieval call binding the contract method 0x17d7de7c.
//
// Solidity: function getName() constant returns(string)
func (_BerryMaster *BerryMasterSession) GetName() (string, error) {
	return _BerryMaster.Contract.GetName(&_BerryMaster.CallOpts)
}

// GetName is a free data retrieval call binding the contract method 0x17d7de7c.
//
// Solidity: function getName() constant returns(string)
func (_BerryMaster *BerryMasterCallerSession) GetName() (string, error) {
	return _BerryMaster.Contract.GetName(&_BerryMaster.CallOpts)
}

// GetNewValueCountbyRequestId is a free data retrieval call binding the contract method 0x46eee1c4.
//
// Solidity: function getNewValueCountbyRequestId(_requestId uint256) constant returns(uint256)
func (_BerryMaster *BerryMasterCaller) GetNewValueCountbyRequestId(opts *bind.CallOpts, _requestId *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BerryMaster.contract.Call(opts, out, "getNewValueCountbyRequestId", _requestId)
	return *ret0, err
}

// GetNewValueCountbyRequestId is a free data retrieval call binding the contract method 0x46eee1c4.
//
// Solidity: function getNewValueCountbyRequestId(_requestId uint256) constant returns(uint256)
func (_BerryMaster *BerryMasterSession) GetNewValueCountbyRequestId(_requestId *big.Int) (*big.Int, error) {
	return _BerryMaster.Contract.GetNewValueCountbyRequestId(&_BerryMaster.CallOpts, _requestId)
}

// GetNewValueCountbyRequestId is a free data retrieval call binding the contract method 0x46eee1c4.
//
// Solidity: function getNewValueCountbyRequestId(_requestId uint256) constant returns(uint256)
func (_BerryMaster *BerryMasterCallerSession) GetNewValueCountbyRequestId(_requestId *big.Int) (*big.Int, error) {
	return _BerryMaster.Contract.GetNewValueCountbyRequestId(&_BerryMaster.CallOpts, _requestId)
}

// GetRequestIdByQueryHash is a free data retrieval call binding the contract method 0x1db842f0.
//
// Solidity: function getRequestIdByQueryHash(_request bytes32) constant returns(uint256)
func (_BerryMaster *BerryMasterCaller) GetRequestIdByQueryHash(opts *bind.CallOpts, _request [32]byte) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BerryMaster.contract.Call(opts, out, "getRequestIdByQueryHash", _request)
	return *ret0, err
}

// GetRequestIdByQueryHash is a free data retrieval call binding the contract method 0x1db842f0.
//
// Solidity: function getRequestIdByQueryHash(_request bytes32) constant returns(uint256)
func (_BerryMaster *BerryMasterSession) GetRequestIdByQueryHash(_request [32]byte) (*big.Int, error) {
	return _BerryMaster.Contract.GetRequestIdByQueryHash(&_BerryMaster.CallOpts, _request)
}

// GetRequestIdByQueryHash is a free data retrieval call binding the contract method 0x1db842f0.
//
// Solidity: function getRequestIdByQueryHash(_request bytes32) constant returns(uint256)
func (_BerryMaster *BerryMasterCallerSession) GetRequestIdByQueryHash(_request [32]byte) (*big.Int, error) {
	return _BerryMaster.Contract.GetRequestIdByQueryHash(&_BerryMaster.CallOpts, _request)
}

// GetRequestIdByRequestQIndex is a free data retrieval call binding the contract method 0x6173c0b8.
//
// Solidity: function getRequestIdByRequestQIndex(_index uint256) constant returns(uint256)
func (_BerryMaster *BerryMasterCaller) GetRequestIdByRequestQIndex(opts *bind.CallOpts, _index *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BerryMaster.contract.Call(opts, out, "getRequestIdByRequestQIndex", _index)
	return *ret0, err
}

// GetRequestIdByRequestQIndex is a free data retrieval call binding the contract method 0x6173c0b8.
//
// Solidity: function getRequestIdByRequestQIndex(_index uint256) constant returns(uint256)
func (_BerryMaster *BerryMasterSession) GetRequestIdByRequestQIndex(_index *big.Int) (*big.Int, error) {
	return _BerryMaster.Contract.GetRequestIdByRequestQIndex(&_BerryMaster.CallOpts, _index)
}

// GetRequestIdByRequestQIndex is a free data retrieval call binding the contract method 0x6173c0b8.
//
// Solidity: function getRequestIdByRequestQIndex(_index uint256) constant returns(uint256)
func (_BerryMaster *BerryMasterCallerSession) GetRequestIdByRequestQIndex(_index *big.Int) (*big.Int, error) {
	return _BerryMaster.Contract.GetRequestIdByRequestQIndex(&_BerryMaster.CallOpts, _index)
}

// GetRequestIdByTimestamp is a free data retrieval call binding the contract method 0x0f0b424d.
//
// Solidity: function getRequestIdByTimestamp(_timestamp uint256) constant returns(uint256)
func (_BerryMaster *BerryMasterCaller) GetRequestIdByTimestamp(opts *bind.CallOpts, _timestamp *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BerryMaster.contract.Call(opts, out, "getRequestIdByTimestamp", _timestamp)
	return *ret0, err
}

// GetRequestIdByTimestamp is a free data retrieval call binding the contract method 0x0f0b424d.
//
// Solidity: function getRequestIdByTimestamp(_timestamp uint256) constant returns(uint256)
func (_BerryMaster *BerryMasterSession) GetRequestIdByTimestamp(_timestamp *big.Int) (*big.Int, error) {
	return _BerryMaster.Contract.GetRequestIdByTimestamp(&_BerryMaster.CallOpts, _timestamp)
}

// GetRequestIdByTimestamp is a free data retrieval call binding the contract method 0x0f0b424d.
//
// Solidity: function getRequestIdByTimestamp(_timestamp uint256) constant returns(uint256)
func (_BerryMaster *BerryMasterCallerSession) GetRequestIdByTimestamp(_timestamp *big.Int) (*big.Int, error) {
	return _BerryMaster.Contract.GetRequestIdByTimestamp(&_BerryMaster.CallOpts, _timestamp)
}

// GetRequestQ is a free data retrieval call binding the contract method 0xb5413029.
//
// Solidity: function getRequestQ() constant returns(uint256[51])
func (_BerryMaster *BerryMasterCaller) GetRequestQ(opts *bind.CallOpts) ([51]*big.Int, error) {
	var (
		ret0 = new([51]*big.Int)
	)
	out := ret0
	err := _BerryMaster.contract.Call(opts, out, "getRequestQ")
	return *ret0, err
}

// GetRequestQ is a free data retrieval call binding the contract method 0xb5413029.
//
// Solidity: function getRequestQ() constant returns(uint256[51])
func (_BerryMaster *BerryMasterSession) GetRequestQ() ([51]*big.Int, error) {
	return _BerryMaster.Contract.GetRequestQ(&_BerryMaster.CallOpts)
}

// GetRequestQ is a free data retrieval call binding the contract method 0xb5413029.
//
// Solidity: function getRequestQ() constant returns(uint256[51])
func (_BerryMaster *BerryMasterCallerSession) GetRequestQ() ([51]*big.Int, error) {
	return _BerryMaster.Contract.GetRequestQ(&_BerryMaster.CallOpts)
}

// GetRequestUintVars is a free data retrieval call binding the contract method 0xe0ae93c1.
//
// Solidity: function getRequestUintVars(_requestId uint256, _data bytes32) constant returns(uint256)
func (_BerryMaster *BerryMasterCaller) GetRequestUintVars(opts *bind.CallOpts, _requestId *big.Int, _data [32]byte) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BerryMaster.contract.Call(opts, out, "getRequestUintVars", _requestId, _data)
	return *ret0, err
}

// GetRequestUintVars is a free data retrieval call binding the contract method 0xe0ae93c1.
//
// Solidity: function getRequestUintVars(_requestId uint256, _data bytes32) constant returns(uint256)
func (_BerryMaster *BerryMasterSession) GetRequestUintVars(_requestId *big.Int, _data [32]byte) (*big.Int, error) {
	return _BerryMaster.Contract.GetRequestUintVars(&_BerryMaster.CallOpts, _requestId, _data)
}

// GetRequestUintVars is a free data retrieval call binding the contract method 0xe0ae93c1.
//
// Solidity: function getRequestUintVars(_requestId uint256, _data bytes32) constant returns(uint256)
func (_BerryMaster *BerryMasterCallerSession) GetRequestUintVars(_requestId *big.Int, _data [32]byte) (*big.Int, error) {
	return _BerryMaster.Contract.GetRequestUintVars(&_BerryMaster.CallOpts, _requestId, _data)
}

// GetRequestVars is a free data retrieval call binding the contract method 0xe1eee6d6.
//
// Solidity: function getRequestVars(_requestId uint256) constant returns(string, string, bytes32, uint256, uint256, uint256)
func (_BerryMaster *BerryMasterCaller) GetRequestVars(opts *bind.CallOpts, _requestId *big.Int) (string, string, [32]byte, *big.Int, *big.Int, *big.Int, error) {
	var (
		ret0 = new(string)
		ret1 = new(string)
		ret2 = new([32]byte)
		ret3 = new(*big.Int)
		ret4 = new(*big.Int)
		ret5 = new(*big.Int)
	)
	out := &[]interface{}{
		ret0,
		ret1,
		ret2,
		ret3,
		ret4,
		ret5,
	}
	err := _BerryMaster.contract.Call(opts, out, "getRequestVars", _requestId)
	return *ret0, *ret1, *ret2, *ret3, *ret4, *ret5, err
}

// GetRequestVars is a free data retrieval call binding the contract method 0xe1eee6d6.
//
// Solidity: function getRequestVars(_requestId uint256) constant returns(string, string, bytes32, uint256, uint256, uint256)
func (_BerryMaster *BerryMasterSession) GetRequestVars(_requestId *big.Int) (string, string, [32]byte, *big.Int, *big.Int, *big.Int, error) {
	return _BerryMaster.Contract.GetRequestVars(&_BerryMaster.CallOpts, _requestId)
}

// GetRequestVars is a free data retrieval call binding the contract method 0xe1eee6d6.
//
// Solidity: function getRequestVars(_requestId uint256) constant returns(string, string, bytes32, uint256, uint256, uint256)
func (_BerryMaster *BerryMasterCallerSession) GetRequestVars(_requestId *big.Int) (string, string, [32]byte, *big.Int, *big.Int, *big.Int, error) {
	return _BerryMaster.Contract.GetRequestVars(&_BerryMaster.CallOpts, _requestId)
}

// GetStakerInfo is a free data retrieval call binding the contract method 0x733bdef0.
//
// Solidity: function getStakerInfo(_staker address) constant returns(uint256, uint256)
func (_BerryMaster *BerryMasterCaller) GetStakerInfo(opts *bind.CallOpts, _staker common.Address) (*big.Int, *big.Int, error) {
	var (
		ret0 = new(*big.Int)
		ret1 = new(*big.Int)
	)
	out := &[]interface{}{
		ret0,
		ret1,
	}
	err := _BerryMaster.contract.Call(opts, out, "getStakerInfo", _staker)
	return *ret0, *ret1, err
}

// GetStakerInfo is a free data retrieval call binding the contract method 0x733bdef0.
//
// Solidity: function getStakerInfo(_staker address) constant returns(uint256, uint256)
func (_BerryMaster *BerryMasterSession) GetStakerInfo(_staker common.Address) (*big.Int, *big.Int, error) {
	return _BerryMaster.Contract.GetStakerInfo(&_BerryMaster.CallOpts, _staker)
}

// GetStakerInfo is a free data retrieval call binding the contract method 0x733bdef0.
//
// Solidity: function getStakerInfo(_staker address) constant returns(uint256, uint256)
func (_BerryMaster *BerryMasterCallerSession) GetStakerInfo(_staker common.Address) (*big.Int, *big.Int, error) {
	return _BerryMaster.Contract.GetStakerInfo(&_BerryMaster.CallOpts, _staker)
}

// GetSubmissionsByTimestamp is a free data retrieval call binding the contract method 0x11c98512.
//
// Solidity: function getSubmissionsByTimestamp(_requestId uint256, _timestamp uint256) constant returns(uint256[5])
func (_BerryMaster *BerryMasterCaller) GetSubmissionsByTimestamp(opts *bind.CallOpts, _requestId *big.Int, _timestamp *big.Int) ([5]*big.Int, error) {
	var (
		ret0 = new([5]*big.Int)
	)
	out := ret0
	err := _BerryMaster.contract.Call(opts, out, "getSubmissionsByTimestamp", _requestId, _timestamp)
	return *ret0, err
}

// GetSubmissionsByTimestamp is a free data retrieval call binding the contract method 0x11c98512.
//
// Solidity: function getSubmissionsByTimestamp(_requestId uint256, _timestamp uint256) constant returns(uint256[5])
func (_BerryMaster *BerryMasterSession) GetSubmissionsByTimestamp(_requestId *big.Int, _timestamp *big.Int) ([5]*big.Int, error) {
	return _BerryMaster.Contract.GetSubmissionsByTimestamp(&_BerryMaster.CallOpts, _requestId, _timestamp)
}

// GetSubmissionsByTimestamp is a free data retrieval call binding the contract method 0x11c98512.
//
// Solidity: function getSubmissionsByTimestamp(_requestId uint256, _timestamp uint256) constant returns(uint256[5])
func (_BerryMaster *BerryMasterCallerSession) GetSubmissionsByTimestamp(_requestId *big.Int, _timestamp *big.Int) ([5]*big.Int, error) {
	return _BerryMaster.Contract.GetSubmissionsByTimestamp(&_BerryMaster.CallOpts, _requestId, _timestamp)
}

// GetSymbol is a free data retrieval call binding the contract method 0x15070401.
//
// Solidity: function getSymbol() constant returns(string)
func (_BerryMaster *BerryMasterCaller) GetSymbol(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _BerryMaster.contract.Call(opts, out, "getSymbol")
	return *ret0, err
}

// GetSymbol is a free data retrieval call binding the contract method 0x15070401.
//
// Solidity: function getSymbol() constant returns(string)
func (_BerryMaster *BerryMasterSession) GetSymbol() (string, error) {
	return _BerryMaster.Contract.GetSymbol(&_BerryMaster.CallOpts)
}

// GetSymbol is a free data retrieval call binding the contract method 0x15070401.
//
// Solidity: function getSymbol() constant returns(string)
func (_BerryMaster *BerryMasterCallerSession) GetSymbol() (string, error) {
	return _BerryMaster.Contract.GetSymbol(&_BerryMaster.CallOpts)
}

// GetTimestampbyRequestIDandIndex is a free data retrieval call binding the contract method 0x77fbb663.
//
// Solidity: function getTimestampbyRequestIDandIndex(_requestID uint256, _index uint256) constant returns(uint256)
func (_BerryMaster *BerryMasterCaller) GetTimestampbyRequestIDandIndex(opts *bind.CallOpts, _requestID *big.Int, _index *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BerryMaster.contract.Call(opts, out, "getTimestampbyRequestIDandIndex", _requestID, _index)
	return *ret0, err
}

// GetTimestampbyRequestIDandIndex is a free data retrieval call binding the contract method 0x77fbb663.
//
// Solidity: function getTimestampbyRequestIDandIndex(_requestID uint256, _index uint256) constant returns(uint256)
func (_BerryMaster *BerryMasterSession) GetTimestampbyRequestIDandIndex(_requestID *big.Int, _index *big.Int) (*big.Int, error) {
	return _BerryMaster.Contract.GetTimestampbyRequestIDandIndex(&_BerryMaster.CallOpts, _requestID, _index)
}

// GetTimestampbyRequestIDandIndex is a free data retrieval call binding the contract method 0x77fbb663.
//
// Solidity: function getTimestampbyRequestIDandIndex(_requestID uint256, _index uint256) constant returns(uint256)
func (_BerryMaster *BerryMasterCallerSession) GetTimestampbyRequestIDandIndex(_requestID *big.Int, _index *big.Int) (*big.Int, error) {
	return _BerryMaster.Contract.GetTimestampbyRequestIDandIndex(&_BerryMaster.CallOpts, _requestID, _index)
}

// GetUintVar is a free data retrieval call binding the contract method 0x612c8f7f.
//
// Solidity: function getUintVar(_data bytes32) constant returns(uint256)
func (_BerryMaster *BerryMasterCaller) GetUintVar(opts *bind.CallOpts, _data [32]byte) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BerryMaster.contract.Call(opts, out, "getUintVar", _data)
	return *ret0, err
}

// GetUintVar is a free data retrieval call binding the contract method 0x612c8f7f.
//
// Solidity: function getUintVar(_data bytes32) constant returns(uint256)
func (_BerryMaster *BerryMasterSession) GetUintVar(_data [32]byte) (*big.Int, error) {
	return _BerryMaster.Contract.GetUintVar(&_BerryMaster.CallOpts, _data)
}

// GetUintVar is a free data retrieval call binding the contract method 0x612c8f7f.
//
// Solidity: function getUintVar(_data bytes32) constant returns(uint256)
func (_BerryMaster *BerryMasterCallerSession) GetUintVar(_data [32]byte) (*big.Int, error) {
	return _BerryMaster.Contract.GetUintVar(&_BerryMaster.CallOpts, _data)
}

// GetVariablesOnDeck is a free data retrieval call binding the contract method 0x19e8e03b.
//
// Solidity: function getVariablesOnDeck() constant returns(uint256, uint256, string)
func (_BerryMaster *BerryMasterCaller) GetVariablesOnDeck(opts *bind.CallOpts) (*big.Int, *big.Int, string, error) {
	var (
		ret0 = new(*big.Int)
		ret1 = new(*big.Int)
		ret2 = new(string)
	)
	out := &[]interface{}{
		ret0,
		ret1,
		ret2,
	}
	err := _BerryMaster.contract.Call(opts, out, "getVariablesOnDeck")
	return *ret0, *ret1, *ret2, err
}

// GetVariablesOnDeck is a free data retrieval call binding the contract method 0x19e8e03b.
//
// Solidity: function getVariablesOnDeck() constant returns(uint256, uint256, string)
func (_BerryMaster *BerryMasterSession) GetVariablesOnDeck() (*big.Int, *big.Int, string, error) {
	return _BerryMaster.Contract.GetVariablesOnDeck(&_BerryMaster.CallOpts)
}

// GetVariablesOnDeck is a free data retrieval call binding the contract method 0x19e8e03b.
//
// Solidity: function getVariablesOnDeck() constant returns(uint256, uint256, string)
func (_BerryMaster *BerryMasterCallerSession) GetVariablesOnDeck() (*big.Int, *big.Int, string, error) {
	return _BerryMaster.Contract.GetVariablesOnDeck(&_BerryMaster.CallOpts)
}

// IsInDispute is a free data retrieval call binding the contract method 0x3df0777b.
//
// Solidity: function isInDispute(_requestId uint256, _timestamp uint256) constant returns(bool)
func (_BerryMaster *BerryMasterCaller) IsInDispute(opts *bind.CallOpts, _requestId *big.Int, _timestamp *big.Int) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _BerryMaster.contract.Call(opts, out, "isInDispute", _requestId, _timestamp)
	return *ret0, err
}

// IsInDispute is a free data retrieval call binding the contract method 0x3df0777b.
//
// Solidity: function isInDispute(_requestId uint256, _timestamp uint256) constant returns(bool)
func (_BerryMaster *BerryMasterSession) IsInDispute(_requestId *big.Int, _timestamp *big.Int) (bool, error) {
	return _BerryMaster.Contract.IsInDispute(&_BerryMaster.CallOpts, _requestId, _timestamp)
}

// IsInDispute is a free data retrieval call binding the contract method 0x3df0777b.
//
// Solidity: function isInDispute(_requestId uint256, _timestamp uint256) constant returns(bool)
func (_BerryMaster *BerryMasterCallerSession) IsInDispute(_requestId *big.Int, _timestamp *big.Int) (bool, error) {
	return _BerryMaster.Contract.IsInDispute(&_BerryMaster.CallOpts, _requestId, _timestamp)
}

// RetrieveData is a free data retrieval call binding the contract method 0x93fa4915.
//
// Solidity: function retrieveData(_requestId uint256, _timestamp uint256) constant returns(uint256)
func (_BerryMaster *BerryMasterCaller) RetrieveData(opts *bind.CallOpts, _requestId *big.Int, _timestamp *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BerryMaster.contract.Call(opts, out, "retrieveData", _requestId, _timestamp)
	return *ret0, err
}

// RetrieveData is a free data retrieval call binding the contract method 0x93fa4915.
//
// Solidity: function retrieveData(_requestId uint256, _timestamp uint256) constant returns(uint256)
func (_BerryMaster *BerryMasterSession) RetrieveData(_requestId *big.Int, _timestamp *big.Int) (*big.Int, error) {
	return _BerryMaster.Contract.RetrieveData(&_BerryMaster.CallOpts, _requestId, _timestamp)
}

// RetrieveData is a free data retrieval call binding the contract method 0x93fa4915.
//
// Solidity: function retrieveData(_requestId uint256, _timestamp uint256) constant returns(uint256)
func (_BerryMaster *BerryMasterCallerSession) RetrieveData(_requestId *big.Int, _timestamp *big.Int) (*big.Int, error) {
	return _BerryMaster.Contract.RetrieveData(&_BerryMaster.CallOpts, _requestId, _timestamp)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_BerryMaster *BerryMasterCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BerryMaster.contract.Call(opts, out, "totalSupply")
	return *ret0, err
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_BerryMaster *BerryMasterSession) TotalSupply() (*big.Int, error) {
	return _BerryMaster.Contract.TotalSupply(&_BerryMaster.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_BerryMaster *BerryMasterCallerSession) TotalSupply() (*big.Int, error) {
	return _BerryMaster.Contract.TotalSupply(&_BerryMaster.CallOpts)
}

// ChangeDeity is a paid mutator transaction binding the contract method 0x47abd7f1.
//
// Solidity: function changeDeity(_newDeity address) returns()
func (_BerryMaster *BerryMasterTransactor) ChangeDeity(opts *bind.TransactOpts, _newDeity common.Address) (*types.Transaction, error) {
	return _BerryMaster.contract.Transact(opts, "changeDeity", _newDeity)
}

// ChangeDeity is a paid mutator transaction binding the contract method 0x47abd7f1.
//
// Solidity: function changeDeity(_newDeity address) returns()
func (_BerryMaster *BerryMasterSession) ChangeDeity(_newDeity common.Address) (*types.Transaction, error) {
	return _BerryMaster.Contract.ChangeDeity(&_BerryMaster.TransactOpts, _newDeity)
}

// ChangeDeity is a paid mutator transaction binding the contract method 0x47abd7f1.
//
// Solidity: function changeDeity(_newDeity address) returns()
func (_BerryMaster *BerryMasterTransactorSession) ChangeDeity(_newDeity common.Address) (*types.Transaction, error) {
	return _BerryMaster.Contract.ChangeDeity(&_BerryMaster.TransactOpts, _newDeity)
}

// ChangeBerryContract is a paid mutator transaction binding the contract method 0xae0a8279.
//
// Solidity: function changeBerryContract(_berryContract address) returns()
func (_BerryMaster *BerryMasterTransactor) ChangeBerryContract(opts *bind.TransactOpts, _berryContract common.Address) (*types.Transaction, error) {
	return _BerryMaster.contract.Transact(opts, "changeBerryContract", _berryContract)
}

// ChangeBerryContract is a paid mutator transaction binding the contract method 0xae0a8279.
//
// Solidity: function changeBerryContract(_berryContract address) returns()
func (_BerryMaster *BerryMasterSession) ChangeBerryContract(_berryContract common.Address) (*types.Transaction, error) {
	return _BerryMaster.Contract.ChangeBerryContract(&_BerryMaster.TransactOpts, _berryContract)
}

// ChangeBerryContract is a paid mutator transaction binding the contract method 0xae0a8279.
//
// Solidity: function changeBerryContract(_berryContract address) returns()
func (_BerryMaster *BerryMasterTransactorSession) ChangeBerryContract(_berryContract common.Address) (*types.Transaction, error) {
	return _BerryMaster.Contract.ChangeBerryContract(&_BerryMaster.TransactOpts, _berryContract)
}

// BerryMasterNewBerryAddressIterator is returned from FilterNewBerryAddress and is used to iterate over the raw logs and unpacked data for NewBerryAddress events raised by the BerryMaster contract.
type BerryMasterNewBerryAddressIterator struct {
	Event *BerryMasterNewBerryAddress // Event containing the contract specifics and raw log

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
func (it *BerryMasterNewBerryAddressIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BerryMasterNewBerryAddress)
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
		it.Event = new(BerryMasterNewBerryAddress)
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
func (it *BerryMasterNewBerryAddressIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BerryMasterNewBerryAddressIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BerryMasterNewBerryAddress represents a NewBerryAddress event raised by the BerryMaster contract.
type BerryMasterNewBerryAddress struct {
	NewBerry common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterNewBerryAddress is a free log retrieval operation binding the contract event 0xc2d1449eb0b6547aa426e09d9942a77fa4fc8cd3296305b3163e22452e0bcb8d.
//
// Solidity: e NewBerryAddress(_newBerry address)
func (_BerryMaster *BerryMasterFilterer) FilterNewBerryAddress(opts *bind.FilterOpts) (*BerryMasterNewBerryAddressIterator, error) {

	logs, sub, err := _BerryMaster.contract.FilterLogs(opts, "NewBerryAddress")
	if err != nil {
		return nil, err
	}
	return &BerryMasterNewBerryAddressIterator{contract: _BerryMaster.contract, event: "NewBerryAddress", logs: logs, sub: sub}, nil
}

// WatchNewBerryAddress is a free log subscription operation binding the contract event 0xc2d1449eb0b6547aa426e09d9942a77fa4fc8cd3296305b3163e22452e0bcb8d.
//
// Solidity: e NewBerryAddress(_newBerry address)
func (_BerryMaster *BerryMasterFilterer) WatchNewBerryAddress(opts *bind.WatchOpts, sink chan<- *BerryMasterNewBerryAddress) (event.Subscription, error) {

	logs, sub, err := _BerryMaster.contract.WatchLogs(opts, "NewBerryAddress")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BerryMasterNewBerryAddress)
				if err := _BerryMaster.contract.UnpackLog(event, "NewBerryAddress", log); err != nil {
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
