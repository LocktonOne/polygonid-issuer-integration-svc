package requests

import (
	"github.com/go-chi/chi"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"net/http"
	"regexp"
	"strings"
)

type AddressRequest struct {
	Address string
}

func NewAddressRequestRequest(r *http.Request) (string, error) {
	var req AddressRequest
	req.Address = strings.ToLower(chi.URLParam(r, "address"))
	return req.Address, req.validate()
}

func (r AddressRequest) validate() error {
	return validation.Errors{
		"address": validation.Validate(
			&r.Address,
			validation.Required,
			validation.Match(regexp.MustCompile("^0x[0-9a-f]{40}$")),
		),
	}.Filter()
}
