package service

import (
	"github.com/XDBerry29/e-commerce-go+react/domain/models"
	"github.com/XDBerry29/e-commerce-go+react/repositories"
)

type ProductService struct {
	productRepository repositories.ProductRepository
}

func NewProductService(repo repositories.ProductRepository) *ProductService {
	return &ProductService{
		productRepository: repo,
	}
}

func (service *ProductService) GetAllProducts() ([]models.ProductUserResponse, error) {
	products, err := service.productRepository.GetAllProducts()
	if err != nil {
		return nil, err
	}
	var userProducts []models.ProductUserResponse
	for _, product := range products {
		userProducts = append(userProducts, product.ToUserResponse())
	}

	return userProducts, nil
}

func (service *ProductService) AddProduct(name string, description string, photo string, stock int) error {
	product := models.NewProduct(name, description, photo, stock)
	if err := service.productRepository.CreateProduct(product); err != nil {
		return err
	}
	return nil
}
