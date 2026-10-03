package income

import (
	"context"
	"errors"
)

type incomeService struct {
	repo incomeRepository
}

func newService(repository incomeRepository) *incomeService {
	return &incomeService{
		repo: repository,
	}
}

func (is *incomeService) CreateIncomeService(ctx context.Context, name, category string, userId int, value float32) (*income, error) {
	newIncome := &income{
		IncomeName: name,
		IncomeValue: value,
		IncomeCategory: category,
		IncomeUser: userId,
	}

	if err := is.repo.CreateIncome(ctx, newIncome); err != nil {
		return nil, errors.New("failed to create income") 
	}

	return newIncome, nil
}

func (is *incomeService) GetIncomeByIdService(ctx context.Context, incomeId, userId int) (*income, error) {
	return is.repo.GetIncomeById(ctx, incomeId, userId)
}

func (is *incomeService) GetAllIncomesService(ctx context.Context, userId int) ([]*income, error) {
	return is.repo.GetAllIncomes(ctx, userId) 
}

func (is *incomeService) UpdateIncomeService(ctx context.Context, incomeId, userId int, name, category string, value float32) (*income, error) {
	updatedIncome := &income{
		IncomeId: incomeId,
		IncomeName: name,
		IncomeValue: value,
		IncomeCategory: category,
		IncomeUser: userId,
	}

	if err := is.repo.UpdateIncome(ctx, updatedIncome); err != nil {
		return nil, errors.New("failed to update income")
	}

	return updatedIncome, nil
}

func (is *incomeService) DeleteIncomeService(ctx context.Context, incomeId, userId int) error {
	return is.repo.DeleteIncome(ctx, incomeId, userId)
}

func (is *incomeService) FilterByCategoryService(ctx context.Context, category string, userId int) ([]*income, error) {
	return is.repo.FilterByCategory(ctx, category, userId)
}

func (is *incomeService) FilterByDateService(ctx context.Context, dateFrom, dateTo string, userId int) ([]*income, error) {
	return is.repo.FilterByDate(ctx, dateFrom, dateTo, userId)
}
