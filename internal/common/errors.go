package common

import (
	"errors"
	"fmt"
)

var (
	ErrInvalidArgument = errors.New("invalid argument")
	ErrNotFound        = errors.New("not found")
)

func NewInvalidArgument(message string) error {
	if message == "" {
		return ErrInvalidArgument
	}

	return fmt.Errorf("%w: %s", ErrInvalidArgument, message)
}

func NewNotFound(message string) error {
	if message == "" {
		return ErrNotFound
	}

	return fmt.Errorf("%w: %s", ErrNotFound, message)
}
