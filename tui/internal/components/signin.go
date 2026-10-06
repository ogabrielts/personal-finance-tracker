package components

import (
	"encoding/json"
	"expense-api/tui/internal/api"
	"fmt"
	"io"
	"net/http"
	"net/http/cookiejar"
)

type signinResponse struct {
	Message string `json:"message"`
	Token string `json:"token"`
}

func SignIn(email, password string) (*http.Client, error) {
	jar, err := cookiejar.New(nil)
	if err != nil {
		return nil, err
	}

	client := &http.Client{
		Jar: jar,
	}

	res, err := petapi.SignIn(client, email, password)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	
	if res.StatusCode != 200 {
		return nil, fmt.Errorf("Failed to sign in: %s", res.Status)
	}

	resBody, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	body := signinResponse{}
	err = json.Unmarshal(resBody, &body)

	fmt.Println(body.Message)

	return client, nil
}
