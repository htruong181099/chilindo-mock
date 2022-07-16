package dtos

import "chilindo/src/order-service/models"

type OrderDTO struct {
	OrderId int
}

type CreateOrderDTO struct {
	Order *models.Order
}

type UserOrderDTO struct {
	UserId int
}
