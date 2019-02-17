package stubborn

// SReceive is the method that receives the messages through the channel
func (channel *Channel) SReceive() (pack *Package) {
	for {

		pack := <-channel.InputBuffer
		channel.Metrics.incrementMessagesReceived(pack.ID)
		if pack.IsACK {
			oldPack := channel.OutputBuffer.GetElement(pack.ID)
			oldPack.Arrived = true
			channel.OutputBuffer.InsertElement(pack.ID, oldPack)
		} else {
			go channel.ack(pack.ID)
			return pack
		}
	}
}

func (channel *Channel) ack(id int) {
	pack := newPackage(channel.PeerID, nil, true)

	channel.sendDirect(id, pack)
}
