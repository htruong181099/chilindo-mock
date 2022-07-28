package controllers

import (
	"chilindo/src/product-service/config"
	"chilindo/src/product-service/dtos"
	"chilindo/src/product-service/models"
	"chilindo/src/product-service/services"
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
	GetOptions(c *gin.Context)     //Done
	GetOptionById(c *gin.Context)  //Done
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
		c.JSON(http.StatusNotFound, gin.H{
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
	oid, errCv := strconv.Atoi(c.Param(optionId))
	if errCv != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Message": "Error update option",
		})
		log.Println("UpdateOption: Error parse param", errCv)
		c.Abort()
		return
	}
	optionBody, obErr := p.ProductService.GetOptionById(&dtos.OptionIdDTO{
		OptionId: oid,
	})
	if obErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Message": "Error update option",
		})
		log.Println("UpdateOption: Error get option", obErr)
		c.Abort()
		return
	}
	if optionBody == nil {
		c.JSON(http.StatusNotFound, gin.H{
			"Message": "Not found",
		})
		c.Abort()
		return
	}

	if err := c.ShouldBindJSON(&optionBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Message": "Error update option",
		})
		log.Println("UpdateOption: Error ShouldBindJSON ", err)
		c.Abort()
		return
	}

	dto := dtos.NewUpdateOptionDTO(optionBody)

	option, oErr := p.ProductService.UpdateOption(dto)
	if oErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Message": "Error update option",
		})
		log.Println("UpdateOption: Error call service", oErr)
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
} //Done

func (p ProductController) DeleteOption(c *gin.Context) {
	oId, errGetOId := strconv.Atoi(c.Param(optionId))
	if errGetOId != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Message": "Error to delete option",
		})
		log.Println("DeleteOption: Error to parse oId", errGetOId)
		c.Abort()
		return
	}
	var dto dtos.OptionIdDTO
	dto.OptionId = oId
	_, err := p.ProductService.DeleteOption(&dto)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Message": "Error to delete option",
		})
		log.Println("DeleteOption: Error to parse oId", err)
		c.Abort()
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"Message": "Delete success",
	})
} //Done

//Product Controller

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
	if productBody.Id == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"Message": "Fail to create product",
		})
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
	if c.Request.ContentLength == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"Message": "Empty body",
		})
		c.Abort()
		return
	}
	productId := c.Param(productId)
	productBody, err := p.ProductService.GetProductById(&dtos.ProductDTO{ProductId: productId})
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Message": "Error to update product",
		})
		log.Println("UpdateProduct: Error Update in package controller", err)
		c.Abort()
		return
	}
	if productBody == nil {
		c.JSON(http.StatusNotFound, gin.H{
			"Message": "Not Found",
		})
		c.Abort()
		return
	}
	if err := c.ShouldBindJSON(&productBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Message": "Error to update product",
		})
		log.Println("UpdateProduct: Error ShouldBindJSON in package controller", err)
		c.Abort()
		return
	}
	dtoUpdate := dtos.NewUpdateProductDTO(productBody)
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
	if product == nil {
		c.JSON(http.StatusNotFound, gin.H{
			"Message": "Not Found",
		})
		c.Abort()
		return
	}
	c.JSON(http.StatusOK, product)

} //Done

func (p ProductController) DeleteProduct(c *gin.Context) {
	var dto dtos.ProductDTO
	dto.ProductId = c.Param(productId)
	_, err := p.ProductService.DeleteProduct(&dto)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Message": "Error to delete product",
		})
		log.Println("DeleteProduct: Error to get id product in package controller", err)
		c.Abort()
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"Message": "Delete Success",
	})
} //Done

func NewProductController(productService services.IProductService) *ProductController {
	return &ProductController{ProductService: productService}
}
