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

type ProductVariant struct {
	Id uuid.UUID `json:"id" db:"id"`
	ProductId uuid.UUID `json:"product_id" db:"product_id"`
	Color string `json:"color" db:"color"`
	Quantity int `json:"quantity" db:"quantity"`
	Price int64 `json:"price" db:"price"`
	ImageUrls StringArray `json:"image_urls" db:"image_urls"`
	Weight int64 `json:"weight" db:"weight"`
	CreatedAt string `json:"created_at" db:"created_at"`
	UpdatedAt string `json:"updated_at" db:"updated_at"`
}

type VariantArray []ProductVariant

func (v *VariantArray) Scan(value interface{}) error {
	if value == nil {
		*v = []ProductVariant{}
		return nil
	}

	bytes, ok := value.([]byte)
	if !ok {
		log.Println("failed to scan ProductCategoryArray")
		return apperrors.InternalServerError("failed to scan ProductCategoryArray")
	}

	return json.Unmarshal(bytes, v)
}



type CreateProductRequest struct {
	Name string `json:"name" db:"name" validate:"required,max=15"`
	ImageUrls StringArray `json:"image_urls" db:"image_urls" validate:"required"`
	Variants  []VariantRequest `json:"variants" validate:"required,dive"`
	Origin string `json:"origin" db:"origin" validate:"required"`
	About string `json:"about" db:"about" validate:"required"`
}

type VariantRequest struct {
    Color     string   `json:"color" validate:"required"`
    Quantity  int      `json:"quantity" validate:"required,min=1"`
	Weight int64 `json:"weight"`
    Price     int64    `json:"price" validate:"required,min=100"`
    ImageUrls StringArray `json:"image_urls" validate:"required"`
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
	TotalQuantity int `json:"total_quantity" db:"total_quantity"`
	Categories ProductCategoryArray `json:"categories" db:"categories"`
	Origin string `json:"origin" db:"origin"`
	About string `json:"about" db:"about"`
	ImageUrls StringArray `json:"image_urls" db:"image_urls"`
	CreatedAt string `json:"created_at" db:"created_at"`
	Variants VariantArray `json:"variants" db:"variants"`
}

type GetProductInCategoryResponse struct {
	Id uuid.UUID `json:"id" db:"id"`
	Name string `json:"name" db:"name"`
	TotalQuantity int `json:"quantity" db:"total_quantity"`
	Origin string `json:"origin" db:"origin"`
	About string `json:"about" db:"about"`
	ImageUrls StringArray `json:"image_urls" db:"image_urls"`
	Variants VariantArray `json:"variants" db:"variants"`
	CreatedAt string `json:"created_at" db:"created_at"`
}

type ProductCategory struct {
	Id string `json:"id" db:"id"`
	Name string `json:"name" db:"name"`
	Description string `json:"description" db:"description"`
	ImageUrl *string `json:"image_url" db:"image_url"`
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

type RemoveProductFromCategory struct {
	ProductId string `json:"product_id" validate:"required" uuid:"4"`
	CategoryId string `json:"category_id" validate:"required" uuid:"4"`
}

type CreateProductCategoryRequest struct {
	Name string `json:"name" validate:"required,max=15"`
	Description string `json:"description" validate:"required"`
	ImageUrl *string `json:"image_url" db:"image_url"`
}

type UpdateProductCategoryRequest struct {
	Name *string `json:"name"`
	Description *string `json:"description"`
	ImageUrl *string `json:"image_url"`
	CategoryId string `json:"category_id" validate:"required"`
}