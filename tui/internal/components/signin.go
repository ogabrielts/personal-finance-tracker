package components

import (
	"encoding/json"
	"expense-api/tui/internal/api"
	"fmt"
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

	resBody, err := petapi.SignIn(client, email, password)
	if err != nil {
		return nil, err
	}

	body := signinResponse{}
	err = json.Unmarshal(resBody, &body)	
	
	if body.Message == "" {
		return nil, fmt.Errorf("failed to sign in")
	}

	return client, nil
}
