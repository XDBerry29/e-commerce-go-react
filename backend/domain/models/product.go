package models

import "github.com/google/uuid"

type Product struct {
	ID          uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"-"`
	Name        string    `json:"name" gorm:"unique"`
	Description string    `json:"description"`
	Photo       string    `json:"photo"` // should be photo url can be more than 1 photo for future iterations
	Stock       int       `json:"stock"`
}

type ProductUserResponse struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Photo       string `json:"photo"`
	IsOnStock   bool   `json:"is_on_stock"`
}

func NewProduct(name string, description string, photo string, stock int) *Product {
	return &Product{
		Name:        name,
		Description: description,
		Photo:       photo,
		Stock:       stock,
	}
}

func (p *Product) ToUserResponse() ProductUserResponse {
	return ProductUserResponse{
		Name:        p.Name,
		Description: p.Description,
		Photo:       p.Photo,
		IsOnStock:   p.Stock > 0,
	}
}
