package utils

import (
	"github.com/golang-jwt/jwt"
	"log"
	"time"
)

var jwtKey = []byte("supersecretkey")

type IJwtUtil interface {
	GenerateJWT(email string, username string, id int) (tokenString string, err error)
	ExtractToken(tokenString string) (*JWTClaim, error) // Lỗi vì không có giá trị bên trong struct
}
type JWTClaim struct {
	Username string
	Email    string
	Id       int
	Role     string
	jwt.StandardClaims
}

func (j *JWTClaim) GenerateJWT(email string, username string, id int, role string) (tokenString string, err error) {
	expirationTime := time.Now().Add(1 * time.Hour)
	claims := &JWTClaim{
		Email:    email,
		Username: username,
		Id:       id,
		Role:     role,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS384, claims)
	tokenString, err = token.SignedString(jwtKey)
	return
}

func ExtractToken(signedToken string) (*JWTClaim, error) {
	token, err := jwt.ParseWithClaims(
		signedToken,
		&JWTClaim{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(jwtKey), nil
		},
	)
	if err != nil {
		log.Println("ExtractToken : Error in jwt to parse")
		return nil, err
	}
	claims, ok := token.Claims.(*JWTClaim)
	if !ok {
		log.Println("ExtractToken : Error in jwt")
		return nil, err
	}
	return claims, nil
}
