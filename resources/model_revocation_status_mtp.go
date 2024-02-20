/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package resources

type RevocationStatusMtp struct {
	Existence bool                        `json:"existence"`
	NodeAux   *RevocationStatusMtpNodeAux `json:"node_aux,omitempty"`
	Siblings  *[]string                   `json:"siblings,omitempty"`
}
