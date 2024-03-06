// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package statelib

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
	_ = abi.ConvertType
)

// StatelibMetaData contains all meta data concerning the Statelib contract.
var StatelibMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"blockN\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"state\",\"type\":\"uint256\"}],\"name\":\"StateUpdated\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"ID_HISTORY_RETURN_LIMIT\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x6115e0610053600b82828239805160001a607314610046577f4e487b7100000000000000000000000000000000000000000000000000000000600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600436106100b35760003560e01c80637b87f75d1161007b5780637b87f75d1461019a578063893e9739146101ca5780639da68946146101fa578063b1abcc4d1461022a578063b2a11b2b1461025a578063eaa6b26c1461028a576100b3565b80630e21d5fd146100b8578063419e76c1146100e857806344cfc45e146101185780634a2042181461014157806365fe1b5b1461016a575b600080fd5b6100d260048036038101906100cd9190610d48565b6102a8565b6040516100df9190610d97565b60405180910390f35b61010260048036038101906100fd9190610d48565b610318565b60405161010f9190610dcd565b60405180910390f35b81801561012457600080fd5b5061013f600480360381019061013a9190610de8565b61033d565b005b81801561014d57600080fd5b5061016860048036038101906101639190610de8565b61034f565b005b610184600480360381019061017f9190610e3b565b6103ac565b6040516101919190610fee565b60405180910390f35b6101b460048036038101906101af9190611010565b6104ef565b6040516101c19190610fee565b60405180910390f35b6101e460048036038101906101df9190610de8565b61066a565b6040516101f19190611119565b60405180910390f35b610214600480360381019061020f9190610d48565b6106d6565b6040516102219190611119565b60405180910390f35b610244600480360381019061023f9190610de8565b6107ee565b6040516102519190610dcd565b60405180910390f35b610274600480360381019061026f9190610de8565b610825565b6040516102819190610d97565b60405180910390f35b61029261085a565b60405161029f9190610d97565b60405180910390f35b600082826102b68282610318565b6102f5576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016102ec90611191565b60405180910390fd5b846000016000858152602001908152602001600020805490509250505092915050565b6000808360000160008481526020019081526020016000208054905011905092915050565b61034a8383834243610860565b505050565b6103598383610318565b15610399576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161039090611249565b60405180910390fd5b6103a7838383600080610860565b505050565b606084846103ba8282610318565b6103f9576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016103f090611191565b60405180910390fd5b6000806104228960000160008a81526020019081526020016000208054905088886103e8610980565b91509150600082826104349190611298565b67ffffffffffffffff81111561044d5761044c6112cc565b5b60405190808252806020026020018201604052801561048657816020015b610473610c9a565b81526020019060019003908161046b5790505b50905060008390505b828110156104de576104a28b8b83610a77565b8285836104af9190611298565b815181106104c0576104bf6112fb565b5b602002602001018190525080806104d69061132a565b91505061048f565b508095505050505050949350505050565b60608585856104ff8383836107ee565b61053e576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610535906113be565b60405180910390fd5b60008960010160008a81526020019081526020016000206000898152602001908152602001600020905060008061057d83805490508a8a6103e8610980565b915091506000828261058f9190611298565b67ffffffffffffffff8111156105a8576105a76112cc565b5b6040519080825280602002602001820160405280156105e157816020015b6105ce610c9a565b8152602001906001900390816105c65790505b50905060008390505b828110156106565761061a8e8e87848154811061060a576106096112fb565b5b9060005260206000200154610a77565b8285836106279190611298565b81518110610638576106376112fb565b5b6020026020010181905250808061064e9061132a565b9150506105ea565b508097505050505050505095945050505050565b610672610c9a565b8383836106808383836107ee565b6106bf576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016106b6906113be565b60405180910390fd5b6106ca878787610c20565b93505050509392505050565b6106de610c9a565b82826106ea8282610318565b610729576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161072090611191565b60405180910390fd5b60008560000160008681526020019081526020016000209050600081600183805490506107569190611298565b81548110610767576107666112fb565b5b9060005260206000209060030201604051806060016040529081600082015481526020016001820154815260200160028201548152505090506040518060e001604052808781526020018260000151815260200160008152602001826020015181526020016000815260200182604001518152602001600081525094505050505092915050565b6000808460010160008581526020019081526020016000206000848152602001908152602001600020805490501190509392505050565b600083600101600084815260200190815260200160002060008381526020019081526020016000208054905090509392505050565b6103e881565b6000856000016000868152602001908152602001600020905080604051806060016040528086815260200185815260200184815250908060018154018082558091505060019003906000526020600020906003020160009091909190915060008201518160000155602082015181600101556040820151816002015550508560010160008681526020019081526020016000206000858152602001908152602001600020600182805490506109159190611298565b90806001815401808255809150506001900390600052602060002001600090919091909150557f88aef4d78ad30d12a12a98e96007f5b09c1610b5364b2b99960b7d07e00a88388583858760405161097094939291906113ed565b60405180910390a1505050505050565b600080600084116109c6576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016109bd9061147e565b60405180910390fd5b82841115610a09576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610a00906114ea565b60405180910390fd5b858510610a4b576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610a4290611556565b60405180910390fd5b60008486610a599190611576565b905086811115610a67578690505b8581925092505094509492505050565b610a7f610c9a565b6000600185600001600086815260200190815260200160002080549050610aa69190611298565b8314905060008560000160008681526020019081526020016000208481548110610ad357610ad26112fb565b5b906000526020600020906003020190506040518060e001604052808681526020018260000154815260200183610b4d57876000016000888152602001908152602001600020600187610b259190611576565b81548110610b3657610b356112fb565b5b906000526020600020906003020160000154610b50565b60005b81526020018260010154815260200183610bae57876000016000888152602001908152602001600020600187610b869190611576565b81548110610b9757610b966112fb565b5b906000526020600020906003020160010154610bb1565b60005b81526020018260020154815260200183610c0f57876000016000888152602001908152602001600020600187610be79190611576565b81548110610bf857610bf76112fb565b5b906000526020600020906003020160020154610c12565b60005b815250925050509392505050565b610c28610c9a565b60008460010160008581526020019081526020016000206000848152602001908152602001600020905060008160018380549050610c669190611298565b81548110610c7757610c766112fb565b5b90600052602060002001549050610c8f868683610a77565b925050509392505050565b6040518060e00160405280600081526020016000815260200160008152602001600081526020016000815260200160008152602001600081525090565b600080fd5b6000819050919050565b610cef81610cdc565b8114610cfa57600080fd5b50565b600081359050610d0c81610ce6565b92915050565b6000819050919050565b610d2581610d12565b8114610d3057600080fd5b50565b600081359050610d4281610d1c565b92915050565b60008060408385031215610d5f57610d5e610cd7565b5b6000610d6d85828601610cfd565b9250506020610d7e85828601610d33565b9150509250929050565b610d9181610d12565b82525050565b6000602082019050610dac6000830184610d88565b92915050565b60008115159050919050565b610dc781610db2565b82525050565b6000602082019050610de26000830184610dbe565b92915050565b600080600060608486031215610e0157610e00610cd7565b5b6000610e0f86828701610cfd565b9350506020610e2086828701610d33565b9250506040610e3186828701610d33565b9150509250925092565b60008060008060808587031215610e5557610e54610cd7565b5b6000610e6387828801610cfd565b9450506020610e7487828801610d33565b9350506040610e8587828801610d33565b9250506060610e9687828801610d33565b91505092959194509250565b600081519050919050565b600082825260208201905092915050565b6000819050602082019050919050565b610ed781610d12565b82525050565b60e082016000820151610ef36000850182610ece565b506020820151610f066020850182610ece565b506040820151610f196040850182610ece565b506060820151610f2c6060850182610ece565b506080820151610f3f6080850182610ece565b5060a0820151610f5260a0850182610ece565b5060c0820151610f6560c0850182610ece565b50505050565b6000610f778383610edd565b60e08301905092915050565b6000602082019050919050565b6000610f9b82610ea2565b610fa58185610ead565b9350610fb083610ebe565b8060005b83811015610fe1578151610fc88882610f6b565b9750610fd383610f83565b925050600181019050610fb4565b5085935050505092915050565b600060208201905081810360008301526110088184610f90565b905092915050565b600080600080600060a0868803121561102c5761102b610cd7565b5b600061103a88828901610cfd565b955050602061104b88828901610d33565b945050604061105c88828901610d33565b935050606061106d88828901610d33565b925050608061107e88828901610d33565b9150509295509295909350565b60e0820160008201516110a16000850182610ece565b5060208201516110b46020850182610ece565b5060408201516110c76040850182610ece565b5060608201516110da6060850182610ece565b5060808201516110ed6080850182610ece565b5060a082015161110060a0850182610ece565b5060c082015161111360c0850182610ece565b50505050565b600060e08201905061112e600083018461108b565b92915050565b600082825260208201905092915050565b7f4964656e7469747920646f6573206e6f74206578697374000000000000000000600082015250565b600061117b601783611134565b915061118682611145565b602082019050919050565b600060208201905081810360008301526111aa8161116e565b9050919050565b7f5a65726f2074696d657374616d7020616e6420626c6f636b2073686f756c642060008201527f6265206f6e6c7920696e20746865206669727374206964656e7469747920737460208201527f6174650000000000000000000000000000000000000000000000000000000000604082015250565b6000611233604383611134565b915061123e826111b1565b606082019050919050565b6000602082019050818103600083015261126281611226565b9050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b60006112a382610d12565b91506112ae83610d12565b92508282039050818111156112c6576112c5611269565b5b92915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b600061133582610d12565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff820361136757611366611269565b5b600182019050919050565b7f537461746520646f6573206e6f74206578697374000000000000000000000000600082015250565b60006113a8601483611134565b91506113b382611372565b602082019050919050565b600060208201905081810360008301526113d78161139b565b9050919050565b6113e781610d12565b82525050565b600060808201905061140260008301876113de565b61140f60208301866113de565b61141c60408301856113de565b61142960608301846113de565b95945050505050565b7f4c656e6774682073686f756c642062652067726561746572207468616e203000600082015250565b6000611468601f83611134565b915061147382611432565b602082019050919050565b600060208201905081810360008301526114978161145b565b9050919050565b7f4c656e677468206c696d69742065786365656465640000000000000000000000600082015250565b60006114d4601583611134565b91506114df8261149e565b602082019050919050565b60006020820190508181036000830152611503816114c7565b9050919050565b7f537461727420696e646578206f7574206f6620626f756e647300000000000000600082015250565b6000611540601983611134565b915061154b8261150a565b602082019050919050565b6000602082019050818103600083015261156f81611533565b9050919050565b600061158182610d12565b915061158c83610d12565b92508282019050808211156115a4576115a3611269565b5b9291505056fea2646970667358221220f24a682a0b703044b28e4db9fc2f25b2e0d0e856bee4abc1d9d985962c2daa2864736f6c63430008100033",
}

