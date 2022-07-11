package services

import (
	"chilindo/src/product-service/dtos"
	"chilindo/src/product-service/models"
	"chilindo/src/product-service/repository"
	"log"
)

type IImageService interface {
	CreateImage(dto *dtos.ImageDTO) (*models.Image, error) //Done
	GetImage(dto *dtos.ImageByIdDTO) (*[]models.Image, error)
	GetImageById(dto *dtos.ImageByIdDTO) (*models.Image, error)
	UpdateImage(dto *dtos.ImageUpdateDTO) (*models.Image, error)
	DeleteImage(dto *dtos.ImageByIdDTO) (*models.Image, error)
}

type ImageService struct {
	ImageRepository repository.IImageRepository
}

func (i ImageService) CreateImage(dto *dtos.ImageDTO) (*models.Image, error) {
	image, err := i.ImageRepository.CreateImage(dto)
	if err != nil {
		log.Println("CreateImage: Error to create Image in package service", err)
		return nil, err
	}
	return image, nil
}

func (i ImageService) GetImage(dto *dtos.ImageByIdDTO) (*[]models.Image, error) {
	//TODO implement me
	panic("implement me")
}

func (i ImageService) GetImageById(dto *dtos.ImageByIdDTO) (*models.Image, error) {
	//TODO implement me
	panic("implement me")
}

func (i ImageService) UpdateImage(dto *dtos.ImageUpdateDTO) (*models.Image, error) {
	//TODO implement me
	panic("implement me")
}

func (i ImageService) DeleteImage(dto *dtos.ImageByIdDTO) (*models.Image, error) {
	//TODO implement me
	panic("implement me")
}

func NewImageService(imageRepository repository.IImageRepository) *ImageService {
	return &ImageService{ImageRepository: imageRepository}
}
