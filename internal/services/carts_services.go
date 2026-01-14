package services

import (
	"context"
	"e-commerce/internal/errors/apperrors"
	"e-commerce/internal/models"
	"e-commerce/internal/repositories"
	"errors"
)

const (
	MAXINCART=10
)

type CartsService struct {
	cartsRepo repositories.CartsRepository
	productRepo repositories.ProductRepository
}

func NewCartsService(
	cartsRepo repositories.CartsRepository, 
	productRepo repositories.ProductRepository,
	) *CartsService {
	return &CartsService{
		cartsRepo: cartsRepo,
		productRepo: productRepo,
	}
}

func (s *CartsService) AddToCart(ctx context.Context, userId string, payload *models.AddToCart) error {
	// check if product exists
	_, err := s.productRepo.GetProduct(ctx, payload.ProductId)
	if err != nil {
		return err
	}
	// check if user already has a cart and there is space in it
	cart, err := s.cartsRepo.CheckAvailableCart(ctx, userId)
	if err != nil {
		var appErr *apperrors.ErrorResponse
		if errors.As(err, &appErr) && appErr.Code == apperrors.ErrNotFound {
			cart, err = s.cartsRepo.CreateCart(ctx, userId)
			if err != nil {
				return err
			}
		} else {
		return err
		}
	}
	// check if items are too much in cart
	if cart.ItemCount >= MAXINCART {
		return apperrors.ConflictError("cart is full")
	}

	// add item to cart 
	addErr := s.cartsRepo.AddToCart(ctx, cart.CartId, payload)
	if err != nil {
		return addErr
	}

	return nil
}