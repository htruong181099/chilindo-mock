package services

import (
	"chilindo/src/product-service/dtos"
	"chilindo/src/product-service/models"
	"chilindo/src/product-service/repository"
	"log"
)

type IProductService interface {
	CreateProduct(dto *dtos.CreateProductDTO) (*models.Product, error) //Done
	GetProducts() (*[]models.Product, error)                           //Done
	GetProductById(dto *dtos.ProductDTO) (*models.Product, error)      //Done
	UpdateProduct(dto *dtos.UpdateProductDTO) (*models.Product, error) //Done
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
} //Done

func (p ProductService) GetProductById(dto *dtos.ProductDTO) (*models.Product, error) {
	product, err := p.ProductRepository.GetProductById(dto)
	if err != nil {
		log.Println("GetProductById: Error in get product by Id", err)
		return nil, err
	}
	return product, nil
} //Done

func (p ProductService) UpdateProduct(dto *dtos.UpdateProductDTO) (*models.Product, error) {
	product, err := p.ProductRepository.UpdateProduct(dto)
	if err != nil {
		log.Println("UpdateProduct: Error in package service", err)
		return nil, err
	}
	return product, nil
} //Done

func (p ProductService) DeleteProduct(dto *dtos.ProductDTO) (*models.Product, error) {
	product, err := p.ProductRepository.DeleteProduct(dto)
	if err != nil {
		log.Println("DeleteProduct: Error in package service", err)
		return nil, err
	}
	return product, nil
}

func NewProductService(productRepository repository.IProductRepository) *ProductService {
	return &ProductService{ProductRepository: productRepository}
}
