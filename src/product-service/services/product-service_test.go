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
		Id:          0,
		Name:        "",
		Price:       "",
		Description: "",
		Quantity:    "",
	}, nil)
	var dto *dtos.CreateProductDTO
	_, err := ProductSvc.CreateProduct(dto)
	if err != nil {
		t.Fatal("Error")
	}
}