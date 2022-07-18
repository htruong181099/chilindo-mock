package dto

import "chilindo/src/user-service/models"

type GetByUserIdDTO struct {
	UserId int //`json:"userId"`
}

type UpdatePasswordDTO struct {
	UserId          int    `json:"userId"`
	CurrentPassword string `json:"currentPassword"`
	NewPassword     string `json:"newPassword"`
}

type CreateUserDTO struct {
	User *models.User
}
