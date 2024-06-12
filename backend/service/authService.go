package service

import (
	"github.com/XDBerry29/e-commerce-go+react/domain/models"
	"github.com/XDBerry29/e-commerce-go+react/repositories"
	"github.com/XDBerry29/e-commerce-go+react/utils"
)

type AuthService struct {
	userRepository repositories.UserRepository
}

func NewAuthService(userRepository repositories.UserRepository) *AuthService {
	return &AuthService{
		userRepository: userRepository,
	}
}

func (authService *AuthService) RegisterUser(name string, email string, password string, role string) error {
	if err := utils.ValidatePassword(password); err != nil {
		return err
	}
	hashPassword, err := utils.HashPassword(password)
	if err != nil {
		return err
	}

	user := models.NewUser(name, email, hashPassword, role)

	err = authService.userRepository.CreateUser(user)
	if err != nil {
		return err
	}

	return nil

}

func (authService *AuthService) AuthUser(email string, password string) (string, error) {
	user, err := authService.userRepository.FindUserByEmail(email)
	if err != nil {
		return "", err
	}
	if err = user.CheckPassword(password); err != nil {
		return "", err
	}

	jwt, err := utils.GerateJWT(user.GetID(), user.GetRole())
	if err != nil {
		return "", err
	}

	return jwt, nil
}
