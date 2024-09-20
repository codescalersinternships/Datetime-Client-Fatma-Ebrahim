// This package implements a HTTP client in Go that consumes the datetime server.
// It supports two content types: plain text and JSON.
// It has two public functions: Client and Inputhandler.

package client

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
	"flag"
	"os"

	"github.com/cenkalti/backoff/v4"
)

// Client sends a request to the datetime server with certain content type (plain text or json)
// and print the output to the console.
// and returns the response statuscode, the response body and any countered error.
func Client(w io.Writer, url, contenttype string) (int, []byte, error) {
	// host := os.Getenv("SERVER_HOST")
	// port := os.Getenv("SERVER_PORT")
	// url := fmt.Sprintf("http://%s:%s/datetime", host, port)

	var response *http.Response
	connection := func() error {
		c := http.Client{Timeout: time.Duration(1) * time.Second}
		request, err := http.NewRequest("GET", url, nil)
		if err != nil {
			return err
		}
		request.Header.Add("Content-Type", contenttype)
		response, err = c.Do(request)

		if err != nil {
			return err
		}
		return nil
	}

	expBackoff := backoff.NewExponentialBackOff()
	expBackoff.MaxElapsedTime = 10 * time.Second
	err := backoff.Retry(connection, expBackoff)
	if err != nil {
		return http.StatusServiceUnavailable, nil, fmt.Errorf("error in server connection")
	}
	if response.StatusCode != http.StatusOK {
		return response.StatusCode, nil, fmt.Errorf("status code not OK")
	}

	defer response.Body.Close()
	body, err := io.ReadAll(response.Body)
	if err != nil {
		return response.StatusCode, body, err
	}

	data := texttype(body)
	if contenttype == "application/json" {
		data, err = jsontype(body)
		if err != nil {
			return response.StatusCode, body, err
		}
	}

	fmt.Fprintln(w, data)
	return response.StatusCode, body, nil
}

// Inputhandler parses command line arguments and flags to set the url and content type
// if no flags are provided, it uses the environment variables (SERVER_HOST , SERVER_PORT)
// if no environment variables are provided, it uses default values
// then returns the url and content type
func Inputhandler() (string,  string) {
	port := ""
	flag.StringVar(&port, "p", "", "port number")
	if port == "" {
		port = os.Getenv("SERVER_PORT")
		if port == "" {
			port="8080"
		}
	}
	host := ""
	flag.StringVar(&host, "h", "", "host name")
	if host == "" {
		host = os.Getenv("SERVER_HOST")
		if host == "" {
			host="localhost"
		}
	}
	contenttype := ""
	flag.StringVar(&contenttype, "c", "application/json", "content type")
	flag.Parse()
	url := flag.Arg(0)
	if url == "" {
		url = "http://" + host + ":" + port + "/datetime"
	}
	return url,contenttype
}

func texttype(body []byte) string {
	return string(body)
}
func jsontype(body []byte) (string, error) {
	var result map[string]interface{}
	err := json.Unmarshal(body, &result)
	if err != nil {
		return "", fmt.Errorf("error in json unmarshal")
	}
	return result["datetime"].(string), nil
}

