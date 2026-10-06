package petapi

import (
	"fmt"
	"net/http"
)

func List(client *http.Client, account, id string, all bool) (*http.Response, error) {
	var urlAdr string

	if all {
		urlAdr = fmt.Sprintf("http://localhost:8000/v1/%s", account)
	} else {
		urlAdr = fmt.Sprintf("http://localhost:8000/v1/%s/entry/%s", account, id)
	}

	req, err := http.NewRequest("GET", urlAdr, nil)
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
