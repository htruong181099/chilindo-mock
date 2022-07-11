package dtos

import "chilindo/src/product-service/models"

type CreateProductDTO struct {
	Product *models.Product
}

type UpdateProductDTO struct {
	ProductId string
	Product   *models.Product
}

type ProductDTO struct {
	ProductId string
}

func NewCreateProductDTO(product *models.Product) *CreateProductDTO {
	return &CreateProductDTO{Product: product}
}
