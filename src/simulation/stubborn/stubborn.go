package stubborn

import (
    con "simulation/consensusinfo"
    cmap "github.com/orcaman/concurrent-map"
    "strconv"
)

// StubChannel is an interface to abstract the channel
type StubChannel interface {
    SReceive() *Package
    SSend(int, []byte)
    SSendAll([]byte)					  
    Init()	

    // Sets
    SetDelta0(f func(int, *Package) bool)	
    SetDelta(f func(int) bool)
    SetMaxTries(int)
    SetDefaultDelta(int)
    SetCoordinator(int)	
    SetPercentageMiss(float64)
    
    // Gets
    GetPeerID()            int
    GetCoordID()           int
    GetNParticipants()     int 	
    GetConsensusDecision() string
    GetPackage(int)        *Package
    GetConsensusInfo()     *con.ConsensusInfo
}

// NewStubChannel is the constructor of a stubborn channel
func NewStubChannel(peerID, nParticipantes int, peers cmap.ConcurrentMap) (channel StubChannel){
    channel = newChannel(peerID, nParticipantes, peers)
    return 
}

// Channels returns a map with all input channel
func Channels(nPeers int) (channels cmap.ConcurrentMap){
    channels = cmap.New()

    for id := 1; id <= nPeers; id++ {
        channels.Set(strconv.Itoa(id), make(chan *Package))
    }
    
    return
}
