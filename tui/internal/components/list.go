package components

import (
	"expense-api/tui/internal/api"
	"expense-api/tui/internal/utils"
	"flag"
	"fmt"
	"io"
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

	res, err := petapi.List(client, *listAccount, *listId, *listAll)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		return fmt.Errorf("Failed to list %s: %s\n", *listAccount, res.Status)
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}

	if *listAll {
		if err := utils.DisplayMultipleData(*listAccount, body); err != nil {
			return err
		}
	} else {
		if err := utils.DisplaySingleData(*listAccount, body); err != nil {
			return err
		}
	}

	return nil
}
