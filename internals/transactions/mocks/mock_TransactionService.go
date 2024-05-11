// Code generated by MockGen. DO NOT EDIT.
// Source: internals/transactions/services/TransactionService.go
//
// Generated by this command:
//
//	mockgen -source=internals/transactions/services/TransactionService.go -destination internals/transactions/mocks/mock_TransactionService.go -package mocks
//

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	models "github.com/Sampriti-Mitra/transactions/internals/transactions/models"
	utils "github.com/Sampriti-Mitra/transactions/internals/utils"
	gomock "go.uber.org/mock/gomock"
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

// CreateAccount mocks base method.
func (m *MockITransactionService) CreateAccount(acc *models.Account) (*models.Account, *utils.CustError) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateAccount", acc)
	ret0, _ := ret[0].(*models.Account)
	ret1, _ := ret[1].(*utils.CustError)
	return ret0, ret1
}

// CreateAccount indicates an expected call of CreateAccount.
func (mr *MockITransactionServiceMockRecorder) CreateAccount(acc any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateAccount", reflect.TypeOf((*MockITransactionService)(nil).CreateAccount), acc)
}

// CreateTransaction mocks base method.
func (m *MockITransactionService) CreateTransaction(txn *models.Transaction) (*models.Transaction, *utils.CustError) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateTransaction", txn)
	ret0, _ := ret[0].(*models.Transaction)
	ret1, _ := ret[1].(*utils.CustError)
	return ret0, ret1
}

// CreateTransaction indicates an expected call of CreateTransaction.
func (mr *MockITransactionServiceMockRecorder) CreateTransaction(txn any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateTransaction", reflect.TypeOf((*MockITransactionService)(nil).CreateTransaction), txn)
}

// FetchAccount mocks base method.
func (m *MockITransactionService) FetchAccount(accId int64) (*models.Account, *utils.CustError) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FetchAccount", accId)
	ret0, _ := ret[0].(*models.Account)
	ret1, _ := ret[1].(*utils.CustError)
	return ret0, ret1
}

// FetchAccount indicates an expected call of FetchAccount.
func (mr *MockITransactionServiceMockRecorder) FetchAccount(accId any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FetchAccount", reflect.TypeOf((*MockITransactionService)(nil).FetchAccount), accId)
}