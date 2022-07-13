package repository

import (
	"chilindo/src/auction-service/dtos"
	"chilindo/src/auction-service/models"
	"gorm.io/gorm"
)

type IAuctionRepository interface {
	GetAuctions() (*[]models.Auction, error)
	GetAuctionById(dto dtos.AuctionIdDTO) (*models.Auction, error)
	CreateAuction(dto dtos.CreateAuctionDTO)
}

type AuctionRepository struct {
	db *gorm.DB
}

func (a AuctionRepository) GetAuctions() (*[]models.Auction, error) {
	//TODO implement me
	panic("implement me")
}

func (a AuctionRepository) GetAuctionById(dto dtos.AuctionIdDTO) (*models.Auction, error) {
	//TODO implement me
	panic("implement me")
}

func (a AuctionRepository) CreateAuction(dto dtos.CreateAuctionDTO) {
	//TODO implement me
	panic("implement me")
}

func NewAuctionRepository(db *gorm.DB) *AuctionRepository {
	return &AuctionRepository{db: db}
}
