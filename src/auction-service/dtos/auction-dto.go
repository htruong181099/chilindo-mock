package dtos

import "chilindo/src/auction-service/models"

type AuctionIdDTO struct {
	AuctionId int
}

type CreateAuctionDTO struct {
	Auction *models.Auction
}
