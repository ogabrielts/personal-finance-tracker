package components

import (
	petapi "expense-api/tui/internal/api"
	"flag"
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

	return petapi.Report(client, *reportMonth, *reportYear, *reportAll)
}
