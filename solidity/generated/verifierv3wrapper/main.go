// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package verifierv3wrapper

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

// Verifierv3wrapperMetaData contains all meta data concerning the Verifierv3wrapper contract.
var Verifierv3wrapperMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256[2]\",\"name\":\"a\",\"type\":\"uint256[2]\"},{\"internalType\":\"uint256[2][2]\",\"name\":\"b\",\"type\":\"uint256[2][2]\"},{\"internalType\":\"uint256[2]\",\"name\":\"c\",\"type\":\"uint256[2]\"},{\"internalType\":\"uint256[]\",\"name\":\"input\",\"type\":\"uint256[]\"}],\"name\":\"verify\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"r\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[2]\",\"name\":\"_pA\",\"type\":\"uint256[2]\"},{\"internalType\":\"uint256[2][2]\",\"name\":\"_pB\",\"type\":\"uint256[2][2]\"},{\"internalType\":\"uint256[2]\",\"name\":\"_pC\",\"type\":\"uint256[2]\"},{\"internalType\":\"uint256[18]\",\"name\":\"_pubSignals\",\"type\":\"uint256[18]\"}],\"name\":\"verifyProof\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b5061121e806100206000396000f3fe608060405234801561001057600080fd5b50600436106100365760003560e01c80632612907c1461003b5780634483e7211461006b575b600080fd5b61005560048036038101906100509190610d4d565b61009b565b6040516100629190610df2565b60405180910390f35b61008560048036038101906100809190610e2f565b6101cd565b6040516100929190610df2565b60405180910390f35b60006100a5610c77565b601284849050146100eb576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016100e290610ef5565b60405180910390fd5b60005b60128110156101405784848281811061010a57610109610f15565b5b9050602002013582826012811061012457610123610f15565b5b602002018181525050808061013890610f7d565b9150506100ee565b503073ffffffffffffffffffffffffffffffffffffffff16634483e721888888856040518563ffffffff1660e01b81526004016101809493929190611148565b602060405180830381865afa15801561019d573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906101c191906111bb565b91505095945050505050565b6000610b53565b7f30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd478110610205576000805260206000f35b50565b600060405183815284602082015285604082015260408160608360076107d05a03fa91508161023b576000805260206000f35b825160408201526020830151606082015260408360808360066107d05a03fa91508161026b576000805260206000f35b505050505050565b600060808601600087017f130e6e3408fb554799c24818aa7d14c0d42b2d576ee1b247eba1417668bda1e681527f0330e7932823c2740fbc06b5418f0e84a2e48174921640adbaa2fcc354779794602082015261031660008801357f1c7105c33163afea7554457366353f780de5075ae676d832721ebdbbdb35203a7f0641f11e1cfd2d05eb33ecf162383f3753039d4a4d43338e18dc9675fb8d8de084610208565b61036660208801357f2f29fea7f771b615d4a81d70bd96a50d78b3656a798ca1d1dbaf25af2c4ab1187f2cf5174061cfc00099a8157f440e21e1f33f415b077e4fce83e2c3a11ac0c97a84610208565b6103b660408801357f02e3ffe92fe625446905c3aa68b27fa525d0ed299399f4f30b0fbe6d1e4f4c0c7f17455868052a6e32c18c98b5d09dba86c64b7c90f736f2f91da4f736a634a65a84610208565b61040660608801357f053507b0d7033d9e1c1a269597f3e76984b9e31e37e722f8dfe2560d10c1f9ed7f281bebe7040f63d5fb10f54730c5467acd4f97bbca82207c99a8b74045f33bfc84610208565b61045660808801357f1495a7ac31cb36fcceabc236816d29105b41c8406b4a116350954cca4dc238937f0f27228469fdb9bea786f89ae783d34f5cad734a3e4d7094167582f49e194da784610208565b6104a660a08801357f0b8e9894b9dd6fea88241188e69bcb21509a002d64cdf56f7f584719656bfd0d7f11b90da2a35cc382b3525b5dee431a1b350aefa5e1643ca05bd0d05b7650ad3384610208565b6104f660c08801357f0f20d8229d97353b88fe00e8e939e36d42d9d5189956f39a4b82ee62215afe357f05ca6f42a9113d06200df1c6d8123c04993934132ab3bde643342d53a838925b84610208565b61054660e08801357f2ab02ee3e12f8030a77d0cea91b17d5b45c465fb9582d571b2460d19e5dad4897f1ff2ca8bb15963f9e6f57752717980f06df6e302bfe0bf8527d761bad6d4bdfc84610208565b6105976101008801357f069f08156dfc53a0d6135810decea82d08c0a3958dc45f667ac120838c7d88b17f05e81e460408d590c58b639638988054e553ae8fc01b7a9a6fc867c67362452684610208565b6105e76101208801357f2d312368c7cc190476621f4614893268cecdb542f93b78586d0c6aa5d1eca79a7ec2d46071a306d9c2c65df8a5987c2378863fa7bef5c7301115c70ef1cc3dbe84610208565b6106386101408801357f173bedaaa4d0a3ed6cbba535816c234dcd8b05ea242c010e6e1f2a606e08f1b07f25b84e565b9a6db0816026a94ad9483589293ad29dd97016bb74f878878faa8284610208565b6106896101608801357f0259e97f9b3b208004dd8bc61780f5db9104269aab378323fc7c1f15447eaac57f25ed680b09ae126d6282a6d9d066449dd4c444e1ddc9a052bb46faf1a1a0fd8c84610208565b6106da6101808801357f15dcc45394c85a739c1af331b69430fcacc8fba3b276050d621ad610f41eddcd7f1af2bd38bec975ea81dab8289e50a4c4e74d7a08ae32837bc1581afab25261d484610208565b61072b6101a08801357f26e43f3e27e2663c7981553495580f7c7f9866044a1ebd57be1c4ae1a840a53b7f2c87ba0c4e889f5826f413404dbd0bf926baeaeba4c2044b7d7406806c1de62b84610208565b61077c6101c08801357f0c40573f623e539a8cfe464f919f69faf68e8128d5e4d2f9a1a36daef994f2b27f1823ad4608e901689122e453482372b3fbc345d83b2677f95e054fe184e0c31e84610208565b6107cd6101e08801357f0a1ea97ea44779c01ced2cda93a42fa1dd5ad05a6ccb5c7a01fd7d8a67a378317f2d541e2b79984555daf2c1d6aafdbdbaf36f1f3c566ccf636bcd5dc529c5b9f984610208565b61081e6102008801357f2c1cb133232f005e309cb1d17b91011c7de8c4d451fdd60aacf3fa4d6db5db577f1977c17308aadcba3f89ca513de1c17c9d9cfef2b2b44ec3a9e7a3361a58c13784610208565b61086f6102208801357f25b0e9a6ae6f1572ae31311ec300e9f410f7cabc174f30148f6cb804161e90917f02f7aa11c8147f14d48af23792a328fd909fa374e7ab74eaa8390b8e29414e7784610208565b833582527f30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd4760208501357f30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd4703066020830152843560408301526020850135606083015260408501356080830152606085013560a08301527f2d4d9aa7e302d9df41749d5507949d05dbea33fbb16c643b22f599a2be6df2e260c08301527f14bedd503c37ceb061d8ec60209fe345ce89830a19230301f076caff004d192660e08301527f0967032fcbf776d1afc985f88877f182d38480a653f2decaa9794cbc3bf3060c6101008301527f0e187847ad4c798374d0d6732bf501847dd68bc0e071241e0213bc7fc13db7ab6101208301527f304cfbd1e08a704a99f5e847d93f8c3caafddec46b7a0d379da69a4d112346a76101408301527f1739c1b1a457a8c7313123d24d2f9192f896b7c63eea05a9d57f06547ad0cec8610160830152600088015161018083015260206000018801516101a08301527f198e9393920d483a7260bfb731fb5d25f1aa493335a9e71297e485b7aef312c26101c08301527f1800deef121f1e76426a00665e5c4479674322d4f75edadd46debd5cd992f6ed6101e08301527f090689d0585ff075ec9e99ad690c3395bc4b313370b38ef355acdadcd122975b6102008301527f12c85ea5db8c6deb4aab71808dcb408fe3d1e7690c43d37b4ce6cc0166fa7daa610220830152853561024083015260208601356102608301527f0b445631ccd2a052bfe04c325509e21e105e7aa3987acef18908fa3813d9489a6102808301527f2fdbab7e880e8d1cf7f613774684303ae1d1f43841f6d5898f596e3847130b6e6102a08301527f21b11ccfd5877f7d03b20be7bb256d2f190c2e4a56b7c2024c3a295f1c53b7886102c08301527f0b6a60f96b8e060874a67d4fca4eccfa4ca1bd3cf66a0153d8955b1bc6f805036102e08301526020826103008460086107d05a03fa82518116935050505095945050505050565b6040516103808101604052610b6b60008401356101d4565b610b7860208401356101d4565b610b8560408401356101d4565b610b9260608401356101d4565b610b9f60808401356101d4565b610bac60a08401356101d4565b610bb960c08401356101d4565b610bc660e08401356101d4565b610bd46101008401356101d4565b610be26101208401356101d4565b610bf06101408401356101d4565b610bfe6101608401356101d4565b610c0c6101808401356101d4565b610c1a6101a08401356101d4565b610c286101c08401356101d4565b610c366101e08401356101d4565b610c446102008401356101d4565b610c526102208401356101d4565b610c606102408401356101d4565b610c6d818486888a610273565b8060005260206000f35b604051806102400160405280601290602082028036833780820191505090505090565b600080fd5b600080fd5b600080fd5b600081905082602060020282011115610cc557610cc4610ca4565b5b92915050565b600081905082604060020282011115610ce757610ce6610ca4565b5b92915050565b600080fd5b600080fd5b60008083601f840112610d0d57610d0c610ced565b5b8235905067ffffffffffffffff811115610d2a57610d29610cf2565b5b602083019150836020820283011115610d4657610d45610ca4565b5b9250929050565b60008060008060006101208688031215610d6a57610d69610c9a565b5b6000610d7888828901610ca9565b9550506040610d8988828901610ccb565b94505060c0610d9a88828901610ca9565b93505061010086013567ffffffffffffffff811115610dbc57610dbb610c9f565b5b610dc888828901610cf7565b92509250509295509295909350565b60008115159050919050565b610dec81610dd7565b82525050565b6000602082019050610e076000830184610de3565b92915050565b600081905082602060120282011115610e2957610e28610ca4565b5b92915050565b6000806000806103408587031215610e4a57610e49610c9a565b5b6000610e5887828801610ca9565b9450506040610e6987828801610ccb565b93505060c0610e7a87828801610ca9565b925050610100610e8c87828801610e0d565b91505092959194509250565b600082825260208201905092915050565b7f6578706563746564206172726179206c656e6774682069732031380000000000600082015250565b6000610edf601b83610e98565b9150610eea82610ea9565b602082019050919050565b60006020820190508181036000830152610f0e81610ed2565b9050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b6000819050919050565b6000610f8882610f73565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8203610fba57610fb9610f44565b5b600182019050919050565b82818337505050565b610fda60408383610fc5565b5050565b600060029050919050565b600081905092915050565b6000819050919050565b61100a60408383610fc5565b5050565b600061101a8383610ffe565b60408301905092915050565b600082905092915050565b6000604082019050919050565b61104781610fde565b6110518184610fe9565b925061105c82610ff4565b8060005b83811015611095576110728284611026565b61107c878261100e565b965061108783611031565b925050600181019050611060565b505050505050565b600060129050919050565b600081905092915050565b6000819050919050565b6110c681610f73565b82525050565b60006110d883836110bd565b60208301905092915050565b6000602082019050919050565b6110fa8161109d565b61110481846110a8565b925061110f826110b3565b8060005b8381101561114057815161112787826110cc565b9650611132836110e4565b925050600181019050611113565b505050505050565b60006103408201905061115e6000830187610fce565b61116b604083018661103e565b61117860c0830185610fce565b6111866101008301846110f1565b95945050505050565b61119881610dd7565b81146111a357600080fd5b50565b6000815190506111b58161118f565b92915050565b6000602082840312156111d1576111d0610c9a565b5b60006111df848285016111a6565b9150509291505056fea26469706673582212204efda72c9c50c9ac91d5ced45d8e3f97100146fea1e78790246c03e58d4e306464736f6c63430008100033",
}

