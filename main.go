package main

import (
	"fmt"
	"log"
	"os"

	client "github.com/codescalersinternships/Datetime-Client-Fatma-Ebrahim/pkg"
)

func main() {
	url, contenttype := client.Inputhandler()
	statuscode, result, err := client.GetDateTime(os.Stdout, url, contenttype)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(statuscode, string(result), err)
}
