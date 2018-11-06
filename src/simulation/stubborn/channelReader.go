package stubborn

// SReceive is the method that receives the messages through the channel
func (c *Channel) SReceive() (pack *Package){    
    select{
        case pack := <-c.InBuffer :
            c.ack(pack.ID)
            if pack.IsACK {
                oldPack        := c.OutBuffer.GetElem(pack.ID)
                oldPack.Arrived = true
                c.OutBuffer.InsertElem(pack.ID, oldPack)
            } else {
                go func() { c.InBuffer<- pack }()
            }
            
            return pack
    }
}

func (c *Channel) ack(id int){
    pack := newPackage(c.PeerID, nil, true)

    c.sendDirect(id, pack)
}