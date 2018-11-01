package message

import (
	"strconv"
	"log"
	"encoding/json"
	"stubborn"
	ex "exception"
)

// Message sent out to the server
type Message struct {
	PeerID	 int  
	Round    int  
	Phase 	 int
	NPeers   int
	Voters   map[int] bool 
	Estimate *Estimate
}

// Estimate is the estimate value for a consensus
type Estimate struct {
	Value  string
	PeerID int
}

// Creates a new message using the parameters passed in and returns it
func NewMessage(peerID int, round int, phase int, npeers int, voters map[int] bool, estimate *Estimate) (message *Message){
	message 	   	 = new(Message)
	message.PeerID 	 = peerID
	message.Round  	 = round
	message.Phase  	 = phase
	message.NPeers   = npeers
	message.Voters 	 = voters
	message.Estimate = estimate

	return
}

func NewEstimate(value string, id int) (estimate *Estimate){
	estimate 		= new(Estimate)
	estimate.Value  = value
	estimate.PeerID = id

	return
}

func (message *Message) MessageToBytes() (data []byte) {
	data, err := json.Marshal(message)
	ex.CheckError(err)

	return 
}

// Print the message
func (msg *Message) PrintMessage() {
	log.Println("----------------------")
	log.Println("[ " + toString(msg.PeerID) + " ] Message")
	log.Println("R: " + toString(msg.Round) + "| Ph: " + toString(msg.Phase))
	log.Print("Voters: ")
	log.Println(msg.Voters)
	log.Print("Estimate: ") 
	log.Println(msg.Estimate)
	log.Println("----------------------")
}

func PackageToMessage(pack *stubborn.Package) (message *Message) {
	data := pack.GetData()
	err  := json.Unmarshal(data, &message)
	ex.CheckError(err)

	return 
}

// Axiliary Functions 

func toString(integer int) string { 
	return strconv.Itoa(integer)
}