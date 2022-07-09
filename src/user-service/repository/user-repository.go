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
	//GetUserById(dto *dto.GetUserByIdDTO) (*models.User, error)
}

type UserRepository struct {
	db *gorm.DB
}

func (u UserRepository) CreateUser(dto *dto.SignUpDTO) (*models.User, error) {
	record := u.db.Create(dto.User)
	if record.Error != nil {
		log.Println("CreateUser : Error in package repository")
		return nil, record.Error
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
