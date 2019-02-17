package main

import (
    "time"
    "math/rand"

    fd "simulation/failuredetection"
    stb "simulation/stubborn"
    msg "simulation/message"
)

func handleFailures(channel stb.StubChannel, detectors *fd.Detectors) {
    peerID := channel.GetPeerID()
    detector := detectors.GetDetector(peerID)
    numberParticipants := channel.GetNumberParticipants()

    for {
        detector.Heartbeat()
        id := rand.Intn(numberParticipants) + 1

        if !detectors.Ping(id) {
            go suspected(channel, id)
        }

        time.Sleep(100 * time.Millisecond)
    }
}

func suspected(channel stb.StubChannel, id int) {
    consensusInfo := channel.GetConsensusInfo()
    numberParticipants := channel.GetNumberParticipants()
    peerID := channel.GetPeerID()

    if id == ((consensusInfo.Round % numberParticipants) + 1) && consensusInfo.Round == 1 {
        consensusInfo.Round = 2
        consensusInfo.Voters = make(map[int]bool)

        message := msg.NewMessage(peerID, consensusInfo.Round, consensusInfo.Phase, consensusInfo.Voters, consensusInfo.Estimate)
        data := message.MessageToBytes()
        channel.SSendAll(data)
    }
}
