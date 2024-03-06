package requests

import (
	"encoding/json"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/pkg/errors"
	"gitlab.com/tokene/polygonid-issuer-integration/resources"
	"net/http"
)

type DeployValidatorRequest struct {
	Data resources.DeployValidator
}

func NewDeployValidatorRequest(r *http.Request) (*DeployValidatorRequest, error) {
	var request DeployValidatorRequest

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, errors.Wrap(err, "failed to unmarshal request")
	}

	return &request, request.validate()
}

func (r DeployValidatorRequest) validate() error {
	return validation.Errors{
		"data/attributes/state_address": validation.Validate(
			&r.Data.Attributes.StateAddress,
			validation.Required,
		),
		"data/attributes/validator_type": validation.Validate(
			&r.Data.Attributes.ValidatorType,
			validation.Required,
			validation.In("sig", "mtp"),
		),
	}.Filter()
}
