package services

import (
	"chilindo/dto"
	"chilindo/models"
	"chilindo/repository"
	"log"
)

type IAuthService interface {
	SignUp(dto *dto.SignUpDTO) (*models.User, error)
	SignIn(dto *dto.SignInDTO) (*models.User, error)
}

type AuthService struct {
	UserRepository repository.IUserRepository
}

func NewAuthService(userRepository repository.IUserRepository) *AuthService {
	return &AuthService{UserRepository: userRepository}
}

func (u AuthService) SignUp(dto *dto.SignUpDTO) (*models.User, error) {
	userRegis, err := u.UserRepository.CreateUser(dto)
	if err != nil {
		log.Println("SignUp: Error CreateUser in package service")
		return nil, err
	}
	return userRegis, nil
}

func (u AuthService) SignIn(dto *dto.SignInDTO) (*models.User, error) {
	user, err := u.UserRepository.GetUserByUsername(dto)
	if err != nil {
		log.Println("SignIn: Error GetUserByUsername in package service")
		return nil, err
	}
	return user, nil
}
