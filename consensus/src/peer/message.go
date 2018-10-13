package main

import "fmt"

// Message sent out to the server
type Message struct {
	Username  string   // My username
	Info      string   // Message
}

// Creates a new message using the parameters passed in and returns it
func createMessage(username string, info string,) (msg *Message){
	msg = new(Message)
	
	msg.Username = username
	msg.Info     = info
	return
}

// Print the message
func printMessage(msg Message){
	fmt.Print("[ " + msg.Username + " ]\t")
	fmt.Println(msg.Info)
}