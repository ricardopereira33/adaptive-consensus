package stubborn

import (
	"log"
	"encoding/json"
)

func (c *Channel) SSendAll(message []byte) {
	if Debug { log.Println("sSendAll a message") }
	
	for id, _ := range c.Peers {
		c.SSend(id, message)
	}
}


// SSend is the method that sends the messages through the channel
func (c *Channel) SSend(idDest int, message []byte) {
	if Debug { log.Println("sSend a message") }	
	
	packageMsg := newPackage(c.PeerID, message, false)
	isToSend   := c.delta0(idDest, packageMsg)
	c.OutBuffer.insertElem(idDest, packageMsg)

	if isToSend { 
		c.send(idDest) 
	}
}

func (c *Channel) send(idDest int) {
	message 	  := c.OutBuffer.getElem(idDest)
	peerAddr, prs := c.Peers[idDest]
	jsonMsg, err  := json.Marshal(message)
	checkError(err)

	if prs {
		c.Connection.WriteTo(jsonMsg, peerAddr)
	} else{
		log.Println("The address doesn't exist!")
	}
}

func (c *Channel) sendDirect(idDest int, message *Package) {
	peerAddr 	 := c.Peers[idDest]
	jsonMsg, err := json.Marshal(message)
	checkError(err)
	
	c.Connection.WriteTo(jsonMsg, peerAddr)
}