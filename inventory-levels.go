// Code generated by MockGen. DO NOT EDIT.
// Source: inventory-levels.go

// Package mockshopify is a generated GoMock package.
package mockshopify

import (
	reflect "reflect"

	shopify "github.com/MOHC-LTD/shopify/v2"
	gomock "github.com/golang/mock/gomock"
)

// MockInventoryLevelRepository is a mock of InventoryLevelRepository interface
type MockInventoryLevelRepository struct {
	ctrl     *gomock.Controller
	recorder *MockInventoryLevelRepositoryMockRecorder
}

// MockInventoryLevelRepositoryMockRecorder is the mock recorder for MockInventoryLevelRepository
type MockInventoryLevelRepositoryMockRecorder struct {
	mock *MockInventoryLevelRepository
}

// NewMockInventoryLevelRepository creates a new mock instance
func NewMockInventoryLevelRepository(ctrl *gomock.Controller) *MockInventoryLevelRepository {
	mock := &MockInventoryLevelRepository{ctrl: ctrl}
	mock.recorder = &MockInventoryLevelRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockInventoryLevelRepository) EXPECT() *MockInventoryLevelRepositoryMockRecorder {
	return m.recorder
}

// Set mocks base method
func (m *MockInventoryLevelRepository) Set(inventoryItemID, locationID int64, quantity int) (shopify.InventoryLevel, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Set", inventoryItemID, locationID, quantity)
	ret0, _ := ret[0].(shopify.InventoryLevel)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Set indicates an expected call of Set
func (mr *MockInventoryLevelRepositoryMockRecorder) Set(inventoryItemID, locationID, quantity interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Set", reflect.TypeOf((*MockInventoryLevelRepository)(nil).Set), inventoryItemID, locationID, quantity)
}
