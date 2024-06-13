package repositories

import (
	"github.com/XDBerry29/e-commerce-go+react/domain/models"
	"github.com/google/uuid"
)

type ProductRepository interface {
	CreateProduct(product *models.Product) error
	GetAllProducts() ([]models.Product, error)
	GetProduct(id uuid.UUID) (*models.Product, error)
	DeleteProduct(id uuid.UUID) error
	ModifyProduct(updatedProduct *models.Product) error
	GetProductByPartial(product *models.Product) (*models.Product, error)
}
