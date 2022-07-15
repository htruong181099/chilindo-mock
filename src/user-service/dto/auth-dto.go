package dto

import "chilindo/src/user-service/models"

//type SignUpDTO struct {
//
//}

type SignInDTO struct {
	Username string
	Password string
}

//type SignUpDTO *models.User

type SignUpDTO struct {
	User *models.User
}

type CheckAdminDTO struct {
	Token string
}

func NewSignUpDTO(user *models.User) *SignUpDTO {
	return &SignUpDTO{User: user}
}
