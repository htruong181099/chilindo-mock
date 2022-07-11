package models

import "gorm.io/gorm"

// Make repository
type Image struct {
	gorm.Model
	Id        int     `json:"id" gorm:"primaryKey"`
	ProductId string  `json:"productId"`
	Link      string  `json:"link"`
	Product   Product `json:"-" gorm:"references:ProductId"`
}
