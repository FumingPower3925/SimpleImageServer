package controllers

import (
	"net/http"

	"github.com/FumingPower3925/SimpleImageServer/pkg/repositories/image"
	"github.com/FumingPower3925/SimpleImageServer/pkg/repositories/user"
	"github.com/FumingPower3925/SimpleImageServer/pkg/services"
	"github.com/go-chi/chi/v5"
)

func New() (http.Handler, error) {
	r := chi.NewRouter()
	repo, err := image.GetImageRepository()
	if err != nil {
		return nil, err
	}

	ir := NewImageRouter(services.NewImageService(repo))

	r.Mount("/images", ir.Routes())

	repo2, err2 := user.GetUserRepository()
	if err2 != nil {
		return nil, err2
	}

	ur := NewUserRouter(services.NewUserService(repo2))

	r.Mount("/users", ur.Routes())

	return r, nil
}
