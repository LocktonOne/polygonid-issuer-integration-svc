/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package resources

import "time"

type CreateCredentialRequestAttributes struct {
	CredentialSchema  string                 `json:"credential_schema"`
	CredentialSubject map[string]interface{} `json:"credential_subject"`
	Expiration        *time.Duration         `json:"expiration,omitempty"`
	MtProof           *bool                  `json:"mt_proof,omitempty"`
	SignatureProof    *bool                  `json:"signature_proof,omitempty"`
	Type              string                 `json:"type"`
}
