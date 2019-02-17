package main

import (
	con "simulation/consensusInfo"
	msg "simulation/message"
	stb "simulation/stubborn"
)

func handleMessages(channel stb.StubChannel) {
	for {
		pack := channel.SReceive()
		message := msg.PackageToMessage(pack)
		consensusInfo := channel.GetConsensusInfo()
		numberParticipants := channel.GetNumberParticipants()
		peerID := channel.GetPeerID()

		if len(consensusInfo.Voters) <= numberParticipants/2 {
			checkRound(message, consensusInfo)
			existsNewVoters := (consensusInfo.Round == message.Round) && containsNewVoters(message.Voters, consensusInfo)
			isMajority := (consensusInfo.Phase == 1) && (len(message.Voters) > numberParticipants/2)

			if existsNewVoters || isMajority {
				message.Voters[peerID] = true
				consensusInfo.Voters = union(consensusInfo.Voters, message.Voters)

				if message.Estimate.PeerID == ((consensusInfo.Round % numberParticipants) + 1) {
					consensusInfo.Estimate = message.Estimate
				}

				message := msg.NewMessage(peerID, consensusInfo.Round, consensusInfo.Phase, consensusInfo.Voters, consensusInfo.Estimate)
				data := message.MessageToBytes()
				channel.SSendAll(data)
			}
		}
		if len(consensusInfo.Voters) > numberParticipants/2 {
			if checkPhase(message, consensusInfo) || !channel.IsAlive() {
				channel.Finish()
				break
			}
        }
	}
}

func checkRound(message *msg.Message, consensusInfo *con.ConsensusInfo) {
	if consensusInfo.Round < message.Round {
		consensusInfo.Estimate = message.Estimate
		consensusInfo.Round = message.Round
		consensusInfo.Phase = message.Phase
		consensusInfo.Voters = make(map[int]bool)
	}
	if consensusInfo.Round == message.Round && consensusInfo.Phase < message.Phase {
		consensusInfo.Phase = message.Phase
		consensusInfo.Voters = make(map[int]bool)
	}
}

func checkPhase(message *msg.Message, consensusInfo *con.ConsensusInfo) bool {
	if consensusInfo.Phase == 1 {
		consensusInfo.Decision = message.Estimate.Value
		return true
	}

	consensusInfo.Round++
	consensusInfo.Phase = 1
	consensusInfo.Voters = make(map[int]bool)
	return false
}

func containsNewVoters(senderVoters map[int]bool, consensusInfo *con.ConsensusInfo) bool {
	for id := range senderVoters {
		_, present := consensusInfo.Voters[id]
		if !present {
			return true
		}
	}

	return false
}

func union(hash1 map[int]bool, hash2 map[int]bool) map[int]bool {
	for id := range hash2 {
		hash1[id] = true
	}

	return hash1
}
