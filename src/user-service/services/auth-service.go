package services

import (
	"chilindo/models"
	"chilindo/repository"
	"chilindo/services/dto"
	"log"
)

type IAuthService interface {
	SignUp(user *models.User) (*models.User, error)
	SignIn(dto *dto.SignInDTO) (*models.User, error)
}

type AuthService struct {
	UserRepository repository.IUserRepository
}

func NewAuthService(userRepository repository.IUserRepository) *AuthService {
	return &AuthService{UserRepository: userRepository}
}

func (u AuthService) SignUp(user *models.User) (*models.User, error) {
	userRegis, err := u.UserRepository.CreateUser(user)
	if err != nil {
		log.Println("SignUp: Error CreateUser in package service")
		return nil, err
	}
	return userRegis, nil
}

func (u AuthService) SignIn(dto *dto.SignInDTO) (*models.User, error) {
	user, err := u.UserRepository.GetUserByUsername(dto.Username, dto.Password)
	if err != nil {
		log.Println("SignIn: Error GetUserByUsername in package service")
		return nil, err
	}
	return user, nil
}
