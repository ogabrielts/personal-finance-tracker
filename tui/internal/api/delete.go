package petapi

import (
	"fmt"
	"io"
	"net/http"
)

func Delete(client *http.Client, account, id string) error {
	urlAdr := fmt.Sprintf("http://localhost:8000/v1/%s/entry/%s", account, id)
	req, err := http.NewRequest("DELETE", urlAdr, nil)
	if err != nil {
		return err
	}

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
