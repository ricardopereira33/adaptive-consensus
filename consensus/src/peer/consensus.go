package main

import (
	"stubborn"
	"net"
)

func consensus(channel stubborn.StubChannel, value string) {
	voters   = make(map[int] *net.UDPAddr)
	round    = 1 			  
	phase    = 1 	
	
	if peerID == ((round % nParticipants) + 1) {
		
	}
}