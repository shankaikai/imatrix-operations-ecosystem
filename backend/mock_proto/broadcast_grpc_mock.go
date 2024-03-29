// Code generated by MockGen. DO NOT EDIT.
// Source: capstone.operations_ecosystem/backend/proto (interfaces: BroadcastServicesClient)

// Package mock_proto is a generated GoMock package.
package mock_proto

import (
	context "context"
	reflect "reflect"

	operations_ecosys "capstone.operations_ecosystem/backend/proto"
	gomock "github.com/golang/mock/gomock"
	grpc "google.golang.org/grpc"
)

// MockBroadcastServicesClient is a mock of BroadcastServicesClient interface.
type MockBroadcastServicesClient struct {
	ctrl     *gomock.Controller
	recorder *MockBroadcastServicesClientMockRecorder
}

// MockBroadcastServicesClientMockRecorder is the mock recorder for MockBroadcastServicesClient.
type MockBroadcastServicesClientMockRecorder struct {
	mock *MockBroadcastServicesClient
}

// NewMockBroadcastServicesClient creates a new mock instance.
func NewMockBroadcastServicesClient(ctrl *gomock.Controller) *MockBroadcastServicesClient {
	mock := &MockBroadcastServicesClient{ctrl: ctrl}
	mock.recorder = &MockBroadcastServicesClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockBroadcastServicesClient) EXPECT() *MockBroadcastServicesClientMockRecorder {
	return m.recorder
}

// AddBroadcast mocks base method.
func (m *MockBroadcastServicesClient) AddBroadcast(arg0 context.Context, arg1 *operations_ecosys.Broadcast, arg2 ...grpc.CallOption) (*operations_ecosys.Response, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "AddBroadcast", varargs...)
	ret0, _ := ret[0].(*operations_ecosys.Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddBroadcast indicates an expected call of AddBroadcast.
func (mr *MockBroadcastServicesClientMockRecorder) AddBroadcast(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddBroadcast", reflect.TypeOf((*MockBroadcastServicesClient)(nil).AddBroadcast), varargs...)
}

// DeleteBroadcast mocks base method.
func (m *MockBroadcastServicesClient) DeleteBroadcast(arg0 context.Context, arg1 *operations_ecosys.Broadcast, arg2 ...grpc.CallOption) (*operations_ecosys.Response, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DeleteBroadcast", varargs...)
	ret0, _ := ret[0].(*operations_ecosys.Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteBroadcast indicates an expected call of DeleteBroadcast.
func (mr *MockBroadcastServicesClientMockRecorder) DeleteBroadcast(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteBroadcast", reflect.TypeOf((*MockBroadcastServicesClient)(nil).DeleteBroadcast), varargs...)
}

// FindBroadcasts mocks base method.
func (m *MockBroadcastServicesClient) FindBroadcasts(arg0 context.Context, arg1 *operations_ecosys.BroadcastQuery, arg2 ...grpc.CallOption) (operations_ecosys.BroadcastServices_FindBroadcastsClient, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "FindBroadcasts", varargs...)
	ret0, _ := ret[0].(operations_ecosys.BroadcastServices_FindBroadcastsClient)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindBroadcasts indicates an expected call of FindBroadcasts.
func (mr *MockBroadcastServicesClientMockRecorder) FindBroadcasts(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindBroadcasts", reflect.TypeOf((*MockBroadcastServicesClient)(nil).FindBroadcasts), varargs...)
}

// UpdateBroadcast mocks base method.
func (m *MockBroadcastServicesClient) UpdateBroadcast(arg0 context.Context, arg1 *operations_ecosys.Broadcast, arg2 ...grpc.CallOption) (*operations_ecosys.Response, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "UpdateBroadcast", varargs...)
	ret0, _ := ret[0].(*operations_ecosys.Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateBroadcast indicates an expected call of UpdateBroadcast.
func (mr *MockBroadcastServicesClientMockRecorder) UpdateBroadcast(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateBroadcast", reflect.TypeOf((*MockBroadcastServicesClient)(nil).UpdateBroadcast), varargs...)
}

// UpdateBroadcastRecipient mocks base method.
func (m *MockBroadcastServicesClient) UpdateBroadcastRecipient(arg0 context.Context, arg1 *operations_ecosys.BroadcastRecipient, arg2 ...grpc.CallOption) (*operations_ecosys.Response, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "UpdateBroadcastRecipient", varargs...)
	ret0, _ := ret[0].(*operations_ecosys.Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateBroadcastRecipient indicates an expected call of UpdateBroadcastRecipient.
func (mr *MockBroadcastServicesClientMockRecorder) UpdateBroadcastRecipient(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateBroadcastRecipient", reflect.TypeOf((*MockBroadcastServicesClient)(nil).UpdateBroadcastRecipient), varargs...)
}
