package dto

import "chilindo/src/user-service/models"

type AddressDTO struct {
	Address *models.Address
}

func NewAddressDTO(address *models.Address) *AddressDTO {
	return &AddressDTO{Address: address}
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
}
