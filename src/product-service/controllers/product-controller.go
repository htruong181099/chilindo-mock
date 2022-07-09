package controllers

import (
	"chilindo/src/product-service/dtos"
	"chilindo/src/product-service/models"
	"chilindo/src/product-service/services"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type IProductController interface {
	CreateProduct(c *gin.Context)
	GetProducts(c *gin.Context)
	GetProductById(c *gin.Context)
	UpdateProduct(c *gin.Context)
	DeleteProduct(c *gin.Context)
}

type ProductController struct {
	ProductService services.IProductService
}

func (p ProductController) GetProducts(c *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (p ProductController) GetProductById(c *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (p ProductController) UpdateProduct(c *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (p ProductController) DeleteProduct(c *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (p ProductController) CreateProduct(c *gin.Context) {
	var productBody *models.Product
	if err := c.ShouldBindJSON(&productBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Message": "Fail to create product",
		})
		log.Println("CreateProduct: Error to ShouldBindJSON in package controller", err)
		c.Abort()
		return
	}
	dto := dtos.NewCreateProductDTO(productBody)
	product, err := p.ProductService.CreateProduct(dto)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Message": "Fail to create product",
		})
		log.Println("CreateProduct: Error to create product in package controller", err)
		c.Abort()
		return
	}
	c.JSON(http.StatusOK, product)
}

func NewProductController(productService services.IProductService) *ProductController {
	return &ProductController{ProductService: productService}
}
