package products

import (
	"e-commerce/internal/handlers"
	"e-commerce/internal/middlewares"

	"github.com/gorilla/mux"
)

func RegisterProductRoutes(r *mux.Router, ph *handlers.ProductHandler, am *middlewares.AuthMiddleware) {
	products := r.PathPrefix("/products").Subrouter()

	// admin protected routes
	admin := products.NewRoute().Subrouter()
	admin.Use(am.AuthAdmin())

	admin.HandleFunc("/categories", ph.CreateCategory).Methods("POST")
	admin.HandleFunc("/categories/add-product", ph.AddProductToCategory).Methods("POST")


	// unprotected routes
	products.HandleFunc("", ph.GetProducts).Methods("GET")
	products.HandleFunc("/{product_id}", ph.GetProduct).Methods("GET")
}
