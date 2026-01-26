package repositories

import (
	"context"
	"e-commerce/internal/models"
)

type ProductRepository interface {
	AddProduct(ctx context.Context, payload *models.CreateProductRequest) (*models.GetProductResponse, error)
	GetProducts(ctx context.Context, payload *models.GetProductsRequest) (*[]models.GetProductResponse, error)
	GetProduct(ctx context.Context, productId string) (*models.GetProductResponse, error)
	AddproductToCategory(ctx context.Context, productId string, categoryId string) error
	CreateProductCategory(ctx context.Context, payload *models.CreateProductCategoryRequest) (*models.ProductCategory, error)
	RemoveProductFromCategory(ctx context.Context, productId string, categoryId string) error
	DeleteProductCategory(ctx context.Context, categoryId string) error
	UpdateProductCategory(ctx context.Context, name *string, description *string, imageUrl *string, categoryId string) (*models.ProductCategory, error)
	GetProductsInCategory(ctx context.Context, categoryId string, limit int, offset int) (*[]models.GetProductInCategoryResponse, error)
	GetCategory(ctx context.Context, categoryId string) (*models.ProductCategory, error)
	GetSubCategories(ctx context.Context, categoryId string) (*[]models.ProductCategory, error)
}
