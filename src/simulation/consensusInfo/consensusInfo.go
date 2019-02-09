package consensusinfo

// ConsensusInfo is the struct that contains the information about consensus
type ConsensusInfo struct {
	CoordID     int
	Voters      map[int]bool
	Round       int
	Phase       int
	Estimate    *Estimate
	Decision    string
	PercentMiss float64
}

// Estimate is the estimate value for a consensus
type Estimate struct {
	Value  string
	PeerID int
}

// NewConsensusInfo creates a new consensusInfo
func NewConsensusInfo() (consensusInfo *ConsensusInfo) {
	consensusInfo = new(ConsensusInfo)

	return 
}

// NewEstimate creates a new estimate
func NewEstimate(value string, id int) (estimate *Estimate) {
	estimate = new(Estimate)
	estimate.Value = value
	estimate.PeerID = id

	return
}
