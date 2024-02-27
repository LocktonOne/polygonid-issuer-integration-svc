/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package resources

type DeployVerifier struct {
	Key
	Attributes DeployVerifierAttributes `json:"attributes"`
}
type DeployVerifierResponse struct {
	Data     DeployVerifier `json:"data"`
	Included Included       `json:"included"`
}

type DeployVerifierListResponse struct {
	Data     []DeployVerifier `json:"data"`
	Included Included         `json:"included"`
	Links    *Links           `json:"links"`
}

// MustDeployVerifier - returns DeployVerifier from include collection.
// if entry with specified key does not exist - returns nil
// if entry with specified key exists but type or ID mismatches - panics
func (c *Included) MustDeployVerifier(key Key) *DeployVerifier {
	var deployVerifier DeployVerifier
	if c.tryFindEntry(key, &deployVerifier) {
		return &deployVerifier
	}
	return nil
}
