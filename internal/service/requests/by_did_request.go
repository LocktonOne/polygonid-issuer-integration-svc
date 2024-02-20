package requests

import (
	"github.com/go-chi/chi"
	"net/http"
)

type DidRequest struct {
	Did string
}

func NewDidRequestRequest(r *http.Request) string {
	return chi.URLParam(r, "did")
}
