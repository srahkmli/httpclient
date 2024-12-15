package httpclient

import (
	"context"
	"net/http"
	"time"
)

// HttpClient defines an interface for making HTTP requests.
type HttpClient interface {
	PostRequest(ctx context.Context, url string, body any, headers map[string]string) ([]byte, error)
	GetRequest(ctx context.Context, url string, headers map[string]string) ([]byte, error)
}

// Option is a function that configures the HTTP client.
type Option func(*defaultHttpClient)

type defaultHttpClient struct {
	client        *http.Client
	retries       int
	retryDelay    time.Duration
	enableLogging bool
}
