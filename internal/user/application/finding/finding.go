package finding

import (
	"context"

	"github.com/jorgeAM/simple-api/internal/user/domain"
)

type UserRetrieveOneService struct {
	repository domain.Repository
}

func NewUserRetrieveOneService(repository domain.Repository) *UserRetrieveOneService {
	return &UserRetrieveOneService{repository}
}

func (u *UserRetrieveOneService) FindUserByID(ctx context.Context, id string) (*domain.User, error) {
	return u.repository.GetUser(ctx, id)
}
