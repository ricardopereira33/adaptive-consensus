package main

import (
    stb "simulation/stubborn"
    msg "simulation/message"
)

func suspected(id int, channel stb.StubChannel) {
    peerID := channel.GetPeerID()
    consensusInfo := channel.GetConsensusInfo()
    numberParticipants := channel.GetNumberParticipants()

    if id == ((consensusInfo.Round % numberParticipants) + 1) && consensusInfo.Round == 1 {
        consensusInfo.Round = 2
        consensusInfo.Voters = make(map[int]bool)

        message := msg.NewMessage(peerID, consensusInfo.Round, consensusInfo.Phase, consensusInfo.Voters, consensusInfo.Estimate)
        data := message.MessageToBytes()
        channel.SSendAll(data)
    }
}
