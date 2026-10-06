package components

import (
	"expense-api/tui/internal/api"
	"expense-api/tui/internal/utils"
	"flag"
	"fmt"
	"io"
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
		
	res, err := petapi.Edit(client, *editAccount, *editId, *editName, *editCategory, float32(*editValue))
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		return fmt.Errorf("Failed to update %s: %s\n", *editAccount, res.Status)
	}

	body, err := io.ReadAll(res.Body) 
	if err != nil {
		return err
	}

	if err := utils.DisplaySingleData(*editAccount, body); err != nil {
		return err
	}

	return nil
}
