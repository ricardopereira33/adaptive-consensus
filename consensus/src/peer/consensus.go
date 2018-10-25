package main

import (
	"strconv"
	"stubborn"
	"log"
)

func consensus(channel stubborn.StubChannel, value string) {
	voters   = make(map[int] bool)
	round    = 1 			  
	phase    = 1 
	estimate = newEstimate(value, peerID)
	
	if debug { 
		log.Println("CoordID: " + strconv.Itoa(((round % nParticipants) + 1))) 
	}

	if peerID == ((round % nParticipants) + 1) {
		voters[peerID]  = true
		estimate.peerID = peerID

		message := newMessage(peerID, round, phase, voters, estimate)
		data 	:= messageToBytes(message)

		channel.SSendAll(data)
	}
}