package removing

import (
	"context"

	"github.com/jorgeAM/simple-api/internal/user/domain"
)

type UserRemovingService struct {
	repository domain.Repository
}

func NewUserRemovingService(repository domain.Repository) *UserRemovingService {
	return &UserRemovingService{repository}
}

func (u *UserRemovingService) RemoveUserByID(ctx context.Context, id string) error {
	return u.repository.DeleteUser(ctx, id)
}
