package client

import (
	"github.com/zerodays/woocommerce-go/cart"
	"github.com/zerodays/woocommerce-go/internal/backend"
	"github.com/zerodays/woocommerce-go/order"
)

// API is the API client. It should be created with the New function.
type API struct {
	Order *order.Client
	Cart  *cart.Client
}

// Init initializes the API client with given credentials.
// ConsumerKey and consumerSecret are gotten from woocommerce admin console.
// BaseURL is the base URL of the store. For instance if the index URL of the woocommerce API is
// https://example.com/wp-json/wc/v3, then the base URL is https://example.com
func (a *API) Init(baseURL, consumerKey, consumerSecret string) {
	b := backend.New(baseURL, consumerKey, consumerSecret)

	a.Order = order.New(b)
	a.Cart = cart.New(b)
}

// New creates a new API client with given credentials.
// ConsumerKey and consumerSecret are gotten from woocommerce admin console.
// BaseURL is the base URL of the store. For instance if the index URL of the woocommerce API is
// https://example.com/wp-json/wc/v3, then the base URL is https://example.com
func New(baseURL, consumerKey, consumerSecret string) *API {
	api := &API{}
	api.Init(baseURL, consumerKey, consumerSecret)
	return api
}
