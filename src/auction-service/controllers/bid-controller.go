package controllers

import (
	"chilindo/pkg/configs"
	"chilindo/src/auction-service/dtos"
	"chilindo/src/auction-service/models"
	"chilindo/src/auction-service/services"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
	"time"
)

const (
	auctionId = "auctionId"
	bidId     = "bidId"
)

type IBidController interface {
	GetBidsOfAuction(c *gin.Context)
	GetBidById(c *gin.Context)
	CreateBid(c *gin.Context)
}

type BidController struct {
	BidService services.IBidService
}

func (b BidController) GetBidsOfAuction(c *gin.Context) {
	aid, err := strconv.Atoi(c.Param(auctionId))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Message": "Error to get Bid by auction Id",
		})
		log.Println("GetBidsOfAuction: Error get actionID in pkg controller", err)
		c.Abort()
		return
	}
	var dto dtos.AuctionIdDTO
	dto.AuctionId = aid
	bid, errGetBid := b.BidService.GetBidsOfAuction(&dto)
	if errGetBid != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Message": "Not found bid of auction",
		})
		log.Println("GetBidsOfAuction: Error call to BidService in pkg controller", errGetBid)
		c.Abort()
		return
	}
	c.JSON(http.StatusOK, bid)
} //Done

func (b BidController) GetBidById(c *gin.Context) {
	bidId, err := strconv.Atoi(c.Param(bidId))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"Message": "Not found bid of bid",
		})
		log.Println("GetBidById: Error call to GetBidById in pkg controller", err)
		c.Abort()
		return
	}
	var dto dtos.BidIdDTO
	dto.BidId = bidId
	bid, errGetListBid := b.BidService.GetBidById(&dto)
	if errGetListBid != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Message": "Error get bid",
		})
		log.Println("GetBidById: Error call to GetBidById in pkg controller", errGetListBid)
		c.Abort()
		return
	}
	if bid == nil {
		c.JSON(http.StatusNotFound, gin.H{
			"Message": "Not found bid of auction",
		})
		c.Abort()
		return
	}
	c.JSON(http.StatusOK, bid)
}

func (b BidController) CreateBid(c *gin.Context) {
	var bidBody *models.Bid
	if err := c.ShouldBindJSON(&bidBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Message": "Fail to create bid ",
		})
		log.Println("CreateBid: Error ShouldBindJSON in pkg controller", err)
		c.Abort()
		return
	}
	timeBid := time.Now()
	userId, ok := c.Get(configs.UserID)
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{
			"Message": "Fail to create bid ",
		})
		log.Println("CreateBid: Error get bid in pkg controller")
		c.Abort()
		return
	}
	bidderId := userId.(int)
	auctionId, errCv := strconv.Atoi(c.Param(auctionId))
	if errCv != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Message": "Fail to create bid ",
		})
		log.Println("CreateBid: Error to convert AuctionId in pkg controller", errCv)
		c.Abort()
		return
	}
	var dto dtos.CreateBidDTO
	dto.Bid = bidBody
	dto.Bid.AuctionId = auctionId
	dto.Bid.BidderId = bidderId
	dto.Bid.BidTime = timeBid
	bid, errCreateBid := b.BidService.CreateBid(&dto)
	if errCreateBid != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Message": "Fail to create bid ",
		})
		log.Println("CreateBid: Error to call BidService in pkg controller", errCreateBid)
		c.Abort()
		return
	}
	if bid == nil {
		c.JSON(http.StatusNotFound, gin.H{
			"Message": "Not found auction ",
		})
		c.Abort()
		return
	}
	c.JSON(http.StatusCreated, bid)
}

func NewBidController(bidService services.IBidService) *BidController {
	return &BidController{BidService: bidService}
}
