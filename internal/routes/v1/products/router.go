package products

import (
	"e-commerce/internal/handlers"

	"github.com/gorilla/mux"
)

func RegisterProductRoutes(r *mux.Router, ph *handlers.ProductHandler) {
	products := r.PathPrefix("/products").Subrouter()

	// unprotected routes
	products.HandleFunc("", ph.GetProducts).Methods("GET")
	products.HandleFunc("/{product_id}", ph.GetProduct).Methods("GET")
}
