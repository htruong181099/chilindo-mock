package dtos

import "chilindo/src/product-service/models"

type CreateOption struct {
	Option *models.Option
}

func NewCreateOption(option *models.Option) *CreateOption {
	return &CreateOption{Option: option}
}

type OptionById struct {
	OptionId  int
	ProductId string
}

type UpdateOptionBy struct {
	Option    *models.Option
	OptionId  int
	ProductId string
}
