package consensusinfo

import "time"

// Results is the struct that contains the consensus results
type Results struct {
	PeerID       int
	Sent         []float64
    Received     []float64
    DecisionTime time.Time
}

// DurationSlice define a new type of Slice to get order a list of time values
type DurationSlice []time.Duration

// NewResults creates a new estimate
func NewResults(sent, received []float64, decisionTime time.Time, id int) (results *Results) {
	results = new(Results)
	results.PeerID = id
	results.Sent = sent
    results.Received = received
    results.DecisionTime = decisionTime

	return
}

func (p DurationSlice) Len() int {
    return len(p)
}

func (p DurationSlice) Less(i, j int) bool {
    return p[i].Nanoseconds() < p[j].Nanoseconds()
}

func (p DurationSlice) Swap(i, j int) {
    p[i], p[j] = p[j], p[i]
}
