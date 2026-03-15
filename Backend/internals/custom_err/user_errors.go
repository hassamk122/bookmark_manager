package customerr

import "errors"

var (
	EmailTaken         = errors.New("Email already taken")
	UserNotFound       = errors.New("User not found")
	InvalidCredentials = errors.New("Invalid Credentials")
)
