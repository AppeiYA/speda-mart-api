package auth

import (
	"e-commerce/internal/handlers"

	"github.com/gorilla/mux"
)

func RegisterAuthRoute(r *mux.Router ,ah *handlers.AuthHandler){
	auth := r.PathPrefix("/auth").Subrouter()

	// authentication
	auth.HandleFunc("/login", ah.LoginUser).Methods("POST")
	auth.HandleFunc("/exists/{email}", ah.CheckUserExists).Methods("GET")

	// google auth login
	auth.HandleFunc("/login/{provider}", ah.GoogleAuthLogin).Methods("GET")
	auth.HandleFunc("/callback/{provider}", ah.GoogleCallback).Methods("GET")
}