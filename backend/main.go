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
)

func main() {

	e := echo.New()
	e.Use(middleware.Logger())  // Logs each request
	e.Use(middleware.Recover()) // Recovers from panics and returns 500 error

	if err := godotenv.Load(".env"); err != nil {
		log.Fatal("Error loading the env...")
	}

	db, err := utils.ConnectDatabase()
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	//INITIATE THE INSTANCES OF THE APP
	userRepo := repositories.NewUserRepository(db)
	authService := service.NewAuthService(userRepo)
	authController := controller.NewAuthController(authService)
	routes.AuthRoutes(e, authController)

	PORT := os.Getenv("PORT")
	e.Logger.Fatal(e.Start(":" + PORT))
}
