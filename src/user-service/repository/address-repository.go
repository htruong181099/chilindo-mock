package repository

import (
	"chilindo/models"
	"gorm.io/gorm"
	"log"
)

type IAddressRepository interface {
	CreateAddress(address *models.Address) (*models.Address, error)
	GetAddress(userid int) (*models.Address, error)
	UpdateAddress(address *models.Address) (*models.Address, error)
	DeleteAddress(userid int) (*models.Address, error)
}

type AddressRepository struct {
	db *gorm.DB
}

func (a AddressRepository) CreateAddress(address *models.Address) (*models.Address, error) {
	err := a.db.Create(address)
	if err.Error != nil {
		log.Println("CreateAddress: Error Create in package repository")
		return nil, err.Error
	}
	return address, nil
}

func (a AddressRepository) GetAddress(userid int) (*models.Address, error) {
	//TODO implement me
	panic("implement me")
}

func (a AddressRepository) UpdateAddress(address *models.Address) (*models.Address, error) {
	//TODO implement me
	panic("implement me")
}

func (a AddressRepository) DeleteAddress(userid int) (*models.Address, error) {
	//TODO implement me
	panic("implement me")
}

func NewAddressRepository(db *gorm.DB) *AddressRepository {
	return &AddressRepository{db: db}
}
