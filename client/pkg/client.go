// This package implements a HTTP client in Go that consumes the datetime server.
// It supports two content types: plain text and JSON.
// It has two public functions: Client and Inputhandler.
package client

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/cenkalti/backoff/v4"
)

// Logger is used to log messages to logs.log file.
var (
	outfile, _ = os.Create("logs.log")
	logger     = log.New(outfile, "", 0)
)

// Client sends a request to the datetime server with certain content type (plain text or json)
// and print the output to the console.
// and returns the response statuscode, the response body and any countered error.
func Client(w io.Writer, url, contenttype string) (int, []byte, error) {
	var response *http.Response
	connection := func() error {
		c := http.Client{Timeout: time.Duration(1) * time.Second}
		request, err := http.NewRequest("GET", url, nil)
		if err != nil {
			logger.Println("Failed to create request", err)
			return err
		}
		request.Header.Add("Content-Type", contenttype)
		response, err = c.Do(request)
		logger.Println("Sending request to server")
		if err != nil {
			logger.Println("Failed to send request", err)
			return err
		}
		logger.Println("Request sent successfully")
		return nil
	}

	expBackoff := backoff.NewExponentialBackOff()
	expBackoff.MaxElapsedTime = 10 * time.Second
	err := backoff.Retry(connection, expBackoff)
	if err != nil {
		logger.Println("Failed to connect to server", err)
		return http.StatusServiceUnavailable, nil, fmt.Errorf("error in server connection")
	}
	logger.Println("Server returned status code", response.StatusCode)
	if response.StatusCode != http.StatusOK {
		return response.StatusCode, nil, fmt.Errorf("status code not OK")
	}

	defer response.Body.Close()
	body, err := io.ReadAll(response.Body)
	if err != nil {
		logger.Println("Failed to read response body", err)
		return response.StatusCode, body, err
	}

	data := texttype(body)
	if contenttype == "application/json" {
		data, err = jsontype(body)
		if err != nil {
			logger.Println("Failed to read response body", err)
			return response.StatusCode, body, err
		}
	}

	fmt.Fprintln(w, data)
	logger.Println("Response recieved successfully")
	return response.StatusCode, body, nil
}

// Inputhandler parses command line arguments and flags to set the url and content type
// if no flags are provided, it uses the environment variables (SERVER_HOST , SERVER_PORT)
// if no environment variables are provided, it uses default values
// then returns the url and content type
func Inputhandler() (string, string) {
	port := ""
	flag.StringVar(&port, "p", "", "port number")
	if port == "" {
		port = os.Getenv("SERVER_PORT")
		if port == "" {
			port = "8080"
		}
	}

	host := ""
	flag.StringVar(&host, "h", "", "host name")
	if host == "" {
		host = os.Getenv("SERVER_HOST")
		if host == "" {
			host = "localhost"
		}
	}

	contenttype := ""
	flag.StringVar(&contenttype, "c", "application/json", "content type")
	flag.Parse()
	url := flag.Arg(0)
	if url == "" {
		url = "http://" + host + ":" + port + "/datetime"
	}
	logger.Println("Requested URL:", url)
	return url, contenttype
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
