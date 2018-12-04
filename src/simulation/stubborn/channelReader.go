package stubborn

/*** Exported methods ***/

// SReceive is the method that receives the messages through the channel
func (c *Channel) SReceive() (pack *Package) {    
    for {
        pack := <-c.InBuffer
        c.Metrics.IncReceivedMsg(pack.ID)
        if pack.IsACK {
            oldPack        := c.OutBuffer.GetElem(pack.ID)
            oldPack.Arrived = true
            c.OutBuffer.InsertElem(pack.ID, oldPack)
        } else {
            go c.ack(pack.ID)
            return pack
        }
    }
}

/*** Auxiliares Functions ***/

func (c *Channel) ack(id int) {
    pack := newPackage(c.PeerID, nil, true)

    c.sendDirect(id, pack)
}