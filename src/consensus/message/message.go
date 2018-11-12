package message

import (
    "strconv"
    "log"
    "encoding/json"
    "consensus/stubborn"
    ex "consensus/exception"
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
    PeerID int
}

// NewMessage creates a new message using the parameters passed in and returns it
func NewMessage(peerID int, round int, phase int, voters map[int] bool, estimate *Estimate) (message *Message){
    message 	   	 = new(Message)
    message.PeerID 	 = peerID
    message.Round  	 = round
    message.Phase  	 = phase
    message.Voters 	 = voters
    message.Estimate = estimate

    return
}
// NewEstimate creates a new estimate value
func NewEstimate(value string, id int) (estimate *Estimate){
    estimate 		= new(Estimate)
    estimate.Value  = value
    estimate.PeerID = id

    return
}

// MessageToBytes convert a Message to Bytes
func (message *Message) MessageToBytes() (data []byte) {
    data, err := json.Marshal(message)
    ex.CheckError(err)

    return 
}

// PrintMessage print the message
func (message *Message) PrintMessage() {
    log.Println("----------------------")
    log.Println("[ " + toString(message.PeerID) + " ] Message")
    log.Println("R: " + toString(message.Round) + "| Ph: " + toString(message.Phase))
    log.Print("Voters: ")
    log.Println(message.Voters)
    log.Print("Estimate: ") 
    log.Println(message.Estimate)
    log.Println("----------------------")
}

// PackageToMessage convert a Package to a Message
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