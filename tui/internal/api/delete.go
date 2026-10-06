package petapi

import (
	"fmt"
	"net/http"
)

func Delete(client *http.Client, account, id string) (*http.Response, error) {
	urlAdr := fmt.Sprintf("http://localhost:8000/v1/%s/entry/%s", account, id)
	req, err := http.NewRequest("DELETE", urlAdr, nil)
	if err != nil {
		return nil, err
	}

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	return res, nil
}
