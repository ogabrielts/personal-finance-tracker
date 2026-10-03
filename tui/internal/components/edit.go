package components

import (
	"expense-api/tui/internal/api"
	"flag"
	"net/http"
)

func Edit(client *http.Client, args []string) error {
	editCmd := flag.NewFlagSet("edit", flag.ExitOnError)
	editId := editCmd.String("id", "", "The reference code")
	editName := editCmd.String("name", "", "A name description")
	editValue := editCmd.Float64("value", 0.0, "The value")
	editCategory := editCmd.String("category", "", "A category")
	editAccount := editCmd.String("account-type", "", "Account type")

	if err := editCmd.Parse(args); err != nil {
		return err
	}
		
	return petapi.Edit(client, *editAccount, *editId, *editName, *editCategory, float32(*editValue))
}
