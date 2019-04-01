package failuredetection

import (
    "time"
    "strconv"

    fd "github.com/joa/failuredetector"
    cmap "github.com/orcaman/concurrent-map"
    ex "simulation/exception"
)

// Detectors is a struct that detects faulty nodes
type Detectors struct {
    Peers cmap.ConcurrentMap
}

// NewDetectors creates a new struct Detectors
func NewDetectors(threshold float64, maxSampleSize int, numberParticipants int) (detectors *Detectors) {
    detectors = new(Detectors)
    detectors.Peers = cmap.New()

    for id := 1; id <= numberParticipants; id++ {
        value, err :=  fd.New(8.0, 200, 500 * time.Millisecond, 0, 500 * time.Millisecond, nil)

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

// Ping checks if a peer is still alive
func (detectors *Detectors) Ping(id int) bool {
    detector := detectors.GetDetector(id)

    if detector != nil {
        return detector.IsAvailable()
    }

    return false
}


