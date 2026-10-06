package user

import (
	"context"
	"database/sql"
	"errors"
	"expense-api/api/internal/utils"
	"fmt"
)

type userDB struct {
	db *sql.DB
}

func newDB(db *sql.DB) *userDB {
	return &userDB{
		db: db,
	}
}

func (udb *userDB) CreateUser(ctx context.Context, user *user) error {
	_, err := udb.db.ExecContext(ctx, "INSERT INTO pet_user (user_name, user_email, user_password, user_address) VALUES ($1, $2, $3, $4)", user.UserName, user.UserEmail, user.UserPassword, user.UserAddress)
	if err != nil {
		return fmt.Errorf("failed to insert data: %w", err)
	}

	return nil
}

func (udb *userDB) GetUserById(ctx context.Context, id int) (*user, error) {
	var user user

	err := udb.db.QueryRowContext(ctx, "SELECT user_id, user_name, user_email, user_password, user_address FROM pet_user WHERE user_id = $1", id).Scan(&user.UserId, &user.UserName, &user.UserEmail, &user.UserPassword, &user.UserAddress)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, utils.ErrNotFound
		}
		return nil, fmt.Errorf("failed to query data: %w", err)
	}

	return &user, nil
}

func (udb *userDB) GetUserByEmail(ctx context.Context, email string) (*user, error) {
	var user user

	err := udb.db.QueryRowContext(ctx, "SELECT user_id, user_name, user_email, user_password, user_address FROM pet_user WHERE user_email = $1", email).Scan(&user.UserId, &user.UserName, &user.UserEmail, &user.UserPassword, &user.UserAddress)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, utils.ErrNotFound
		}
		return nil, fmt.Errorf("failed to query data: %w", err)
	}

	return &user, nil
}

func (udb *userDB) UpdateUser(ctx context.Context, user *user) error {
	if user.UserName != "" {
		_, err := udb.db.ExecContext(ctx, "UPDATE pet_user SET user_name = $1 WHERE user_id = $2", user.UserName, user.UserId)
		if err != nil {
			return fmt.Errorf("failed to update data: %w", err)
		}
	}

	if user.UserEmail != "" {
		_, err := udb.db.ExecContext(ctx, "UPDATE pet_user SET user_email = $1 WHERE user_id = $2", user.UserEmail, user.UserId)
		if err != nil {
			return fmt.Errorf("failed to update data: %w", err)
		}
	}

	if user.UserPassword != "" {
		_, err := udb.db.ExecContext(ctx, "UPDATE pet_user SET user_password = $1 WHERE user_id = $2", user.UserPassword, user.UserId)
		if err != nil {
			return fmt.Errorf("failed to update data: %w", err)
		}
	}

	if user.UserAddress != "" {
		_, err := udb.db.ExecContext(ctx, "UPDATE pet_user SET user_address = $1 WHERE user_id = $2", user.UserAddress, user.UserId)
		if err != nil {
			return fmt.Errorf("failed to update data: %w", err)
		}
	}

	return nil
}

func (udb *userDB) DeleteUser(ctx context.Context, id int) error {
	_, err := udb.db.ExecContext(ctx, "DELETE FROM pet_user WHERE user_id = $1", id)
	if err != nil {
		return fmt.Errorf("failed to delete data: %w", err)
	}

	return nil
}
