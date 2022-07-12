package models

import "gorm.io/gorm"

type Option struct {
	gorm.Model   `json:"-"`
	Id           int     `json:"id" gorm:"primaryKey"`
	ProductId    string  `json:"productId" gorm:"size:20"`
	Color        string  `json:"color"`
	Size         string  `json:"size"`
	ProductModel string  `json:"productModel"`
	Product      Product `json:"-" gorm:"foreignKey:ProductId;references:Id"`
}
