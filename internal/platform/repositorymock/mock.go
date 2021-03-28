package repositorymock

import (
	"github.com/jorgeAM/simple-api/internal/user/domain"
	"github.com/stretchr/testify/mock"
)

type UserMockRepository struct {
	mock.Mock
}

func (u *UserMockRepository) NewUser(user *domain.User) error {
	args := u.Called(user)
	return args.Error(0)
}

func (u *UserMockRepository) GetUsers() ([]*domain.User, error) {
	args := u.Called()
	return args.Get(0).([]*domain.User), args.Error(1)
}

func (u *UserMockRepository) GetUser(id string) (*domain.User, error) {
	args := u.Called(id)
	return args.Get(0).(*domain.User), args.Error(1)
}

func (u *UserMockRepository) UpdateUser(user *domain.User) (*domain.User, error) {
	args := u.Called(user)
	return args.Get(0).(*domain.User), args.Error(1)
}

func (u *UserMockRepository) DeleteUser(id string) error {
	args := u.Called(id)
	return args.Error(0)
}
