# mockshopify

A [gomock](https://github.com/golang/mock) implementation of the [shopify](https://github.com/MOHC-LTD/shopify) package.

## Contents

- [Installation](#installation)
- [Usage](#usage)
- [How to contribute](#how-to-contribute)

## Installation

Install the module using

```sh
go get -u github.com/MOHC-LTD/mockshopify
```

## Usage

Each implementing repository on the shop is exported, this allows you to expect calls to the shop in your tests.

```go
func TestSomething(t *testing.T) {
    mockCtrl := gomock.NewController(t)
    defer mockCtrl.Finish()

    mockShop := mockshopify.NewShop(mockCtrl)

    service := NewService(mockShop)

    mockShop.OrderRepository.EXPECT().GET(10).Return(shopify.Order{}, nil)

    service.Get(10)
}
```

## How to contribute

Something missing or not working as expected? See our [contribution guide](./CONTRIBUTING.md).
