package models

import (
	"database/sql/driver"
	"e-commerce/internal/errors/apperrors"
	"encoding/json"
	"log"

	"github.com/google/uuid"
)

type ProductModel struct {
	Id uuid.UUID `json:"id" db:"id"`
	Name string `json:"name" db:"name"`
	Quantity int `json:"quantity" db:"quantity"`
	Color string `json:"color" db:"color"`
	Price int64 `json:"price" db:"price"`
	Origin string `json:"origin" db:"origin"`
	About string `json:"about" db:"about"`
	ImageUrls StringArray `json:"image_urls" db:"image_urls"`
	CreatedAt string `json:"created_at" db:"created_at"`
	UpdatedAt string `json:"updated_at" db:"updated_at"`
}

type AddProduct struct {
	Name string `json:"name" db:"name"`
	Quantity int `json:"quantity" db:"quantity"`
	Color string `json:"color" db:"color"`
	Price int64 `json:"price" db:"price"`
	Origin string `json:"origin" db:"origin"`
	About string `json:"about" db:"about"`
}

type GetProductsRequest struct {
	Name     *string
	MinPrice *int64
	MaxPrice *int64
	Color    *string
	Origin   *string
	Limit    int
	Page   int
	Offset int
}

type GetProductResponse struct {
	Id uuid.UUID `json:"id" db:"id"`
	Name string `json:"name" db:"name"`
	Quantity int `json:"quantity" db:"quantity"`
	Color string `json:"color" db:"color"`
	Price int64 `json:"price" db:"price"`
	Categories ProductCategoryArray `json:"categories" db:"categories"`
	Origin string `json:"origin" db:"origin"`
	About string `json:"about" db:"about"`
	ImageUrls StringArray `json:"image_urls" db:"image_urls"`
	CreatedAt string `json:"created_at" db:"created_at"`
}

type GetProductInCategoryResponse struct {
	Id uuid.UUID `json:"id" db:"id"`
	Name string `json:"name" db:"name"`
	Quantity int `json:"quantity" db:"quantity"`
	Color string `json:"color" db:"color"`
	Price int64 `json:"price" db:"price"`
	Origin string `json:"origin" db:"origin"`
	About string `json:"about" db:"about"`
	ImageUrls StringArray `json:"image_urls" db:"image_urls"`
}

type ProductCategory struct {
	Id string `json:"id" db:"id"`
	Name string `json:"name" db:"name"`
	Description string `json:"description" db:"description"`
}

type ProductCategoryArray []ProductCategory

func (p *ProductCategoryArray) Scan(value interface{}) error {
	if value == nil {
		*p = []ProductCategory{}
		return nil
	}

	bytes, ok := value.([]byte)
	if !ok {
		log.Println("failed to scan ProductCategoryArray")
		return apperrors.InternalServerError("failed to scan ProductCategoryArray")
	}

	return json.Unmarshal(bytes, p)
}

type StringArray []string

func (a *StringArray) Scan(value interface{}) error {
	if value == nil {
		*a = []string{}
		return nil
	}

	bytes, ok := value.([]byte)
	if !ok {
		log.Println("failed to scan StringArray")
		return apperrors.InternalServerError("failed to scan StringArray")
	}

	return json.Unmarshal(bytes, a)
}

func (a StringArray) Value() (driver.Value, error) {
	return json.Marshal(a)
}

type AddProductToCategory struct {
	ProductId string `json:"product_id" validate:"required" uuid:"4"`
	CategoryId string `json:"category_id" validate:"required" uuid:"4"`
}

type CreateProductCategoryRequest struct {
	Name string `json:"name" validate:"required,max=15"`
	Description string `json:"description" validate:"required"`
}