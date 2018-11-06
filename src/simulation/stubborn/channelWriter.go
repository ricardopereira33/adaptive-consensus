package stubborn

// SSendAll is the method that sends a message to all participants
func (c *Channel) SSendAll(message []byte) {
    for id := range c.Peers {
        if id != c.GetPeerID() {
            c.SSend(id, message)
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
    c.sendDirect(idDest, message)
}

func (c *Channel) sendDirect(idDest int, message *Package) {
    dest, _     := c.Peers[idDest]
    coefficient := int(1/c.consInfo.PercentMiss)
    percentMsg  := float64(1/(c.consInfo.NMessages % coefficient))*100


    if  percentMsg != c.consInfo.PercentMiss {
        dest <-message
        c.consInfo.NMessages++
    }
}