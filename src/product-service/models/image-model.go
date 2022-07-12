package models

import "gorm.io/gorm"

type Image struct {
	gorm.Model `json:"-"`
	Id         int     `json:"id" gorm:"primaryKey"`
	ProductId  string  `json:"productId" gorm:"size:20"`
	Link       string  `json:"link"`
	Product    Product `json:"-" gorm:"foreignKey:ProductId"`
}
