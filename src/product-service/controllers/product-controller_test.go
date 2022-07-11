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
} //Done

func TestProductController_GetProductById(t *testing.T) {
	ctr := gomock.NewController(t)
	defer ctr.Finish()
	srv := services.NewMockIProductService(ctr)
	productCtr := NewProductController(srv)

	//Mock service
	srv.EXPECT().GetProductById(gomock.Any()).Return(&models.Product{
		Model:       gorm.Model{},
		Id:          "",
		Name:        "",
		Price:       "",
		Description: "",
		Quantity:    0,
	}, nil).Times(1)

	req, err := http.NewRequest("GET", "/api/products:id", nil)
	if err != nil {
		t.Fatal("Error")
	}
	rr := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(rr)
	c.Request = req

	productCtr.GetProductById(c)
	//testing
	if rr.Code != http.StatusOK {
		t.Fatalf("Status expect is 200 but got %v", rr.Code)
	}

} //Done

func TestProductController_UpdateProduct(t *testing.T) {
	ctr := gomock.NewController(t)
	defer ctr.Finish()
	svr := services.NewMockIProductService(ctr)
	productCtr := NewProductController(svr)

	//Mock service
	svr.EXPECT().UpdateProduct(gomock.Any()).Return(&models.Product{
		Model:       gorm.Model{},
		Id:          "",
		Name:        "",
		Price:       "",
		Description: "",
		Quantity:    0,
	}, nil).Times(1)
	//Create mock body
	body := []byte("{}")
	//Create request
	req, err := http.NewRequest("POST", "api/products/:id", bytes.NewReader(body))
	if err != nil {
		t.Fatal("Error")
	}
	rr := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(rr)
	c.Request = req

	//testing
	productCtr.UpdateProduct(c)
	if rr.Code != http.StatusOK {
		t.Fatalf("Status epect is 200 but got %v", rr.Code)
	}
} //Done

func TestProductController_DeleteProduct(t *testing.T) {
	ctr := gomock.NewController(t)
	defer ctr.Finish()
	svr := services.NewMockIProductService(ctr)
	productCtr := NewProductController(svr)

	//Mock service
	svr.EXPECT().DeleteProduct(gomock.Any()).Return(&models.Product{
		Model:       gorm.Model{},
		Id:          "",
		Name:        "",
		Price:       "",
		Description: "",
		Quantity:    0,
	}, nil)

	req, err := http.NewRequest("DELETE", "api/products/:id", nil)
	if err != nil {
		t.Fatal("Error")
	}
	rr := httptest.NewRecorder()

	c, _ := gin.CreateTestContext(rr)
	c.Request = req
	//testing
	productCtr.DeleteProduct(c)
	if rr.Code != http.StatusOK {
		t.Fatalf("Status expect is 200 but got %v", rr.Code)
	}

} //Done
