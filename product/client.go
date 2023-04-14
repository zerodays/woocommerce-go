package product

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/zerodays/woocommerce-go"
	"github.com/zerodays/woocommerce-go/internal/backend"
)

const (
	pathList          = "/products"
	pathRetrieve      = "/products/%d"
	pathListVariation = "/products/%d/variations"
)

// Client is the API client used for working with products.
// It should not be initialized directly. Use client.API instead.
//
// Generic type P represents the type of the product and PV represents type of product variation.
// This library provides default implementation of the Product type as woocommerce.Product
// and ProductVariation type as woocommerce.ProductVariation.
// If you have extensions installed on woocommerce that add additional fields to the product,
// you can create your own type that embeds the woocommerce.Product or woocommerce.ProductVariation
//
//	type and add additional fields.
type Client[P, PV any] struct {
	backend *backend.Backend
}

// New creates a new client for products.
// It should not be called directly.
// Instead, client.API should be used.
func New[P, PV any](backend *backend.Backend) *Client[P, PV] {
	return &Client[P, PV]{
		backend: backend,
	}
}

// List lists products with given parameters.
func (c Client[P, PV]) List(parameters woocommerce.Parameters) ([]P, error) {
	// Execute authenticated request.
	resp, err := c.backend.AuthenticatedRequest(backend.APITypeRest, http.MethodGet, pathList, nil, parameters, nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Unmarshal JSON.
	var products []P
	err = json.NewDecoder(resp.Body).Decode(&products)
	if err != nil {
		return nil, fmt.Errorf("[woocommerce-go]: could not unmarshal products json: %w", err)
	}

	return products, nil
}

// ListVariations lists product variations for a given product.
func (c Client[P, PV]) ListVariations(productID int, parameters woocommerce.Parameters) ([]PV, error) {
	// Execute authenticated request.
	path := fmt.Sprintf(pathListVariation, productID)
	resp, err := c.backend.AuthenticatedRequest(backend.APITypeRest, http.MethodGet, path, nil, parameters, nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Unmarshal JSON.
	var variations []PV
	err = json.NewDecoder(resp.Body).Decode(&variations)
	if err != nil {
		return nil, fmt.Errorf("[woocommerce-go]: could not unmarshal product variations json: %w", err)
	}

	return variations, nil
}

// Retrieve retrieves a single product by its ID.
func (c Client[P, PV]) Retrieve(productID int) (P, error) {
	var product P

	// Execute authenticated request.
	path := fmt.Sprintf(pathRetrieve, productID)
	resp, err := c.backend.AuthenticatedRequest(backend.APITypeRest, http.MethodGet, path, nil, nil, nil)
	if err != nil {
		return product, err
	}
	defer resp.Body.Close()

	// Unmarshal JSON.
	err = json.NewDecoder(resp.Body).Decode(&product)
	if err != nil {
		return product, fmt.Errorf("[woocommerce-go]: could not unmarshal product json: %w", err)
	}

	return product, nil
}
