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

func (ds DurationSlice) Len() int {
    return len(ds)
}

func (ds DurationSlice) Less(i, j int) bool {
    return ds[i].Nanoseconds() < ds[j].Nanoseconds()
}

func (ds DurationSlice) Swap(i, j int) {
    ds[i], ds[j] = ds[j], ds[i]
}
