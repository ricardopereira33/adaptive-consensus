package stubborn

import (
    "log"
	ex "consensus/exception"
)

// SReceive is the method that receives the messages through the channel
func (c *Channel) SReceive() (pack *Package){
    select{
        case pack := <-c.InBuffer :
            c.ack(pack.ID)
            return pack
    }
}

// receive is the loop that receive the message and send through the buffer(pipe)
// for the SReceive method receive the messages
func (c *Channel) receive() {
	if Debug { log.Println("Listener start...") }

    buffer  := make([]byte, MaxDatagramSize)
    channel := c.Connection

    for {
        nBytes, _, err := channel.ReadFromUDP(buffer)
        if ex.CheckUDPError(err) {
            break
        }
        
        data := buffer[:nBytes]
        pack := bytesToPackage(data)

        if pack.IsACK {
            oldPack := c.OutBuffer.getElem(pack.ID)
            oldPack.Arrived = true
            c.OutBuffer.insertElem(pack.ID, oldPack)
        } else {
            go func() { c.InBuffer<- pack }()
        }
    }
}

func (c *Channel) ack(id int){
    pack := newPackage(c.PeerID, nil, true)

    c.sendDirect(id, pack)
}