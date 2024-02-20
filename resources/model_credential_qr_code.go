/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package resources

type CredentialQrCode struct {
	Key
	Attributes CredentialQrCodeAttributes `json:"attributes"`
}
type CredentialQrCodeResponse struct {
	Data     CredentialQrCode `json:"data"`
	Included Included         `json:"included"`
}

type CredentialQrCodeListResponse struct {
	Data     []CredentialQrCode `json:"data"`
	Included Included           `json:"included"`
	Links    *Links             `json:"links"`
}

// MustCredentialQrCode - returns CredentialQrCode from include collection.
// if entry with specified key does not exist - returns nil
// if entry with specified key exists but type or ID mismatches - panics
func (c *Included) MustCredentialQrCode(key Key) *CredentialQrCode {
	var credentialQrCode CredentialQrCode
	if c.tryFindEntry(key, &credentialQrCode) {
		return &credentialQrCode
	}
	return nil
}
