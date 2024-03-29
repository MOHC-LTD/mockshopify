// Code generated by MockGen. DO NOT EDIT.
// Source: metafields.go

// Package mock_shopify is a generated GoMock package.
package mockshopify

import (
	reflect "reflect"

	shopify "github.com/MOHC-LTD/shopify/v2"
	gomock "github.com/golang/mock/gomock"
)

// MockMetafieldRepository is a mock of MetafieldRepository interface.
type MockMetafieldRepository struct {
	ctrl     *gomock.Controller
	recorder *MockMetafieldRepositoryMockRecorder
}

// MockMetafieldRepositoryMockRecorder is the mock recorder for MockMetafieldRepository.
type MockMetafieldRepositoryMockRecorder struct {
	mock *MockMetafieldRepository
}

// NewMockMetafieldRepository creates a new mock instance.
func NewMockMetafieldRepository(ctrl *gomock.Controller) *MockMetafieldRepository {
	mock := &MockMetafieldRepository{ctrl: ctrl}
	mock.recorder = &MockMetafieldRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMetafieldRepository) EXPECT() *MockMetafieldRepositoryMockRecorder {
	return m.recorder
}

// List mocks base method.
func (m *MockMetafieldRepository) List(query shopify.MetafieldQuery) (shopify.Metafields, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", query)
	ret0, _ := ret[0].(shopify.Metafields)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List.
func (mr *MockMetafieldRepositoryMockRecorder) List(query interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockMetafieldRepository)(nil).List), query)
}
