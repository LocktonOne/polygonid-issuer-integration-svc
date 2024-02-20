package requests

import (
	"github.com/go-chi/chi"
	"net/http"
)

type NonceRequest struct {
	DidRequest
	Nonce string
}

func NewNonceRequest(r *http.Request) *NonceRequest {
	var request NonceRequest

	request.Did = NewDidRequestRequest(r)

	request.Nonce = chi.URLParam(r, "nonce")

	return &request
}
