package repositories

import (
	"github.com/XDBerry29/e-commerce-go+react/domain/models"
	"gorm.io/gorm"
)

type productRepository struct {
	DB *gorm.DB
}

func NewProductRepository(db *gorm.DB) ProductRepository {
	return &productRepository{
		DB: db,
	}
}

func (r *productRepository) CreateProduct(product *models.Product) error {
	return r.DB.Create(product).Error
}

func (r *productRepository) GetAllProducts() ([]models.Product, error) {
	var products []models.Product

	err := r.DB.Find(&products).Error
	if err != nil {
		return nil, err
	}

	return products, nil
}
