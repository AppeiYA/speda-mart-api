package repositories

import (
	"context"
	"e-commerce/internal/models"
)

type ProductRepository interface {
	GetProducts(ctx context.Context, payload *models.GetProductsRequest) (*[]models.GetProductResponse, error)
	GetProduct(ctx context.Context, productId string) (*models.GetProductResponse, error)
}
