/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package resources

type CreateIdentityRequest struct {
	Key
	Attributes CreateIdentityRequestAttributes `json:"attributes"`
}
type CreateIdentityRequestRequest struct {
	Data     CreateIdentityRequest `json:"data"`
	Included Included              `json:"included"`
}

type CreateIdentityRequestListRequest struct {
	Data     []CreateIdentityRequest `json:"data"`
	Included Included                `json:"included"`
	Links    *Links                  `json:"links"`
}

// MustCreateIdentityRequest - returns CreateIdentityRequest from include collection.
// if entry with specified key does not exist - returns nil
// if entry with specified key exists but type or ID mismatches - panics
func (c *Included) MustCreateIdentityRequest(key Key) *CreateIdentityRequest {
	var createIdentityRequest CreateIdentityRequest
	if c.tryFindEntry(key, &createIdentityRequest) {
		return &createIdentityRequest
	}
	return nil
}
