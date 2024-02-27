/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package resources

type SetZkpRequestAttributes struct {
	AllowedIssuers           *[]string `json:"allowed_issuers,omitempty"`
	CircuitId                string    `json:"circuit_id"`
	Claim                    string    `json:"claim"`
	ClaimPathNotExists       bool      `json:"claim_path_not_exists"`
	FieldName                string    `json:"field_name"`
	Operator                 string    `json:"operator"`
	Reason                   string    `json:"reason"`
	RequestId                int64     `json:"request_id"`
	SchemaType               string    `json:"schema_type"`
	SchemaUrl                string    `json:"schema_url"`
	SkipClaimRevocationCheck bool      `json:"skip_claim_revocation_check"`
	SlotIndex                int64     `json:"slot_index"`
	ValidatorAddress         *string   `json:"validator_address,omitempty"`
	Value                    []int64   `json:"value"`
	VerifierAddress          string    `json:"verifier_address"`
}
