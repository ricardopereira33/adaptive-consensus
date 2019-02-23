package main

import (
    con "simulation/consensus"
)

func suspected(id int, iPeer interface{}) {
    peer := iPeer.(con.IPeer)
    channel := peer.GetChannel()
    peerID := peer.GetPeerID()
    consensusInfo := peer.GetConsensusInfo()
    numberParticipants := peer.GetNumberParticipants()

    if id == ((consensusInfo.Round % numberParticipants) + 1) && consensusInfo.Round == 1 {
        consensusInfo.Round = 2
        consensusInfo.Voters = make(map[int]bool)

        message := con.NewMessage(peerID, consensusInfo.Round, consensusInfo.Phase, consensusInfo.Voters, consensusInfo.Estimate)
        data := message.MessageToBytes()
        channel.SendAll(data)
    }
}
