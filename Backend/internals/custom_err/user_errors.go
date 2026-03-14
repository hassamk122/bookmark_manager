package customerr

import "errors"

var (
	ErrEmailTaken         = errors.New("Email already taken")
	ErrUserNotFound       = errors.New("User not found")
	ErrInvalidCredentials = errors.New("Invalid Credentials")
)
