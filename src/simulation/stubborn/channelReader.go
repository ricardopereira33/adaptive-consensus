package stubborn

// Receive is the method that receives the messages through the channel
func (channel *Channel) Receive() (pack *Package) {
	pack = <- channel.peerChannel.inputBuffer

	channel.metrics.incrementMessagesReceived(pack.ID)

	if channel.senderVoted(pack.ID, pack){
		oldPack := channel.outputBuffer.GetElement(pack.ID)

		if oldPack != nil {
			oldPack.Arrived = true

			channel.outputBuffer.InsertElement(pack.ID, oldPack)
		}
	}

	return
}

// receiveSuspicious is the method that handles suspicious of fails
func (channel *Channel) receiveSuspicious() {
	for {
		pack := <- channel.peerChannel.inputSuspicions

		channel.suspected(pack.SuspiciousID)

		println("> Suspicious <")
	}
}