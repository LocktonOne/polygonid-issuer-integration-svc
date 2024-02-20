/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package resources

type GetCredential struct {
	Key
	Attributes GetCredentialAttributes `json:"attributes"`
}
type GetCredentialResponse struct {
	Data     GetCredential `json:"data"`
	Included Included      `json:"included"`
}

type GetCredentialListResponse struct {
	Data     []GetCredential `json:"data"`
	Included Included        `json:"included"`
	Links    *Links          `json:"links"`
}

// MustGetCredential - returns GetCredential from include collection.
// if entry with specified key does not exist - returns nil
// if entry with specified key exists but type or ID mismatches - panics
func (c *Included) MustGetCredential(key Key) *GetCredential {
	var getCredential GetCredential
	if c.tryFindEntry(key, &getCredential) {
		return &getCredential
	}
	return nil
}
