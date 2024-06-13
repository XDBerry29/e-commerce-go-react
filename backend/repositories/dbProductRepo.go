package repositories

import (
	"github.com/XDBerry29/e-commerce-go+react/domain/models"
	"github.com/google/uuid"
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

func (r *productRepository) DeleteProduct(id uuid.UUID) error {
	err := r.DB.Delete(&models.Product{}, "id = ?", id).Error
	if err != nil {
		return err
	}
	return nil
}

// GetProduct implements ProductRepository.
func (r *productRepository) GetProduct(id uuid.UUID) (*models.Product, error) {
	var product = &models.Product{ID: id}
	err := r.DB.First(product).Error
	if err != nil {
		return nil, err
	}

	return product, nil

}

// ModifyProduct implements ProductRepository.
func (r *productRepository) ModifyProduct(updatedProduct *models.Product) error {

	err := r.DB.Save(updatedProduct).Error
	if err != nil {
		return err
	}

	return nil

}

// GetProductByPartial implements ProductRepository.
func (r *productRepository) GetProductByPartial(product *models.Product) (*models.Product, error) {
	err := r.DB.Where(product).First(product).Error
	if err != nil {
		return nil, err
	}
	return product, nil
}
