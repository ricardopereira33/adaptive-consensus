package main

import (
	"stubborn"
)

func consensus(channel stubborn.StubChannel, value string) {
	voters   = make([]int, 1)
	round    = 1 			  
	phase    = 1 
	estimate := newEstimate(value, peerID)
	
	if peerID == ((round % nParticipants) + 1) {
		voters 			= append(voters, peerID)
		estimate.peerID = peerID

		message := newMessage(peerID, round, phase, voters, estimate)
		data 	:= messageToBytes(message)

		channel.SSendAll(data)
	}
}