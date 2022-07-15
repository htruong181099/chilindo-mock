package services

import (
	"chilindo/src/product-service/dtos"
	"chilindo/src/product-service/models"
	"chilindo/src/product-service/repository"
	"errors"
	"log"
)

type IProductService interface {
	//product
	CreateProduct(dto *dtos.CreateProductDTO) (*models.Product, error) //Done
	GetProducts() (*[]models.Product, error)                           //Done
	GetProductById(dto *dtos.ProductDTO) (*models.Product, error)      //Done
	UpdateProduct(dto *dtos.UpdateProductDTO) (*models.Product, error) //Done
	DeleteProduct(dto *dtos.ProductDTO) (*models.Product, error)       //Done
	//option
	CreateOption(dto *dtos.CreateOptionDTO) (*models.Option, error) //Done
	GetOptions(dto *dtos.ProductIdDTO) (*[]models.Option, error)
	GetOptionById(dto *dtos.OptionIdDTO) (*models.Option, error)
	UpdateOption(dto *dtos.UpdateOptionDTO) (*models.Option, error)
	DeleteOption(dto *dtos.OptionIdDTO) (*models.Option, error)
}

type ProductService struct {
	ProductRepository repository.IProductRepository
}

func (p ProductService) CreateOption(dto *dtos.CreateOptionDTO) (*models.Option, error) {
	var proDTO dtos.ProductDTO
	proDTO.ProductId = dto.Option.ProductId
	prod, prodErr := p.ProductRepository.GetProductById(&proDTO)
	if prodErr != nil {
		log.Println("CreateOption: Error not found product to create option", prodErr)
		return nil, prodErr
	}
	if prod == nil {
		log.Println("CreateOption: Error not found product to create option", prodErr)
		return nil, errors.New("not found product")
	}
	option, err := p.ProductRepository.CreateOption(dto)
	if err != nil {
		log.Println("CreateOption: Error to create option", err)
		return nil, err
	}
	return option, nil
} //Done

func (p ProductService) GetOptions(dto *dtos.ProductIdDTO) (*[]models.Option, error) {
	options, err := p.ProductRepository.GetOptions(dto)
	if err != nil {
		log.Println("GetOptions: Error get options", err)
		return nil, err
	}
	return options, nil
} //Done

func (p ProductService) GetOptionById(dto *dtos.OptionIdDTO) (*models.Option, error) {
	option, err := p.ProductRepository.GetOptionById(dto)
	if err != nil {
		log.Println("GetOptionById: Error get option", err)
		return nil, err
	}
	return option, nil
} //Done

func (p ProductService) UpdateOption(dto *dtos.UpdateOptionDTO) (*models.Option, error) {
	option, err := p.ProductRepository.UpdateOption(dto)
	if err != nil {
		log.Println("UpdateOption: Error call repo")
		return nil, err
	}
	return option, nil
}

func (p ProductService) DeleteOption(dto *dtos.OptionIdDTO) (*models.Option, error) {
	option, err := p.ProductRepository.DeleteOption(dto)
	if err != nil {
		log.Println("DeleteOption: Error delete option", err)
		return nil, err
	}
	return option, nil
} //Done

//Product Repository

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
} //Done

func NewProductService(productRepository repository.IProductRepository) *ProductService {
	return &ProductService{ProductRepository: productRepository}
}
