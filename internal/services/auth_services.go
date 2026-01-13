package services

import (
	"context"
	"e-commerce/internal/errors/apperrors"
	"e-commerce/internal/models"
	"e-commerce/internal/repositories"
	"e-commerce/internal/utils"
	"e-commerce/package/jwt"
)

type AuthService struct {
	UserRepo repositories.UserRepository
	JwtService *jwt.JwtService
}

func NewAuthService(userRepo repositories.UserRepository, jwtService *jwt.JwtService) *AuthService {
	return &AuthService{
		UserRepo: userRepo,
		JwtService: jwtService,
	}
}

func (s *AuthService) Login(ctx context.Context, payload *models.LoginRequest) (*models.GetUserResponse, string, error) {
	// check if user exists 
	exists, err := s.UserRepo.GetUserForAuth(ctx, payload.Email)
	if err != nil {
		return nil, "", err
	}

	// if user exists, validate password
	matchErr := utils.CompareHashAndPassword(payload.Password, exists.Hash)
	if !matchErr {
		return nil, "", apperrors.UnauthorizedException("Incorrect credentials")
	}

	// if user password is correct, generate token
	token, err := s.JwtService.GenerateToken(
		jwt.UserPayload{
			UserId: exists.Id.String(), 
			Email: exists.Email, 
			Role: exists.Role},
		)

	if err != nil {
		return nil, "", apperrors.InternalServerError("error siginig token")
	}

	resp := models.GetUserResponse{
		FirstName: exists.FirstName,
		LastName: exists.LastName,
		Email: exists.Email,
		Role: exists.Role,
		CreatedAt: exists.CreatedAt,
	}

	return &resp, token, nil
}