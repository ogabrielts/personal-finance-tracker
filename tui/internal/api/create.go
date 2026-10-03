package petapi

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func Create(client *http.Client, account, name, category string, value float32) error {	
	reqBody, err := json.Marshal(map[string]any{
		"name": name,
		"category": category,
		"value": value,
	})
	if err != nil {
		return err
	}

	urlAdr := fmt.Sprintf("http://localhost:8000/v1/%s", account)
	res, err := client.Post(urlAdr, "application/json", bytes.NewBuffer(reqBody))
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
