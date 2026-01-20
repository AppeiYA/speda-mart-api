package main

import (
	"e-commerce/internal/config"
	"e-commerce/internal/db"
	"e-commerce/internal/db/seed"
	"e-commerce/internal/handlers"

	"e-commerce/internal/middlewares"
	"e-commerce/internal/repositories/postgres"
	v1 "e-commerce/internal/routes/v1"
	"e-commerce/internal/routes/v1/auth"
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

	if err := seed.SeedAdmin(db, cfg.AdminKey, cfg.AdminEmail); err != nil {
		log.Fatalln("Failed to seed admin", err)
	}

	// oauth2 gomniauth
	googleAuth := auth.NewConfigGomniAuth(
		cfg.GoogleSecurityKey, 
		cfg.GoogleClientId, 
		cfg.GoogleClientSecret, 
		cfg.GoogleRedirectUrl,
	)
	googleAuth.InitGomniauth()
	
	// repositories
	userRepo:= postgres.NewUserRepository(db)
	productRepo := postgres.NewProductRepository(db)
	cartRepo := postgres.NewCartsRepository(db)

	// services
	jwtService := &jwt.JwtService{
		JwtSecret: cfg.JwtSecret,
	}
	userServ := services.NewUserService(userRepo)
	authServ := services.NewAuthService(userRepo, jwtService)
	productServ := services.NewProductService(productRepo)
	cartServ := services.NewCartsService(cartRepo, productRepo)

	// middlewares
	authMiddleware := middlewares.NewAuthMiddleware(jwtService)

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
	cartHandler := &handlers.CartsHandler{
		CartService: cartServ,
	}

	// v1 router
	router := v1.NewV1Router(userHandler, authHandler, productHandler, cartHandler, authMiddleware)

	// middlewares
	rateLimiter := middlewares.NewRateLimiter(2, 5)
	corsMiddleware := middlewares.CORSMiddleware()

	handler := corsMiddleware.Handler(
	middlewares.SecurityHeaders(
		rateLimiter.Middleware(router),
	),
)


	log.Println("Server running at http://localhost" + cfg.Port + ". Db connected: " + db.DriverName())
	if err := http.ListenAndServe(cfg.Port, handler); err != nil {
		log.Fatal(err)
	}
}