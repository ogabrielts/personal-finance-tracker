package income

type income struct {
	IncomeId int `json:"income_id"`
	IncomeName string `json:"income_name"`
	IncomeValue float32 `json:"income_value"`
	IncomeCategory string `json:"income_category"`
	IncomeUser int `json:"income_user"`
}
