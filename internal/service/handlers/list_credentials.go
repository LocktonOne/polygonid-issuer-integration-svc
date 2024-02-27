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
	"gitlab.com/tokene/polygonid-issuer-integration/internal/service/types"
	"net/http"
	"strings"
)

func ListCredentials(w http.ResponseWriter, r *http.Request) {
	request, err := requests.NewListCredentialsRequest(r)
	if err != nil {
		helpers.Log(r).WithError(err).Error("failed to fetch new sign mint request")
		ape.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	if !strings.HasPrefix(request.Did, "did:") {
		helpers.Log(r).WithError(errors.New("DID is empty or doesn't begin with \"did:\"")).Info("wrong did")
		ape.RenderErr(w, problems.BadRequest(errors.New("DID is empty or doesn't begin with \"did:\""))...)
		return
	}

	code, response, err := isuer_sender.ListCredential(r, request.Did, types.GetClaimsParams{
		SchemaType: request.SchemaType,
		SchemaHash: request.SchemaHash,
		Subject:    request.Subject,
		Revoked:    request.Revoked,
		Self:       request.Self,
		QueryField: request.QueryField,
		QueryValue: request.QueryValue,
	})
	if err != nil {
		helpers.Log(r).WithError(err).Error("failed to get list of credentials")
		ape.RenderErr(w, &jsonapi.ErrorObject{
			Title:  http.StatusText(code),
			Detail: err.Error(),
			Status: cast.ToString(code),
		})
		return
	}

	w.WriteHeader(code)
	ape.Render(w, converter.ToGetCredentialListResource(*response))
}
