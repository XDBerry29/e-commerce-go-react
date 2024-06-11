package models

import (
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID       uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	Name     string    `json:"name"`
	Email    string    `json:"email" gorm:"unique"`
	Password string    `json:"-"`
	Role     string    `json:"role"`
}

func NewUser(name string, email string, password string, role string) *User {
	return &User{
		Name:     name,
		Email:    email,
		Password: password,
		Role:     role,
	}
}

func (user *User) GetID() uuid.UUID {
	return user.ID
}

func (user *User) GetPassword() string {
	return user.Password
}

func (user *User) GetRole() string {
	return user.Role
}

func (user *User) CheckPassword(password string) error {
	return bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
}