// StatelibABI is the input ABI used to generate the binding from.
// Deprecated: Use StatelibMetaData.ABI instead.
var StatelibABI = StatelibMetaData.ABI

// StatelibBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use StatelibMetaData.Bin instead.
var StatelibBin = StatelibMetaData.Bin

// DeployStatelib deploys a new Ethereum contract, binding an instance of Statelib to it.
func DeployStatelib(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Statelib, error) {
	parsed, err := StatelibMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(StatelibBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Statelib{StatelibCaller: StatelibCaller{contract: contract}, StatelibTransactor: StatelibTransactor{contract: contract}, StatelibFilterer: StatelibFilterer{contract: contract}}, nil
}

// Statelib is an auto generated Go binding around an Ethereum contract.
type Statelib struct {
	StatelibCaller     // Read-only binding to the contract
	StatelibTransactor // Write-only binding to the contract
	StatelibFilterer   // Log filterer for contract events
}

// StatelibCaller is an auto generated read-only Go binding around an Ethereum contract.
type StatelibCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StatelibTransactor is an auto generated write-only Go binding around an Ethereum contract.
type StatelibTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StatelibFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type StatelibFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StatelibSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type StatelibSession struct {
	Contract     *Statelib         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// StatelibCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type StatelibCallerSession struct {
	Contract *StatelibCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// StatelibTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type StatelibTransactorSession struct {
	Contract     *StatelibTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// StatelibRaw is an auto generated low-level Go binding around an Ethereum contract.
type StatelibRaw struct {
	Contract *Statelib // Generic contract binding to access the raw methods on
}

// StatelibCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type StatelibCallerRaw struct {
	Contract *StatelibCaller // Generic read-only contract binding to access the raw methods on
}

// StatelibTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type StatelibTransactorRaw struct {
	Contract *StatelibTransactor // Generic write-only contract binding to access the raw methods on
}

