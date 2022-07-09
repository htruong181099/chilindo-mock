package models

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Id          int    `json:"id" gorm:"primaryKey"`
	Name        string `json:"name"`
	Price       string `json:"price"`
	Description string `json:"description"`
	Quantity    string `json:"quantity"`
}
