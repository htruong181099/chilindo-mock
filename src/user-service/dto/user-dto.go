package dto

import "chilindo/src/user-service/models"

type GetUserByIdDTO struct {
	//field
	UserId int
}

type CreateUserDTO struct {
	User *models.User
}
