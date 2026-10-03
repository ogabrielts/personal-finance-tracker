package user

import (
	"context"
	"errors"
	"expense-api/api/internal/utils"
)

type userService struct {
	repo userRepository
}

func newService(repository userRepository) *userService {
	return &userService{
		repo: repository,
	}
}

func (us *userService) CreateUserService(ctx context.Context, name, email, password, address string) (*user, error) {
	ok, err := us.repo.GetUserByEmail(ctx, email)
	if err == nil && ok != nil {
		return nil, errors.New("email already exists")
	}

	hashedPassword, err := utils.HashPassword(password)
	if err != nil {
		return nil, errors.New("system failure")
	}

	user := &user{
		UserName: name,
		UserEmail: email,
		UserPassword: hashedPassword,
		UserAddress: address,
	}

	if err := us.repo.CreateUser(ctx, user); err != nil {
		return nil, errors.New("failed to create user")
	}

	return user, nil
}

func (us *userService) GetUserByIdService(ctx context.Context, id int) (*user, error) {
	return us.repo.GetUserById(ctx, id)
}

func (us *userService) GetUserByEmailService(ctx context.Context, email string) (*user, error) {
	return us.repo.GetUserByEmail(ctx, email)
}

func (us *userService) UpdateUserService(ctx context.Context, id int, name, email, password, address string) (*user, error) {
	_, err := us.repo.GetUserById(ctx, id)
	if err != nil {
		return nil, errors.New("failed to find user")
	}

	var hashedPassword string
	if password != "" {
		hashedPassword, err = utils.HashPassword(password)
		if err != nil {
			return nil, errors.New("system failure")
		}
	}

	user := &user{
		UserId: id,
		UserName: name,
		UserEmail: email,
		UserPassword: hashedPassword,
		UserAddress: address,
	}

	if err := us.repo.UpdateUser(ctx, user); err != nil {
		return nil, errors.New("failed to update user")
	}

	return user, nil
}

func (us *userService) DeleteUserService(ctx context.Context, id int) error {
	return us.repo.DeleteUser(ctx, id)
}
