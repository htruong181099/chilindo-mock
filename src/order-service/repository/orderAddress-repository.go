package repository

import (
	"chilindo/src/order-service/dtos"
	"gorm.io/gorm"
)

type IOrderAddressRepository interface {
	Create(dto *dtos.CreateOrderAddressDTO)
	Get(dto *dtos.OrderAddressDTO)
}

type OrderAddressRepository struct {
	db *gorm.DB
}

func (o OrderAddressRepository) Create() {
	//TODO implement me
	panic("implement me")
}

func (o OrderAddressRepository) Get() {
	//TODO implement me
	panic("implement me")
}

func NewOrderAddressRepository(db *gorm.DB) *OrderAddressRepository {
	return &OrderAddressRepository{db: db}
}
