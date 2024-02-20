/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package resources

import "time"

type IssuerIdentityState struct {
	BlockNumber        *int32    `json:"blockNumber,omitempty"`
	BlockTimestamp     *int32    `json:"blockTimestamp,omitempty"`
	ClaimsTreeRoot     *string   `json:"claimsTreeRoot,omitempty"`
	CreatedAt          time.Time `json:"createdAt"`
	Identifier         string    `json:"identifier"`
	ModifiedAt         time.Time `json:"modifiedAt"`
	PreviousState      *string   `json:"previousState,omitempty"`
	RevocationTreeRoot *string   `json:"revocationTreeRoot,omitempty"`
	RootOfRoots        *string   `json:"rootOfRoots,omitempty"`
	State              *string   `json:"state,omitempty"`
	StateID            int64     `json:"stateID"`
	Status             string    `json:"status"`
	TxID               *string   `json:"txID,omitempty"`
}
