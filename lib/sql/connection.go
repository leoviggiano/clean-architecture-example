package sql

import (
	"context"
	"database/sql"
	"io"
	"time"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq" // Postgres dependency
	"github.com/sirupsen/logrus"
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

type Settings struct {
	Conn            string
	MaxIdleCons     int
	MaxOpenCons     int
	ConnMaxLifetime time.Duration
}

func NewConnection(settings Settings, logger logrus.FieldLogger) (Connection, error) {
	db, err := sqlx.Connect("postgres", settings.Conn)
	if err != nil {
		return nil, err
	}

	db.SetMaxIdleConns(settings.MaxIdleCons)
	db.SetMaxOpenConns(settings.MaxOpenCons)
	db.SetConnMaxLifetime(settings.ConnMaxLifetime)

	return Wrapper{db, logger}, err
}
