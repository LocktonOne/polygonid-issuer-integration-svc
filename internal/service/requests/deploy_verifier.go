package requests

import (
	"encoding/json"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"gitlab.com/distributed_lab/logan/v3/errors"
	"gitlab.com/tokene/polygonid-issuer-integration/resources"
	"net/http"
)

type DeployVerifierRequest struct {
	Data resources.DeployVerifierRequest
}

func NewDeployVerifierRequest(r *http.Request) (*DeployVerifierRequest, error) {
	var request DeployVerifierRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, errors.Wrap(err, "failed to unmarshal request")
	}

	return &request, request.validate()
}

func (r DeployVerifierRequest) validate() error {
	return validation.Errors{
		"data/attributes/name": validation.Validate(
			&r.Data.Attributes.Name,
			validation.Required,
		),
		"data/attributes/symbol": validation.Validate(
			&r.Data.Attributes.Symbol,
			validation.Required,
		),
	}.Filter()
}
