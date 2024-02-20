/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package resources

type RevocationStatus struct {
	Key
	Attributes RevocationStatusAttributes `json:"attributes"`
}
type RevocationStatusResponse struct {
	Data     RevocationStatus `json:"data"`
	Included Included         `json:"included"`
}

type RevocationStatusListResponse struct {
	Data     []RevocationStatus `json:"data"`
	Included Included           `json:"included"`
	Links    *Links             `json:"links"`
}

// MustRevocationStatus - returns RevocationStatus from include collection.
// if entry with specified key does not exist - returns nil
// if entry with specified key exists but type or ID mismatches - panics
func (c *Included) MustRevocationStatus(key Key) *RevocationStatus {
	var revocationStatus RevocationStatus
	if c.tryFindEntry(key, &revocationStatus) {
		return &revocationStatus
	}
	return nil
}
