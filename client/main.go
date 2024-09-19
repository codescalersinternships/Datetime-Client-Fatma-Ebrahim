package main

import (
	"os"

	client "github.com/codescalersinternships/Datetime-Client-Fatma-Ebrahim/pkg"
)

func main() {
	client.Client(os.Stdout, "application/json")
}
