package creating

import (
	"context"

	"github.com/jorgeAM/api/internal/user/domain"
)

type UserCreatingService struct {
	repository domain.Repository
}

func NewUserCreatingService(repository domain.Repository) *UserCreatingService {
	return &UserCreatingService{repository}
}

func (u *UserCreatingService) CreateNewUser(_ context.Context, id, username, firstName, lastName string) error {
	user := &domain.User{
		ID:        id,
		Username:  username,
		FirstName: firstName,
		LastName:  lastName,
	}

	return u.repository.NewUser(user)
}
