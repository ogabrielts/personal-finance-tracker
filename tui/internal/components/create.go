package components

import (
	"expense-api/tui/internal/api"
	"flag"
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

	return petapi.Create(client, *createAccount, *createName, *createCategory, float32(*createValue))
}
