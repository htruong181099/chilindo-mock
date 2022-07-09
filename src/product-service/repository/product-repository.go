package repository

import (
	"chilindo/src/product-service/dtos"
	"chilindo/src/product-service/models"
	"gorm.io/gorm"
)

type IProductRepository interface {
	CreateProduct(dto dtos.CreateProductDTO) (*models.Product, error)
	GetProducts(dto dtos.GetProductsDTO) (*[]models.Product, error)
	GetProductById(dto dtos.ProductDTO) (*models.Product, error)
	UpdateProduct(dto dtos.UpdateProductDTO) (*models.Product, error)
	DeleteProduct(dto dtos.ProductDTO) (*models.Product, error)
}

type ProductRepository struct {
	db *gorm.DB
}

func (p ProductRepository) CreateProduct(dto dtos.CreateProductDTO) (*models.Product, error) {
	//TODO implement me
	panic("implement me")
}

func (p ProductRepository) GetProducts(dto dtos.GetProductsDTO) (*[]models.Product, error) {
	//TODO implement me
	panic("implement me")
}

func (p ProductRepository) GetProductById(dto dtos.ProductDTO) (*models.Product, error) {
	//TODO implement me
	panic("implement me")
}

func (p ProductRepository) UpdateProduct(dto dtos.UpdateProductDTO) (*models.Product, error) {
	//TODO implement me
	panic("implement me")
}

func (p ProductRepository) DeleteProduct(dto dtos.ProductDTO) (*models.Product, error) {
	//TODO implement me
	panic("implement me")
}

func NewProductRepository(db *gorm.DB) *ProductRepository {
	return &ProductRepository{db: db}
}
