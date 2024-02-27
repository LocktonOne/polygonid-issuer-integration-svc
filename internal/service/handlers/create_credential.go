package handlers

import (
	"github.com/google/jsonapi"
	"github.com/spf13/cast"
	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
	"gitlab.com/tokene/polygonid-issuer-integration/internal/service/helpers"
	isuer_sender "gitlab.com/tokene/polygonid-issuer-integration/internal/service/helpers/issuer-sender"
	"gitlab.com/tokene/polygonid-issuer-integration/internal/service/requests"
	"gitlab.com/tokene/polygonid-issuer-integration/internal/service/types"
	"gitlab.com/tokene/polygonid-issuer-integration/resources"
	"net/http"
	"strings"
	"time"
)

func CreateCredential(w http.ResponseWriter, r *http.Request) {
	request, err := requests.NewCreateCredentialRequest(r)
	if err != nil {
		helpers.Log(r).WithError(err).Info("wrong request")
		ape.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	if !strings.HasPrefix(request.Did, "did:") {
		helpers.Log(r).WithError(err).Info("wrong did")
		ape.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	issuerRequest := types.CreateClaimRequest{
		CredentialSchema:  request.Data.Attributes.CredentialSchema,
		Expiration:        nil,
		Type:              request.Data.Attributes.Type,
		CredentialSubject: request.Data.Attributes.CredentialSubject,
	}

	if request.Data.Attributes.Expiration != nil {
		expiration := time.Now().Add(*request.Data.Attributes.Expiration).Unix()
		issuerRequest.Expiration = &expiration
	}

	code, credential, err := isuer_sender.CreateCredential(r, request.Did, issuerRequest)
	if err != nil {
		helpers.Log(r).WithError(err).Error("failed to create credential")
		ape.RenderErr(w, &jsonapi.ErrorObject{
			Title:  http.StatusText(code),
			Detail: err.Error(),
			Status: cast.ToString(code),
		})
		return
	}

	w.WriteHeader(code)
	ape.Render(w, resources.Key{
		ID:   credential,
		Type: resources.CREATE_CREDENTIAL,
	}.AsRelation(),
	)
}
