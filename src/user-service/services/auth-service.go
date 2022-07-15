package services

import (
	"chilindo/pkg/pb/admin"
	jwtUtil "chilindo/pkg/utils"
	"chilindo/src/user-service/dto"
	"chilindo/src/user-service/models"
	"chilindo/src/user-service/repository"
	"log"
	"strings"
)

type IAuthService interface {
	SignUp(dto *dto.SignUpDTO) (*models.User, error)
	SignIn(dto *dto.SignInDTO) (*models.User, error)
	CheckIsAdmin(req *admin.CheckIsAdminRequest) (*admin.CheckIsAdminResponse, error)
}

type AuthService struct {
	UserRepository repository.IUserRepository
}

func NewAuthService(userRepository repository.IUserRepository) *AuthService {
	return &AuthService{UserRepository: userRepository}
}

func (u *AuthService) SignUp(dto *dto.SignUpDTO) (*models.User, error) {
	userRegis, err := u.UserRepository.CreateUser(dto)
	if err != nil {
		log.Println("SignUp: Error CreateUser in package service")
		return nil, err
	}
	return userRegis, nil
}

func (u AuthService) SignIn(dto *dto.SignInDTO) (*models.User, error) {
	user, err := u.UserRepository.GetUserByUsername(dto)
	if err != nil {
		log.Println("SignIn: Error GetUserByUsername in package service")
		return nil, err
	}
	return user, nil
}

func (u AuthService) CheckIsAdmin(req *admin.CheckIsAdminRequest) (*admin.CheckIsAdminResponse, error) {
	isAuth := false
	isAdmin := false
	tokenString := req.Token

	tokenResult := strings.TrimPrefix(tokenString, "Bearer ")
	//fmt.Println("Check Token", tokenResult)

	claims, err := jwtUtil.ExtractToken(tokenResult)
	if err != nil {
		log.Println("CheckIsAdmin: ", err)
		return nil, err
	}

	isAuth = true
	log.Println(claims)
	log.Println(claims.Role)
	if claims.Role == "admin" {
		isAdmin = true
	}

	return &admin.CheckIsAdminResponse{
		IsAuth:  isAuth,
		IsAdmin: isAdmin,
	}, nil
}
