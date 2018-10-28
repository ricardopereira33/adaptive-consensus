package main

import (
	"stubborn"
)

func handleMessages(channel stubborn.StubChannel) {
	for {
		pack 	:= channel.SReceive()
		message := bytesToMessage(pack)

		message.printMessage()

		if len(voters) <= nParticipants/2 {
			checkRound(message)
			existsNewVoters := (round == message.Round) && containsNewVoters(message.Voters)
			isMajority 		:= (phase == 1) && (len(message.Voters) > nParticipants/2)
			
			if existsNewVoters || isMajority {
				message.Voters[peerID] = true
				voters = union(voters, message.Voters)
				
				if message.Estimate.peerID == ((round % nParticipants) + 1){
					estimate = message.Estimate
				}

				message := newMessage(peerID, round, phase, voters, estimate)
				data 	:= message.messageToBytes()
				channel.SSendAll(data)
			}
		} else if checkPhase(message) { 
			break 
		}
	}
}

func checkRound(message *Message) {
	if round < message.Round {
		estimate = message.Estimate
		round 	 = message.Round
		phase    = message.Phase
		voters   = make(map[int] bool)
	} else if round == message.Round && phase < message.Phase {
		phase  = message.Phase
		voters = make(map[int] bool)
	}
}

func checkPhase(message *Message) bool {
	if phase == 1 {
		consensusDecision = message.Estimate.Value
		return true
	} else {
		round++
		phase  = 1
		voters = make(map[int] bool)
		return false
	}
}


// Auxiliary Funtions 

func containsNewVoters(senderVoters map[int] bool) bool {
	for id, _ := range senderVoters {
		_, isPresent := voters[id]
		if !isPresent {
			return true
		}  
	}

	return false
}

func union(hash1 map[int] bool, hash2 map[int] bool) map[int] bool{
	for id, _ := range hash2 {
		hash1[id] = true
	}

	return hash1
}