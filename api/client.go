package client

import (
	"encoding/json"
	"io"
	"net/http"
)

func FetchData(url string, t interface{}) error {
	res, err := http.Get(url)
	if err != nil {
		return err
	}

	dataBytes, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}

	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return err
	}

	json.Unmarshal(dataBytes, &t)

	return nil
}
