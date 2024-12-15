package httpclient

import (
	"context"
)

// HttpClient defines an interface for making HTTP requests.
type HttpClient interface {
	PostRequest(ctx context.Context, url string, body any, headers map[string]string) ([]byte, error)
	GetRequest(ctx context.Context, url string, headers map[string]string) ([]byte, error)
}
