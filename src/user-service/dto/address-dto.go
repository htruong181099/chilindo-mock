package dto

import "chilindo/models"

type CreateAddressDTO struct {
	Address *models.Address
}

type GetAddressDTO struct {
	UserId int
}

type GetAddressByIdDTO struct {
	AddressId int
}

type UpdateAddressDTO struct {
	//data transfer object -> request body
	AddressId int
	UserId    int
}
