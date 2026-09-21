package auth

import (
	"context"
	"database/sql"
	"fmt"
)

type authDB struct {
	db *sql.DB
}

func newDB(db *sql.DB) *authDB {
	return &authDB{
		db: db,
	}
}

func (adb *authDB) GetUser(ctx context.Context, email string) (*authUser, error) {
	var user authUser

	err := adb.db.QueryRowContext(ctx, "SELECT user_id, user_name, user_email, user_password FROM pet_user WHERE user_email = $1", email).Scan(&user.UserId, &user.UserName, &user.UserEmail, &user.UserPassword)
	if err != nil {
		return nil, fmt.Errorf("failed to query data: %w", err)
	}

	return &user, nil
}
