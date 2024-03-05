package issuer_sender

import (
	"encoding/json"
	"fmt"
	"gitlab.com/distributed_lab/logan/v3/errors"
	"gitlab.com/tokene/polygonid-issuer-integration/internal/service/helpers"
	"gitlab.com/tokene/polygonid-issuer-integration/internal/service/types"
	"net/http"
	"strconv"
)

func GetCredential(r *http.Request, did, claim string) (int, *types.GetClaimResponse, error) {
	issuerConfig := helpers.IssuerConfig(r)

	issuerPath := fmt.Sprintf("%s/v1/%s/claims/%s", issuerConfig.Endpoint, did, claim)

	code, raw, err := IssuerGetRequestWithBasicAuth(r, issuerPath, nil)
	if err != nil {
		return code, nil, errors.Wrap(err, "failed to read notification service response")
	}

	var issuerResp types.GetClaimResponse
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

func GetCredentialQRCode(r *http.Request, did, claim string) (int, *types.GetClaimQrCodeResponse, error) {
	issuerConfig := helpers.IssuerConfig(r)

	issuerPath := fmt.Sprintf("%s/v1/%s/claims/%s/qrcode", issuerConfig.Endpoint, did, claim)

	code, raw, err := IssuerGetRequestWithBasicAuth(r, issuerPath, nil)
	if err != nil {
		return code, nil, errors.Wrap(err, "failed to read notification service response")
	}

	var issuerResp types.GetClaimQrCodeResponse
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

func ListCredential(r *http.Request, did string, params types.GetClaimsParams) (int, *types.GetClaimsResponse, error) {
	issuerConfig := helpers.IssuerConfig(r)

	issuerPath := fmt.Sprintf("%s/v1/%s/claims", issuerConfig.Endpoint, did)

	paramsMap := map[string]string{}
	if params.SchemaHash != nil {
		paramsMap["schemaHash"] = *params.SchemaHash
	}
	if params.SchemaType != nil {
		paramsMap["schemaType"] = *params.SchemaType
	}
	if params.Subject != nil {
		paramsMap["subject"] = *params.Subject
	}
	if params.Revoked != nil {
		paramsMap["revoked"] = strconv.FormatBool(*params.Revoked)
	}
	if params.Self != nil {
		paramsMap["self"] = strconv.FormatBool(*params.Self)
	}
	if params.QueryField != nil {
		paramsMap["query_field"] = *params.QueryField
	}
	if params.QueryValue != nil {
		paramsMap["query_value"] = *params.QueryValue
	}

	code, raw, err := IssuerGetRequestWithBasicAuth(r, issuerPath, paramsMap)
	if err != nil {
		return code, nil, errors.Wrap(err, "failed to read notification service response")
	}

	var issuerResp types.GetClaimsResponse
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

func GetCredentialRevocationStatus(r *http.Request, did, nonce string) (int, *types.RevocationStatusResponse, error) {
	issuerConfig := helpers.IssuerConfig(r)

	issuerPath := fmt.Sprintf("%s/v1/%s/claims/revocation/status/%s", issuerConfig.Endpoint, did, nonce)

	code, raw, err := IssuerGetRequestWithBasicAuth(r, issuerPath, nil)
	if err != nil {
		return code, nil, errors.Wrap(err, "failed to read notification service response")
	}

	var issuerResp types.RevocationStatusResponse
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

func CreateCredential(r *http.Request, did string, request types.CreateClaimRequest) (int, string, error) {
	issuerConfig := helpers.IssuerConfig(r)

	issuerPath := fmt.Sprintf("%s/v1/%s/claims", issuerConfig.Endpoint, did)

	marshaled, err := json.Marshal(request)
	if err != nil {
		return http.StatusInternalServerError, "", errors.Wrap(err, "failed to marshal request to issuer")
	}

	code, raw, err := IssuerRequestWithBasicAuth(r, issuerPath, marshaled)
	if err != nil {
		return code, "", errors.Wrap(err, "failed to read notification service response")
	}

	var issuerResp types.CreateClaimResponse
	if code >= http.StatusBadRequest {
		var errResp helpers.Error
		if err = json.Unmarshal(raw, &errResp); err != nil {
			return http.StatusInternalServerError, "", errors.Wrap(err, "failed to unmarshal service response")
		}
		return code, "", errors.New(errResp.Message)
	}
	if err = json.Unmarshal(raw, &issuerResp); err != nil {
		return http.StatusInternalServerError, "", errors.Wrap(err, "failed to unmarshal service response")
	}

	return code, issuerResp.Id, nil
}

func RevokeCredential(r *http.Request, did, nonce string) (int, *types.RevokeClaimResponse, error) {
	issuerConfig := helpers.IssuerConfig(r)

	issuerPath := fmt.Sprintf("%s/v1/%s/claims/revoke/%s", issuerConfig.Endpoint, did, nonce)

	code, raw, err := IssuerRequestWithBasicAuth(r, issuerPath, nil)
	if err != nil {
		return code, nil, errors.Wrap(err, "failed to read notification service response")
	}

	var issuerResp types.RevokeClaimResponse
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
