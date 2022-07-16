package dtos

import "chilindo/src/auction-service/models"

type BidIdDTO struct {
	BidId int `json:"bidId"`
}

type CreateBidDTO struct {
	Bid *models.Bid
}
