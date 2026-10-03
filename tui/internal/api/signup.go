package petapi

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func SignUp(name, email, password, address string) error {
	reqBody, err := json.Marshal(map[string]string{
		"name": name,
		"email": email,
		"password": password,
		"address": address,
	})
	if err != nil {
		return err
	}

	res, err := http.Post("http://localhost:8000/v1/user", "application/json", bytes.NewBuffer(reqBody))
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
