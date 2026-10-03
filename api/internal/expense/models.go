package expense

type expense struct {
	ExpenseId int `json:"expense_id"`
	ExpenseName string `json:"expense_name"`
	ExpenseValue float32 `json:"expense_value"`
	ExpenseCategory string  `json:"expense_category"`
	ExpenseUser int `json:"expense_user"`
}