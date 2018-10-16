package stubborn

import (
	"log"
	"encoding/json"
)

func (c Channel) sSend(idDest int, message []byte) {
	if Debug { log.Println("sSend a message") }	
	
	packageMsg := newPackage(idDest, message, false)
	isToSend   := c.delta0Func(idDest, packageMsg)
	c.outBuffer.insertElem(packageMsg)

	if isToSend { 
		c.send(idDest) 
	}
}

func (c Channel) send(idDest int) {
	message 	 := c.outBuffer.getElem(idDest)
	peerAddr	 := c.peers[idDest]
	jsonMsg, err := json.Marshal(message)
	checkError(err, false)

	c.connection.WriteTo(jsonMsg, peerAddr)
}

func (c Channel) sendDirect(idDest int, message *Package) {
	peerAddr := c.peers[idDest]
	jsonMsg, err := json.Marshal(message)
	checkError(err, false)
	
	c.connection.WriteTo(jsonMsg, peerAddr)
}