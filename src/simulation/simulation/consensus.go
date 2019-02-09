package main

import (
	con "simulation/consensusInfo"
	msg "simulation/message"
	stb "simulation/stubborn"
)

func consensus(channel stb.StubChannel, value string) {
	c := channel.GetConsensusInfo()
	c.Voters = make(map[int]bool)
	c.Round = 1
	c.Phase = 1
	c.Estimate = con.NewEstimate(value, channel.GetPeerID())
	c.CoordID = (c.Round % channel.GetNParticipants()) + 1
	peerID := channel.GetPeerID()

	if peerID == channel.GetCoordID() {
		c.Voters[peerID] = true
		c.Estimate.PeerID = peerID

		message := msg.NewMessage(peerID, c.Round, c.Phase, c.Voters, c.Estimate)
		data := message.MessageToBytes()

		channel.SetCoordinator(c.CoordID)
		channel.SSendAll(data)
	}
}
