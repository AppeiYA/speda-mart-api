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

func (s *ProductService) CreateCategory(ctx context.Context, req *models.CreateProductCategoryRequest) (*models.ProductCategory, error) {
	category, err := s.ProductRepo.CreateProductCategory(ctx, req)
	if err != nil {
		return nil, err
	}
	return category, nil
}

func (s *ProductService) AddProductToCategory(ctx context.Context, productId string, categoryId string) error {
	if _, err := uuid.Parse(productId); !(err == nil) {
		return apperrors.BadException("product does not exist")
	}

	if _, err := uuid.Parse(categoryId); !(err == nil) {
		return apperrors.BadException("category does not exist")
	}
	// check if product exists 
	_, err := s.ProductRepo.GetProduct(ctx, productId)
	if err != nil {
		return err
	}
	
	// check if category exist
	_, err = s.ProductRepo.GetCategory(ctx, categoryId)
	if err != nil {
		return err
	}

	// add product to category
	err = s.ProductRepo.AddproductToCategory(ctx, productId, categoryId)
	if err != nil {
		return err
	}
	return nil
}