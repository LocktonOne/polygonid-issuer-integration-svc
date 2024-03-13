package requests

import (
	"encoding/json"
	"github.com/pkg/errors"
	"gitlab.com/tokene/polygonid-issuer-integration/resources"
	"net/http"
)

type DeployStateRequest struct {
	Data resources.DeployStateRequest
}

func NewDeployStateRequest(r *http.Request) (*DeployStateRequest, error) {
	var request DeployStateRequest

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, errors.Wrap(err, "failed to unmarshal request")
	}

	return &request, nil
}
