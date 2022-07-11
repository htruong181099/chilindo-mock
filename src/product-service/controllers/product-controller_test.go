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
	req, err := http.NewRequest("POST", "admin/product", bytes.NewReader(body))
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
}
