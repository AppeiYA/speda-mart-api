package middlewares

import (
	"strings"

	"github.com/rs/cors"
)

func CORSMiddleware() *cors.Cors {
	return cors.New(cors.Options{
		// AllowedOrigins:   []string{"http://localhost:3000", "http://localhost:5173"},
		AllowOriginFunc: func(origin string) bool {
			return strings.HasPrefix(origin, "http://localhost:") || origin == "http://deployment.com"
		},
		AllowedMethods:   []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Authorization", "Content-Type", "X-Requested-With", "Application"},
		ExposedHeaders:   []string{"Authorization"},
		AllowCredentials: true,
		MaxAge:           3600,
		Debug:            false,
	})
}
