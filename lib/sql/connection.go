package sql

import (
	"context"
	"database/sql"
	"io"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq" // Postgres dependency
)

type Connection interface {
	io.Closer
	Query(ctx context.Context, query string, args ...interface{}) (*sql.Rows, error)
	Exec(ctx context.Context, query string, args ...interface{}) (sql.Result, error)
	NamedExec(ctx context.Context, query string, arg interface{}) (sql.Result, error)
	In(query string, args ...interface{}) (string, []interface{}, error)
	Get(ctx context.Context, dest interface{}, query string, args ...interface{}) error
	Select(ctx context.Context, dest interface{}, query string, args ...interface{}) error
}

func NewConnection(connectionString string) (Connection, error) {
	db, err := sqlx.Connect("postgres", connectionString)
	if err != nil {
		return nil, err
	}

	return Wrapper{db}, err
}
