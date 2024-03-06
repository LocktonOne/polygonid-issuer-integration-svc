package handlers

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
	"gitlab.com/tokene/polygonid-issuer-integration/internal/service/helpers"
	"gitlab.com/tokene/polygonid-issuer-integration/internal/service/helpers/converter"
	"gitlab.com/tokene/polygonid-issuer-integration/solidity/generated/state"
	"gitlab.com/tokene/polygonid-issuer-integration/solidity/generated/statelib"
	"gitlab.com/tokene/polygonid-issuer-integration/solidity/generated/verifierstatetransition"
	"math/big"
	"net/http"
)

func DeployState(w http.ResponseWriter, r *http.Request) {
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

	verifierStateTransitionAddress, tx, _, err := verifierstatetransition.DeployVerifierstatetransition(auth, client)
	if err != nil {
		helpers.Log(r).WithError(err).Info("failed to deploy contract")
		ape.RenderErr(w, problems.InternalError())
		return
	}
	fmt.Println(verifierStateTransitionAddress)
	statelibAddress, tx, _, err := statelib.DeployStatelib(auth, client)
	if err != nil {
		helpers.Log(r).WithError(err).Info("failed to deploy contract")
		ape.RenderErr(w, problems.InternalError())
		return
	}
	fmt.Println(statelibAddress)

	address, tx, _, err := state.DeployState(auth, client)
	if err != nil {
		helpers.Log(r).WithError(err).Info("failed to deploy contract")
		ape.RenderErr(w, problems.InternalError())
		return
	}

	ape.Render(w, converter.ToTransactionResource("", address.Hex(), tx.Hash().Hex()))
}
