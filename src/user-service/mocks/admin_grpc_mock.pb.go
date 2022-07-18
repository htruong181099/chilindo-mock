// Code generated by MockGen. DO NOT EDIT.
// Source: .\pkg\pb\admin\admin_grpc.pb.go

// Package admin is a generated GoMock package.
package admin

import (
	admin "chilindo/pkg/pb/admin"
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	grpc "google.golang.org/grpc"
)

// MockAdminServiceClient is a mock of AdminServiceClient interface.
type MockAdminServiceClient struct {
	ctrl     *gomock.Controller
	recorder *MockAdminServiceClientMockRecorder
}

// MockAdminServiceClientMockRecorder is the mock recorder for MockAdminServiceClient.
type MockAdminServiceClientMockRecorder struct {
	mock *MockAdminServiceClient
}

// NewMockAdminServiceClient creates a new mock instance.
func NewMockAdminServiceClient(ctrl *gomock.Controller) *MockAdminServiceClient {
	mock := &MockAdminServiceClient{ctrl: ctrl}
	mock.recorder = &MockAdminServiceClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAdminServiceClient) EXPECT() *MockAdminServiceClientMockRecorder {
	return m.recorder
}

// CheckIsAdmin mocks base method.
func (m *MockAdminServiceClient) CheckIsAdmin(ctx context.Context, in *admin.CheckIsAdminRequest, opts ...grpc.CallOption) (*admin.CheckIsAdminResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "CheckIsAdmin", varargs...)
	ret0, _ := ret[0].(*admin.CheckIsAdminResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CheckIsAdmin indicates an expected call of CheckIsAdmin.
func (mr *MockAdminServiceClientMockRecorder) CheckIsAdmin(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckIsAdmin", reflect.TypeOf((*MockAdminServiceClient)(nil).CheckIsAdmin), varargs...)
}

// CheckUserAuth mocks base method.
func (m *MockAdminServiceClient) CheckUserAuth(ctx context.Context, in *admin.CheckUserAuthRequest, opts ...grpc.CallOption) (*admin.CheckUserAuthResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "CheckUserAuth", varargs...)
	ret0, _ := ret[0].(*admin.CheckUserAuthResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CheckUserAuth indicates an expected call of CheckUserAuth.
func (mr *MockAdminServiceClientMockRecorder) CheckUserAuth(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckUserAuth", reflect.TypeOf((*MockAdminServiceClient)(nil).CheckUserAuth), varargs...)
}

// MockAdminServiceServer is a mock of AdminServiceServer interface.
type MockAdminServiceServer struct {
	ctrl     *gomock.Controller
	recorder *MockAdminServiceServerMockRecorder
}

// MockAdminServiceServerMockRecorder is the mock recorder for MockAdminServiceServer.
type MockAdminServiceServerMockRecorder struct {
	mock *MockAdminServiceServer
}

// NewMockAdminServiceServer creates a new mock instance.
func NewMockAdminServiceServer(ctrl *gomock.Controller) *MockAdminServiceServer {
	mock := &MockAdminServiceServer{ctrl: ctrl}
	mock.recorder = &MockAdminServiceServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAdminServiceServer) EXPECT() *MockAdminServiceServerMockRecorder {
	return m.recorder
}

// CheckIsAdmin mocks base method.
func (m *MockAdminServiceServer) CheckIsAdmin(arg0 context.Context, arg1 *admin.CheckIsAdminRequest) (*admin.CheckIsAdminResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CheckIsAdmin", arg0, arg1)
	ret0, _ := ret[0].(*admin.CheckIsAdminResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CheckIsAdmin indicates an expected call of CheckIsAdmin.
func (mr *MockAdminServiceServerMockRecorder) CheckIsAdmin(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckIsAdmin", reflect.TypeOf((*MockAdminServiceServer)(nil).CheckIsAdmin), arg0, arg1)
}

// CheckUserAuth mocks base method.
func (m *MockAdminServiceServer) CheckUserAuth(arg0 context.Context, arg1 *admin.CheckUserAuthRequest) (*admin.CheckUserAuthResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CheckUserAuth", arg0, arg1)
	ret0, _ := ret[0].(*admin.CheckUserAuthResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CheckUserAuth indicates an expected call of CheckUserAuth.
func (mr *MockAdminServiceServerMockRecorder) CheckUserAuth(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckUserAuth", reflect.TypeOf((*MockAdminServiceServer)(nil).CheckUserAuth), arg0, arg1)
}

// mustEmbedUnimplementedAdminServiceServer mocks base method.
func (m *MockAdminServiceServer) mustEmbedUnimplementedAdminServiceServer() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "mustEmbedUnimplementedAdminServiceServer")
}

// mustEmbedUnimplementedAdminServiceServer indicates an expected call of mustEmbedUnimplementedAdminServiceServer.
func (mr *MockAdminServiceServerMockRecorder) mustEmbedUnimplementedAdminServiceServer() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "mustEmbedUnimplementedAdminServiceServer", reflect.TypeOf((*MockAdminServiceServer)(nil).mustEmbedUnimplementedAdminServiceServer))
}

// MockUnsafeAdminServiceServer is a mock of UnsafeAdminServiceServer interface.
type MockUnsafeAdminServiceServer struct {
	ctrl     *gomock.Controller
	recorder *MockUnsafeAdminServiceServerMockRecorder
}

// MockUnsafeAdminServiceServerMockRecorder is the mock recorder for MockUnsafeAdminServiceServer.
type MockUnsafeAdminServiceServerMockRecorder struct {
	mock *MockUnsafeAdminServiceServer
}

// NewMockUnsafeAdminServiceServer creates a new mock instance.
func NewMockUnsafeAdminServiceServer(ctrl *gomock.Controller) *MockUnsafeAdminServiceServer {
	mock := &MockUnsafeAdminServiceServer{ctrl: ctrl}
	mock.recorder = &MockUnsafeAdminServiceServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUnsafeAdminServiceServer) EXPECT() *MockUnsafeAdminServiceServerMockRecorder {
	return m.recorder
}

// mustEmbedUnimplementedAdminServiceServer mocks base method.
func (m *MockUnsafeAdminServiceServer) mustEmbedUnimplementedAdminServiceServer() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "mustEmbedUnimplementedAdminServiceServer")
}

// mustEmbedUnimplementedAdminServiceServer indicates an expected call of mustEmbedUnimplementedAdminServiceServer.
func (mr *MockUnsafeAdminServiceServerMockRecorder) mustEmbedUnimplementedAdminServiceServer() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "mustEmbedUnimplementedAdminServiceServer", reflect.TypeOf((*MockUnsafeAdminServiceServer)(nil).mustEmbedUnimplementedAdminServiceServer))
}