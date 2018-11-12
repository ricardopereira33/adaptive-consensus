package consensusinfo

import "sync/atomic"

// ConsensusInfo is the struct that contains the information about consensus
type ConsensusInfo struct {
    CoordID		int
    Voters	    map[int] bool
    Round 	    int
    Phase       int
    Estimate    *Estimate
    CDecision   string
    PercentMiss float64
    NMessages   uint64
}

// Estimate is the estimate value for a consensus
type Estimate struct {
    Value  string
    PeerID int
}

// NewConsensusInfo creates a new consensusInfo
func NewConsensusInfo() *ConsensusInfo {
    consensusInfo := new(ConsensusInfo)

    return consensusInfo
}

// NewEstimate creates a new estimate
func NewEstimate(value string, id int) (estimate *Estimate){
    estimate 		= new(Estimate)
    estimate.Value  = value
    estimate.PeerID = id

    return
}

// IncMsg increments the number of messages exchanged
func (ci *ConsensusInfo) IncMsg() {
    atomic.AddUint64(&ci.NMessages, 1)
}

// GetMsg returns the number of messages exchanged
func (ci *ConsensusInfo) GetMsg() (value int) {
    value = int(atomic.LoadUint64(&ci.NMessages))
    
    return 
}