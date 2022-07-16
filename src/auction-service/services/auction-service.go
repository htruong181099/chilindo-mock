package services

import (
	"chilindo/src/auction-service/dtos"
	"chilindo/src/auction-service/models"
	"chilindo/src/auction-service/repository"
	"log"
)

type IAuctionService interface {
	GetAuctions() (*[]models.Auction, error)
	GetAuctionById(dto *dtos.AuctionIdDTO) (*models.Auction, error)
	CreateAuction(dto *dtos.CreateAuctionDTO) (*models.Auction, error)
}

type AuctionService struct {
	AuctionRepository repository.IAuctionRepository
}

func (a AuctionService) GetAuctions() (*[]models.Auction, error) {
	//TODO implement me
	panic("implement me")
}

func (a AuctionService) GetAuctionById(dto *dtos.AuctionIdDTO) (*models.Auction, error) {
	//TODO implement me
	panic("implement me")
}

func (a AuctionService) CreateAuction(dto *dtos.CreateAuctionDTO) (*models.Auction, error) {
	auth, err := a.AuctionRepository.CreateAuction(dto)
	if err != nil {
		log.Println("Error call to repository in package service", err)
		return nil, err
	}
	return auth, nil
} //Done

func NewAuctionService(auctionRepository repository.IAuctionRepository) *AuctionService {
	return &AuctionService{AuctionRepository: auctionRepository}
}
