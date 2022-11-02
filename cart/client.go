package cart

import (
	"encoding/json"
	"fmt"
	"github.com/zerodays/woocommerce-go"
	"github.com/zerodays/woocommerce-go/internal/backend"
	"net/http"
)

const (
	pathCart               = "/cart"
	pathAddItem            = "/cart/add-item"
	pathRemoveItem         = "/cart/remove-item"
	pathUpdateItem         = "/cart/update-item"
	pathUpdateCustomer     = "/cart/update-customer"
	pathSelectShippingRate = "/cart/select-shipping-rate"
)

const (
	headerNonce     = "nonce"
	headerCartToken = "cart-token"
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
		headerCartToken: cartToken,
	}
	resp, err := c.backend.AuthenticatedRequest(backend.APITypeBlocks, http.MethodGet, pathCart, nil, nil, headers)
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

// New creates a new cart. It returns the cart token, an error that might have occurred.
func (c Client) New() (token string, err error) {
	resp, err := c.backend.AuthenticatedRequest(backend.APITypeBlocks, http.MethodGet, pathCart, nil, nil, nil)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	token = resp.Header.Get(headerCartToken)
	return token, nil
}

// getNonce gets the nonce for the cart with given cart token.
func (c Client) getNonce(cartToken string) (string, error) {
	// Execute request
	headers := map[string]string{
		headerCartToken: cartToken,
	}
	resp, err := c.backend.AuthenticatedRequest(backend.APITypeBlocks, http.MethodGet, pathCart, nil, nil, headers)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	return resp.Header.Get(headerNonce), nil
}

// AddItem adds an item to the cart with given cart token.
func (c Client) AddItem(cartToken string, itemID, quantity int, variations []woocommerce.CartItemVariation) (*woocommerce.Cart, error) {
	// Get nonce for the cart.
	nonce, err := c.getNonce(cartToken)
	if err != nil {
		return nil, fmt.Errorf("[woocommerce-go] could not get nonce: %w", err)
	}

	// Construct request body
	type addItemRequest struct {
		ID         int                             `json:"id"`
		Quantity   int                             `json:"quantity"`
		Variations []woocommerce.CartItemVariation `json:"variation,omitempty"`
	}
	req := addItemRequest{
		ID:         itemID,
		Quantity:   quantity,
		Variations: variations,
	}

	// Construct headers
	headers := map[string]string{
		headerNonce:     nonce,
		headerCartToken: cartToken,
	}

	// Execute request.
	resp, err := c.backend.AuthenticatedRequest(backend.APITypeBlocks, http.MethodPost, pathAddItem, req, nil, headers)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Unmarshal JSON response.
	cart := &woocommerce.Cart{}
	err = json.NewDecoder(resp.Body).Decode(&cart)
	if err != nil {
		return nil, fmt.Errorf("[woocommerce-go] could not unmarshal cart json: %w", err)
	}
	return cart, nil
}

// RemoveItem removes an item from the cart.
func (c Client) RemoveItem(cartToken string, itemKey string) (*woocommerce.Cart, error) {
	// Get nonce for the cart.
	nonce, err := c.getNonce(cartToken)
	if err != nil {
		return nil, fmt.Errorf("[woocommerce-go] could not get nonce: %w", err)
	}

	// Construct request body
	type removeItemRequest struct {
		Key string `json:"key"`
	}
	req := removeItemRequest{
		Key: itemKey,
	}

	// Construct headers
	headers := map[string]string{
		headerNonce:     nonce,
		headerCartToken: cartToken,
	}

	// Execute request.
	resp, err := c.backend.AuthenticatedRequest(backend.APITypeBlocks, http.MethodPost, pathRemoveItem, req, nil, headers)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Unmarshal JSON response.
	cart := &woocommerce.Cart{}
	err = json.NewDecoder(resp.Body).Decode(&cart)
	if err != nil {
		return nil, fmt.Errorf("[woocommerce-go] could not unmarshal cart json: %w", err)
	}
	return cart, nil
}

