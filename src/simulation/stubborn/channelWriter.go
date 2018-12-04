package stubborn

import (
	"time"
	"strconv"
)

// SSendAll is the method that sends a message to all participants
func (c *Channel) SSendAll(message []byte) {
    for id := 1; id <= c.NParticipants; id++ {
        if id != c.GetPeerID() {
            go c.SSend(id, message)
        }
    }
}

// SSend is the method that sends the messages through the channel
func (c *Channel) SSend(idDest int, message []byte) {
    packageMsg := newPackage(c.PeerID, message, false)
    isToSend   := c.delta0(idDest, packageMsg)
    c.OutBuffer.InsertElem(idDest, packageMsg)

    if isToSend {
        c.send(idDest) 
    }
}

func (c *Channel) send(idDest int) {
    message := c.OutBuffer.GetElem(idDest)
    go c.sendDirect(idDest, message)
}

func (c *Channel) sendDirect(idDest int, message *Package) {
    coefficient := int(1/c.consInfo.PercentMiss) * 100
    missingMsg  := (c.Metrics.GetSendedMsg(idDest) + 1) % coefficient

    if missingMsg != 0 {
        // Simulate the message delay
        time.Sleep(time.Second)

        c.sendToBuffer(idDest, message)
        c.Metrics.IncSendedMsg(idDest)
    }
}

func (c *Channel) sendToBuffer(id int, pack *Package) {
    value, prs := c.Peers.Get(strconv.Itoa(id))
    
    if prs {
        channel := value.(chan *Package)
        channel <- pack 
    }
}