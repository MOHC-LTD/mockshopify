// Code generated by MockGen. DO NOT EDIT.
// Source: ../shopify/products.go

// Package mockshopify is a generated GoMock package.
package mockshopify

import (
	context "context"
	reflect "reflect"

	shopify "github.com/MOHC-LTD/shopify/v2"
	gomock "github.com/golang/mock/gomock"
)

// MockProductRepository is a mock of ProductRepository interface.
type MockProductRepository struct {
	ctrl     *gomock.Controller
	recorder *MockProductRepositoryMockRecorder
}

// MockProductRepositoryMockRecorder is the mock recorder for MockProductRepository.
type MockProductRepositoryMockRecorder struct {
	mock *MockProductRepository
}

// NewMockProductRepository creates a new mock instance.
func NewMockProductRepository(ctrl *gomock.Controller) *MockProductRepository {
	mock := &MockProductRepository{ctrl: ctrl}
	mock.recorder = &MockProductRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockProductRepository) EXPECT() *MockProductRepositoryMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockProductRepository) Create(product shopify.Product) (shopify.Product, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", product)
	ret0, _ := ret[0].(shopify.Product)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockProductRepositoryMockRecorder) Create(product interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockProductRepository)(nil).Create), product)
}

// Delete mocks base method.
func (m *MockProductRepository) Delete(productID int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", productID)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockProductRepositoryMockRecorder) Delete(productID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockProductRepository)(nil).Delete), productID)
}

// Get mocks base method.
func (m *MockProductRepository) Get(id int64) (shopify.Product, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", id)
	ret0, _ := ret[0].(shopify.Product)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockProductRepositoryMockRecorder) Get(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockProductRepository)(nil).Get), id)
}

// List mocks base method.
func (m *MockProductRepository) List(query shopify.ProductQuery) (shopify.Products, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", query)
	ret0, _ := ret[0].(shopify.Products)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List.
func (mr *MockProductRepositoryMockRecorder) List(query interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockProductRepository)(nil).List), query)
}

// Update mocks base method.
func (m *MockProductRepository) Update(product shopify.Product) (shopify.Product, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", product)
	ret0, _ := ret[0].(shopify.Product)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update.
func (mr *MockProductRepositoryMockRecorder) Update(product interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockProductRepository)(nil).Update), product)
}

// MockProductCreator is a mock of ProductCreator interface.
type MockProductCreator struct {
	ctrl     *gomock.Controller
	recorder *MockProductCreatorMockRecorder
}

// MockProductCreatorMockRecorder is the mock recorder for MockProductCreator.
type MockProductCreatorMockRecorder struct {
	mock *MockProductCreator
}

// NewMockProductCreator creates a new mock instance.
func NewMockProductCreator(ctrl *gomock.Controller) *MockProductCreator {
	mock := &MockProductCreator{ctrl: ctrl}
	mock.recorder = &MockProductCreatorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockProductCreator) EXPECT() *MockProductCreatorMockRecorder {
	return m.recorder
}

// Save mocks base method.
func (m *MockProductCreator) Save(arg0 context.Context, arg1 shopify.Product) (shopify.Product, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Save", arg0, arg1)
	ret0, _ := ret[0].(shopify.Product)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Save indicates an expected call of Save.
func (mr *MockProductCreatorMockRecorder) Save(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Save", reflect.TypeOf((*MockProductCreator)(nil).Save), arg0, arg1)
}
