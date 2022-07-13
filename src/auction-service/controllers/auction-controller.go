package controllers

import "chilindo/src/auction-service/services"

type IAuctionController interface {
}

type AuctionController struct {
	AuctionService services.IAuctionService
}

func NewAuctionController(auctionService services.IAuctionService) *AuctionController {
	return &AuctionController{AuctionService: auctionService}
}
