package services

import (
	"context"

	"github.com/FumingPower3925/SimpleImageServer/pkg/models"
	"github.com/FumingPower3925/SimpleImageServer/pkg/repositories/image"
)

type ImageService struct {
	repository image.IImageRepository
}

func NewImageService(repo image.IImageRepository) *ImageService {
	return &ImageService{repository: repo}
}

func (is *ImageService) CreateService(ctx context.Context, image *models.ControllersImage) error {
	return is.repository.CreateImage(ctx, image)
}

func (is *ImageService) GetByAttributeService(ctx context.Context, image *models.ControllersImage) (*[]models.RepositoryImage, error) {
	return is.repository.GetByAttributeImage(ctx, image)
}

func (is *ImageService) GetAllService(ctx context.Context) (*[]models.RepositoryImage, error) {
	return is.repository.GetAllImages(ctx)
}

func (is *ImageService) UpdateService(ctx context.Context, id string, image *models.ControllersImage) error {
	return is.repository.UpdateImage(ctx, id, image)
}

func (is *ImageService) DeleteService(ctx context.Context, id string) error {
	return is.repository.DeleteImage(ctx, id)
}
