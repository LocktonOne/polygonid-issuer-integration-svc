package requests

import (
	"encoding/json"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"gitlab.com/distributed_lab/logan/v3/errors"
	"gitlab.com/tokene/polygonid-issuer-integration/resources"
	"net/http"
)

type CreateIdentityRequest struct {
	Data resources.CreateIdentityRequest
}

func NewCreateIdentityRequest(r *http.Request) (*CreateIdentityRequest, error) {
	var request CreateIdentityRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, errors.Wrap(err, "failed to unmarshal request")
	}

	return &request, request.validate()
}

func (r CreateIdentityRequest) validate() error {
	return validation.Errors{
		"data/attributes/method": validation.Validate(
			&r.Data.Attributes.Method,
			validation.Required,
		),
		"data/attributes/network": validation.Validate(
			&r.Data.Attributes.Network,
			validation.Required,
		),
		"data/attributes/blockchain": validation.Validate(
			&r.Data.Attributes.Blockchain,
			validation.Required,
		),
		"data/attributes/type": validation.Validate(
			&r.Data.Attributes.Type,
			validation.Required,
		),
	}.Filter()
}
