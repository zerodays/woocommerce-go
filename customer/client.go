package customer

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/zerodays/woocommerce-go"
	"github.com/zerodays/woocommerce-go/internal/backend"
)

const (
	pathList     = "/customers"
	pathRetrieve = "/customers/%s"
)

// Client is the API client used for working with customers.
// It should not be initialized directly. Use client.API instead.
//
// Generic type C represents the type of the customer.
// This library provides default implementation of the Customer type as woocommerce.Customer.
// If you have extensions installed on woocommerce that add additional fields to the customer,
// you can create your own type that embeds the woocommerce.Customer type and add additional fields.
type Client[C any] struct {
	backend *backend.Backend
}

// New creates a new client for customers.
// It should not be called directly.
// Instead, client.API should be used.
func New[C any](backend *backend.Backend) *Client[C] {
	return &Client[C]{
		backend: backend,
	}
}

// List lists customers with given parameters.
func (c Client[C]) List(parameters woocommerce.Parameters) ([]C, error) {
	// Execute authenticated request.
	resp, err := c.backend.AuthenticatedRequest(backend.APITypeRest, http.MethodGet, pathList, nil, parameters, nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Unmarshal JSON.
	var customers []C
	err = json.NewDecoder(resp.Body).Decode(&customers)
	if err != nil {
		return nil, fmt.Errorf("[woocommerce-go]: could not unmarshal customers json: %w", err)
	}

	return customers, nil
}

// Retrieve retrieves a single customer by its ID.
func (c Client[C]) Retrieve(id string) (C, error) {
	var empty C

	// Execute authenticated request.
	path := fmt.Sprintf(pathRetrieve, id)
	resp, err := c.backend.AuthenticatedRequest(backend.APITypeRest, http.MethodGet, path, nil, nil, nil)
	if err != nil {
		return empty, err
	}
	defer resp.Body.Close()

	// Unmarshal JSON.
	var customer C
	err = json.NewDecoder(resp.Body).Decode(&customer)
	if err != nil {
		return empty, fmt.Errorf("[woocommerce-go]: could not unmarshal customer json: %w", err)
	}

	return customer, nil
}

// Update updates the given customer. ID of the customer must be set.
func (c Client[C]) Update(customer *C, id int) error {
	path := fmt.Sprintf(pathRetrieve, strconv.Itoa(id))
	resp, err := c.backend.AuthenticatedRequest(backend.APITypeRest, http.MethodPut, path, customer, nil, nil)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Unmarshal JSON.
	err = json.NewDecoder(resp.Body).Decode(customer)
	if err != nil {
		return fmt.Errorf("[woocommerce-go]: could not unmarshal customer json: %w", err)
	}

	return nil
}
