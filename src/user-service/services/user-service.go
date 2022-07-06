package services

import (
	"chilindo/dto"
	"chilindo/models"
	"chilindo/repository"
)

type IUserService interface {
	GetAddress(dto dto.GetAddressDTO) ([]models.Address, error)
	GetAddressById(dto dto.GetAddressDTO) (*models.Address, error)
}

type UserService struct {
	UserRepository    repository.IUserRepository
	AddressRepository repository.IAddressRepository
}

func (u UserService) GetAddressById(dto dto.GetAddressDTO) (*models.Address, error) {
	//TODO implement me
	//find user
	//userId := dto.userId
	//user, error := u.UserRepository.GetUserByID(userId)
	//if error != nil {
	//
	//}

	//address, err := u.AddressRepository.GetAddressById(dto)

	panic("implement me")
}

func NewUserService(userRepository repository.IUserRepository, addressRepository repository.IAddressRepository) *UserService {
	return &UserService{UserRepository: userRepository, AddressRepository: addressRepository}
}
