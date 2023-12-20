package image

import (
	"context"

	"github.com/FumingPower3925/SimpleImageServer/pkg/models"
)

type IImageRepository interface {
	CreateImage(ctx context.Context, image *models.ControllersImage) error
	GetByAttributeImage(ctx context.Context, image *models.ControllersImage) (*[]models.RepositoryImage, error)
	GetAllImages(ctx context.Context) (*[]models.RepositoryImage, error)
	UpdateImage(ctx context.Context, filename string, image *models.ControllersImage) error
	DeleteImage(ctx context.Context, filename string) error
}
