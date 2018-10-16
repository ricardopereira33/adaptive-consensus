package stubborn

import (
	"log"
)

// SReceive is the method that receives the messages through the channel
func (c *Channel) SReceive() (pack *Package){
    if Debug { log.Println("sReceive a message") }	

    select{
        case pack := <-c.inBuffer :
            c.ack(pack.ID)
            return pack
    }
}

// receive is the loop that receive the message and send through the buffer(pipe)
// for the SReceive method receive the messages
func (c *Channel) receive() (int, []byte) {
	if Debug { log.Println("receive") }

    buffer  := make([]byte, MaxDatagramSize)
    channel := c.connection

    for {
        nBytes, _, err := channel.ReadFromUDP(buffer)

        if checkError(err, true){
            data := buffer[:nBytes]
            pack := bytesToPackage(data)

            if pack.IsACK {
                oldPack := c.outBuffer.getElem(pack.ID)
                oldPack.Arrived = true
                c.outBuffer.insertElem(pack.ID, oldPack)
            } else {
                go func() { c.inBuffer<- pack }()
            }
        }
    }
}

func (c *Channel) ack(id int){
    pack := newPackage(c.peerID, nil, true)

    c.sendDirect(id, pack)
}