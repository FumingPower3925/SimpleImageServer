package user

import (
	"context"

	"github.com/FumingPower3925/SimpleImageServer/pkg/models"
)

type IUserRepository interface {
	CheckCredentialsUser(ctx context.Context, id string, password string) (bool, error)
	CreateUser(ctx context.Context, user *models.User) error
	DeleteUser(ctx context.Context, id string) error
	ExistsCreator(ctx context.Context, id string) (bool, error)
}
