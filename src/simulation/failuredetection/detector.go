package failuredetection

import (
	"strconv"
	"time"

	fd "github.com/joa/failuredetector"
	cmap "github.com/orcaman/concurrent-map"
	ex "simulation/exception"
)

// Detectors is a struct that detects faulty nodes
type Detectors struct {
	Peers              cmap.ConcurrentMap
	percentageFaults   float64
	numberFaults       int
	numberParticipants int
}

// NewDetectors creates a new struct Detectors
func NewDetectors(threshold float64, maxSampleSize int, numberParticipants int, percentageFaults float64) (detectors *Detectors) {
	detectors = new(Detectors)
	detectors.Peers = cmap.New()
	detectors.numberParticipants = numberParticipants
	detectors.numberFaults = 0
	detectors.percentageFaults = percentageFaults

	for id := 1; id <= numberParticipants; id++ {
		value, err := fd.New(8.0, 200, 500*time.Millisecond, 0, 500*time.Millisecond, nil)

		ex.CheckError(err)

		detectors.Peers.Set(strconv.Itoa(id), value)
	}

	return
}

// GetDetector returns the detector for a given peer (id)
func (detectors *Detectors) GetDetector(id int) *fd.PhiAccuralFailureDetector {
	detector, present := detectors.Peers.Get(strconv.Itoa(id))

	if present {
		return detector.(*fd.PhiAccuralFailureDetector)
	}

	return nil
}

// IsAvailable checks if a peer is still alive
func (detectors *Detectors) IsAvailable(id int) bool {
	detector := detectors.GetDetector(id)

	if detector != nil {
		return detector.IsAvailable()
	}

	return false
}

// CanPeerDie verifies if a peer can die, taking into account the current number of faults
func (detectors *Detectors) CanPeerDie() bool {
	currentPercentageFaults := (float64(detectors.numberFaults) / float64(detectors.numberParticipants)) * 100

	if currentPercentageFaults <= detectors.percentageFaults {
		return true
	}

	return false
}

// IncrementFaults increments the number of faults
func (detectors *Detectors) IncrementFaults() {
	detectors.numberFaults++
}
