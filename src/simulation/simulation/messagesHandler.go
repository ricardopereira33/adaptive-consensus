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
				peer.ReceiveMessage(message.PeerID)

				message.Voters[peerID] = 1
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
		consensusInfo.Voters = make(map[int]int)
	}
	if consensusInfo.Round == message.Round && consensusInfo.Phase < message.Phase {
		consensusInfo.Phase = message.Phase
		consensusInfo.Voters = make(map[int]int)
	}
}

func checkPhase(message *con.Message, consensusInfo *con.Info) bool {
	if consensusInfo.Phase == 1 {
		consensusInfo.Decision = message.Estimate.Value
		return true
	}

	consensusInfo.Round++
	consensusInfo.Phase = 1
	consensusInfo.Voters = make(map[int]int)
	return false
}

func containsNewVoters(senderVoters map[int]int, consensusInfo *con.Info) bool {
	for id := range senderVoters {
		_, present := consensusInfo.Voters[id]
		if !present {
			return true
		}
	}

	return false
}

func union(hash1 map[int]int, hash2 map[int]int) map[int]int {
	for id := range hash2 {
		hash1[id] = 1
	}

	return hash1
}
