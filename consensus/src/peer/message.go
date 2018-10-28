package main

import (
	"strconv"
	"log"
	"encoding/json"
	"stubborn"
)

// Message sent out to the server
type Message struct {
	PeerID	 int  
	Round    int  
	Phase 	 int
	Voters   map[int] bool 
	Estimate *Estimate
}

// Estimate is the estimate value for a consensus
type Estimate struct {
	Value  string
	peerID int
}

// Creates a new message using the parameters passed in and returns it
func newMessage(peerID int, round int, phase int, voters map[int] bool, estimate *Estimate) (message *Message){
	message 	   	 = new(Message)
	message.PeerID 	 = peerID
	message.Round  	 = round
	message.Phase  	 = phase
	message.Voters 	 = voters
	message.Estimate = estimate

	return
}

func newEstimate(value string, id int) (estimate *Estimate){
	estimate 		= new(Estimate)
	estimate.Value  = value
	estimate.peerID = id

	return
}

func (message *Message) messageToBytes() (data []byte) {
	data, err := json.Marshal(message)
	checkError(err, false)

	return 
}

// Print the message
func (msg *Message) printMessage() {
	log.Println("----------------------")
	log.Println("[ " + toString(msg.PeerID) + " ] Message")
	log.Println("Round: " + toString(msg.Round))
	log.Println("Phase: " + toString(msg.Phase))
	log.Println("Voters: ")
	log.Println(msg.Voters)
	log.Println("Estimate: ") 
	log.Println(msg.Estimate)
	log.Println("----------------------")
}

func bytesToMessage(pack *stubborn.Package) (message *Message) {
	data := pack.GetData()
	err  := json.Unmarshal(data, &message)
	checkError(err, false)

	return 
}

// Axiliary Functions 

func toString(integer int) string { 
	return strconv.Itoa(integer)
}