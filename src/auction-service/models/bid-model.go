package models

import (
	"gorm.io/gorm"
	"time"
)

type Bid struct {
	gorm.Model `json:"-"`
	Id         int       `json:"id" gorm:"primaryKey"`
	BidderId   int       `json:"bidderId"`
	AuctionId  int       `json:"auctionId"`
	BidTime    time.Time `json:"bidTime" gorm:"default"`
	Amount     float32   `json:"amount"`
	Auction    Auction   `json:"-" gorm:"foreignKey:AuctionId"`
}
