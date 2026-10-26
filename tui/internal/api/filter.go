package petapi

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

func FilterCategory(client *http.Client, account, category string) (*http.Response, error) {
	reqBody, err := json.Marshal(map[string]string{
		"category": category,
	})
	if err != nil {
		return nil, err
	}

	urlAdr := fmt.Sprintf("http://localhost:8000/v1/%s/filter/category", account)
	res, err := client.Post(urlAdr, "application/json", bytes.NewBuffer(reqBody))
	if err != nil {
		return nil, err
	}

	return res, nil
}

func FilterDate(client *http.Client, account, dateFrom, dateTo string) (*http.Response, error) {
	reqBody, err := json.Marshal(map[string]string{
		"datefrom": dateFrom,
		"dateto": dateTo,
	})
	if err != nil {
		return nil, err
	}

	urlAdr := fmt.Sprintf("http://localhost:8000/v1/%s/filter/date", account)
	res, err := client.Post(urlAdr, "application/json", bytes.NewBuffer(reqBody))
	if err != nil {
		return nil, err
	}

	return res, nil
}
