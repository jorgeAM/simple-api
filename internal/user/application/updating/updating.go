package updating

import (
	"context"

	"github.com/jorgeAM/api/internal/user/domain"
)

type UserUpdatingService struct {
	repository domain.Repository
}

func NewUserUpdatingService(repository domain.Repository) *UserUpdatingService {
	return &UserUpdatingService{repository}
}

func (u *UserUpdatingService) UpdateUser(_ context.Context, user *domain.User) (*domain.User, error) {
	return u.repository.UpdateUser(user)
}
