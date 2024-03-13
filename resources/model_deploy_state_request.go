/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package resources

type DeployStateRequest struct {
	Key
	Attributes DeployStateRequestAttributes `json:"attributes"`
}
type DeployStateRequestRequest struct {
	Data     DeployStateRequest `json:"data"`
	Included Included           `json:"included"`
}

type DeployStateRequestListRequest struct {
	Data     []DeployStateRequest `json:"data"`
	Included Included             `json:"included"`
	Links    *Links               `json:"links"`
}

// MustDeployStateRequest - returns DeployStateRequest from include collection.
// if entry with specified key does not exist - returns nil
// if entry with specified key exists but type or ID mismatches - panics
func (c *Included) MustDeployStateRequest(key Key) *DeployStateRequest {
	var deployStateRequest DeployStateRequest
	if c.tryFindEntry(key, &deployStateRequest) {
		return &deployStateRequest
	}
	return nil
}
