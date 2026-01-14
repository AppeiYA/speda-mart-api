package v1

import (
	"e-commerce/internal/handlers"
	"e-commerce/internal/middlewares"
	"e-commerce/internal/routes/v1/auth"
	"e-commerce/internal/routes/v1/carts"
	"e-commerce/internal/routes/v1/products"
	"e-commerce/internal/routes/v1/users"

	"github.com/gorilla/mux"
)

func NewV1Router(
	uh *handlers.UserHandler, 
	ah *handlers.AuthHandler, 
	ph *handlers.ProductHandler,
	ch *handlers.CartsHandler, 
	am *middlewares.AuthMiddleware,
	) *mux.Router {
	r := mux.NewRouter()

	api := r.PathPrefix("/api/v1").Subrouter()

	users.RegisterUserRoutes(api, uh)
	auth.RegisterAuthRoute(api, ah)
	products.RegisterProductRoutes(api, ph)
	carts.RegisterCartsRoutes(api, ch, am)

	return r
}