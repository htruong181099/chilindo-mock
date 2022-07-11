package controllers

import (
	"chilindo/src/product-service/dtos"
	"chilindo/src/product-service/models"
	"chilindo/src/product-service/services"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type IOptionController interface {
	CreateOption(c *gin.Context) //Done
	GetOption(c *gin.Context)
	GetOptionById(c *gin.Context)
	UpdateOption(c *gin.Context)
	DeleteOption(c *gin.Context)
}

type OptionController struct {
	OptionService services.IOptionService
}

func (o OptionController) CreateOption(c *gin.Context) {
	var optionBody *models.Option
	if err := c.ShouldBindJSON(&optionBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Message": "Create option fail",
		})
		log.Println("CreateOption: Error in ShouldBindJSON", err)
		c.Abort()
		return
	}
	dto := dtos.NewCreateOption(optionBody)
	option, errCreateOption := o.OptionService.CreateOption(dto)
	if errCreateOption != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Message": "Create option fail",
		})
		log.Println("CreateOption: Error in CreateOption controller to service in package controller", errCreateOption)
		c.Abort()
		return
	}
	c.JSON(http.StatusCreated, option)
} //Done

func (o OptionController) GetOption(c *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (o OptionController) GetOptionById(c *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (o OptionController) UpdateOption(c *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (o OptionController) DeleteOption(c *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func NewOptionController(optionService services.IOptionService) *OptionController {
	return &OptionController{OptionService: optionService}
}
