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
	CheckUserAuth(req *admin.CheckUserAuthRequest) (*admin.CheckUserAuthResponse, error)
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

	claims, err := jwtUtil.ExtractToken(tokenResult)
	if err != nil {
		log.Println("CheckIsAdmin: ", err)
		return nil, err
	}

	isAuth = true
	if claims.Role == "admin" {
		isAdmin = true
	}

	return &admin.CheckIsAdminResponse{
		IsAuth:  isAuth,
		IsAdmin: isAdmin,
	}, nil
}

func (u AuthService) CheckUserAuth(req *admin.CheckUserAuthRequest) (*admin.CheckUserAuthResponse, error) {
	isAuth := false
	tokenString := req.Token
	log.Println("Check here")
	tokenResult := strings.TrimPrefix(tokenString, "Bearer ")

	claims, err := jwtUtil.ExtractToken(tokenResult)
	if err != nil {
		log.Println("CheckUserAuth: ", err)
		//return nil, err
		return &admin.CheckUserAuthResponse{
			IsAuth: false,
		}, nil
	}
	isAuth = true
	userId := int32(claims.Id)
	log.Println(claims)
	log.Println(claims.Role)

	return &admin.CheckUserAuthResponse{
		IsAuth: isAuth,
		UserId: userId,
	}, nil
}
