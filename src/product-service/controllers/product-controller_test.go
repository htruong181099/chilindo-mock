package controllers

import (
	"bytes"
	"chilindo/src/product-service/models"
	services "chilindo/src/product-service/services/mocks"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"gorm.io/gorm"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestProductController_CreateProduct(t *testing.T) {
	ctr := gomock.NewController(t)
	defer ctr.Finish()
	mockSvc := services.NewMockIProductService(ctr)
	productCtr := NewProductController(mockSvc)

	//Mock service
	mockSvc.EXPECT().CreateProduct(gomock.Any()).Return(&models.Product{
		Model:       gorm.Model{},
		Id:          "00-00",
		Name:        "",
		Price:       "",
		Description: "",
		Quantity:    1,
	}, nil).Times(1)

	// Create Body request
	body := []byte("{}")
	//request
	req, err := http.NewRequest("POST", "api/products", bytes.NewReader(body))
	if err != nil {
		t.Fatal("Error")
		return
	}
	rr := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(rr)
	c.Request = req
	productCtr.CreateProduct(c)

	if rr.Code != http.StatusCreated {
		t.Fatalf("Status epect is 201 code but got %v", rr.Code)
	}
} //Done

func TestProductController_GetProducts(t *testing.T) {
	ctr := gomock.NewController(t)
	defer ctr.Finish()
	svc := services.NewMockIProductService(ctr)
	productCtr := NewProductController(svc)

	//Mock service
	svc.EXPECT().GetProducts().Return(&[]models.Product{{
		Model:       gorm.Model{},
		Id:          "",
		Name:        "",
		Price:       "",
		Description: "",
		Quantity:    0,
	}}, nil).Times(1)

	//Create mock request
	req, err := http.NewRequest("GET", "api/products", nil)
	if err != nil {
		t.Fatalf("Error")
	}
	rr := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(rr)
	c.Request = req

	productCtr.GetProducts(c)

	if rr.Code != http.StatusOK {
		t.Fatalf("Status epect is 201 code but got %v", rr.Code)
	}
}
