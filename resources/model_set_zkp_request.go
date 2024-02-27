/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package resources

type SetZkpRequest struct {
	Key
	Attributes SetZkpRequestAttributes `json:"attributes"`
}
type SetZkpRequestRequest struct {
	Data     SetZkpRequest `json:"data"`
	Included Included      `json:"included"`
}

type SetZkpRequestListRequest struct {
	Data     []SetZkpRequest `json:"data"`
	Included Included        `json:"included"`
	Links    *Links          `json:"links"`
}

// MustSetZkpRequest - returns SetZkpRequest from include collection.
// if entry with specified key does not exist - returns nil
// if entry with specified key exists but type or ID mismatches - panics
func (c *Included) MustSetZkpRequest(key Key) *SetZkpRequest {
	var setZKPRequest SetZkpRequest
	if c.tryFindEntry(key, &setZKPRequest) {
		return &setZKPRequest
	}
	return nil
}
