package services

import (
	"chilindo/src/product-service/dtos"
	"chilindo/src/product-service/models"
	repository "chilindo/src/product-service/repository/mocks"
	"github.com/golang/mock/gomock"
	"gorm.io/gorm"
	"testing"
)

func TestProductService_CreateProduct(t *testing.T) {
	ctr := gomock.NewController(t)
	defer ctr.Finish()
	repo := repository.NewMockIProductRepository(ctr)
	ProductSvc := NewProductService(repo)

	//Mock repo
	repo.EXPECT().CreateProduct(gomock.Any()).Return(&models.Product{
		Model:       gorm.Model{},
		Id:          "",
		Name:        "",
		Price:       "",
		Description: "",
		Quantity:    0,
	}, nil).Times(1)
	var dto *dtos.CreateProductDTO
	_, err := ProductSvc.CreateProduct(dto)
	if err != nil {
		t.Fatal("Error")
	}
} //Done

func TestProductService_GetProducts(t *testing.T) {
	ctr := gomock.NewController(t)
	defer ctr.Finish()
	repo := repository.NewMockIProductRepository(ctr)
	productSCV := NewProductService(repo)

	// Mock Repository
	repo.EXPECT().GetProducts().Return(&[]models.Product{{
		Model:       gorm.Model{},
		Id:          "",
		Name:        "",
		Price:       "",
		Description: "",
		Quantity:    0,
	}}, nil).Times(1)

	_, err := productSCV.GetProducts()
	if err != nil {
		t.Fatal("Error", err)
	}
} //Done

func TestProductService_GetProductById(t *testing.T) {
	ctr := gomock.NewController(t)
	defer ctr.Finish()
	repo := repository.NewMockIProductRepository(ctr)
	productSRV := NewProductService(repo)

	//Mock repository product
	repo.EXPECT().GetProductById(gomock.Any()).Return(&models.Product{
		Model:       gorm.Model{},
		Id:          "",
		Name:        "",
		Price:       "",
		Description: "",

		Quantity: 0,
	}, nil).Times(1)

	//Testing
	var dtoTest *dtos.ProductDTO
	_, err := productSRV.GetProductById(dtoTest)
	if err != nil {
		t.Fatalf("Error")
	}
} //Done

func TestProductService_UpdateProduct(t *testing.T) {
	ctr := gomock.NewController(t)
	defer ctr.Finish()
	repo := repository.NewMockIProductRepository(ctr)
	productSrv := NewProductService(repo)

	//Mock repository
	repo.EXPECT().UpdateProduct(gomock.Any()).Return(&models.Product{
		Model:       gorm.Model{},
		Id:          "",
		Name:        "",
		Price:       "",
		Description: "",
		Quantity:    0,
	}, nil).Times(1)

	//testing
	var dtoTest *dtos.UpdateProductDTO
	_, err := productSrv.UpdateProduct(dtoTest)
	if err != nil {
		t.Fatal("Error")
	}

} //Done
