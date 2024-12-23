// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package verifiersigwrapper

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

// VerifiersigwrapperMetaData contains all meta data concerning the Verifiersigwrapper contract.
var VerifiersigwrapperMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256[2]\",\"name\":\"a\",\"type\":\"uint256[2]\"},{\"internalType\":\"uint256[2][2]\",\"name\":\"b\",\"type\":\"uint256[2][2]\"},{\"internalType\":\"uint256[2]\",\"name\":\"c\",\"type\":\"uint256[2]\"},{\"internalType\":\"uint256[]\",\"name\":\"input\",\"type\":\"uint256[]\"}],\"name\":\"verify\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"r\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[2]\",\"name\":\"_pA\",\"type\":\"uint256[2]\"},{\"internalType\":\"uint256[2][2]\",\"name\":\"_pB\",\"type\":\"uint256[2][2]\"},{\"internalType\":\"uint256[2]\",\"name\":\"_pC\",\"type\":\"uint256[2]\"},{\"internalType\":\"uint256[11]\",\"name\":\"_pubSignals\",\"type\":\"uint256[11]\"}],\"name\":\"verifyProof\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b50610f85806100206000396000f3fe608060405234801561001057600080fd5b50600436106100365760003560e01c80632612907c1461003b578063b9c6ea871461006b575b600080fd5b61005560048036038101906100509190610ab4565b61009b565b6040516100629190610b59565b60405180910390f35b61008560048036038101906100809190610b96565b6101cd565b6040516100929190610b59565b60405180910390f35b60006100a56109de565b600b84849050146100eb576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016100e290610c5c565b60405180910390fd5b60005b600b8110156101405784848281811061010a57610109610c7c565b5b905060200201358282600b811061012457610123610c7c565b5b602002018181525050808061013890610ce4565b9150506100ee565b503073ffffffffffffffffffffffffffffffffffffffff1663b9c6ea87888888856040518563ffffffff1660e01b81526004016101809493929190610eaf565b602060405180830381865afa15801561019d573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906101c19190610f22565b91505095945050505050565b600061091c565b7f30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd478110610205576000805260206000f35b50565b600060405183815284602082015285604082015260408160608360076107d05a03fa91508161023b576000805260206000f35b825160408201526020830151606082015260408360808360066107d05a03fa91508161026b576000805260206000f35b505050505050565b600060808601600087017f0992267e58f93dac7a7beb752274c98029cfc357435f878ea72b030d1085100a81527ea1152a41bb3aa4fbe1aa7882c2e2a8e1b6e968f077ff4af210d42e0d7072b5602082015261031560008801357f1622d9ce6cd690e847704358fe99032bcf337ba5c8f17181b36216c8a61b89b47f11788a050c473369bbac1d2e011ec3236dc2e426bf33950035d783539419e03b84610208565b61036560208801357f0b70d29ed3ae4b77ab974d5624d17cef7bc8a6c91d157a36637dbc85a311a5977f19dcba7a8a64168495e184fa3ce0be0d08ab53712d2bb51ebd016cab0679c89b84610208565b6103b560408801357f10412904e7135f93df873a768f47e26eaf7649cf526233820d4e705ee7cc81687f03d16811819c4ff82df949af146566ada7540c0e620d0a3a243927b0a4f6b0db84610208565b61040560608801357f0a837cd8825cf5fa6feeb51b4b85f3c411db9db2df64da913a0fedc14d12d77d7f161628f234fb50ecce9c863ea7a616d4790276ed14f410062964ce932e77bf0684610208565b61045560808801357f10a808acd38d95f293e7595ac1d254314023fe5a4b417d46a457216250683b1a7f233e0ce24b641fa92543e2251944350e1ba373aff7a2f6296403855ba1163b2d84610208565b6104a560a08801357f20cf4b8faf4cc3ddef5428e3ec7ac6ab8fac15d3a408ab9be92c7c3636959bff7f0c2d90736caddcdbe291f31a2543dd6edaa258a516960ffb9c3cea7e2da18d1584610208565b6104f560c08801357f24ea85494a29b0493f759cdecc86f90166567ba7fa3dfafc612441df821298707f2af675475a9a8c80732328cca338fdc7b0d082049982557e0e663a9fd85ce44484610208565b61054560e08801357f070d9dbbd07134b5bbee59c1d7f8e3d631430dd6394698fbba5f53384848bdf07f0da6d80e9f321df5ac17ec71cf7b1d5a7564f79b51be0c74bedb8c2266a4ba8e84610208565b6105966101008801357f1d6c48d0c956333d81951c4ec0240786f431aa7899010a6360bd50a063dd50977f2d23bb3f195568e6d21110df201b9520e2a0e49f433ac618afc3536ebef3adb884610208565b6105e76101208801357f105e3dd896cb03b12043762dce2f72ea2a58a985fea546ba5bbd115d0bef7aaa7f032eda83efb2182f8abd68311a50bd1c1acde67ab2959d45fbc26941639e5ad484610208565b6106386101408801357f21ef8dfeae61f0040d207a37d86209217a484587294fcde52f50d2095167820f7f293ec4a0dd36216f02bcdc9c329cffaa450d2b5a44f55b61cabfc8300929a92a84610208565b833582527f30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd4760208501357f30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd4703066020830152843560408301526020850135606083015260408501356080830152606085013560a08301527f2d4d9aa7e302d9df41749d5507949d05dbea33fbb16c643b22f599a2be6df2e260c08301527f14bedd503c37ceb061d8ec60209fe345ce89830a19230301f076caff004d192660e08301527f0967032fcbf776d1afc985f88877f182d38480a653f2decaa9794cbc3bf3060c6101008301527f0e187847ad4c798374d0d6732bf501847dd68bc0e071241e0213bc7fc13db7ab6101208301527f304cfbd1e08a704a99f5e847d93f8c3caafddec46b7a0d379da69a4d112346a76101408301527f1739c1b1a457a8c7313123d24d2f9192f896b7c63eea05a9d57f06547ad0cec8610160830152600088015161018083015260206000018801516101a08301527f198e9393920d483a7260bfb731fb5d25f1aa493335a9e71297e485b7aef312c26101c08301527f1800deef121f1e76426a00665e5c4479674322d4f75edadd46debd5cd992f6ed6101e08301527f090689d0585ff075ec9e99ad690c3395bc4b313370b38ef355acdadcd122975b6102008301527f12c85ea5db8c6deb4aab71808dcb408fe3d1e7690c43d37b4ce6cc0166fa7daa610220830152853561024083015260208601356102608301527f1469e40d51ed9ca8f5677275a3111e52ef96f34c58c29d156f870870109074296102808301527f03c7e502253d6cfe7a9314775ac2d0fe4704763560fd70e78bec59fd4f6da94b6102a08301527f2b90580b4fd31c34ac334b8581fd0de1b6d8af80163ae34b60a5fc851797a24b6102c08301527f2a1bf8e0abf527ccfb58440a7364da4944f71aac2c9b5faa5c9a26f37034f2cc6102e08301526020826103008460086107d05a03fa82518116935050505095945050505050565b604051610380810160405261093460008401356101d4565b61094160208401356101d4565b61094e60408401356101d4565b61095b60608401356101d4565b61096860808401356101d4565b61097560a08401356101d4565b61098260c08401356101d4565b61098f60e08401356101d4565b61099d6101008401356101d4565b6109ab6101208401356101d4565b6109b96101408401356101d4565b6109c76101608401356101d4565b6109d4818486888a610273565b8060005260206000f35b604051806101600160405280600b90602082028036833780820191505090505090565b600080fd5b600080fd5b600080fd5b600081905082602060020282011115610a2c57610a2b610a0b565b5b92915050565b600081905082604060020282011115610a4e57610a4d610a0b565b5b92915050565b600080fd5b600080fd5b60008083601f840112610a7457610a73610a54565b5b8235905067ffffffffffffffff811115610a9157610a90610a59565b5b602083019150836020820283011115610aad57610aac610a0b565b5b9250929050565b60008060008060006101208688031215610ad157610ad0610a01565b5b6000610adf88828901610a10565b9550506040610af088828901610a32565b94505060c0610b0188828901610a10565b93505061010086013567ffffffffffffffff811115610b2357610b22610a06565b5b610b2f88828901610a5e565b92509250509295509295909350565b60008115159050919050565b610b5381610b3e565b82525050565b6000602082019050610b6e6000830184610b4a565b92915050565b6000819050826020600b0282011115610b9057610b8f610a0b565b5b92915050565b6000806000806102608587031215610bb157610bb0610a01565b5b6000610bbf87828801610a10565b9450506040610bd087828801610a32565b93505060c0610be187828801610a10565b925050610100610bf387828801610b74565b91505092959194509250565b600082825260208201905092915050565b7f6578706563746564206172726179206c656e6774682069732031310000000000600082015250565b6000610c46601b83610bff565b9150610c5182610c10565b602082019050919050565b60006020820190508181036000830152610c7581610c39565b9050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b6000819050919050565b6000610cef82610cda565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8203610d2157610d20610cab565b5b600182019050919050565b82818337505050565b610d4160408383610d2c565b5050565b600060029050919050565b600081905092915050565b6000819050919050565b610d7160408383610d2c565b5050565b6000610d818383610d65565b60408301905092915050565b600082905092915050565b6000604082019050919050565b610dae81610d45565b610db88184610d50565b9250610dc382610d5b565b8060005b83811015610dfc57610dd98284610d8d565b610de38782610d75565b9650610dee83610d98565b925050600181019050610dc7565b505050505050565b6000600b9050919050565b600081905092915050565b6000819050919050565b610e2d81610cda565b82525050565b6000610e3f8383610e24565b60208301905092915050565b6000602082019050919050565b610e6181610e04565b610e6b8184610e0f565b9250610e7682610e1a565b8060005b83811015610ea7578151610e8e8782610e33565b9650610e9983610e4b565b925050600181019050610e7a565b505050505050565b600061026082019050610ec56000830187610d35565b610ed26040830186610da5565b610edf60c0830185610d35565b610eed610100830184610e58565b95945050505050565b610eff81610b3e565b8114610f0a57600080fd5b50565b600081519050610f1c81610ef6565b92915050565b600060208284031215610f3857610f37610a01565b5b6000610f4684828501610f0d565b9150509291505056fea26469706673582212203d97d159d5ba30b2a99031d5ed9d9e2459bc13c5deece8ecd1f9e337f6b398d864736f6c63430008100033",
}

