package creating

import (
	"context"
	"errors"

	"github.com/jorgeAM/simple-api/kit/command"
)

type CreateNewUserHandler struct {
	service UserCreatingService
}

func (h CreateNewUserHandler) Handle(ctx context.Context, cmd command.Command) error {
	createNewUserCmd, ok := cmd.(CreateNewUserComand)

	if !ok {
		return errors.New("unexpected command")
	}

	return h.service.CreateNewUser(ctx, createNewUserCmd.id, createNewUserCmd.username, createNewUserCmd.firstName, createNewUserCmd.lastName)
}

func NewCreateNewUserHandler(service UserCreatingService) CreateNewUserHandler {
	return CreateNewUserHandler{
		service: service,
	}
}
