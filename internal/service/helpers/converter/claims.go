package converter

import (
	"gitlab.com/tokene/polygonid-issuer-integration/internal/service/types"
	"gitlab.com/tokene/polygonid-issuer-integration/resources"
)

func ToRevokeCredentialResource(message string) resources.RevokeCredentialResponse {
	return resources.RevokeCredentialResponse{
		Data: resources.RevokeCredential{
			Key: resources.Key{
				ID:   "",
				Type: resources.REVOKE_CREDENTIAL,
			},
			Attributes: resources.RevokeCredentialAttributes{Message: message},
		},
	}
}

func ToRevocationStatusResource(obj types.RevocationStatusResponse) (resp resources.RevocationStatusResponse) {
	resp = resources.RevocationStatusResponse{
		Data: resources.RevocationStatus{
			Key: resources.Key{
				ID:   "",
				Type: resources.REVOCATION_STATUS,
			},
			Attributes: resources.RevocationStatusAttributes{
				Issuer: resources.RevocationStatusIssuer{
					ClaimsTreeRoot:     obj.Issuer.ClaimsTreeRoot,
					RevocationTreeRoot: obj.Issuer.RevocationTreeRoot,
					RootOfRoots:        obj.Issuer.RootOfRoots,
					State:              obj.Issuer.State,
				},
				Mtp: resources.RevocationStatusMtp{
					Existence: obj.Mtp.Existence,
					Siblings:  obj.Mtp.Siblings,
				},
			},
		},
	}
	if obj.Mtp.NodeAux != nil {
		resp.Data.Attributes.Mtp.NodeAux = &resources.RevocationStatusMtpNodeAux{
			NodeAuxKey: obj.Mtp.NodeAux.Key,
			Value:      obj.Mtp.NodeAux.Value,
		}
	}
	return resp
}

func ToQrCodeResource(obj types.GetClaimQrCodeResponse) resources.CredentialQrCodeResponse {
	var credentials []resources.CredentialQrCodeBodyCredentials
	for _, cred := range obj.Body.Credentials {
		credentials = append(credentials, resources.CredentialQrCodeBodyCredentials{
			Id:          cred.Id,
			Description: cred.Description,
		})
	}

	return resources.CredentialQrCodeResponse{
		Data: resources.CredentialQrCode{
			Key: resources.Key{
				ID:   obj.Id,
				Type: resources.QR_CODE,
			},
			Attributes: resources.CredentialQrCodeAttributes{
				Body: resources.CredentialQrCodeBody{
					Credentials: credentials,
					Url:         obj.Body.Url,
				},
				From: obj.From,
				Thid: obj.Thid,
				To:   obj.To,
				Typ:  obj.Typ,
				Type: obj.Type,
			},
		},
		Included: resources.Included{},
	}
}

func ToGetCredentialResource(obj types.GetClaimResponse) resources.GetCredentialResponse {
	resp := resources.GetCredentialResponse{
		Data: resources.GetCredential{
			Key: resources.Key{
				ID:   obj.Id,
				Type: resources.CREATE_IDENTITY,
			},
			Attributes: resources.GetCredentialAttributes{
				Context: obj.Context,
				CredentialSchema: resources.CredentialSchema{
					Id:   obj.CredentialSchema.Id,
					Type: obj.CredentialSchema.Type,
				},
				CredentialStatus:  obj.CredentialStatus,
				CredentialSubject: obj.CredentialSubject,
				ExpirationDate:    obj.ExpirationDate,
				IssuanceDate:      obj.IssuanceDate,
				Issuer:            obj.Issuer,
				Proof:             obj.Proof,
				Type:              obj.Type,
			},
		},
	}
	if obj.RefreshService != nil {
		resp.Data.Attributes.RefreshService = &resources.RefreshService{
			Id:   obj.RefreshService.Id,
			Type: string(obj.RefreshService.Type),
		}
	}
	if obj.DisplayMethod != nil {
		resp.Data.Attributes.DisplayMethod = &resources.DisplayMethod{
			Id:   obj.DisplayMethod.Id,
			Type: string(obj.DisplayMethod.Type),
		}
	}
	return resp
}

func ToGetCredentialListResource(list types.GetClaimsResponse) resources.GetCredentialListResponse {
	var resp resources.GetCredentialListResponse
	for _, obj := range list {
		el := ToGetCredentialResource(obj)
		resp.Data = append(resp.Data, el.Data)
	}
	return resp
}
