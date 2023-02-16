// Code generated by MockGen. DO NOT EDIT.
// Source: shopify/product-images.go

// Package mock_shopify is a generated GoMock package.
package mockshopify

import (
	shopify "github.com/MOHC-LTD/shopify/v2"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockProductImageRepository is a mock of ProductImageRepository interface
type MockProductImageRepository struct {
	ctrl     *gomock.Controller
	recorder *MockProductImageRepositoryMockRecorder
}

// MockProductImageRepositoryMockRecorder is the mock recorder for MockProductImageRepository
type MockProductImageRepositoryMockRecorder struct {
	mock *MockProductImageRepository
}

// NewMockProductImageRepository creates a new mock instance
func NewMockProductImageRepository(ctrl *gomock.Controller) *MockProductImageRepository {
	mock := &MockProductImageRepository{ctrl: ctrl}
	mock.recorder = &MockProductImageRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockProductImageRepository) EXPECT() *MockProductImageRepositoryMockRecorder {
	return m.recorder
}

// List mocks base method
func (m *MockProductImageRepository) List(productID int64, query shopify.ProductImageQuery) (shopify.ProductImages, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", productID, query)
	ret0, _ := ret[0].(shopify.ProductImages)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List
func (mr *MockProductImageRepositoryMockRecorder) List(productID, query interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockProductImageRepository)(nil).List), productID, query)
}
