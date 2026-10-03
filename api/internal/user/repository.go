package user

import "context"

type userRepository interface {
	CreateUser(ctx context.Context, user *user) error
	GetUserById(ctx context.Context, id int) (*user, error)
	GetUserByEmail(ctx context.Context, email string) (*user, error)
	UpdateUser(ctx context.Context, user *user) error
	DeleteUser(ctx context.Context, id int) error
}