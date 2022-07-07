package dto

import "chilindo/models"

//type SignUpDTO struct {
//
//}

type SignInDTO struct {
	Username string
	Password string
}

type SignUpDTO struct {
	User *models.User
}
