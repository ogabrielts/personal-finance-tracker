package petapi

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

func Create(client *http.Client, account, name, category string, value float32) (*http.Response, error) {	
	reqBody, err := json.Marshal(map[string]any{
		"name": name,
		"category": category,
		"value": value,
	})
	if err != nil {
		return nil, err
	}

	urlAdr := fmt.Sprintf("http://localhost:8000/v1/%s", account)
	res, err := client.Post(urlAdr, "application/json", bytes.NewBuffer(reqBody))
	if err != nil {
		return nil, err
	}
	
	return res, nil
}
