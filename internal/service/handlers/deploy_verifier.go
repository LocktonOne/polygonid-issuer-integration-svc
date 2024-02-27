package handlers

import (
	"context"
	"crypto/ecdsa"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
	"gitlab.com/tokene/polygonid-issuer-integration/internal/service/helpers"
	"gitlab.com/tokene/polygonid-issuer-integration/internal/service/helpers/converter"
	"gitlab.com/tokene/polygonid-issuer-integration/internal/service/requests"
	"gitlab.com/tokene/polygonid-issuer-integration/solidity/generated/erc20verifier"
	"math/big"
	"net/http"
)

func DeployOnChainVerifier(w http.ResponseWriter, r *http.Request) {
	request, err := requests.NewDeployVerifierRequest(r)
	if err != nil {
		helpers.Log(r).WithError(err).Info("wrong request")
		ape.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	client, err := ethclient.Dial(helpers.NetworkConfig(r).RpcUrl.String())
	if err != nil {
		helpers.Log(r).WithError(err).Info("failed to connect client to the RPC URL")
		ape.RenderErr(w, problems.InternalError())
		return
	}
	publicKeyECDSA, ok := helpers.NetworkConfig(r).PrivateKey.Public().(*ecdsa.PublicKey)
	if !ok {
		helpers.Log(r).WithError(err).Info("failed to casting public key to ECDSA")
		ape.RenderErr(w, problems.InternalError())
		return
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		helpers.Log(r).WithError(err).Info("failed to to get account nonce")
		ape.RenderErr(w, problems.InternalError())
		return
	}

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		helpers.Log(r).WithError(err).Info("failed to get gas price")
		ape.RenderErr(w, problems.InternalError())
		return
	}

	auth, err := bind.NewKeyedTransactorWithChainID(helpers.NetworkConfig(r).PrivateKey, big.NewInt(helpers.NetworkConfig(r).ChainId))
	if err != nil {
		helpers.Log(r).WithError(err).Info("failed to deploy contract")
		ape.RenderErr(w, problems.InternalError())
		return
	}

	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0) // in wei
	auth.GasPrice = gasPrice
	if helpers.NetworkConfig(r).GasLimit != 0 {
		auth.GasLimit = helpers.NetworkConfig(r).GasLimit
	}

	address, tx, _, err := erc20verifier.DeployErc20verifier(auth, client, request.Data.Attributes.Name, request.Data.Attributes.Symbol)
	if err != nil {
		helpers.Log(r).WithError(err).Info("failed to deploy contract")
		ape.RenderErr(w, problems.InternalError())
		return
	}

	ape.Render(w, converter.ToDeployVerifierResource(address.Hex(), tx.Hash().Hex()))
}
