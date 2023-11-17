package handlers

import (
	"encoding/json"
	"io"
	"net/http"
)

// GetRequest respData must be an address
func GetRequest(urlReq string, respData interface{}) error {

	resp, err := http.Get(urlReq)
	if err != nil {
		return err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	err = json.Unmarshal(body, respData)
	if err != nil {
		return err
	}

	return nil
}
