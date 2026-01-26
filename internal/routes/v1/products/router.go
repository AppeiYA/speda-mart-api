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

	admin.HandleFunc("", ph.AddProduct).Methods("POST")
	admin.HandleFunc("/categories", ph.CreateCategory).Methods("POST")
	admin.HandleFunc("/categories", ph.UpdateProductCategory).Methods("PATCH")
	admin.HandleFunc("/categories/add-product", ph.AddProductToCategory).Methods("POST")
	admin.HandleFunc("/categories/remove-product", ph.RemoveProductFromCategory).Methods("POST")
	admin.HandleFunc("/categories/{category_id}", ph.DeleteProductCategory).Methods("DELETE")


	// unprotected routes
	products.HandleFunc("", ph.GetProducts).Methods("GET")
	products.HandleFunc("/categories/{category_id}/subcategories", ph.GetSubCategories).Methods("GET")
	products.HandleFunc("/categories/{category_id}", ph.GetProductsInCategory).Methods("GET")
	products.HandleFunc("/{product_id:[a-fA-F0-9\\-]+}", ph.GetProduct).Methods("GET")
}
