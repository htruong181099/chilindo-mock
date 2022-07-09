package services

import (
	"chilindo/src/product-service/dtos"
	"chilindo/src/product-service/models"
	"chilindo/src/product-service/repository"
)

type IProductService interface {
	CreateProduct(dto dtos.CreateProductDTO) (*models.Product, error)
	GetProducts(dto dtos.GetProductsDTO) (*[]models.Product, error)
	GetProductById(dto dtos.ProductDTO) (*models.Product, error)
	UpdateProduct(dto dtos.UpdateProductDTO) (*models.Product, error)
	DeleteProduct(dto dtos.ProductDTO) (*models.Product, error)
}

type ProductService struct {
	ProductRepository repository.IProductRepository
}
