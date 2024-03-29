package order

import (
	"encoding/json"
	"fmt"
	"github.com/zerodays/woocommerce-go"
	"github.com/zerodays/woocommerce-go/internal/backend"
	"net/http"
	"strconv"
)

const (
	pathList = "/orders"
	pathEdit = "/orders/%d"
)

// Client is the API client used for working with orders.
// It should not be initialized directly. Use client.API instead.
type Client struct {
	backend *backend.Backend
}

// New creates a new client for orders.
// It should not be called directly.
// Instead, client.API should be used.
func New(backend *backend.Backend) *Client {
	return &Client{
		backend: backend,
	}
}

// List returns a list of orders with given parameters and total order count.
func (c Client) List(parameters woocommerce.Parameters) ([]*woocommerce.Order, int, error) {
	// Execute authenticated request.
	resp, err := c.backend.AuthenticatedRequest(backend.APITypeRest, http.MethodGet, pathList, nil, parameters, nil)
	if err != nil {
		return nil, 0, err
	}
	defer resp.Body.Close()

	// Unmarshal JSON.
	var orders []*woocommerce.Order
	err = json.NewDecoder(resp.Body).Decode(&orders)
	if err != nil {
		return nil, 0, fmt.Errorf("[woocommerce-go]: could not unmarshal orders json: %w", err)
	}

	// Get total order count
	countStr := resp.Header.Get(backend.TotalCountHeader)
	var count int
	if countStr != "" {
		count, err = strconv.Atoi(countStr)
		if err != nil {
			return nil, 0, fmt.Errorf("[woocommerce-go]: could not parse total order count: %w", err)
		}
	}

	return orders, count, nil
}

// Create creates a new order.
func (c Client) Create(orderCreate *woocommerce.OrderCreate) (*woocommerce.Order, error) {
	// Execute authenticated request.
	resp, err := c.backend.AuthenticatedRequest(backend.APITypeRest, http.MethodPost, pathList, orderCreate, nil, nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Unmarshal JSON.
	order := &woocommerce.Order{}
	err = json.NewDecoder(resp.Body).Decode(&order)
	if err != nil {
		return nil, fmt.Errorf("[woocommerce-go]: could not unmarshal order json: %w", err)
	}

	return order, nil
}

// Update updates the order with a given ID
func (c Client) Update(orderID int, orderUpdate woocommerce.OrderUpdate) (*woocommerce.Order, error) {
	// Execute authenticated request.
	path := fmt.Sprintf(pathEdit, orderID)
	resp, err := c.backend.AuthenticatedRequest(backend.APITypeRest, http.MethodPut, path, orderUpdate, nil, nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Unmarshal JSON.
	order := &woocommerce.Order{}
	err = json.NewDecoder(resp.Body).Decode(&order)
	if err != nil {
		return nil, fmt.Errorf("[woocommerce-go]: could not unmarshal order json: %w", err)
	}

	return order, nil
}
