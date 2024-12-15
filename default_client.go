package httpclient

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

var DefaultClient = &defaultHttpClient{
	client: &http.Client{},
}

type defaultHttpClient struct {
	client *http.Client
}

func (c *defaultHttpClient) PostRequest(ctx context.Context, url string, body any, headers map[string]string) ([]byte, error) {
	byterequest, err := json.Marshal(body)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request body: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, url, bytes.NewReader(byterequest))
	if err != nil {
		return nil, fmt.Errorf("failed to create POST request: %w", err)
	}

	for key, value := range headers {
		req.Header.Set(key, value)
	}

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("POST request failed: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return nil, fmt.Errorf("POST request returned status code %d", resp.StatusCode)
	}

	return io.ReadAll(resp.Body)
}

func (c *defaultHttpClient) GetRequest(ctx context.Context, url string, headers map[string]string) ([]byte, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create GET request: %w", err)
	}

	for key, value := range headers {
		req.Header.Set(key, value)
	}

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("GET request failed: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return nil, fmt.Errorf("GET request returned status code %d", resp.StatusCode)
	}

	return io.ReadAll(resp.Body)
}
