package middleware

import (
	"net/http"
	"strings"

	"github.com/XDBerry29/e-commerce-go+react/utils"
	"github.com/labstack/echo/v4"
)

func JWTAuth(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		authHeader := c.Request().Header.Get("Authorization")
		if authHeader == "" {
			return c.JSON(http.StatusUnauthorized, map[string]string{"error": "Missing or invalid token"})
		}

		token := strings.TrimPrefix(authHeader, "Bearer ")
		claims, err := utils.ValidateJWT(token)
		if err != nil {
			return c.JSON(http.StatusUnauthorized, map[string]string{"error": "Invalid token"})
		}

		c.Set("claims", claims)
		return next(c)
	}
}

func AdminOnly(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		claims, ok := c.Get("claims").(*utils.Claims)
		if !ok {
			return c.JSON(http.StatusUnauthorized, map[string]string{"error": "Invalid token claims"})
		}

		if claims.Role != "admin" {
			return c.JSON(http.StatusForbidden, map[string]string{"error": "Forbidden: You do not have access to this resource"})
		}

		return next(c)
	}
}
