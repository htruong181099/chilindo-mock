package repository

import (
	"chilindo/src/auction-service/dtos"
	"chilindo/src/auction-service/models"
	"gorm.io/gorm"
	"log"
)

type IBidRepository interface {
	GetBidsOfAuction(dto *dtos.AuctionIdDTO) (*[]models.Bid, error) //Done
	GetBidById(dto *dtos.BidIdDTO) (*models.Bid, error)             //Done
	CreateBid(dto *dtos.CreateBidDTO) (*models.Bid, error)          //Done
}

type BidRepository struct {
	db *gorm.DB
}

func (b BidRepository) GetBidsOfAuction(dto *dtos.AuctionIdDTO) (*[]models.Bid, error) {
	var bid *[]models.Bid

	record := b.db.Where("auction_id = ?", dto.AuctionId).Find(&bid)
	if record.Error != nil {
		log.Println("GetBidsOfAuction: Error to find bid", record.Error)
		return nil, record.Error
	}
	return bid, nil
} //Done

func (b BidRepository) GetBidById(dto *dtos.BidIdDTO) (*models.Bid, error) {
	var bid *models.Bid
	var count int64
	record := b.db.Where("id = ?", dto.BidId).Find(&bid).Count(&count)
	if record.Error != nil {
		log.Println("GetBidById: Error to find bid", record.Error)
		return nil, record.Error
	}
	if count == 0 {
		return nil, nil
	}
	return bid, nil
} //Done

func (b BidRepository) CreateBid(dto *dtos.CreateBidDTO) (*models.Bid, error) {
	var bid *models.Bid
	bid = dto.Bid
	record := b.db.Create(&bid)
	if record.Error != nil {
		log.Println("CreateBid: Error to create bid", record.Error)
		return nil, record.Error
	}
	return bid, nil
} //Done

func NewBidRepository(db *gorm.DB) *BidRepository {
	return &BidRepository{db: db}
}
