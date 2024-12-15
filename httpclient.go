package httpclient

import (
	"context"
	"crypto/tls"
	"net/http"
	"time"
)

// IHttpClient defines an interface for making HTTP requests.
type IHttpClient interface {
	PostRequest(ctx context.Context, url string, body any, headers map[string]string) ([]byte, error)
	GetRequest(ctx context.Context, url string, headers map[string]string) ([]byte, error)
}

// Option is a function that configures the HTTP client.
type Option func(*httpClient)

type httpClient struct {
	client            *http.Client
	retries           int
	retryDelay        time.Duration
	enableLogging     bool
	userAgent         string
	transport         http.RoundTripper
	enableBodyLogging bool
	proxyURL          string
	tlsConfig         *tls.Config
}
