package postdude

import "fmt"

type Error struct {
	Message string
	Code    int
}

// NewError initialize an instance of 'Error' struct.
func NewError(message string, code int, a ...any) *Error {
	return &Error{Message: fmt.Sprintf(message, a), Code: code}
}
