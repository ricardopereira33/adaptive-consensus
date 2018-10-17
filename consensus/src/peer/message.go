package main

import (
	"log"
	"encoding/json"
)

// Message sent out to the server
type Message struct {
	Username  string   // My username
	Info      string   // Message
}

// Creates a new message using the parameters passed in and returns it
func newMessage(username string, info string) (message *Message){
	message 		 = new(Message)
	message.Username = username
	message.Info     = info

	return
}

func messageToBytes(message *Message) (data []byte) {
	data, err := json.Marshal(message)
	checkError(err, false)

	return 
}

// Print the message
func printMessage(msg Message){
	log.Print("[ " + msg.Username + " ] ")
	log.Println(" " + msg.Info)
}