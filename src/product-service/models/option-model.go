package models

import "gorm.io/gorm"

type Option struct {
	gorm.Model
	Id           int     `json:"id" gorm:"primaryKey"`
	ProductId    string  `json:"productId"`
	Color        string  `json:"color"`
	Size         string  `json:"size"`
	ProductModel string  `json:"productModel"`
	Product      Product `json:"-" gorm:"references:ProductId"`
}
