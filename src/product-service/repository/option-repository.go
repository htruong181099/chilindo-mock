package repository

import (
	"chilindo/src/product-service/dtos"
	"chilindo/src/product-service/models"
	"gorm.io/gorm"
	"log"
)

type IOptionRepository interface {
	//create option
	CreateOption(dto *dtos.CreateOption) (*models.Option, error) //Done
	GetOption(dto *dtos.OptionById) (*[]models.Option, error)
	GetOptionById(dto *dtos.OptionById) (*models.Option, error)
	UpdateOption(dto *dtos.UpdateOptionBy) (*models.Option, error)
	DeleteOption(dto *dtos.OptionById) (*models.Option, error)
}

type OptionRepository struct {
	db *gorm.DB
}

func (o OptionRepository) CreateOption(dto *dtos.CreateOption) (*models.Option, error) {
	record := o.db.Create(&dto.Option)
	if record.Error != nil {
		log.Println("CreateOption: Error to create repository")
		return nil, record.Error
	}
	return dto.Option, nil
}

func (o OptionRepository) GetOption(dto *dtos.OptionById) (*[]models.Option, error) {
	var option *[]models.Option
	record := o.db.Where("ProductId = ?", dto.ProductId).Find(&option)
	if record.Error != nil {
		log.Println("GetOption : Error to get all option", record.Error)
		return nil, record.Error
	}
	return option, nil

}

func (o OptionRepository) GetOptionById(dto *dtos.OptionById) (*models.Option, error) {
	//TODO implement me
	panic("implement me")
}

func (o OptionRepository) UpdateOption(dto *dtos.UpdateOptionBy) (*models.Option, error) {
	//TODO implement me
	panic("implement me")
}

func (o OptionRepository) DeleteOption(dto *dtos.OptionById) (*models.Option, error) {
	//TODO implement me
	panic("implement me")
}

func NewOptionRepository(db *gorm.DB) *OptionRepository {
	return &OptionRepository{db: db}
}
