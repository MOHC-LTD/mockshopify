// Code generated by MockGen. DO NOT EDIT.
// Source: ../../src/shopify/customer-addresses.go

// Package mockshopify is a generated GoMock package.
package mockshopify

import (
	v2 "github.com/MOHC-LTD/shopify/v2"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockCustomerAddressRepository is a mock of CustomerAddressRepository interface
type MockCustomerAddressRepository struct {
	ctrl     *gomock.Controller
	recorder *MockCustomerAddressRepositoryMockRecorder
}

// MockCustomerAddressRepositoryMockRecorder is the mock recorder for MockCustomerAddressRepository
type MockCustomerAddressRepositoryMockRecorder struct {
	mock *MockCustomerAddressRepository
}

// NewMockCustomerAddressRepository creates a new mock instance
func NewMockCustomerAddressRepository(ctrl *gomock.Controller) *MockCustomerAddressRepository {
	mock := &MockCustomerAddressRepository{ctrl: ctrl}
	mock.recorder = &MockCustomerAddressRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockCustomerAddressRepository) EXPECT() *MockCustomerAddressRepositoryMockRecorder {
	return m.recorder
}

// List mocks base method
func (m *MockCustomerAddressRepository) List(id int64) (v2.CustomerAddresses, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", id)
	ret0, _ := ret[0].(v2.CustomerAddresses)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List
func (mr *MockCustomerAddressRepositoryMockRecorder) List(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockCustomerAddressRepository)(nil).List), id)
}

// Create mocks base method
func (m *MockCustomerAddressRepository) Create(id int64, address v2.CustomerAddress) (v2.CustomerAddress, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", id, address)
	ret0, _ := ret[0].(v2.CustomerAddress)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create
func (mr *MockCustomerAddressRepositoryMockRecorder) Create(id, address interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockCustomerAddressRepository)(nil).Create), id, address)
}

// Delete mocks base method
func (m *MockCustomerAddressRepository) Delete(id, addressID int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", id, addressID)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete
func (mr *MockCustomerAddressRepositoryMockRecorder) Delete(id, addressID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockCustomerAddressRepository)(nil).Delete), id, addressID)
}

// Update mocks base method
func (m *MockCustomerAddressRepository) Update(id int64, address v2.CustomerAddress) (v2.CustomerAddress, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", id, address)
	ret0, _ := ret[0].(v2.CustomerAddress)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update
func (mr *MockCustomerAddressRepositoryMockRecorder) Update(id, address interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockCustomerAddressRepository)(nil).Update), id, address)
}

// SetDefault mocks base method
func (m *MockCustomerAddressRepository) SetDefault(id, addressID int64) (v2.CustomerAddress, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetDefault", id, addressID)
	ret0, _ := ret[0].(v2.CustomerAddress)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SetDefault indicates an expected call of SetDefault
func (mr *MockCustomerAddressRepositoryMockRecorder) SetDefault(id, addressID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetDefault", reflect.TypeOf((*MockCustomerAddressRepository)(nil).SetDefault), id, addressID)
}
