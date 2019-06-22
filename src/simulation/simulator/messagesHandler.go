package main

import (
	"strconv"

	con "simulation/consensus"
	cmap "github.com/orcaman/concurrent-map"
)

func handleMessages(peer *con.Peer) {
	for {
		channel := peer.GetChannel()
		pack := channel.Receive()
		message := con.PackageToMessage(pack)
		consensusInfo := peer.GetConsensusInfo()
		numberParticipants := peer.GetNumberParticipants()
		peerID := peer.GetPeerID()

		if consensusInfo.Voters.Count() <= numberParticipants/2 {
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

				newMessage := con.NewMessage(peerID, consensusInfo.Round, consensusInfo.Phase, consensusInfo.Voters, consensusInfo.Estimate)
				data := newMessage.MessageToBytes()

				peer.RecordMetrics(newMessage)
				channel.SendAll(data)
			}
		}

		if consensusInfo.Voters.Count() > numberParticipants/2 {
			if checkPhase(message, consensusInfo) {
				newMessage := con.NewMessage(peerID, consensusInfo.Round, consensusInfo.Phase, consensusInfo.Voters, consensusInfo.Estimate)
				peer.RecordMetrics(newMessage)
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
		consensusInfo.Voters = cmap.New()
	}

	if consensusInfo.Round == message.Round && consensusInfo.Phase < message.Phase {
		consensusInfo.Phase = message.Phase
		consensusInfo.Voters = cmap.New()
	}
}

func checkPhase(message *con.Message, consensusInfo *con.Info) bool {
	if consensusInfo.Phase == 1 {
		consensusInfo.Decision = message.Estimate.Value

		return true
	}

	consensusInfo.Round++
	consensusInfo.Phase = 1
	consensusInfo.Voters = cmap.New()

	return false
}

func containsNewVoters(senderVoters map[int]int, consensusInfo *con.Info) bool {
	for id := range senderVoters {
		present := consensusInfo.Voters.Has(strconv.Itoa(id))

		if !present {
			return true
		}
	}

	return false
}

func union(hash1 cmap.ConcurrentMap, hash2 map[int]int) cmap.ConcurrentMap {
	for id := range hash2 {
		hash1.Set(strconv.Itoa(id), 1)
	}

	return hash1
}
