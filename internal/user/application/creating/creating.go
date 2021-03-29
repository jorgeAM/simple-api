package creating

import (
	"context"

	"github.com/jorgeAM/simple-api/internal/user/domain"
)

type UserCreatingService struct {
	repository domain.Repository
}

func NewUserCreatingService(repository domain.Repository) *UserCreatingService {
	return &UserCreatingService{repository}
}

func (u *UserCreatingService) CreateNewUser(ctx context.Context, id, username, firstName, lastName string) error {
	user, err := domain.NewUser(id, username, firstName, lastName)

	if err != nil {
		return err
	}

	return u.repository.NewUser(ctx, user)
}
