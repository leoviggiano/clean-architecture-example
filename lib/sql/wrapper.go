package sql

import (
	"context"
	"database/sql"
	"regexp"

	"github.com/jmoiron/sqlx"

	"clean/lib/log"
)

var (
	space     = regexp.MustCompile(`\s+`)
	ErrNoRows = sql.ErrNoRows
)

type Wrapper struct {
	DB     *sqlx.DB
	Logger log.Logger
}

func (w Wrapper) DriverName() string {
	return w.DB.DriverName()
}

func (w Wrapper) Rebind(s string) string {
	return w.DB.Rebind(s)
}

func (w Wrapper) BindNamed(s string, i interface{}) (string, []interface{}, error) {
	return w.DB.BindNamed(s, i)
}

func (w Wrapper) log(query string, args ...interface{}) {
	if w.Logger != nil {
		w.Logger.Infof("[POSTGRES] %s %v", space.ReplaceAllString(query, " "), args)
	}
}

func (w Wrapper) Close() error {
	w.log("closed")
	return w.DB.Close()
}

func (w Wrapper) SelectContext(ctx context.Context, dest interface{}, query string, args ...interface{}) error {
	w.log(query, args...)
	return w.DB.SelectContext(ctx, dest, query, args...)
}

func (w Wrapper) Select(ctx context.Context, dest interface{}, query string, args ...interface{}) error {
	w.log(query, args...)
	return w.DB.SelectContext(ctx, dest, query, args...)
}

func (w Wrapper) Query(ctx context.Context, query string, args ...interface{}) (*sql.Rows, error) {
	w.log(query, args...)
	return w.DB.QueryContext(ctx, query, args...)
}

func (w Wrapper) Exec(ctx context.Context, query string, args ...interface{}) (sql.Result, error) {
	w.log(query, args...)
	return w.DB.ExecContext(ctx, query, args...)
}

func (w Wrapper) Get(ctx context.Context, dest interface{}, query string, args ...interface{}) error {
	w.log(query, args...)
	return w.DB.GetContext(ctx, dest, query, args...)
}

func (w Wrapper) In(query string, args ...interface{}) (string, []interface{}, error) {
	return sqlx.In(query, args...)
}

func (w Wrapper) NamedExec(ctx context.Context, query string, arg interface{}) (sql.Result, error) {
	return w.DB.NamedExecContext(ctx, query, arg)
}
