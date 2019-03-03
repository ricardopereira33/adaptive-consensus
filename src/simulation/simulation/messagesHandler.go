package main

import (
	con "simulation/consensus"
)

func handleMessages(peer *con.Peer) {
	for {
        channel := peer.GetChannel()
		pack := channel.Receive()
		message := con.PackageToMessage(pack)
		consensusInfo := peer.GetConsensusInfo()
		numberParticipants := peer.GetNumberParticipants()
		peerID := peer.GetPeerID()

		if len(consensusInfo.Voters) <= numberParticipants/2 {
			checkRound(message, consensusInfo)
			withNewVoters := (consensusInfo.Round == message.Round) && containsNewVoters(message.Voters, consensusInfo)
			isMajority := (consensusInfo.Phase == 1) && (len(message.Voters) > numberParticipants/2)

			if withNewVoters || isMajority {
				message.Voters[peerID] = true
				consensusInfo.Voters = union(consensusInfo.Voters, message.Voters)

				if message.Estimate.PeerID == ((consensusInfo.Round % numberParticipants) + 1) {
					consensusInfo.Estimate = message.Estimate
				}

				message := con.NewMessage(peerID, consensusInfo.Round, consensusInfo.Phase, consensusInfo.Voters, consensusInfo.Estimate)
				data := message.MessageToBytes()
				channel.SendAll(data)
			}
        }

		if len(consensusInfo.Voters) > numberParticipants/2 {
			if checkPhase(message, consensusInfo) {
				channel.Finish()
				break
			}
        }

        if !peer.IsAlive() {
            channel.Finish()
			break
        }
	}
}

func checkRound(message *con.Message, consensusInfo *con.Info) {
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

func checkPhase(message *con.Message, consensusInfo *con.Info) bool {
	if consensusInfo.Phase == 1 {
		consensusInfo.Decision = message.Estimate.Value
		return true
	}

	consensusInfo.Round++
	consensusInfo.Phase = 1
	consensusInfo.Voters = make(map[int]bool)
	return false
}

func containsNewVoters(senderVoters map[int]bool, consensusInfo *con.Info) bool {
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
