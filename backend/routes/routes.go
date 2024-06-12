package routes

import (
	"github.com/XDBerry29/e-commerce-go+react/controller"
	"github.com/labstack/echo/v4"
)

func AuthRoutes(e *echo.Echo, authController *controller.AuthController) {
	e.POST("/register", authController.Register)
	e.POST("/login", authController.Login)
}
