// Code generated by MockGen. DO NOT EDIT.
// Source: delivery.go

// Package authinterfaces is a generated GoMock package.
package authinterfaces

import (
	http "net/http"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockIAuthDelivery is a mock of IAuthDelivery interface.
type MockIAuthDelivery struct {
	ctrl     *gomock.Controller
	recorder *MockIAuthDeliveryMockRecorder
}

// MockIAuthDeliveryMockRecorder is the mock recorder for MockIAuthDelivery.
type MockIAuthDeliveryMockRecorder struct {
	mock *MockIAuthDelivery
}

// NewMockIAuthDelivery creates a new mock instance.
func NewMockIAuthDelivery(ctrl *gomock.Controller) *MockIAuthDelivery {
	mock := &MockIAuthDelivery{ctrl: ctrl}
	mock.recorder = &MockIAuthDeliveryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIAuthDelivery) EXPECT() *MockIAuthDeliveryMockRecorder {
	return m.recorder
}

// AuthHandler mocks base method.
func (m *MockIAuthDelivery) AuthHandler(arg0 http.ResponseWriter, arg1 *http.Request) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "AuthHandler", arg0, arg1)
}

// AuthHandler indicates an expected call of AuthHandler.
func (mr *MockIAuthDeliveryMockRecorder) AuthHandler(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AuthHandler", reflect.TypeOf((*MockIAuthDelivery)(nil).AuthHandler), arg0, arg1)
}