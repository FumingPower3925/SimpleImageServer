package services

import (
	"context"

	"github.com/FumingPower3925/SimpleImageServer/pkg/models"
	"github.com/FumingPower3925/SimpleImageServer/pkg/repositories/user"
)

type UserService struct {
	repository user.IUserRepository
}

func NewUserService(repo user.IUserRepository) *UserService {
	return &UserService{repository: repo}
}

func (us *UserService) CreateService(ctx context.Context, user *models.User) error {
	return us.repository.CreateUser(ctx, user)
}

func (us *UserService) CheckCredentialsService(ctx context.Context, id string, password string) (bool, error) {
	return us.repository.CheckCredentialsUser(ctx, id, password)
}

func (us *UserService) ExistsService(ctx context.Context, id string) (bool, error) {
	return us.repository.ExistsCreator(ctx, id)
}

func (us *UserService) DeleteService(ctx context.Context, id string) error {
	return us.repository.DeleteUser(ctx, id)
}
