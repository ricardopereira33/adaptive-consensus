package stubborn

import (
	"log"
	"net"
	"encoding/json"
)

const debug = true

// Channel to send and receive messages between peers
type Channel struct {
	peers 		map[int] *net.UDPAddr
	outBuffer	Buffer
	inBuffer	Buffer
}

func newChannel(peers map[int] *net.UDPAddr) (channel *Channel) {
	channel = new(Channel)

	channel.peers	  = peers
	channel.outBuffer = newBuffer(len(peers))
	channel.inBuffer  = newBuffer(len(peers))
	return
}

func (c Channel) sSend(idDest int, message []byte) bool {
	if debug { log.Println("sSend a message") }	
	
	res := c.outBuffer.insertElem(idDest, message)
	return res
}

func (c Channel) send() {

}

func (c Channel) sReceive() []byte {
	if debug { log.Println("sReceive a message") }	

	message := c.inBuffer.getElem()
	return message
}

func (c Channel) receive() []byte {

}