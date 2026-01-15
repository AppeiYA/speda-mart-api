package carts

import (
	"e-commerce/internal/handlers"
	"e-commerce/internal/middlewares"

	"github.com/gorilla/mux"
)

func RegisterCartsRoutes(r *mux.Router ,ch *handlers.CartsHandler, am *middlewares.AuthMiddleware) {
	carts := r.PathPrefix("/carts").Subrouter()

	// protected routes
	protected := carts.NewRoute().Subrouter()
	protected.Use(am.Auth())

	protected.HandleFunc("", ch.AddToCart).Methods("POST")
	protected.HandleFunc("", ch.GetUserCart).Methods("GET")
	protected.HandleFunc("/{product_id}", ch.DeleteItemFromCart).Methods("DELETE")
}