package income

import "context"

type incomeRepository interface {
	CreateIncome(ctx context.Context, income *income) error
	GetIncomeById(ctx context.Context, incomeId, userId int) (*income, error)
	GetAllIncomes(ctx context.Context, userId int) ([]*income, error)
	UpdateIncome(ctx context.Context, income *income) error
	DeleteIncome(ctx context.Context, incomeId, userId int) error

	FilterByCategory(ctx context.Context, category string, userId int) ([]*income, error)
	FilterByDate(ctx context.Context, dateFrom, dateTo string, userId int) ([]*income, error) 
}
