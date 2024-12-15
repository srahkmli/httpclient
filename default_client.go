package httpclient

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

// PostRequest sends a POST request with JSON body and custom headers.
func (c *defaultHttpClient) PostRequest(ctx context.Context, url string, body any, headers map[string]string) ([]byte, error) {
	// Marshal body to JSON
	byterequest, err := json.Marshal(body)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request body: %w", err)
	}

	// Create request with context
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, url, bytes.NewReader(byterequest))
	if err != nil {
		return nil, fmt.Errorf("failed to create POST request: %w", err)
	}

	// Set headers
	for key, value := range headers {
		req.Header.Set(key, value)
	}

	// Perform the request with retries
	var resp *http.Response
	for i := 0; i <= c.retries; i++ {
		if c.enableLogging {
			log.Printf("Attempt %d: Sending POST request to %s", i+1, url)
		}
		resp, err = c.client.Do(req)
		if err != nil {
			if c.enableLogging {
				log.Printf("Attempt %d failed: %v", i+1, err)
			}
			if i < c.retries {
				time.Sleep(c.retryDelay)
				continue
			}
			return nil, fmt.Errorf("POST request failed after %d retries: %w", c.retries, err)
		}
		defer resp.Body.Close()

		// Check for successful status code
		if resp.StatusCode >= 200 && resp.StatusCode < 300 {
			break
		}

		if c.enableLogging {
			log.Printf("Attempt %d failed with status code: %d", i+1, resp.StatusCode)
		}
		if i < c.retries {
			time.Sleep(c.retryDelay)
		}
	}

	// Read the response body
	bodyResp, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	return bodyResp, nil
}

// GetRequest sends a GET request with custom headers.
func (c *defaultHttpClient) GetRequest(ctx context.Context, url string, headers map[string]string) ([]byte, error) {
	// Create request with context
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create GET request: %w", err)
	}

	// Set headers
	for key, value := range headers {
		req.Header.Set(key, value)
	}

	// Perform the request with retries
	var resp *http.Response
	for i := 0; i <= c.retries; i++ {
		if c.enableLogging {
			log.Printf("Attempt %d: Sending GET request to %s", i+1, url)
		}
		resp, err = c.client.Do(req)
		if err != nil {
			if c.enableLogging {
				log.Printf("Attempt %d failed: %v", i+1, err)
			}
			if i < c.retries {
				time.Sleep(c.retryDelay)
				continue
			}
			return nil, fmt.Errorf("GET request failed after %d retries: %w", c.retries, err)
		}
		defer resp.Body.Close()

		// Check for successful status code
		if resp.StatusCode >= 200 && resp.StatusCode < 300 {
			break
		}

		if c.enableLogging {
			log.Printf("Attempt %d failed with status code: %d", i+1, resp.StatusCode)
		}
		if i < c.retries {
			time.Sleep(c.retryDelay)
		}
	}

	// Read the response body
	bodyResp, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	return bodyResp, nil
}
