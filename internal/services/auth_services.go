package services

import (
	"context"
	"e-commerce/internal/errors/apperrors"
	"e-commerce/internal/models"
	"e-commerce/internal/repositories"
	"e-commerce/internal/utils"
	"e-commerce/package/jwt"
	"errors"
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

func (s *AuthService) GoogleAuthLogin(ctx context.Context, payload *models.CreateUser) (*models.GetUserResponse, string, error) {
	// check if email already exists for user
	user, err := s.UserRepo.GetUserForAuth(ctx, payload.Email)
	if err != nil {
		var appErr *apperrors.ErrorResponse
		// if user is not found, create a new user
		if errors.As(err, &appErr) && appErr.Code == apperrors.ErrNotFound {
			// hash random password
			password, err := utils.GenerateRandomStringForHashing(32)
			if err != nil {
				return nil, "", err
			}
			hash, err := utils.HashPassword(password)
			if err != nil {
				return nil, "", err
			}

			newUser := &models.CreateUser{
				FirstName: payload.FirstName,
				LastName: payload.LastName,
				Email: payload.Email,
				Password: hash,
			}

			user, err = s.UserRepo.CreateUser(ctx, newUser)
			if err != nil {
				return nil, "", err
			}
		}else {
			return nil, "", err
		}
	}

	token, err := s.JwtService.GenerateToken(
		jwt.UserPayload{
			UserId: user.Id.String(), 
			Email: user.Email, 
			Role: user.Role},
		)

	if err != nil {
		return nil, "", apperrors.InternalServerError("error siginig token")
	}

	resp := models.GetUserResponse{
		FirstName: user.FirstName,
		LastName: user.LastName,
		Email: user.Email,
		Role: user.Role,
		CreatedAt: user.CreatedAt,
	}

	return &resp, token, nil

}

func (s *AuthService) UserExists(ctx context.Context, email string) (bool, error) {
	_, err := s.UserRepo.GetUserByEmail(ctx, email)
	if err != nil {
		if appErr, ok := err.(*apperrors.ErrorResponse); ok {
			if appErr.Code == apperrors.ErrNotFound {
				return false, nil
			}
			return false, err 
		}
		return false, err
	}
	return true, nil
}