package components

import (
	"expense-api/tui/internal/api"
	"expense-api/tui/internal/utils"
	"flag"
	"fmt"
	"io"
	"net/http"
)

func Filter(client *http.Client, args[]string) error {
	filterCmd := flag.NewFlagSet("filter", flag.ExitOnError)
	filterAccount :=  filterCmd.String("account-type", "", "Account type")
	filterCategory := filterCmd.String("category", "", "A category")
	filterDateFrom := filterCmd.String("from", "", "Starting date")
	filterDateTo := filterCmd.String("to", "", "End date")

	if err := filterCmd.Parse(args); err != nil {
		return err
	}
	
	var res *http.Response
	var err error
	if *filterCategory != "" && *filterDateFrom == "" && *filterDateTo == "" {
		res, err = petapi.FilterCategory(client, *filterAccount, *filterCategory)
		if err != nil {
			return err
		}
	} else if *filterCategory == "" && *filterDateFrom != "" && *filterDateTo != "" {
		res, err = petapi.FilterDate(client, *filterAccount, *filterDateFrom, *filterDateTo)
		if err != nil {
			return err
		}
	}
	defer res.Body.Close()
	
	if res.StatusCode != 200 {
		return fmt.Errorf("Failed to filter %s: %s\n", *filterAccount, res.Status)
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}

	if err := utils.DisplayMultipleData(*filterAccount, body); err != nil {
		return err
	}

	return nil
}
