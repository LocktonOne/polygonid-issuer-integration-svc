package types

import "time"

const (
	BasicAuthScopes = "basicAuth.Scopes"
)

// Defines values for CreateIdentityRequestDidMetadataType.
const (
	BJJ CreateIdentityRequestDidMetadataType = "BJJ"
	ETH CreateIdentityRequestDidMetadataType = "ETH"
)

// Defines values for DisplayMethodType.
const (
	Iden3BasicDisplayMethodV1 DisplayMethodType = "Iden3BasicDisplayMethodV1"
)

// Defines values for RefreshServiceType.
const (
	Iden3RefreshService2023 RefreshServiceType = "Iden3RefreshService2023"
)

// AgentResponse defines model for AgentResponse.
type AgentResponse struct {
	Body     interface{} `json:"body"`
	From     string      `json:"from"`
	Id       string      `json:"id"`
	ThreadID string      `json:"threadID"`
	To       string      `json:"to"`
	Typ      string      `json:"typ"`
	Type     string      `json:"type"`
}

// Config defines model for Config.
type Config = []KeyValue

// CreateClaimRequest defines model for CreateClaimRequest.
type CreateClaimRequest struct {
	CredentialSchema      string                 `json:"credentialSchema"`
	CredentialSubject     map[string]interface{} `json:"credentialSubject"`
	DisplayMethod         *DisplayMethod         `json:"displayMethod,omitempty"`
	Expiration            *int64                 `json:"expiration,omitempty"`
	MerklizedRootPosition *string                `json:"merklizedRootPosition,omitempty"`
	RefreshService        *RefreshService        `json:"refreshService,omitempty"`
	RevNonce              *uint64                `json:"revNonce,omitempty"`
	SubjectPosition       *string                `json:"subjectPosition,omitempty"`
	Type                  string                 `json:"type"`
	Version               *uint32                `json:"version,omitempty"`
}

// CreateClaimResponse defines model for CreateClaimResponse.
type CreateClaimResponse struct {
	Id string `json:"id"`
}

// CreateIdentityRequest defines model for CreateIdentityRequest.
type CreateIdentityRequest struct {
	DidMetadata struct {
		Blockchain string                               `json:"blockchain"`
		Method     string                               `json:"method"`
		Network    string                               `json:"network"`
		Type       CreateIdentityRequestDidMetadataType `json:"type"`
	} `json:"didMetadata"`
}

// CreateIdentityRequestDidMetadataType defines model for CreateIdentityRequest.DidMetadata.Type.
type CreateIdentityRequestDidMetadataType string

// CreateIdentityResponse defines model for CreateIdentityResponse.
type CreateIdentityResponse struct {
	Address    *string        `json:"address"`
	Identifier *string        `json:"identifier,omitempty"`
	State      *IdentityState `json:"state,omitempty"`
}

// CredentialSchema defines model for CredentialSchema.
type CredentialSchema struct {
	Id   string `json:"id"`
	Type string `json:"type"`
}

// DisplayMethod defines model for DisplayMethod.
type DisplayMethod struct {
	Id   string            `json:"id"`
	Type DisplayMethodType `json:"type"`
}

// DisplayMethodType defines model for DisplayMethod.Type.
type DisplayMethodType string

// GenericErrorMessage defines model for GenericErrorMessage.
type GenericErrorMessage struct {
	Message string `json:"message"`
}

// GetClaimQrCodeResponse defines model for GetClaimQrCodeResponse.
type GetClaimQrCodeResponse struct {
	Body struct {
		Credentials []struct {
			Description string `json:"description"`
			Id          string `json:"id"`
		} `json:"credentials"`
		Url string `json:"url"`
	} `json:"body"`
	From string `json:"from"`
	Id   string `json:"id"`
	Thid string `json:"thid"`
	To   string `json:"to"`
	Typ  string `json:"typ"`
	Type string `json:"type"`
}

// GetClaimResponse defines model for GetClaimResponse.
type GetClaimResponse struct {
	Context           []string               `json:"@context"`
	CredentialSchema  CredentialSchema       `json:"credentialSchema"`
	CredentialStatus  interface{}            `json:"credentialStatus"`
	CredentialSubject map[string]interface{} `json:"credentialSubject"`
	DisplayMethod     *DisplayMethod         `json:"displayMethod,omitempty"`
	ExpirationDate    *time.Time             `json:"expirationDate"`
	Id                string                 `json:"id"`
	IssuanceDate      *time.Time             `json:"issuanceDate"`
	Issuer            string                 `json:"issuer"`
	Proof             interface{}            `json:"proof"`
	RefreshService    *RefreshService        `json:"refreshService,omitempty"`
	Type              []string               `json:"type"`
}

// GetClaimsResponse defines model for GetClaimsResponse.
type GetClaimsResponse = []GetClaimResponse

// GetIdentityDetailsResponse defines model for GetIdentityDetailsResponse.
type GetIdentityDetailsResponse struct {
	Address    *string        `json:"address,omitempty"`
	Balance    *string        `json:"balance,omitempty"`
	Identifier *string        `json:"identifier,omitempty"`
	State      *IdentityState `json:"state,omitempty"`
}

// Health defines model for Health.
type Health map[string]bool

