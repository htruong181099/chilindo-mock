package repository

import (
	"chilindo/src/product-service/dtos"
	"chilindo/src/product-service/models"
	"fmt"
	"gorm.io/gorm"
	"log"
)

type IProductRepository interface {
	CreateProduct(dto *dtos.CreateProductDTO) (*models.Product, error)
	GetProducts() (*[]models.Product, error)
	GetProductById(dto *dtos.ProductDTO) (*models.Product, error)
	UpdateProduct(dto *dtos.UpdateProductDTO) (*models.Product, error)
	DeleteProduct(dto *dtos.ProductDTO) (*models.Product, error)
}

type ProductRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) *ProductRepository {
	return &ProductRepository{db: db}
}

func (p ProductRepository) CreateProduct(dto *dtos.CreateProductDTO) (*models.Product, error) {
	fmt.Println(dto.Product)
	record := p.db.Create(&dto.Product)
	if record.Error != nil {
		log.Println("CreateProduct: Error to create repository")
		return nil, record.Error
	}
	return dto.Product, nil
} // Done

func (p ProductRepository) GetProductById(dto *dtos.ProductDTO) (*models.Product, error) {
	//TODO implement me
	panic("implement me")
}

func (p ProductRepository) GetProducts() (*[]models.Product, error) {
	var products *[]models.Product
	record := p.db.Find(&products)
	if record.Error != nil {
		log.Println("GetProducts: Error get all in package", record.Error)
		return nil, record.Error
	}
	return products, nil
}

func (p ProductRepository) UpdateProduct(dto *dtos.UpdateProductDTO) (*models.Product, error) {
	//TODO implement me
	panic("implement me")
}

func (p ProductRepository) DeleteProduct(dto *dtos.ProductDTO) (*models.Product, error) {
	//TODO implement me
	panic("implement me")
}
