package issuer_sender

import (
	"encoding/json"
	"fmt"
	"gitlab.com/distributed_lab/logan/v3/errors"
	"gitlab.com/tokene/polygonid-issuer-integration/internal/service/helpers"
	"gitlab.com/tokene/polygonid-issuer-integration/internal/service/types"
	"gitlab.com/tokene/polygonid-issuer-integration/resources"
	"net/http"
)

func GetIdentities(r *http.Request) (int, []string, error) {
	issuerConfig := helpers.IssuerConfig(r)

	issuerPath := fmt.Sprintf("%s/v1/identities", issuerConfig.Endpoint)

	code, raw, err := IssuerGetRequestWithBasicAuth(r, issuerPath, nil)
	if err != nil {
		return code, nil, errors.Wrap(err, "failed to read notification service response")
	}
	fmt.Println(string(raw))
	var issuerResp []string
	if code >= http.StatusBadRequest {
		var errResp helpers.Error
		if err = json.Unmarshal(raw, &errResp); err != nil {
			return http.StatusInternalServerError, nil, errors.Wrap(err, "failed to unmarshal service response")
		}
		return code, nil, errors.New(errResp.Message)
	}
	if err = json.Unmarshal(raw, &issuerResp); err != nil {
		return http.StatusInternalServerError, nil, errors.Wrap(err, "failed to unmarshal service response")
	}

	return code, issuerResp, nil
}

func GetIdentityDetail(r *http.Request, did string) (int, *types.GetIdentityDetailsResponse, error) {
	issuerConfig := helpers.IssuerConfig(r)

	issuerPath := fmt.Sprintf("%s/v1/identities/%s/details", issuerConfig.Endpoint, did)

	code, raw, err := IssuerGetRequestWithBasicAuth(r, issuerPath, nil)
	if err != nil {
		return code, nil, errors.Wrap(err, "failed to read notification service response")
	}
	fmt.Println(string(raw))
	var issuerResp types.GetIdentityDetailsResponse
	if code >= http.StatusBadRequest {
		var errResp helpers.Error
		if err = json.Unmarshal(raw, &errResp); err != nil {
			return http.StatusInternalServerError, nil, errors.Wrap(err, "failed to unmarshal service response")
		}
		return code, nil, errors.New(errResp.Message)
	}
	if err = json.Unmarshal(raw, &issuerResp); err != nil {
		return http.StatusInternalServerError, nil, errors.Wrap(err, "failed to unmarshal service response")
	}

	return code, &issuerResp, nil
}

func CreateIdentity(r *http.Request, request resources.IssuerCreateIdentityRequest) (int, *types.CreateIdentityResponse, error) {
	issuerConfig := helpers.IssuerConfig(r)

	issuerPath := fmt.Sprintf("%s/v1/identities", issuerConfig.Endpoint)

	marshaled, err := json.Marshal(request)
	if err != nil {
		return http.StatusInternalServerError, nil, errors.Wrap(err, "failed to marshal request to issuer")
	}

	code, raw, err := IssuerRequestWithBasicAuth(r, issuerPath, marshaled)
	if err != nil {
		return code, nil, errors.Wrap(err, "failed to read notification service response")
	}

	fmt.Println(string(raw))
	if code >= http.StatusBadRequest {
		var errResp helpers.Error
		if err = json.Unmarshal(raw, &errResp); err != nil {
			return http.StatusInternalServerError, nil, errors.Wrap(err, "failed to unmarshal service response")
		}
		return code, nil, errors.New(errResp.Message)
	}

	var issuerResp types.CreateIdentityResponse
	if err = json.Unmarshal(raw, &issuerResp); err != nil {
		return http.StatusInternalServerError, nil, errors.Wrap(err, "failed to unmarshal service response")
	}

	return code, &issuerResp, nil
}

func PublishIdentityState(r *http.Request, did string) (int, *types.PublishIdentityStateResponse, error) {
	issuerConfig := helpers.IssuerConfig(r)

	issuerPath := fmt.Sprintf("%s/v1/%s/state/publish", issuerConfig.Endpoint, did)

	code, raw, err := IssuerRequestWithBasicAuth(r, issuerPath, nil)
	if err != nil {
		return code, nil, errors.Wrap(err, "failed to read notification service response")
	}
	fmt.Println(string(raw))
	var issuerResp types.PublishIdentityStateResponse
	if code >= http.StatusBadRequest || code == http.StatusOK {
		var errResp helpers.Error
		if err = json.Unmarshal(raw, &errResp); err != nil {
			return http.StatusInternalServerError, nil, errors.Wrap(err, "failed to unmarshal service response")
		}
		return code, nil, errors.New(errResp.Message)
	}

	if err = json.Unmarshal(raw, &issuerResp); err != nil {
		return http.StatusInternalServerError, nil, errors.Wrap(err, "failed to unmarshal service response")
	}

	return code, &issuerResp, nil
}

func RetryPublishIdentityState(r *http.Request, did string) (int, *types.PublishIdentityStateResponse, error) {
	issuerConfig := helpers.IssuerConfig(r)

	issuerPath := fmt.Sprintf("%s/v1/%s/state/retry", issuerConfig.Endpoint, did)

	code, raw, err := IssuerRequestWithBasicAuth(r, issuerPath, nil)
	if err != nil {
		return code, nil, errors.Wrap(err, "failed to read notification service response")
	}
	fmt.Println(string(raw))
	var issuerResp types.PublishIdentityStateResponse
	if code >= http.StatusBadRequest {
		var errResp helpers.Error
		if err = json.Unmarshal(raw, &errResp); err != nil {
			return http.StatusInternalServerError, nil, errors.Wrap(err, "failed to unmarshal service response")
		}
		if code == http.StatusOK {
			code = http.StatusBadRequest
		}
		return code, nil, errors.New(errResp.Message)
	}
	if err = json.Unmarshal(raw, &issuerResp); err != nil {
		return http.StatusInternalServerError, nil, errors.Wrap(err, "failed to unmarshal service response")
	}

	return code, &issuerResp, nil
}
