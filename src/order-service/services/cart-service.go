package services

import "chilindo/src/order-service/repository"

type ICartService interface {
}

type CartService struct {
	CartRepository repository.ICartRepository
}

func NewCartService(cartRepository repository.ICartRepository) *CartService {
	return &CartService{CartRepository: cartRepository}
}
