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
