package consensusinfo

// ConsensusInfo is the struct that contains the information about consensus
type ConsensusInfo struct {
    CoordID     int
    Voters      map[int] bool
    Round       int
    Phase       int
    Estimate    *Estimate
    CDecision   string
    PercentMiss float64
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
    estimate        = new(Estimate)
    estimate.Value  = value
    estimate.PeerID = id

    return
}