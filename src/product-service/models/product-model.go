package models

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Id          string `json:"id" gorm:"primaryKey"`
	Name        string `json:"name"`
	Price       string `json:"price"`
	Description string `json:"description"`
	Quantity    int    `json:"quantity"`
}
