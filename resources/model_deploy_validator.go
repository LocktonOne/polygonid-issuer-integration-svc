/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package resources

type DeployValidator struct {
	Key
	Attributes DeployValidatorAttributes `json:"attributes"`
}
type DeployValidatorRequest struct {
	Data     DeployValidator `json:"data"`
	Included Included        `json:"included"`
}

type DeployValidatorListRequest struct {
	Data     []DeployValidator `json:"data"`
	Included Included          `json:"included"`
	Links    *Links            `json:"links"`
}

// MustDeployValidator - returns DeployValidator from include collection.
// if entry with specified key does not exist - returns nil
// if entry with specified key exists but type or ID mismatches - panics
func (c *Included) MustDeployValidator(key Key) *DeployValidator {
	var deployValidator DeployValidator
	if c.tryFindEntry(key, &deployValidator) {
		return &deployValidator
	}
	return nil
}
