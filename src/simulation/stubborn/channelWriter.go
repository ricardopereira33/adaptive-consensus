package stubborn

// SSendAll is the method that sends a message to all participants
func (c *Channel) SSendAll(message []byte) {
    for id := range c.Peers {
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
    numberMsg   := (c.consInfo.NMessages+1) % coefficient
    percentMsg  := float64(100)/float64(numberMsg)

    /*
    c.print("nMsg: " + strconv.Itoa(numberMsg))
    c.print("percMsg:" + strconv.FormatFloat(percentMsg,'f',-1,64))
    */

    if percentMsg != c.consInfo.PercentMiss {
        c.sendToBuffer(idDest, message)
        c.consInfo.NMessages++
    }
}