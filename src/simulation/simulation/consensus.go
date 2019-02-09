package main

import (
	con "simulation/consensusInfo"
	msg "simulation/message"
	stb "simulation/stubborn"
)

func consensus(channel stb.StubChannel, value string) {
	peerID := channel.GetPeerID()
	consensusInfo := channel.GetConsensusInfo()
	consensusInfo.Voters = make(map[int]bool)
	consensusInfo.Round = 1
	consensusInfo.Phase = 1
	consensusInfo.Estimate = con.NewEstimate(value, channel.GetPeerID())
	consensusInfo.CoordID = (consensusInfo.Round % channel.GetNumberParticipants()) + 1

	if peerID == channel.GetCoordID() {
		consensusInfo.Voters[peerID] = true
		consensusInfo.Estimate.PeerID = peerID

		message := msg.NewMessage(peerID, consensusInfo.Round, consensusInfo.Phase, consensusInfo.Voters, consensusInfo.Estimate)
		data := message.MessageToBytes()

		channel.SetCoordinator(consensusInfo.CoordID)
		channel.SSendAll(data)
	}
}
