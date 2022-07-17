package services

import (
	"chilindo/src/auction-service/dtos"
	"chilindo/src/auction-service/models"
	"chilindo/src/auction-service/repository"
	"log"
)

type IBidService interface {
	GetBidsOfAuction(dto *dtos.AuctionIdDTO) (*[]models.Bid, error)
	GetBidById(dto *dtos.BidIdDTO) (*[]models.Bid, error)
	CreateBid(dto *dtos.CreateBidDTO) (*models.Bid, error)
}

type BidService struct {
	BidRepository repository.IBidRepository
}

func (b BidService) GetBidsOfAuction(dto *dtos.AuctionIdDTO) (*[]models.Bid, error) {
	bid, err := b.BidRepository.GetBidsOfAuction(dto)
	if err != nil {
		log.Println("GetBidsOfAuction: Error to get bid in package service", err)
	}
	return bid, nil
} //Done

func (b BidService) GetBidById(dto *dtos.BidIdDTO) (*[]models.Bid, error) {
	bids, err := b.BidRepository.GetBidById(dto)
	if err != nil {
		log.Println("GetBidById: Error to get bid in package service", err)
	}
	return bids, nil
} //Done

func (b BidService) CreateBid(dto *dtos.CreateBidDTO) (*models.Bid, error) {
	bid, err := b.BidRepository.CreateBid(dto)
	if err != nil {
		log.Println("CreateBid: Error to create bid in package service", err)
	}
	return bid, nil
} //Done

func NewBidService(bidRepository repository.IBidRepository) *BidService {
	return &BidService{BidRepository: bidRepository}
}
