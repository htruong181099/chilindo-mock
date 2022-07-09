package models

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Id          string `json:"id" gorm:"primaryKey"`
	Name        string `json:"name"`
	Price       string `json:"name"`
	Description string `json:"description"`
	Quantity    string `json:"quantity"`
}
