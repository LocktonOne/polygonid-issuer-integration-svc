/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package resources

type DeployVerifierRequest struct {
	Key
	Attributes DeployVerifierRequestAttributes `json:"attributes"`
}
type DeployVerifierRequestRequest struct {
	Data     DeployVerifierRequest `json:"data"`
	Included Included              `json:"included"`
}

type DeployVerifierRequestListRequest struct {
	Data     []DeployVerifierRequest `json:"data"`
	Included Included                `json:"included"`
	Links    *Links                  `json:"links"`
}

// MustDeployVerifierRequest - returns DeployVerifierRequest from include collection.
// if entry with specified key does not exist - returns nil
// if entry with specified key exists but type or ID mismatches - panics
func (c *Included) MustDeployVerifierRequest(key Key) *DeployVerifierRequest {
	var deployVerifierRequest DeployVerifierRequest
	if c.tryFindEntry(key, &deployVerifierRequest) {
		return &deployVerifierRequest
	}
	return nil
}
