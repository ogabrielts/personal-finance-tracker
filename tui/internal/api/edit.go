package petapi

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

func Edit(client *http.Client, account, id, name, category string, value float32) (*http.Response, error) {	
	reqBody, err := json.Marshal(map[string]any{
		"name": name,
		"category": category,
		"value": value,
	})
	if err != nil {
		return nil, err
	}
	
	urlAdr := fmt.Sprintf("http://localhost:8000/v1/%s/entry/%s", account, id)
	req, err := http.NewRequest("PUT", urlAdr, bytes.NewBuffer(reqBody))
	if err != nil {
		return nil, err
	}
	
	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	return res, nil
}
