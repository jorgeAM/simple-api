package domain

import (
	"errors"
	"fmt"
)

var ErrInvalidFirstName = errors.New("invalid first name")

type FirstName struct {
	value string
}

func NewFirstName(value string) (FirstName, error) {
	if len(value) == 0 {
		return FirstName{}, fmt.Errorf("%w: %s", ErrInvalidFirstName, value)
	}

	return FirstName{value: value}, nil
}

func (u FirstName) String() string {
	return u.value
}
