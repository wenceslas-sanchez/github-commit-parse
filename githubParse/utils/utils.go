package utils

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func GetRequest(client *http.Client, url string) (*http.Response, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("connection error: %s", err)
	}
	req.Header = GITHUBHEADER

	res, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("bad response: %s", err)
	}
	if res.StatusCode != 200 {
		return nil, fmt.Errorf("response error: %s", res.Status)
	}
	return res, nil
}

func DecodeResponseJSON[T any](res *http.Response, schema *T) error {
	defer res.Body.Close()
	if err := json.NewDecoder(res.Body).Decode(schema); err != nil {
		return err
	}

	return nil
}
