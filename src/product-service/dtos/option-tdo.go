package dtos

import "chilindo/src/product-service/models"

type CreateOptionDTO struct {
	Option *models.Option
}

func NewCreateOption(option *models.Option) *CreateOptionDTO {
	return &CreateOptionDTO{Option: option}
}

type OptionIdDTO struct {
	OptionId int
}

type OptionByIdDTO struct {
	OptionId  int
	ProductId string
}

type UpdateOptionDTO struct {
	Option *models.Option
}

func NewUpdateOptionDTO(option *models.Option) *UpdateOptionDTO {
	return &UpdateOptionDTO{Option: option}
}

type ProductIdDTO struct {
	ProductId string
}
