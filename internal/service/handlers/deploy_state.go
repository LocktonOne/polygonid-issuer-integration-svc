package handlers

import (
	"context"
	"encoding/hex"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/pkg/errors"
	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
	"gitlab.com/tokene/polygonid-issuer-integration/internal/service/helpers"
	"gitlab.com/tokene/polygonid-issuer-integration/internal/service/helpers/converter"
	"gitlab.com/tokene/polygonid-issuer-integration/internal/service/requests"
	"gitlab.com/tokene/polygonid-issuer-integration/solidity/generated/smtlib"
	"gitlab.com/tokene/polygonid-issuer-integration/solidity/generated/state"
	"gitlab.com/tokene/polygonid-issuer-integration/solidity/generated/statelib"
	"gitlab.com/tokene/polygonid-issuer-integration/solidity/generated/verifierstatetransition"
	"math/big"
	"net/http"
	"regexp"
	"strings"
)

func DeployState(w http.ResponseWriter, r *http.Request) {
	request, err := requests.NewDeployStateRequest(r)
	if err != nil {
		helpers.Log(r).WithError(err).Info("wrong request")
		ape.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	poseidonContracts := helpers.PoseidonContractsConfig(r)
	if poseidonContracts.El1 == nil ||
		poseidonContracts.El2 == nil ||
		poseidonContracts.El3 == nil {
		helpers.Log(r).Info("poseidon contract addresses are not specified")
		ape.RenderErr(w, problems.InternalError())
		return
	}

	auth, client, err := helpers.GetEthClient(r)
	var verifierStateTransitionAddress common.Address
	if request.Data.Attributes.VerifierStateTransition == nil {
		address, tx, _, err := verifierstatetransition.DeployVerifierstatetransition(auth, client)
		if err != nil {
			helpers.Log(r).WithError(err).Info("failed to deploy contract")
			ape.RenderErr(w, problems.InternalError())
			return
		}
		auth.Nonce = auth.Nonce.Add(auth.Nonce, big.NewInt(1))
		verifierStateTransitionAddress = address
		if _, err = bind.WaitMined(context.Background(), client, tx); err != nil {
			helpers.Log(r).WithError(err).Info(errors.Wrap(err, "failed to wait deploy state contract"))
			ape.RenderErr(w, problems.InternalError())
			return
		}
	} else {
		verifierStateTransitionAddress = common.HexToAddress(*request.Data.Attributes.VerifierStateTransition)
	}
	fmt.Println("verifierStateTransitionAddress", verifierStateTransitionAddress)

	var smtLibAddress common.Address
	if request.Data.Attributes.SmtLibAddress == nil {
		address, tx, _, err := DeployContractWithLinks(auth, client, smtlib.SmtlibABI, smtlib.SmtlibBin, map[string]common.Address{
			"PoseidonUnit2L": common.HexToAddress(*poseidonContracts.El2),
			"PoseidonUnit3L": common.HexToAddress(*poseidonContracts.El3),
		})
		//address, tx, _, err := smtlib.DeploySmtlib(auth, client)
		if err != nil {
			helpers.Log(r).WithError(err).Info("failed to smt library contract")
			ape.RenderErr(w, problems.InternalError())
			return
		}
		smtLibAddress = address
		if _, err = bind.WaitMined(context.Background(), client, tx); err != nil {
			helpers.Log(r).WithError(err).Info(errors.Wrap(err, "failed to wait deploy state contract"))
			ape.RenderErr(w, problems.InternalError())
			return
		}
		auth.Nonce = auth.Nonce.Add(auth.Nonce, big.NewInt(1))
	} else {
		smtLibAddress = common.HexToAddress(*request.Data.Attributes.SmtLibAddress)
	}
	fmt.Println("smtLibAddress", smtLibAddress)

	var stateLibAddress common.Address
	if request.Data.Attributes.VerifierStateTransition == nil {
		address, tx, _, err := statelib.DeployStatelib(auth, client)
		if err != nil {
			helpers.Log(r).WithError(err).Info("failed to deploy state library contract")
			ape.RenderErr(w, problems.InternalError())
			return
		}
		if _, err = bind.WaitMined(context.Background(), client, tx); err != nil {
			helpers.Log(r).WithError(err).Info(errors.Wrap(err, "failed to wait deploy state contract"))
			ape.RenderErr(w, problems.InternalError())
			return
		}
		auth.Nonce = auth.Nonce.Add(auth.Nonce, big.NewInt(1))
		stateLibAddress = address
	} else {
		stateLibAddress = common.HexToAddress(*request.Data.Attributes.StateLibAddress)
	}

	fmt.Println("stateLibAddress", stateLibAddress)
	stateAddress, tx, _, err := DeployContractWithLinks(auth, client, state.StateMetaData.ABI, state.StateMetaData.Bin, map[string]common.Address{
		"PoseidonUnit1L": common.HexToAddress(*poseidonContracts.El1),
		"StateLib":       stateLibAddress,
		"SmtLib":         smtLibAddress,
	})
	//stateAddress, tx, _, err := state.DeployState(auth, client)
	if err != nil {
		helpers.Log(r).WithError(err).Info("failed to deploy contract")
		ape.RenderErr(w, problems.InternalError())
		return
	}
	stateContract, err := state.NewState(stateAddress, client)
	if err != nil {
		helpers.Log(r).WithError(err).Info("failed to get state contract")
		ape.RenderErr(w, problems.InternalError())
		return
	}
	fmt.Println("stateAddress", stateAddress)
	var defaultTypeId [2]byte
	if request.Data.Attributes.DefaultTypeId == nil {
		defaultTypeId, err = getDefaultTypeId(helpers.NetworkConfig(r).ChainId)
		if err != nil {
			helpers.Log(r).WithError(err).Info("failed to get default type id")
			ape.RenderErr(w, problems.BadRequest(err)...)
			return
		}
	} else {
		_, err = hex.Decode(defaultTypeId[:], []byte(*request.Data.Attributes.DefaultTypeId))
		if err != nil {
			helpers.Log(r).WithError(err).Info("failed to get default type id")
			ape.RenderErr(w, problems.BadRequest(err)...)
			return
		}
	}
	if _, err = bind.WaitMined(context.Background(), client, tx); err != nil {
		helpers.Log(r).WithError(err).Info(errors.Wrap(err, "failed to wait deploy state contract"))
		ape.RenderErr(w, problems.InternalError())
		return
	}
	auth.Nonce = auth.Nonce.Add(auth.Nonce, big.NewInt(1))
	fmt.Println("defaultIdType", defaultTypeId)
	if _, err = stateContract.Initialize(auth, verifierStateTransitionAddress, defaultTypeId); err != nil {
		helpers.Log(r).WithError(err).Info("failed to initialize")
		ape.RenderErr(w, problems.InternalError())
		return
	}
	ape.Render(w, converter.ToTransactionResource("", stateAddress.Hex(), tx.Hash().Hex()))
}

func ABILinkLibrary(bin string, libraryName string, libraryAddress common.Address) string {
	libstr := fmt.Sprintf("_+%v_+", libraryName)
	libraryRexp := regexp.MustCompile(libstr)

	// Remove the 0x prefix from those addresses, just need the actual hex string
	cleanLibraryAddr := strings.Replace(libraryAddress.Hex(), "0x", "", -1)

	modifiedBin := libraryRexp.ReplaceAllString(bin, cleanLibraryAddr)

	return modifiedBin
}

// DeployContractWithLinks patches a contract bin with provided library addresses
func DeployContractWithLinks(
	opts *bind.TransactOpts,
	backend bind.ContractBackend,
	abiString string,
	bin string,
	libraries map[string]common.Address,
	params ...interface{},
) (common.Address, *types.Transaction, *bind.BoundContract, error) {

	for libraryName, libraryAddress := range libraries {
		bin = ABILinkLibrary(bin, libraryName, libraryAddress)
	}

	parsed, err := abi.JSON(strings.NewReader(abiString))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	return bind.DeployContract(opts, parsed, common.FromHex(bin), backend, params...)
}

func getDefaultTypeId(chainId int64) ([2]byte, error) {
	var res [2]byte
	switch chainId {
	case 31337: // hardhat
		temp := common.Hex2Bytes("0212")
		fmt.Println("temp", temp)
		copy(res[:], temp[:])
		return res, nil
	case 80001: // polygon mumbai
		return [2]byte{02, 12}, nil
	case 1101: // zkEVM
		_, err := hex.Decode(res[:], []byte("0231"))
		return res, err
	case 1442: // zkEVM testnet
		_, err := hex.Decode(res[:], []byte("0232"))
		return res, err
	case 137: // polygon main
		_, err := hex.Decode(res[:], []byte("0211"))
		return res, err
	}
	return res, errors.New("no default type id for this network")
}
