package httpclient

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	url2 "net/url"
	"time"
)

// PostRequest sends a POST request with JSON body and custom headers.
func (c *httpClient) PostRequest(ctx context.Context, url string, body any, headers map[string]string) ([]byte, error) {
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

	// If user-agent is set, add it to headers
	if c.userAgent != "" {
		req.Header.Set("User-Agent", c.userAgent)
	}

	// If a proxy is configured, configure the transport
	if c.proxyURL != "" {
		proxyURL, err := url2.Parse(c.proxyURL)
		if err != nil {
			return nil, fmt.Errorf("invalid proxy URL: %w", err)
		}
		c.client.Transport = &http.Transport{
			Proxy: http.ProxyURL(proxyURL),
		}
	}

	// If TLS config is set, set the custom TLS configuration
	if c.tlsConfig != nil {
		c.client.Transport = &http.Transport{
			TLSClientConfig: c.tlsConfig,
		}
	}

	// Perform the request with retries
	var resp *http.Response
	for i := 0; i <= c.retries; i++ {
		if c.enableLogging {
			log.Printf("Attempt %d: Sending POST request to %s", i+1, url)
			if c.enableBodyLogging {
				log.Printf("Request body: %s", string(byterequest))
			}
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

	// Log response body if logging is enabled
	if c.enableLogging && c.enableBodyLogging {
		bodyResp, _ := io.ReadAll(resp.Body)
		log.Printf("Response body: %s", string(bodyResp))
	}

	// Read the response body
	bodyResp, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	return bodyResp, nil
}

// GetRequest sends a GET request with custom headers.
func (c *httpClient) GetRequest(ctx context.Context, url string, headers map[string]string) ([]byte, error) {
	// Create request with context
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create GET request: %w", err)
	}

	// Set headers
	for key, value := range headers {
		req.Header.Set(key, value)
	}

	// If user-agent is set, add it to headers
	if c.userAgent != "" {
		req.Header.Set("User-Agent", c.userAgent)
	}

	// If proxy is configured, configure the transport
	if c.proxyURL != "" {
		proxyURL, err := url2.Parse(c.proxyURL)
		if err != nil {
			return nil, fmt.Errorf("invalid proxy URL: %w", err)
		}
		c.client.Transport = &http.Transport{
			Proxy: http.ProxyURL(proxyURL),
		}
	}

	// If TLS config is set, set the custom TLS configuration
	if c.tlsConfig != nil {
		c.client.Transport = &http.Transport{
			TLSClientConfig: c.tlsConfig,
		}
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

	// Log response body if logging is enabled
	if c.enableLogging && c.enableBodyLogging {
		bodyResp, _ := io.ReadAll(resp.Body)
		log.Printf("Response body: %s", string(bodyResp))
	}

	// Read the response body
	bodyResp, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	return bodyResp, nil
}
