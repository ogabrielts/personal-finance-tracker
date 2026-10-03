package petapi

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
)

func SignIn(client *http.Client, email, password string) ([]byte, error) {
	reqBody, err := json.Marshal(map[string]string{
		"email": email,
		"password": password,
	})
	if err != nil {
		return nil, err
	}

	res, err := client.Post("http://localhost:8000/v1/login", "application/json", bytes.NewBuffer(reqBody))
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}
