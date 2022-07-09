package repository

import "gorm.io/gorm"

type IImageRepository interface {
}

type ImageRepository struct {
	db *gorm.DB
}

func NewImageRepository(db *gorm.DB) *ImageRepository {
	return &ImageRepository{db: db}
}
