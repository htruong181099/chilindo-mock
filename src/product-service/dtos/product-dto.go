package dtos

import "chilindo/src/product-service/models"

type CreateProductDTO struct {
	Product *models.Product
}

type GetProductsDTO struct {
}

type UpdateProductDTO struct {
	productId string
	Product   *models.Product
}

type ProductDTO struct {
	productId string
}

func NewCreateProductDTO(product *models.Product) *CreateProductDTO {
	return &CreateProductDTO{Product: product}
}
