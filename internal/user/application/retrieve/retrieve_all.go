package retrieve

import (
	"context"

	"github.com/jorgeAM/simple-api/internal/user/domain"
)

type UserRetrieveAllService struct {
	repository domain.Repository
}

func NewUserRetrieveAllService(repository domain.Repository) *UserRetrieveAllService {
	return &UserRetrieveAllService{repository}
}

func (u *UserRetrieveAllService) GetAllUser(ctx context.Context) ([]*domain.User, error) {
	return u.repository.GetUsers(ctx)
}
