package postgres

import (
	"database/sql"
	sq "github.com/Masterminds/squirrel"
	"github.com/fatih/structs"
	"gitlab.com/distributed_lab/kit/pgdb"
	"gitlab.com/distributed_lab/logan/v3/errors"
	"gitlab.com/tokene/polygonid-issuer-integration/internal/data"
)

const didTableName = "dids"

func NewDidQ(db *pgdb.DB) data.DidQ {
	return &didQ{
		db:  db,
		sql: sq.StatementBuilder,
	}
}

type didQ struct {
	db  *pgdb.DB
	sql sq.StatementBuilderType
}

func (q *didQ) New() data.DidQ {
	return NewDidQ(q.db)
}

func (q *didQ) Get() (*data.Did, error) {
	var result data.Did
	err := q.db.Get(&result, q.sql.Select("*").From(didTableName))
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, errors.Wrap(err, "failed to get did from db")
	}
	return &result, nil
}

func (q *didQ) Select() ([]data.Did, error) {
	var result []data.Did
	err := q.db.Select(&result, q.sql.Select("*").From(didTableName))
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, errors.Wrap(err, "failed to select dids from db")
	}
	return result, nil
}

func (q *didQ) Insert(value data.Did) error {
	clauses := structs.Map(value)

	stmt := sq.Insert(didTableName).SetMap(clauses)
	err := q.db.Exec(stmt)
	if err != nil {
		return errors.Wrap(err, "failed to insert did to db")
	}
	return nil
}

func (q *didQ) Update(value data.Did) (*data.Did, error) {
	clauses := structs.Map(value)

	var result data.Did
	stmt := q.sql.Update(didTableName).SetMap(clauses).Suffix("returning *")
	err := q.db.Get(&result, stmt)
	if err != nil {
		return nil, errors.Wrap(err, "failed to update did in db")
	}
	return &result, nil
}

func (q *didQ) Delete() error {
	err := q.db.Exec(q.sql.Delete(didTableName))
	if err != nil {
		return errors.Wrap(err, "failed to delete dids from db")
	}
	return nil
}

func (q *didQ) FilterByDid(dids ...string) data.DidQ {
	pred := sq.Eq{"did": dids}
	q.sql = q.sql.Where(pred)
	return q
}

func (q *didQ) FilterByAddress(addresses ...string) data.DidQ {
	pred := sq.Eq{"address": addresses}
	q.sql = q.sql.Where(pred)
	return q
}
