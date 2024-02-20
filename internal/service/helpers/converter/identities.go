package converter

import (
	"gitlab.com/tokene/polygonid-issuer-integration/internal/service/types"
	"gitlab.com/tokene/polygonid-issuer-integration/resources"
)

func ToListIdentitiesResource(arr []string) resources.RelationCollection {
	var resp resources.RelationCollection
	for _, did := range arr {
		resp.Data = append(resp.Data, resources.Key{
			ID:   did,
			Type: resources.IDENTITY,
		})
	}
	return resp
}

func ToPublishIdentityResource(obj types.PublishIdentityStateResponse) resources.PublishStateResponse {
	return resources.PublishStateResponse{
		Data: resources.PublishState{
			Key: resources.Key{
				ID:   "",
				Type: resources.PUBLISH_STATE,
			},
			Attributes: resources.PublishStateAttributes{
				ClaimsTreeRoot:     obj.ClaimsTreeRoot,
				RevocationTreeRoot: obj.RevocationTreeRoot,
				RootOfRoots:        obj.RootOfRoots,
				State:              obj.State,
				TxID:               obj.TxID,
			},
		},
	}
}

func ToGetIdentityDetailResource(obj types.GetIdentityDetailsResponse) resources.GetIdentityDetailResponse {
	return resources.GetIdentityDetailResponse{
		Data: resources.GetIdentityDetail{
			Key: resources.Key{
				ID:   *obj.Identifier,
				Type: resources.IDENTITY,
			},
			Attributes: resources.GetIdentityDetailAttributes{
				Address: obj.Address,
				Balance: obj.Balance,
				State: resources.IdentityState{
					StateID:            obj.State.StateID,
					BlockNumber:        obj.State.BlockNumber,
					BlockTimestamp:     obj.State.BlockTimestamp,
					ClaimsTreeRoot:     obj.State.ClaimsTreeRoot,
					CreatedAt:          obj.State.CreatedAt,
					Identifier:         obj.State.Identifier,
					ModifiedAt:         obj.State.ModifiedAt,
					PreviousState:      obj.State.PreviousState,
					RevocationTreeRoot: obj.State.RevocationTreeRoot,
					RootOfRoots:        obj.State.RootOfRoots,
					State:              obj.State.State,
					Status:             obj.State.Status,
					TxID:               obj.State.TxID,
				},
			},
		},
	}
}

func ToCreateIdentityResource(obj types.CreateIdentityResponse) resources.CreateIdentityResponse {
	return resources.CreateIdentityResponse{
		Data: resources.CreateIdentity{
			Key: resources.Key{
				ID:   *obj.Identifier,
				Type: resources.CREATE_IDENTITY,
			},
			Attributes: resources.CreateIdentityAttributes{
				Address: obj.Address,
				State: resources.IdentityState{
					StateID:            obj.State.StateID,
					BlockNumber:        obj.State.BlockNumber,
					BlockTimestamp:     obj.State.BlockTimestamp,
					ClaimsTreeRoot:     obj.State.ClaimsTreeRoot,
					CreatedAt:          obj.State.CreatedAt,
					Identifier:         obj.State.Identifier,
					ModifiedAt:         obj.State.ModifiedAt,
					PreviousState:      obj.State.PreviousState,
					RevocationTreeRoot: obj.State.RevocationTreeRoot,
					RootOfRoots:        obj.State.RootOfRoots,
					State:              obj.State.State,
					Status:             obj.State.Status,
					TxID:               obj.State.TxID,
				},
			},
		},
	}
}
