package models

import (
	"github.com/google/uuid"
)

type UserModel struct {
	Id        uuid.UUID `json:"id" db:"id"`
	FirstName string    `json:"first_name" db:"first_name"`
	LastName  string    `json:"last_name" db:"last_name"`
	Email     string    `json:"email" db:"email"`
	Hash      string    `json:"hash" db:"hash"`
	Role      string    `json:"role" db:"role"`
	CreatedAt string    `json:"created_at" db:"created_at"`
	UpdatedAt string    `json:"updated_at" db:"updated_at"`
}

type CreateUser struct {
	FirstName string `json:"first_name" db:"first_name" validate:"required"`
	LastName  string `json:"last_name" db:"last_name" validate:"required"`
	Email     string `json:"email" db:"email" validate:"email,required"`
	Password  string `json:"password" db:"hash" validate:"required,password"`
}

type GetUserResponse struct {
	FirstName string `json:"first_name" db:"first_name"`
	LastName  string `json:"last_name" db:"last_name"`
	Email     string `json:"email" db:"email"`
	Role      string `json:"role" db:"role"`
	CreatedAt string `json:"created_at" db:"created_at"`
}

type LoginRequest struct {
	Email    string `json:"email" db:"email" validate:"email,required"`
	Password string `json:"password" db:"hash" validate:"required"`
}
