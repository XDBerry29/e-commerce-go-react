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
	if err := productController.productService.AddProduct(input.Name, input.Description, input.Photo, input.Stock); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "Product added successfully"})
}

func (productController *ProductController) GetProducts(c echo.Context) error {
	products, err := productController.productService.GetAllProducts()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, products)

}
