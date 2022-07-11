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

type IProductController interface {
	CreateProduct(c *gin.Context)  //Done
	GetProducts(c *gin.Context)    //Done
	GetProductById(c *gin.Context) //Done
	UpdateProduct(c *gin.Context)  //Done
	DeleteProduct(c *gin.Context)
}

const idProduct = "id"

type ProductController struct {
	ProductService services.IProductService
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
	c.JSON(http.StatusCreated, product)
} //done

func (p ProductController) GetProducts(c *gin.Context) {
	products, err := p.ProductService.GetProducts()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Message": "Fail to get all products",
		})
		log.Println("GetProducts: Error get all product in package controller", err)
		c.Abort()
	}
	c.JSON(http.StatusOK, products)
} //Done

func (p ProductController) GetProductById(c *gin.Context) {
	var dto dtos.ProductDTO
	dto.ProductId = c.Param(idProduct)
	c.Set(config.ProductID, dto.ProductId)
	product, err := p.ProductService.GetProductById(&dto)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Message": "Error to get Product ",
		})
		log.Println("GetProductById: Error in package controller", err)
		c.Abort()
		return
	}
	c.JSON(http.StatusOK, product)
} //Done

func (p ProductController) UpdateProduct(c *gin.Context) {
	productId := c.Param(idProduct)
	var productUpdateBody *models.Product
	if err := c.ShouldBindJSON(&productUpdateBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Message": "Error to update product",
		})
		log.Println("UpdateProduct: Error ShouldBindJSON in package controller", err)
		c.Abort()
		return
	}
	dtoUpdate := dtos.NewUpdateProductDTO(productUpdateBody)
	dtoUpdate.ProductId = productId
	product, err := p.ProductService.UpdateProduct(dtoUpdate)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Message": "Error to update product",
		})
		log.Println("UpdateProduct: Error Update in package controller", err)
		c.Abort()
		return
	}
	c.JSON(http.StatusOK, product)

} //Done

func (p ProductController) DeleteProduct(c *gin.Context) {
	var dto dtos.ProductDTO
	fmt.Println(dto.ProductId)
	dto.ProductId = c.Param(idProduct)
	fmt.Println(dto.ProductId)
	product, err := p.ProductService.DeleteProduct(&dto)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Message": "Error to delete product",
		})
		log.Println("DeleteProduct: Error to get id product in package controller", err)
		c.Abort()
		return
	}
	c.JSON(http.StatusOK, product)
} //Done

func NewProductController(productService services.IProductService) *ProductController {
	return &ProductController{ProductService: productService}
}
