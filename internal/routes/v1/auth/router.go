package auth

import (
	"e-commerce/internal/handlers"

	"github.com/gorilla/mux"
)

func RegisterAuthRoute(r *mux.Router ,ah *handlers.AuthHandler){
	auth := r.PathPrefix("/auth").Subrouter()

	// authentication
	auth.HandleFunc("/login", ah.LoginUser).Methods("POST")
}