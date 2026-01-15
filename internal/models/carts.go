package models

import (
	"e-commerce/internal/errors/apperrors"
	"encoding/json"
	"log"
)

type Status string

const (
	StatusActive     Status = "active"
	StatusCheckedOut Status = "checked_out"
	StatusAbandoned  Status = "abandoned"
)

func (s Status) IsValid() bool {
	switch s {
	case StatusActive, StatusCheckedOut, StatusAbandoned:
		return true
	}
	return false
}

type Carts struct {
	Id        string `json:"id" db:"id"`
	UserId    string `json:"user_id" db:"user_id"`
	Status    Status `json:"status" db:"status"`
	CreatedAt string `json:"created_at" db:"created_at"`
	UpdatedAt string `json:"updated_at" db:"updated_at"`
}

type CartItems struct {
	Id        string `json:"id" db:"id"`
	CartId string `json:"cart_id" db:"cart_id"`
	ProductId string `json:"product_id" db:"product_id"`
	Quantity int `json:"quantity" db:"quantity"`
	UnitPrice int64 `json:"unit_price" db:"unit_price"`
	CreatedAt string `json:"created_at" db:"created_at"`
	UpdatedAt string `json:"updated_at" db:"updated_at"`
}

type AddToCart struct {
	ProductId string `json:"product_id" db:"product_id" validate:"required"`
	Quantity int `json:"quantity" db:"quantity" validate:"required,min=1"`
	UnitPrice int64 `json:"unit_price" db:"unit_price" validate:"required,min=100"`
}

type ItemsInCart struct {
	ProductId string `json:"product_id" db:"product_id"`
	Quantity int `json:"quantity" db:"quantity"`
	SnapShotPrice int64 `json:"snapshot_price" db:"snapshot_price"`
	ProductDetails ProductDetails `json:"product_details" db:"product_details"`
}

type ItemsInCartList []ItemsInCart

func (i *ItemsInCartList) Scan(value any) error {
	if value == nil {
		*i = []ItemsInCart{}
		return nil
	}

	bytes, ok := value.([]byte)
	if !ok {
		log.Println("failed to scan ItemsInCart")
		return apperrors.InternalServerError("error scanning cart items")
	}

	return json.Unmarshal(bytes, i)
}

type ProductDetails struct {
	Name string `json:"name" db:"name"`
	Color string `json:"color" db:"color"`
	Origin string `json:"origin" db:"origin"`
	About string `json:"about" db:"about"`
}

type GetCartResponse struct {
	Id        string `json:"id" db:"id"`
	UserId    string `json:"user_id" db:"user_id"`
	Status    Status `json:"status" db:"status"`
	ItemCount int `json:"item_count" db:"item_count"`
	Items ItemsInCartList `json:"items" db:"items"`
}

type UpdateProductQuantityInCart struct {
	CartId string `json:"cart_id" db:"cart_id"`
	ProductId string `json:"product_id" db:"product_id"`
	Quantity int `json:"quantity" db:"quantity"`
}

type CheckAvailableCart struct {
	Id string `json:"id" db:"id"`
	ItemCount int `json:"item_count" db:"item_count"`
}