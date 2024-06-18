package data

type DidQ interface {
	New() DidQ

	Get() (*Did, error)
	Select() ([]Did, error)
	Insert(value Did) error
	Update(value Did) (*Did, error)
	Delete() error

	FilterByDid(dids ...string) DidQ
	FilterByAddress(addresses ...string) DidQ
}

type Did struct {
	DID     string `db:"did" structs:"did"`
	Address string `db:"address" structs:"address"`
}
