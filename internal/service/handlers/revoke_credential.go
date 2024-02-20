package handlers

import (
	"errors"
	"github.com/google/jsonapi"
	"github.com/spf13/cast"
	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
	"gitlab.com/tokene/polygonid-issuer-integration/internal/service/helpers"
	"gitlab.com/tokene/polygonid-issuer-integration/internal/service/helpers/converter"
	isuer_sender "gitlab.com/tokene/polygonid-issuer-integration/internal/service/helpers/issuer-sender"
	"gitlab.com/tokene/polygonid-issuer-integration/internal/service/requests"
	"net/http"
	"strings"
)

func RevokeCredential(w http.ResponseWriter, r *http.Request) {
	request := requests.NewNonceRequest(r)

	if !strings.HasPrefix(request.Did, "did:") {
		helpers.Log(r).WithError(errors.New("DID is empty or doesn't begin with \"did:\"")).Info("wrong did")
		ape.RenderErr(w, problems.BadRequest(errors.New("DID is empty or doesn't begin with \"did:\""))...)
		return
	}

	code, response, err := isuer_sender.RevokeCredential(r, request.Did, request.Nonce)
	if err != nil {
		helpers.Log(r).WithError(err).Error("failed to revoke credential")
		ape.RenderErr(w, &jsonapi.ErrorObject{
			Title:  http.StatusText(code),
			Detail: err.Error(),
			Status: cast.ToString(code),
		})
		return
	}

	w.Header().Set("content-type", jsonapi.MediaType)
	w.WriteHeader(code)
	ape.Render(w, converter.ToRevokeCredentialResource(response.Message))
}