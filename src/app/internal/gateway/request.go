package gateway

import (
	"io/ioutil"
	"log"
	"net/http"
)

// MakeOMDBRequest makes a request to the OMDB
func MakeOMDBRequest(title string, omdbURL string, apiKey string) ([]byte, error) {

	req, err := http.NewRequest(http.MethodGet, omdbURL, nil)
	if err != nil {
		log.Printf("error with new request: %+v\n", err)
		return nil, err
	}

	queryParams := req.URL.Query()

	queryParams.Set("apiKey", apiKey)
	queryParams.Set("t", title)

	req.URL.RawQuery = queryParams.Encode()

	client := http.Client{}
	resp, err := client.Do(req)
	if resp != nil {
		defer resp.Body.Close()
	}
	if err != nil {
		log.Printf("failed to perform request: %v %+v", req.Method, req.URL.String())
		return nil, err
	}

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("error reading body from performed request: %v %+v response code: %+v", req.Method, req.URL.String(), resp.StatusCode)
		return nil, err
	}

	return respBody, nil
}
