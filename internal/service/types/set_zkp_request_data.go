package types

import "math/big"

type CredentialAtomicQuery struct {
	Schema                   *big.Int   `json:"schema"`
	ClaimPathKey             *big.Int   `json:"claimPathKey"`
	Operator                 *big.Int   `json:"operator"`
	SlotIndex                *big.Int   `json:"slotIndex"`
	Value                    []*big.Int `json:"value"`
	QueryHash                *big.Int   `json:"queryHash"`
	AllowedIssuers           []*big.Int `json:"allowedIssuers"`
	CircuitIds               []string
	SkipClaimRevocationCheck bool     `json:"skipClaimRevocationCheck"`
	ClaimPathNotExists       *big.Int `json:"claimPathNotExists"`
}
