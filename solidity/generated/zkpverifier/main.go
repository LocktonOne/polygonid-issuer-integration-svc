// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package zkpverifier

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

// IZKPVerifierZKPRequest is an auto generated low-level Go binding around an user-defined struct.
type IZKPVerifierZKPRequest struct {
	Metadata  string
	Validator common.Address
	Data      []byte
}

// ZkpverifierMetaData contains all meta data concerning the Zkpverifier contract.
var ZkpverifierMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"REQUESTS_RETURN_LIMIT\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"requestId\",\"type\":\"uint64\"}],\"name\":\"getZKPRequest\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"metadata\",\"type\":\"string\"},{\"internalType\":\"contractICircuitValidator\",\"name\":\"validator\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"internalType\":\"structIZKPVerifier.ZKPRequest\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"startIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"length\",\"type\":\"uint256\"}],\"name\":\"getZKPRequests\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"metadata\",\"type\":\"string\"},{\"internalType\":\"contractICircuitValidator\",\"name\":\"validator\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"internalType\":\"structIZKPVerifier.ZKPRequest[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getZKPRequestsCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"name\":\"proofs\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"requestId\",\"type\":\"uint64\"}],\"name\":\"requestIdExists\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"requestId\",\"type\":\"uint64\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"metadata\",\"type\":\"string\"},{\"internalType\":\"contractICircuitValidator\",\"name\":\"validator\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"internalType\":\"structIZKPVerifier.ZKPRequest\",\"name\":\"request\",\"type\":\"tuple\"}],\"name\":\"setZKPRequest\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"requestId\",\"type\":\"uint64\"},{\"internalType\":\"uint256[]\",\"name\":\"inputs\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[2]\",\"name\":\"a\",\"type\":\"uint256[2]\"},{\"internalType\":\"uint256[2][2]\",\"name\":\"b\",\"type\":\"uint256[2][2]\"},{\"internalType\":\"uint256[2]\",\"name\":\"c\",\"type\":\"uint256[2]\"}],\"name\":\"submitZKPResponse\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60806040523480156200001157600080fd5b5062000032620000266200003860201b60201c565b6200004060201b60201c565b62000104565b600033905090565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050816000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a35050565b6126cd80620001146000396000f3fe608060405234801561001057600080fd5b50600436106100a95760003560e01c80639f5223e0116100715780639f5223e014610142578063ab7bcfb71461015e578063b45c0fdf1461018e578063b68967e2146101be578063c76d0845146101da578063f2fde38b1461020a576100a9565b80631905e7b1146100ae5780635f9e60d7146100cc5780636508e1b4146100fc578063715018a61461011a5780638da5cb5b14610124575b600080fd5b6100b6610226565b6040516100c39190611127565b60405180910390f35b6100e660048036038101906100e19190611178565b61022c565b6040516100f39190611435565b60405180910390f35b6101046104ee565b6040516101119190611127565b60405180910390f35b6101226104fc565b005b61012c610510565b6040516101399190611478565b60405180910390f35b61015c600480360381019061015791906114f7565b610539565b005b61017860048036038101906101739190611553565b6105d4565b604051610185919061159b565b60405180910390f35b6101a860048036038101906101a391906115e2565b61066b565b6040516101b5919061159b565b60405180910390f35b6101d860048036038101906101d391906116cb565b61069b565b005b6101f460048036038101906101ef9190611553565b610919565b60405161020191906117be565b60405180910390f35b610224600480360381019061021f91906117e0565b610b1d565b005b6103e881565b60606000806102456101f78054905086866103e8610ba0565b9150915060008282610257919061183c565b67ffffffffffffffff8111156102705761026f611870565b5b6040519080825280602002602001820160405280156102a957816020015b6102966110d7565b81526020019060019003908161028e5790505b50905060008390505b828110156104e1576101f660006101f783815481106102d4576102d361189f565b5b90600052602060002090600491828204019190066008029054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff16815260200190815260200160002060405180606001604052908160008201805461033b906118fd565b80601f0160208091040260200160405190810160405280929190818152602001828054610367906118fd565b80156103b45780601f10610389576101008083540402835291602001916103b4565b820191906000526020600020905b81548152906001019060200180831161039757829003601f168201915b505050505081526020016001820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001600282018054610423906118fd565b80601f016020809104026020016040519081016040528092919081815260200182805461044f906118fd565b801561049c5780601f106104715761010080835404028352916020019161049c565b820191906000526020600020905b81548152906001019060200180831161047f57829003601f168201915b5050505050815250508285836104b2919061183c565b815181106104c3576104c261189f565b5b602002602001018190525080806104d99061192e565b9150506102b2565b5080935050505092915050565b60006101f780549050905090565b610504610c97565b61050e6000610d15565b565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b610541610c97565b806101f660008467ffffffffffffffff1667ffffffffffffffff16815260200190815260200160002081816105769190611f58565b9050506101f78290806001815401808255809150506001900390600052602060002090600491828204019190066008029091909190916101000a81548167ffffffffffffffff021916908367ffffffffffffffff1602179055505050565b600080600090505b6101f780549050811015610660578267ffffffffffffffff166101f7828154811061060a5761060961189f565b5b90600052602060002090600491828204019190066008029054906101000a900467ffffffffffffffff1667ffffffffffffffff160361064d576001915050610666565b80806106589061192e565b9150506105dc565b50600090505b919050565b6101f56020528160005260406000206020528060005260406000206000915091509054906101000a900460ff1681565b600073ffffffffffffffffffffffffffffffffffffffff166101f660008867ffffffffffffffff1667ffffffffffffffff16815260200190815260200160002060010160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1603610755576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161074c90611fe9565b60405180910390fd5b6107ec86868680806020026020016040519081016040528093929190818152602001838360200280828437600081840152601f19601f820116905080830192505050505050506101f660008a67ffffffffffffffff1667ffffffffffffffff16815260200190815260200160002060010160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16610dd9565b6107fb86868686868633610dde565b5060016101f560003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008867ffffffffffffffff1667ffffffffffffffff16815260200190815260200160002060006101000a81548160ff02191690831515021790555061091186868680806020026020016040519081016040528093929190818152602001838360200280828437600081840152601f19601f820116905080830192505050505050506101f660008a67ffffffffffffffff1667ffffffffffffffff16815260200190815260200160002060010160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff166110ca565b505050505050565b6109216110d7565b61092a826105d4565b610969576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161096090612055565b60405180910390fd5b6101f660008367ffffffffffffffff1667ffffffffffffffff1681526020019081526020016000206040518060600160405290816000820180546109ac906118fd565b80601f01602080910402602001604051908101604052809291908181526020018280546109d8906118fd565b8015610a255780601f106109fa57610100808354040283529160200191610a25565b820191906000526020600020905b815481529060010190602001808311610a0857829003601f168201915b505050505081526020016001820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001600282018054610a94906118fd565b80601f0160208091040260200160405190810160405280929190818152602001828054610ac0906118fd565b8015610b0d5780601f10610ae257610100808354040283529160200191610b0d565b820191906000526020600020905b815481529060010190602001808311610af057829003601f168201915b5050505050815250509050919050565b610b25610c97565b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1603610b94576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610b8b906120e7565b60405180910390fd5b610b9d81610d15565b50565b60008060008411610be6576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610bdd90612153565b60405180910390fd5b82841115610c29576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610c20906121bf565b60405180910390fd5b858510610c6b576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610c629061222b565b60405180910390fd5b60008486610c79919061224b565b905086811115610c87578690505b8581925092505094509492505050565b610c9f6110cf565b73ffffffffffffffffffffffffffffffffffffffff16610cbd610510565b73ffffffffffffffffffffffffffffffffffffffff1614610d13576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610d0a906122cb565b60405180910390fd5b565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050816000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a35050565b505050565b6000806101f660008a67ffffffffffffffff1667ffffffffffffffff168152602001908152602001600020604051806060016040529081600082018054610e24906118fd565b80601f0160208091040260200160405190810160405280929190818152602001828054610e50906118fd565b8015610e9d5780601f10610e7257610100808354040283529160200191610e9d565b820191906000526020600020905b815481529060010190602001808311610e8057829003601f168201915b505050505081526020016001820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001600282018054610f0c906118fd565b80601f0160208091040260200160405190810160405280929190818152602001828054610f38906118fd565b8015610f855780601f10610f5a57610100808354040283529160200191610f85565b820191906000526020600020905b815481529060010190602001808311610f6857829003601f168201915b50505050508152505090506000635307e79f60e01b90506000818a8a8a8a8a8860400151604051602001610fbe9695949392919061247f565b60405160208183030381529060405286604051602001610fe0939291906125b5565b6040516020818303038152906040529050600080846020015173ffffffffffffffffffffffffffffffffffffffff168360405161101d91906125ee565b6000604051808303816000865af19150503d806000811461105a576040519150601f19603f3d011682016040523d82523d6000602084013e61105f565b606091505b5091509150816110b75760008151111561107c5780518082602001fd5b6040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016110ae90612677565b60405180910390fd5b8195505050505050979650505050505050565b505050565b600033905090565b604051806060016040528060608152602001600073ffffffffffffffffffffffffffffffffffffffff168152602001606081525090565b6000819050919050565b6111218161110e565b82525050565b600060208201905061113c6000830184611118565b92915050565b600080fd5b600080fd5b6111558161110e565b811461116057600080fd5b50565b6000813590506111728161114c565b92915050565b6000806040838503121561118f5761118e611142565b5b600061119d85828601611163565b92505060206111ae85828601611163565b9150509250929050565b600081519050919050565b600082825260208201905092915050565b6000819050602082019050919050565b600081519050919050565b600082825260208201905092915050565b60005b8381101561121e578082015181840152602081019050611203565b60008484015250505050565b6000601f19601f8301169050919050565b6000611246826111e4565b61125081856111ef565b9350611260818560208601611200565b6112698161122a565b840191505092915050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b60006112b96112b46112af84611274565b611294565b611274565b9050919050565b60006112cb8261129e565b9050919050565b60006112dd826112c0565b9050919050565b6112ed816112d2565b82525050565b600081519050919050565b600082825260208201905092915050565b600061131a826112f3565b61132481856112fe565b9350611334818560208601611200565b61133d8161122a565b840191505092915050565b60006060830160008301518482036000860152611365828261123b565b915050602083015161137a60208601826112e4565b5060408301518482036040860152611392828261130f565b9150508091505092915050565b60006113ab8383611348565b905092915050565b6000602082019050919050565b60006113cb826111b8565b6113d581856111c3565b9350836020820285016113e7856111d4565b8060005b858110156114235784840389528151611404858261139f565b945061140f836113b3565b925060208a019950506001810190506113eb565b50829750879550505050505092915050565b6000602082019050818103600083015261144f81846113c0565b905092915050565b600061146282611274565b9050919050565b61147281611457565b82525050565b600060208201905061148d6000830184611469565b92915050565b600067ffffffffffffffff82169050919050565b6114b081611493565b81146114bb57600080fd5b50565b6000813590506114cd816114a7565b92915050565b600080fd5b6000606082840312156114ee576114ed6114d3565b5b81905092915050565b6000806040838503121561150e5761150d611142565b5b600061151c858286016114be565b925050602083013567ffffffffffffffff81111561153d5761153c611147565b5b611549858286016114d8565b9150509250929050565b60006020828403121561156957611568611142565b5b6000611577848285016114be565b91505092915050565b60008115159050919050565b61159581611580565b82525050565b60006020820190506115b0600083018461158c565b92915050565b6115bf81611457565b81146115ca57600080fd5b50565b6000813590506115dc816115b6565b92915050565b600080604083850312156115f9576115f8611142565b5b6000611607858286016115cd565b9250506020611618858286016114be565b9150509250929050565b600080fd5b600080fd5b600080fd5b60008083601f84011261164757611646611622565b5b8235905067ffffffffffffffff81111561166457611663611627565b5b6020830191508360208202830111156116805761167f61162c565b5b9250929050565b6000819050826020600202820111156116a3576116a261162c565b5b92915050565b6000819050826040600202820111156116c5576116c461162c565b5b92915050565b60008060008060008061014087890312156116e9576116e8611142565b5b60006116f789828a016114be565b965050602087013567ffffffffffffffff81111561171857611717611147565b5b61172489828a01611631565b9550955050604061173789828a01611687565b935050608061174889828a016116a9565b92505061010061175a89828a01611687565b9150509295509295509295565b60006060830160008301518482036000860152611784828261123b565b915050602083015161179960208601826112e4565b50604083015184820360408601526117b1828261130f565b9150508091505092915050565b600060208201905081810360008301526117d88184611767565b905092915050565b6000602082840312156117f6576117f5611142565b5b6000611804848285016115cd565b91505092915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b60006118478261110e565b91506118528361110e565b925082820390508181111561186a5761186961180d565b5b92915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b6000600282049050600182168061191557607f821691505b602082108103611928576119276118ce565b5b50919050565b60006119398261110e565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff820361196b5761196a61180d565b5b600182019050919050565b600080fd5b600080fd5b600080fd5b600080833560016020038436030381126119a2576119a1611976565b5b80840192508235915067ffffffffffffffff8211156119c4576119c361197b565b5b6020830192506001820236038313156119e0576119df611980565b5b509250929050565b600082905092915050565b60008190508160005260206000209050919050565b60006020601f8301049050919050565b600082821b905092915050565b600060088302611a557fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82611a18565b611a5f8683611a18565b95508019841693508086168417925050509392505050565b6000611a92611a8d611a888461110e565b611294565b61110e565b9050919050565b6000819050919050565b611aac83611a77565b611ac0611ab882611a99565b848454611a25565b825550505050565b600090565b611ad5611ac8565b611ae0818484611aa3565b505050565b5b81811015611b0457611af9600082611acd565b600181019050611ae6565b5050565b601f821115611b4957611b1a816119f3565b611b2384611a08565b81016020851015611b32578190505b611b46611b3e85611a08565b830182611ae5565b50505b505050565b600082821c905092915050565b6000611b6c60001984600802611b4e565b1980831691505092915050565b6000611b858383611b5b565b9150826002028217905092915050565b611b9f83836119e8565b67ffffffffffffffff811115611bb857611bb7611870565b5b611bc282546118fd565b611bcd828285611b08565b6000601f831160018114611bfc5760008415611bea578287013590505b611bf48582611b79565b865550611c5c565b601f198416611c0a866119f3565b60005b82811015611c3257848901358255600182019150602085019450602081019050611c0d565b86831015611c4f5784890135611c4b601f891682611b5b565b8355505b6001600288020188555050505b50505050505050565b611c70838383611b95565b505050565b6000611c8082611457565b9050919050565b611c9081611c75565b8114611c9b57600080fd5b50565b60008135611cab81611c87565b80915050919050565b60008160001b9050919050565b600073ffffffffffffffffffffffffffffffffffffffff611ce184611cb4565b9350801983169250808416831791505092915050565b6000611d028261129e565b9050919050565b6000611d1482611cf7565b9050919050565b6000819050919050565b611d2e82611d09565b611d41611d3a82611d1b565b8354611cc1565b8255505050565b60008083356001602003843603038112611d6557611d64611976565b5b80840192508235915067ffffffffffffffff821115611d8757611d8661197b565b5b602083019250600182023603831315611da357611da2611980565b5b509250929050565b600082905092915050565b60008190508160005260206000209050919050565b601f821115611e0c57611ddd81611db6565b611de684611a08565b81016020851015611df5578190505b611e09611e0185611a08565b830182611ae5565b50505b505050565b611e1b8383611dab565b67ffffffffffffffff811115611e3457611e33611870565b5b611e3e82546118fd565b611e49828285611dcb565b6000601f831160018114611e785760008415611e66578287013590505b611e708582611b79565b865550611ed8565b601f198416611e8686611db6565b60005b82811015611eae57848901358255600182019150602085019450602081019050611e89565b86831015611ecb5784890135611ec7601f891682611b5b565b8355505b6001600288020188555050505b50505050505050565b611eec838383611e11565b505050565b6000810160008301611f038185611985565b611f0e818386611c65565b50505050600181016020830180611f2481611c9e565b9050611f308184611d25565b5050506002810160408301611f458185611d48565b611f50818386611ee1565b505050505050565b611f628282611ef1565b5050565b600082825260208201905092915050565b7f76616c696461746f72206973206e6f742073657420666f72207468697320726560008201527f7175657374206964000000000000000000000000000000000000000000000000602082015250565b6000611fd3602883611f66565b9150611fde82611f77565b604082019050919050565b6000602082019050818103600083015261200281611fc6565b9050919050565b7f7265717565737420696420646f65736e27742065786973740000000000000000600082015250565b600061203f601883611f66565b915061204a82612009565b602082019050919050565b6000602082019050818103600083015261206e81612032565b9050919050565b7f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160008201527f6464726573730000000000000000000000000000000000000000000000000000602082015250565b60006120d1602683611f66565b91506120dc82612075565b604082019050919050565b60006020820190508181036000830152612100816120c4565b9050919050565b7f4c656e6774682073686f756c642062652067726561746572207468616e203000600082015250565b600061213d601f83611f66565b915061214882612107565b602082019050919050565b6000602082019050818103600083015261216c81612130565b9050919050565b7f4c656e677468206c696d69742065786365656465640000000000000000000000600082015250565b60006121a9601583611f66565b91506121b482612173565b602082019050919050565b600060208201905081810360008301526121d88161219c565b9050919050565b7f537461727420696e646578206f7574206f6620626f756e647300000000000000600082015250565b6000612215601983611f66565b9150612220826121df565b602082019050919050565b6000602082019050818103600083015261224481612208565b9050919050565b60006122568261110e565b91506122618361110e565b92508282019050808211156122795761227861180d565b5b92915050565b7f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572600082015250565b60006122b5602083611f66565b91506122c08261227f565b602082019050919050565b600060208201905081810360008301526122e4816122a8565b9050919050565b600082825260208201905092915050565b600080fd5b82818337505050565b600061231683856122eb565b93507f07ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff831115612349576123486122fc565b5b60208302925061235a838584612301565b82840190509392505050565b61237260408383612301565b5050565b600060029050919050565b600081905092915050565b6000819050919050565b6123a260408383612301565b5050565b60006123b28383612396565b60408301905092915050565b600082905092915050565b6000604082019050919050565b6123df81612376565b6123e98184612381565b92506123f48261238c565b8060005b8381101561242d5761240a82846123be565b61241487826123a6565b965061241f836123c9565b9250506001810190506123f8565b505050505050565b600082825260208201905092915050565b6000612451826112f3565b61245b8185612435565b935061246b818560208601611200565b6124748161122a565b840191505092915050565b600061014082019050818103600083015261249b81888a61230a565b90506124aa6020830187612366565b6124b760608301866123d6565b6124c460e0830185612366565b8181036101208301526124d78184612446565b9050979650505050505050565b60007fffffffff0000000000000000000000000000000000000000000000000000000082169050919050565b6000819050919050565b61252b612526826124e4565b612510565b82525050565b600081905092915050565b6000612547826112f3565b6125518185612531565b9350612561818560208601611200565b80840191505092915050565b60008160601b9050919050565b60006125858261256d565b9050919050565b60006125978261257a565b9050919050565b6125af6125aa82611457565b61258c565b82525050565b60006125c1828661251a565b6004820191506125d1828561253c565b91506125dd828461259e565b601482019150819050949350505050565b60006125fa828461253c565b915081905092915050565b7f4661696c656420746f207665726966792070726f6f6620776974686f7574207260008201527f657665727420726561736f6e0000000000000000000000000000000000000000602082015250565b6000612661602c83611f66565b915061266c82612605565b604082019050919050565b6000602082019050818103600083015261269081612654565b905091905056fea264697066735822122027308e0f87f42c2e546bb311e89f9a8473aa64585693852f0872cfd8b9dda96d64736f6c63430008100033",
}

