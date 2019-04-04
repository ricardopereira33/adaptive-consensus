package stubborn

import (
	"math/rand"
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
	}
	// else {
	//     channel.metrics.initialDelay(idDestination)
	// }
}

func (channel *Channel) sendMessage(idDestination int) {
	message := channel.outputBuffer.GetElement(idDestination)
	go channel.sendDirect(idDestination, message)
}

func (channel *Channel) sendDirect(idDestination int, message *Package) {
	if successMessage(channel.percentageMiss) {
        //Bandwidth
        channel.limiter.Take()

		channel.sendToBuffer(idDestination, message)
		channel.metrics.incrementMessagesSent(idDestination)
	}
}

func (channel *Channel) sendToBuffer(id int, pack *Package) {
	value, present := channel.peers.Get(strconv.Itoa(id))

	if present {
        // Latency
        time.Sleep(channel.latency)

		peerChannel := value.(chan *Package)
		peerChannel <- pack
	}
}

func successMessage(percentageMiss float64) bool {
	randomValue := rand.Float64()
	percentage := percentageMiss / 100.0

	if randomValue < percentage {
		return false
	}

	return true
}
