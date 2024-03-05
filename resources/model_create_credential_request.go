/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package resources

type CreateCredentialRequest struct {
	Key
	Attributes CreateCredentialRequestAttributes `json:"attributes"`
}
type CreateCredentialRequestRequest struct {
	Data     CreateCredentialRequest `json:"data"`
	Included Included                `json:"included"`
}

type CreateCredentialRequestListRequest struct {
	Data     []CreateCredentialRequest `json:"data"`
	Included Included                  `json:"included"`
	Links    *Links                    `json:"links"`
}

// MustCreateCredentialRequest - returns CreateCredentialRequest from include collection.
// if entry with specified key does not exist - returns nil
// if entry with specified key exists but type or ID mismatches - panics
func (c *Included) MustCreateCredentialRequest(key Key) *CreateCredentialRequest {
	var createCredentialRequest CreateCredentialRequest
	if c.tryFindEntry(key, &createCredentialRequest) {
		return &createCredentialRequest
	}
	return nil
}
