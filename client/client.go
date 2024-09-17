package client

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
)

func Client(w io.Writer, url string) error {

	resp, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("error: %v", err)
	}
	defer resp.Body.Close()

	scanner := bufio.NewScanner(resp.Body)
	for scanner.Scan() {
		fmt.Fprintln(w, scanner.Text())
	}
	return nil
}
