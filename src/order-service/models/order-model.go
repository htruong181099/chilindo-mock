package models

import "gorm.io/gorm"

type Order struct {
	gorm.Model     `json:"-"`
	Id             int          `json:"id" gorm:"primaryKey"`
	UserId         int          `json:"userId" gorm:"index"`
	CartId         int          `json:"cartId" gorm:"unique"`
	OrderAddressId int          `json:"orderAddressId" gorm:"unique"`
	Amount         float32      `json:"amount"`
	Cart           Cart         `json:"-" gorm:"foreignKey:CartId"`
	OrderAddress   OrderAddress `json:"-" gorm:"foreignKey:OrderAddressId"`
}
