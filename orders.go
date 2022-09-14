// Code generated by MockGen. DO NOT EDIT.
// Source: shopify/orders.go

// Package mockshopify is a generated GoMock package.
package mockshopify

import (
	reflect "reflect"

	shopify "github.com/MOHC-LTD/shopify"
	gomock "github.com/golang/mock/gomock"
)

// MockOrderRepository is a mock of OrderRepository interface.
type MockOrderRepository struct {
	ctrl     *gomock.Controller
	recorder *MockOrderRepositoryMockRecorder
}

// MockOrderRepositoryMockRecorder is the mock recorder for MockOrderRepository.
type MockOrderRepositoryMockRecorder struct {
	mock *MockOrderRepository
}

// NewMockOrderRepository creates a new mock instance.
func NewMockOrderRepository(ctrl *gomock.Controller) *MockOrderRepository {
	mock := &MockOrderRepository{ctrl: ctrl}
	mock.recorder = &MockOrderRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockOrderRepository) EXPECT() *MockOrderRepositoryMockRecorder {
	return m.recorder
}

// Close mocks base method.
func (m *MockOrderRepository) Close(id int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close", id)
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close.
func (mr *MockOrderRepositoryMockRecorder) Close(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockOrderRepository)(nil).Close), id)
}

// Create mocks base method.
func (m *MockOrderRepository) Create(order shopify.Order) (shopify.Order, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", order)
	ret0, _ := ret[0].(shopify.Order)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockOrderRepositoryMockRecorder) Create(order interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockOrderRepository)(nil).Create), order)
}

// CreateMetafield mocks base method.
func (m *MockOrderRepository) CreateMetafield(orderID int64, metafield shopify.Metafield) (shopify.Metafield, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateMetafield", orderID, metafield)
	ret0, _ := ret[0].(shopify.Metafield)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateMetafield indicates an expected call of CreateMetafield.
func (mr *MockOrderRepositoryMockRecorder) CreateMetafield(orderID, metafield interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateMetafield", reflect.TypeOf((*MockOrderRepository)(nil).CreateMetafield), orderID, metafield)
}

// Get mocks base method.
func (m *MockOrderRepository) Get(id int64) (shopify.Order, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", id)
	ret0, _ := ret[0].(shopify.Order)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockOrderRepositoryMockRecorder) Get(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockOrderRepository)(nil).Get), id)
}

// List mocks base method.
func (m *MockOrderRepository) List(query shopify.OrderQuery) (shopify.Orders, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", query)
	ret0, _ := ret[0].(shopify.Orders)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List.
func (mr *MockOrderRepositoryMockRecorder) List(query interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockOrderRepository)(nil).List), query)
}

// Open mocks base method.
func (m *MockOrderRepository) Open(id int64) (shopify.Order, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Open", id)
	ret0, _ := ret[0].(shopify.Order)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Open indicates an expected call of Open.
func (mr *MockOrderRepositoryMockRecorder) Open(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Open", reflect.TypeOf((*MockOrderRepository)(nil).Open), id)
}

// Update mocks base method.
func (m *MockOrderRepository) Update(order shopify.Order) (shopify.Order, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", order)
	ret0, _ := ret[0].(shopify.Order)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update.
func (mr *MockOrderRepositoryMockRecorder) Update(order interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockOrderRepository)(nil).Update), order)
}

// UpdateMetafield mocks base method.
func (m *MockOrderRepository) UpdateMetafield(orderID int64, metafield shopify.Metafield) (shopify.Metafield, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateMetafield", orderID, metafield)
	ret0, _ := ret[0].(shopify.Metafield)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateMetafield indicates an expected call of UpdateMetafield.
func (mr *MockOrderRepositoryMockRecorder) UpdateMetafield(orderID, metafield interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateMetafield", reflect.TypeOf((*MockOrderRepository)(nil).UpdateMetafield), orderID, metafield)
}
