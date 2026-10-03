package reports

import (
	"context"
	"database/sql"
	"fmt"
)

type reportsDB struct {
	db *sql.DB
}

func newDB(db *sql.DB) *reportsDB {
	return &reportsDB{
		db: db,
	}
}

func (rdb *reportsDB) GetAllExpenses(ctx context.Context, userId int) ([]*expense, error) {
	var expenses []*expense

	rows, err := rdb.db.QueryContext(ctx, "SELECT expense_id, expense_name, expense_value, expense_category, expense_user FROM pet_expense WHERE expense_user = $1", userId)
	if err != nil {
		return nil, fmt.Errorf("failed to query data: %w", err)
	}
	defer rows.Close()

	var id, user int
	var name, category string
	var value float32
	
	for rows.Next() {
		if err := rows.Scan(&id, &name, &value, &category, &user); err != nil {
			return nil, fmt.Errorf("failed to query data: %w", err)
		}

		expense := &expense{
			ExpenseId: id,
			ExpenseName: name,
			ExpenseValue: value,
			ExpenseCategory: category,
			ExpenseUser: user,
		}

		expenses = append(expenses, expense)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("failed to query data: %w", err)
	}

	return expenses, nil
}

func (rdb *reportsDB) GetAllIncomes(ctx context.Context, userId int) ([]*income, error) {
	var incomes []*income

	rows, err := rdb.db.QueryContext(ctx, "SELECT income_id, income_name, income_value, income_category, income_user FROM pet_income WHERE income_user = $1", userId)
	if err != nil {
		return nil, fmt.Errorf("failed to query data: %w", err)
	}
	defer rows.Close()

	var id, user int
	var name, category string
	var value float32
	
	for rows.Next() {
		if err := rows.Scan(&id, &name, &value, &category, &user); err != nil {
			return nil, fmt.Errorf("failed to query data: %w", err)
		}
		
		income := &income{
			IncomeId: id,
			IncomeName: name,
			IncomeValue: value,
			IncomeCategory: category,
			IncomeUser: user,
		}

		incomes = append(incomes, income)

	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("failed to query data: %w", err)
	}

	return incomes, nil
}

func (rdb *reportsDB) GetExpensesByDate(ctx context.Context, userId int, dateFrom, dateTo string) ([]*expense, error) {
	var expenses []*expense

	rows, err := rdb.db.QueryContext(ctx, "SELECT expense_id, expense_name, expense_value, expense_category, expense_user FROM pet_expense WHERE expense_user = $1 AND created_at >= $2 AND created_at < $3", userId, dateFrom, dateTo)
	if err != nil {
		return nil, fmt.Errorf("failed to query data: %w", err)
	}
	defer rows.Close()

	var id, user int
	var name, category string
	var value float32
	
	for rows.Next() {
		if err := rows.Scan(&id, &name, &value, &category, &user); err != nil {
			return nil, fmt.Errorf("failed to query data: %w", err)
		}

		expense := &expense{
			ExpenseId: id,
			ExpenseName: name,
			ExpenseValue: value,
			ExpenseCategory: category,
			ExpenseUser: user,
		}

		expenses = append(expenses, expense)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("failed to query data: %w", err)
	}

	return expenses, nil
}

func (rdb *reportsDB) GetIncomesByDate(ctx context.Context, userId int, dateFrom, dateTo string) ([]*income, error) {
	var incomes []*income

	rows, err := rdb.db.QueryContext(ctx, "SELECT income_id, income_name, income_value, income_category, income_user FROM pet_income WHERE income_user = $1 AND created_at >= $2 AND created_at < $3", userId, dateFrom, dateTo)
	if err != nil {
		return nil, fmt.Errorf("failed to query data: %w", err)
	}
	defer rows.Close()

	var id, user int
	var name, category string
	var value float32
	
	for rows.Next() {
		if err := rows.Scan(&id, &name, &value, &category, &user); err != nil {
			return nil, fmt.Errorf("failed to query data: %w", err)
		}
		
		income := &income{
			IncomeId: id,
			IncomeName: name,
			IncomeValue: value,
			IncomeCategory: category,
			IncomeUser: user,
		}

		incomes = append(incomes, income)

	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("failed to query data: %w", err)
	}

	return incomes, nil
}
