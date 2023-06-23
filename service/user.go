//go:generate go run github.com/golang/mock/mockgen -package=mocks -source=$GOFILE  -destination=../testdata/mocks/user.go

package service

import (
	"context"
	"errors"

	"clean/entity"
	"clean/lib/cache"
	"clean/lib/log"
	"clean/lib/sql"
	"clean/service/internal/repository"
)

var (
	ErrUserNotFound = errors.New("user not found")
)

type User interface {
	Create(ctx context.Context, user *entity.User) error
	Get(ctx context.Context, id int) (*entity.User, error)
	GetAll(ctx context.Context) ([]*entity.User, error)
	Update(ctx context.Context, user *entity.User) error
	GiveExp(ctx context.Context, user *entity.User, exp int) error
	Delete(ctx context.Context, id int) error
}

type user struct {
	userRepository repository.User
	log            log.Logger
}

func NewUser(db sql.Connection, cache cache.Cache, logger log.Logger) User {
	return user{
		userRepository: repository.NewUser(db, cache),
		log:            logger,
	}
}

func (u user) Create(ctx context.Context, user *entity.User) error {
	err := u.userRepository.Create(ctx, user)
	if err != nil {
		u.log.Errorf("error when creating user:", err.Error())
		return err
	}

	u.log.Infof("user created: %s", user.Name)
	return err
}

func (u user) Get(ctx context.Context, id int) (*entity.User, error) {
	return u.userRepository.Get(ctx, id)
}

func (u user) GetAll(ctx context.Context) ([]*entity.User, error) {
	return u.userRepository.GetAll(ctx)
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
