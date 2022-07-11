package services

import (
	"chilindo/src/product-service/dtos"
	"chilindo/src/product-service/models"
	repository "chilindo/src/product-service/repository/mocks"
	"github.com/golang/mock/gomock"
	"gorm.io/gorm"
	"testing"
)

func TestOptionService_CreateOption(t *testing.T) {
	ctr := gomock.NewController(t)
	defer ctr.Finish()
	repo := repository.NewMockIOptionRepository(ctr)
	optionSvr := NewOptionService(repo)

	//Mock repository
	repo.EXPECT().CreateOption(gomock.Any()).Return(&models.Option{
		Model:     gorm.Model{},
		Id:        0,
		ProductId: "",
		Link:      "",
		Product:   models.Product{},
	}, nil).Times(1)
	var dto *dtos.CreateOption
	_, err := optionSvr.CreateOption(dto)
	if err != nil {
		t.Fatalf("Error")
	}
} //Done