// ZkpverifierABI is the input ABI used to generate the binding from.
// Deprecated: Use ZkpverifierMetaData.ABI instead.
var ZkpverifierABI = ZkpverifierMetaData.ABI

// ZkpverifierBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ZkpverifierMetaData.Bin instead.
var ZkpverifierBin = ZkpverifierMetaData.Bin

// DeployZkpverifier deploys a new Ethereum contract, binding an instance of Zkpverifier to it.
func DeployZkpverifier(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Zkpverifier, error) {
	parsed, err := ZkpverifierMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ZkpverifierBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Zkpverifier{ZkpverifierCaller: ZkpverifierCaller{contract: contract}, ZkpverifierTransactor: ZkpverifierTransactor{contract: contract}, ZkpverifierFilterer: ZkpverifierFilterer{contract: contract}}, nil
}

// Zkpverifier is an auto generated Go binding around an Ethereum contract.
type Zkpverifier struct {
	ZkpverifierCaller     // Read-only binding to the contract
	ZkpverifierTransactor // Write-only binding to the contract
	ZkpverifierFilterer   // Log filterer for contract events
}

// ZkpverifierCaller is an auto generated read-only Go binding around an Ethereum contract.
type ZkpverifierCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ZkpverifierTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ZkpverifierTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ZkpverifierFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ZkpverifierFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ZkpverifierSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ZkpverifierSession struct {
	Contract     *Zkpverifier      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ZkpverifierCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ZkpverifierCallerSession struct {
	Contract *ZkpverifierCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// ZkpverifierTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ZkpverifierTransactorSession struct {
	Contract     *ZkpverifierTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// ZkpverifierRaw is an auto generated low-level Go binding around an Ethereum contract.
type ZkpverifierRaw struct {
	Contract *Zkpverifier // Generic contract binding to access the raw methods on
}

// ZkpverifierCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ZkpverifierCallerRaw struct {
	Contract *ZkpverifierCaller // Generic read-only contract binding to access the raw methods on
}

// ZkpverifierTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ZkpverifierTransactorRaw struct {
	Contract *ZkpverifierTransactor // Generic write-only contract binding to access the raw methods on
}

