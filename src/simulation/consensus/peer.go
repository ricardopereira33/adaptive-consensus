package consensus

import (
	"math/rand"
	"sync/atomic"
	"time"
	"strconv"

	cmap "github.com/orcaman/concurrent-map"
	fd "simulation/failuredetection"
	stb "simulation/stubborn"
)

var (
	// DefaultDelta is the default time to relay the messages to the others peers
	DefaultDelta = time.Second * 5
)

// IPeer is an interface for a peer
type IPeer interface {
	Init()
	ReceiveMessage(int)
	NeedAck(int) bool
	IsAlive() bool
	RecordMetrics(*Message)

	GetPeerID() int
	GetNumberParticipants() int
	GetCoordID() int
	GetConsensusDecision() string
	GetConsensusInfo() *Info
	GetChannel() stb.SChannel
	GetMessageNumber() int
	GetMetrics() *Metrics

	SetCoordinator(int)
	SetDefaultDelta(float64)
	SetProbabilityToFail(float64)
}

// Peer is a struct that represents a peer
type Peer struct {
	id                 int
	numberParticipants int
	messageNumber      uint64
	alive              bool
	probabilityToFail  float64
	channel            stb.SChannel
	consensusInfo      *Info
	detectors          *fd.Detectors
	metrics 		   *Metrics
	lastMessages	   cmap.ConcurrentMap
}

// NewPeer creates a new Peer
func NewPeer(peerID, numberParticipants int, peers cmap.ConcurrentMap, detectors *fd.Detectors, startTime time.Time) (peer *Peer) {
	peer = new(Peer)
	peer.id = peerID
	peer.numberParticipants = numberParticipants
	peer.consensusInfo = NewConsensusInfo()
	peer.detectors = detectors
	peer.alive = true
	peer.channel = stb.NewSChannel(peerID, numberParticipants, peer, peers, startTime)
	peer.messageNumber = 0
	peer.lastMessages = newLastMessagesMap(numberParticipants)
	peer.metrics = newMetrics()

	return
}

// ReceiveMessage update map of received messages
func (peer *Peer) ReceiveMessage(peerID int) {
	peer.lastMessages.Set(strconv.Itoa(peerID), true)
}

// RecordMetrics records algorithm variables
func (peer *Peer) RecordMetrics(message *Message) {
	go peer.metrics.recordSnapshot(peer.consensusInfo, peer.consensusInfo.Voters, peer.channel, peer, peer.numberParticipants, message)
}

// GetMetrics returns collected metrics
func (peer *Peer) GetMetrics() *Metrics {
	return peer.metrics
}

// NeedAck checks if is needed send an ACK
func (peer *Peer) NeedAck(peerID int) bool {
	value, _ := peer.lastMessages.Get(strconv.Itoa(peerID))
	needAck := value.(bool)

	if needAck {
		peer.lastMessages.Set(strconv.Itoa(peerID), false)
	}

	return needAck
}

func (peer *Peer) triggerFailure() {
	for {
		if peer.peerFailed() {
			peer.alive = false
			peer.detectors.IncrementFaults()

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

		if !detectors.IsAvailable(id) {
			channel.SendSuspicion(peerID, id)
		}

		if !peer.alive {
			break
		}

		time.Sleep(100 * time.Millisecond)
	}
}

// Init is the method that starts routines which handle failures
func (peer *Peer) Init() {
	peer.channel.Init(DefaultDelta)

	go peer.handleFailures()
	go peer.triggerFailure()
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

// GetMessageNumber returns the next message number
func (peer *Peer) GetMessageNumber() int {
	atomic.AddUint64(&peer.messageNumber, 1)

	return int(atomic.LoadUint64(&peer.messageNumber))
}

// SetCoordinator saves the coordainator ID
func (peer *Peer) SetCoordinator(coordID int) {
	peer.consensusInfo.CoordID = coordID
}

// SetDefaultDelta sets the value of DefaultDelta
func (peer *Peer) SetDefaultDelta(defaultDelta float64) {
	DefaultDelta = time.Millisecond * time.Duration(int(defaultDelta * 1000))
}

// SetProbabilityToFail sets the value of probabilityToFail
func (peer *Peer) SetProbabilityToFail(probabilityToFail float64) {
	peer.probabilityToFail = probabilityToFail
}

func (peer *Peer) peerFailed() bool {
	randomValue := rand.Float64()
	percentage := peer.probabilityToFail / 100.0
	canPeerDie := peer.detectors.CanPeerDie()

	if randomValue < percentage && canPeerDie {
		return true
	}

	return false
}

func newLastMessagesMap(numberParticipants int) cmap.ConcurrentMap {
	lastMessagesMap := cmap.New()

	for id := 1; id <= numberParticipants; id++ {
		lastMessagesMap.Set(strconv.Itoa(id), false)
	}

	return lastMessagesMap
}
