package repositories

import (
	"context"
	"e-commerce/internal/models"
)

type CartsRepository interface {
	CreateCart(ctx context.Context, userId string) (*models.CheckAvailableCart , error)
	GetuserCart(ctx context.Context, userId string) (*models.GetCartResponse, error)
	AddToCart(ctx context.Context, cartId string, payload *models.AddToCart) error
	UpdateProductQuantity(ctx context.Context, payload *models.UpdateProductQuantityInCart) error
	DeleteFromCart(ctx context.Context, cartId string, productId string) error
	CheckAvailableCart(ctx context.Context, userId string) (*models.CheckAvailableCart, error)
}