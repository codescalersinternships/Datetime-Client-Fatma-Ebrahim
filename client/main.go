package main

import (
	"fmt"
	"os"

	client "github.com/codescalersinternships/Datetime-Client-Fatma-Ebrahim/pkg"
)

func main() {
	url,contenttype:=client.Inputhandler()
	statuscode, result, err :=client.Client(os.Stdout, url,contenttype)
	fmt.Println(statuscode, string(result), err)
}
