package models

import "gorm.io/gorm"

type Product struct {
	gorm.Model  `json:"-"`
	Id          string   `json:"id" gorm:"primaryKey size(20)"`
	Name        string   `json:"name"`
	Price       string   `json:"price"`
	Description string   `json:"description"`
	Quantity    int      `json:"quantity"`
	Options     []Option `json:"options"`
}
