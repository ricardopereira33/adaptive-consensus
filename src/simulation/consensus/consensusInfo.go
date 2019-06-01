package consensus

import (
	"strconv"

	ex "simulation/exception"
	cmap "github.com/orcaman/concurrent-map"
)

// Info is the struct that contains the information about consensus
type Info struct {
	CoordID  int
	Voters   cmap.ConcurrentMap
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

func convertCmapToMap(concurrentMap cmap.ConcurrentMap) map[int] int {
	normalMap := make(map[int]int)

	for value := range concurrentMap.Iter() {
		id, err := strconv.Atoi(value.Key)
		ex.CheckError(err)

		normalMap[id] = value.Val.(int)
	}

	return normalMap
}