package main

import (
	"strconv"

	con "simulation/consensus"
	cmap "github.com/orcaman/concurrent-map"
)

func consensus(peer *con.Peer, value string) {
	peerID := peer.GetPeerID()
	consensusInfo := peer.GetConsensusInfo()
	consensusInfo.Voters = cmap.New()
	consensusInfo.Round = 1
	consensusInfo.Phase = 1
	consensusInfo.Estimate = con.NewEstimate(value, peer.GetPeerID())
	consensusInfo.CoordID = (consensusInfo.Round % peer.GetNumberParticipants()) + 1
	consensusInfo.PeerID = peerID

	if peerID == peer.GetCoordID() {
		go broadcastMessage(peer, consensusInfo, peerID)
	}
}

func broadcastMessage(peer *con.Peer, consensusInfo *con.Info, peerID int) {
	channel := peer.GetChannel()
	consensusInfo.Voters.Set(strconv.Itoa(peerID), 1)
	consensusInfo.Estimate.PeerID = peerID

	message := con.NewMessage(peerID, consensusInfo.Round, consensusInfo.Phase, consensusInfo.Voters, consensusInfo.Estimate)
	data := message.MessageToBytes()

	peer.SetCoordinator(consensusInfo.CoordID)
	peer.RecordMetrics(message)

	channel.SendAll(data)
}
