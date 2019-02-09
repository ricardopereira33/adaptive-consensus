package stubborn

import (
	cmap "github.com/orcaman/concurrent-map"
	"strconv"
	"time"
)

// Metrics contains all the metrics of a simulation for 1 peer
type Metrics struct {
	msgReceived cmap.ConcurrentMap
	msgSent     cmap.ConcurrentMap
	start    time.Time
	decision time.Time
}

// NewMetrics creates a new metrics struct
func NewMetrics(nPeers int) (metrics *Metrics) {
	metrics = new(Metrics)
	metrics.msgReceived = newMap(nPeers)
	metrics.msgSent = newMap(nPeers)
	metrics.start = time.Now()

	return
}

func (m *Metrics) finish() {
	m.decision = time.Now()
}

// incReceivedMsg increments the number of received messages
func (m *Metrics) incMsgReceived(peerID int) {
	strID := strconv.Itoa(peerID)
	value, _ := m.msgReceived.Get(strID)
	m.msgReceived.Set(strID, value.(int)+1)
}

// incSendedMsg increments the number of sended messages
func (m *Metrics) incMsgSent(peerID int) {
	strID := strconv.Itoa(peerID)
	value, _ := m.msgSent.Get(strID)
	m.msgSent.Set(strID, value.(int)+1)
}

// getSendedMsg returns the number of received messages
func (m *Metrics) getMsgSent(peerID int) int {
	strID := strconv.Itoa(peerID)
	value, _ := m.msgSent.Get(strID)

	return value.(int)
}

func (m *Metrics) results() ([]float64, []float64) {
	size := m.msgReceived.Count()
	sent := make([]float64, size)
	received := make([]float64, size)

	for _, id := range m.msgReceived.Keys() {
		msgReceived, _ := m.msgReceived.Get(id)
		msgSent, _ := m.msgSent.Get(id)
		id, _ := strconv.Atoi(id)

		sent[id-1] = float64(msgSent.(int))
		received[id-1] = float64(msgReceived.(int))
	}

	return sent, received
}

func newMap(nPeers int) (channels cmap.ConcurrentMap) {
	channels = cmap.New()

	for id := 1; id <= nPeers; id++ {
		channels.Set(strconv.Itoa(id), 0)
	}

	return
}

// "[" + strconv.Itoa(receivedMsg.(int)) +
// ", " + strconv.Itoa(sendedMsg.(int)) +
// ", " + m.start.Format("15:04:05") +
// ", " + m.decision.Format("15:04:05") + "]"
