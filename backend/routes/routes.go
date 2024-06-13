package routes

import (
	"github.com/XDBerry29/e-commerce-go+react/controller"
	"github.com/XDBerry29/e-commerce-go+react/routes/middleware"
	"github.com/labstack/echo/v4"
)

func AuthRoutes(e *echo.Echo, authController *controller.AuthController) {
	e.POST("/register", authController.Register)
	e.POST("/login", authController.Login)
}

func ProductRoutes(e *echo.Echo, productController *controller.ProductController) {
	e.GET("/products", productController.GetProducts)

	admin := e.Group("/admin")
	admin.Use(middleware.JWTAuth)
	admin.Use(middleware.AdminOnly)

	admin.POST("/product", productController.AddProduct)
	admin.GET("/products", productController.GetProductsAdmin)
	admin.DELETE("/product", productController.DeleteProduct)
	admin.PUT("/product", productController.ModifyProduct)

}
