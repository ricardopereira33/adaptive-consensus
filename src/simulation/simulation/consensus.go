package main

import (
	"strconv"
	"log"
	stb "simulation/stubborn"
    msg "simulation/message"
    con "simulation/consensusinfo"
)

func consensus(channel stb.StubChannel, value string) {
    c         := channel.GetConsensusInfo()
    c.Voters   = make(map[int] bool)
	c.Round    = 1 			  
	c.Phase    = 1 
    c.Estimate = con.NewEstimate(value, channel.GetPeerID())
    peerID    := channel.GetPeerID()
	
	if debug { 
		log.Println("CoordID: " + strconv.Itoa(((c.Round % channel.GetNParticipants()) + 1))) 
	}
	
	coord := (c.Round % channel.GetNParticipants()) + 1
	if peerID == channel.GetCoordID(){
		c.Voters[peerID]  = true
		c.Estimate.PeerID = peerID

		message := msg.NewMessage(peerID, c.Round, c.Phase, c.Voters, c.Estimate)
		data 	:= message.MessageToBytes()

		channel.SetCoordinator(coord)
		channel.SSendAll(data)
	}
}