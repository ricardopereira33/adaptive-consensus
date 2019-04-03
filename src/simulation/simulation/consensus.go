package main

import (
	con "simulation/consensus"
)

func consensus(peer *con.Peer, value string) {
	channel := peer.GetChannel()
	peerID := peer.GetPeerID()
	consensusInfo := peer.GetConsensusInfo()
	consensusInfo.Voters = make(map[int]bool)
	consensusInfo.Round = 1
	consensusInfo.Phase = 1
	consensusInfo.Estimate = con.NewEstimate(value, peer.GetPeerID())
	consensusInfo.CoordID = (consensusInfo.Round % peer.GetNumberParticipants()) + 1

	if peerID == peer.GetCoordID() {
		consensusInfo.Voters[peerID] = true
		consensusInfo.Estimate.PeerID = peerID

		message := con.NewMessage(peerID, consensusInfo.Round, consensusInfo.Phase, consensusInfo.Voters, consensusInfo.Estimate)
		data := message.MessageToBytes()

		peer.SetCoordinator(consensusInfo.CoordID)
		channel.SendAll(data)
	}
}
