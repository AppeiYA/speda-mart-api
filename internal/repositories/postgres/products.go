package postgres

import (
	"context"
	"database/sql"
	"e-commerce/internal/db"
	"e-commerce/internal/errors/apperrors"
	"e-commerce/internal/models"
	"e-commerce/internal/repositories/postgres/product_queries"
	"errors"
	"log"
)

type ProductRepository struct {
	db *db.DB
}

func NewProductRepository(db *db.DB) *ProductRepository {
	return &ProductRepository{
		db: db,
	}
}

func (r *ProductRepository) AddProduct(ctx context.Context, payload *models.CreateProductRequest) (*models.GetProductResponse, error) {
	var productId string
	var product models.GetProductResponse
	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}

	defer tx.Rollback()

	createMProductErr := tx.QueryRowContext(ctx, product_queries.INSERTMASTERPRODUCT, payload.Name, payload.Origin, payload.About, payload.ImageUrls).Scan(
		&productId,
	)

	if createMProductErr != nil {
        return nil, createMProductErr
    }

	for _, v := range payload.Variants{
		// INSERT PRODUCT VARIANT
		_,err:=tx.ExecContext(ctx, product_queries.ADDPRODUCTVARIANT, productId, v.Color, v.Quantity, v.Weight, v.Price, v.ImageUrls)
		if err != nil {
			return nil, err
		}
	}

	// get the product and its variant and return it
	getErr := tx.QueryRowContext(ctx, product_queries.GETPRODUCTBYID, productId).Scan(
		&product.Id,
		&product.Name,
		&product.Origin,
		&product.About,
		&product.ImageUrls,
		&product.CreatedAt,
		&product.TotalQuantity,
		&product.Categories,
		&product.Variants,
	)
	if getErr != nil {
		return nil, getErr
	}
	tx.Commit()
	return &product, nil
}

func (r *ProductRepository) GetProducts(ctx context.Context, payload *models.GetProductsRequest) (*[]models.GetProductResponse, error) {
	var products []models.GetProductResponse

	err := r.db.SelectContext(
		ctx,
		&products,
		product_queries.GETPRODUCTS,
		payload.Name,
		payload.MinPrice,
		payload.MaxPrice,
		payload.Color,
		payload.Origin,
		payload.Limit,
		payload.Offset,
	)
	if err != nil {
		log.Println("Error getting from db: ", err)
		return nil, apperrors.InternalServerError("error getting products")
	}

	return &products, nil
}

func (r *ProductRepository) GetProduct(ctx context.Context, productId string) (*models.GetProductResponse, error) {
	var resp models.GetProductResponse

	err := r.db.GetContext(ctx, &resp, product_queries.GETPRODUCTBYID, productId)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, apperrors.NotFoundError("product not found")
		}
		log.Println("Error getting in db: ", err)
		return nil, apperrors.InternalServerError("error in db get")
	}
	return &resp, nil
}

func (r *ProductRepository) AddproductToCategory(ctx context.Context, productId string, categoryId string) error {
	_, err := r.db.ExecContext(ctx, product_queries.ADDPRODUCTTOCATEGORY, productId, categoryId)
	if err != nil {
		log.Println("Error adding product to category", err)
		return apperrors.InternalServerError(err.Error())
	}
	return nil
}

func (r *ProductRepository) CreateProductCategory(ctx context.Context, payload *models.CreateProductCategoryRequest) (*models.ProductCategory, error) {
	var category models.ProductCategory
	err := r.db.QueryRowContext(ctx, product_queries.CREATEPRODUCTCATEGORY, payload.Name, payload.Description, payload.ImageUrl).Scan(
		&category.Id,
		&category.Name,
		&category.Description,
		&category.ImageUrl,
	)
	if err != nil {
		log.Println("Error creating product category", err)
		return nil, apperrors.InternalServerError(err.Error())
	}
	return &category, nil
}

func (r *ProductRepository) RemoveProductFromCategory(ctx context.Context, productId string, categoryId string) error {
	_, err := r.db.ExecContext(ctx, product_queries.REMOVEPRODUCTFROMCATEGORY, productId, categoryId)
	if err != nil {
		log.Println("Error remove product from category", err)
		return apperrors.InternalServerError(err.Error())
	}
	return nil
}

func (r *ProductRepository) DeleteProductCategory(ctx context.Context, categoryId string) error {
	_, err := r.db.ExecContext(ctx, product_queries.DELETEPRODUCTCATEGORY, categoryId)
	if err != nil {
		log.Println("Error deleting product category", err)
		return apperrors.InternalServerError(err.Error())
	}
	return nil
}

func (r *ProductRepository) UpdateProductCategory(ctx context.Context, name *string, description *string, imageUrl *string, categoryId string) (*models.ProductCategory, error){
	var updateCategory models.ProductCategory
	var productCategoryId string

	err := r.db.QueryRowContext(ctx, product_queries.UPDATEPRODUCTCATEGORY, name, description, imageUrl, categoryId).Scan(
		&productCategoryId,
		&updateCategory.Name,
		&updateCategory.Description,
		&updateCategory.ImageUrl,
	)

	if err != nil {
		log.Println("Error updating product category", err)
		return nil, apperrors.InternalServerError(err.Error())
	}

	return &updateCategory, nil
}

func (r *ProductRepository) GetProductsInCategory(ctx context.Context, categoryId string, limit int, offset int) (*[]models.GetProductInCategoryResponse, error) {
	var products []models.GetProductInCategoryResponse

	err := r.db.SelectContext(ctx, &products, product_queries.GETPRODUCTSINCATEGORY, categoryId, limit, offset)
	if err != nil {
		log.Println("Error getting products in category from db: ", err)
		return nil, apperrors.InternalServerError("error getting products in category")
	}
	return &products, nil
}

func (r *ProductRepository) GetCategory(ctx context.Context, categoryId string) (*models.ProductCategory, error) {
	var category models.ProductCategory

	err := r.db.GetContext(ctx, &category, product_queries.GETCATEGORY, categoryId)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, apperrors.NotFoundError("category not found")
		}
		log.Println("Error getting in db: ", err)
		return nil, apperrors.InternalServerError("error in db get")
	}
	return &category, nil
}

func (r *ProductRepository) GetSubCategories(ctx context.Context, categoryId string) (*[]models.ProductCategory, error) {
	var categories []models.ProductCategory

	err := r.db.SelectContext(ctx, &categories, product_queries.GETSUBCATEGORIES, categoryId)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, apperrors.NotFoundError("category not found")
		}
		log.Println("Error getting in db: ", err)
		return nil, apperrors.InternalServerError("error in db get")
	}

	return &categories, nil
}