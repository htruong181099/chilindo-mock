package services

import (
	"chilindo/dto"
	"chilindo/models"
	"chilindo/repository"
	"log"
)

type IUserService interface {
	GetAddress(dto *dto.GetAddressDTO) (*[]models.Address, error)
	GetAddressById(dto *dto.GetAddressByIdDTO) (*models.Address, error)
	CreateAddress(dto *dto.CreateAddressDTO) (*models.Address, error)
	DeletedAddressById(dto *dto.DeleteAddressByIdDTO) (*models.Address, error)
}

type UserService struct {
	UserRepository    repository.IUserRepository
	AddressRepository repository.IAddressRepository
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
	//TODO implement me
	//find user
	//userId := dto.userId
	//user, error := u.UserRepository.GetUserByID(userId)
	//if error != nil {
	//
	//}
	//address, err := u.AddressRepository.GetAddressById(dto)
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

func (u *UserService) DeletedAddress(dto *dto.DeleteAddressByIdDTO) (*models.Address, error) {
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
