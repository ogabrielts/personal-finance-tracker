package utils

import (
	"encoding/json"
	"expense-api/tui/internal/models"
	"fmt"
)

func DisplaySingleData(account string, resBody []byte) error {
	switch account {
		case "income":
			var body models.Income
			if err := json.Unmarshal(resBody, &body); err != nil {
				return err
			}

			fmt.Printf("Id: %d\n", body.Id)
			fmt.Printf("Name: %s\n", body.Name)
			fmt.Printf("Value: %.2f\n", body.Value)
			fmt.Printf("Category: %s\n", body.Category)

		case "expense":
			var body models.Expense
			if err := json.Unmarshal(resBody, &body); err != nil {
				return err
			}

			fmt.Printf("Id: %d\n", body.Id)
			fmt.Printf("Name: %s\n", body.Name)
			fmt.Printf("Value: %.2f\n", body.Value)
			fmt.Printf("Category: %s\n\n", body.Category)
		
	}

	return nil
}

func DisplayMultipleData(account string, resBody []byte) error {
	switch account {
	case "income":
		var incomes []models.Income

		if err := json.Unmarshal(resBody, &incomes); err != nil {
			return err
		}

		fmt.Println("All my incomes:")
		for _, v := range incomes {
			fmt.Printf("Id: %d\n", v.Id)
			fmt.Printf("Name: %s\n", v.Name)
			fmt.Printf("Value: %.2f\n", v.Value)
			fmt.Printf("Category: %s\n\n", v.Category)
		}

	case "expense":
		var expenses []models.Expense

		if err := json.Unmarshal(resBody, &expenses); err != nil {
			return err
		}

		fmt.Println("All my expenses:")
		for _, v := range expenses {
			fmt.Printf("Id: %d\n", v.Id)
			fmt.Printf("Name: %s\n", v.Name)
			fmt.Printf("Value: %.2f\n", v.Value)
			fmt.Printf("Category: %s\n\n", v.Category)
		}
	}

	return nil
}

func DisplayReport(resBody []byte) error {
	var report models.Report

	if err := json.Unmarshal(resBody, &report); err != nil {
		return err
	}
	
	for k, v := range report.Incomes {
		fmt.Printf("%s: %.2f\n", k, v)
	}
	fmt.Printf("Total Income: %.2f\n", report.TotalIncome)

	for k, v := range report.Expenses {
		fmt.Printf("%s: %.2f\n", k, v)
	}
	fmt.Printf("Total Expenses: %.2f\n", report.TotalExpense)
	fmt.Printf("Final Balance: %.2f\n\n", report.Balance)

	return nil
}
