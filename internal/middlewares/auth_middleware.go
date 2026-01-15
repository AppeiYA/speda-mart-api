package middlewares

import (
	"context"
	s "e-commerce/internal/shared"
	"e-commerce/package/jwt"
	"net/http"
	"strings"
)

type AuthMiddleware struct {
	JwtService *jwt.JwtService
}

func NewAuthMiddleware(jwtService *jwt.JwtService) *AuthMiddleware {
	return &AuthMiddleware{
		JwtService: jwtService,
	}
}

func (a *AuthMiddleware) Auth() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			authHeader := r.Header.Get("Authorization")
			if authHeader == "" {
				s.ReqResponse(w, http.StatusUnauthorized, s.Payload{Message: "no token provided"})
				return
			}

			tokenStr := strings.TrimPrefix(authHeader, "Bearer ")
			claims, err := a.JwtService.VerifyToken(tokenStr)
			if err != nil {
				s.ReqResponse(w, http.StatusUnauthorized, s.Payload{Message: "invalid token"})
				return
			}

			ctx := context.WithValue(r.Context(), userCtxKey, claims)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
func (a *AuthMiddleware) AuthAdmin() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			authHeader := r.Header.Get("Authorization")
			if authHeader == "" {
				s.ReqResponse(w, http.StatusUnauthorized, s.Payload{Message: "no token provided"})
				return
			}

			tokenStr := strings.TrimPrefix(authHeader, "Bearer ")
			claims, err := a.JwtService.VerifyToken(tokenStr)
			if err != nil {
				s.ReqResponse(w, http.StatusUnauthorized, s.Payload{Message: "invalid token"})
				return
			}
			if claims.Role != "admin" {
				s.ReqResponse(w, http.StatusUnauthorized, s.Payload{Message: "User unauthorized to access route"})
				return
			}

			ctx := context.WithValue(r.Context(), userCtxKey, claims)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}