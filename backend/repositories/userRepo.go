package repositories

import "github.com/XDBerry29/e-commerce-go+react/domain/models"

type UserRepository interface {
	CreateUser(user *models.User) error
	FindUserByEmail(email string) (*models.User, error)
}
