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
	"strconv"
)

type IProductController interface {
	CreateProduct(c *gin.Context)  //Done
	GetProducts(c *gin.Context)    //Done
	GetProductById(c *gin.Context) //Done
	UpdateProduct(c *gin.Context)  //Done
	DeleteProduct(c *gin.Context)  //Done
	CreateOption(c *gin.Context)   //Done
	GetOptions(c *gin.Context)
	GetOptionById(c *gin.Context)
	UpdateOption(c *gin.Context)
	DeleteOption(c *gin.Context)
}

const (
	productId = "productId"
	optionId  = "optionId"
)

type ProductController struct {
	ProductService services.IProductService
}

func (p ProductController) CreateOption(c *gin.Context) {
	var optionBody *models.Option
	if err := c.ShouldBindJSON(&optionBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Message": "Fail to create option",
		})
		log.Println("CreateOption: Error to ShouldBindJSON in package controller", err)
		c.Abort()
		return
	}
	dto := dtos.NewCreateOption(optionBody)
	dto.Option.ProductId = c.Param(productId)
	option, err := p.ProductService.CreateOption(dto)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Message": "Fail to create option",
		})
		log.Println("CreateOption: Error to CreateOption in package controller", err)
		c.Abort()
		return
	}
	c.JSON(http.StatusCreated, option)
} //Done

func (p ProductController) GetOptions(c *gin.Context) {
	id := c.Param(productId)
	var dto dtos.ProductIdDTO
	dto.ProductId = id
	options, err := p.ProductService.GetOptions(&dto)
	if err != nil {
		log.Println("GetOptions: error in controller package", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"Message": "Fail to get options",
		})
		c.Abort()
		return
	}
	if options == nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Message": "Not found options",
		})
		c.Abort()
		return
	}
	c.JSON(http.StatusOK, options)
} //Done

func (p ProductController) GetOptionById(c *gin.Context) {
	oid := c.Param(optionId)
	var dto dtos.OptionIdDTO
	oidInt, conErr := strconv.Atoi(oid)
	if conErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Message": "Error get option",
		})
		log.Println("GetOptionById: Error parse option param to id", conErr)
		c.Abort()
		return
	}

	dto.OptionId = oidInt

	option, err := p.ProductService.GetOptionById(&dto)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Message": "Error get option",
		})
		log.Println("GetOptionById: Error call service in pkg controller", conErr)
		c.Abort()
		return
	}
	if option == nil {
		c.JSON(http.StatusNotFound, gin.H{
			"Message": "Option not found",
		})
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, option)

} //done

func (p ProductController) UpdateOption(c *gin.Context) {
	var optionBody *models.Option
	if err := c.ShouldBindJSON(&optionBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Message": "Error update option",
		})
		log.Println("UpdateOption: Error ShouldBindJSON ", err)
		c.Abort()
		return
	}
	oid, errCv := strconv.Atoi(c.Param(optionId))
	if errCv != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Message": "Error update option",
		})
		log.Println("UpdateOption: Error parse param", errCv)
		c.Abort()
		return
	}
	optionBody.Id = oid
	dto := dtos.NewUpdateOptionDTO(optionBody)

	option, err := p.ProductService.UpdateOption(dto)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Message": "Error update option",
		})
		log.Println("UpdateOption: Error call service", errCv)
		c.Abort()
		return
	}
	if option == nil {
		c.JSON(http.StatusNotFound, gin.H{
			"Message": "Not found",
		})
		c.Abort()
		return
	}
	c.JSON(http.StatusOK, option)
}

func (p ProductController) DeleteOption(c *gin.Context) {
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
	dto.ProductId = c.Param(productId)
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
	if product == nil {
		c.JSON(http.StatusNotFound, gin.H{
			"Message": "Not found product",
		})
		c.Abort()
		return
	}
	c.JSON(http.StatusOK, product)
} //Done

func (p ProductController) UpdateProduct(c *gin.Context) {
	productId := c.Param(productId)
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
	dto.ProductId = c.Param(productId)
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
