package message

import (
    "strconv"
    "log"
    "encoding/json"
    stb "simulation/stubborn"
    con "simulation/consensusinfo"
    ex "simulation/exception"
)

// Message sent out to the server
type Message struct {
    PeerID   int  
    Round    int  
    Phase    int
    Voters   map[int] bool 
    Estimate *con.Estimate
}

// NewMessage creates a new message using the parameters passed in and returns it
func NewMessage(peerID int, round int, phase int, voters map[int] bool, estimate *con.Estimate) (message *Message){
    message          = new(Message)
    message.PeerID   = peerID
    message.Round    = round
    message.Phase    = phase
    message.Voters   = voters
    message.Estimate = estimate

    return
}

/*** Exported Methods ***/

// MessageToBytes convert a Message to Bytes
func (message *Message) MessageToBytes() (data []byte) {
    data, err := json.Marshal(message)
    ex.CheckError(err)

    return 
}

// PrintMessage prints the message
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
func PackageToMessage(pack *stb.Package) (message *Message) {
    data := pack.GetData()
    err  := json.Unmarshal(data, &message)
    ex.CheckError(err)

    return 
}

/*** Axiliary Functions ***/

func toString(integer int) string { 
    return strconv.Itoa(integer)
}