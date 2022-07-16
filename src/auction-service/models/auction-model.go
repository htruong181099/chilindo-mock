package models

import (
	"gorm.io/gorm"
	"time"
)

type Auction struct {
	gorm.Model   `json:"-"`
	Id           int       `json:"id" gorm:"primaryKey"`
	ProductId    string    `json:"productId"`
	StartingTime time.Time `json:"starting_time"`
	EndingTime   time.Time `json:"ending_time"`
	IsActive     bool      `json:"isActive" gorm:"default:false"`
	LowestBid    float32   `json:"lowestBid"`
}

func NewAuction(id int, productId string, startingTime time.Time, endingTime time.Time, isActive bool, lowestBid float32) *Auction {
	return &Auction{Id: id, ProductId: productId, StartingTime: startingTime, EndingTime: endingTime, IsActive: isActive, LowestBid: lowestBid}
}
