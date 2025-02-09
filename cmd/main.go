package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/prakoso-id/go-windsurf/internal/application/services"
	"github.com/prakoso-id/go-windsurf/internal/infrastructure/middleware"
	"github.com/prakoso-id/go-windsurf/internal/infrastructure/persistence"
	"github.com/prakoso-id/go-windsurf/internal/interfaces/handlers"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	db, err := persistence.NewPostgresDB()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Initialize repositories
	userRepo := persistence.NewUserRepository(db)
	productRepo := persistence.NewProductRepository(db)

	// Initialize services
	authService := services.NewAuthService(os.Getenv("JWT_SECRET"))
	userService := services.NewUserService(userRepo)
	productService := services.NewProductService(productRepo)

	// Initialize handlers
	authHandler := handlers.NewAuthHandler(authService, userService)
	userHandler := handlers.NewUserHandler(userService)
	productHandler := handlers.NewProductHandler(productService)

	// Initialize router
	r := gin.Default()

	// Public routes
	api := r.Group("/api/v1")
	{
		api.POST("/login", authHandler.Login)
		api.POST("/register", userHandler.Register)

		// Protected routes
		protected := api.Group("/")
		protected.Use(middleware.AuthMiddleware(authService))
		{
			// User routes
			users := protected.Group("/users")
			{
				users.GET("/profile", userHandler.GetProfile)
				users.PUT("/profile", userHandler.UpdateProfile)
				users.POST("/change-password", userHandler.ChangePassword)
			}

			// Product routes
			products := protected.Group("/products")
			{
				products.POST("/", productHandler.CreateProduct)
				products.GET("/:id", productHandler.GetProduct)
				products.PUT("/:id", productHandler.UpdateProduct)
				products.DELETE("/:id", productHandler.DeleteProduct)
				products.GET("/", productHandler.GetAllProducts)
			}
		}
	}

	if err := r.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
