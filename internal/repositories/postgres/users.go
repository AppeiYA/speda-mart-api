package postgres

import (
	"context"
	"database/sql"
	"e-commerce/internal/db"
	"e-commerce/internal/errors/apperrors"
	"e-commerce/internal/models"
	"e-commerce/internal/repositories/postgres/user_queries"
	"errors"
	"log"
)

type UsersRepository struct {
	db *db.DB
}

func NewUserRepository(db *db.DB) *UsersRepository {
	return &UsersRepository{db: db}
}

func (r *UsersRepository) CreateUser(ctx context.Context, payload *models.CreateUser) (*models.UserModel ,error) {
	var user models.UserModel
	err := r.db.QueryRowContext(ctx, 
		user_queries.CREATEUSER, 
		payload.FirstName, 
		payload.LastName, 
		payload.Email, 
		payload.Password).Scan(
			&user.Id,
			&user.FirstName,
			&user.LastName,
			&user.Email,
			&user.Hash,
			&user.Role,
			&user.CreatedAt,
			&user.UpdatedAt,
		)
	if err != nil {
		log.Println("Error creating user ", err)
		return nil, apperrors.InternalServerError("Error creating user in db")
	}
	return &user, nil
}

func (r *UsersRepository) GetUserByEmail(ctx context.Context, email string) (*models.GetUserResponse, error) {
	var getResponse models.GetUserResponse

	err := r.db.GetContext(ctx, &getResponse, user_queries.GETUSERBYEMAIL, email)
	switch {
	case err == nil:
		return &getResponse, nil
	case errors.Is(err, sql.ErrNoRows):
		return nil, apperrors.NotFoundError("user not found")
	default:
		log.Println("Db error: ", err)
		return nil, apperrors.InternalServerError("Db error::GetUserByEmail")
	}
}

func (r *UsersRepository) GetUserForAuth(ctx context.Context, email string) (*models.UserModel, error) {
	var user models.UserModel

	err := r.db.GetContext(ctx, &user, user_queries.GETUSERFORAUTH, email)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, apperrors.NotFoundError("user not found")
		}
		log.Println("Error getting in db: ", err)
		return nil, apperrors.InternalServerError("error in db get")
	}
	return &user, nil
}
