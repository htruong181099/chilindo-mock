package models

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"log"
)

type User struct {
	gorm.Model
	Id          int    `json:"id" gorm:"primaryKey"`
	FirstName   string `json:"firstName"`
	LastName    string `json:"lastName"`
	Username    string `json:"username" gorm:"unique"`
	Password    string `json:"password"`
	Email       string `json:"email" gorm:"unique"`
	PhoneNumber string `json:"phoneNumber"`
	Gender      string `json:"gender"`
	Language    string `json:"language"`
	Role        string `json:"Role" gorm:"default:user"`
}

func (u *User) HashPassword(password string) error {
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		log.Println("HashPassword: error hash password")
		return err
	}
	u.Password = string(passwordHash)
	return nil
}

func (u *User) CheckPassword(password string) error {
	if err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password)); err != nil {
		log.Println("CheckPassword: error check password")
		return err
	}
	return nil
}