// NewStatelib creates a new instance of Statelib, bound to a specific deployed contract.
func NewStatelib(address common.Address, backend bind.ContractBackend) (*Statelib, error) {
	contract, err := bindStatelib(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Statelib{StatelibCaller: StatelibCaller{contract: contract}, StatelibTransactor: StatelibTransactor{contract: contract}, StatelibFilterer: StatelibFilterer{contract: contract}}, nil
}

// NewStatelibCaller creates a new read-only instance of Statelib, bound to a specific deployed contract.
func NewStatelibCaller(address common.Address, caller bind.ContractCaller) (*StatelibCaller, error) {
	contract, err := bindStatelib(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &StatelibCaller{contract: contract}, nil
}

// NewStatelibTransactor creates a new write-only instance of Statelib, bound to a specific deployed contract.
func NewStatelibTransactor(address common.Address, transactor bind.ContractTransactor) (*StatelibTransactor, error) {
	contract, err := bindStatelib(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &StatelibTransactor{contract: contract}, nil
}

// NewStatelibFilterer creates a new log filterer instance of Statelib, bound to a specific deployed contract.
func NewStatelibFilterer(address common.Address, filterer bind.ContractFilterer) (*StatelibFilterer, error) {
	contract, err := bindStatelib(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &StatelibFilterer{contract: contract}, nil
}

// bindStatelib binds a generic wrapper to an already deployed contract.
func bindStatelib(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := StatelibMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Statelib *StatelibRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Statelib.Contract.StatelibCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Statelib *StatelibRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Statelib.Contract.StatelibTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Statelib *StatelibRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Statelib.Contract.StatelibTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Statelib *StatelibCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Statelib.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Statelib *StatelibTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Statelib.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Statelib *StatelibTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Statelib.Contract.contract.Transact(opts, method, params...)
}

// IDHISTORYRETURNLIMIT is a free data retrieval call binding the contract method 0xeaa6b26c.
//
// Solidity: function ID_HISTORY_RETURN_LIMIT() view returns(uint256)
func (_Statelib *StatelibCaller) IDHISTORYRETURNLIMIT(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Statelib.contract.Call(opts, &out, "ID_HISTORY_RETURN_LIMIT")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// IDHISTORYRETURNLIMIT is a free data retrieval call binding the contract method 0xeaa6b26c.
//
// Solidity: function ID_HISTORY_RETURN_LIMIT() view returns(uint256)
func (_Statelib *StatelibSession) IDHISTORYRETURNLIMIT() (*big.Int, error) {
	return _Statelib.Contract.IDHISTORYRETURNLIMIT(&_Statelib.CallOpts)
}

// IDHISTORYRETURNLIMIT is a free data retrieval call binding the contract method 0xeaa6b26c.
//
// Solidity: function ID_HISTORY_RETURN_LIMIT() view returns(uint256)
func (_Statelib *StatelibCallerSession) IDHISTORYRETURNLIMIT() (*big.Int, error) {
	return _Statelib.Contract.IDHISTORYRETURNLIMIT(&_Statelib.CallOpts)
}

// StatelibStateUpdatedIterator is returned from FilterStateUpdated and is used to iterate over the raw logs and unpacked data for StateUpdated events raised by the Statelib contract.
type StatelibStateUpdatedIterator struct {
	Event *StatelibStateUpdated // Event containing the contract specifics and raw log

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
func (it *StatelibStateUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StatelibStateUpdated)
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
		it.Event = new(StatelibStateUpdated)
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
func (it *StatelibStateUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StatelibStateUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StatelibStateUpdated represents a StateUpdated event raised by the Statelib contract.
type StatelibStateUpdated struct {
	Id        *big.Int
	BlockN    *big.Int
	Timestamp *big.Int
	State     *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterStateUpdated is a free log retrieval operation binding the contract event 0x88aef4d78ad30d12a12a98e96007f5b09c1610b5364b2b99960b7d07e00a8838.
//
// Solidity: event StateUpdated(uint256 id, uint256 blockN, uint256 timestamp, uint256 state)
func (_Statelib *StatelibFilterer) FilterStateUpdated(opts *bind.FilterOpts) (*StatelibStateUpdatedIterator, error) {

	logs, sub, err := _Statelib.contract.FilterLogs(opts, "StateUpdated")
	if err != nil {
		return nil, err
	}
	return &StatelibStateUpdatedIterator{contract: _Statelib.contract, event: "StateUpdated", logs: logs, sub: sub}, nil
}

// WatchStateUpdated is a free log subscription operation binding the contract event 0x88aef4d78ad30d12a12a98e96007f5b09c1610b5364b2b99960b7d07e00a8838.
//
// Solidity: event StateUpdated(uint256 id, uint256 blockN, uint256 timestamp, uint256 state)
func (_Statelib *StatelibFilterer) WatchStateUpdated(opts *bind.WatchOpts, sink chan<- *StatelibStateUpdated) (event.Subscription, error) {

	logs, sub, err := _Statelib.contract.WatchLogs(opts, "StateUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StatelibStateUpdated)
				if err := _Statelib.contract.UnpackLog(event, "StateUpdated", log); err != nil {
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

// ParseStateUpdated is a log parse operation binding the contract event 0x88aef4d78ad30d12a12a98e96007f5b09c1610b5364b2b99960b7d07e00a8838.
//
// Solidity: event StateUpdated(uint256 id, uint256 blockN, uint256 timestamp, uint256 state)
func (_Statelib *StatelibFilterer) ParseStateUpdated(log types.Log) (*StatelibStateUpdated, error) {
	event := new(StatelibStateUpdated)
	if err := _Statelib.contract.UnpackLog(event, "StateUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
