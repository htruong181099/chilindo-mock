package repository

import (
	"chilindo/src/order-service/dtos"
	"chilindo/src/order-service/models"
	"gorm.io/gorm"
	"log"
)

type ICartRepository interface {
	GetCart(dto *dtos.CartDTO) (*models.Cart, error)
	GetListCarts(dto *dtos.UserCartDTO) (*[]models.Cart, error)
}

type CartRepository struct {
	db *gorm.DB
}

func (c CartRepository) GetCart(dto *dtos.CartDTO) (*models.Cart, error) {
	var cart *models.Cart
	record := c.db.Where("cart_id = ? ", dto.CartId).Find(&cart)
	if record.Error != nil {
		log.Println("GetCart: Error to query", record.Error)
		return nil, record.Error
	}
	return cart, nil
}

func (c CartRepository) GetListCarts(dto *dtos.UserCartDTO) (*[]models.Cart, error) {
	var cart *[]models.Cart
	record := c.db.Where("user_id = ? ", dto.UserId).Find(&cart)
	if record.Error != nil {
		log.Println("GetCart: Error to query", record.Error)
		return nil, record.Error
	}
	return cart, nil
}

func NewCartRepository(db *gorm.DB) *CartRepository {
	return &CartRepository{db: db}
}
