package retrieve

import (
	"context"

	"github.com/jorgeAM/simple-api/kit/query"
)

type GetAllUsersHandler struct {
	service UserRetrieveAllService
}

func (h GetAllUsersHandler) Handle(ctx context.Context, cmd query.Command) (interface{}, error) {
	return h.service.GetAllUser(ctx)
}

func NewGetAllUsersHandler(service UserRetrieveAllService) GetAllUsersHandler {
	return GetAllUsersHandler{
		service: service,
	}
}
