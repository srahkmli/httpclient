package httpclient

import (
	"context"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

// Post sends a POST request with the specified body and headers and decodes the response into the provided type T.
func Post[T any](ctx context.Context, client httpClient, url string, req any, headers map[string]string) (T, error) {
	var resStruct T

	resp, err := client.PostRequest(ctx, url, req, headers)
	if err != nil {
		return resStruct, fmt.Errorf("http POST request failed: %w", err)
	}

	if err = json.Unmarshal(resp, &resStruct); err != nil {
		return resStruct, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	return resStruct, nil
}

// Get sends a GET request with the specified headers and decodes the response into the provided type T.
func Get[T any](ctx context.Context, client httpClient, url string, headers map[string]string) (T, error) {
	var resStruct T

	resp, err := client.GetRequest(ctx, url, headers)
	if err != nil {
		return resStruct, fmt.Errorf("http GET request failed: %w", err)
	}

	if err = json.Unmarshal(resp, &resStruct); err != nil {
		return resStruct, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	return resStruct, nil
}

// WithTimeout sets a custom timeout for the HTTP client.
func WithTimeout(timeout time.Duration) Option {
	return func(c *httpClient) {
		c.client.Timeout = timeout
	}
}

// WithLogging enables or disables request/response logging.
func WithLogging(enable bool) Option {
	return func(c *httpClient) {
		c.enableLogging = enable
	}
}

// WithCustomHeader adds a custom header to the HTTP request.
func WithCustomHeader(key, value string) Option {
	return func(c *httpClient) {
		c.client.Transport = http.DefaultTransport
	}
}

// WithRetries enables retries for failed requests and specifies retry count and delay.
func WithRetries(retries int, delay time.Duration) Option {
	return func(c *httpClient) {
		c.retries = retries
		c.retryDelay = delay
	}
}

// New creates a new HTTP client with the provided options.
func New(options ...Option) *httpClient {
	client := &httpClient{
		client: &http.Client{
			Timeout: 30 * time.Second, // Default timeout
		},
	}
	for _, opt := range options {
		opt(client)
	}
	return client
}

// GetJSON is a helper function to simplify GET requests and decode JSON responses.
func (c *httpClient) GetJSON(ctx context.Context, url string, headers map[string]string, result any) error {
	respBody, err := c.GetRequest(ctx, url, headers)
	if err != nil {
		return fmt.Errorf("failed to make GET request: %w", err)
	}

	// Decode JSON response into the provided result
	if err := json.Unmarshal(respBody, result); err != nil {
		return fmt.Errorf("failed to decode JSON response: %w", err)
	}
	return nil
}

// GetWithResponseTime sends a GET request and returns response time along with the data.
func (c *httpClient) GetWithResponseTime(ctx context.Context, url string, headers map[string]string) ([]byte, time.Duration, error) {
	start := time.Now()
	respBody, err := c.GetRequest(ctx, url, headers)
	if err != nil {
		return nil, 0, fmt.Errorf("failed to make GET request: %w", err)
	}
	duration := time.Since(start)
	return respBody, duration, nil
}

// WithUserAgent sets a custom User-Agent header globally.
func WithUserAgent(userAgent string) Option {
	return func(c *httpClient) {
		c.userAgent = userAgent
	}
}

// WithTransport allows setting a custom HTTP transport for fine-grained control.
func WithTransport(transport http.RoundTripper) Option {
	return func(c *httpClient) {
		c.transport = transport
	}
}

// WithProxy sets the proxy URL for the HTTP client.
func WithProxy(proxyURL string) Option {
	return func(c *httpClient) {
		c.proxyURL = proxyURL
	}
}

// WithTLSConfig allows configuring custom TLS settings, e.g., disabling SSL verification.
func WithTLSConfig(tlsConfig *tls.Config) Option {
	return func(c *httpClient) {
		c.tlsConfig = tlsConfig
	}
}

// WithBodyLogging enables logging of request and response bodies (use with caution).
func WithBodyLogging(enable bool) Option {
	return func(c *httpClient) {
		c.enableBodyLogging = enable
	}
}
