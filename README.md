# mockshopify

A [gomock](https://github.com/golang/mock) implementation of the [shopify](https://github.com/MOHC-LTD/shopify) package.

## Contents

- [Github authentication](#github-authentication)
- [Installation](#installation)
- [Usage](#usage)
- [Docker](#docker)
- [How to contribute](#how-to-contribute)

## Github authentication

As this is a private Github repository, you will need to set up private go modules.

First set the `GOPRIVATE` go environment variable.

```sh
go env -w GOPRIVATE=github.com/MOHC-LTD
```

Generate a [Github personal access token](https://github.com/settings/tokens), and set up
global Github authentication on your machine

```sh
git config --global url."https://${username}:${access_token}@github.com".insteadOf "https://github.com"
```

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

## Docker

To build applications that consuming this module using docker, you will need to allow the docker container to authenticate with Github.

Do this by adding the following lines to your Dockerfile.

```sh
ARG authToken

RUN go env -w GOPRIVATE=github.com/MOHC-LTD

RUN apk add git

RUN git config --global url."https://golang:$authToken@github.com".insteadOf "https://github.com"
```

Then, when building your container, set the docker argument `authToken` to the value of your Github access token.

## How to contribute

Something missing or not working as expected? See our [contribution guide](./CONTRIBUTING.md).
