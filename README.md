# Datetime-Client-Fatma-Ebrahim

This package implements a HTTP client in Go that consumes the datetime server. It supports two content types: plain text and JSON.

## Installation

To install the client package, run the following command:

```shell
go get /github.com/codescalersinternships/Datetime-Client-Fatma-Ebrahim/client
```

## Usage


Here's an example of how to use the `Client` package:

```go
package main

import (
    "fmt"
    "os"
    "/github.com/codescalersinternships/Datetime-Client-Fatma-Ebrahim/client"
)

func main() {
	client.Client(os.Stdout, "application/json")
}

```
Now You can see the response (date and time) in the terminal.

