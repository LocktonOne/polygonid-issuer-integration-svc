package converter

import (
	"gitlab.com/tokene/polygonid-issuer-integration/resources"
)

func ToTransactionResource(id, address, tx string) resources.DeployVerifierResponse {
	return resources.DeployVerifierResponse{
		Data: resources.DeployVerifier{
			Key: resources.Key{
				ID:   id,
				Type: resources.TRANSACTION,
			},
			Attributes: resources.DeployVerifierAttributes{
				Address: address,
				Tx:      tx,
			},
		},
	}
}
