package customer

import (
	"encoding/json"
	"fmt"
	"github.com/zerodays/woocommerce-go"
	"github.com/zerodays/woocommerce-go/internal/backend"
	"net/http"
)

const (
	pathList     = "/customers"
	pathRetrieve = "/customers/%s"
)

// Client is the API client used for working with customers.
// It should not be initialized directly. Use client.API instead.
type Client struct {
	backend *backend.Backend
}

// New creates a new client for customers.
// It should not be called directly.
// Instead, client.API should be used.
func New(backend *backend.Backend) *Client {
	return &Client{
		backend: backend,
	}
}

// List lists customers with given parameters.
func (c Client) List(parameters woocommerce.Parameters) ([]*woocommerce.Customer, error) {
	// Execute authenticated request.
	resp, err := c.backend.AuthenticatedRequest(backend.APITypeRest, http.MethodGet, pathList, nil, parameters, nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Unmarshal JSON.
	var customers []*woocommerce.Customer
	err = json.NewDecoder(resp.Body).Decode(&customers)
	if err != nil {
		return nil, fmt.Errorf("[woocommerce-go]: could not unmarshal customers json: %w", err)
	}

	return customers, nil
}

// Retrieve retrieves a single customer by its ID.
func (c Client) Retrieve(id string) (*woocommerce.Customer, error) {
	// Execute authenticated request.
	path := fmt.Sprintf(pathRetrieve, id)
	resp, err := c.backend.AuthenticatedRequest(backend.APITypeRest, http.MethodGet, path, nil, nil, nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Unmarshal JSON.
	customer := &woocommerce.Customer{}
	err = json.NewDecoder(resp.Body).Decode(customer)
	if err != nil {
		return nil, fmt.Errorf("[woocommerce-go]: could not unmarshal customer json: %w", err)
	}

	return customer, nil
}
