package requests

import (
	"encoding/json"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"gitlab.com/distributed_lab/logan/v3/errors"
	"gitlab.com/tokene/polygonid-issuer-integration/resources"
	"net/http"
)

type SetZKPRequest struct {
	Data resources.SetZkpRequest
}

func NewSetZKPRequest(r *http.Request) (*SetZKPRequest, error) {
	var request SetZKPRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, errors.Wrap(err, "failed to unmarshal request")
	}

	return &request, request.validate()
}

func (r SetZKPRequest) validate() error {
	return validation.Errors{
		"data/attributes/schema_url": validation.Validate(
			&r.Data.Attributes.SchemaUrl,
			validation.Required,
		),
		"data/attributes/schema_type": validation.Validate(
			&r.Data.Attributes.SchemaType,
			validation.Required,
		),
		"data/attributes/circuit_id": validation.Validate(
			&r.Data.Attributes.CircuitId,
			validation.Required,
		),
		"data/attributes/claim": validation.Validate(
			&r.Data.Attributes.Claim,
			validation.Required,
		),
		"data/attributes/field_name": validation.Validate(
			&r.Data.Attributes.FieldName,
			validation.Required,
		),
		"data/attributes/operator": validation.Validate(
			&r.Data.Attributes.Operator,
			validation.Required,
		),
		"data/attributes/reason": validation.Validate(
			&r.Data.Attributes.Reason,
			validation.Required,
		),
		"data/attributes/request_id": validation.Validate(
			&r.Data.Attributes.RequestId,
			validation.Required,
		),
		"data/attributes/validator_address": validation.Validate(
			&r.Data.Attributes.ValidatorAddress,
			validation.Required,
		),
		"data/attributes/value": validation.Validate(
			&r.Data.Attributes.Value,
			validation.Required,
		),
		"data/attributes/verifier_address": validation.Validate(
			&r.Data.Attributes.VerifierAddress,
			validation.Required,
		),
	}.Filter()
}
