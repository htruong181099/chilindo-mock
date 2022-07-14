package repository

import (
	"chilindo/src/user-service/dto"
	"chilindo/src/user-service/models"
	"gorm.io/gorm"
	"log"
)

type IUserRepository interface {
	CreateUser(dto *dto.SignUpDTO) (*models.User, error)
	GetUserByUsername(dto *dto.SignInDTO) (*models.User, error)
	GetUserById(dto *dto.GetByUserIdDTO) (*models.User, error)
	UpdatePassword(dto *dto.UpdatePasswordDTO) (*models.User, error)
}

type UserRepository struct {
	db *gorm.DB
}

func (u UserRepository) UpdatePassword(dto *dto.UpdatePasswordDTO) (*models.User, error) {
	var user *models.User
	result := u.db.Where("id = ?", dto.UserId).Find(&user)
	if result.Error != nil {
		log.Println("UpdatePassword: Error in package repository", result.Error)
		return nil, result.Error
	}
	if checkPassErr := user.CheckPassword(dto.CurrentPassword); checkPassErr != nil {
		log.Println("UpdatePassword: Error checking password package repository", checkPassErr)
		return nil, checkPassErr
	}
	if err := user.HashPassword(dto.NewPassword); err != nil {
		log.Println("UpdatePassword: Error hash password in package repository", err)
		return nil, err
	}
	u.db.Save(&user)
	return user, nil
}

func (u UserRepository) GetUserById(dto *dto.GetByUserIdDTO) (*models.User, error) {
	var user *models.User
	var count int64
	result := u.db.Where("id = ?", dto.UserId).Find(&user).Count(&count)
	if result.Error != nil {
		log.Println("GetUserById: Error in package repository", result.Error)
		return nil, result.Error
	}
	if count == 0 {
		return nil, nil
	}
	return user, nil
}

func (u UserRepository) CreateUser(dto *dto.SignUpDTO) (*models.User, error) {
	if err := dto.User.HashPassword(dto.User.Password); err != nil {
		log.Println("CreateUser: Error in package repository", err)
		return nil, err
	}
	result := u.db.Create(&dto.User)
	if result.Error != nil {
		log.Println("CreateUser: Error in package repository", result.Error)
		return nil, result.Error
	}
	return dto.User, nil
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
