package controllers

import (
	"bytes"
	"chilindo/src/user-service/config"
	"chilindo/src/user-service/models"
	services "chilindo/src/user-service/services/mocks"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"gorm.io/gorm"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func CreateTestUser(t *testing.T) (*services.MockIUserService, *UserController) {
	ctr := gomock.NewController(t)
	defer ctr.Finish()
	mockSvr := services.NewMockIUserService(ctr)
	userCtr := NewUserController(mockSvr)
	return mockSvr, userCtr
}

func TestUserController_CreateAddressByUserId(t *testing.T) {
	mockSrv, userCtr := CreateTestUser(t)

	//Mock
	mockSrv.EXPECT().CreateAddress(gomock.Any()).Return(&models.Address{
		Model:       gorm.Model{},
		Id:          0,
		UserId:      0,
		FirstName:   "",
		LastName:    "",
		PhoneNumber: "",
		Province:    "",
		District:    "",
		SubDistrict: "",
		Address:     "",
		TypeAddress: "",
		User:        models.User{},
	}, nil).Times(1)

	body := []byte("{}")

	req, err := http.NewRequest("POST", "api/users/address/:addressId", bytes.NewBuffer(body))

	if err != nil {
		t.Fatalf("Error")
	}

	w := httptest.NewRecorder()

	c, _ := gin.CreateTestContext(w)
	c.Set(config.UserID, 1)

	c.Request = req

	userCtr.CreateAddressByUserId(c)
	if w.Code != http.StatusCreated {
		t.Fatalf("201 but got %v", w.Code)
	}
} //Done

func TestUserController_ChangePassword(t *testing.T) {
	mockSrv, userCtr := CreateTestUser(t)

	mockSrv.EXPECT().UpdatePassword(gomock.Any()).Return(&models.User{
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

	bodyRequest := `{"currentPassword":"abcdefg","newPassword":"abcd123efg"}`

	req, err := http.NewRequest("PATCH", "api/users/password", strings.NewReader(bodyRequest))
	if err != nil {
		t.Fatalf("Error to create new req")
	}
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Set(config.UserID, 1)
	c.Request = req
	userCtr.ChangePassword(c)
	if w.Code != http.StatusOK {
		t.Fatalf("200 but got %v", w.Code)
	}

} //Done

func TestUserController_DeleteAddressById(t *testing.T) {
	mockSvr, userCtr := CreateTestUser(t)

	mockSvr.EXPECT().DeletedAddressById(gomock.Any()).Return(&models.Address{
		Model:       gorm.Model{},
		Id:          0,
		UserId:      0,
		FirstName:   "",
		LastName:    "",
		PhoneNumber: "",
		Province:    "",
		District:    "",
		SubDistrict: "",
		Address:     "",
		TypeAddress: "",
		User:        models.User{},
	}, nil)

	req, err := http.NewRequest("DELETE", "api/users/address/:addressId", nil)

	if err != nil {
		t.Fatalf("Error to create new request")
	}

	w := httptest.NewRecorder()

	c, _ := gin.CreateTestContext(w)

	c.Set(config.UserID, 1)

	c.Params = []gin.Param{gin.Param{Key: addressId, Value: "1"}}

	c.Request = req

	userCtr.DeleteAddressById(c)

	if w.Code != http.StatusOK {
		t.Fatalf("200 but got %v", w.Code)

	}
} //Done

func TestUserController_GetAddress(t *testing.T) {
	mockSvr, userCtr := CreateTestUser(t)

	mockSvr.EXPECT().GetAddress(gomock.Any()).Return(&[]models.Address{{
		Model:       gorm.Model{},
		Id:          0,
		UserId:      0,
		FirstName:   "",
		LastName:    "",
		PhoneNumber: "",
		Province:    "",
		District:    "",
		SubDistrict: "",
		Address:     "",
		TypeAddress: "",
		User:        models.User{},
	}}, nil).Times(1)

	req, err := http.NewRequest("GET", "api/users/address", nil)

	if err != nil {
		t.Fatalf("Error")
	}

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	c.Request = req

	c.Set(config.UserID, 1)
	userCtr.GetAddress(c)
	if w.Code != http.StatusOK {
		t.Fatalf("200 but got %v", w.Code)
	}

} //Done

func TestUserController_GetAddressById(t *testing.T) {
	mockSvr, userCtr := CreateTestUser(t)

	mockSvr.EXPECT().GetAddressById(gomock.Any()).Return(&models.Address{
		Model:       gorm.Model{},
		Id:          0,
		UserId:      0,
		FirstName:   "",
		LastName:    "",
		PhoneNumber: "",
		Province:    "",
		District:    "",
		SubDistrict: "",
		Address:     "",
		TypeAddress: "",
		User:        models.User{},
	}, nil)

	req, err := http.NewRequest("GET", "api/users/address/:addressId", nil)

	if err != nil {
		t.Fatalf("Error to create new request")
	}

	w := httptest.NewRecorder()

	c, _ := gin.CreateTestContext(w)

	c.Set(config.UserID, 1)

	c.Params = []gin.Param{gin.Param{Key: addressId, Value: "1"}}

	c.Request = req

	userCtr.GetAddressById(c)

	if w.Code != http.StatusOK {
		t.Fatalf("200 but got %v", w.Code)

	}
} //Done

func TestUserController_GetUser(t *testing.T) {
	mockSrv, userCtr := CreateTestUser(t)

	//Mock
	mockSrv.EXPECT().GetUserById(gomock.Any()).Return(&models.User{
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

	req, err := http.NewRequest("POST", "api/users", nil)

	if err != nil {
		t.Fatalf("Error")
	}

	w := httptest.NewRecorder()

	c, _ := gin.CreateTestContext(w)
	c.Set(config.UserID, 1)

	c.Request = req

	userCtr.GetUser(c)
	if w.Code != http.StatusOK {
		t.Fatalf("Expect status is 200 but got %v", w.Code)
	}
} /Done
