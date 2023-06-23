package service

import (
	"context"
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"

	"clean/entity"
	mocks "clean/testdata/mocks"
	repoMocks "clean/testdata/mocks/repository"
)

func TestNewUser(t *testing.T) {
	require.NotNil(t, NewUser(nil, nil, nil))
}

func TestUser_Create(t *testing.T) {
	ctx := context.Background()
	ctrl := gomock.NewController(t)

	userMock := repoMocks.NewMockUser(ctrl)
	loggerMock := mocks.NewMockLogger(ctrl)

	service := user{
		userRepository: userMock,
		log:            loggerMock,
	}

	t.Run("success", func(t *testing.T) {
		user := &entity.User{
			Name: "Léo",
		}

		userMock.EXPECT().Create(ctx, user)
		loggerMock.EXPECT().Infof("user created: %s", user.Name)

		err := service.Create(ctx, user)
		require.NoError(t, err)
	})

	t.Run("error", func(t *testing.T) {
		user := &entity.User{
			Name: "Léo",
		}

		expectedErr := errors.New("deu ruim")

		userMock.EXPECT().Create(ctx, user).
			Return(expectedErr)

		loggerMock.EXPECT().Errorf("error when creating user:", expectedErr.Error())

		err := service.Create(ctx, user)
		require.Error(t, err)
	})
}
