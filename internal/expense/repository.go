package expense

import "context"

type expenseRepository interface {
	CreateExpense(ctx context.Context, expense *expense) error
	GetExpenseById(ctx context.Context, expenseId, userId int) (*expense, error)
	GetAllExpenses(ctx context.Context, userId int) ([]*expense, error)
	UpdateExpense(ctx context.Context, expense *expense) error
	DeleteExpense(ctx context.Context, expenseId, userId int) error

	FilterByCategory(ctx context.Context, category string, userId int) ([]*expense, error)
}
