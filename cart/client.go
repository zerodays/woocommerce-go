package cart

import (
	"encoding/json"
	"fmt"
	"github.com/zerodays/woocommerce-go"
	"github.com/zerodays/woocommerce-go/internal/backend"
	"net/http"
)

// Client is the API client used for working with cart.
// It should not be initialized directly. Use client.API instead.
type Client struct {
	backend *backend.Backend
}

// New creates a new client for cart.
// It should not be called directly.
// Instead, client.API should be used.
func New(backend *backend.Backend) *Client {
	return &Client{
		backend: backend,
	}
}

// Get gets the cart with given cart token.
func (c Client) Get(cartToken string) (*woocommerce.Cart, error) {
	// Execute request
	headers := map[string]string{
		"cart-token": cartToken,
	}
	resp, err := c.backend.AuthenticatedRequest(backend.APITypeBlocks, http.MethodGet, "/cart", nil, nil, headers)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Unmarshal JSON.
	cart := &woocommerce.Cart{}
	err = json.NewDecoder(resp.Body).Decode(&cart)
	if err != nil {
		return nil, fmt.Errorf("[woocommerce-go] could not unmarshal cart json: %w", err)
	}

	return cart, nil
}
