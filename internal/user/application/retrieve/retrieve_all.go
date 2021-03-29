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

func (u *UserRetrieveAllService) GetAllUser(ctx context.Context) ([]*userResponse, error) {
	users, err := u.repository.GetUsers(ctx)

	if err != nil {
		return nil, err
	}

	var usersRes []*userResponse

	for _, user := range users {
		userRes := NewUserResponseFromAggregate(user)
		usersRes = append(usersRes, userRes)
	}

	return usersRes, nil
}
