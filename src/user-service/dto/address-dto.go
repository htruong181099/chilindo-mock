package dto

import "chilindo/models"

type CreateAddressDTO struct {
	Address *models.Address
	UserId  int
}

type GetAddressDTO struct {
	UserId int
}

type GetAddressByIdDTO struct {
	AddressId int
	UserId    int
}

type UpdateAddressDTO struct {
	Address *models.Address
	UserId  int
}

type DeleteAddressByIdDTO struct {
	AddressId int
	UserId    int
}
