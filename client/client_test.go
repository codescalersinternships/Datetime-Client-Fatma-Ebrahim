package client

import (
	"bytes"
	"testing"
)

func TestClient(t *testing.T) {

	buffer := bytes.Buffer{}
	err := Client(&buffer, "http://localhost:8080/datetime")

	if err != nil {
		t.Errorf("expected nil, got %v", err)
	}
}
