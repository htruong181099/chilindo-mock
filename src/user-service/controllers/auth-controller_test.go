package controllers

import (
	"bytes"
	"chilindo/src/user-service/models"
	services "chilindo/src/user-service/services/mocks"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"gorm.io/gorm"
	"net/http"
	"net/http/httptest"
	"testing"
)

func CreateTestAuth(t *testing.T) (*services.MockIAuthService, *AuthController) {
	ctr := gomock.NewController(t)
	defer ctr.Finish()
	mockSvr := services.NewMockIAuthService(ctr)
	//adminClient := adminMock.NewMockAdminServiceClient(ctr)
	authCtr := NewAuthController(mockSvr)
	return mockSvr, authCtr
}

func TestAuthController_SignUp(t *testing.T) {
	mockSvr, authCtr := CreateTestAuth(t)

	mockSvr.EXPECT().SignUp(gomock.Any()).Return(&models.User{
		Model:       gorm.Model{},
		Id:          0,
		FirstName:   "",
		LastName:    "",
		Username:    "",
		Password:    "",
		Email:       "",
		PhoneNumber: "",
		Gender:      "",
		Language:    "",
		Role:        "",
	}, nil).Times(1)

	body := []byte("{}")
	req, err := http.NewRequest("POST", "api/auth/signup", bytes.NewBuffer(body))

	if err != nil {
		t.Fatalf("Error")
	}
	w := httptest.NewRecorder()

	c, _ := gin.CreateTestContext(w)

	c.Request = req
	authCtr.SignUp(c)

	if w.Code != http.StatusCreated {
		t.Fatalf("Expected Status is 201 but got %v", w.Code)
	}
} //Done

func TestAuthController_SignIn(t *testing.T) {
	mockSvr, authCtr := CreateTestAuth(t)

	mockSvr.EXPECT().SignIn(gomock.Any()).Return(&models.User{
		Model:       gorm.Model{},
		Id:          0,
		FirstName:   "",
		LastName:    "",
		Username:    "",
		Password:    "",
		Email:       "",
		PhoneNumber: "",
		Gender:      "",
		Language:    "",
		Role:        "",
	}, nil).Times(1)

	body := []byte("{}")

	req, err := http.NewRequest("POST", "api/auth/signup", bytes.NewBuffer(body))

	if err != nil {
		t.Fatalf("Error")
	}
	w := httptest.NewRecorder()

	c, _ := gin.CreateTestContext(w)

	c.Request = req

	authCtr.SignIn(c)
	if w.Code != http.StatusOK {
		t.Fatalf("200 but got %v", w.Code)
	}
} //Done
