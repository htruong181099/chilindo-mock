package repository

import (
	"chilindo/dto"
	"chilindo/models"
	"fmt"
	"gorm.io/gorm"
	"log"
)

type IUserRepository interface {
	CreateUser(user *models.User) (*models.User, error)
	GetUserByUsername(dto *dto.SignInDTO) (*models.User, error)
	GetUserById(dto *dto.GetUserByIdDTO) (*models.User, error)
}

type UserRepository struct {
	db *gorm.DB
}

func (u UserRepository) GetUserById(dto *dto.GetUserByIdDTO) (*models.User, error) {
	//TODO implement me
	panic("implement me")
}

func (u UserRepository) CreateUser(user *models.User) (*models.User, error) {
	fmt.Println(user.Password)
	if err := user.HashPassword(user.Password); err != nil {
		log.Println("CreateUser: Error in package repository")
		return nil, err
	}
	fmt.Println(user.Password)
	if err := u.db.Create(user); err != nil {
		log.Println("CreateUser: Error in package repository")
		return nil, err.Error
	}
	return user, nil
}

func (u UserRepository) GetUserByUsername(dto *dto.SignInDTO) (*models.User, error) {
	username := dto.Username
	password := dto.Password
	var user *models.User
	result := u.db.Where("username = ?", username).Find(&user)
	if result.Error != nil {
		log.Println("GetUserByUsername: Error find username in package repository", result.Error)
		return nil, result.Error
	}
	//Compose
	if err := user.CheckPassword(password); err != nil {
		log.Println("GetUserByUsername: Error in check password package repository")
		return nil, err
	}
	return user, nil
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}
