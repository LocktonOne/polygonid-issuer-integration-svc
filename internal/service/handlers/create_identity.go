package handlers

import (
	"github.com/google/jsonapi"
	"github.com/spf13/cast"
	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
	"gitlab.com/tokene/polygonid-issuer-integration/internal/data"
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
	user, err := helpers.DidQ(r).FilterByAddress(helpers.Address(r)).Get()
	if err != nil {
		helpers.Log(r).WithError(err).Error("failed to insert identity")
		ape.RenderErr(w, problems.InternalError())
		return
	}
	if user != nil {
		helpers.Log(r).WithError(err).Error("identity is already exists")
		ape.RenderErr(w, problems.Conflict())
		return
	}

	issuerRequest := resources.IssuerCreateIdentityRequest{
		DidMetadata: resources.DidMetadata{
			Blockchain: request.Data.Attributes.Blockchain,
			Method:     request.Data.Attributes.Method,
			Network:    request.Data.Attributes.Network,
			Type:       request.Data.Attributes.Type,
		},
	}

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
	addr := helpers.Address(r)
	if err = helpers.DidQ(r).Insert(data.Did{
		DID:     *response.Identifier,
		Address: helpers.Address(r),
	}); err != nil {
		helpers.Log(r).WithError(err).Error("failed to insert identity")
		ape.RenderErr(w, problems.InternalError())
		return
	}
	response.Address = &addr
	w.WriteHeader(code)
	ape.Render(w, converter.ToCreateIdentityResource(*response))
}
