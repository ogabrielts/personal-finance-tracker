package reports

import "context"

type reportsRepository interface {
	GetAllExpenses(ctx context.Context, userId int) ([]*expense, error)
	GetAllIncomes(ctx context.Context, userId int) ([]*income, error)
	GetExpensesByDate(ctx context.Context, userId int, dateFrom, dateTo string) ([]*expense, error)
	GetIncomesByDate(ctx context.Context, userId int, dateFrom, dateTo string) ([]*income, error)
}
