package petapi

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
)

func SignUp(name, email, password, address string) ([]byte, error) {
	reqBody, err := json.Marshal(map[string]string{
		"name": name,
		"email": email,
		"password": password,
		"address": address,
	})
	if err != nil {
		return nil, err
	}

	res, err := http.Post("http://localhost:8000/v1/user", "application/json", bytes.NewBuffer(reqBody))
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
