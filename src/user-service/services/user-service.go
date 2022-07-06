package services

import (
	"chilindo/models"
	"chilindo/repository"
	"chilindo/services/dto"
	"log"
)

type IUserService interface {
	SignUp(user *models.User) (*models.User, error)
	SignIn(dto *dto.SignInDTO) (*models.User, error)
}

type UserService struct {
	UserRepository repository.IUserRepository
}

func NewUserService(userRepository repository.IUserRepository) *UserService {
	return &UserService{UserRepository: userRepository}
}

func (u UserService) SignUp(user *models.User) (*models.User, error) {
	userRegis, err := u.UserRepository.CreateUser(user)
	if err != nil {
		log.Println("SignUp: Error CreateUser in package service")
		return nil, err
	}
	return userRegis, nil
}

func (u UserService) SignIn(dto *dto.SignInDTO) (*models.User, error) {
	user, err := u.UserRepository.GetUserByUsername(dto.Username, dto.Password)
	if err != nil {
		log.Println("SignIn: Error GetUserByUsername in package service")
		return nil, err
	}
	return user, nil
}
