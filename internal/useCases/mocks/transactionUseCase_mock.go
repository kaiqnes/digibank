// Code generated by MockGen. DO NOT EDIT.
// Source: ./transactionUseCase.go

// Package mock_useCases is a generated GoMock package.
package mock_useCases

import (
	entities "digibank/internal/domain/entities"
	errorx "digibank/internal/frameworks/errorx"
	dto "digibank/internal/interfaceAdapters/dto"
	reflect "reflect"

	gin "github.com/gin-gonic/gin"
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
func (m *MockTransactionUseCase) CreateTransaction(ctx *gin.Context, transactionContent dto.CreateTransactionInput) (entities.Transaction, errorx.Errorx) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateTransaction", ctx, transactionContent)
	ret0, _ := ret[0].(entities.Transaction)
	ret1, _ := ret[1].(errorx.Errorx)
	return ret0, ret1
}

// CreateTransaction indicates an expected call of CreateTransaction.
func (mr *MockTransactionUseCaseMockRecorder) CreateTransaction(ctx, transactionContent interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateTransaction", reflect.TypeOf((*MockTransactionUseCase)(nil).CreateTransaction), ctx, transactionContent)
}
