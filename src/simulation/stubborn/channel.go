package stubborn

import (
    "time"
    "log"
    "fmt"
    "strconv"
    con "simulation/consensusInfo"
    cmap "github.com/orcaman/concurrent-map"
)	

var ( 
    // MaxTries is the maximum value of tries 
    MaxTries     = 3
    // DefaultDelta is the default time to relay the messages to the others peers 
    DefaultDelta = time.Second * 3
    // Debug flag, when it is true, prints some debug infos	
    Debug        = true	
)

// Channel to send and receive messages between peers
type Channel struct {
    Peers         cmap.ConcurrentMap
    OutBuffer     Buffer
    InBuffer      chan *Package
    Delta0Func    func(int, *Package) bool
    DeltaFunc     func(int) bool
    consInfo      *con.ConsensusInfo
    PeerID        int
    NParticipants int
    Metrics       *Metrics
}

func newChannel(peerID, nParticipants int, peers cmap.ConcurrentMap) (channel *Channel) {
    channel               = new(Channel)
    channel.Peers         = peers
    channel.OutBuffer     = newBuffer(nParticipants)
    channel.PeerID        = peerID
    channel.consInfo      = con.NewConsensusInfo()
    channel.NParticipants = nParticipants
    channel.Metrics       = NewMetrics(nParticipants)
    
    value, prs := peers.Get(strconv.Itoa(peerID))
    if prs {
        channel.InBuffer  = value.(chan *Package)
    }

    return
}

func (c *Channel) delta0(id int, pack *Package) bool {
    return c.Delta0Func(id, pack)
}

func (c *Channel) delta(id int) bool {
    return c.DeltaFunc(id)
}

func (c *Channel) retransmission() {
    tries := 0

    for {
        time.Sleep(DefaultDelta)	
        for id := 1; id <= c.NParticipants; id++ {
            if c.delta(id) || tries > MaxTries {
                pack := c.OutBuffer.GetElem(id)
                
                if pack != nil && !pack.Arrived { 
                    c.send(id) 
                }
            }
            tries++
        }
    }
}

// Init is the method that start receipt of the message 
func (c *Channel) Init() {
    go c.retransmission()
}

// Results returns the metrics results
func (c *Channel) Results() map[string] string {
    return c.Metrics.results()
}

// Finish is the method that finish the consensus protocol
func (c *Channel) Finish() {
    c.Metrics.finish()
}

// GetPeerID returns the peer ID
func (c *Channel) GetPeerID() int {
    return c.PeerID
}

// GetPackage returns the last package sent to id
func (c *Channel) GetPackage(id int) *Package {
    pack := c.OutBuffer.GetElem(id)

    return pack
}

// GetNParticipants returns the number of participants
func (c *Channel) GetNParticipants() int {
    return c.NParticipants
}

// GetCoordID returns the coordinator ID
func (c *Channel) GetCoordID() int {
    return c.consInfo.CoordID
}

// GetConsensusDecision returns the consensus decision value
func (c *Channel) GetConsensusDecision() string {
    return c.consInfo.CDecision
}

// GetConsensusInfo returns the consensus information
func (c *Channel) GetConsensusInfo() *con.ConsensusInfo {
    return c.consInfo
}

// SetMaxTries sets the MaxTries value.
func (c *Channel) SetMaxTries(max int) {
    MaxTries = max
}

// SetDefaultDelta sets the DefaultDelta value.
func (c *Channel) SetDefaultDelta(ddelta int) {
    DefaultDelta = time.Second * time.Duration(ddelta)
}

// SetDelta0 is the method to define the delta0 implemention
func (c *Channel) SetDelta0(f func(int, *Package) bool) {
    c.Delta0Func = f
}

// SetDelta is the method to define the delta0 implemention
func (c *Channel) SetDelta(f func(int) bool) {
    c.DeltaFunc = f	
}

// SetCoordinator saves the coordainator ID
func (c *Channel) SetCoordinator(coordID int) {
    c.consInfo.CoordID = coordID
}

// SetPercentageMiss sets percentMiss value
func (c *Channel) SetPercentageMiss(miss float64) {
    c.consInfo.PercentMiss = miss
}

// Print prints a message
func (c Channel) print(message interface{}) {
    fmt.Print("[Peer " + strconv.Itoa(c.GetPeerID()) + "] ")
    log.Println(message)
}

func (c Channel) printStatus() {
    log.Println(c.Peers) 	
    log.Println(c.OutBuffer) 	
    log.Println(c.Delta0Func != nil) 	
    log.Println(c.DeltaFunc != nil) 	
}