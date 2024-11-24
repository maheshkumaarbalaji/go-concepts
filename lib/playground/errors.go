package playground

import (
	"fmt"
)

type CustomError struct {
	Lineno int
	Message string
}

func (ce *CustomError) Error() string {
	return fmt.Sprintf("Error at line %d: %s", ce.Lineno, ce.Message)
}

func SampleFunction() error {
	cusError := CustomError{
		Lineno: 1,
		Message: "This is a test message",
	}

	return &cusError
}
