package handlers

import (
	"github.com/pkg/errors"
	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
	"gitlab.com/tokene/polygonid-issuer-integration/internal/service/helpers"
	"gitlab.com/tokene/polygonid-issuer-integration/internal/service/requests"
	"gitlab.com/tokene/polygonid-issuer-integration/resources"
	"net/http"
)

func GetIdentityByAddress(w http.ResponseWriter, r *http.Request) {
	addr, err := requests.NewAddressRequestRequest(r)

	if err != nil {
		helpers.Log(r).WithError(errors.Wrap(err, "failed to parse request"))
		ape.RenderErr(w, problems.BadRequest(errors.Wrap(err, "failed to parse request"))...)
		return
	}

	did, err := helpers.DidQ(r).FilterByAddress(addr).Get()
	if err != nil {
		helpers.Log(r).WithError(errors.Wrap(err, "failed to get identity"))
		ape.RenderErr(w, problems.InternalError())
		return
	}
	if did == nil {
		helpers.Log(r).WithError(errors.Wrap(err, "identity not found"))
		ape.RenderErr(w, problems.NotFound())
		return
	}
	ape.Render(w, resources.Key{
		ID:   did.DID,
		Type: resources.IDENTITY,
	})
}
