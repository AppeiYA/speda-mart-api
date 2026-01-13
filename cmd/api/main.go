package main

import (
	"e-commerce/internal/config"
	"e-commerce/internal/db"
	"e-commerce/internal/handlers"
	// "e-commerce/internal/middlewares"
	"e-commerce/internal/repositories/postgres"
	v1 "e-commerce/internal/routes/v1"
	"e-commerce/internal/services"
	"e-commerce/package/jwt"
	"log"
	"net/http"
)

func main() {
	cfg := config.LoadConfig()
	
	db, dbErr := db.ConnectDB(cfg.DatabaseUrl)
	if dbErr != nil {
		log.Fatal(dbErr)
	}
	
	// repositories
	userRepo:= postgres.NewUserRepository(db)
	productRepo := postgres.NewProductRepository(db)

	// services
	jwtService := &jwt.JwtService{
		JwtSecret: cfg.JwtSecret,
	}
	userServ := services.NewUserService(userRepo)
	authServ := services.NewAuthService(userRepo, jwtService)
	productServ := services.NewProductService(productRepo)

	// middlewares
	// authMiddleware := middlewares.NewAuthMiddleware(jwtService)

	// handlers
	userHandler := &handlers.UserHandler{
		UserServ: userServ,
	}
	authHandler := &handlers.AuthHandler{
		AuthServ: authServ,
	}
	productHandler := &handlers.ProductHandler{
		ProductServ: productServ,
	}

	// v1 router
	router := v1.NewV1Router(userHandler, authHandler, productHandler)


	log.Println("Server running at http://localhost" + cfg.Port + ". Db connected: " + db.DriverName())
	if err := http.ListenAndServe(cfg.Port, router); err != nil {
		log.Fatal(err)
	}
}