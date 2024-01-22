// Code generated by MockGen. DO NOT EDIT.
// Source: ../shopify/articles.go

// Package mockshopify is a generated GoMock package.
package mockshopify

import (
	reflect "reflect"

	shopify "github.com/MOHC-LTD/shopify/v2"
	gomock "github.com/golang/mock/gomock"
)

// MockArticleRepository is a mock of ArticleRepository interface.
type MockArticleRepository struct {
	ctrl     *gomock.Controller
	recorder *MockArticleRepositoryMockRecorder
}

// MockArticleRepositoryMockRecorder is the mock recorder for MockArticleRepository.
type MockArticleRepositoryMockRecorder struct {
	mock *MockArticleRepository
}

// NewMockArticleRepository creates a new mock instance.
func NewMockArticleRepository(ctrl *gomock.Controller) *MockArticleRepository {
	mock := &MockArticleRepository{ctrl: ctrl}
	mock.recorder = &MockArticleRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockArticleRepository) EXPECT() *MockArticleRepositoryMockRecorder {
	return m.recorder
}

// Get mocks base method.
func (m *MockArticleRepository) Get(blogID, id int64) (shopify.Article, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", blogID, id)
	ret0, _ := ret[0].(shopify.Article)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockArticleRepositoryMockRecorder) Get(blogID, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockArticleRepository)(nil).Get), blogID, id)
}

// GetAll mocks base method.
func (m *MockArticleRepository) GetAll(blogID int64) (shopify.Articles, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAll", blogID)
	ret0, _ := ret[0].(shopify.Articles)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAll indicates an expected call of GetAll.
func (mr *MockArticleRepositoryMockRecorder) GetAll(blogID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAll", reflect.TypeOf((*MockArticleRepository)(nil).GetAll), blogID)
}