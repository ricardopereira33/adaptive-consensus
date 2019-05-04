package stubborn

// Receive is the method that receives the messages through the channel
func (channel *Channel) Receive() (pack *Package) {
	for {
		pack := <- channel.inputBuffer

		if pack.Suspicious {
			channel.suspected(pack.SuspiciousID)
		} else {
			channel.metrics.incrementMessagesReceived(pack.ID)

			if channel.senderVoted(pack.ID, pack){
				oldPack := channel.outputBuffer.GetElement(pack.ID)

				if oldPack != nil {
					oldPack.Arrived = true

					channel.outputBuffer.InsertElement(pack.ID, oldPack)
				}
			}

			return pack
		}
	}
}
