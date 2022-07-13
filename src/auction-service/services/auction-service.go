package services

import "chilindo/src/auction-service/repository"

type IAuctionService interface {
}

type AuctionService struct {
	AuctionRepository repository.IAuctionRepository
}

func NewAuctionService(auctionRepository repository.IAuctionRepository) *AuctionService {
	return &AuctionService{AuctionRepository: auctionRepository}
}
