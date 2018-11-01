package exception

import (
	"log"
	"fmt"
	"os"
)

//handles errors
func HandleErr(err error) {
	if err != nil {
		log.Println("No one in the chat yet")
	}
}

//check errors
func CheckError(err error) bool {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1) 
	}
	return true 
}

func CheckUDPError(err error) bool {
	if err != nil {
		log.Println("Conection closed.")
		return true 
	}
	return false
}