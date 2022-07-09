package controllers

import (
	"chilindo/src/product-service/services"
	"github.com/gin-gonic/gin"
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
