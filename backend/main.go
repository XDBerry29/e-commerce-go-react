package main

import (
	"log"
	"os"

	"github.com/XDBerry29/e-commerce-go+react/controller"
	"github.com/XDBerry29/e-commerce-go+react/repositories"
	"github.com/XDBerry29/e-commerce-go+react/routes"
	"github.com/XDBerry29/e-commerce-go+react/service"
	"github.com/XDBerry29/e-commerce-go+react/utils"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gorm.io/gorm"
)

func InitiateServer(e *echo.Echo, db *gorm.DB) {

	e.Use(middleware.Logger())  // Logs each request
	e.Use(middleware.Recover()) // Recovers from panics and returns 500 error

	userRepo := repositories.NewUserRepository(db)
	authService := service.NewAuthService(userRepo)
	authController := controller.NewAuthController(authService)
	routes.AuthRoutes(e, authController)

	productRepo := repositories.NewProductRepository(db)
	productService := service.NewProductService(productRepo)
	productController := controller.NewProductController(productService)
	routes.ProductRoutes(e, productController)

}

func main() {

	e := echo.New()

	if err := godotenv.Load(".env"); err != nil {
		log.Fatal("Error loading the env...")
	}

	db, err := utils.ConnectDatabase()
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	InitiateServer(e, db)

	PORT := os.Getenv("PORT")
	e.Logger.Fatal(e.Start(":" + PORT))
}
