package domain

import (
	"errors"
	"fmt"
)

var ErrInvalidUsername = errors.New("invalid username")

type Username struct {
	value string
}

func NewUsername(value string) (Username, error) {
	if len(value) == 0 {
		return Username{}, fmt.Errorf("%w: %s", ErrInvalidUsername, value)
	}

	return Username{value: value}, nil
}

func (u Username) String() string {
	return u.value
}
