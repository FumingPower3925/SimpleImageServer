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

type ImageRouter struct {
	service *services.ImageService
}

func NewImageRouter(service *services.ImageService) *ImageRouter {
	return &ImageRouter{service: service}
}

func (ir *ImageRouter) CreateHandler(w http.ResponseWriter, r *http.Request) {
	var image models.ControllersImage
	err := json.NewDecoder(r.Body).Decode(&image)
	if err != nil {
		responses.HTTPError(w, r, http.StatusBadRequest, err.Error())
		return
	}

	defer r.Body.Close()

	ctx := r.Context()
	err = ir.service.CreateService(ctx, &image)
	if err != nil {
		responses.HTTPError(w, r, http.StatusBadRequest, err.Error())
		return
	}

	w.Header().Add("Location", fmt.Sprintf("%s%s", r.URL.String(), image.Filename))
	responses.JSON(w, r, http.StatusCreated, responses.Map{"image": image})
}

func (ir *ImageRouter) GetAllHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	images, err := ir.service.GetAllService(ctx)
	if err != nil {
		responses.HTTPError(w, r, http.StatusNotFound, err.Error())
		return
	}

	responses.JSON(w, r, http.StatusOK, responses.Map{"images": images})
}

func (ir *ImageRouter) GetByAttributeHandler(w http.ResponseWriter, r *http.Request) {
	var image models.ControllersImage
	err := json.NewDecoder(r.Body).Decode(&image)
	if err != nil {
		responses.HTTPError(w, r, http.StatusBadRequest, err.Error())
		return
	}

	defer r.Body.Close()

	ctx := r.Context()
	images, err2 := ir.service.GetByAttributeService(ctx, &image)
	if err2 != nil {
		responses.HTTPError(w, r, http.StatusNotFound, err2.Error())
		return
	}

	responses.JSON(w, r, http.StatusOK, responses.Map{"images": images})
}

func (ir *ImageRouter) UpdateHandler(w http.ResponseWriter, r *http.Request) {
	filename := chi.URLParam(r, "filename")

	var image models.ControllersImage
	err := json.NewDecoder(r.Body).Decode(&image)
	if err != nil {
		responses.HTTPError(w, r, http.StatusBadRequest, err.Error())
		return
	}

	defer r.Body.Close()

	ctx := r.Context()
	err = ir.service.UpdateService(ctx, filename, &image)
	if err != nil {
		responses.HTTPError(w, r, http.StatusNotFound, err.Error())
		return
	}

	responses.JSON(w, r, http.StatusOK, responses.Map{"image": image})
}

func (ir *ImageRouter) DeleteHandler(w http.ResponseWriter, r *http.Request) {
	filename := chi.URLParam(r, "filename")

	ctx := r.Context()
	err := ir.service.DeleteService(ctx, filename)
	if err != nil {
		responses.HTTPError(w, r, http.StatusNotFound, err.Error())
		return
	}

	responses.JSON(w, r, http.StatusNoContent, responses.Map{})
}

func (ir *ImageRouter) Routes() http.Handler {
	r := chi.NewRouter()

	r.Get("/", ir.GetAllHandler)

	r.Post("/", ir.CreateHandler)

	r.Get("/search", ir.GetByAttributeHandler)

	r.Put("/{filename}", ir.UpdateHandler)

	r.Delete("/{filename}", ir.DeleteHandler)

	return r
}
