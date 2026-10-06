package components

import (
	petapi "expense-api/tui/internal/api"
	"expense-api/tui/internal/utils"
	"flag"
	"fmt"
	"io"
	"net/http"
	"time"
)

func Report(client *http.Client, args []string) error {
	reportCmd := flag.NewFlagSet("report", flag.ExitOnError)
	reportMonth := reportCmd.String("month", "", "Which month to generate report")
	reportYear := reportCmd.Int("year", time.Now().Year(), "The year of the month to report")
	reportAll := reportCmd.Bool("all", false, "Report since the beginning")

	if err := reportCmd.Parse(args); err != nil {
		return err
	}

	res, err := petapi.Report(client, *reportMonth, *reportYear, *reportAll)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		return fmt.Errorf("Failed to generate report: %s\n", res.Status)
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}

	if err := utils.DisplayReport(body); err != nil {
		return err
	}

	return nil
}
