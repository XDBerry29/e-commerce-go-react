package controller

import (
	"net/http"

	"github.com/XDBerry29/e-commerce-go+react/service"
	"github.com/labstack/echo/v4"
)

type ProductController struct {
	productService *service.ProductService
}

func NewProductController(productService *service.ProductService) *ProductController {
	return &ProductController{
		productService: productService,
	}
}

func (productController *ProductController) AddProduct(c echo.Context) error {
	var input struct {
		Name        string `json:"name" `
		Description string `json:"description"`
		Photo       string `json:"photo"` // should be photo url can be more than 1 photo for future iterations
		Stock       int    `json:"stock"`
	}
	if err := c.Bind(&input); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusNotImplemented, map[string]string{"message": "Not Implemented yet!"})
}
