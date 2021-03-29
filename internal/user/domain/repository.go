package domain

import "context"

type Repository interface {
	NewUser(ctx context.Context, user *User) error
	GetUsers(ctx context.Context) ([]*User, error)
	GetUser(ctx context.Context, id string) (*User, error)
	DeleteUser(ctx context.Context, id string) error
}
