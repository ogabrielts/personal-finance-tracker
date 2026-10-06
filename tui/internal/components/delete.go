package components

import (
	"expense-api/tui/internal/api"
	"flag"
	"fmt"
	"net/http"
)

func Delete(client *http.Client, args []string) error {
	deleteCmd := flag.NewFlagSet("delete", flag.ExitOnError)
	deleteId := deleteCmd.String("id", "", "The reference code")
	deleteAccount := deleteCmd.String("account-type", "", "Account type")

	if err := deleteCmd.Parse(args); err != nil {
		return err
	}

	res, err := petapi.Delete(client, *deleteAccount, *deleteId)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		return fmt.Errorf("Failed to delete %s: %s\n", *deleteAccount, res.Status)
	}
	
	fmt.Printf("Successfully deleted %s\n\n", *deleteAccount)

	return nil
}
