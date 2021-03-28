package domain

import (
	"errors"
	"fmt"

	"github.com/google/uuid"
)

var ErrInvalidUserID = errors.New("invalid user ID")

type UserID struct {
	value string
}

func NewUserID(value string) (UserID, error) {
	v, err := uuid.Parse(value)

	if err != nil {
		return UserID{}, fmt.Errorf("%w: %s", ErrInvalidUserID, value)
	}

	return UserID{value: v.String()}, nil
}

func (id UserID) String() string {
	return id.value
}
