package stubborn

import (
    "time"
    "strconv"
    cmap "github.com/orcaman/concurrent-map"
)

// Metrics contains all the metrics of a simulation for 1 peer
type Metrics struct {
    receivedMsg cmap.ConcurrentMap
    sendedMsg   cmap.ConcurrentMap
    firstMsg    time.Time
    start       time.Time 
    decision    time.Time 
}

// NewMetrics creates a new metrics struct
func NewMetrics(nPeers int) (metrics *Metrics) {
    metrics             = new(Metrics)
    metrics.receivedMsg = newMap(nPeers)
    metrics.sendedMsg   = newMap(nPeers)
    metrics.start       = time.Now()

    return 
}

// IncReceivedMsg increments the number of received messages
func (m *Metrics) IncReceivedMsg(peerID int) {
    strID    := strconv.Itoa(peerID)
    value, _ := m.receivedMsg.Get(strID)
    m.receivedMsg.Set(strID, value.(int) + 1)
}

// IncSendedMsg increments the number of sended messages
func (m *Metrics) IncSendedMsg(peerID int) {
    strID    := strconv.Itoa(peerID)
    value, _ := m.sendedMsg.Get(strID)
    m.sendedMsg.Set(strID, value.(int) + 1)
}

// GetSendedMsg returns the number of received messages 
func (m *Metrics) GetSendedMsg(peerID int) int {
    strID    := strconv.Itoa(peerID)
    value, _ := m.sendedMsg.Get(strID)

    return value.(int)
}


func newMap(nPeers int) (channels cmap.ConcurrentMap){
    channels = cmap.New()

    for id := 1; id <= nPeers; id++ {
        channels.Set(strconv.Itoa(id), 0)
    }
    
    return
}