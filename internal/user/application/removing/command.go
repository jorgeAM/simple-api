package removing

import "github.com/jorgeAM/simple-api/kit/command"

const RemoveUserCommandType command.Type = "command.removing.user"

type RemoveUserCommand struct {
	id string
}

func (c RemoveUserCommand) Type() command.Type {
	return RemoveUserCommandType
}

func NewRemoveUserCommand(id string) RemoveUserCommand {
	return RemoveUserCommand{
		id: id,
	}
}
