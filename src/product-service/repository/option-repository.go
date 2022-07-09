package repository

import "gorm.io/gorm"

type IOptionRepository interface {
	//create option
}

type OptionRepository struct {
	db *gorm.DB
}

func NewOptionRepository(db *gorm.DB) *OptionRepository {
	return &OptionRepository{db: db}
}
