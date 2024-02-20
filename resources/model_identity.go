/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package resources

type Identity struct {
	Key
	Attributes IdentityAttributes `json:"attributes"`
}
type IdentityResponse struct {
	Data     Identity `json:"data"`
	Included Included `json:"included"`
}

type IdentityListResponse struct {
	Data     []Identity `json:"data"`
	Included Included   `json:"included"`
	Links    *Links     `json:"links"`
}

// MustIdentity - returns Identity from include collection.
// if entry with specified key does not exist - returns nil
// if entry with specified key exists but type or ID mismatches - panics
func (c *Included) MustIdentity(key Key) *Identity {
	var identity Identity
	if c.tryFindEntry(key, &identity) {
		return &identity
	}
	return nil
}
