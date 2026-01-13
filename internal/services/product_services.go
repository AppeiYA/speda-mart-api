package services

import (
	"context"
	"e-commerce/internal/errors/apperrors"
	"e-commerce/internal/models"
	"e-commerce/internal/repositories"

	"github.com/google/uuid"
)

type ProductService struct {
	ProductRepo repositories.ProductRepository
}

func NewProductService(productRepo repositories.ProductRepository) *ProductService {
	return &ProductService{
		ProductRepo: productRepo,
	}
}

func(s *ProductService) GetProducts(ctx context.Context, payload *models.GetProductsRequest) (*[]models.GetProductResponse, error) {
	products, err := s.ProductRepo.GetProducts(ctx, payload)
	if err != nil {
		return nil, err
	}
	return products, nil
}

func (s *ProductService) GetProduct(ctx context.Context, productId string) (*models.GetProductResponse, error) {
	if _, err := uuid.Parse(productId); !(err == nil) {
		return nil, apperrors.BadException("product does not exist")
	}
	product, err := s.ProductRepo.GetProduct(ctx, productId)
	if err != nil{
		return nil, err
	}
	return product, nil
}