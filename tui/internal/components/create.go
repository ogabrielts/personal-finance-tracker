package components

import (
	"expense-api/tui/internal/api"
	"expense-api/tui/internal/utils"
	"flag"
	"fmt"
	"io"
	"net/http"
)

func Create(client *http.Client, args []string) error {
	createCmd := flag.NewFlagSet("create", flag.ExitOnError)
	createName := createCmd.String("name", "", "A name description")
	createValue := createCmd.Float64("value", 0.0, "The value")
	createCategory := createCmd.String("category", "", "A category")
	createAccount := createCmd.String("account-type", "", "Account type")
	
	if err := createCmd.Parse(args); err != nil {
		return err
	}

	res, err := petapi.Create(client, *createAccount, *createName, *createCategory, float32(*createValue))
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if res.StatusCode != 201 {
		return fmt.Errorf("Failed to create new %s: %s\n", *createAccount, res.Status)
	}
	
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}

	if err := utils.DisplaySingleData(*createAccount, body); err != nil {
		return err
	}

	return nil
}


