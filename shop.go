package mockshopify

import (
	"github.com/MOHC-LTD/shopify/v2"

	gomock "github.com/golang/mock/gomock"
)

// Shop is a gomock shopify shop
type Shop struct {
	CustomerRepository         *MockCustomerRepository
	CustomerAddressRepository  *MockCustomerAddressRepository
	OrderRepository            *MockOrderRepository
	FulfillmentRepository      *MockFulfillmentRepository
	FulfillmentEventRepository *MockFulfillmentEventRepository
	FulfillmentOrderRepository *MockFulfillmentOrderRepository
	VariantRepository          *MockVariantRepository
	ProductRepository          *MockProductRepository
	InventoryLevelRepository   *MockInventoryLevelRepository
	InventoryItemRepository    *MockInventoryItemRepository
	CollectionRepository       *MockCollectionRepository
	ProductImageRepository     *MockProductImageRepository
	MetafieldRepository        *MockMetafieldRepository
	BlogRepository             *MockBlogRepository
	ArticleRepository          *MockArticleRepository
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
		CustomerRepository:         NewMockCustomerRepository(ctrl),
		CustomerAddressRepository:  NewMockCustomerAddressRepository(ctrl),
		OrderRepository:            NewMockOrderRepository(ctrl),
		FulfillmentRepository:      NewMockFulfillmentRepository(ctrl),
		FulfillmentEventRepository: NewMockFulfillmentEventRepository(ctrl),
		FulfillmentOrderRepository: NewMockFulfillmentOrderRepository(ctrl),
		VariantRepository:          NewMockVariantRepository(ctrl),
		ProductRepository:          NewMockProductRepository(ctrl),
		InventoryLevelRepository:   NewMockInventoryLevelRepository(ctrl),
		InventoryItemRepository:    NewMockInventoryItemRepository(ctrl),
		CollectionRepository:       NewMockCollectionRepository(ctrl),
		ProductImageRepository:     NewMockProductImageRepository(ctrl),
		MetafieldRepository:        NewMockMetafieldRepository(ctrl),
		BlogRepository:             NewMockBlogRepository(ctrl),
		ArticleRepository:          NewMockArticleRepository(ctrl),
	}
}

// Customers returns a mock implementation of a shopify customer repository
func (shop Shop) Customers() shopify.CustomerRepository {
	return shop.CustomerRepository
}

// CustomerAddresses returns an HTTP implementation of a Shopify customer addresses repository
func (shop Shop) CustomerAddresses() shopify.CustomerAddressRepository {
	return shop.CustomerAddressRepository
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

// FulfillmentEvents returns a mock implementation of a shopify fulfillment event repository
func (shop Shop) FulfillmentOrders() shopify.FulfillmentOrderRepository {
	return shop.FulfillmentOrderRepository
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

// InventoryItems returns a mock implementation of a shopify inventory items repository
func (shop Shop) InventoryItems() shopify.InventoryItemRepository {
	return shop.InventoryItemRepository
}

// Collections returns a mock implementation of a shopify collection repository
func (shop Shop) Collections() shopify.CollectionRepository {
	return shop.CollectionRepository
}

// ProductImages returns a mock implementation of a shopify product image repository
func (shop Shop) ProductImages() shopify.ProductImageRepository {
	return shop.ProductImageRepository
}

// Metafields returns a mock implementation of a shopify product image repository
func (shop Shop) Metafields() shopify.MetafieldRepository {
	return shop.MetafieldRepository
}

// Blogs returns a mock implementation of a shopify blog repository
func (shop Shop) Blogs() shopify.BlogRepository {
	return shop.BlogRepository
}

// Articles returns a mock implementation of a shopify article repository
func (shop Shop) Articles() shopify.ArticleRepository {
	return shop.ArticleRepository
}