// Verifierv3wrapperABI is the input ABI used to generate the binding from.
// Deprecated: Use Verifierv3wrapperMetaData.ABI instead.
var Verifierv3wrapperABI = Verifierv3wrapperMetaData.ABI

// Verifierv3wrapperBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use Verifierv3wrapperMetaData.Bin instead.
var Verifierv3wrapperBin = Verifierv3wrapperMetaData.Bin

// DeployVerifierv3wrapper deploys a new Ethereum contract, binding an instance of Verifierv3wrapper to it.
func DeployVerifierv3wrapper(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Verifierv3wrapper, error) {
	parsed, err := Verifierv3wrapperMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(Verifierv3wrapperBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Verifierv3wrapper{Verifierv3wrapperCaller: Verifierv3wrapperCaller{contract: contract}, Verifierv3wrapperTransactor: Verifierv3wrapperTransactor{contract: contract}, Verifierv3wrapperFilterer: Verifierv3wrapperFilterer{contract: contract}}, nil
}

// Verifierv3wrapper is an auto generated Go binding around an Ethereum contract.
type Verifierv3wrapper struct {
	Verifierv3wrapperCaller     // Read-only binding to the contract
	Verifierv3wrapperTransactor // Write-only binding to the contract
	Verifierv3wrapperFilterer   // Log filterer for contract events
}

// Verifierv3wrapperCaller is an auto generated read-only Go binding around an Ethereum contract.
type Verifierv3wrapperCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Verifierv3wrapperTransactor is an auto generated write-only Go binding around an Ethereum contract.
type Verifierv3wrapperTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Verifierv3wrapperFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type Verifierv3wrapperFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Verifierv3wrapperSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type Verifierv3wrapperSession struct {
	Contract     *Verifierv3wrapper // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// Verifierv3wrapperCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type Verifierv3wrapperCallerSession struct {
	Contract *Verifierv3wrapperCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// Verifierv3wrapperTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type Verifierv3wrapperTransactorSession struct {
	Contract     *Verifierv3wrapperTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// Verifierv3wrapperRaw is an auto generated low-level Go binding around an Ethereum contract.
type Verifierv3wrapperRaw struct {
	Contract *Verifierv3wrapper // Generic contract binding to access the raw methods on
}

// Verifierv3wrapperCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type Verifierv3wrapperCallerRaw struct {
	Contract *Verifierv3wrapperCaller // Generic read-only contract binding to access the raw methods on
}

// Verifierv3wrapperTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type Verifierv3wrapperTransactorRaw struct {
	Contract *Verifierv3wrapperTransactor // Generic write-only contract binding to access the raw methods on
}

