package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	client "github.com/codescalersinternships/Datetime-Client-Fatma-Ebrahim/pkg"
)

func TestClient(t *testing.T) {
	url := "http://localhost:8080/datetime"
	t.Run("Client with Json content type", func(t *testing.T) {
		buffer := bytes.Buffer{}
		contenttype := "application/json"

		statuscode, result, err := client.Client(&buffer, url, contenttype)
		if err != nil {
			t.Errorf("expected nil, got error %v", err)
		}
		if statuscode != http.StatusOK {
			t.Errorf("expected status code %d, got %d", http.StatusOK, statuscode)
		}
		var jsonbody map[string]interface{}
		if json.Unmarshal(result, &jsonbody) != nil {
			t.Errorf("expected nil ,got error in json format")
		}
	})

	t.Run("Client with plain text type", func(t *testing.T) {
		buffer := bytes.Buffer{}
		contenttype := "plain text"
		statuscode, result, err := client.Client(&buffer, url, contenttype)
		if err != nil {
			t.Errorf("expected nil, got error %v", err)
		}
		if statuscode != http.StatusOK {
			t.Errorf("expected status code %d, got %d", http.StatusOK, statuscode)
		}
		var jsonbody map[string]interface{}
		if json.Unmarshal(result, &jsonbody) == nil {
			t.Errorf("expected error in json format got nil")
		}
	})

}

func TestClientbyMock(t *testing.T) {
	t.Run("Client with Json content type", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			contenttype := r.Header.Get("Content-Type")
			if contenttype == "application/json" {
				datetime := time.Now().Local().Format("Monday 02-01-2006 15:04:05")
				data, _ := json.Marshal(map[string]string{"datetime": datetime})
				fmt.Fprint(w, string(data))
			} else {

				fmt.Fprint(w, time.Now().Local().Format("Monday 02-01-2006 15:04:05"))
			}
		}))
		defer server.Close()
		buffer := bytes.Buffer{}
		statuscode, _, err := client.Client(&buffer, server.URL, "application/json")
		if err != nil {
			t.Errorf("expected nil, got error %v", err)
		}
		if statuscode != http.StatusOK {
			t.Errorf("expected status code %d, got %d", http.StatusOK, statuscode)
		}
	})
}
