package main

import (
	"stubborn"
	msg "message"
)

func handleMessages(channel stubborn.StubChannel) {
	for {
		pack 	:= channel.SReceive()
		message := msg.PackageToMessage(pack)

		message.PrintMessage()

		if len(voters) <= nParticipants/2 {
			checkRound(message)
			existsNewVoters := (round == message.Round) && containsNewVoters(message.Voters)
			isMajority 		:= (phase == 1) && (len(message.Voters) > nParticipants/2)
			
			if existsNewVoters || isMajority {
				message.Voters[peerID] = true
				voters = union(voters, message.Voters)
				
				if message.Estimate.PeerID == ((round % nParticipants) + 1){
					estimate = message.Estimate
				}

				message := msg.NewMessage(peerID, round, phase, nParticipants, voters, estimate)
				data 	:= message.MessageToBytes()
				channel.SSendAll(data)
			}
		} else if checkPhase(message) { 
			break 
		}
	}
}

func checkRound(message *msg.Message) {
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

func checkPhase(message *msg.Message) bool {
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