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
	"strings"
	"testing"
)

// Test Option Controller

func TestProductController_CreateOption(t *testing.T) {
	ctr := gomock.NewController(t)
	defer ctr.Finish()
	mockSvc := services.NewMockIProductService(ctr)
	testCtl := NewProductController(mockSvc)

	//mock service
	mockSvc.EXPECT().CreateOption(gomock.Any()).Return(&models.Option{
		Model:        gorm.Model{},
		Id:           0,
		ProductId:    "",
		Color:        "",
		Size:         "",
		ProductModel: "",
		Product:      models.Product{},
	}, nil).Times(1)

	body := []byte("{}")
	req, err := http.NewRequest("POST", "api/products/:productId/options", bytes.NewBuffer(body))

	if err != nil {
		t.Fatal("Error")
	}
	rr := httptest.NewRecorder()

	c, _ := gin.CreateTestContext(rr)
	c.Request = req
	testCtl.CreateOption(c)
	if rr.Code != http.StatusCreated {
		t.Fatalf("Status expected is 201 but %v", rr.Code)
	}
} // Done

func TestProductController_GetOptionById(t *testing.T) {
	ctr := gomock.NewController(t)
	defer ctr.Finish()
	mockSvc := services.NewMockIProductService(ctr)
	testCtl := NewProductController(mockSvc)

	mockSvc.EXPECT().GetOptionById(gomock.Any()).Return(&models.Option{
		Model:        gorm.Model{},
		Id:           0,
		ProductId:    "",
		Color:        "",
		Size:         "",
		ProductModel: "",
		Product:      models.Product{},
	}, nil).Times(1)

	req, err := http.NewRequest("POST", "api/products/:productId/options", nil)

	if err != nil {
		t.Fatal("Error")
	}
	rr := httptest.NewRecorder()

	c, _ := gin.CreateTestContext(rr)
	c.Params = []gin.Param{gin.Param{Key: "optionId", Value: "2"}}
	c.Request = req
	testCtl.GetOptionById(c)
	if rr.Code != http.StatusOK {
		t.Fatalf("Status expected is 200 but %v", rr.Code)
	}
} // Done

func TestProductController_GetOptions(t *testing.T) {
	ctr := gomock.NewController(t)
	defer ctr.Finish()
	mockSvc := services.NewMockIProductService(ctr)
	testCtl := NewProductController(mockSvc)

	mockSvc.EXPECT().GetOptions(gomock.Any()).Return(&[]models.Option{{Model: gorm.Model{},
		Id:           0,
		ProductId:    "",
		Color:        "",
		Size:         "",
		ProductModel: "",
		Product:      models.Product{}}}, nil).Times(1)

	req, err := http.NewRequest("Get", "api/products/:productId/options", nil)
	if err != nil {
		t.Fatal("Error")
	}
	rr := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(rr)

	c.Params = []gin.Param{gin.Param{Key: "optionId", Value: "2"}}
	c.Request = req
	testCtl.GetOptions(c)
	if rr.Code != http.StatusOK {
		t.Fatalf("Status expected is 200 but %v", rr.Code)
	}

} //Done

func TestProductController_UpdateOption(t *testing.T) {
	ctr := gomock.NewController(t)
	defer ctr.Finish()
	mockSvc := services.NewMockIProductService(ctr)
	testCtl := NewProductController(mockSvc)

	//Mock svc
	mockSvc.EXPECT().GetOptionById(gomock.Any()).Return(&models.Option{
		Model:        gorm.Model{},
		Id:           0,
		ProductId:    "",
		Color:        "",
		Size:         "",
		ProductModel: "",
		Product:      models.Product{},
	}, nil).Times(1)
	mockSvc.EXPECT().UpdateOption(gomock.Any()).Return(&models.Option{
		Model:        gorm.Model{},
		Id:           0,
		ProductId:    "",
		Color:        "",
		Size:         "",
		ProductModel: "",
		Product:      models.Product{},
	}, nil).Times(1)

	body := []byte("{}")

	req, err := http.NewRequest("PATCH", "api/options/:optionId", bytes.NewBuffer(body))

	if err != nil {
		t.Fatal("Error")
	}

	rr := httptest.NewRecorder()

	c, _ := gin.CreateTestContext(rr)
	c.Request = req

	c.Params = []gin.Param{gin.Param{Key: "optionId", Value: "2"}}
	c.Request = req
	testCtl.UpdateOption(c)
	if rr.Code != http.StatusOK {
		t.Fatalf("Status expected is 200 but %v", rr.Code)
	}
} //Done

func TestProductController_DeleteOption(t *testing.T) {
	ctr := gomock.NewController(t)
	defer ctr.Finish()
	mockSvc := services.NewMockIProductService(ctr)
	testCtl := NewProductController(mockSvc)

	mockSvc.EXPECT().DeleteOption(gomock.Any()).Return(&models.Option{
		Model:        gorm.Model{},
		Id:           0,
		ProductId:    "",
		Color:        "",
		Size:         "",
		ProductModel: "",
		Product:      models.Product{},
	}, nil).Times(1)

	req, err := http.NewRequest("DELETE", "api/options/:optionId", nil)
	if err != nil {
		t.Fatal("Error")
	}
	rr := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(rr)
	c.Request = req
	c.Params = []gin.Param{gin.Param{Key: "optionId", Value: "2"}}
	testCtl.DeleteOption(c)
	if rr.Code != http.StatusOK {
		t.Fatalf("Status is expected 200 but %v", rr.Code)
	}

} //Done

// Test Product Controller
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
	body := `{"id": "siu"}`
	//request
	req, err := http.NewRequest("POST", "api/products", strings.NewReader(body))
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
	svr.EXPECT().GetProductById(gomock.Any()).Return(&models.Product{
		Model:       gorm.Model{},
		Id:          "",
		Name:        "",
		Price:       "",
		Description: "",
		Quantity:    0,
	}, nil).Times(1)
	//Create mock body
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
