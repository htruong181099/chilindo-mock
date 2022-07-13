package services

import "chilindo/src/auction-service/repository"

type IBidService interface {
}

type BidService struct {
	BidRepository repository.IBidRepository
}