// NewZkpverifier creates a new instance of Zkpverifier, bound to a specific deployed contract.
func NewZkpverifier(address common.Address, backend bind.ContractBackend) (*Zkpverifier, error) {
	contract, err := bindZkpverifier(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Zkpverifier{ZkpverifierCaller: ZkpverifierCaller{contract: contract}, ZkpverifierTransactor: ZkpverifierTransactor{contract: contract}, ZkpverifierFilterer: ZkpverifierFilterer{contract: contract}}, nil
}

// NewZkpverifierCaller creates a new read-only instance of Zkpverifier, bound to a specific deployed contract.
func NewZkpverifierCaller(address common.Address, caller bind.ContractCaller) (*ZkpverifierCaller, error) {
	contract, err := bindZkpverifier(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ZkpverifierCaller{contract: contract}, nil
}

// NewZkpverifierTransactor creates a new write-only instance of Zkpverifier, bound to a specific deployed contract.
func NewZkpverifierTransactor(address common.Address, transactor bind.ContractTransactor) (*ZkpverifierTransactor, error) {
	contract, err := bindZkpverifier(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ZkpverifierTransactor{contract: contract}, nil
}

// NewZkpverifierFilterer creates a new log filterer instance of Zkpverifier, bound to a specific deployed contract.
func NewZkpverifierFilterer(address common.Address, filterer bind.ContractFilterer) (*ZkpverifierFilterer, error) {
	contract, err := bindZkpverifier(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ZkpverifierFilterer{contract: contract}, nil
}

// bindZkpverifier binds a generic wrapper to an already deployed contract.
func bindZkpverifier(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ZkpverifierMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Zkpverifier *ZkpverifierRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Zkpverifier.Contract.ZkpverifierCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Zkpverifier *ZkpverifierRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Zkpverifier.Contract.ZkpverifierTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Zkpverifier *ZkpverifierRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Zkpverifier.Contract.ZkpverifierTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Zkpverifier *ZkpverifierCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Zkpverifier.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Zkpverifier *ZkpverifierTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Zkpverifier.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Zkpverifier *ZkpverifierTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Zkpverifier.Contract.contract.Transact(opts, method, params...)
}

// REQUESTSRETURNLIMIT is a free data retrieval call binding the contract method 0x1905e7b1.
//
// Solidity: function REQUESTS_RETURN_LIMIT() view returns(uint256)
func (_Zkpverifier *ZkpverifierCaller) REQUESTSRETURNLIMIT(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Zkpverifier.contract.Call(opts, &out, "REQUESTS_RETURN_LIMIT")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// REQUESTSRETURNLIMIT is a free data retrieval call binding the contract method 0x1905e7b1.
//
// Solidity: function REQUESTS_RETURN_LIMIT() view returns(uint256)
func (_Zkpverifier *ZkpverifierSession) REQUESTSRETURNLIMIT() (*big.Int, error) {
	return _Zkpverifier.Contract.REQUESTSRETURNLIMIT(&_Zkpverifier.CallOpts)
}

// REQUESTSRETURNLIMIT is a free data retrieval call binding the contract method 0x1905e7b1.
//
// Solidity: function REQUESTS_RETURN_LIMIT() view returns(uint256)
func (_Zkpverifier *ZkpverifierCallerSession) REQUESTSRETURNLIMIT() (*big.Int, error) {
	return _Zkpverifier.Contract.REQUESTSRETURNLIMIT(&_Zkpverifier.CallOpts)
}

// GetZKPRequest is a free data retrieval call binding the contract method 0xc76d0845.
//
// Solidity: function getZKPRequest(uint64 requestId) view returns((string,address,bytes))
func (_Zkpverifier *ZkpverifierCaller) GetZKPRequest(opts *bind.CallOpts, requestId uint64) (IZKPVerifierZKPRequest, error) {
	var out []interface{}
	err := _Zkpverifier.contract.Call(opts, &out, "getZKPRequest", requestId)

	if err != nil {
		return *new(IZKPVerifierZKPRequest), err
	}

	out0 := *abi.ConvertType(out[0], new(IZKPVerifierZKPRequest)).(*IZKPVerifierZKPRequest)

	return out0, err

}

// GetZKPRequest is a free data retrieval call binding the contract method 0xc76d0845.
//
// Solidity: function getZKPRequest(uint64 requestId) view returns((string,address,bytes))
func (_Zkpverifier *ZkpverifierSession) GetZKPRequest(requestId uint64) (IZKPVerifierZKPRequest, error) {
	return _Zkpverifier.Contract.GetZKPRequest(&_Zkpverifier.CallOpts, requestId)
}

// GetZKPRequest is a free data retrieval call binding the contract method 0xc76d0845.
//
// Solidity: function getZKPRequest(uint64 requestId) view returns((string,address,bytes))
func (_Zkpverifier *ZkpverifierCallerSession) GetZKPRequest(requestId uint64) (IZKPVerifierZKPRequest, error) {
	return _Zkpverifier.Contract.GetZKPRequest(&_Zkpverifier.CallOpts, requestId)
}

// GetZKPRequests is a free data retrieval call binding the contract method 0x5f9e60d7.
//
// Solidity: function getZKPRequests(uint256 startIndex, uint256 length) view returns((string,address,bytes)[])
func (_Zkpverifier *ZkpverifierCaller) GetZKPRequests(opts *bind.CallOpts, startIndex *big.Int, length *big.Int) ([]IZKPVerifierZKPRequest, error) {
	var out []interface{}
	err := _Zkpverifier.contract.Call(opts, &out, "getZKPRequests", startIndex, length)

	if err != nil {
		return *new([]IZKPVerifierZKPRequest), err
	}

	out0 := *abi.ConvertType(out[0], new([]IZKPVerifierZKPRequest)).(*[]IZKPVerifierZKPRequest)

	return out0, err

}

// GetZKPRequests is a free data retrieval call binding the contract method 0x5f9e60d7.
//
// Solidity: function getZKPRequests(uint256 startIndex, uint256 length) view returns((string,address,bytes)[])
func (_Zkpverifier *ZkpverifierSession) GetZKPRequests(startIndex *big.Int, length *big.Int) ([]IZKPVerifierZKPRequest, error) {
	return _Zkpverifier.Contract.GetZKPRequests(&_Zkpverifier.CallOpts, startIndex, length)
}

// GetZKPRequests is a free data retrieval call binding the contract method 0x5f9e60d7.
//
// Solidity: function getZKPRequests(uint256 startIndex, uint256 length) view returns((string,address,bytes)[])
func (_Zkpverifier *ZkpverifierCallerSession) GetZKPRequests(startIndex *big.Int, length *big.Int) ([]IZKPVerifierZKPRequest, error) {
	return _Zkpverifier.Contract.GetZKPRequests(&_Zkpverifier.CallOpts, startIndex, length)
}

// GetZKPRequestsCount is a free data retrieval call binding the contract method 0x6508e1b4.
//
// Solidity: function getZKPRequestsCount() view returns(uint256)
func (_Zkpverifier *ZkpverifierCaller) GetZKPRequestsCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Zkpverifier.contract.Call(opts, &out, "getZKPRequestsCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetZKPRequestsCount is a free data retrieval call binding the contract method 0x6508e1b4.
//
// Solidity: function getZKPRequestsCount() view returns(uint256)
func (_Zkpverifier *ZkpverifierSession) GetZKPRequestsCount() (*big.Int, error) {
	return _Zkpverifier.Contract.GetZKPRequestsCount(&_Zkpverifier.CallOpts)
}

// GetZKPRequestsCount is a free data retrieval call binding the contract method 0x6508e1b4.
//
// Solidity: function getZKPRequestsCount() view returns(uint256)
func (_Zkpverifier *ZkpverifierCallerSession) GetZKPRequestsCount() (*big.Int, error) {
	return _Zkpverifier.Contract.GetZKPRequestsCount(&_Zkpverifier.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Zkpverifier *ZkpverifierCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Zkpverifier.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Zkpverifier *ZkpverifierSession) Owner() (common.Address, error) {
	return _Zkpverifier.Contract.Owner(&_Zkpverifier.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Zkpverifier *ZkpverifierCallerSession) Owner() (common.Address, error) {
	return _Zkpverifier.Contract.Owner(&_Zkpverifier.CallOpts)
}

// Proofs is a free data retrieval call binding the contract method 0xb45c0fdf.
//
// Solidity: function proofs(address , uint64 ) view returns(bool)
func (_Zkpverifier *ZkpverifierCaller) Proofs(opts *bind.CallOpts, arg0 common.Address, arg1 uint64) (bool, error) {
	var out []interface{}
	err := _Zkpverifier.contract.Call(opts, &out, "proofs", arg0, arg1)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Proofs is a free data retrieval call binding the contract method 0xb45c0fdf.
//
// Solidity: function proofs(address , uint64 ) view returns(bool)
func (_Zkpverifier *ZkpverifierSession) Proofs(arg0 common.Address, arg1 uint64) (bool, error) {
	return _Zkpverifier.Contract.Proofs(&_Zkpverifier.CallOpts, arg0, arg1)
}

// Proofs is a free data retrieval call binding the contract method 0xb45c0fdf.
//
// Solidity: function proofs(address , uint64 ) view returns(bool)
func (_Zkpverifier *ZkpverifierCallerSession) Proofs(arg0 common.Address, arg1 uint64) (bool, error) {
	return _Zkpverifier.Contract.Proofs(&_Zkpverifier.CallOpts, arg0, arg1)
}

// RequestIdExists is a free data retrieval call binding the contract method 0xab7bcfb7.
//
// Solidity: function requestIdExists(uint64 requestId) view returns(bool)
func (_Zkpverifier *ZkpverifierCaller) RequestIdExists(opts *bind.CallOpts, requestId uint64) (bool, error) {
	var out []interface{}
	err := _Zkpverifier.contract.Call(opts, &out, "requestIdExists", requestId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// RequestIdExists is a free data retrieval call binding the contract method 0xab7bcfb7.
//
// Solidity: function requestIdExists(uint64 requestId) view returns(bool)
func (_Zkpverifier *ZkpverifierSession) RequestIdExists(requestId uint64) (bool, error) {
	return _Zkpverifier.Contract.RequestIdExists(&_Zkpverifier.CallOpts, requestId)
}

// RequestIdExists is a free data retrieval call binding the contract method 0xab7bcfb7.
//
// Solidity: function requestIdExists(uint64 requestId) view returns(bool)
func (_Zkpverifier *ZkpverifierCallerSession) RequestIdExists(requestId uint64) (bool, error) {
	return _Zkpverifier.Contract.RequestIdExists(&_Zkpverifier.CallOpts, requestId)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Zkpverifier *ZkpverifierTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Zkpverifier.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Zkpverifier *ZkpverifierSession) RenounceOwnership() (*types.Transaction, error) {
	return _Zkpverifier.Contract.RenounceOwnership(&_Zkpverifier.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Zkpverifier *ZkpverifierTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Zkpverifier.Contract.RenounceOwnership(&_Zkpverifier.TransactOpts)
}

// SetZKPRequest is a paid mutator transaction binding the contract method 0x9f5223e0.
//
// Solidity: function setZKPRequest(uint64 requestId, (string,address,bytes) request) returns()
func (_Zkpverifier *ZkpverifierTransactor) SetZKPRequest(opts *bind.TransactOpts, requestId uint64, request IZKPVerifierZKPRequest) (*types.Transaction, error) {
	return _Zkpverifier.contract.Transact(opts, "setZKPRequest", requestId, request)
}

// SetZKPRequest is a paid mutator transaction binding the contract method 0x9f5223e0.
//
// Solidity: function setZKPRequest(uint64 requestId, (string,address,bytes) request) returns()
func (_Zkpverifier *ZkpverifierSession) SetZKPRequest(requestId uint64, request IZKPVerifierZKPRequest) (*types.Transaction, error) {
	return _Zkpverifier.Contract.SetZKPRequest(&_Zkpverifier.TransactOpts, requestId, request)
}

// SetZKPRequest is a paid mutator transaction binding the contract method 0x9f5223e0.
//
// Solidity: function setZKPRequest(uint64 requestId, (string,address,bytes) request) returns()
func (_Zkpverifier *ZkpverifierTransactorSession) SetZKPRequest(requestId uint64, request IZKPVerifierZKPRequest) (*types.Transaction, error) {
	return _Zkpverifier.Contract.SetZKPRequest(&_Zkpverifier.TransactOpts, requestId, request)
}

// SubmitZKPResponse is a paid mutator transaction binding the contract method 0xb68967e2.
//
// Solidity: function submitZKPResponse(uint64 requestId, uint256[] inputs, uint256[2] a, uint256[2][2] b, uint256[2] c) returns()
func (_Zkpverifier *ZkpverifierTransactor) SubmitZKPResponse(opts *bind.TransactOpts, requestId uint64, inputs []*big.Int, a [2]*big.Int, b [2][2]*big.Int, c [2]*big.Int) (*types.Transaction, error) {
	return _Zkpverifier.contract.Transact(opts, "submitZKPResponse", requestId, inputs, a, b, c)
}

// SubmitZKPResponse is a paid mutator transaction binding the contract method 0xb68967e2.
//
// Solidity: function submitZKPResponse(uint64 requestId, uint256[] inputs, uint256[2] a, uint256[2][2] b, uint256[2] c) returns()
func (_Zkpverifier *ZkpverifierSession) SubmitZKPResponse(requestId uint64, inputs []*big.Int, a [2]*big.Int, b [2][2]*big.Int, c [2]*big.Int) (*types.Transaction, error) {
	return _Zkpverifier.Contract.SubmitZKPResponse(&_Zkpverifier.TransactOpts, requestId, inputs, a, b, c)
}

// SubmitZKPResponse is a paid mutator transaction binding the contract method 0xb68967e2.
//
// Solidity: function submitZKPResponse(uint64 requestId, uint256[] inputs, uint256[2] a, uint256[2][2] b, uint256[2] c) returns()
func (_Zkpverifier *ZkpverifierTransactorSession) SubmitZKPResponse(requestId uint64, inputs []*big.Int, a [2]*big.Int, b [2][2]*big.Int, c [2]*big.Int) (*types.Transaction, error) {
	return _Zkpverifier.Contract.SubmitZKPResponse(&_Zkpverifier.TransactOpts, requestId, inputs, a, b, c)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Zkpverifier *ZkpverifierTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Zkpverifier.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Zkpverifier *ZkpverifierSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Zkpverifier.Contract.TransferOwnership(&_Zkpverifier.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Zkpverifier *ZkpverifierTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Zkpverifier.Contract.TransferOwnership(&_Zkpverifier.TransactOpts, newOwner)
}

// ZkpverifierOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Zkpverifier contract.
type ZkpverifierOwnershipTransferredIterator struct {
	Event *ZkpverifierOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *ZkpverifierOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZkpverifierOwnershipTransferred)
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
		it.Event = new(ZkpverifierOwnershipTransferred)
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
func (it *ZkpverifierOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZkpverifierOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZkpverifierOwnershipTransferred represents a OwnershipTransferred event raised by the Zkpverifier contract.
type ZkpverifierOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Zkpverifier *ZkpverifierFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*ZkpverifierOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Zkpverifier.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &ZkpverifierOwnershipTransferredIterator{contract: _Zkpverifier.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Zkpverifier *ZkpverifierFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ZkpverifierOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Zkpverifier.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZkpverifierOwnershipTransferred)
				if err := _Zkpverifier.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Zkpverifier *ZkpverifierFilterer) ParseOwnershipTransferred(log types.Log) (*ZkpverifierOwnershipTransferred, error) {
	event := new(ZkpverifierOwnershipTransferred)
	if err := _Zkpverifier.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
