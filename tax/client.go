package tax

import (
	"encoding/json"
	"fmt"
	"github.com/zerodays/woocommerce-go"
	"github.com/zerodays/woocommerce-go/internal/backend"
	"net/http"
)

const pathList = "/taxes"

// Client is the API client used for working with taxes.
// It should not be initialized directly. Use client.API instead.
type Client struct {
	backend *backend.Backend
}

// New creates a new client for taxes.
// It should not be called directly.
// Instead, client.API should be used.
func New(backend *backend.Backend) *Client {
	return &Client{
		backend: backend,
	}
}

// List lists taxes with given parameters.
func (c Client) List(parameters woocommerce.Parameters) ([]*woocommerce.Tax, error) {
	// Execute authenticated request.
	resp, err := c.backend.AuthenticatedRequest(backend.APITypeRest, http.MethodGet, pathList, nil, parameters, nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Unmarshal JSON.
	var taxes []*woocommerce.Tax
	err = json.NewDecoder(resp.Body).Decode(&taxes)
	if err != nil {
		return nil, fmt.Errorf("[woocommerce-go]: could not unmarshal taxes json: %w", err)
	}

	return taxes, nil
}
