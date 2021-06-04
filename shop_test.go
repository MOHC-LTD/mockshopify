package mockshopify_test

import (
	"testing"

	"github.com/MOHC-LTD/mockshopify"
	"github.com/MOHC-LTD/shopify"
)

func Test_ShopImplementsShopify(t *testing.T) {
	var _ shopify.Shop = new(mockshopify.Shop)
}
