package dto

import "chilindo/models"

type GetUserByIdDTO struct {
	//field
	UserId int
}

type CreateUserDTO struct {
	User *models.User
}
