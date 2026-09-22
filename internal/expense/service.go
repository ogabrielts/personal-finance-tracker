package expense

import (
	"context"
	"errors"
)

type expenseService struct {
	repo expenseRepository
}

func newService(repository expenseRepository) *expenseService {
	return &expenseService{
		repo: repository,
	}
}

func (es *expenseService) CreateExpenseService(ctx context.Context, name, category string, userId int, value float32) (*expense, error) {
	newExpense := &expense{
		ExpenseName: name,
		ExpenseValue: value,
		ExpenseCategory: category,
		ExpenseUser: userId,
	}

	if err := es.repo.CreateExpense(ctx, newExpense); err != nil {
		return nil, errors.New("failed to create expense")
	}

	return newExpense, nil
}

func (es *expenseService) GetExpenseByIdService(ctx context.Context, expenseId, userId int) (*expense, error) {
	return es.repo.GetExpenseById(ctx, expenseId, userId)
}

func (es *expenseService) GetAllExpensesService(ctx context.Context, userId int) ([]*expense, error) {
	return es.repo.GetAllExpenses(ctx, userId)
}

func (es *expenseService) UpdateExpenseService(ctx context.Context, expenseId, userId int, name, category string, value float32) (*expense, error) {
	updatedExpense := &expense{
		ExpenseId: expenseId,
		ExpenseName: name,
		ExpenseValue: value,
		ExpenseCategory: category,
		ExpenseUser: userId,
	}

	if err := es.repo.UpdateExpense(ctx, updatedExpense); err != nil {
		return nil, errors.New("failed to update expense")
	}

	return updatedExpense, nil
}

func (es *expenseService) DeleteExpenseService(ctx context.Context, expenseId, userId int) error {
	return es.repo.DeleteExpense(ctx, expenseId, userId)
}

func (es *expenseService) FilterByCategoryService(ctx context.Context, category string, userId int) ([]*expense, error) {
	return es.repo.FilterByCategory(ctx, category, userId)
}

func (es *expenseService) FilterByDateService(ctx context.Context, dateFrom, dateTo string, userId int) ([]*expense, error) {
	return es.repo.FilterByDate(ctx, dateFrom, dateTo, userId)
}
