package consensus

import "time"

// Results is the struct that contains the consensus results
type Results struct {
	PeerID               int
	Sent                 []float64
	Received             []float64
    Delays               []float64
    ListOfBandwidthUsage []string
    ListOfRetransmission []string
    DecisionTime         time.Time
    BandwidthExceeded    bool
}

// DurationSlice define a new type of Slice to get order a list of time values
type DurationSlice []time.Duration

// NewResults creates a new estimate
func NewResults(sent, received []float64, decisionTime time.Time, listOfBandwidthUsage []string, listOfRetransmission []string, bandwidthExceeded bool, id int) (results *Results) {
	results = new(Results)
	results.PeerID = id
	results.Sent = sent
	results.Received = received
	results.DecisionTime = decisionTime
    results.ListOfBandwidthUsage = listOfBandwidthUsage
    results.ListOfRetransmission = listOfRetransmission
    results.BandwidthExceeded = bandwidthExceeded

	return
}

// NewResultsOfDelays creates a new estimate
func NewResultsOfDelays(delays []float64, peerID int) (results *Results) {
	results = new(Results)
	results.PeerID = peerID
	results.Delays = delays

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
