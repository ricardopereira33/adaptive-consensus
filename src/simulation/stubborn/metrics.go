package stubborn

import (
    "strconv"
	"time"
	cmap "github.com/orcaman/concurrent-map"
)

// Metrics contains all the metrics of a simulation for 1 peer
type Metrics struct {
	messagesReceived cmap.ConcurrentMap
	messagesSent     cmap.ConcurrentMap
	decision         time.Time
}

// NewMetrics creates a new metrics struct
func NewMetrics(numberParticipants int) (metrics *Metrics) {
	metrics = new(Metrics)
	metrics.messagesReceived = newMap(numberParticipants)
	metrics.messagesSent = newMap(numberParticipants)

	return
}

func (metrics *Metrics) finish() {
	metrics.decision = time.Now()
}

// incrementMessagesReceived increments the number of received messages
func (metrics *Metrics) incrementMessagesReceived(peerID int) {
	strID := strconv.Itoa(peerID)
	value, _ := metrics.messagesReceived.Get(strID)
	metrics.messagesReceived.Set(strID, value.(int)+1)
}

// incrementMessagesSent increments the number of sended messages
func (metrics *Metrics) incrementMessagesSent(peerID int) {
	strID := strconv.Itoa(peerID)
	value, _ := metrics.messagesSent.Get(strID)
	metrics.messagesSent.Set(strID, value.(int)+1)
}

// getMessagesSent returns the number of received messages
func (metrics *Metrics) getMessagesSent(peerID int) int {
	strID := strconv.Itoa(peerID)
	value, _ := metrics.messagesSent.Get(strID)

	return value.(int)
}

func (metrics *Metrics) results() ([]float64, []float64, time.Time) {
	size := metrics.messagesReceived.Count()
	sent := make([]float64, size)
	received := make([]float64, size)

	for _, id := range metrics.messagesReceived.Keys() {
		messageReceived, _ := metrics.messagesReceived.Get(id)
		messageSent, _ := metrics.messagesSent.Get(id)
		id, _ := strconv.Atoi(id)

		sent[id-1] = float64(messageSent.(int))
		received[id-1] = float64(messageReceived.(int))
	}

	return sent, received, metrics.decision
}

func newMap(numberParticipantes int) (channels cmap.ConcurrentMap) {
	channels = cmap.New()

	for id := 1; id <= numberParticipantes; id++ {
		channels.Set(strconv.Itoa(id), 0)
	}

	return
}
