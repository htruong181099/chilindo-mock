package controllers

import (
	"chilindo/src/product-service/config"
	"chilindo/src/product-service/dtos"
	"chilindo/src/product-service/models"
	"chilindo/src/product-service/services"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type IImageController interface {
	CreateImage(c *gin.Context)
	GetImage(c *gin.Context)
	GetImageById(c *gin.Context)
	UpdateImage(c *gin.Context)
	DeleteImage(c *gin.Context)
}

type ImageController struct {
	ImageService services.IImageService
}

func (i ImageController) CreateImage(c *gin.Context) {
	var imageBody *models.Image
	if err := c.ShouldBindJSON(&imageBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Message": "Error to create Image for product",
		})
		log.Println("CreateImage: Error ShouldBindJSON in package controller", err)
		c.Abort()
		return
	}
	productId, ok := c.Get(config.ProductID)
	fmt.Println("Check here", productId)
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{
			"Message": "Error to create Image for product",
		})
		log.Println("CreateImage: Error get id Product in package controller")
		c.Abort()
		return
	}
	dto := dtos.NewImageDTO(imageBody, productId.(string))
	image, err := i.ImageService.CreateImage(dto)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Message": "Error to create Image for product",
		})
		log.Println("CreateImage: Error call service CreateImage in package controller", err)
		c.Abort()
		return
	}
	c.JSON(http.StatusCreated, image)

}

func (i ImageController) GetImage(c *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (i ImageController) GetImageById(c *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (i ImageController) UpdateImage(c *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (i ImageController) DeleteImage(c *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func NewImageController(imageService services.IImageService) *ImageController {
	return &ImageController{ImageService: imageService}
}
