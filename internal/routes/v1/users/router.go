package users

import (
	"e-commerce/internal/handlers"

	"github.com/gorilla/mux"
)

func RegisterUserRoutes(r *mux.Router, uh *handlers.UserHandler){
	users := r.PathPrefix("/users").Subrouter()

	// public routes
	users.HandleFunc("/register", uh.CreateUser).Methods("POST")
}