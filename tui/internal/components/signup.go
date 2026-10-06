package components

import (
	petapi "expense-api/tui/internal/api"
	"fmt"
)

func SignUp(name, email, password, address string) error {
	res, err := petapi.SignUp(name, email, password, address)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if res.StatusCode != 201 {
		return fmt.Errorf("Failed to create account: %s\n", res.Status)
	}
	
	return nil
}
