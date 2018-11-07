package stubborn

import (
    con "simulation/consensusinfo"
    "sync"
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
func NewStubChannel(peerID, nParticipantes int, peers map[int] chan *Package, mutex *sync.Mutex) (channel StubChannel){
	channel = newChannel(peerID, nParticipantes, peers, mutex)
	return 
}

// Channels returns a map with all input channel
func Channels(nPeers int) (list map[int] chan *Package){
	list = make(map[int] chan *Package)

	for id := 1; id <= nPeers; id++ {		
        list[id] = make(chan *Package)
    }
    
	return
}
