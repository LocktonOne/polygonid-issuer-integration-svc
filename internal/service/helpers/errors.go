package helpers

import "errors"

type Error struct {
	Message string `json:"message"`
}

var (
	WrongDidErr = errors.New("DID is empty or doesn't begin with \"did:\"")
)
