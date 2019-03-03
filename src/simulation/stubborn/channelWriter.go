package stubborn

import (
	"strconv"
	"time"
)

// SendAll is the method that sends a message to all participants
func (channel *Channel) SendAll(message []byte) {
	for id := 1; id <= channel.numberParticipants; id++ {
		if id != channel.peerID {
			go channel.Send(id, message)
		}
	}
}

// SendSuspicion is the method that notify for a possible suspicion
func (channel *Channel) SendSuspicion(peerID int, targetPeerID int) {
    message := newSuspect(targetPeerID, true)
    channel.sendDirect(peerID, message)
}

// Send is the method that sends the messages through the channel
func (channel *Channel) Send(idDestination int, message []byte) {
	packageMsg := newPackage(channel.peerID, message, false)
	isToSend := channel.delta0(idDestination, packageMsg)
	channel.outputBuffer.InsertElement(idDestination, packageMsg)

	if isToSend {
        channel.sendMessage(idDestination)
	} else {
        channel.metrics.initialDelay(idDestination)
    }
}

func (channel *Channel) sendMessage(idDestination int) {
	message := channel.outputBuffer.GetElement(idDestination)
	go channel.sendDirect(idDestination, message)
}

func (channel *Channel) sendDirect(idDestination int, message *Package) {
	coefficient := int(1 / channel.percentMiss) * 100
	missingMsg := (channel.metrics.getMessagesSent(idDestination) + 1) % coefficient

	if missingMsg != 0 {
		// Simulate the message delay
		time.Sleep(time.Second)

		channel.sendToBuffer(idDestination, message)
		channel.metrics.incrementMessagesSent(idDestination)
	}
}

func (channel *Channel) sendToBuffer(id int, pack *Package) {
	value, present := channel.peers.Get(strconv.Itoa(id))

	if present {
		peerChannel := value.(chan *Package)
		peerChannel <- pack
	}
}
