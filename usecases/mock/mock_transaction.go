// Code generated by MockGen. DO NOT EDIT.
// Source: transaction.go

// Package mock is a generated GoMock package.
package mock

import (
	transfers "golang-coding-challenge/transfers"
	usecases "golang-coding-challenge/usecases"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockTransactionUseCase is a mock of TransactionUseCase interface.
type MockTransactionUseCase struct {
	ctrl     *gomock.Controller
	recorder *MockTransactionUseCaseMockRecorder
}

// MockTransactionUseCaseMockRecorder is the mock recorder for MockTransactionUseCase.
type MockTransactionUseCaseMockRecorder struct {
	mock *MockTransactionUseCase
}

// NewMockTransactionUseCase creates a new mock instance.
func NewMockTransactionUseCase(ctrl *gomock.Controller) *MockTransactionUseCase {
	mock := &MockTransactionUseCase{ctrl: ctrl}
	mock.recorder = &MockTransactionUseCaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTransactionUseCase) EXPECT() *MockTransactionUseCaseMockRecorder {
	return m.recorder
}

// CreateTransaction mocks base method.
func (m *MockTransactionUseCase) CreateTransaction(reqBody usecases.CreateTransactionRequest, userID int) (*transfers.TransactionJson, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateTransaction", reqBody, userID)
	ret0, _ := ret[0].(*transfers.TransactionJson)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateTransaction indicates an expected call of CreateTransaction.
func (mr *MockTransactionUseCaseMockRecorder) CreateTransaction(reqBody, userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateTransaction", reflect.TypeOf((*MockTransactionUseCase)(nil).CreateTransaction), reqBody, userID)
}

// GetTransactions mocks base method.
func (m *MockTransactionUseCase) GetTransactions(userID, accountID int) ([]transfers.TransactionJson, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTransactions", userID, accountID)
	ret0, _ := ret[0].([]transfers.TransactionJson)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTransactions indicates an expected call of GetTransactions.
func (mr *MockTransactionUseCaseMockRecorder) GetTransactions(userID, accountID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTransactions", reflect.TypeOf((*MockTransactionUseCase)(nil).GetTransactions), userID, accountID)
}
