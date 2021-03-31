package creating

import "github.com/jorgeAM/simple-api/kit/command"

const CreateNewUserComandType command.Type = "command.creating.user"

type CreateNewUserComand struct {
	id        string
	username  string
	firstName string
	lastName  string
}

func (cmd CreateNewUserComand) Type() command.Type {
	return CreateNewUserComandType
}

func NewCreateNewUserComand(id, username, firstName, lastName string) CreateNewUserComand {
	return CreateNewUserComand{
		id:        id,
		username:  username,
		firstName: firstName,
		lastName:  lastName,
	}
}
