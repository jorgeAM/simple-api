package domain

import (
	"errors"
	"fmt"
)

var ErrInvalidLastName = errors.New("invalid last name")

type LastName struct {
	value string
}

func NewLastName(value string) (LastName, error) {
	if len(value) == 0 {
		return LastName{}, fmt.Errorf("%w: %s", ErrInvalidLastName, value)
	}

	return LastName{value: value}, nil
}

func (u LastName) String() string {
	return u.value
}
