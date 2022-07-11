package dtos

import "chilindo/src/product-service/models"

type CreateOption struct {
	Option *models.Option
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
