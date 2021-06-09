// Code generated by MockGen. DO NOT EDIT.
// Source: internal/shopify/fulfillment-events.go

// Package mockshopify is a generated GoMock package.
package mockshopify

import (
	shopify "github.com/MOHC-LTD/shopify"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockFulfillmentEventRepository is a mock of FulfillmentEventRepository interface
type MockFulfillmentEventRepository struct {
	ctrl     *gomock.Controller
	recorder *MockFulfillmentEventRepositoryMockRecorder
}

// MockFulfillmentEventRepositoryMockRecorder is the mock recorder for MockFulfillmentEventRepository
type MockFulfillmentEventRepositoryMockRecorder struct {
	mock *MockFulfillmentEventRepository
}

// NewMockFulfillmentEventRepository creates a new mock instance
func NewMockFulfillmentEventRepository(ctrl *gomock.Controller) *MockFulfillmentEventRepository {
	mock := &MockFulfillmentEventRepository{ctrl: ctrl}
	mock.recorder = &MockFulfillmentEventRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockFulfillmentEventRepository) EXPECT() *MockFulfillmentEventRepositoryMockRecorder {
	return m.recorder
}

// Create mocks base method
func (m *MockFulfillmentEventRepository) Create(orderID, fulfillmentID int64, event shopify.FulfillmentEvent) (shopify.FulfillmentEvent, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", orderID, fulfillmentID, event)
	ret0, _ := ret[0].(shopify.FulfillmentEvent)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create
func (mr *MockFulfillmentEventRepositoryMockRecorder) Create(orderID, fulfillmentID, event interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockFulfillmentEventRepository)(nil).Create), orderID, fulfillmentID, event)
}

// Delete mocks base method
func (m *MockFulfillmentEventRepository) Delete(orderID, fulfillmentID, eventID int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", orderID, fulfillmentID, eventID)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete
func (mr *MockFulfillmentEventRepositoryMockRecorder) Delete(orderID, fulfillmentID, eventID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockFulfillmentEventRepository)(nil).Delete), orderID, fulfillmentID, eventID)
}

// List mocks base method
func (m *MockFulfillmentEventRepository) List(orderID, fulfillmentID int64) ([]shopify.FulfillmentEvent, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", orderID, fulfillmentID)
	ret0, _ := ret[0].([]shopify.FulfillmentEvent)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List
func (mr *MockFulfillmentEventRepositoryMockRecorder) List(orderID, fulfillmentID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockFulfillmentEventRepository)(nil).List), orderID, fulfillmentID)
}