package stubborn

// Receive is the method that receives the messages through the channel
func (channel *Channel) Receive() (pack *Package) {
	for {
		pack := <-channel.inputBuffer

		if pack.Suspicious {
			channel.suspected(pack.SuspiciousID)
		} else {
			// channel.metrics.incrementMessagesReceived(pack.ID)

			if pack.Ack {
				oldPack := channel.outputBuffer.GetElement(pack.ID)
				oldPack.Arrived = true
				channel.outputBuffer.InsertElement(pack.ID, oldPack)
			} else {
				go channel.ack(pack.ID)
				return pack
			}
		}
	}
}

func (channel *Channel) ack(id int) {
	pack := newPackage(channel.peerID, nil, true)

	channel.sendDirect(id, pack)
}
