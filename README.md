# HTTP Client Package

A simple and efficient Go package for making HTTP requests, handling responses, and managing common HTTP client utilities like retries, timeouts, and headers.

## Table of Contents

- [Installation](#installation)
- [Usage](#usage)
  - [Basic Example](#basic-example)
  - [Custom Configuration](#custom-configuration)
  - [Handling Timeouts](#handling-timeouts)
  - [Retries](#retries)
- [Features](#features)
- [License](#license)

## Installation

To install the `httpclient` package, use Go's package manager:

## Usage
This package simplifies working with HTTP requests and responses. Below are examples of how to use the package.

## Basic Example
Here’s how to use the httpclient package with a simple GET request:
 
```go

    package main

    import (
    "fmt"
    "log"
    "github.com/srahkmli/httpclient"
      )

    func main() {
    // Create a new HTTP client
    client := httpclient.New()

    // Make a GET request
    response, err := client.Get("https://api.example.com/data")
    if err != nil {
        log.Fatalf("Error making GET request: %v", err)
    }
    
    // Print the response
    fmt.Println("Response:", response)
    }
```

## Custom Configuration
You can customize the HTTP client with options like custom headers, timeouts, or proxies.
```go
    package main

    import (
    "fmt"
    "log"
    "github.com/srahkmli/httpclient"
    )

    func main() {
    // Create a new HTTP client with custom settings
    client := httpclient.New(
        httpclient.WithTimeout(10),                // Set timeout to 10 seconds
        httpclient.WithCustomHeader("User-Agent", "MyCustomClient/1.0"),
    )

    // Make a GET request
    response, err := client.Get("https://api.example.com/data")
    if err != nil {
        log.Fatalf("Error making GET request: %v", err)
    }
    
    // Print the response
    fmt.Println("Response:", response)
    }
```
## Handling Timeouts
To handle timeouts, you can configure the timeout value when initializing the client. The client will automatically cancel the request if it exceeds the specified timeout.
```go
    package main

    import (
    "fmt"
    "log"
    "github.com/srahkmli/httpclient"
    )

    func main() {
    // Create a new client with a 5-second timeout
    client := httpclient.New(httpclient.WithTimeout(5))

    // Make a GET request
    response, err := client.Get("https://api.example.com/data")
    if err != nil {
        log.Fatalf("Request timed out or error occurred: %v", err)
    }

    // Print the response
    fmt.Println("Response:", response)
    }

```
## Handling Timeouts
To handle timeouts, you can configure the timeout value when initializing the client. The client will automatically cancel the request if it exceeds the specified timeout.
```go

    package main

    import (
    "fmt"
    "log"
    "github.com/srahkmli/httpclient"
    )

    func main() {
    // Create a new client with a 5-second timeout
    client := httpclient.New(httpclient.WithTimeout(5))

    // Make a GET request
    response, err := client.Get("https://api.example.com/data")
    if err != nil {
        log.Fatalf("Request timed out or error occurred: %v", err)
    }

    // Print the response
    fmt.Println("Response:", response)
    }
```
## Retries
The httpclient package supports automatic retries for failed requests. You can configure the number of retries and the delay between retries.
```go
    package main

    import (
    "fmt"
    "log"
    "github.com/srahkmli/httpclient"
    )

    func main() {
    // Create a new client with retry configuration
    client := httpclient.New(
        httpclient.WithRetries(3),                // Retry up to 3 times
        httpclient.WithRetryDelay(2),             // 2-second delay between retries
    )

    // Make a GET request
    response, err := client.Get("https://api.example.com/data")
    if err != nil {
        log.Fatalf("Error making GET request: %v", err)
    }

    // Print the response
    fmt.Println("Response:", response)
    }
```

## Handling JSON Responses
This package also simplifies working with JSON responses. You can decode JSON directly into your structs:
```go
    package main

    import (
    "fmt"
    "log"
    "github.com/srahkmli/httpclient"
    )

    type ApiResponse struct {
    Name  string `json:"name"`
    Value string `json:"value"`
    }

    func main() {
    // Create a new HTTP client
    client := httpclient.New()

    // Make a GET request
    var data ApiResponse
    err := client.GetJSON("https://api.example.com/data", &data)
    if err != nil {
        log.Fatalf("Error making GET request: %v", err)
    }
    
    // Print the decoded JSON response
    fmt.Printf("Name: %s, Value: %s\n", data.Name, data.Value)
    }
```
## Features 💡 🧠
Customizable timeouts for requests.
Automatic retries on failed requests with configurable retry count and delay.
Custom headers to send with each request.
Flexible HTTP methods (GET, POST, PUT, DELETE).
Supports JSON encoding/decoding for APIs.
Error handling with structured responses.
Request/response logging (optional).


## Configuration Options
 * The package allows you to configure the HTTP client in several ways:

 * Timeout: Set a custom timeout (in seconds).

 * Example: httpclient.WithTimeout(10)

 * Custom Headers: Add custom headers to requests.

 * Example: httpclient.WithCustomHeader("User-Agent", "MyCustomClient/1.0")

 * Retries: Enable automatic retries and configure retry count and delay.

 * Example: httpclient.WithRetries(3), httpclient.WithRetryDelay(2)

 * Logging: Enable or disable request/response logging for debugging.

 * Example: httpclient.WithLogging(true)

## Contributing
I welcome contributions to improve the package! To contribute, please fork this repository, make your changes, and submit a pull request.

## License
This project is licensed under the MIT License - see the LICENSE file for details.
```css 
This version addresses formatting consistency, making sure each section has a clear header, and ensures the examples are correctly displayed with proper indentation. The package usage examples are consistent, and the overall readability of the document is improved.
```
```bash
go get github.com/srahkmli/httpclient
