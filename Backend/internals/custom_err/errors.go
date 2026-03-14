package customerr

import "errors"

var (
	ErrInvalidRequestPayload = errors.New("Invalid request payload")
	ErrSomethingWentWrong    = errors.New("Something went Wrong")
)