// NewVerifierv3wrapper creates a new instance of Verifierv3wrapper, bound to a specific deployed contract.
func NewVerifierv3wrapper(address common.Address, backend bind.ContractBackend) (*Verifierv3wrapper, error) {
	contract, err := bindVerifierv3wrapper(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Verifierv3wrapper{Verifierv3wrapperCaller: Verifierv3wrapperCaller{contract: contract}, Verifierv3wrapperTransactor: Verifierv3wrapperTransactor{contract: contract}, Verifierv3wrapperFilterer: Verifierv3wrapperFilterer{contract: contract}}, nil
}

// NewVerifierv3wrapperCaller creates a new read-only instance of Verifierv3wrapper, bound to a specific deployed contract.
func NewVerifierv3wrapperCaller(address common.Address, caller bind.ContractCaller) (*Verifierv3wrapperCaller, error) {
	contract, err := bindVerifierv3wrapper(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &Verifierv3wrapperCaller{contract: contract}, nil
}

// NewVerifierv3wrapperTransactor creates a new write-only instance of Verifierv3wrapper, bound to a specific deployed contract.
func NewVerifierv3wrapperTransactor(address common.Address, transactor bind.ContractTransactor) (*Verifierv3wrapperTransactor, error) {
	contract, err := bindVerifierv3wrapper(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &Verifierv3wrapperTransactor{contract: contract}, nil
}

// NewVerifierv3wrapperFilterer creates a new log filterer instance of Verifierv3wrapper, bound to a specific deployed contract.
func NewVerifierv3wrapperFilterer(address common.Address, filterer bind.ContractFilterer) (*Verifierv3wrapperFilterer, error) {
	contract, err := bindVerifierv3wrapper(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &Verifierv3wrapperFilterer{contract: contract}, nil
}

// bindVerifierv3wrapper binds a generic wrapper to an already deployed contract.
func bindVerifierv3wrapper(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := Verifierv3wrapperMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Verifierv3wrapper *Verifierv3wrapperRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Verifierv3wrapper.Contract.Verifierv3wrapperCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Verifierv3wrapper *Verifierv3wrapperRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Verifierv3wrapper.Contract.Verifierv3wrapperTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Verifierv3wrapper *Verifierv3wrapperRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Verifierv3wrapper.Contract.Verifierv3wrapperTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Verifierv3wrapper *Verifierv3wrapperCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Verifierv3wrapper.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Verifierv3wrapper *Verifierv3wrapperTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Verifierv3wrapper.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Verifierv3wrapper *Verifierv3wrapperTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Verifierv3wrapper.Contract.contract.Transact(opts, method, params...)
}

// Verify is a free data retrieval call binding the contract method 0x2612907c.
//
// Solidity: function verify(uint256[2] a, uint256[2][2] b, uint256[2] c, uint256[] input) view returns(bool r)
func (_Verifierv3wrapper *Verifierv3wrapperCaller) Verify(opts *bind.CallOpts, a [2]*big.Int, b [2][2]*big.Int, c [2]*big.Int, input []*big.Int) (bool, error) {
	var out []interface{}
	err := _Verifierv3wrapper.contract.Call(opts, &out, "verify", a, b, c, input)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Verify is a free data retrieval call binding the contract method 0x2612907c.
//
// Solidity: function verify(uint256[2] a, uint256[2][2] b, uint256[2] c, uint256[] input) view returns(bool r)
func (_Verifierv3wrapper *Verifierv3wrapperSession) Verify(a [2]*big.Int, b [2][2]*big.Int, c [2]*big.Int, input []*big.Int) (bool, error) {
	return _Verifierv3wrapper.Contract.Verify(&_Verifierv3wrapper.CallOpts, a, b, c, input)
}

// Verify is a free data retrieval call binding the contract method 0x2612907c.
//
// Solidity: function verify(uint256[2] a, uint256[2][2] b, uint256[2] c, uint256[] input) view returns(bool r)
func (_Verifierv3wrapper *Verifierv3wrapperCallerSession) Verify(a [2]*big.Int, b [2][2]*big.Int, c [2]*big.Int, input []*big.Int) (bool, error) {
	return _Verifierv3wrapper.Contract.Verify(&_Verifierv3wrapper.CallOpts, a, b, c, input)
}

// VerifyProof is a free data retrieval call binding the contract method 0x4483e721.
//
// Solidity: function verifyProof(uint256[2] _pA, uint256[2][2] _pB, uint256[2] _pC, uint256[18] _pubSignals) view returns(bool)
func (_Verifierv3wrapper *Verifierv3wrapperCaller) VerifyProof(opts *bind.CallOpts, _pA [2]*big.Int, _pB [2][2]*big.Int, _pC [2]*big.Int, _pubSignals [18]*big.Int) (bool, error) {
	var out []interface{}
	err := _Verifierv3wrapper.contract.Call(opts, &out, "verifyProof", _pA, _pB, _pC, _pubSignals)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// VerifyProof is a free data retrieval call binding the contract method 0x4483e721.
//
// Solidity: function verifyProof(uint256[2] _pA, uint256[2][2] _pB, uint256[2] _pC, uint256[18] _pubSignals) view returns(bool)
func (_Verifierv3wrapper *Verifierv3wrapperSession) VerifyProof(_pA [2]*big.Int, _pB [2][2]*big.Int, _pC [2]*big.Int, _pubSignals [18]*big.Int) (bool, error) {
	return _Verifierv3wrapper.Contract.VerifyProof(&_Verifierv3wrapper.CallOpts, _pA, _pB, _pC, _pubSignals)
}

// VerifyProof is a free data retrieval call binding the contract method 0x4483e721.
//
// Solidity: function verifyProof(uint256[2] _pA, uint256[2][2] _pB, uint256[2] _pC, uint256[18] _pubSignals) view returns(bool)
func (_Verifierv3wrapper *Verifierv3wrapperCallerSession) VerifyProof(_pA [2]*big.Int, _pB [2][2]*big.Int, _pC [2]*big.Int, _pubSignals [18]*big.Int) (bool, error) {
	return _Verifierv3wrapper.Contract.VerifyProof(&_Verifierv3wrapper.CallOpts, _pA, _pB, _pC, _pubSignals)
}
