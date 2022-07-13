package repository

import (
	"chilindo/src/auction-service/dtos"
	"chilindo/src/auction-service/models"
	"gorm.io/gorm"
)

type IBidRepository interface {
	GetBidsOfAuction(dto dtos.AuctionIdDTO) (*[]models.Auction, error)
	GetBidById(dto dtos.BidIdDTO) (*[]models.Auction, error)
	CreateBid(dto dtos.CreateBidDTO)
}

type BidRepository struct {
	db *gorm.DB
}

func NewBidRepository(db *gorm.DB) *BidRepository {
	return &BidRepository{db: db}
}
