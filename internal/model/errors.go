package model

import (
	"errors"
	"fmt"
)

var (
	ErrNotFound     = errors.New("resource not found")
	ErrUnauthorized = errors.New("unauthorized access")
	ErrInvalidInput = errors.New("invalid input data")
)

type ValidateError struct {
	Message string
	Err     error
}

func (ve *ValidateError) Error() string {
	if ve.Err != nil {
		return fmt.Sprintf("%s:%v", ve.Message, ve.Err)
	}
	return ve.Message
}
