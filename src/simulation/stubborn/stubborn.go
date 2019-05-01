package stubborn

import (
	"strconv"
	"time"

	cmap "github.com/orcaman/concurrent-map"
)

// SChannel is an interface to abstract the channel
type SChannel interface {
	Receive() *Package
	Send(int, []byte)
	SendAll([]byte)
	SendSuspicion(int, int)
	Results(time.Time) ([]float64, []float64, time.Time, []string, []string, bool)
	Init(time.Duration)
	Finish()

	LastPackageBuffered(peerID int) *Package

	GetPackage(id int) *Package
	GetBandwidthExceeded() bool

    SetMaxTries(int)
	SetPercentageMiss(float64)
	SetSuspectedFunc(func(int, interface{}))
	SetBandwidth(int)
    SetLatency(float64)

    SetDelta0(function func(int, *Package) bool)
	SetDelta(function func(int) bool)
    SetSenderVoted(func(int, *Package) bool)
}

// NewSChannel is the constructor of a stubborn channel
func NewSChannel(peerID int, numberParticipants int, peer interface{}, peers cmap.ConcurrentMap, startTime time.Time) (channel SChannel) {
	channel = newChannel(peerID, numberParticipants, peer, peers, startTime)

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
