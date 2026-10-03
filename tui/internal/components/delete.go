package components

import (
	"expense-api/tui/internal/api"
	"flag"
	"net/http"
)

func Delete(client *http.Client, args []string) error {
	deleteCmd := flag.NewFlagSet("delete", flag.ExitOnError)
	deleteId := deleteCmd.String("id", "", "The reference code")
	deleteAccount := deleteCmd.String("account-type", "", "Account type")

	if err := deleteCmd.Parse(args); err != nil {
		return err
	}

	return petapi.Delete(client, *deleteAccount, *deleteId)
}
