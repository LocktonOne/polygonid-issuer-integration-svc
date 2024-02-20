package handlers

import (
	"github.com/google/jsonapi"
	"github.com/spf13/cast"
	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
	"gitlab.com/tokene/polygonid-issuer-integration/internal/service/helpers"
	"gitlab.com/tokene/polygonid-issuer-integration/internal/service/helpers/converter"
	"gitlab.com/tokene/polygonid-issuer-integration/internal/service/helpers/issuer-sender"
	"gitlab.com/tokene/polygonid-issuer-integration/internal/service/requests"
	"gitlab.com/tokene/polygonid-issuer-integration/resources"
	"net/http"
)

func CreateIdentity(w http.ResponseWriter, r *http.Request) {
	request, err := requests.NewCreateIdentityRequest(r)
	if err != nil {
		helpers.Log(r).WithError(err).Info("wrong request")
		ape.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	issuerRequest := resources.IssuerCreateIdentityRequest{
		DidMetadata: resources.DidMetadata{
			Blockchain: request.Data.Attributes.Blockchain,
			Method:     request.Data.Attributes.Method,
			Network:    request.Data.Attributes.Network,
			Type:       request.Data.Attributes.Type,
		}}

	code, response, err := issuer_sender.CreateIdentity(r, issuerRequest)
	if err != nil {
		helpers.Log(r).WithError(err).Error("failed to create identity")
		ape.RenderErr(w, &jsonapi.ErrorObject{
			Title:  http.StatusText(code),
			Detail: err.Error(),
			Status: cast.ToString(code),
		})
		return
	}

	w.Header().Set("content-type", jsonapi.MediaType)
	w.WriteHeader(code)
	ape.Render(w, converter.ToCreateIdentityResource(*response))
}
