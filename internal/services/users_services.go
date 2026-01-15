package services

import (
	"context"
	"e-commerce/internal/errors/apperrors"
	"e-commerce/internal/models"
	"e-commerce/internal/repositories"
	"e-commerce/internal/utils"
)

type UserService struct {
	userRepo repositories.UserRepository
}

func NewUserService(userRepo repositories.UserRepository) *UserService {
	return &UserService{
		userRepo: userRepo,
	}
}

func (s *UserService) CreateUser(ctx context.Context, payload *models.CreateUser) error {
	// check if user already exists
	exists, _ := s.userRepo.GetUserByEmail(ctx, payload.Email)
	if exists != nil {
		return apperrors.ConflictError("user already exists")
	}

	// hash password
	hash, err := utils.HashPassword(payload.Password)

	if err != nil {
		return apperrors.InternalServerError("Error hashing password")
	}

	// overwrite password in object
	payload.Password = hash

	// create user
	_, createErr := s.userRepo.CreateUser(ctx, payload)
	return createErr
}