package repository

import (
	"chilindo/dto"
	"chilindo/models"
	"gorm.io/gorm"
	"log"
)

type IAddressRepository interface {
	CreateAddress(dto *dto.CreateAddressDTO) (*models.Address, error)
	GetAddress(dto *dto.GetAddressDTO) (*[]models.Address, error)
	GetAddressById(dto *dto.GetAddressByIdDTO) (*models.Address, error)
	UpdateAddress(dto *dto.UpdateAddressDTO) (*models.Address, error)
	DeleteAddressById(dto *dto.DeleteAddressByIdDTO) (*models.Address, error)
}

type AddressRepository struct {
	db *gorm.DB
}

func (a *AddressRepository) CreateAddress(dto *dto.CreateAddressDTO) (*models.Address, error) {
	var address *models.Address
	address.UserId = dto.UserId
	result := a.db.Create(&address)
	if result.Error != nil {
		log.Println("CreateAddress: Error Create in package repository", result)
		return nil, result.Error
	}
	return address, nil
}

func (a *AddressRepository) GetAddress(dto *dto.GetAddressDTO) (*[]models.Address, error) {
	var address *[]models.Address
	result := a.db.Where("user_id = ?", dto.UserId).Find(&address)
	if result.Error != nil {
		log.Println("GetAddress: Error Find in package repository", result.Error)
		return nil, result.Error
	}
	return address, nil
}

func (a *AddressRepository) GetAddressById(dto *dto.GetAddressByIdDTO) (*models.Address, error) {
	var address *models.Address
	result := a.db.Where("id = ? And user_id =?", dto.AddressId, dto.UserId).Find(&address)
	if result.Error != nil {
		log.Println("GetAddress: Error Find in package repository", result.Error)
		return nil, result.Error
	}
	return address, nil
}

func (a *AddressRepository) UpdateAddress(dto *dto.UpdateAddressDTO) (*models.Address, error) {
	var addressFind *models.Address
	err := a.db.Where("user_id = ?", dto.UserId).Find(&addressFind)
	if err.Error != nil {
		log.Println("UpdateAddress: Error to Find in package repository", err)
		return nil, err.Error
	}
	addressFind = dto.Address
	a.db.Save(&addressFind)
	return addressFind, nil
}

func (a *AddressRepository) DeleteAddressById(dto *dto.DeleteAddressByIdDTO) (*models.Address, error) {
	var deleteAddress *models.Address
	resultFind := a.db.Where("user_id = ? AND id = ?", dto.UserId, dto.AddressId).Find(&deleteAddress)
	if resultFind.Error != nil {
		log.Println("DeleteAddress: Error to find Address  in package repository", resultFind)
		return nil, resultFind.Error
	}
	resultDelete := a.db.Delete(&deleteAddress)
	if resultDelete.Error != nil {
		log.Println("DeleteAddress: Error to Deleted Address  in package repository", resultDelete)
		return nil, resultDelete.Error
	}
	return deleteAddress, nil
}

func NewAddressRepository(db *gorm.DB) *AddressRepository {
	return &AddressRepository{db: db}
}
