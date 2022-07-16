package dtos

import "chilindo/src/order-service/models"

type CreateOrderAddressDTO struct {
	OrderAddress *models.OrderAddress
}

type OrderAddressDTO struct {
	OrderAddressId int
}
