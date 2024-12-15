package httpclient

import (
	"context"
	"fmt"
	httpclient "github.com/srahkmli/httpclient"
	"log"
	"time"
)

func main() {
	// Create a new HTTP client with custom options
	client := httpclient.New(
		httpclient.WithTimeout(5*time.Second),
		httpclient.WithRetries(3, 2*time.Second),
		httpclient.WithLogging(true),
		httpclient.WithUserAgent("MyCustomUserAgent/1.0"),
		httpclient.WithBodyLogging(true),
		httpclient.WithProxy("http://proxy.example.com"),
	)

	// Make a GET request
	response, err := client.GetRequest(context.Background(), "https://api.example.com/data", nil)
	if err != nil {
		log.Fatalf("Error making GET request: %v", err)
	}

	// Print the response
	fmt.Println("Response:", string(response))
}
