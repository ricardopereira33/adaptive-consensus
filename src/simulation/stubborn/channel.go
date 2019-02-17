package stubborn

import (
	"fmt"
	"log"
	"time"
    "strconv"
    "math/rand"

	cmap "github.com/orcaman/concurrent-map"
	con "simulation/consensusInfo"
)

var (
	// MaxTries is the maximum value of tries
	MaxTries = 3
	// DefaultDelta is the default time to relay the messages to the others peers
	DefaultDelta = time.Second * 3
	// Debug flag, when it is true, prints some debug infos
	Debug = true
)

// Channel to send and receive messages between peers
type Channel struct {
	Peers              cmap.ConcurrentMap
	OutputBuffer       Buffer
	InputBuffer        chan *Package
	Delta0Func         func(int, *Package) bool
	DeltaFunc          func(int) bool
	PeerID             int
	Metrics            *Metrics
	consensusInfo      *con.ConsensusInfo
    NumberParticipants int
    alive              bool
}

func newChannel(peerID, numberParticipants int, peers cmap.ConcurrentMap) (channel *Channel) {
	channel = new(Channel)
	channel.Peers = peers
	channel.OutputBuffer = newBuffer(numberParticipants)
	channel.PeerID = peerID
	channel.consensusInfo = con.NewConsensusInfo()
	channel.NumberParticipants = numberParticipants
    channel.Metrics = NewMetrics(numberParticipants)
    channel.alive = true

	value, present := peers.Get(strconv.Itoa(peerID))
	if present {
		channel.InputBuffer = value.(chan *Package)
	}

	return
}

func (channel *Channel) delta0(id int, pack *Package) bool {
	return channel.Delta0Func(id, pack)
}

func (channel *Channel) delta(id int) bool {
	return channel.DeltaFunc(id)
}

func (channel *Channel) retransmission() {
	tries := 0

	for {
		time.Sleep(DefaultDelta)
		for id := 1; id <= channel.NumberParticipants; id++ {
			if channel.delta(id) || tries > MaxTries {
				pack := channel.OutputBuffer.GetElement(id)

				if pack != nil && !pack.Arrived {
					channel.send(id)
				}
			}
			tries++
		}
	}
}

func (channel *Channel) triggerFailure() {
    numberParticipants := channel.NumberParticipants
    peerID := channel.PeerID

    for {
        id := rand.Intn(numberParticipants) + 1

        if id == peerID {
            channel.alive = false
            break
        }

        time.Sleep(500 * time.Millisecond)
    }
}

// Init is the method that start receipt of the message
func (channel *Channel) Init() {
    go channel.retransmission()
    go channel.triggerFailure()
}

// IsAlive returns the status of the peer
func (channel *Channel) IsAlive() bool {
	return channel.alive
}

// Results returns the metrics results
func (channel *Channel) Results() ([]float64, []float64, time.Time) {
	return channel.Metrics.results()
}

// Finish is the method that finish the consensus protocol
func (channel *Channel) Finish() {
	channel.Metrics.finish()
}

// GetPeerID returns the peer ID
func (channel *Channel) GetPeerID() int {
	return channel.PeerID
}

// GetPackage returns the last package sent to id
func (channel *Channel) GetPackage(id int) *Package {
	pack := channel.OutputBuffer.GetElement(id)

	return pack
}

// GetNumberParticipants returns the number of participants
func (channel *Channel) GetNumberParticipants() int {
	return channel.NumberParticipants
}

// GetCoordID returns the coordinator ID
func (channel *Channel) GetCoordID() int {
	return channel.consensusInfo.CoordID
}

// GetConsensusDecision returns the consensus decision value
func (channel *Channel) GetConsensusDecision() string {
	return channel.consensusInfo.Decision
}

// GetConsensusInfo returns the consensus information
func (channel *Channel) GetConsensusInfo() *con.ConsensusInfo {
	return channel.consensusInfo
}

// SetMaxTries sets the MaxTries value.
func (channel *Channel) SetMaxTries(max int) {
	MaxTries = max
}

// SetDefaultDelta sets the DefaultDelta value.
func (channel *Channel) SetDefaultDelta(ddelta int) {
	DefaultDelta = time.Second * time.Duration(ddelta)
}

// SetDelta0 is the method to define the delta0 implemention
func (channel *Channel) SetDelta0(f func(int, *Package) bool) {
	channel.Delta0Func = f
}

// SetDelta is the method to define the delta0 implemention
func (channel *Channel) SetDelta(f func(int) bool) {
	channel.DeltaFunc = f
}

// SetCoordinator saves the coordainator ID
func (channel *Channel) SetCoordinator(coordID int) {
	channel.consensusInfo.CoordID = coordID
}

// SetPercentageMiss sets percentMiss value
func (channel *Channel) SetPercentageMiss(miss float64) {
	channel.consensusInfo.PercentMiss = miss
}

// Print prints a message
func (channel Channel) print(message interface{}) {
	fmt.Print("[Peer " + strconv.Itoa(channel.GetPeerID()) + "] ")
	log.Println(message)
}

func (channel Channel) printStatus() {
	log.Println(channel.Peers)
	log.Println(channel.OutputBuffer)
	log.Println(channel.Delta0Func != nil)
	log.Println(channel.DeltaFunc != nil)
}
