package models

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string `json:"name"`
	Username string `json:"username" gorm:"unique"`
	Email    string `json:"email" gorm:"unique"`
	Password string `json:"password"`
}

func (user *User) HashPassword(passowrd string) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(passowrd), 14)
	if err != nil {
		return err
	}
	user.Password = string(bytes)
	return nil
}

func (user *User) CheckPassword(providerPassword string) error {
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(providerPassword))
	if err != nil {
		return err
	}
	return nil
}
