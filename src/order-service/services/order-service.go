package services

import "chilindo/src/order-service/repository"

type IOrderService interface {
}

type OrderService struct {
	OrderRepository repository.IOrderRepository
}

func NewOrderService(orderRepository repository.IOrderRepository) *OrderService {
	return &OrderService{OrderRepository: orderRepository}
}
