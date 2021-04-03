package finding

import (
	"context"
	"errors"

	"github.com/jorgeAM/simple-api/internal/user/domain"
	"github.com/jorgeAM/simple-api/kit/query"
)

type FindUserByIDHandler struct {
	service UserRetrieveOneService
}

func (h FindUserByIDHandler) Handle(ctx context.Context, cmd query.Command) (interface{}, error) {
	findUserByIDQuery, ok := cmd.(FindUserByIDQuery)

	if !ok {
		return nil, errors.New("unexpected command")
	}

	userID, err := domain.NewUserID(findUserByIDQuery.id)

	if err != nil {
		return nil, err
	}

	return h.service.FindUserByID(ctx, userID)
}

func NewFindUserByIDHandler(service UserRetrieveOneService) FindUserByIDHandler {
	return FindUserByIDHandler{
		service: service,
	}
}
