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
	Results() ([]float64, []float64, time.Time, []float64, bool)
	Init(time.Duration)
	Finish()

    LastPackageBuffered(peerID int) *Package

    GetPackage(id int) *Package
    GetBandwidthExceeded() bool
	SetMaxTries(int)
	SetPercentageMiss(float64)
	SetDelta0(function func(int, *Package) bool)
	SetDelta(function func(int) bool)
    SetSuspectedFunc(func(int, interface{}))
    SetBandwidth(int)
    SetLatency(float64)
}

// NewSChannel is the constructor of a stubborn channel
func NewSChannel(peerID int, numberParticipants int, peer interface{}, peers cmap.ConcurrentMap) (channel SChannel) {
	channel = newChannel(peerID, numberParticipants, peer, peers)

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
