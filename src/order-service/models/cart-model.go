package models

import "gorm.io/gorm"

type Cart struct {
	gorm.Model `json:"-"`
	UserId     int     `json:"userId" gorm:"index"`
	AuctionId  int     `json:"auctionId" gorm:"unique"`
	BidId      int     `json:"bidId" gorm:"unique"`
	Sold       bool    `json:"sold" gorm:"default:false"`
	Saved      bool    `json:"saved" gorm:"default:false"`
	Amount     float32 `json:"amount"`
}
