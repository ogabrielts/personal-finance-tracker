package income

import (
	"context"
	"database/sql"
	"fmt"
)

type incomeDB struct {
	db *sql.DB
}

func newDB(db *sql.DB) *incomeDB {
	return &incomeDB{
		db: db,
	}
}

func (idb *incomeDB) CreateIncome(ctx context.Context, income *income) error {
	_, err := idb.db.ExecContext(ctx, "INSERT INTO pet_income (income_name, income_value, income_category, income_user) VALUES ($1, $2, $3, $4)", income.IncomeName, income.IncomeValue, income.IncomeCategory, income.IncomeUser)
	if err != nil {
		return fmt.Errorf("failed to insert data: %w", err)
	}

	return nil
}

func (idb *incomeDB) GetIncomeById(ctx context.Context, incomeId, userId int) (*income, error) {
	var income income

	err := idb.db.QueryRowContext(ctx, "SELECT income_id, income_name, income_value, income_category, income_user FROM pet_income WHERE income_id = $1 AND income_user = $2", incomeId, userId).Scan(&income.IncomeId, &income.IncomeName, &income.IncomeValue, &income.IncomeCategory, &income.IncomeUser)
	if err != nil {
		return nil, fmt.Errorf("failed to query data: %w", err)
	}

	return &income, nil
}

func (idb *incomeDB) GetAllIncomes(ctx context.Context, userId int) ([]*income, error) {
	var incomes []*income

	rows, err := idb.db.QueryContext(ctx, "SELECT income_id, income_name, income_value, income_category, income_user FROM pet_income WHERE income_user = $1", userId)
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

func (idb *incomeDB) UpdateIncome(ctx context.Context, income *income) error {
	if income.IncomeName != "" {
		_, err := idb.db.ExecContext(ctx, "UPDATE pet_income SET income_name = $1 WHERE income_id = $2 AND income_user = $3", income.IncomeName, income.IncomeId, income.IncomeUser)
		if err != nil {
			return fmt.Errorf("failed to update data: %w", err)
		}
	}

	if income.IncomeValue != 0.0 {
		_, err := idb.db.ExecContext(ctx, "UPDATE pet_income SET income_value = $1 WHERE income_id = $2 AND income_user = $3", income.IncomeValue, income.IncomeId, income.IncomeUser)
		if err != nil {
			return fmt.Errorf("failed to update data: %w", err)
		}
	}

	if income.IncomeCategory != "" {
		_, err := idb.db.ExecContext(ctx, "UPDATE pet_income SET income_category = $1 WHERE income_id = $2 AND income_user = $3", income.IncomeCategory, income.IncomeId, income.IncomeUser)
		if err != nil {
			return fmt.Errorf("failed to update data: %w", err)
		}
	}

	return nil
}

func (idb *incomeDB) DeleteIncome(ctx context.Context, incomeId, userId int) error {
	_, err := idb.db.ExecContext(ctx, "DELETE FROM pet_income WHERE income_id = $1 AND income_user = $2", incomeId, userId)
	if err != nil {
		return fmt.Errorf("failed to delete data: %w", err)
	}

	return nil
}
