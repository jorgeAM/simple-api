package repositorymock

import (
	"context"

	"github.com/jorgeAM/simple-api/internal/user/domain"
	"github.com/stretchr/testify/mock"
)

type UserMockRepository struct {
	mock.Mock
}

func (u *UserMockRepository) NewUser(ctx context.Context, user *domain.User) error {
	args := u.Called(ctx, user)
	return args.Error(0)
}

func (u *UserMockRepository) GetUsers(ctx context.Context) ([]*domain.User, error) {
	args := u.Called(ctx)
	return args.Get(0).([]*domain.User), args.Error(1)
}

func (u *UserMockRepository) GetUser(ctx context.Context, id string) (*domain.User, error) {
	args := u.Called(ctx, id)
	return args.Get(0).(*domain.User), args.Error(1)
}

func (u *UserMockRepository) DeleteUser(ctx context.Context, id string) error {
	args := u.Called(ctx, id)
	return args.Error(0)
}
