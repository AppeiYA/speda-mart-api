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
	product, err := s.productRepo.GetProduct(ctx, payload.ProductId)
	if err != nil {
		return err
	}
	// check if quantity of product requested to cart matches available amount
	if product.Quantity < payload.Quantity {
		return apperrors.ConflictError("Requested quantity is unavailable")
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
	addErr := s.cartsRepo.AddToCart(ctx, cart.Id, payload)
	if err != nil {
		return addErr
	}

	return nil
}

func (s *CartsService) GetUserCart(ctx context.Context, userId string) (*models.GetCartResponse, error) {
	return s.cartsRepo.GetuserCart(ctx, userId)
}

func (s *CartsService) DeleteItemFromCart(ctx context.Context, userId string, itemId string) error {
	// search if item exists
	_, err := s.productRepo.GetProduct(ctx, itemId)
	if err != nil {
		return err
	}

	// check if item is in cart
	cart, err := s.cartsRepo.GetuserCart(ctx, userId)
	if err != nil {
		return nil
	}

	existsInCart := ProductInCart(cart.Items, itemId)
	if !existsInCart {
		return apperrors.NotFoundError("Product not in cart")
	}

	// delete item from cart
	deleteErr := s.cartsRepo.DeleteFromCart(ctx, cart.Id, itemId)
	if deleteErr != nil {
		return deleteErr
	}
	
	return nil
}

func ProductInCart(items []models.ItemsInCart, productId string) bool {
	for _, v := range items {
		if v.ProductId == productId {
			return true
		}
	}
	return false
}