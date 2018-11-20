package stubborn

import "time"

// Metrics contains all the metrics of a simulation for 1 peer
type Metrics struct {
    receivedMsg map[int] int
    sendedMsg   map[int] int
    firstMsg    time.Time
    start       time.Time 
    decision    time.Time 
}

// NewMetrics creates a new metrics struct
func NewMetrics() (metrics *Metrics) {
    metrics             = new(Metrics)
    metrics.receivedMsg = make(map[int] int)
    metrics.sendedMsg   = make(map[int] int)

    return 
}

// TODO integrate metrics !!