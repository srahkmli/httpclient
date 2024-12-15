package httpclient

import (
	"context"
	"encoding/json"
	"fmt"
)

// Post sends a POST request with the specified body and headers and decodes the response into the provided type T.
func Post[T any](ctx context.Context, client HttpClient, url string, req any, headers map[string]string) (T, error) {
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
func Get[T any](ctx context.Context, client HttpClient, url string, headers map[string]string) (T, error) {
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
