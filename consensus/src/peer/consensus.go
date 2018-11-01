package main

import (
	"strconv"
	"stubborn"
	"log"
	msg "message"
)

func consensus(channel stubborn.StubChannel, value string) {
	voters   = make(map[int] bool)
	round    = 1 			  
	phase    = 1 
	estimate = msg.NewEstimate(value, peerID)
	
	if debug { 
		log.Println("CoordID: " + strconv.Itoa(((round % nParticipants) + 1))) 
	}

	if peerID == ((round % nParticipants) + 1) {
		voters[peerID]  = true
		estimate.PeerID = peerID

		message := msg.NewMessage(peerID, round, phase, nParticipants, voters, estimate)
		data 	:= message.MessageToBytes()

		channel.SSendAll(data)
	}
}