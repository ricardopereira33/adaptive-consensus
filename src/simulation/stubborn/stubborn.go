package stubborn

import (
    "strconv"
    "time"
	cmap "github.com/orcaman/concurrent-map"
	con "simulation/consensusInfo"
)

// StubChannel is an interface to abstract the channel
type StubChannel interface {
	SReceive() *Package
	SSend(int, []byte)
	SSendAll([]byte)
	Init()
	Results() ([]float64, []float64, time.Time)
	Finish()

	// Sets
	SetDelta0(f func(int, *Package) bool)
	SetDelta(f func(int) bool)
	SetMaxTries(int)
	SetDefaultDelta(int)
	SetCoordinator(int)
	SetPercentageMiss(float64)

	// Gets
	GetPeerID() int
	GetCoordID() int
	GetNumberParticipants() int
	GetConsensusDecision() string
	GetPackage(int) *Package
	GetConsensusInfo() *con.ConsensusInfo
}

// NewStubChannel is the constructor of a stubborn channel
func NewStubChannel(peerID, numberParticipantes int, peers cmap.ConcurrentMap) (channel StubChannel) {
	channel = newChannel(peerID, numberParticipantes, peers)
    
    return
}

// Channels returns a map with all input channel
func Channels(numberParticipantes int) (channels cmap.ConcurrentMap) {
	channels = cmap.New()

	for id := 1; id <= numberParticipantes; id++ {
		channels.Set(strconv.Itoa(id), make(chan *Package))
	}

	return
}
