package expense

import (
	"context"
	"database/sql"
	"fmt"
)

type expenseDB struct {
	db *sql.DB
}

func newDB(db *sql.DB) *expenseDB {
	return &expenseDB{
		db: db,
	}
}

func (edb *expenseDB) CreateExpense(ctx context.Context, expense *expense) error {
	_, err := edb.db.ExecContext(ctx, "INSERT INTO pet_expense (expense_name, expense_value, expense_category, expense_user) VALUES ($1, $2, $3, $4)", expense.ExpenseName, expense.ExpenseValue, expense.ExpenseCategory, expense.ExpenseUser)
	if err != nil {
		return fmt.Errorf("failed to insert data: %w", err)
	}

	return nil
}

func (edb *expenseDB) GetExpenseById(ctx context.Context, expenseId, userId int) (*expense, error) {
	var expense expense

	err := edb.db.QueryRowContext(ctx, "SELECT expense_id, expense_name, expense_value, expense_category, expense_user FROM pet_expense WHERE expense_id = $1 AND expense_user = $2", expenseId, userId).Scan(&expense.ExpenseId, &expense.ExpenseName, &expense.ExpenseValue, &expense.ExpenseCategory, &expense.ExpenseUser)
	if err != nil {
		return nil, fmt.Errorf("failed to query data: %w", err)
	}

	return &expense, nil
}

func (edb *expenseDB) GetAllExpenses(ctx context.Context, userId int) ([]*expense, error) {
	var expenses []*expense

	rows, err := edb.db.QueryContext(ctx, "SELECT expense_id, expense_name, expense_value, expense_category, expense_user FROM pet_expense WHERE expense_user = $1", userId)
	if err != nil {
		return nil, fmt.Errorf("failed to query data: %w", err)
	}
	defer rows.Close()

	var name, category string
	var id, user int
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

func (edb *expenseDB) UpdateExpense(ctx context.Context, expense *expense) error {
	if expense.ExpenseName != "" {
		_, err := edb.db.ExecContext(ctx, "UPDATE pet_expense SET expense_name = $1 WHERE expense_id = $2 AND expense_user = $3", expense.ExpenseName, expense.ExpenseId, expense.ExpenseUser)
		if err != nil {
			return fmt.Errorf("failed to update data: %w", err)
		}
	}

	if expense.ExpenseValue != 0.0 {
		_, err := edb.db.ExecContext(ctx, "UPDATE pet_expense SET expense_value = $1 WHERE expense_id = $2 AND expense_user = $3", expense.ExpenseValue, expense.ExpenseId, expense.ExpenseUser)
		if err != nil {
			return fmt.Errorf("failed to update data: %w", err)
		}
	}

	if expense.ExpenseCategory != "" {
		_, err := edb.db.ExecContext(ctx, "UPDATE pet_expense SET expense_category = $1 WHERE expense_id = $2 AND expense_user = $3", expense.ExpenseCategory, expense.ExpenseId, expense.ExpenseUser)
		if err != nil {
			return fmt.Errorf("failed to update data: %w", err)
		}
	}

	return nil
}

func (edb *expenseDB) DeleteExpense(ctx context.Context, expenseId, userId int) error {
	_, err := edb.db.ExecContext(ctx, "DELETE FROM pet_expense WHERE expense_id = $1 AND expense_user = $2", expenseId, userId)
	if err != nil {
		return fmt.Errorf("failed to delete data: %w", err)
	}

	return nil
}
