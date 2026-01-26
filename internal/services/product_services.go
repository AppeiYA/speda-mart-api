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

func (s *ProductService) AddProduct(ctx context.Context, payload *models.CreateProductRequest) (*models.GetProductResponse, error){
	// create product 
	product, err := s.ProductRepo.AddProduct(ctx, payload)
	if err != nil {
		return nil, err
	}
	return product, nil
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

func (s *ProductService) RemoveProductFromCategory(ctx context.Context, productId, categoryId string) error {
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

	// check if product is in category
	product, err := s.ProductRepo.GetProduct(ctx, productId)
	if err != nil {
		return err
	}
	inCategory := ProductInCategory(product, categoryId)
	if !inCategory {
		return apperrors.BadException("product not in category")
	}

	// remove product from category
	err = s.ProductRepo.RemoveProductFromCategory(ctx, productId, categoryId)
	if err != nil {
		return err
	}
	return nil
}

func (s *ProductService) GetProductsInCategory(ctx context.Context, categoryId string, limit, offset int) (*[]models.GetProductInCategoryResponse, error) {
	if _, err := uuid.Parse(categoryId); !(err == nil) {
		return nil, apperrors.BadException("category does not exist")
	}

	// check if category exist
	_, err := s.ProductRepo.GetCategory(ctx, categoryId)
	if err != nil {
		return nil, err
	}

	// get products in category
	products, err := s.ProductRepo.GetProductsInCategory(ctx, categoryId, limit, offset)
	if err != nil {
		return nil, err
	}

	return  products, nil
}

func (s *ProductService) GetSubCategories(ctx context.Context, categoryId string) (*[]models.ProductCategory, error) {
	if _, err := uuid.Parse(categoryId); !(err == nil) {
		return nil, apperrors.BadException("category does not exist")
	}

	// check if category exist
	_, err := s.ProductRepo.GetCategory(ctx, categoryId)
	if err != nil {
		return nil, err
	}

	subCategories, err := s.ProductRepo.GetSubCategories(ctx, categoryId)
	if err != nil {
		return nil, err
	}

	return subCategories, nil
}

func (s *ProductService) DeleteProductCategory(ctx context.Context, categoryId string) error {
	if _, err := uuid.Parse(categoryId); !(err == nil) {
		return apperrors.BadException("category does not exist")
	}

	// check if category exist
	_, err := s.ProductRepo.GetCategory(ctx, categoryId)
	if err != nil {
		return err
	}

	// delete category
	err = s.ProductRepo.DeleteProductCategory(ctx, categoryId)
	if err != nil {
		return err
	}
	return nil
}

func (s *ProductService) UpdateProductCategory(ctx context.Context, req *models.UpdateProductCategoryRequest) (*models.ProductCategory, error) {
	if _, err := uuid.Parse(req.CategoryId); !(err == nil) {
		return nil, apperrors.BadException("category does not exist")
	}

	// check if category exist
	_, err := s.ProductRepo.GetCategory(ctx, req.CategoryId)
	if err != nil {
		return nil, err
	}

	// update category
	updatedCategory, err := s.ProductRepo.UpdateProductCategory(ctx, req.Name, req.Description, req.ImageUrl, req.CategoryId)
	if err != nil {
		return nil, err
	}

	return updatedCategory, nil
}

func ProductInCategory(product *models.GetProductResponse, categoryId string) bool {
	if len(product.Categories) == 0 {
		return false
	}

	for _, category := range product.Categories {
		if category.Id == categoryId {
			return true
		}
	}
	return false
}