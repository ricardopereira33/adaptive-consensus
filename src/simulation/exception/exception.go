package exception

import (
	"errors"
	"fmt"
	"os"
)

// CheckError checks the errors
func CheckError(err error) bool {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
	return true
}

// NewError creates a new error
func NewError(message string) error {
	return errors.New(message)
}
