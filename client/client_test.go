package client

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"testing"
)

func TestClient(t *testing.T) {
	t.Run("Client with Json content type", func(t *testing.T) {
		buffer := bytes.Buffer{}
		contenttype := "application/json"
		statuscode, result,err := Client(&buffer, contenttype)
		if err !=nil{
			log.Fatal(err)
		}
		fmt.Print(string(result))
		if statuscode != http.StatusOK {
			t.Errorf("expected status code %d, got %d", http.StatusOK, statuscode)
		}
		_, err = jsontype(result)
		if err != nil {
			t.Errorf("expected nil, got %v", err)
		}
		fmt.Println(buffer.String())
	})

	t.Run("Client with plain text type", func(t *testing.T) {
		buffer := bytes.Buffer{}
		contenttype := "plain text"
		statuscode, result,err := Client(&buffer, contenttype)
		if err !=nil{
			log.Fatal(err)
		}
		if statuscode != http.StatusOK {
			t.Errorf("expected status code %d, got %d", http.StatusOK, statuscode)
		}
		_, err = jsontype(result)
		if err == nil {
			t.Errorf("expected error in json unmarshal, got nil")
		}
		fmt.Println(buffer.String())
	})

}
