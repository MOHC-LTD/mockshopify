// Code generated by MockGen. DO NOT EDIT.
// Source: shopify/variants.go

// Package mockshopify is a generated GoMock package.
package mockshopify

import (
	reflect "reflect"

	shopify "github.com/MOHC-LTD/shopify/v2"
	gomock "github.com/golang/mock/gomock"
)

// MockVariantRepository is a mock of VariantRepository interface.
type MockVariantRepository struct {
	ctrl     *gomock.Controller
	recorder *MockVariantRepositoryMockRecorder
}

// MockVariantRepositoryMockRecorder is the mock recorder for MockVariantRepository.
type MockVariantRepositoryMockRecorder struct {
	mock *MockVariantRepository
}

// NewMockVariantRepository creates a new mock instance.
func NewMockVariantRepository(ctrl *gomock.Controller) *MockVariantRepository {
	mock := &MockVariantRepository{ctrl: ctrl}
	mock.recorder = &MockVariantRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockVariantRepository) EXPECT() *MockVariantRepositoryMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockVariantRepository) Create(productID int64, variant shopify.Variant) (shopify.Variant, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", productID, variant)
	ret0, _ := ret[0].(shopify.Variant)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockVariantRepositoryMockRecorder) Create(productID, variant interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockVariantRepository)(nil).Create), productID, variant)
}

// Get mocks base method.
func (m *MockVariantRepository) Get(id int64) (shopify.Variant, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", id)
	ret0, _ := ret[0].(shopify.Variant)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockVariantRepositoryMockRecorder) Get(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockVariantRepository)(nil).Get), id)
}
