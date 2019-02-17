package stubborn

import (
    "strconv"
    "time"

    fd "simulation/failuredetection"
	con "simulation/consensusInfo"
	cmap "github.com/orcaman/concurrent-map"
)

// StubChannel is an interface to abstract the channel
type StubChannel interface {
	SReceive() *Package
	SSend(int, []byte)
	SSendAll([]byte)
    Init()
    IsAlive() bool
    Results() ([]float64, []float64, time.Time)
    Finish()

	// Sets
	SetDelta0(func(int, *Package) bool)
	SetDelta(func(int) bool)
	SetMaxTries(int)
	SetDefaultDelta(int)
	SetCoordinator(int)
    SetPercentageMiss(float64)
    SetSuspectedFunc(func(int, StubChannel))

	// Gets
	GetPeerID() int
	GetCoordID() int
	GetNumberParticipants() int
	GetConsensusDecision() string
	GetPackage(int) *Package
	GetConsensusInfo() *con.ConsensusInfo
}

// NewStubChannel is the constructor of a stubborn channel
func NewStubChannel(peerID, numberParticipants int, peers cmap.ConcurrentMap, detectors *fd.Detectors) (channel StubChannel) {
	channel = newChannel(peerID, numberParticipants, peers, detectors)

    return
}

// Channels returns a map with all input channels
func Channels(numberParticipants int) (channels cmap.ConcurrentMap) {
	channels = cmap.New()

	for id := 1; id <= numberParticipants; id++ {
		channels.Set(strconv.Itoa(id), make(chan *Package))
	}

	return
}
