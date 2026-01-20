package users

import (
	"e-commerce/internal/handlers"
	"e-commerce/internal/middlewares"

	"github.com/gorilla/mux"
)

func RegisterUserRoutes(r *mux.Router, uh *handlers.UserHandler, am *middlewares.AuthMiddleware){
	users := r.PathPrefix("/users").Subrouter()

	// public routes
	users.HandleFunc("/register", uh.CreateUser).Methods("POST")

	// protected routes (users)

	// protected route (admin)
	admin := users.NewRoute().Subrouter()
	admin.Use(am.AuthAdmin())

	
}