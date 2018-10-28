package stubborn

import (
	"log"
	"fmt"
	"os"
)

//handles errors
func handleErr(err error) {
	if err != nil {
		log.Println("No one in the chat yet")
	}
}

//check errors
func checkError(err error) bool{
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1) 
	}
	return true 
}

func checkUDPError(err error) bool{
	if err != nil {
		log.Println("Conection closed.")
		return true 
	}
	return false
}