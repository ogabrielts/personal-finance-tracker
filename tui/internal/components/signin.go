package components

import (
	"expense-api/tui/internal/api"
	"fmt"
	"net/http"
	"net/http/cookiejar"
)

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

	return client, nil
}
