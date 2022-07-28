package models

import "gorm.io/gorm"

type Address struct {
	gorm.Model  `json:"-"`
	Id          int    `json:"id" gorm:"primaryKey""`
	UserId      int    `json:"userId"`
	FirstName   string `json:"firstName"`
	LastName    string `json:"lastName"`
	PhoneNumber string `json:"phoneNumber"`
	Province    string `json:"province"`
	District    string `json:"district"`
	SubDistrict string `json:"subDistrict"`
	Address     string `json:"address"`
	TypeAddress string `json:"typeAddress"`
	User        User   `json:"-" gorm:"foreignKey:UserId"`
}
