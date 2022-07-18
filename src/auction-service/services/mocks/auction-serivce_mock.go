// Code generated by MockGen. DO NOT EDIT.
// Source: ./src/auction-service/services/auction-service.go

// Package services is a generated GoMock package.
package services

import (
	dtos "chilindo/src/auction-service/dtos"
	models "chilindo/src/auction-service/models"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockIAuctionService is a mock of IAuctionService interface.
type MockIAuctionService struct {
	ctrl     *gomock.Controller
	recorder *MockIAuctionServiceMockRecorder
}

// MockIAuctionServiceMockRecorder is the mock recorder for MockIAuctionService.
type MockIAuctionServiceMockRecorder struct {
	mock *MockIAuctionService
}

// NewMockIAuctionService creates a new mock instance.
func NewMockIAuctionService(ctrl *gomock.Controller) *MockIAuctionService {
	mock := &MockIAuctionService{ctrl: ctrl}
	mock.recorder = &MockIAuctionServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIAuctionService) EXPECT() *MockIAuctionServiceMockRecorder {
	return m.recorder
}

// CreateAuction mocks base method.
func (m *MockIAuctionService) CreateAuction(dto *dtos.CreateAuctionDTO) (*models.Auction, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateAuction", dto)
	ret0, _ := ret[0].(*models.Auction)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateAuction indicates an expected call of CreateAuction.
func (mr *MockIAuctionServiceMockRecorder) CreateAuction(dto interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateAuction", reflect.TypeOf((*MockIAuctionService)(nil).CreateAuction), dto)
}

// GetAuctionById mocks base method.
func (m *MockIAuctionService) GetAuctionById(dto *dtos.AuctionIdDTO) (*models.Auction, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAuctionById", dto)
	ret0, _ := ret[0].(*models.Auction)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAuctionById indicates an expected call of GetAuctionById.
func (mr *MockIAuctionServiceMockRecorder) GetAuctionById(dto interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAuctionById", reflect.TypeOf((*MockIAuctionService)(nil).GetAuctionById), dto)
}

// GetAuctions mocks base method.
func (m *MockIAuctionService) GetAuctions() (*[]models.Auction, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAuctions")
	ret0, _ := ret[0].(*[]models.Auction)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAuctions indicates an expected call of GetAuctions.
func (mr *MockIAuctionServiceMockRecorder) GetAuctions() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAuctions", reflect.TypeOf((*MockIAuctionService)(nil).GetAuctions))
}
