package controllers

import (
	"chilindo/pkg/pb/product"
	productMock "chilindo/src/auction-service/mocks"
	"chilindo/src/auction-service/models"
	"chilindo/src/auction-service/services/mocks"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"gorm.io/gorm"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"
)

func CreateTest(t *testing.T) (*services.MockIAuctionService, *productMock.MockProductServiceClient, *AuctionController) {
	ctr := gomock.NewController(t)
	defer ctr.Finish()
	auctionService := services.NewMockIAuctionService(ctr)
	productClient := productMock.NewMockProductServiceClient(ctr)
	auctionCtl := NewAuctionController(auctionService, productClient)

	return auctionService, productClient, auctionCtl
}

func TestAuctionController_CreateAuction(t *testing.T) {
	auctionService, productClient, auctionCtl := CreateTest(t)
	productClient.EXPECT().
		GetProduct(gomock.Any(), gomock.Any()).
		Return(&product.GetProductResponse{
			IsFound:     true,
			Id:          "1",
			Name:        "",
			Price:       "",
			Description: "",
			Quantity:    0,
		}, nil).
		Times(1)
	auctionService.EXPECT().
		CreateAuction(gomock.Any()).
		Return(&models.Auction{
			Model:        gorm.Model{},
			Id:           0,
			ProductId:    "",
			StartingTime: time.Time{},
			EndingTime:   time.Time{},
			IsActive:     false,
			LowestBid:    0,
		}, nil).
		Times(1)

	w := httptest.NewRecorder()
	body := `{
		"starting_time"	:	"2022-07-15T12:38:40+00:00",
		"ending_time"	:	"2022-07-15T12:38:40+00:00"
	}`
	req, reqErr := http.NewRequest("POST", "api/auction", strings.NewReader(body))
	if reqErr != nil {
		t.Fatalf("Error %v", reqErr)
	}
	c, _ := gin.CreateTestContext(w)
	c.Request = req

	auctionCtl.CreateAuction(c)

	if w.Code != http.StatusCreated {
		t.Fatalf("Expect %v but got %v", http.StatusCreated, w.Code)
	}

}

func TestAuctionController_CreateAuction_ErrorParseTime(t *testing.T) {
	auctionService, productClient, auctionCtl := CreateTest(t)
	productClient.EXPECT().
		GetProduct(gomock.Any(), gomock.Any()).
		Return(&product.GetProductResponse{
			IsFound:     true,
			Id:          "1",
			Name:        "",
			Price:       "",
			Description: "",
			Quantity:    0,
		}, nil).
		Times(1)
	auctionService.EXPECT().
		CreateAuction(gomock.Any()).
		Return(&models.Auction{
			Model:        gorm.Model{},
			Id:           0,
			ProductId:    "",
			StartingTime: time.Time{},
			EndingTime:   time.Time{},
			IsActive:     false,
			LowestBid:    0,
		}, nil).
		Times(1)

	w := httptest.NewRecorder()
	body := `{}`
	req, reqErr := http.NewRequest("POST", "api/auction", strings.NewReader(body))
	if reqErr != nil {
		t.Fatalf("Error %v", reqErr)
	}
	c, _ := gin.CreateTestContext(w)
	c.Request = req

	auctionCtl.CreateAuction(c)

	if w.Code != http.StatusBadRequest {
		t.Fatalf("Expect %v but got %v", http.StatusCreated, w.Code)
	}

}
