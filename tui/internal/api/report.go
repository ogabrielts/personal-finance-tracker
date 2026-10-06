package petapi

import (
	"bytes"
	"encoding/json"
	"net/http"
)

func Report(client *http.Client, month string, year int, all bool) (*http.Response, error) {
	reqBody, err := json.Marshal(map[string]any{
		"month": month,
		"year": year,
	})
	if err != nil {
		return nil, err
	}

	var	urlAdr string
	if all {	
		urlAdr = "http://localhost:8000/v1/reports/full"
	} else {
		urlAdr = "http://localhost:8000/v1/reports/monthly"
	}

	req, err := http.NewRequest("POST", urlAdr, bytes.NewBuffer(reqBody))
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
