package repository

import (
	"chilindo/src/order-service/dtos"
	"chilindo/src/order-service/models"
	"gorm.io/gorm"
)

type IOrderRepository interface {
	CreateOrder(dto *dtos.CreateOrderDTO) (*models.Order, error)
	GetOrder(dto *dtos.OrderDTO) (*models.Order, error)
	GetListOrders(dto *dtos.UserOrderDTO) (*[]models.Order, error)
}

type OrderRepository struct {
	db *gorm.DB
}

func (o OrderRepository) CreateOrder(dto *dtos.CreateOrderDTO) (*models.Order, error) {
	//TODO implement me
	panic("implement me")
}

func (o OrderRepository) GetOrder(dto *dtos.OrderDTO) (*models.Order, error) {
	//TODO implement me
	panic("implement me")
}

func (o OrderRepository) GetListOrders(dto *dtos.UserOrderDTO) (*[]models.Order, error) {
	//TODO implement me
	panic("implement me")
}

func NewOrderRepository(db *gorm.DB) *OrderRepository {
	return &OrderRepository{db: db}
}
