package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"clean/entity"
	"clean/lib/log"
	"clean/service"
)

type Handler interface {
	CreateUser(w http.ResponseWriter, r *http.Request)
	GetUser(w http.ResponseWriter, r *http.Request)
	GetUsers(w http.ResponseWriter, r *http.Request)
	UpdateUser(w http.ResponseWriter, r *http.Request)
	DeleteUser(w http.ResponseWriter, r *http.Request)
}

type handler struct {
	services service.All
	logger   log.Logger
}

func NewHandler(services service.All, logger log.Logger) Handler {
	return handler{services, logger}
}

func (h handler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var user entity.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		h.logger.Error(err)
		w.WriteHeader(http.StatusServiceUnavailable)
		return
	}

	err = h.services.User.Create(r.Context(), &user)
	if err != nil {
		h.logger.Error(err)
		w.WriteHeader(http.StatusServiceUnavailable)
		return
	}

	h.logger.Infof("user created: %s", user.Name)
	w.WriteHeader(http.StatusCreated)
}

func (h handler) GetUser(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		h.logger.Error(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	user, err := h.services.User.Get(r.Context(), id)
	if err != nil {
		h.logger.Error(err)
		w.WriteHeader(http.StatusServiceUnavailable)
		return
	}

	userJSON, err := json.Marshal(user)
	if err != nil {
		h.logger.Error(err)
		w.WriteHeader(http.StatusServiceUnavailable)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(userJSON)
}

func (h handler) GetUsers(w http.ResponseWriter, r *http.Request) {
	users, err := h.services.User.GetAll(r.Context())
	if err != nil {
		h.logger.Error(err)
		w.WriteHeader(http.StatusServiceUnavailable)
		return
	}

	usersJSON, err := json.Marshal(users)
	if err != nil {
		h.logger.Error(err)
		w.WriteHeader(http.StatusServiceUnavailable)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(usersJSON)
}

func (h handler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	var user entity.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		h.logger.Error(err)
		w.WriteHeader(http.StatusServiceUnavailable)
		return
	}

	err = h.services.User.Update(r.Context(), &user)
	if err != nil {
		h.logger.Error(err)
		w.WriteHeader(http.StatusServiceUnavailable)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h handler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		h.logger.Error(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = h.services.User.Delete(r.Context(), id)
	if err != nil {
		h.logger.Error(err)
		w.WriteHeader(http.StatusServiceUnavailable)
		return
	}

	w.WriteHeader(http.StatusOK)
}
