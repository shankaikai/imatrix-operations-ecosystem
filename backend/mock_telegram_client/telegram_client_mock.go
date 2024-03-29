// Code generated by MockGen. DO NOT EDIT.
// Source: capstone.operations_ecosystem/backend/telegram_client (interfaces: TelegramClientInterface)

// Package mock_telegram_client is a generated GoMock package.
package mock_telegram_client

import (
	reflect "reflect"

	operations_ecosys "capstone.operations_ecosystem/backend/proto"
	gomock "github.com/golang/mock/gomock"
	grpc "google.golang.org/grpc"
)

// MockTelegramClientInterface is a mock of TelegramClientInterface interface.
type MockTelegramClientInterface struct {
	ctrl     *gomock.Controller
	recorder *MockTelegramClientInterfaceMockRecorder
}

// MockTelegramClientInterfaceMockRecorder is the mock recorder for MockTelegramClientInterface.
type MockTelegramClientInterfaceMockRecorder struct {
	mock *MockTelegramClientInterface
}

// NewMockTelegramClientInterface creates a new mock instance.
func NewMockTelegramClientInterface(ctrl *gomock.Controller) *MockTelegramClientInterface {
	mock := &MockTelegramClientInterface{ctrl: ctrl}
	mock.recorder = &MockTelegramClientInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTelegramClientInterface) EXPECT() *MockTelegramClientInterfaceMockRecorder {
	return m.recorder
}

// CreateBroadcastClient mocks base method.
func (m *MockTelegramClientInterface) CreateBroadcastClient(arg0 *string, arg1 *int) (operations_ecosys.BroadcastServicesClient, *grpc.ClientConn) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateBroadcastClient", arg0, arg1)
	ret0, _ := ret[0].(operations_ecosys.BroadcastServicesClient)
	ret1, _ := ret[1].(*grpc.ClientConn)
	return ret0, ret1
}

// CreateBroadcastClient indicates an expected call of CreateBroadcastClient.
func (mr *MockTelegramClientInterfaceMockRecorder) CreateBroadcastClient(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateBroadcastClient", reflect.TypeOf((*MockTelegramClientInterface)(nil).CreateBroadcastClient), arg0, arg1)
}

// InsertBroadcast mocks base method.
func (m *MockTelegramClientInterface) InsertBroadcast(arg0 *string, arg1 *int, arg2 *operations_ecosys.Broadcast) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InsertBroadcast", arg0, arg1, arg2)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// InsertBroadcast indicates an expected call of InsertBroadcast.
func (mr *MockTelegramClientInterfaceMockRecorder) InsertBroadcast(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InsertBroadcast", reflect.TypeOf((*MockTelegramClientInterface)(nil).InsertBroadcast), arg0, arg1, arg2)
}
