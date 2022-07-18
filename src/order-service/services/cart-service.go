package services

import (
	"chilindo/src/order-service/dtos"
	"chilindo/src/order-service/models"
	"chilindo/src/order-service/repository"
	"log"
)

type ICartService interface {
	GetCart(dto *dtos.CartDTO) (*models.Cart, error)
	GetListCarts(dto *dtos.UserCartDTO) (*[]models.Cart, error)
}

type CartService struct {
	CartRepository repository.ICartRepository
}

func (c CartService) GetCart(dto *dtos.CartDTO) (*models.Cart, error) {
	cart, err := c.CartRepository.GetCart(dto)
	if err != nil {
		log.Println("GetCart: Error to call repo in pkg service", err)
		return nil, err
	}
	return cart, nil
}

func (c CartService) GetListCarts(dto *dtos.UserCartDTO) (*[]models.Cart, error) {
	cart, err := c.CartRepository.GetListCarts(dto)
	if err != nil {
		log.Println("GetListCarts: Error to call repo in pkg service", err)
		return nil, err
	}
	return cart, nil
}

func NewCartService(cartRepository repository.ICartRepository) *CartService {
	return &CartService{CartRepository: cartRepository}
}
