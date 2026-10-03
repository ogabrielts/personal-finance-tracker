package petapi

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func Report(client *http.Client, month string, year int, all bool) error {
	reqBody, err := json.Marshal(map[string]any{
		"month": month,
		"year": year,
	})
	if err != nil {
		return err
	}

	var	urlAdr string
	if all {	
		urlAdr = "http://localhost:8000/v1/reports/full"
	} else {
		urlAdr = "http://localhost:8000/v1/reports/monthly"
	}

	req, err := http.NewRequest("POST", urlAdr, bytes.NewBuffer(reqBody))
	if err != nil {
		return err
	}

	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}

	fmt.Println(string(body))

	return nil
}
