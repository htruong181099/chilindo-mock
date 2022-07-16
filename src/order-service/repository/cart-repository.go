package repository

import (
	"chilindo/src/order-service/dtos"
	"chilindo/src/order-service/models"
	"gorm.io/gorm"
)

type ICartRepository interface {
	GetCart(dto *dtos.CartDTO) (*models.Cart, error)
	GetListCarts(dto *dtos.UserCartDTO) (*models.Cart, error)
}

type CartRepository struct {
	db *gorm.DB
}

func (c CartRepository) GetCart(dto *dtos.CartDTO) (*models.Cart, error) {
	//TODO implement me
	panic("implement me")
}

func (c CartRepository) GetListCarts(dto *dtos.UserCartDTO) (*models.Cart, error) {
	//TODO implement me
	panic("implement me")
}

func NewCartRepository(db *gorm.DB) *CartRepository {
	return &CartRepository{db: db}
}
