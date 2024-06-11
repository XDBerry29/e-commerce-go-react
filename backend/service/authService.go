package service

import (
	"github.com/XDBerry29/e-commerce-go+react/domain/models"
	"github.com/XDBerry29/e-commerce-go+react/repositories"
	"github.com/XDBerry29/e-commerce-go+react/utils"
)

type UserService struct {
	userRepository repositories.UserRepository
}

func NewUserService(userRepository repositories.UserRepository) *UserService {
	return &UserService{
		userRepository: userRepository,
	}
}

func (userService *UserService) RegisterUser(name string, email string, password string, role string) error {
	hashPassword, err := utils.HashPassword(password)
	if err != nil {
		return err
	}

	user := models.NewUser(name, email, hashPassword, role)

	err = userService.userRepository.CreateUser(user)
	if err != nil {
		return err
	}

	return nil

}
