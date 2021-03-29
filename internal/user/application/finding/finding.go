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

func (u *UserRetrieveOneService) FindUserByID(ctx context.Context, userID domain.UserID) (*userResponse, error) {
	user, err := u.repository.GetUser(ctx, userID.String())

	if err != nil {
		return nil, err
	}

	userRes := NewUserResponseFromAggregate(user)

	return userRes, nil
}
