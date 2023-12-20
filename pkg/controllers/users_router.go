package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/FumingPower3925/SimpleImageServer/pkg/api/responses"
	"github.com/FumingPower3925/SimpleImageServer/pkg/models"
	"github.com/FumingPower3925/SimpleImageServer/pkg/services"
	"github.com/go-chi/chi/v5"
)

type UserRouter struct {
	service *services.UserService
}

func NewUserRouter(service *services.UserService) *UserRouter {
	return &UserRouter{service: service}
}

func (ur *UserRouter) CreateHandler(w http.ResponseWriter, r *http.Request) {
	var user models.User

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		responses.HTTPError(w, r, http.StatusBadRequest, err.Error())
		return
	}

	defer r.Body.Close()

	ctx := r.Context()

	err = ur.service.CreateService(ctx, &user)
	if err != nil {
		responses.HTTPError(w, r, http.StatusBadRequest, err.Error())
		return
	}

	w.Header().Add("Location", fmt.Sprintf("%s%s", r.URL.String(), user.Id))
	responses.JSON(w, r, http.StatusCreated, responses.Map{"user": user})
}

func (ur *UserRouter) CheckCredentialsHandler(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	password := chi.URLParam(r, "password")
	ctx := r.Context()

	valid, err := ur.service.CheckCredentialsService(ctx, id, password)
	if err != nil {
		responses.HTTPError(w, r, http.StatusNotFound, err.Error())
		return
	}

	if !valid {
		responses.JSON(w, r, http.StatusForbidden, responses.Map{})
	}

	responses.JSON(w, r, http.StatusOK, responses.Map{"valid": valid})
}

func (ur *UserRouter) ExistsHandler(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	ctx := r.Context()
	exists, err := ur.service.ExistsService(ctx, id)
	if err != nil {
		responses.HTTPError(w, r, http.StatusNotFound, err.Error())
		return
	}

	responses.JSON(w, r, http.StatusOK, exists)
}

func (ur *UserRouter) DeleteHandler(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	ctx := r.Context()
	err := ur.service.DeleteService(ctx, id)
	if err != nil {
		responses.HTTPError(w, r, http.StatusNotFound, err.Error())
		return
	}

	responses.JSON(w, r, http.StatusNoContent, responses.Map{})
}

func (ur *UserRouter) Routes() http.Handler {
	r := chi.NewRouter()

	r.Post("/id", ur.CreateHandler)

	r.Get("/{id}{password}", ur.CheckCredentialsHandler)

	r.Get("/{id}", ur.ExistsHandler)

	r.Delete("/{id}", ur.DeleteHandler)

	return r
}
