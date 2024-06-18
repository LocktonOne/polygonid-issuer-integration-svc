package handlers

import (
	"github.com/google/jsonapi"
	"github.com/spf13/cast"
	"gitlab.com/distributed_lab/ape"
	"gitlab.com/tokene/polygonid-issuer-integration/internal/service/helpers"
	"gitlab.com/tokene/polygonid-issuer-integration/internal/service/helpers/converter"
	"gitlab.com/tokene/polygonid-issuer-integration/internal/service/helpers/issuer-sender"
	"net/http"
)

func GetIdentities(w http.ResponseWriter, r *http.Request) {
	code, response, err := issuer_sender.GetIdentities(r)
	if err != nil {
		helpers.Log(r).WithError(err).Error("failed to get list of identities")
		ape.RenderErr(w, &jsonapi.ErrorObject{
			Title:  http.StatusText(code),
			Detail: err.Error(),
			Status: cast.ToString(code),
		})
		return
	}

	w.WriteHeader(code)
	ape.Render(w, converter.ToListIdentitiesResource(response))
}
