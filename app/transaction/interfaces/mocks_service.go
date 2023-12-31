// Code generated by MockGen. DO NOT EDIT.
// Source: service.go

// Package transactioninterfaces is a generated GoMock package.
package transactioninterfaces

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	utils "github.com/masharpik/TransactionalSystem/app/auth/utils"
	utils0 "github.com/masharpik/TransactionalSystem/app/transaction/utils"
)

// MockITransactionService is a mock of ITransactionService interface.
type MockITransactionService struct {
	ctrl     *gomock.Controller
	recorder *MockITransactionServiceMockRecorder
}

// MockITransactionServiceMockRecorder is the mock recorder for MockITransactionService.
type MockITransactionServiceMockRecorder struct {
	mock *MockITransactionService
}

// NewMockITransactionService creates a new mock instance.
func NewMockITransactionService(ctrl *gomock.Controller) *MockITransactionService {
	mock := &MockITransactionService{ctrl: ctrl}
	mock.recorder = &MockITransactionServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockITransactionService) EXPECT() *MockITransactionServiceMockRecorder {
	return m.recorder
}

// InputMoney mocks base method.
func (m *MockITransactionService) InputMoney(arg0 string, arg1 float64) (utils.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InputMoney", arg0, arg1)
	ret0, _ := ret[0].(utils.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// InputMoney indicates an expected call of InputMoney.
func (mr *MockITransactionServiceMockRecorder) InputMoney(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InputMoney", reflect.TypeOf((*MockITransactionService)(nil).InputMoney), arg0, arg1)
}

// OutputMoney mocks base method.
func (m *MockITransactionService) OutputMoney(arg0 string, arg1 float64, arg2 string) (utils0.OutputTransactionResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "OutputMoney", arg0, arg1, arg2)
	ret0, _ := ret[0].(utils0.OutputTransactionResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// OutputMoney indicates an expected call of OutputMoney.
func (mr *MockITransactionServiceMockRecorder) OutputMoney(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OutputMoney", reflect.TypeOf((*MockITransactionService)(nil).OutputMoney), arg0, arg1, arg2)
}
