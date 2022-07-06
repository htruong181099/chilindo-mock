package services

import (
	"chilindo/repository"
)

type IUserService interface {
}

type UserService struct {
	UserRepository repository.IUserRepository
}

func NewUserService(userRepository repository.IUserRepository) *UserService {
	return &UserService{UserRepository: userRepository}
}
