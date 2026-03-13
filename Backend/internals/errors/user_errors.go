package errors

import "errors"

var (
	ErrEmailTaken         = errors.New("Email already taken").Error()
	ErrUserNotFound       = errors.New("User not found").Error()
	ErrInvalidCredentials = errors.New("Invalid Credentials").Error()
)
