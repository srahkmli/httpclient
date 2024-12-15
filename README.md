
HTTP Client Package
A simple and efficient Go package for making HTTP requests, handling responses, and managing common HTTP client utilities like retries, timeouts, and headers.

Table of Contents
Installation
Usage
Basic Example
Custom Configuration
Handling Timeouts
Retries
Features
Contributing
License
Installation
To install the httpclient package, use Go's package manager:

bash
Copy code
go get github.com/yourusername/httpclient
Usage
Basic Example
Here's how to use the httpclient package with a simple GET request:

go
Copy code
package main

import (
    "fmt"
    "log"
    "github.com/yourusername/httpclient"
)

func main() {
    client := httpclient.New()

    response, err := client.Get("https://api.example.com/data")
    if err != nil {
        log.Fatalf("Error making GET request: %v", err)
    }
    fmt.Println("Response:", response)
}
Custom Configuration
You can customize the HTTP client with options like custom headers, timeouts, or proxies.

go
Copy code
package main

import (
    "fmt"
    "log"
    "github.com/yourusername/httpclient"
)

func main() {
    client := httpclient.New(
        httpclient.WithTimeout(10),                // Set timeout to 10 seconds
        httpclient.WithCustomHeader("User-Agent", "MyCustomClient/1.0"),
    )

    response, err := client.Get("https://api.example.com/data")
    if err != nil {
        log.Fatalf("Error making GET request: %v", err)
    }
    fmt.Println("Response:", response)
}
Handling Timeouts
To handle timeouts, you can configure the timeout value when initializing the client. The client will automatically cancel the request if it exceeds the specified timeout.

go
Copy code
package main

import (
    "fmt"
    "log"
    "github.com/yourusername/httpclient"
)

func main() {
    client := httpclient.New(httpclient.WithTimeout(5)) // Set timeout to 5 seconds

    response, err := client.Get("https://api.example.com/data")
    if err != nil {
        log.Fatalf("Request timed out or error occurred: %v", err)
    }
    fmt.Println("Response:", response)
}
Retries
The httpclient package supports automatic retries for failed requests. You can configure the number of retries and the delay between retries.

go
Copy code
package main

import (
    "fmt"
    "log"
    "github.com/yourusername/httpclient"
)

func main() {
    client := httpclient.New(
        httpclient.WithRetries(3),                // Retry up to 3 times
        httpclient.WithRetryDelay(2),             // 2-second delay between retries
    )

    response, err := client.Get("https://api.example.com/data")
    if err != nil {
        log.Fatalf("Error making GET request: %v", err)
    }
    fmt.Println("Response:", response)
}
Features
Customizable timeouts for requests.
Automatic retries on failed requests with configurable retry count and delay.
Custom headers to send with each request.
Flexible HTTP methods (GET, POST, PUT, DELETE).
Supports JSON encoding/decoding for APIs.
Error handling with structured responses.
Contributing
We welcome contributions to improve the package! To contribute, please fork this repository, make your changes, and submit a pull request.

Steps to Contribute:
Fork the repository.
Create a new branch (git checkout -b feature-branch).
Make your changes.
Commit your changes (git commit -am 'Add new feature').
Push to the branch (git push origin feature-branch).
Create a pull request.
License
This project is licensed under the MIT License - see the LICENSE file for details.
