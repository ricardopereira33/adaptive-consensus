package stubborn

import (
	"strconv"
	"time"
)

// SSendAll is the method that sends a message to all participants
func (channel *Channel) SSendAll(message []byte) {
	for id := 1; id <= channel.NumberParticipants; id++ {
		if id != channel.GetPeerID() {
			go channel.SSend(id, message)
		}
	}
}

// SSend is the method that sends the messages through the channel
func (channel *Channel) SSend(idDestination int, message []byte) {
	packageMsg := newPackage(channel.PeerID, message, false)
	isToSend := channel.delta0(idDestination, packageMsg)
	channel.OutputBuffer.InsertElement(idDestination, packageMsg)

	if isToSend {
		channel.send(idDestination)
	}
}

func (channel *Channel) send(idDestination int) {
	message := channel.OutputBuffer.GetElement(idDestination)
	go channel.sendDirect(idDestination, message)
}

func (channel *Channel) sendDirect(idDestination int, message *Package) {
	coefficient := int(1 / channel.consensusInfo.PercentMiss) * 100
	missingMsg := (channel.Metrics.getMessagesSent(idDestination) + 1) % coefficient

	if missingMsg != 0 {
		// Simulate the message delay
		time.Sleep(time.Second)

		channel.sendToBuffer(idDestination, message)
		channel.Metrics.incrementMessagesSent(idDestination)
	}
}

func (channel *Channel) sendToBuffer(id int, pack *Package) {
	value, present := channel.Peers.Get(strconv.Itoa(id))

	if present {
		peerChannel := value.(chan *Package)
		peerChannel <- pack
	}
}
