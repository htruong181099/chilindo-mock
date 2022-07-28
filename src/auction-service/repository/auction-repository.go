package repository

import (
	"chilindo/src/auction-service/dtos"
	"chilindo/src/auction-service/models"
	"gorm.io/gorm"
	"log"
)

type IAuctionRepository interface {
	GetAuctions() (*[]models.Auction, error)
	GetAuctionById(dto *dtos.AuctionIdDTO) (*models.Auction, error)
	CreateAuction(dto *dtos.CreateAuctionDTO) (*models.Auction, error)
}

type AuctionRepository struct {
	db *gorm.DB
}

func (a AuctionRepository) GetAuctions() (*[]models.Auction, error) {
	var auction *[]models.Auction
	record := a.db.Find(&auction)
	if record.Error != nil {
		log.Println("GetAuctions: Error to get auth", record.Error)
		return nil, record.Error
	}
	return auction, nil
}

func (a AuctionRepository) GetAuctionById(dto *dtos.AuctionIdDTO) (*models.Auction, error) {
	var auction *models.Auction
	var count int64
	record := a.db.Where("id = ?", dto.AuctionId).Find(&auction).Count(&count)
	if record.Error != nil {
		log.Println("GetAuctionById: Error to get auth", record.Error)
		return nil, record.Error
	}
	if count == 0 {
		return nil, nil
	}
	return auction, nil
}

func (a AuctionRepository) CreateAuction(dto *dtos.CreateAuctionDTO) (*models.Auction, error) {
	var auth *models.Auction
	auth = dto.Auction
	record := a.db.Create(&auth)
	if record.Error != nil {
		log.Println("CreateAuction: Error to create auth", record.Error)
		return nil, record.Error
	}
	return auth, nil
} //Done

func NewAuctionRepository(db *gorm.DB) *AuctionRepository {
	return &AuctionRepository{db: db}
}
