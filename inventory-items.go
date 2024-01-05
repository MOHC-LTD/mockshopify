// Code generated by MockGen. DO NOT EDIT.
// Source: inventory-items.go

// Package mock_shopify is a generated GoMock package.
package mockshopify

import (
	reflect "reflect"

	v2 "github.com/MOHC-LTD/shopify/v2"
	gomock "github.com/golang/mock/gomock"
)

// MockInventoryItemRepository is a mock of InventoryItemRepository interface.
type MockInventoryItemRepository struct {
	ctrl     *gomock.Controller
	recorder *MockInventoryItemRepositoryMockRecorder
}

// MockInventoryItemRepositoryMockRecorder is the mock recorder for MockInventoryItemRepository.
type MockInventoryItemRepositoryMockRecorder struct {
	mock *MockInventoryItemRepository
}

// NewMockInventoryItemRepository creates a new mock instance.
func NewMockInventoryItemRepository(ctrl *gomock.Controller) *MockInventoryItemRepository {
	mock := &MockInventoryItemRepository{ctrl: ctrl}
	mock.recorder = &MockInventoryItemRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockInventoryItemRepository) EXPECT() *MockInventoryItemRepositoryMockRecorder {
	return m.recorder
}

// Get mocks base method.
func (m *MockInventoryItemRepository) Get(id int64) (v2.InventoryItem, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", id)
	ret0, _ := ret[0].(v2.InventoryItem)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockInventoryItemRepositoryMockRecorder) Get(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockInventoryItemRepository)(nil).Get), id)
}
