package domain

type User struct {
	ID        UserID
	Username  Username
	FirstName FirstName
	LastName  LastName
}

func NewUser(id, username, firstName, lastName string) (*User, error) {
	idVO, err := NewUserID(id)

	if err != nil {
		return nil, err
	}

	usernameVO, err := NewUsername(username)

	if err != nil {
		return nil, err
	}

	firstNameVO, err := NewFirstName(firstName)

	if err != nil {
		return nil, err
	}

	lastNameVO, err := NewLastName(lastName)

	if err != nil {
		return nil, err
	}

	return &User{
		ID:        idVO,
		Username:  usernameVO,
		FirstName: firstNameVO,
		LastName:  lastNameVO,
	}, nil
}
