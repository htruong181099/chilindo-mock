package services

import (
	"chilindo/src/product-service/dtos"
	"chilindo/src/product-service/models"
	"chilindo/src/product-service/repository"
	"log"
)

type IOptionService interface {
	CreateOption(dto *dtos.CreateOption) (*models.Option, error) //Done
	GetOption(dto *dtos.OptionById) (*[]models.Option, error)
	GetOptionById(dto *dtos.OptionById) (*models.Option, error)
	UpdateOption(dto *dtos.UpdateOptionBy) (*models.Option, error)
	DeleteOption(dto *dtos.OptionById) (*models.Option, error)
}

type OptionService struct {
	OptionRepository repository.IOptionRepository
}

func (o OptionService) CreateOption(dto *dtos.CreateOption) (*models.Option, error) {
	option, err := o.OptionRepository.CreateOption(dto)
	if err != nil {
		log.Println("CreateOption: Error in package service")
		return nil, err
	}
	return option, nil
} //Done

func (o OptionService) GetOption(dto *dtos.OptionById) (*[]models.Option, error) {
	//TODO implement me
	panic("implement me")
}

func (o OptionService) GetOptionById(dto *dtos.OptionById) (*models.Option, error) {
	//TODO implement me
	panic("implement me")
}

func (o OptionService) UpdateOption(dto *dtos.UpdateOptionBy) (*models.Option, error) {
	//TODO implement me
	panic("implement me")
}

func (o OptionService) DeleteOption(dto *dtos.OptionById) (*models.Option, error) {
	//TODO implement me
	panic("implement me")
}

func NewOptionService(optionRepository repository.IOptionRepository) *OptionService {
	return &OptionService{OptionRepository: optionRepository}
}