// IdentityState defines model for IdentityState.
type IdentityState struct {
	BlockNumber        *int      `json:"blockNumber,omitempty"`
	BlockTimestamp     *int      `json:"blockTimestamp,omitempty"`
	ClaimsTreeRoot     *string   `json:"claimsTreeRoot,omitempty"`
	CreatedAt          time.Time `json:"createdAt"`
	Identifier         string    `json:"-"`
	ModifiedAt         time.Time `json:"modifiedAt"`
	PreviousState      *string   `json:"previousState,omitempty"`
	RevocationTreeRoot *string   `json:"revocationTreeRoot,omitempty"`
	RootOfRoots        *string   `json:"rootOfRoots,omitempty"`
	State              *string   `json:"state,omitempty"`
	StateID            int64     `json:"-"`
	Status             string    `json:"status"`
	TxID               *string   `json:"txID,omitempty"`
}

// KeyValue defines model for KeyValue.
type KeyValue struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

// PublishIdentityStateResponse defines model for PublishIdentityStateResponse.
type PublishIdentityStateResponse struct {
	ClaimsTreeRoot     *string `json:"claimsTreeRoot,omitempty"`
	RevocationTreeRoot *string `json:"revocationTreeRoot,omitempty"`
	RootOfRoots        *string `json:"rootOfRoots,omitempty"`
	State              *string `json:"state,omitempty"`
	TxID               *string `json:"txID,omitempty"`
}

// RefreshService defines model for RefreshService.
type RefreshService struct {
	Id   string             `json:"id"`
	Type RefreshServiceType `json:"type"`
}

// RefreshServiceType defines model for RefreshService.Type.
type RefreshServiceType string

// RevocationStatusResponse defines model for RevocationStatusResponse.
type RevocationStatusResponse struct {
	Issuer struct {
		ClaimsTreeRoot     *string `json:"claimsTreeRoot,omitempty"`
		RevocationTreeRoot *string `json:"revocationTreeRoot,omitempty"`
		RootOfRoots        *string `json:"rootOfRoots,omitempty"`
		State              *string `json:"state,omitempty"`
	} `json:"issuer"`
	Mtp struct {
		Existence bool `json:"existence"`
		NodeAux   *struct {
			Key   *string `json:"key,omitempty"`
			Value *string `json:"value,omitempty"`
		} `json:"node_aux,omitempty"`
		Siblings *[]string `json:"siblings"`
	} `json:"mtp"`
}

// RevokeClaimResponse defines model for RevokeClaimResponse.
type RevokeClaimResponse struct {
	Message string `json:"message"`
}

// PathClaim defines model for pathClaim.
type PathClaim = string

// PathIdentifier defines model for pathIdentifier.
type PathIdentifier = string

// PathNonce defines model for pathNonce.
type PathNonce = int64

// N400 defines model for 400.
type N400 = GenericErrorMessage

// N401 defines model for 401.
type N401 = GenericErrorMessage

// N403 defines model for 403.
type N403 = GenericErrorMessage

// N404 defines model for 404.
type N404 = GenericErrorMessage

// N409 defines model for 409.
type N409 = GenericErrorMessage

// N422 defines model for 422.
type N422 = GenericErrorMessage

// N500 defines model for 500.
type N500 = GenericErrorMessage

// N500CreateIdentity defines model for 500-CreateIdentity.
type N500CreateIdentity struct {
	Code      *int    `json:"code,omitempty"`
	Error     *string `json:"error,omitempty"`
	RequestID *string `json:"requestID,omitempty"`
}

// AgentTextBody defines parameters for Agent.
type AgentTextBody = string

// GetClaimsParams defines parameters for GetClaims.
type GetClaimsParams struct {
	// SchemaType Filter per schema type. Example - KYCAgeCredential
	SchemaType *string `form:"schemaType,omitempty" json:"schemaType,omitempty"`

	// SchemaHash Filter per schema hash. Example - c9b2370371b7fa8b3dab2a5ba81b6838
	SchemaHash *string `form:"schemaHash,omitempty" json:"schemaHash,omitempty"`

	// Subject Filter per subject. Example - did:polygonid:polygon:mumbai:2qE1BZ7gcmEoP2KppvFPCZqyzyb5tK9T6Gec5HFANQ
	Subject *string `form:"subject,omitempty" json:"subject,omitempty"`

	// Revoked Filter per claims revoked or not - Example - true.
	Revoked *bool `form:"revoked,omitempty" json:"revoked,omitempty"`

	// Self Filter per retrieve claims of the provided identifier. Example - true
	Self *bool `form:"self,omitempty" json:"self,omitempty"`

	// QueryField Filter this field inside the data of the claim
	QueryField *string `form:"query_field,omitempty" json:"query_field,omitempty"`

	// QueryValue Filter this value inside the data of the claim for the specified field in query_field
	QueryValue *string `form:"query_value,omitempty" json:"query_value,omitempty"`
}

// AgentTextRequestBody defines body for Agent for text/plain ContentType.
type AgentTextRequestBody = AgentTextBody

// CreateIdentityJSONRequestBody defines body for CreateIdentity for application/json ContentType.
type CreateIdentityJSONRequestBody = CreateIdentityRequest

// CreateClaimJSONRequestBody defines body for CreateClaim for application/json ContentType.
type CreateClaimJSONRequestBody = CreateClaimRequest
