package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"testing"

	client "github.com/codescalersinternships/Datetime-Client-Fatma-Ebrahim/pkg"
)

func TestClient(t *testing.T) {
	t.Run("Client with Json content type", func(t *testing.T) {
		buffer := bytes.Buffer{}
		contenttype := "application/json"
		statuscode, result, err := client.Client(&buffer, contenttype)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Print(string(result))
		if statuscode != http.StatusOK {
			t.Errorf("expected status code %d, got %d", http.StatusOK, statuscode)
		}
		var jsonbody map[string]interface{}
		if json.Unmarshal(result, &jsonbody) != nil {
			t.Errorf("expected nil ,got error in json format")
		}
		fmt.Println(buffer.String())
	})

	t.Run("Client with plain text type", func(t *testing.T) {
		buffer := bytes.Buffer{}
		contenttype := "plain text"
		statuscode, result, err := client.Client(&buffer, contenttype)
		if err != nil {
			log.Fatal(err)
		}
		if statuscode != http.StatusOK {
			t.Errorf("expected status code %d, got %d", http.StatusOK, statuscode)
		}
		var jsonbody map[string]interface{}
		if json.Unmarshal(result, &jsonbody) == nil {
			t.Errorf("expected error in json format got nil")
		}
		fmt.Println(buffer.String())
	})

}
