package exception

import (
    "log"
    "fmt"
    "os"
    "errors"
)

// HandleErr handles the errors
func HandleErr(err error) {
    if err != nil {
        log.Println("No one in the chat yet")
    }
}

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