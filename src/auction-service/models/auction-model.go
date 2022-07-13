package models

import (
	"gorm.io/gorm"
	"time"
)

type Auction struct {
	gorm.Model   `json:"-"`
	Id           int       `json:"id" gorm:"primaryKey autoIncrement"`
	ProductId    string    `json:"productId"`
	StartingTime time.Time `json:"startTime"`
	EndingTime   time.Time `json:"endTime"`
	IsActive     bool      `json:"isActive" gorm:"default:false"`
	LowestBid    float32   `json:"lowestBid"`
}
