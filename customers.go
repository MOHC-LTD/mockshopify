// Code generated by MockGen. DO NOT EDIT.
// Source: shopify/customers.go

// Package mockshopify is a generated GoMock package.
package mockshopify

import (
	reflect "reflect"

	shopify "github.com/MOHC-LTD/shopify"
	gomock "github.com/golang/mock/gomock"
)

// MockCustomerRepository is a mock of CustomerRepository interface.
type MockCustomerRepository struct {
	ctrl     *gomock.Controller
	recorder *MockCustomerRepositoryMockRecorder
}

// MockCustomerRepositoryMockRecorder is the mock recorder for MockCustomerRepository.
type MockCustomerRepositoryMockRecorder struct {
	mock *MockCustomerRepository
}

// NewMockCustomerRepository creates a new mock instance.
func NewMockCustomerRepository(ctrl *gomock.Controller) *MockCustomerRepository {
	mock := &MockCustomerRepository{ctrl: ctrl}
	mock.recorder = &MockCustomerRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCustomerRepository) EXPECT() *MockCustomerRepositoryMockRecorder {
	return m.recorder
}

// Get mocks base method.
func (m *MockCustomerRepository) Get(id int64) (shopify.Customer, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", id)
	ret0, _ := ret[0].(shopify.Customer)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockCustomerRepositoryMockRecorder) Get(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockCustomerRepository)(nil).Get), id)
}

// GetByQuery mocks base method.
func (m *MockCustomerRepository) GetByQuery(fields []string, query shopify.CustomerSearchQuery) (shopify.Customers, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByQuery", fields, query)
	ret0, _ := ret[0].(shopify.Customers)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByQuery indicates an expected call of GetByQuery.
func (mr *MockCustomerRepositoryMockRecorder) GetByQuery(fields, query interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByQuery", reflect.TypeOf((*MockCustomerRepository)(nil).GetByQuery), fields, query)
}

// Update mocks base method.
func (m *MockCustomerRepository) Update(customer shopify.Customer) (shopify.Customer, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", customer)
	ret0, _ := ret[0].(shopify.Customer)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update.
func (mr *MockCustomerRepositoryMockRecorder) Update(customer interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockCustomerRepository)(nil).Update), customer)
}