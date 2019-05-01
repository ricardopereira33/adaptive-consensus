package consensus

// Info is the struct that contains the information about consensus
type Info struct {
	CoordID  int
	Voters   map[int]int
	Round    int
	Phase    int
	Estimate *Estimate
	Decision string
}

// Estimate is the estimate value for a consensus
type Estimate struct {
	Value  string
	PeerID int
}

// NewConsensusInfo creates a new consensusInfo
func NewConsensusInfo() (consensusInfo *Info) {
	consensusInfo = new(Info)

	return
}

// NewEstimate creates a new estimate
func NewEstimate(value string, id int) (estimate *Estimate) {
	estimate = new(Estimate)
	estimate.Value = value
	estimate.PeerID = id

	return
}
