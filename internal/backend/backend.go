package backend

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/zerodays/woocommerce-go"
)

const (
	timeoutDuration     = 1 * time.Minute
	urlPathPrefixRest   = "/wp-json/wc/v3"
	urlPathPrefixBlocks = "/wp-json/wc/store/v1"
)

// APIType is the type of the API to call. Woocommerce has two APIs:
// - [REST API](https://woocommerce.github.io/woocommerce-rest-api-docs/)
// - [Blocks API](https://github.com/woocommerce/woocommerce-blocks/tree/trunk/src/StoreApi)
type APIType string

const (
	APITypeRest   APIType = "rest"
	APITypeBlocks APIType = "blocks"
)

type ErrInvalidStatusCode struct {
	StatusCode int
	Body       string
}

func (e ErrInvalidStatusCode) Error() string {
	// Very explicit error to make debugging easier.
	return fmt.Sprintf("[woocommerce-go]: got invalid status code %d. Body: %s", e.StatusCode, e.Body)
}

// Backend holds the backend that handles
// execution of authenticate requests.
// It should be initialized with the New method.
type Backend struct {
	baseURL             string
	basicAuthentication string
}

// New creates a new Backend with passed user credentials.
// ConsumerKey and consumerSecret are gotten from woocommerce admin console.
// BaseURL is the base URL of the store. For instance if the index URL of the woocommerce API is
// https://example.com/wp-json/wc/v3, then the base URL is https://example.com
func New(baseURL, consumerKey, consumerSecret string) *Backend {
	auth := consumerKey + ":" + consumerSecret
	auth = base64.StdEncoding.EncodeToString([]byte(auth))

	return &Backend{
		baseURL:             baseURL,
		basicAuthentication: auth,
	}
}

type filterReader struct {
	io.ReadCloser
}

func (f *filterReader) Read(p []byte) (int, error) {
	n, err := f.ReadCloser.Read(p)
	if err != nil {
		return n, err
	}

	// Filter out null bytes
	j := 0
	for i := 0; i < n; i++ {
		if p[i] != 0 {
			p[j] = p[i]
			j++
		}
	}

	if j == 0 {
		return 0, io.EOF
	}

	return j, nil
}

// AuthenticatedRequest executes an authenticated request to the woocommerce server.
// Body can be se to nil to execute a request with an empty body.
// Parameters can be set to nil to execute a request without GET parameters.
// Responses with status code not in range of [200, 300) are being treated as an error.
// The function returns http response and errors that might have occurred during the request execution.
// If the error is nil, caller is responsible for closing the response body.
func (b *Backend) AuthenticatedRequest(apiType APIType, method, path string, body interface{}, parameters woocommerce.Parameters, headers map[string]string) (*http.Response, error) {
	client := &http.Client{
		Timeout: timeoutDuration,
	}

	// Parse the given body if it is not nil.
	var bodyReader io.Reader = nil
	if body != nil {
		bodyBytes, err := json.Marshal(body)
		if err != nil {
			return nil, fmt.Errorf("[woocommerce-go]: could not marshal body to JSON: %w", err)
		}

		bodyReader = bytes.NewReader(bodyBytes)
	}

	// Build the URL.
	reqURL := b.baseURL
	switch apiType {
	case APITypeRest:
		reqURL += urlPathPrefixRest
	case APITypeBlocks:
		reqURL += urlPathPrefixBlocks
	default:
		return nil, fmt.Errorf("[woocommerce-go]: invalid API type: %s", apiType)
	}
	reqURL += path
	if parameters != nil {
		reqURL += "?" + parameters.Values().Encode()
	}

	// Create a new request and set its headers
	req, err := http.NewRequest(method, reqURL, bodyReader)
	if err != nil {
		return nil, fmt.Errorf("[woocommerce-go]: could not create a new request: %w", err)
	}

	// Remove User-Agent header, because go's default one is blocked by neoserve.
	req.Header.Set("User-Agent", "")

	if apiType == APITypeRest {
		req.Header.Set("Authorization", "Basic "+b.basicAuthentication)
	}
	req.Header.Set("Content-Type", "application/json")

	// Set custom headers
	for key, value := range headers {
		req.Header.Set(key, value)
	}

	// Execute the request
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("[woocommerce-go]: could not execute the request: %w", err)
	}

	// Check valid response code range.
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		if resp.StatusCode >= 400 {
			var err = &woocommerce.Error{}
			_ = json.NewDecoder(resp.Body).Decode(&err)
			_ = resp.Body.Close()

			err.StatusCode = resp.StatusCode
			return resp, err
		} else {
			data, _ := io.ReadAll(resp.Body)
			_ = resp.Body.Close()

			return resp, ErrInvalidStatusCode{
				StatusCode: resp.StatusCode,
				Body:       string(data),
			}
		}
	}

	resp.Body = &filterReader{resp.Body}

	return resp, nil
}
