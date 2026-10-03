package petapi

import (
	"fmt"
	"io"
	"net/http"
)

func List(client *http.Client, account, id string, all bool) error {
	var urlAdr string

	if all {
		urlAdr = fmt.Sprintf("http://localhost:8000/v1/%s", account)
	} else {
		urlAdr = fmt.Sprintf("http://localhost:8000/v1/%s/entry/%s", account, id)
	}

	req, err := http.NewRequest("GET", urlAdr, nil)
	if err != nil {
		return err
	}

	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}

	fmt.Println(string(body))

	return nil
}
