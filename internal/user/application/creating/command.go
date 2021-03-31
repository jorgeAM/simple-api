package creating

import "github.com/jorgeAM/simple-api/kit/command"

const CreateNewUserCommandType command.Type = "command.creating.user"

type CreateNewUserCommand struct {
	id        string
	username  string
	firstName string
	lastName  string
}

func (cmd CreateNewUserCommand) Type() command.Type {
	return CreateNewUserCommandType
}

func NewCreateNewUserComand(id, username, firstName, lastName string) CreateNewUserCommand {
	return CreateNewUserCommand{
		id:        id,
		username:  username,
		firstName: firstName,
		lastName:  lastName,
	}
}
