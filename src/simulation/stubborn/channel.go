package stubborn

import (
	"strconv"
	"time"

	cmap "github.com/orcaman/concurrent-map"
	rl "go.uber.org/ratelimit"
)

var (
	// MaxTries is the maximum value of tries
	MaxTries = 3
	// Debug flag, when it is true, prints some debug infos
	Debug = true
)

// Channel to send and receive messages between peers
type Channel struct {
	peerID             int
	numberParticipants int
	peer               interface{}
	peers              cmap.ConcurrentMap
	outputBuffer       Buffer
	inputBuffer        chan *Package
	metrics            *Metrics
	percentageMiss     float64
	suspectedFunc      func(int, interface{})
	delta0Func         func(int, *Package) bool
	deltaFunc          func(int) bool
    limiter            rl.Limiter
    latency            time.Duration
}

func newChannel(peerID int, numberParticipants int, peer interface{}, peers cmap.ConcurrentMap) (channel *Channel) {
	channel = new(Channel)
	channel.peerID = peerID
	channel.peers = peers
	channel.peer = peer
	channel.outputBuffer = newBuffer(numberParticipants)
	channel.numberParticipants = numberParticipants
	channel.metrics = NewMetrics(numberParticipants)

	value, present := peers.Get(strconv.Itoa(peerID))

	if present {
		channel.inputBuffer = value.(chan *Package)
	}

	return
}

func (channel *Channel) delta0(id int, pack *Package) bool {
	return channel.delta0Func(id, pack)
}

func (channel *Channel) delta(id int) bool {
	return channel.deltaFunc(id)
}

func (channel *Channel) suspected(id int) {
	channel.suspectedFunc(id, channel.peer)
}

func (channel *Channel) retransmission(defaultDelta time.Duration) {
	tries := 0

	for {
		time.Sleep(defaultDelta)
		for id := 1; id <= channel.numberParticipants; id++ {
			if channel.delta(id) || tries > MaxTries {
				pack := channel.outputBuffer.GetElement(id)

				if pack != nil && !pack.Arrived {
					// channel.metrics.logDelay(id)
					channel.sendMessage(id)
				}
			}
			tries++
		}
	}
}

// Init starts the channel
func (channel *Channel) Init(deltaDefault time.Duration) {
	go channel.retransmission(deltaDefault)
}

// Results returns the metrics results
func (channel *Channel) Results() ([]float64, []float64, time.Time, []float64) {
	return channel.metrics.results()
}

// Finish is the method that finish the consensus protocol
func (channel *Channel) Finish() {
	channel.metrics.finish()
}

// LastPackageBuffered returns the last packages buffered and probably sent
func (channel *Channel) LastPackageBuffered(peerID int) *Package {
    return channel.outputBuffer.GetElement(peerID)
}

// GetPackage returns the last package sent to id
func (channel *Channel) GetPackage(id int) *Package {
	pack := channel.outputBuffer.GetElement(id)

	return pack
}

// SetMaxTries sets the MaxTries value.
func (channel *Channel) SetMaxTries(max int) {
	MaxTries = max
}

// SetDelta0 is the method to define the delta0 implemention
func (channel *Channel) SetDelta0(function func(int, *Package) bool) {
	channel.delta0Func = function
}

// SetDelta is the method to define the delta implemention
func (channel *Channel) SetDelta(function func(int) bool) {
	channel.deltaFunc = function
}

// SetSuspectedFunc is the method to define the suspected implementation
func (channel *Channel) SetSuspectedFunc(function func(int, interface{})) {
	channel.suspectedFunc = function
}

// SetPercentageMiss sets percentageMiss value
func (channel *Channel) SetPercentageMiss(percentage float64) {
	channel.percentageMiss = percentage
}

// SetBandwidth sets bandwidth value
func (channel *Channel) SetBandwidth(bandwidth int) {
	channel.limiter = rl.New(bandwidth)
}

// SetLatency sets latency value
func (channel *Channel) SetLatency(latency float64) {
    duration := time.Duration(int(channel.latency / 2))
    channel.latency = time.Millisecond * duration
}
