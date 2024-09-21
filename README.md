# Datetime Client - Fatma Ebrahim

This package implements an HTTP client in Go that consumes the datetime server. It supports two content types: plain text and JSON.

## Installation

To install the client package, run the following command:

```shell
go get github.com/codescalersinternships/Datetime-Client-Fatma-Ebrahim/client
```
To install the needed dependencies:

```shell
go mod download
```

## Usage

Here's an example of how to use the `Client` package:

In a `main.go` file:

```go
package main

import (
    "fmt"
    "os"
    "github.com/codescalersinternships/Datetime-Client-Fatma-Ebrahim/client"
)

func main() {
    url, contentType := client.InputHandler()
    statusCode, result, err := client.Client(os.Stdout, url, contentType)
    fmt.Println(statusCode, string(result), err)
}
```

In the terminal, you can use flags: `-p` for port, `-h` for host, and `-c` for content type:

```shell
go run main.go -p 8000 -h localhost -c application/json
```

Or you can use environment variables `SERVER_PORT` for the port and `SERVER_HOST` for the host:

```shell
export SERVER_PORT=8080 SERVER_HOST=localhost

go run main.go
```

Alternatively, you can simply use the default port number and host to run on localhost using port 8080:

```shell
go run main.go
```

Now you can see the response (date and time) in the terminal!
