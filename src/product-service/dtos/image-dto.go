package dtos

import "chilindo/src/product-service/models"

type ImageDTO struct {
	Image     *models.Image
	ProductId string
}

func NewImageDTO(image *models.Image, productId string) *ImageDTO {
	return &ImageDTO{Image: image, ProductId: productId}
}

type ImageByIdDTO struct {
	ImageId   int
	ProductId string
}

type ImageUpdateDTO struct {
	Image     *models.Image
	ImageId   int
	ProductId string
}
