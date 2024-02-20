/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package resources

import "time"

type GetCredentialAttributes struct {
	Context           []string               `json:"@context"`
	CredentialSchema  CredentialSchema       `json:"credentialSchema"`
	CredentialStatus  interface{}            `json:"credentialStatus"`
	CredentialSubject map[string]interface{} `json:"credentialSubject"`
	DisplayMethod     *DisplayMethod         `json:"displayMethod,omitempty"`
	ExpirationDate    *time.Time             `json:"expirationDate,omitempty"`
	IssuanceDate      *time.Time             `json:"issuanceDate,omitempty"`
	Issuer            string                 `json:"issuer"`
	Proof             interface{}            `json:"proof"`
	RefreshService    *RefreshService        `json:"refreshService,omitempty"`
	Type              []string               `json:"type"`
}
