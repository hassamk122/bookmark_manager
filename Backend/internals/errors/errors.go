package errors

import "errors"

var (
	ErrInvalidRequestPayload = errors.New("Invalid request payload").Error()
)
