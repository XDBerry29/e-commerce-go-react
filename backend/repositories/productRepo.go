package repositories

import "github.com/XDBerry29/e-commerce-go+react/domain/models"

type ProductRepository interface {
	CreateProduct(product *models.Product) error
	GetAllProducts() ([]models.Product, error)
}