// UpdateItem updates the quantity of an item in the cart.
func (c Client) UpdateItem(cartToken, itemKey string, quantity int) (*woocommerce.Cart, error) {
	// Get nonce for the cart.
	nonce, err := c.getNonce(cartToken)
	if err != nil {
		return nil, fmt.Errorf("[woocommerce-go] could not get nonce: %w", err)
	}

	// Construct request body
	type updateItemRequest struct {
		Key      string `json:"key"`
		Quantity int    `json:"quantity"`
	}
	req := updateItemRequest{
		Key:      itemKey,
		Quantity: quantity,
	}

	// Construct headers
	headers := map[string]string{
		headerNonce:     nonce,
		headerCartToken: cartToken,
	}

	// Execute request.
	resp, err := c.backend.AuthenticatedRequest(backend.APITypeBlocks, http.MethodPost, pathUpdateItem, req, nil, headers)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Unmarshal JSON response.
	cart := &woocommerce.Cart{}
	err = json.NewDecoder(resp.Body).Decode(&cart)
	if err != nil {
		return nil, fmt.Errorf("[woocommerce-go] could not unmarshal cart json: %w", err)
	}
	return cart, nil
}

// UpdateCustomer updates the customer shipping and billing address.
func (c Client) UpdateCustomer(cartToken string, billingAddress, shippingAddress *woocommerce.CartAddress) (*woocommerce.Cart, error) {
	// Get nonce for the cart.
	nonce, err := c.getNonce(cartToken)
	if err != nil {
		return nil, fmt.Errorf("[woocommerce-go] could not get nonce: %w", err)
	}

	// Construct request body
	type updateCustomerRequest struct {
		BillingAddress  *woocommerce.CartAddress `json:"billing_address,omitempty"`
		ShippingAddress *woocommerce.CartAddress `json:"shipping_address,omitempty"`
	}
	req := updateCustomerRequest{
		BillingAddress:  billingAddress,
		ShippingAddress: shippingAddress,
	}

	// Construct headers
	headers := map[string]string{
		headerNonce:     nonce,
		headerCartToken: cartToken,
	}

	// Execute request.
	resp, err := c.backend.AuthenticatedRequest(backend.APITypeBlocks, http.MethodPost, pathUpdateCustomer, req, nil, headers)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Unmarshal JSON response.
	cart := &woocommerce.Cart{}
	err = json.NewDecoder(resp.Body).Decode(&cart)
	if err != nil {
		return nil, fmt.Errorf("[woocommerce-go] could not unmarshal cart json: %w", err)
	}
	return cart, nil
}

// SelectShippingRate selects a shipping rate for the cart.
func (c Client) SelectShippingRate(cartToken string, packageID int, rateID string) (*woocommerce.Cart, error) {
	// Get nonce for the cart.
	nonce, err := c.getNonce(cartToken)
	if err != nil {
		return nil, fmt.Errorf("[woocommerce-go] could not get nonce: %w", err)
	}

	// Construct request body
	type selectShippingRateRequest struct {
		PackageID int    `json:"package_id"`
		RateID    string `json:"rate_id"`
	}
	req := selectShippingRateRequest{
		PackageID: packageID,
		RateID:    rateID,
	}

	// Construct headers
	headers := map[string]string{
		headerNonce:     nonce,
		headerCartToken: cartToken,
	}

	// Execute request.
	resp, err := c.backend.AuthenticatedRequest(backend.APITypeBlocks, http.MethodPost, pathSelectShippingRate, req, nil, headers)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Unmarshal JSON response.
	cart := &woocommerce.Cart{}
	err = json.NewDecoder(resp.Body).Decode(&cart)
	if err != nil {
		return nil, fmt.Errorf("[woocommerce-go] could not unmarshal cart json: %w", err)
	}
	return cart, nil
}
