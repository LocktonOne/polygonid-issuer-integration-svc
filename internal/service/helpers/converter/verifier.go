package converter

import (
	"gitlab.com/tokene/polygonid-issuer-integration/resources"
)

func ToTransactionResource(address, tx string) resources.DeployVerifierResponse {
	return resources.DeployVerifierResponse{
		Data: resources.DeployVerifier{
			Key: resources.Key{
				ID:   "",
				Type: resources.TRANSACTION,
			},
			Attributes: resources.DeployVerifierAttributes{
				Address: address,
				Tx:      tx,
			},
		},
	}
}
