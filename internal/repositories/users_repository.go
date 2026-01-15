package repositories

import (
	"context"
	"e-commerce/internal/models"
)

type UserRepository interface {
	CreateUser(ctx context.Context, payload *models.CreateUser) (*models.UserModel ,error)
	GetUserByEmail(ctx context.Context, email string) (*models.GetUserResponse, error)
	GetUserForAuth(ctx context.Context, email string) (*models.UserModel, error)
}
