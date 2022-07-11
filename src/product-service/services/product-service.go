package services

import (
	"chilindo/src/product-service/dtos"
	"chilindo/src/product-service/models"
	"chilindo/src/product-service/repository"
	"log"
)

type IProductService interface {
	CreateProduct(dto *dtos.CreateProductDTO) (*models.Product, error)
	GetProducts() (*[]models.Product, error)
	GetProductById(dto *dtos.ProductDTO) (*models.Product, error)
	UpdateProduct(dto *dtos.UpdateProductDTO) (*models.Product, error)
	DeleteProduct(dto *dtos.ProductDTO) (*models.Product, error)
}

type ProductService struct {
	ProductRepository repository.IProductRepository
}

func (p ProductService) CreateProduct(dto *dtos.CreateProductDTO) (*models.Product, error) {
	createProduct, err := p.ProductRepository.CreateProduct(dto)
	if err != nil {
		log.Println("CreateProduct: Error in create product in service", err)
		return nil, err
	}
	return createProduct, nil
} //Done

func (p ProductService) GetProducts() (*[]models.Product, error) {
	products, err := p.ProductRepository.GetProducts()
	if err != nil {
		log.Println("GetProducts : Error get products in package service", err)
	}
	return products, nil
}

func (p ProductService) GetProductById(dto *dtos.ProductDTO) (*models.Product, error) {
	//TODO implement me
	panic("implement me")
}

func (p ProductService) UpdateProduct(dto *dtos.UpdateProductDTO) (*models.Product, error) {
	//TODO implement me
	panic("implement me")
}

func (p ProductService) DeleteProduct(dto *dtos.ProductDTO) (*models.Product, error) {
	//TODO implement me
	panic("implement me")
}

func NewProductService(productRepository repository.IProductRepository) *ProductService {
	return &ProductService{ProductRepository: productRepository}
}
