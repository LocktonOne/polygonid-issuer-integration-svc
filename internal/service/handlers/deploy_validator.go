package handlers

import (
	"context"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/pkg/errors"
	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
	"gitlab.com/tokene/polygonid-issuer-integration/internal/service/helpers"
	"gitlab.com/tokene/polygonid-issuer-integration/internal/service/helpers/converter"
	"gitlab.com/tokene/polygonid-issuer-integration/internal/service/requests"
	"gitlab.com/tokene/polygonid-issuer-integration/solidity/generated/credentialatomicquerymtpvalidator"
	"gitlab.com/tokene/polygonid-issuer-integration/solidity/generated/credentialatomicquerysigvalidator"
	"gitlab.com/tokene/polygonid-issuer-integration/solidity/generated/verifiermtpwrapper"
	"gitlab.com/tokene/polygonid-issuer-integration/solidity/generated/verifiersigwrapper"
	"math/big"
	"net/http"
)

func DeployValidator(w http.ResponseWriter, r *http.Request) {
	request, err := requests.NewDeployValidatorRequest(r)
	if err != nil {
		helpers.Log(r).WithError(err).Info("wrong request")
		ape.RenderErr(w, problems.BadRequest(err)...)
		return
	}
	validatorContract, tx, err := deployValidator(r, request.Data.Attributes.ValidatorType, request.Data.Attributes.StateAddress)
	if err != nil {
		helpers.Log(r).WithError(err).Info("failed to deploy validator")
		ape.RenderErr(w, problems.InternalError())
		return
	}

	ape.Render(w, converter.ToTransactionResource("", validatorContract.Hex(), tx.Hash().Hex()))
}

func deployValidator(r *http.Request, validatorType, state string) (*common.Address, *types.Transaction, error) {
	auth, client, err := helpers.GetEthClient(r)
	if err != nil {
		return nil, nil, errors.Wrap(err, "failed to get eth client")
	}

	switch validatorType {
	case "mtp":
		verifierMTPWrapperAddress, _, _, err := verifiermtpwrapper.DeployVerifiermtpwrapper(auth, client)
		if err != nil {
			return nil, nil, errors.Wrap(err, "failed to deploy verifier MTP wrapper contract")
		}
		auth.Nonce = auth.Nonce.Add(auth.Nonce, big.NewInt(1))

		credentialAtomicQueryMTPValidatorAddress, tx, credentialAtomicQueryMTPValidatorContract, err := credentialatomicquerymtpvalidator.DeployCredentialatomicquerymtpvalidator(auth, client)
		if err != nil {
			return nil, nil, errors.Wrap(err, "failed to deploy credential Atomic Query MTP Validator")
		}
		if _, err = bind.WaitMined(context.Background(), client, tx); err != nil {
			return nil, nil, errors.Wrap(err, "failed to wait deploy credential Atomic Query MTP Validator")
		}
		auth.Nonce = auth.Nonce.Add(auth.Nonce, big.NewInt(1))

		if _, err = credentialAtomicQueryMTPValidatorContract.Initialize(auth, verifierMTPWrapperAddress, common.HexToAddress(state)); err != nil {
			return nil, nil, errors.Wrap(err, "failed to initialize credential Atomic Query MTP Validator")
		}
		return &credentialAtomicQueryMTPValidatorAddress, tx, nil
	case "sig":
		verifierSigWrapperAddress, _, _, err := verifiersigwrapper.DeployVerifiersigwrapper(auth, client)
		if err != nil {
			return nil, nil, errors.Wrap(err, "failed to deploy verifier Sig wrapper contract")
		}
		auth.Nonce = auth.Nonce.Add(auth.Nonce, big.NewInt(1))

		credentialAtomicQuerySigValidatorAddress, tx, credentialAtomicQuerySigValidatorContract, err := credentialatomicquerysigvalidator.DeployCredentialatomicquerysigvalidator(auth, client)
		if err != nil {
			return nil, nil, errors.Wrap(err, "failed to deploy credential Atomic Query Sig Validator")
		}
		if _, err = bind.WaitMined(context.Background(), client, tx); err != nil {
			return nil, nil, errors.Wrap(err, "failed to wait deploy credential Atomic Query MTP Validator")
		}
		auth.Nonce = auth.Nonce.Add(auth.Nonce, big.NewInt(1))

		if _, err = credentialAtomicQuerySigValidatorContract.Initialize(auth, verifierSigWrapperAddress, common.HexToAddress(state)); err != nil {
			return nil, nil, errors.Wrap(err, "failed to initialize credential Atomic Query Sig Validator")
		}
		return &credentialAtomicQuerySigValidatorAddress, tx, nil
	default:
		return nil, nil, errors.New("no such type " + validatorType)
	}
}
