package stubborn

import (
	"log"
	"encoding/json"
)

// SSend is the method that sends the messages through the channel
func (c *Channel) SSend(idDest int, message []byte) {
	if Debug { log.Println("sSend a message") }	
	
	packageMsg := newPackage(c.peerID, message, false)
	isToSend   := c.delta0(idDest, packageMsg)
	c.outBuffer.insertElem(idDest, packageMsg)

	if isToSend { 
		c.send(idDest) 
	}
}

func (c *Channel) send(idDest int) {
	message 	  := c.outBuffer.getElem(idDest)
	peerAddr, prs := c.peers[idDest]
	jsonMsg, err  := json.Marshal(message)
	checkError(err, false)

	if prs {
		c.connection.WriteTo(jsonMsg, peerAddr)
	} else{
		log.Println("The address doesn't exist!")
	}
}

func (c *Channel) sendDirect(idDest int, message *Package) {
	peerAddr 	 := c.peers[idDest]
	jsonMsg, err := json.Marshal(message)
	checkError(err, false)
	
	c.connection.WriteTo(jsonMsg, peerAddr)
}