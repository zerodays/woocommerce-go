package client

import (
	"github.com/zerodays/woocommerce-go/cart"
	"github.com/zerodays/woocommerce-go/customer"
	"github.com/zerodays/woocommerce-go/internal/backend"
	"github.com/zerodays/woocommerce-go/order"
	"github.com/zerodays/woocommerce-go/product"
	"github.com/zerodays/woocommerce-go/tax"
)

// API is the API client. It should be created with the New function.
//
// Since woocommerce supports extensions that add additional fields to types,
// API is defined with generic types. This allows you to create your own types
// that extend the default implementation in this library.
// Parameters represent the following types:
// - C: Customer, default implementation is woocommerce.Customer
// - P: Product, default implementation is woocommerce.Product
// - PV: ProductVariation, default implementation is woocommerce.ProductVariation
type API[C, P, PV any] struct {
	Order    *order.Client
	Cart     *cart.Client
	Tax      *tax.Client
	Customer *customer.Client[C]
	Product  *product.Client[P, PV]
}

// Init initializes the API client with given credentials.
// ConsumerKey and consumerSecret are gotten from woocommerce admin console.
// BaseURL is the base URL of the store. For instance if the index URL of the woocommerce API is
// https://example.com/wp-json/wc/v3, then the base URL is https://example.com
func (a *API[C, P, PV]) Init(baseURL, consumerKey, consumerSecret string) {
	b := backend.New(baseURL, consumerKey, consumerSecret)

	a.Order = order.New(b)
	a.Cart = cart.New(b)
	a.Tax = tax.New(b)
	a.Customer = customer.New[C](b)
	a.Product = product.New[P, PV](b)
}

// New creates a new API client with given credentials.
// ConsumerKey and consumerSecret are gotten from woocommerce admin console.
// BaseURL is the base URL of the store. For instance if the index URL of the woocommerce API is
// https://example.com/wp-json/wc/v3, then the base URL is https://example.com
//
// Generic parameters are documented in the definition of the API type.
func New[C, P, PV any](baseURL, consumerKey, consumerSecret string) *API[C, P, PV] {
	api := &API[C, P, PV]{}
	api.Init(baseURL, consumerKey, consumerSecret)
	return api
}
