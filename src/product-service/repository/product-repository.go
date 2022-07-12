package repository

import (
	"chilindo/src/product-service/dtos"
	"chilindo/src/product-service/models"
	"fmt"
	"gorm.io/gorm"
	"log"
)

type IProductRepository interface {
	CreateProduct(dto *dtos.CreateProductDTO) (*models.Product, error) //Done
	GetProducts() (*[]models.Product, error)                           //Done
	GetProductById(dto *dtos.ProductDTO) (*models.Product, error)      //Done
	UpdateProduct(dto *dtos.UpdateProductDTO) (*models.Product, error) //Done
	DeleteProduct(dto *dtos.ProductDTO) (*models.Product, error)       //Done
	CreateOption(dto *dtos.CreateOptionDTO) (*models.Option, error)    //Done
	GetOptions(dto *dtos.ProductIdDTO) (*[]models.Option, error)
	GetOptionById(dto *dtos.OptionIdDTO) (*models.Option, error)
	UpdateOption(dto *dtos.UpdateOptionDTO) (*models.Option, error)
	DeleteOption(dto *dtos.OptionIdDTO) (*models.Option, error)
	CountProductById(dto *dtos.ProductDTO) (int64, error)
}

type ProductRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) *ProductRepository {
	return &ProductRepository{db: db}
}
func (p ProductRepository) CreateOption(dto *dtos.CreateOptionDTO) (*models.Option, error) {
	record := p.db.Create(&dto.Option)
	if record.Error != nil {
		log.Println("CreateOption: Error to create repository")
		return nil, record.Error
	}
	return dto.Option, nil
} //Done

func (p ProductRepository) GetOptions(dto *dtos.ProductIdDTO) (*[]models.Option, error) {
	var options *[]models.Option
	var count int64
	record := p.db.Where("product_id = ?", dto.ProductId).Find(&options).Count(&count)
	if record.Error != nil {
		log.Println("GetOptions : Error to get all option", record.Error)
		return nil, record.Error
	}
	if count == 0 {
		log.Println("GetOptions : Not found Options", count)
		return nil, nil
	}
	return options, nil
} //Done

func (p ProductRepository) GetOptionById(dto *dtos.OptionIdDTO) (*models.Option, error) {
	var option *models.Option
	var count int64
	record := p.db.Where("id = ?", dto.OptionId).
		Find(&option).
		Count(&count)
	if record.Error != nil {
		log.Println("GetOptionById: Error to get option in repo pkg", record.Error)
		return nil, record.Error
	}
	if count == 0 {
		log.Println("GetOptionById: Not found option", count)
		return nil, nil
	}
	return option, nil
}

func (p ProductRepository) UpdateOption(dto *dtos.UpdateOptionDTO) (*models.Option, error) {
	var count int64
	var option *models.Option
	getRecord := p.db.Where("id = ?", dto.Option.Id).
		Find(&option).
		Count(&count)
	if getRecord.Error != nil {
		log.Println("UpdateOption: Error to get in package repository", getRecord.Error)
		return nil, getRecord.Error
	}
	if count == 0 {
		return nil, nil
	}
	option.Color = dto.Option.Color
	option.Size = dto.Option.Size
	option.ProductModel = dto.Option.ProductModel

	saveRecord := p.db.Save(&option)
	if saveRecord.Error != nil {
		log.Println("UpdateOption: Error to save in package repository", saveRecord.Error)
		return nil, saveRecord.Error
	}
	return option, nil
} //Done

func (p ProductRepository) DeleteOption(dto *dtos.OptionIdDTO) (*models.Option, error) {
	//TODO implement me
	panic("implement me")
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
	var product *models.Product
	var count int64
	record := p.db.Where("id = ?", dto.ProductId).
		Find(&product).
		Count(&count)
	if record.Error != nil {
		log.Println("GetProductById: Get product by ID", record.Error)
		return nil, record.Error
	}
	if count == 0 {
		log.Println("GetProductById: Not found product", count)
		return nil, nil
	}
	return product, nil
} //Done

func (p ProductRepository) GetProducts() (*[]models.Product, error) {
	var products *[]models.Product
	record := p.db.Find(&products)
	if record.Error != nil {
		log.Println("GetProducts: Error get all in package", record.Error)
		return nil, record.Error
	}
	return products, nil
} //Done

func (p ProductRepository) UpdateProduct(dto *dtos.UpdateProductDTO) (*models.Product, error) {
	var updateProduct *models.Product
	record := p.db.Where("id = ?", dto.ProductId).Find(&updateProduct)
	if record.Error != nil {
		log.Println("UpdateProduct: Error to find product product in package repository", record.Error)
		return nil, record.Error
	}
	updateProduct = dto.Product
	recordSave := p.db.Save(&updateProduct)
	if recordSave.Error != nil {
		log.Println("UpdateProduct: Error save to update produce in package repository", recordSave.Error)
		return nil, recordSave.Error
	}
	return updateProduct, nil
} //Done

func (p ProductRepository) DeleteProduct(dto *dtos.ProductDTO) (*models.Product, error) {
	var product *models.Product
	recordFind := p.db.Where("id = ?", dto.ProductId).Delete(&product)
	if recordFind.Error != nil {
		log.Println("DeleteProduct: Error in find product to delete in package repository", recordFind.Error)
		return nil, recordFind.Error
	}
	return product, nil
} //Done

func (p ProductRepository) CountProductById(dto *dtos.ProductDTO) (int64, error) {
	var count int64
	record := p.db.Model(&models.Product{}).Where("id = ?", dto.ProductId).Count(&count)
	if record.Error != nil {
		log.Println("CountProductById: Get product by ID", record.Error)
		return count, record.Error
	}
	return count, nil
} //Done
