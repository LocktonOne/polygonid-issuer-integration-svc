/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package resources

type IssuerCreateCredentialRequest struct {
	CredentialSchema  string                 `json:"credentialSchema"`
	CredentialSubject map[string]interface{} `json:"credentialSubject"`
	Expiration        *int64                 `json:"expiration,omitempty"`
	MtProof           *bool                  `json:"mtProof,omitempty"`
	SignatureProof    *bool                  `json:"signatureProof,omitempty"`
	Type              string                 `json:"type"`
}
