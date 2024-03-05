package requests

import (
	"github.com/go-chi/chi"
	"net/http"
)

type GetCredentialRequest struct {
	DidRequest
	Claim string
}

func NewGetCredentialRequest(r *http.Request) *GetCredentialRequest {
	var request GetCredentialRequest

	request.Did = NewDidRequestRequest(r)

	request.Claim = chi.URLParam(r, "claim")

	return &request
}
