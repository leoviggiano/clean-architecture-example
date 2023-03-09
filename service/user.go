package service

import (
	"context"
	"errors"

	"clean/entity"
	"clean/lib/cache"
	"clean/lib/sql"
	"clean/service/internal/repository"
)

var (
	ErrUserNotFound = errors.New("user not found")
)

type User interface {
	Create(ctx context.Context, user *entity.User) error
	Get(ctx context.Context, id int) (*entity.User, error)
	Update(ctx context.Context, user *entity.User) error
	GiveExp(ctx context.Context, user *entity.User, exp int) error
	Delete(ctx context.Context, id int) error
}

type user struct {
	userRepository repository.User
}

func NewUser(db sql.Connection, cache cache.Cache) User {
	return user{
		userRepository: repository.NewUser(db, cache),
	}
}

func (u user) Create(ctx context.Context, user *entity.User) error {
	return u.userRepository.Create(ctx, user)
}

func (u user) Get(ctx context.Context, id int) (*entity.User, error) {
	return u.userRepository.Get(ctx, id)
}

func (u user) Update(ctx context.Context, user *entity.User) error {
	return u.userRepository.Update(ctx, user)
}

func (u user) Delete(ctx context.Context, id int) error {
	return u.userRepository.Delete(ctx, id)
}

func (u user) GiveExp(ctx context.Context, user *entity.User, exp int) error {
	if user.Exp+exp > user.NextLevelExp {
		user.Exp = 0
		user.Level += 1
	} else {
		user.Exp += exp
	}

	return u.Update(ctx, user)
}
