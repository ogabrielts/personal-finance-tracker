package auth

import "context"

type authRepository interface {
	GetUser(ctx context.Context, email string) (*authUser, error)
}
