package consensus

import (
    "time"
    "math/rand"

    cmap "github.com/orcaman/concurrent-map"
    stb "simulation/stubborn"
    fd "simulation/failuredetection"
)

var (
    // DefaultDelta is the default time to relay the messages to the others peers
	DefaultDelta = time.Second * 5
)

// IPeer is an interface for a peer
type IPeer interface {
    Init()
    IsAlive() bool

    GetPeerID() int
    GetNumberParticipants() int
    GetCoordID() int
    GetConsensusDecision() string
    GetConsensusInfo() *Info
    GetChannel() stb.SChannel

    SetCoordinator(coordID int)
    SetDefaultDelta(defaultDelta int)
}

// Peer is a struct that represents a peer
type Peer struct {
    id                 int
    numberParticipants int
    alive              bool
    channel            stb.SChannel
    consensusInfo      *Info
    detectors          *fd.Detectors
}

// NewPeer creates a new Peer
func NewPeer(peerID, numberParticipants int, peers cmap.ConcurrentMap, detectors *fd.Detectors) (peer *Peer) {
    peer = new(Peer)
	peer.id = peerID
	peer.numberParticipants = numberParticipants
	peer.consensusInfo = NewConsensusInfo()
    peer.detectors = detectors
    peer.alive = true
    peer.channel = stb.NewSChannel(peerID, numberParticipants, peer, peers)

    return
}

func (peer *Peer) triggerFailure() {
    numberParticipants := peer.numberParticipants

    for {
        id := rand.Intn(numberParticipants) + 1

        if id == peer.id {
            peer.alive = false
            break
        }

        time.Sleep(500 * time.Millisecond)
    }
}

func (peer *Peer) handleFailures() {
    peerID := peer.id
    channel := peer.channel
    detectors := peer.detectors
    detector := detectors.GetDetector(peerID)
    numberParticipants := peer.numberParticipants

    for {
        detector.Heartbeat()
        id := rand.Intn(numberParticipants) + 1

        if !detectors.Ping(id) {
            channel.SendSuspicion(peerID, id)
        }

        if !peer.alive {
            break
        }

        time.Sleep(100 * time.Millisecond)
    }
}

// Init is the method that starts routines which handle failures
func (peer *Peer) Init(withFaults bool) {
    peer.channel.Init(DefaultDelta)

    if withFaults {
        go peer.handleFailures()
        go peer.triggerFailure()
    }
}

// IsAlive returns the status of the peer
func (peer *Peer) IsAlive() bool {
	return peer.alive
}

// GetPeerID returns the peer ID
func (peer *Peer) GetPeerID() int {
	return peer.id
}

// GetNumberParticipants returns the number of participants
func (peer *Peer) GetNumberParticipants() int {
	return peer.numberParticipants
}

// GetCoordID returns the coordinator ID
func (peer *Peer) GetCoordID() int {
	return peer.consensusInfo.CoordID
}

// GetConsensusDecision returns the consensus decision value
func (peer *Peer) GetConsensusDecision() string {
	return peer.consensusInfo.Decision
}

// GetConsensusInfo returns the consensus information
func (peer *Peer) GetConsensusInfo() *Info {
	return peer.consensusInfo
}

// GetChannel returns the channel
func (peer *Peer) GetChannel() stb.SChannel {
	return peer.channel
}

// SetCoordinator saves the coordainator ID
func (peer *Peer) SetCoordinator(coordID int) {
	peer.consensusInfo.CoordID = coordID
}

// SetDefaultDelta sets the value of DefaultDelta
func (peer *Peer) SetDefaultDelta(defaultDelta int) {
    DefaultDelta = time.Second * time.Duration(defaultDelta)
}
