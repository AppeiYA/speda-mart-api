package postgres

import (
	"context"
	"database/sql"
	"e-commerce/internal/db"
	"e-commerce/internal/errors/apperrors"
	"e-commerce/internal/models"
	"e-commerce/internal/repositories/postgres/cart_queries"
	"errors"
	"log"
)

type CartsRepository struct {
	db *db.DB
}

func NewCartsRepository(db *db.DB) *CartsRepository {
	return &CartsRepository{
		db: db,
	}
}

func (r *CartsRepository) CreateCart(ctx context.Context, userId string) (*models.CheckAvailableCart , error) {
	var cartId string
	err := r.db.QueryRowContext(ctx, cart_queries.CREATECART, userId).Scan(
		&cartId,
	)
	if err != nil {
		log.Println("Error creating cart", err)
		return nil, err
	}
	return &models.CheckAvailableCart{Id: cartId, ItemCount: 0}, nil
}

func (r *CartsRepository) GetuserCart(ctx context.Context, userId string) (*models.GetCartResponse, error) {
	var userCart models.GetCartResponse
	err := r.db.GetContext(ctx, &userCart, cart_queries.GETUSERCART, userId)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, apperrors.NotFoundError("cart not found")
		}
		log.Println("Error getting in db: ", err)
		return nil, apperrors.InternalServerError("error in db get")
	}

	return &userCart, nil
}

func (r *CartsRepository) AddToCart(ctx context.Context, cartId string, payload *models.AddToCart) error {
	_, err := r.db.ExecContext(ctx, cart_queries.ADDTOCART, cartId, payload.ProductId, payload.Quantity, payload.UnitPrice)
	if err != nil {
		log.Println("Error adding item to cart")
		return err
	}
	return nil
}

func (r *CartsRepository) UpdateProductQuantity(ctx context.Context, payload *models.UpdateProductQuantityInCart) error {
	_, err := r.db.ExecContext(ctx, cart_queries.UPDATEPRODUCTQUANTITY, payload.Quantity, payload.CartId, payload.ProductId)
	if err != nil {
		log.Println("Error updating product quantity ", err)
		return err
	}
	return nil
}

func (r *CartsRepository) DeleteFromCart(ctx context.Context, cartId string, productId string) error {
	_, err := r.db.ExecContext(ctx, cart_queries.DELETEFROMCART, cartId, productId)
	if err != nil {
		log.Println("Error deleting product from cart ", err)
		return err
	}
	return nil
}

func (r *CartsRepository) CheckAvailableCart(ctx context.Context, userId string) (*models.CheckAvailableCart, error) {
	var cart models.CheckAvailableCart

	err := r.db.GetContext(ctx, &cart, cart_queries.CHECKAVAILABLECART, userId)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, apperrors.NotFoundError("cart not found")
		}
		log.Println("Error getting in db: ", err)
		return nil, apperrors.InternalServerError("error in db get")
	}
	return &cart, nil
}