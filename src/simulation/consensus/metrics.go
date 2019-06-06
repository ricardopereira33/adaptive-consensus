package consensus

import (
	"github.com/orcaman/concurrent-map"
	"sync"
	"time"

	stb "simulation/stubborn"
)

// Metrics contains all the metrics of a simulation for 1 peer
type Metrics struct {
	mutex     sync.Mutex
	snapshots []*Snapshot
}

// Snapshot is a snapshot of mutable consensus variables
type Snapshot struct {
	PeerID          int
	CoordID			int
	Round			int
	Phase 			int
	EstimatePeerID 	int
	EstimateValue 	string
	Decision		string
	Voters 			map[int] int
	Delays			[]float64
	Timestamp 		time.Time
}

func newMetrics() (metrics *Metrics) {
	metrics = new(Metrics)
	metrics.snapshots = make([]*Snapshot, 0)

	return
}

func newSnapshot(consensusInfo *Info, voters cmap.ConcurrentMap, channel stb.SChannel) (snapshot *Snapshot) {
	snapshot = new(Snapshot)
	snapshot.PeerID = consensusInfo.PeerID
	snapshot.CoordID = consensusInfo.CoordID
	snapshot.Round = consensusInfo.Round
	snapshot.Phase = consensusInfo.Phase
	snapshot.EstimatePeerID = consensusInfo.Estimate.PeerID
	snapshot.EstimateValue = consensusInfo.Estimate.Value
	snapshot.Decision = consensusInfo.Decision
	snapshot.Voters = convertCmapToMap(voters)
	snapshot.Delays = channel.GetDelays()
	snapshot.Timestamp = time.Now()

	return
}

// GetFirst returns the first snapshot
func (metrics *Metrics) GetFirst() *Snapshot {
	return metrics.snapshots[0]
}

// GetOther returns a snapshot in a given moment
func (metrics *Metrics) GetOther(index int) *Snapshot {
	if len(metrics.snapshots) > index {
		return metrics.snapshots[index]
	}

	return nil
}

func (metrics *Metrics) recordSnapshot(consensusInfo *Info, voters cmap.ConcurrentMap, channel stb.SChannel) {
	metrics.mutex.Lock()

	snapshot := newSnapshot(consensusInfo, voters, channel)
	metrics.snapshots = append(metrics.snapshots, snapshot)

	metrics.mutex.Unlock()
}

// CopyMap copys a map
func CopyMap(mapStruct map[int] int) map[int] int {
	copyMap := make(map[int] int)

	for id, value := range mapStruct {
		copyMap[id] = value
	}

	return copyMap
}
