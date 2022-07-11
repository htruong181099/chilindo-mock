package services

import (
	"chilindo/src/user-service/dto"
	"chilindo/src/user-service/models"
	"chilindo/src/user-service/repository"
	"log"
)

type IUserService interface {
	GetUserById(dto *dto.GetByUserIdDTO) (*models.User, error)
	UpdatePassword(dto *dto.UpdatePasswordDTO) (*models.User, error)
	GetAddress(dto *dto.GetAddressDTO) (*[]models.Address, error)
	GetAddressById(dto *dto.GetAddressByIdDTO) (*models.Address, error)
	CreateAddress(dto *dto.CreateAddressDTO) (*models.Address, error)
	DeletedAddressById(dto *dto.DeleteAddressByIdDTO) (*models.Address, error)
}

type UserService struct {
	UserRepository    repository.IUserRepository
	AddressRepository repository.IAddressRepository
}

func (u *UserService) UpdatePassword(dto *dto.UpdatePasswordDTO) (*models.User, error) {
	user, repoErr := u.UserRepository.UpdatePassword(dto)
	if repoErr != nil {
		log.Println("ChangePassword: error in package service", repoErr)
		return nil, repoErr
	}
	return user, nil
}

func (u *UserService) GetUserById(dto *dto.GetByUserIdDTO) (*models.User, error) {
	user, repoErr := u.UserRepository.GetUserById(dto)
	if repoErr != nil {
		log.Println("GetUserById: Error Get User in package Service", repoErr)
		return nil, repoErr
	}
	return user, nil
}

func (u *UserService) CreateAddress(dto *dto.CreateAddressDTO) (*models.Address, error) {
	address, err := u.AddressRepository.CreateAddress(dto)
	if err != nil {
		log.Println("CreateAddress: Error Create address in package service", err)
		return nil, err
	}
	return address, nil
}

func (u *UserService) GetAddressById(dto *dto.GetAddressByIdDTO) (*models.Address, error) {
	address, err := u.AddressRepository.GetAddressById(dto)
	if err != nil {
		log.Println("GetAddressById: Error in get address by id in package uer-service", err)
		return nil, err
	}
	return address, nil
}

func (u *UserService) GetAddress(dto *dto.GetAddressDTO) (*[]models.Address, error) {
	address, err := u.AddressRepository.GetAddress(dto)
	if err != nil {
		log.Println("GetAddress: Error GetAddress in package user-service", err)
		return nil, err
	}
	return address, nil
}

func (u *UserService) DeletedAddressById(dto *dto.DeleteAddressByIdDTO) (*models.Address, error) {
	address, err := u.AddressRepository.DeleteAddressById(dto)
	if err != nil {
		log.Println("DeletedAddress: Error Delete Address in package service")
		return nil, err
	}
	return address, nil
}
func NewUserService(userRepository repository.IUserRepository, addressRepository repository.IAddressRepository) *UserService {
	return &UserService{UserRepository: userRepository, AddressRepository: addressRepository}
}
