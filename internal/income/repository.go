package income

import "context"

type incomeRepository interface {
	CreateIncome(ctx context.Context, income *income) error
	GetIncomeById(ctx context.Context, incomeId, userId int) (*income, error)
	GetAllIncomes(ctx context.Context, userId int) ([]*income, error)
	UpdateIncome(ctx context.Context, income *income) error
	DeleteIncome(ctx context.Context, incomeId, userId int) error
}
