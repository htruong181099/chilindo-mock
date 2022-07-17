package controllers

import (
	productMock "chilindo/src/auction-service/mocks"
	services "chilindo/src/auction-service/services/mocks"
	"fmt"
	"github.com/golang/mock/gomock"
	"testing"
)

func TestAuctionController_CreateAuction(t *testing.T) {
	ctr := gomock.NewController(t)
	defer ctr.Finish()
	auctionService := services.NewMockIAuctionService(ctr)
	productClient := productMock.NewMockProductServiceClient(ctr)
	auctionCtl := NewAuctionController(auctionService, productClient)
	fmt.Println(auctionCtl)

	//productClient.EXPECT().
	//	GetProduct(gomock.Any(), gomock.Any()).
	//	Return(&product
	//	.GetProductResponse{
	//		IsFound:     true,
	//		Id:          "1",
	//		Name:        "",
	//		Price:       "",
	//		Description: "",
	//		Quantity:    0,
	//	}, nil).
	//	Times(1)
	//auctionService.EXPECT().
	//	CreateAuction(gomock.Any()).
	//	Return(&models.Auction{
	//		Model:        gorm.Model{},
	//		Id:           0,
	//		ProductId:    "",
	//		StartingTime: time.Time{},
	//		EndingTime:   time.Time{},
	//		IsActive:     false,
	//		LowestBid:    0,
	//	}, nil).
	//	Times(1)

	//ez
	//ez

	//auctionService.EXPECT().
}
