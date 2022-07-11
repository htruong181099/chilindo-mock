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

func TestOptionController_CreateOption(t *testing.T) {
	ctr := gomock.NewController(t)
	defer ctr.Finish()
	srv := services.NewMockIOptionService(ctr)
	optionCtr := NewOptionController(srv)

	//Mock Service
	srv.EXPECT().CreateOption(gomock.Any()).Return(&models.Option{
		Model:     gorm.Model{},
		Id:        0,
		ProductId: "",
		Link:      "",
		Product:   models.Product{},
	}, nil).Times(1)

	body := []byte("{}")

	//create request
	req, err := http.NewRequest("POST", "api/options", bytes.NewBuffer(body))

	if err != nil {
		t.Fatalf("Error")
	}

	//create respond
	rr := httptest.NewRecorder()

	c, _ := gin.CreateTestContext(rr)
	c.Request = req
	//testing
	optionCtr.CreateOption(c)

	if rr.Code != http.StatusCreated {
		t.Fatalf("Status expected 201 but %v", rr.Code)
	}
} //Done
