package auth

import "context"

type authService struct {
	repo authRepository
}

func newService(repository authRepository) *authService {
	return &authService{
		repo: repository,
	}
}

func (as *authService) LogInService(ctx context.Context, email string) (*authUser, error) {
	return as.repo.GetUser(ctx, email)
}
