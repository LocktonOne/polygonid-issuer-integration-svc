package issuer_sender

import (
	"bytes"
	"gitlab.com/distributed_lab/logan/v3/errors"
	"gitlab.com/tokene/polygonid-issuer-integration/internal/service/helpers"
	"io"
	"net/http"
	"time"
)

func IssuerRequestWithBasicAuth(r *http.Request, path string, request []byte) (int, []byte, error) {
	issuerConfig := helpers.IssuerConfig(r)

	client := http.Client{Timeout: 10 * time.Second}
	req, err := http.NewRequest(http.MethodPost, path, bytes.NewReader(request))
	if err != nil {
		return http.StatusBadRequest, nil, errors.Wrap(err, "failed to create request")
	}

	req.SetBasicAuth(issuerConfig.Username, issuerConfig.Password)

	resp, err := client.Do(req)
	if err != nil {
		return http.StatusInternalServerError, nil, errors.Wrap(err, "failed to send request")
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusNotFound {
		return resp.StatusCode, nil, nil
	}

	raw, err := io.ReadAll(resp.Body)
	if err != nil {
		return http.StatusInternalServerError, nil, errors.Wrap(err, "failed to read notification service response")
	}

	return resp.StatusCode, raw, nil
}

func IssuerGetRequestWithBasicAuth(r *http.Request, path string, params map[string]string) (int, []byte, error) {
	issuerConfig := helpers.IssuerConfig(r)

	client := http.Client{Timeout: 10 * time.Second}
	req, err := http.NewRequest(http.MethodGet, path, nil)
	if err != nil {
		return http.StatusBadRequest, nil, errors.Wrap(err, "failed to create request")
	}
	q := req.URL.Query()
	for key, value := range params {
		q.Add(key, value)
	}
	req.URL.RawQuery = q.Encode()

	req.SetBasicAuth(issuerConfig.Username, issuerConfig.Password)

	resp, err := client.Do(req)
	if err != nil {
		return http.StatusInternalServerError, nil, errors.Wrap(err, "failed to send request")
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusNotFound {
		return resp.StatusCode, nil, nil
	}

	raw, err := io.ReadAll(resp.Body)
	if err != nil {
		return http.StatusInternalServerError, nil, errors.Wrap(err, "failed to read notification service response")
	}

	return resp.StatusCode, raw, nil
}
