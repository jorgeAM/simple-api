package domain

type Repository interface {
	NewUser(user *User) error
	GetUsers() ([]*User, error)
	GetUser(id string) (*User, error)
	UpdateUser(user *User) (*User, error)
	DeleteUser(id string) error
}
