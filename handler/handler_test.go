package handler

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"

	"clean/entity"
	"clean/service"
	"clean/testdata/mocks"
)

func TestHandler_CreateUser(t *testing.T) {
	ctrl := gomock.NewController(t)

	userMock := mocks.NewMockUser(ctrl)
	loggerMock := mocks.NewMockLogger(ctrl)

	services := service.All{
		User: userMock,
	}

	handler := NewHandler(services, loggerMock)

	t.Run("success", func(t *testing.T) {
		w := httptest.NewRecorder()
		r := httptest.NewRequest("POST", "/", nil)

		user := &entity.User{Name: "Léo"}
		userJSON, err := json.Marshal(user)
		require.NoError(t, err)

		r.Body = io.NopCloser(bytes.NewBuffer(userJSON))

		userMock.EXPECT().Create(r.Context(), user)
		loggerMock.EXPECT().Infof("user created: %s", user.Name)

		handler.CreateUser(w, r)
		require.Equal(t, http.StatusCreated, w.Result().StatusCode)
	})

	t.Run("error decode body", func(t *testing.T) {
		w := httptest.NewRecorder()
		r := httptest.NewRequest("POST", "/", nil)

		r.Body = io.NopCloser(bytes.NewBuffer([]byte("invalid json")))

		loggerMock.EXPECT().Error(gomock.Any())

		handler.CreateUser(w, r)
		require.Equal(t, http.StatusServiceUnavailable, w.Result().StatusCode)
	})

	t.Run("error create user", func(t *testing.T) {
		w := httptest.NewRecorder()
		r := httptest.NewRequest("POST", "/", nil)

		r.Body = io.NopCloser(bytes.NewBuffer([]byte("invalid json")))

		user := &entity.User{Name: "Léo"}
		userJSON, err := json.Marshal(user)
		require.NoError(t, err)

		r.Body = io.NopCloser(bytes.NewBuffer(userJSON))

		expectedErr := errors.New("error creating user")
		userMock.EXPECT().Create(r.Context(), user).
			Return(expectedErr)

		loggerMock.EXPECT().Error(expectedErr)

		handler.CreateUser(w, r)
		require.Equal(t, http.StatusServiceUnavailable, w.Result().StatusCode)
	})
}
