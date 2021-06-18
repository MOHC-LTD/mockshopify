package mockshopify

import (
	"github.com/MOHC-LTD/shopify"

	gomock "github.com/golang/mock/gomock"
)

// Shop is a gomock shopify shop
type Shop struct {
	OrderRepository            *MockOrderRepository
	FulfillmentRepository      *MockFulfillmentRepository
	FulfillmentEventRepository *MockFulfillmentEventRepository
	VariantRepository          *MockVariantRepository
	ProductRepository          *MockProductRepository
	InventoryLevelRepository   *MockInventoryLevelRepository
}

// NewShop builds a gomock implementation of a shopify shop.
/*
	This shop can be used in the same way as any gomock mock.
	Example:
	mockShop := mockshopify.NewShop(mockCtrl)
	mockShop.OrderRepository.Expect().Get(orderID).Return(order, nil)
	See https://github.com/golang/mock
*/
func NewShop(ctrl *gomock.Controller) Shop {
	return Shop{
		OrderRepository:            NewMockOrderRepository(ctrl),
		FulfillmentRepository:      NewMockFulfillmentRepository(ctrl),
		FulfillmentEventRepository: NewMockFulfillmentEventRepository(ctrl),
		VariantRepository:          NewMockVariantRepository(ctrl),
		InventoryLevelRepository:   NewMockInventoryLevelRepository(ctrl),
	}
}

// Orders returns a mock implementation of a shopify order repository
func (shop Shop) Orders() shopify.OrderRepository {
	return shop.OrderRepository
}

// Fulfillments returns a mock implementation of a shopify fulfillment repository
func (shop Shop) Fulfillments() shopify.FulfillmentRepository {
	return shop.FulfillmentRepository
}

// FulfillmentEvents returns a mock implementation of a shopify fulfillment event repository
func (shop Shop) FulfillmentEvents() shopify.FulfillmentEventRepository {
	return shop.FulfillmentEventRepository
}

// Variants returns a mock implementation of a shopify variant repository
func (shop Shop) Variants() shopify.VariantRepository {
	return shop.VariantRepository
}

// Products returns a mock implementation of a shopify product repository
func (shop Shop) Products() shopify.ProductRepository {
	return shop.ProductRepository
}

// InventoryLevels returns a mock implementation of a shopify inventory levels repository
func (shop Shop) InventoryLevels() shopify.InventoryLevelRepository {
	return shop.InventoryLevelRepository
}
