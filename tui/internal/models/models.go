package models

type Income struct {
	Id int `json:"income_id"`
	Name string `json:"income_name"`
	Value float32 `json:"income_value"`
	Category string `json:"income_category"`
	User int `json:"income_user"`
}

type Expense struct {
	Id int `json:"expense_id"`
	Name string `json:"expense_name"`
	Value float32 `json:"expense_value"`
	Category string `json:"expense_category"`
	User int `json:"expense_user"`
}

type Report struct {
	Balance float32 `json:"balance"`
	Expenses map[string]float32 `json:"expenses"`
	Incomes map[string]float32 `json:"incomes"`
	TotalIncome float32 `json:"total_incomes"`
	TotalExpense float32 `json:"total_expenses"`
}
