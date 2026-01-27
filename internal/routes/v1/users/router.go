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
	protected := users.NewRoute().Subrouter()
	protected.Use(am.Auth())

	protected.HandleFunc("/me", uh.GetUserProfile).Methods("GET")

	// protected route (admin)
	admin := users.NewRoute().Subrouter()
	admin.Use(am.AuthAdmin())

	
}