package repository

import (
	"chilindo/dto"
	"chilindo/models"
	"gorm.io/gorm"
	"log"
)

type IAddressRepository interface {
	CreateAddress(address *models.Address) (*models.Address, error)
	GetAddress(dto dto.GetAddressDTO) (*models.Address, error)
	UpdateAddress(userid int, address *models.Address) (*models.Address, error)
	DeleteAddress(userid int) (*models.Address, error)
}

type AddressRepository struct {
	db *gorm.DB
}

func (a AddressRepository) CreateAddress(address *models.Address) (*models.Address, error) {
	err := a.db.Create(address)
	if err.Error != nil {
		log.Println("CreateAddress: Error Create in package repository", err)
		return nil, err.Error
	}
	return address, nil
}

func (a AddressRepository) GetAddress(dto *dto.GetAddressDTO) ([]models.Address, error) {
	var address []models.Address
	result := a.db.Where("user_id = ?", dto.UserId).Find(&address)
	if result.Error != nil {
		log.Println("GetAddress: Error Find in package repository", result.Error)
		return nil, result.Error
	}
	return address, nil
}

func (a AddressRepository) GetAddressById(dto dto.GetAddressByIdDTO) (*models.Address, error) {
	var address *models.Address
	result := a.db.Where("id = ?", dto.AddressId).Find(&address)
	if result.Error != nil {
		log.Println("GetAddress: Error Find in package repository", result.Error)
		return nil, result.Error
	}
	return address, nil
}

func (a AddressRepository) UpdateAddress(userid int, address *models.Address) (*models.Address, error) {
	var addressFind *models.Address
	err := a.db.Where("user_id = ?", userid).Find(&addressFind)
	if err.Error != nil {
		log.Println("UpdateAddress: Error to Find in package repository", err)
		return nil, err.Error
	}
	addressFind = address
	a.db.Save(&addressFind)
	return addressFind, nil
}

func (a AddressRepository) DeleteAddress(userid int) (*models.Address, error) {
	var deleteAddress *models.Address
	errFind := a.db.Where("user_id = ? AND id = ?", userid).Find(&deleteAddress)
	if errFind.Error != nil {
		log.Println("DeleteAddress: Error to find Address  in package repository", errFind)
		return nil, errFind.Error
	}
	errDelete := a.db.Delete(&deleteAddress)
	if errDelete.Error != nil {
		log.Println("DeleteAddress: Error to find Address  in package repository", errDelete)
		return nil, errDelete.Error
	}
	return deleteAddress, nil
}

func NewAddressRepository(db *gorm.DB) *AddressRepository {
	return &AddressRepository{db: db}
}