// VerifiersigwrapperABI is the input ABI used to generate the binding from.
// Deprecated: Use VerifiersigwrapperMetaData.ABI instead.
var VerifiersigwrapperABI = VerifiersigwrapperMetaData.ABI

// VerifiersigwrapperBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use VerifiersigwrapperMetaData.Bin instead.
var VerifiersigwrapperBin = VerifiersigwrapperMetaData.Bin

// DeployVerifiersigwrapper deploys a new Ethereum contract, binding an instance of Verifiersigwrapper to it.
func DeployVerifiersigwrapper(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Verifiersigwrapper, error) {
	parsed, err := VerifiersigwrapperMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(VerifiersigwrapperBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Verifiersigwrapper{VerifiersigwrapperCaller: VerifiersigwrapperCaller{contract: contract}, VerifiersigwrapperTransactor: VerifiersigwrapperTransactor{contract: contract}, VerifiersigwrapperFilterer: VerifiersigwrapperFilterer{contract: contract}}, nil
}

// Verifiersigwrapper is an auto generated Go binding around an Ethereum contract.
type Verifiersigwrapper struct {
	VerifiersigwrapperCaller     // Read-only binding to the contract
	VerifiersigwrapperTransactor // Write-only binding to the contract
	VerifiersigwrapperFilterer   // Log filterer for contract events
}

// VerifiersigwrapperCaller is an auto generated read-only Go binding around an Ethereum contract.
type VerifiersigwrapperCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VerifiersigwrapperTransactor is an auto generated write-only Go binding around an Ethereum contract.
type VerifiersigwrapperTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VerifiersigwrapperFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type VerifiersigwrapperFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VerifiersigwrapperSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type VerifiersigwrapperSession struct {
	Contract     *Verifiersigwrapper // Generic contract binding to set the session for
	CallOpts     bind.CallOpts       // Call options to use throughout this session
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// VerifiersigwrapperCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type VerifiersigwrapperCallerSession struct {
	Contract *VerifiersigwrapperCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts             // Call options to use throughout this session
}

// VerifiersigwrapperTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type VerifiersigwrapperTransactorSession struct {
	Contract     *VerifiersigwrapperTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// VerifiersigwrapperRaw is an auto generated low-level Go binding around an Ethereum contract.
type VerifiersigwrapperRaw struct {
	Contract *Verifiersigwrapper // Generic contract binding to access the raw methods on
}

// VerifiersigwrapperCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type VerifiersigwrapperCallerRaw struct {
	Contract *VerifiersigwrapperCaller // Generic read-only contract binding to access the raw methods on
}

// VerifiersigwrapperTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type VerifiersigwrapperTransactorRaw struct {
	Contract *VerifiersigwrapperTransactor // Generic write-only contract binding to access the raw methods on
}

// NewVerifiersigwrapper creates a new instance of Verifiersigwrapper, bound to a specific deployed contract.
func NewVerifiersigwrapper(address common.Address, backend bind.ContractBackend) (*Verifiersigwrapper, error) {
	contract, err := bindVerifiersigwrapper(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Verifiersigwrapper{VerifiersigwrapperCaller: VerifiersigwrapperCaller{contract: contract}, VerifiersigwrapperTransactor: VerifiersigwrapperTransactor{contract: contract}, VerifiersigwrapperFilterer: VerifiersigwrapperFilterer{contract: contract}}, nil
}

// NewVerifiersigwrapperCaller creates a new read-only instance of Verifiersigwrapper, bound to a specific deployed contract.
func NewVerifiersigwrapperCaller(address common.Address, caller bind.ContractCaller) (*VerifiersigwrapperCaller, error) {
	contract, err := bindVerifiersigwrapper(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &VerifiersigwrapperCaller{contract: contract}, nil
}

// NewVerifiersigwrapperTransactor creates a new write-only instance of Verifiersigwrapper, bound to a specific deployed contract.
func NewVerifiersigwrapperTransactor(address common.Address, transactor bind.ContractTransactor) (*VerifiersigwrapperTransactor, error) {
	contract, err := bindVerifiersigwrapper(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &VerifiersigwrapperTransactor{contract: contract}, nil
}

// NewVerifiersigwrapperFilterer creates a new log filterer instance of Verifiersigwrapper, bound to a specific deployed contract.
func NewVerifiersigwrapperFilterer(address common.Address, filterer bind.ContractFilterer) (*VerifiersigwrapperFilterer, error) {
	contract, err := bindVerifiersigwrapper(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &VerifiersigwrapperFilterer{contract: contract}, nil
}

// bindVerifiersigwrapper binds a generic wrapper to an already deployed contract.
func bindVerifiersigwrapper(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := VerifiersigwrapperMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Verifiersigwrapper *VerifiersigwrapperRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Verifiersigwrapper.Contract.VerifiersigwrapperCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Verifiersigwrapper *VerifiersigwrapperRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Verifiersigwrapper.Contract.VerifiersigwrapperTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Verifiersigwrapper *VerifiersigwrapperRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Verifiersigwrapper.Contract.VerifiersigwrapperTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Verifiersigwrapper *VerifiersigwrapperCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Verifiersigwrapper.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Verifiersigwrapper *VerifiersigwrapperTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Verifiersigwrapper.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Verifiersigwrapper *VerifiersigwrapperTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Verifiersigwrapper.Contract.contract.Transact(opts, method, params...)
}

// Verify is a free data retrieval call binding the contract method 0x2612907c.
//
// Solidity: function verify(uint256[2] a, uint256[2][2] b, uint256[2] c, uint256[] input) view returns(bool r)
func (_Verifiersigwrapper *VerifiersigwrapperCaller) Verify(opts *bind.CallOpts, a [2]*big.Int, b [2][2]*big.Int, c [2]*big.Int, input []*big.Int) (bool, error) {
	var out []interface{}
	err := _Verifiersigwrapper.contract.Call(opts, &out, "verify", a, b, c, input)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Verify is a free data retrieval call binding the contract method 0x2612907c.
//
// Solidity: function verify(uint256[2] a, uint256[2][2] b, uint256[2] c, uint256[] input) view returns(bool r)
func (_Verifiersigwrapper *VerifiersigwrapperSession) Verify(a [2]*big.Int, b [2][2]*big.Int, c [2]*big.Int, input []*big.Int) (bool, error) {
	return _Verifiersigwrapper.Contract.Verify(&_Verifiersigwrapper.CallOpts, a, b, c, input)
}

// Verify is a free data retrieval call binding the contract method 0x2612907c.
//
// Solidity: function verify(uint256[2] a, uint256[2][2] b, uint256[2] c, uint256[] input) view returns(bool r)
func (_Verifiersigwrapper *VerifiersigwrapperCallerSession) Verify(a [2]*big.Int, b [2][2]*big.Int, c [2]*big.Int, input []*big.Int) (bool, error) {
	return _Verifiersigwrapper.Contract.Verify(&_Verifiersigwrapper.CallOpts, a, b, c, input)
}

// VerifyProof is a free data retrieval call binding the contract method 0xb9c6ea87.
//
// Solidity: function verifyProof(uint256[2] _pA, uint256[2][2] _pB, uint256[2] _pC, uint256[11] _pubSignals) view returns(bool)
func (_Verifiersigwrapper *VerifiersigwrapperCaller) VerifyProof(opts *bind.CallOpts, _pA [2]*big.Int, _pB [2][2]*big.Int, _pC [2]*big.Int, _pubSignals [11]*big.Int) (bool, error) {
	var out []interface{}
	err := _Verifiersigwrapper.contract.Call(opts, &out, "verifyProof", _pA, _pB, _pC, _pubSignals)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// VerifyProof is a free data retrieval call binding the contract method 0xb9c6ea87.
//
// Solidity: function verifyProof(uint256[2] _pA, uint256[2][2] _pB, uint256[2] _pC, uint256[11] _pubSignals) view returns(bool)
func (_Verifiersigwrapper *VerifiersigwrapperSession) VerifyProof(_pA [2]*big.Int, _pB [2][2]*big.Int, _pC [2]*big.Int, _pubSignals [11]*big.Int) (bool, error) {
	return _Verifiersigwrapper.Contract.VerifyProof(&_Verifiersigwrapper.CallOpts, _pA, _pB, _pC, _pubSignals)
}

// VerifyProof is a free data retrieval call binding the contract method 0xb9c6ea87.
//
// Solidity: function verifyProof(uint256[2] _pA, uint256[2][2] _pB, uint256[2] _pC, uint256[11] _pubSignals) view returns(bool)
func (_Verifiersigwrapper *VerifiersigwrapperCallerSession) VerifyProof(_pA [2]*big.Int, _pB [2][2]*big.Int, _pC [2]*big.Int, _pubSignals [11]*big.Int) (bool, error) {
	return _Verifiersigwrapper.Contract.VerifyProof(&_Verifiersigwrapper.CallOpts, _pA, _pB, _pC, _pubSignals)
}
