package requests

import (
	"encoding/json"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"gitlab.com/distributed_lab/logan/v3/errors"
	"gitlab.com/tokene/polygonid-issuer-integration/resources"
	"net/http"
)

type CreateCredentialRequest struct {
	DidRequest
	Data resources.CreateCredentialRequest
}

func NewCreateCredentialRequest(r *http.Request) (*CreateCredentialRequest, error) {
	var request CreateCredentialRequest

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, errors.Wrap(err, "failed to unmarshal request")
	}
	request.Did = NewDidRequestRequest(r)

	return &request, request.validate()
}

func (r CreateCredentialRequest) validate() error {
	return validation.Errors{
		"data/attributes/credential_schema": validation.Validate(
			&r.Data.Attributes.CredentialSchema,
			validation.Required,
		),
		"data/attributes/credential_subject": validation.Validate(
			&r.Data.Attributes.CredentialSubject,
			validation.Required,
		),
		"data/attributes/type": validation.Validate(
			&r.Data.Attributes.Type,
			validation.Required,
		),
	}.Filter()
}
