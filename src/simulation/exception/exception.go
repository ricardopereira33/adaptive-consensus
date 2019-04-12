package exception

import (
	"errors"
	"fmt"
	"log"
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

// CheckUDPError checks the UDP errors
func CheckUDPError(err error) bool {
	if err != nil {
		log.Println("Conection closed.")
		return true
	}
	return false
}

// NewError creates a new error
func NewError(message string) error {
	return errors.New(message)
}
