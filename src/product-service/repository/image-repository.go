package repository

import (
	"chilindo/src/product-service/dtos"
	"chilindo/src/product-service/models"
	"gorm.io/gorm"
	"log"
)

type IImageRepository interface {
	CreateImage(dto *dtos.ImageDTO) (*models.Image, error) //Done
	GetImage(dto *dtos.ImageByIdDTO) (*[]models.Image, error)
	GetImageById(dto *dtos.ImageByIdDTO) (*models.Image, error)
	UpdateImage(dto *dtos.ImageUpdateDTO) (*models.Image, error)
	DeleteImage(dto *dtos.ImageByIdDTO) (*models.Image, error)
}

type ImageRepository struct {
	db *gorm.DB
}

func (i ImageRepository) CreateImage(dto *dtos.ImageDTO) (*models.Image, error) {
	dto.Image.ProductId = dto.ProductId
	record := i.db.Create(&dto.Image)
	if record.Error != nil {
		log.Println("CreateImage: Error to create Image", record.Error)
		return nil, record.Error
	}
	return dto.Image, nil
}

func (i ImageRepository) GetImage(dto *dtos.ImageByIdDTO) (*[]models.Image, error) {
	//TODO implement me
	panic("implement me")
}

func (i ImageRepository) GetImageById(dto *dtos.ImageByIdDTO) (*models.Image, error) {
	//TODO implement me
	panic("implement me")
}

func (i ImageRepository) UpdateImage(dto *dtos.ImageUpdateDTO) (*models.Image, error) {
	//TODO implement me
	panic("implement me")
}

func (i ImageRepository) DeleteImage(dto *dtos.ImageByIdDTO) (*models.Image, error) {
	//TODO implement me
	panic("implement me")
}

func NewImageRepository(db *gorm.DB) *ImageRepository {
	return &ImageRepository{db: db}
}
