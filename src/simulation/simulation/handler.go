package main

import (
	stb "simulation/stubborn"
    msg "simulation/message"
    con "simulation/consensusinfo"
)

func handleMessages(channel stb.StubChannel) {
	for {
		pack 	      := channel.SReceive()
		message       := msg.PackageToMessage(pack)
        c             := channel.GetConsensusInfo()
        nParticipants := channel.GetNParticipants()
        peerID        := channel.GetPeerID()

		if len(c.Voters) <= nParticipants/2 {
			checkRound(message, c)
			existsNewVoters := (c.Round == message.Round) && containsNewVoters(message.Voters, c)
			isMajority 		:= (c.Phase == 1) && (len(message.Voters) > nParticipants/2)
			
			if existsNewVoters || isMajority {
				message.Voters[peerID] = true
				c.Voters = union(c.Voters, message.Voters)
				
				if message.Estimate.PeerID == ((c.Round % nParticipants) + 1){
					c.Estimate = message.Estimate
                }

				message := msg.NewMessage(peerID, c.Round, c.Phase, c.Voters, c.Estimate)
				data 	:= message.MessageToBytes()
				channel.SSendAll(data)
			}
		} else if checkPhase(message, c) { 
			break 
		}
	}
}

func checkRound(message *msg.Message, cInfo *con.ConsensusInfo) {
	if cInfo.Round < message.Round {
		cInfo.Estimate = message.Estimate
		cInfo.Round    = message.Round
	    cInfo.Phase    = message.Phase
		cInfo.Voters   = make(map[int] bool)
	} else if cInfo.Round == message.Round && cInfo.Phase < message.Phase {
		cInfo.Phase  = message.Phase
		cInfo.Voters = make(map[int] bool)
	}
}

func checkPhase(message *msg.Message, cInfo *con.ConsensusInfo) bool {
	if cInfo.Phase == 1 {
		cInfo.CDecision = message.Estimate.Value
		return true
	} 
	
	cInfo.Round++
	cInfo.Phase  = 1
	cInfo.Voters = make(map[int] bool)
	return false
}


// Auxiliary Funtions 

func containsNewVoters(senderVoters map[int] bool, cInfo *con.ConsensusInfo) bool {
	for id := range senderVoters {
		_, isPresent := cInfo.Voters[id]
		if !isPresent {
			return true
		}  
	}

	return false
}

func union(hash1 map[int] bool, hash2 map[int] bool) map[int] bool{
	for id := range hash2 {
		hash1[id] = true
	}

	return hash1
}