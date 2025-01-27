package consensus

import (
	"encoding/json"
	"strconv"

	cmap "github.com/orcaman/concurrent-map"
	ex "simulation/exception"
	stb "simulation/stubborn"
)

// Message sent out to the server
type Message struct {
	PeerID   int
	Round    int
	Phase    int
	Voters   map[int]int
	Estimate *Estimate
}

// NewMessage creates a new message using the parameters passed in and returns it
func NewMessage(peerID int, round int, phase int, voters cmap.ConcurrentMap, estimate *Estimate) (message *Message) {
	message = new(Message)
	message.PeerID = peerID
	message.Round = round
	message.Phase = phase
	message.Voters = convertCmapToMap(voters)
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

// PackageToMessage convert a Package to a Message
func PackageToMessage(pack *stb.Package) (message *Message) {
	data := pack.GetData()
	err := json.Unmarshal(data, &message)
	ex.CheckError(err)

	return
}

// SenderVoted checks if the sender voted
func SenderVoted(senderID int, pack *stb.Package) bool {
    message := PackageToMessage(pack)

    if message.Voters[senderID] == 1 {
        return true
    }

    return false
}

/*** Axiliary Functions ***/

func toString(integer int) string {
	return strconv.Itoa(integer)
}
