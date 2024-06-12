package controller

import (
	"net/http"

	"github.com/XDBerry29/e-commerce-go+react/service"
	"github.com/labstack/echo/v4"
)

type AuthController struct {
	UserService *service.UserService
}

func NewAuthController(userService *service.UserService) *AuthController {
	return &AuthController{
		UserService: userService,
	}
}

func (ctrl *AuthController) Register(c echo.Context) error {
	var input struct {
		Name     string `json:"name" validate:"required"`
		Email    string `json:"email" validate:"required,email"`
		Password string `json:"password" validate:"required,min=6"`
		Role     string `json:"role" validate:"required"` //this will be auto-assigned in a future iteration
	}
	if err := c.Bind(&input); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	if err := ctrl.UserService.RegisterUser(input.Name, input.Email, input.Password, input.Role); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "User registered successfully"})
}
