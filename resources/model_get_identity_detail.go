/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package resources

type GetIdentityDetail struct {
	Key
	Attributes GetIdentityDetailAttributes `json:"attributes"`
}
type GetIdentityDetailResponse struct {
	Data     GetIdentityDetail `json:"data"`
	Included Included          `json:"included"`
}

type GetIdentityDetailListResponse struct {
	Data     []GetIdentityDetail `json:"data"`
	Included Included            `json:"included"`
	Links    *Links              `json:"links"`
}

// MustGetIdentityDetail - returns GetIdentityDetail from include collection.
// if entry with specified key does not exist - returns nil
// if entry with specified key exists but type or ID mismatches - panics
func (c *Included) MustGetIdentityDetail(key Key) *GetIdentityDetail {
	var getIdentityDetail GetIdentityDetail
	if c.tryFindEntry(key, &getIdentityDetail) {
		return &getIdentityDetail
	}
	return nil
}
