package client

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/cenkalti/backoff/v4"
)

func Client(w io.Writer, contenttype string) error {
	host := os.Getenv("SERVER_HOST")
	port := os.Getenv("SERVER_PORT")
	url := fmt.Sprintf("http://%s:%s/datetime", host, port)
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
		log.Fatal("failed to connect to datetime server:", err)
	}

	defer response.Body.Close()
	body, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	data := texttype(body)
	if contenttype == "application/json" {
		data, err = jsontype(body)
		if err != nil {
			log.Fatal(err)
		}
	}

	fmt.Fprintln(w, data)
	return nil
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

// func connection(url, contenttype string) (*http.Response, error) {
// 	c := http.Client{Timeout: time.Duration(1) * time.Second}
// 	request, err := http.NewRequest("GET", url, nil)
// 	if err != nil {
// 		return nil, err
// 	}
// 	request.Header.Add("Content-Type", contenttype)
// 	response, err := c.Do(request)

// 	if err != nil {
// 		return nil, err
// 	}
// 	return response, nil
// }
