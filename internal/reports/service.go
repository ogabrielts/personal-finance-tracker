package reports

import (
	"context"
	"time"
)

type reportsService struct {
	repo reportsRepository
}

func newService(repository reportsRepository) *reportsService {
	return &reportsService{
		repo: repository,
	}
}

func convertTimestamp(month string, year int) (time.Time, time.Time, error) {
	t, err := time.Parse("January", month)
	if err != nil {
		return time.Time{}, time.Time{}, err
	}

	monthNum := t.Month()

	start := time.Date(year, monthNum, 1, 0, 0, 0, 0, time.UTC)
	end := time.Date(year, monthNum + 1, 0, 0, 0, 0, 0, time.UTC)

	return start, end, nil
}

func (rs *reportsService) GetAllExpensesService(ctx context.Context, userId int) (map[string]any, error) {
	report := make(map[string]any)

	expenses, err := rs.repo.GetAllExpenses(ctx, userId)
	if err != nil {
		return nil, err
	}

	incomes, err := rs.repo.GetAllIncomes(ctx, userId)
	if err != nil {
		return nil, err
	}
	
	expenseMap := map[string]float32{}
	var totalExpenses float32
	for _, v := range expenses {
		totalExpenses += v.ExpenseValue
		expenseMap[v.ExpenseName] = v.ExpenseValue
	}
	report["expenses"] = expenseMap

	incomesMap := map[string]float32{}
	var totalIncomes float32
	for _, v := range incomes {
		totalIncomes += v.IncomeValue
		incomesMap[v.IncomeName] = v.IncomeValue
	}
	report["incomes"] = incomesMap

	report["total_expenses"] = totalExpenses
	report["total_incomes"] = totalIncomes

	balance := totalIncomes - totalExpenses
	report["balance"] = balance

	return report, nil
}

func (rs *reportsService) MonthlyReportService(ctx context.Context, userId int, month string, year int) (map[string]any, error) {
	report := make(map[string]any)

	start, end, err := convertTimestamp(month, year)
	if err != nil {
		return nil, err
	}
	
	expenses, err := rs.repo.GetExpensesByDate(ctx, userId, start.Format("2006-01-02 15:04:05"), end.Format("2006-01-02 15:04:05"))
	if err != nil {
		return nil, err
	}

	incomes, err := rs.repo.GetIncomesByDate(ctx, userId, start.Format("2006-01-02 15:04:05"), end.Format("2006-01-02 15:04:05"))
	if err != nil {
		return nil, err
	}

	expenseMap := map[string]float32{}
	var totalExpenses float32
	for _, v := range expenses {
		totalExpenses += v.ExpenseValue
		expenseMap[v.ExpenseName] = v.ExpenseValue
	}
	report["expenses"] = expenseMap

	incomesMap := map[string]float32{}
	var totalIncomes float32
	for _, v := range incomes {
		totalIncomes += v.IncomeValue
		incomesMap[v.IncomeName] = v.IncomeValue
	}
	report["incomes"] = incomesMap

	report["total_expenses"] = totalExpenses
	report["total_incomes"] = totalIncomes

	balance := totalIncomes - totalExpenses
	report["balance"] = balance

	return report, nil
}
