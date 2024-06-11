package utils

import (
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
)

type Claims struct {
	UserID uuid.UUID
	Role   string
	jwt.StandardClaims
}

func GerateJWT(userID uuid.UUID, role string) (string, error) {
	expirationTime := time.Now().Add(24 * time.Hour)
	claims := &Claims{
		UserID: userID,
		Role:   role,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	var jwtKey = []byte(os.Getenv("JWT_SECRET"))

	token := jwt.NewWithClaims(jwt.SigningMethodES256, claims)

	return token.SignedString(jwtKey)

}

func ValidateJWT(tokenString string) (*Claims, error) {
	claims := &Claims{}
	var jwtKey = []byte(os.Getenv("JWT_SECRET"))
	_, err := jwt.ParseWithClaims(tokenString, claims, func(t *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	if err != nil {
		return nil, err
	}
	return claims, nil

}
