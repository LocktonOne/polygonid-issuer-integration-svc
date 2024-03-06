package handlers

import (
	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
	"gitlab.com/tokene/polygonid-issuer-integration/internal/service/helpers"
	"gitlab.com/tokene/polygonid-issuer-integration/internal/service/helpers/converter"
	"gitlab.com/tokene/polygonid-issuer-integration/solidity/generated/zkpverifier"
	"net/http"
)

func DeployOnChainVerifier(w http.ResponseWriter, r *http.Request) {
	auth, client, err := helpers.GetEthClient(r)
	if err != nil {
		helpers.Log(r).WithError(err).Info("failed to get eth client")
		ape.RenderErr(w, problems.InternalError())
		return
	}

	address, tx, _, err := zkpverifier.DeployZkpverifier(auth, client)
	if err != nil {
		helpers.Log(r).WithError(err).Info("failed to deploy contract")
		ape.RenderErr(w, problems.InternalError())
		return
	}

	ape.Render(w, converter.ToTransactionResource("", address.Hex(), tx.Hash().Hex()))
}
