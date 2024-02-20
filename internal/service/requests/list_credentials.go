package requests

import (
	"github.com/pkg/errors"
	"net/http"

	"gitlab.com/distributed_lab/urlval"
)

type ListCredentialsRequest struct {
	DidRequest

	SchemaType *string `filter:"schema_type"`
	SchemaHash *string `filter:"schema_hash"`
	Subject    *string `filter:"subject"`
	Revoked    *bool   `filter:"revoked"`
	Self       *bool   `filter:"self"`
	QueryField *string `filter:"query_field"`
	QueryValue *string `filter:"query_value"`
}

func NewListCredentialsRequest(r *http.Request) (*ListCredentialsRequest, error) {
	var request ListCredentialsRequest
	request.Did = NewDidRequestRequest(r)

	err := urlval.Decode(r.URL.Query(), &request)
	if err != nil {
		return nil, errors.Wrap(err, "failed to unmarshal list task request")
	}

	return &request, nil
}
