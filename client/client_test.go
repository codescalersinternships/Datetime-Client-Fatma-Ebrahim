package client

import (
	"bytes"
	"fmt"
	"testing"
)

func TestClient(t *testing.T) {
	buffer := bytes.Buffer{}
	contenttype := "application/json"
	err := Client(&buffer, contenttype)

	fmt.Println(buffer.String())

	if err != nil {
		t.Errorf("expected nil, got %v", err)
	}
}
