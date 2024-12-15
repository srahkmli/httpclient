package httpclient

import (
	"context"
	"errors"
	"testing"
)

type mockHttpClient struct{}

func (m *mockHttpClient) PostRequest(ctx context.Context, url string, body any, headers map[string]string) ([]byte, error) {
	if url == "https://error.example.com" {
		return nil, errors.New("mocked error")
	}
	return []byte(`{"mocked":"post-response"}`), nil
}

func (m *mockHttpClient) GetRequest(ctx context.Context, url string, headers map[string]string) ([]byte, error) {
	if url == "https://error.example.com" {
		return nil, errors.New("mocked error")
	}
	return []byte(`{"mocked":"get-response"}`), nil
}

func TestPost(t *testing.T) {
	client := &mockHttpClient{}
	ctx := context.Background()
	url := "https://api.example.com"
	headers := map[string]string{"Authorization": "Bearer token"}

	type Response struct {
		Mocked string `json:"mocked"`
	}

	resp, err := Post[Response](ctx, client, url, map[string]string{}, headers)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if resp.Mocked != "post-response" {
		t.Fatalf("unexpected response: %+v", resp)
	}
}

func TestGet(t *testing.T) {
	client := &mockHttpClient{}
	ctx := context.Background()
	url := "https://api.example.com"
	headers := map[string]string{"Authorization": "Bearer token"}

	type Response struct {
		Mocked string `json:"mocked"`
	}

	resp, err := Get[Response](ctx, client, url, headers)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if resp.Mocked != "get-response" {
		t.Fatalf("unexpected response: %+v", resp)
	}
}
