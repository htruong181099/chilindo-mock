package controllers

import (
	"chilindo/src/auction-service/dtos"
	"chilindo/src/auction-service/models"
	"chilindo/src/auction-service/services"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
)

type IAuctionController interface {
	GetAuctions(c *gin.Context)
	GetAuctionById(c *gin.Context)
	CreateAuction(c *gin.Context)
}

type AuctionController struct {
	AuctionService services.IAuctionService
}

func (a AuctionController) GetAuctions(c *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (a AuctionController) GetAuctionById(c *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (a AuctionController) CreateAuction(c *gin.Context) {
	type auctionBody struct {
		Id           int     `json:"id" gorm:"primaryKey"`
		ProductId    string  `json:"productId"`
		StartingTime string  `json:"starting_time"`
		EndingTime   string  `json:"ending_time"`
		IsActive     bool    `json:"isActive" gorm:"default:false"`
		LowestBid    float32 `json:"lowestBid"`
	}
	var auctionBodyReq *auctionBody
	if err := c.ShouldBindJSON(&auctionBodyReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Message": "Fail to create Auction",
		})
		log.Println("CreateAuction: Error to ShouldBindJSON in package controller", err)
		c.Abort()
		return
	}
	fmt.Println("body", auctionBodyReq)
	startingTime, errStat := time.Parse(time.RFC3339, auctionBodyReq.StartingTime)
	if errStat != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Message": "Fail to create Auction",
		})
		log.Println("CreateAuction: Error to parse startingTime in package controller ", errStat)
		c.Abort()
		return
	}
	endingTime, errEnd := time.Parse(time.RFC3339, auctionBodyReq.EndingTime)
	if errEnd != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Message": "Fail to create Auction",
		})
		log.Println("CreateAuction: Error to parse endingTime in package controller", errEnd)
		c.Abort()
		return
	}
	var auction = &models.Auction{
		Id:           auctionBodyReq.Id,
		ProductId:    auctionBodyReq.ProductId,
		StartingTime: startingTime,
		EndingTime:   endingTime,
		IsActive:     auctionBodyReq.IsActive,
		LowestBid:    auctionBodyReq.LowestBid,
	}

	fmt.Println(auction)
	var dto dtos.CreateAuctionDTO
	dto.Auction = auction
	auc, err := a.AuctionService.CreateAuction(&dto)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Message": "Fail to create Auction",
		})
		log.Println("CreateAuction: Error to CreateAuction in package controller", err)
		c.Abort()
		return
	}
	c.JSON(http.StatusCreated, auc)

} //Done

func NewAuctionController(auctionService services.IAuctionService) *AuctionController {
	return &AuctionController{AuctionService: auctionService}
}
