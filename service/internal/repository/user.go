package repository

import (
	"context"
	_ "embed"
	"fmt"

	"github.com/nleof/goyesql"

	"clean/entity"
	"clean/lib/cache"
	"clean/lib/sql"
)

//go:embed user.sql
var userQueries []byte

type User interface {
	Create(ctx context.Context, user *entity.User) error
	Get(ctx context.Context, id int) (*entity.User, error)
	Update(ctx context.Context, user *entity.User) error
	Delete(ctx context.Context, id int) error
}

type user struct {
	queries goyesql.Queries
	db      sql.Connection
	cache   cache.Cache
}

func NewUser(db sql.Connection, cache cache.Cache) User {
	return user{
		queries: goyesql.MustParseBytes(userQueries),
		db:      db,
		cache:   cache,
	}
}

func (u user) Create(ctx context.Context, user *entity.User) error {
	_, err := u.db.Exec(ctx, u.queries["create"], user.ID, user.Name)
	if err != nil {
		return err
	}
	return nil
}

func (u user) Get(ctx context.Context, id int) (*entity.User, error) {
	var user entity.User

	if err := u.cache.Get(u.cacheKeyUser(id), &user); err == nil {
		return &user, nil
	}

	err := u.db.Get(ctx, &user, u.queries["get"], id)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (u user) Update(ctx context.Context, user *entity.User) error {
	_, err := u.db.NamedExec(ctx, u.queries["update"], user)
	if err != nil {
		return err
	}

	return nil
}

func (u user) Delete(ctx context.Context, id int) error {
	_, err := u.db.Exec(ctx, u.queries["delete"], id)
	if err != nil {
		return err
	}

	return nil
}

func (u user) cacheKeyUser(id int) string {
	return fmt.Sprintf("user-%d", id)
}
