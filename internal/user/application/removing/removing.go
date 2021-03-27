package removing

import (
	"context"

	"github.com/jorgeAM/api/internal/user/domain"
)

type UserRemovingService struct {
	repository domain.Repository
}

func NewUserRemovingService(repository domain.Repository) *UserRemovingService {
	return &UserRemovingService{repository}
}

func (u *UserRemovingService) RemoveUserByID(_ context.Context, id string) error {
	return u.repository.DeleteUser(id)
}
