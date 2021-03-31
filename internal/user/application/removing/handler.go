package removing

import (
	"context"
	"errors"

	"github.com/jorgeAM/simple-api/kit/command"
)

type RemoveUserHandler struct {
	service UserRemovingService
}

func (h RemoveUserHandler) Handle(ctx context.Context, cmd command.Command) error {
	removeUserCmd, ok := cmd.(RemoveUserCommand)

	if !ok {
		return errors.New("unexpected command")
	}

	return h.service.RemoveUserByID(ctx, removeUserCmd.id)
}

func NewRemoveUserHandler(service UserRemovingService) RemoveUserHandler {
	return RemoveUserHandler{
		service: service,
	}
}
