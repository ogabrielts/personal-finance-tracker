package reports

type income struct {
	IncomeId int `json:"income_id"`
	IncomeName string `json:"income_name"`
	IncomeValue float32 `json:"income_value"`
	IncomeCategory string `json:"income_category"`
	IncomeUser int `json:"income_user"`
}

type expense struct {
	ExpenseId int `json:"expense_id"`
	ExpenseName string `json:"expense_name"`
	ExpenseValue float32 `json:"expense_value"`
	ExpenseCategory string `json:"expense_category"`
	ExpenseUser int `json:"expense_user"`
}
