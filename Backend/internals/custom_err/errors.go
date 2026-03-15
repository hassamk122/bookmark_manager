package customerr

import "errors"

var (
	InvalidRequestPayload = errors.New("Invalid request payload")
	SomethingWentWrong    = errors.New("Something went Wrong")
)
