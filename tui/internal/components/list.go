package components

import (
	"expense-api/tui/internal/api"
	"flag"
	"net/http"
)

func List(client *http.Client, args []string) error {
	listCmd := flag.NewFlagSet("list", flag.ExitOnError)
	listAccount := listCmd.String("account-type", "", "Account type")
	listId := listCmd.String("id", "", "The reference code")
	listAll := listCmd.Bool("all", false, "List all")

	if err := listCmd.Parse(args); err != nil {
		return err
	}

	return petapi.List(client, *listAccount, *listId, *listAll)
}
