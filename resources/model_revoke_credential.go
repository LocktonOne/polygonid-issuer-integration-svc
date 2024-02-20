/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package resources

type RevokeCredential struct {
	Key
	Attributes RevokeCredentialAttributes `json:"attributes"`
}
type RevokeCredentialResponse struct {
	Data     RevokeCredential `json:"data"`
	Included Included         `json:"included"`
}

type RevokeCredentialListResponse struct {
	Data     []RevokeCredential `json:"data"`
	Included Included           `json:"included"`
	Links    *Links             `json:"links"`
}

// MustRevokeCredential - returns RevokeCredential from include collection.
// if entry with specified key does not exist - returns nil
// if entry with specified key exists but type or ID mismatches - panics
func (c *Included) MustRevokeCredential(key Key) *RevokeCredential {
	var revokeCredential RevokeCredential
	if c.tryFindEntry(key, &revokeCredential) {
		return &revokeCredential
	}
	return nil
}
