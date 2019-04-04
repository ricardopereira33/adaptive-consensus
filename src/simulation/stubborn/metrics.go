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
	delays           cmap.ConcurrentMap
	decision         time.Time
}

// Delay contains all relevant info
type Delay struct {
	lastRegister time.Time
	value        time.Duration
	alreadySent  bool
}

// NewMetrics creates a new metrics struct
func NewMetrics(numberParticipants int) (metrics *Metrics) {
	metrics = new(Metrics)
	metrics.messagesReceived = newMap(numberParticipants, 0)
	metrics.messagesSent = newMap(numberParticipants, 0)
	// metrics.delays = newMap(numberParticipants, defaultDelay())

	return
}

func (metrics *Metrics) finish() {
	metrics.decision = time.Now()
}

// incrementMessagesReceived increments the number of received messages
func (metrics *Metrics) incrementMessagesReceived(peerID int) {
	strID := strconv.Itoa(peerID)
	value, _ := metrics.messagesReceived.Get(strID)
	metrics.messagesReceived.Set(strID, value.(int) + 1)
}

// incrementMessagesSent increments the number of sended messages
func (metrics *Metrics) incrementMessagesSent(peerID int) {
	strID := strconv.Itoa(peerID)
	value, _ := metrics.messagesSent.Get(strID)
	// delayInterface, _ := metrics.delays.Get(strID)
	// delay := delayInterface.(*Delay)

	metrics.messagesSent.Set(strID, value.(int) + 1)

	// delay.alreadySent = true
	// metrics.delays.Set(strID, delay)
}

// getMessagesSent returns the number of received messages
func (metrics *Metrics) getMessagesSent(peerID int) int {
	strID := strconv.Itoa(peerID)
	value, _ := metrics.messagesSent.Get(strID)

	return value.(int)
}

// logDelay log a delay for a given peer
func (metrics *Metrics) initialDelay(peerID int) {
	strID := strconv.Itoa(peerID)
	metrics.delays.Set(strID, defaultDelay())
}

// logDelay log a delay for a given peer
func (metrics *Metrics) logDelay(peerID int) {
	strID := strconv.Itoa(peerID)
	oldDelay, _ := metrics.delays.Get(strID)
	newDelay := defaultDelay()
	newDelay.value = newDelay.lastRegister.Sub(oldDelay.(*Delay).lastRegister)

	metrics.delays.Set(strID, newDelay)
}

func (metrics *Metrics) results() ([]float64, []float64, time.Time, []float64) {
	size := metrics.messagesReceived.Count()
	sent := make([]float64, size)
	received := make([]float64, size)
	delays := make([]float64, size)

	for _, id := range metrics.messagesReceived.Keys() {
		messageReceived, _ := metrics.messagesReceived.Get(id)
		messageSent, _ := metrics.messagesSent.Get(id)
		// delay, _ := metrics.delays.Get(id)
		id, _ := strconv.Atoi(id)

		sent[id-1] = float64(messageSent.(int))
		received[id-1] = float64(messageReceived.(int))

		// if delay.(*Delay).alreadySent {
		//     delays[id-1] = float64(delay.(*Delay).value) / float64(time.Millisecond)
		// } else {
		//     delays[id-1] = 0.0
		// }
	}

	return sent, received, metrics.decision, delays
}

func defaultDelay() (delay *Delay) {
	delay = new(Delay)
	delay.lastRegister = time.Now()
	delay.value = time.Duration(-1)

	return
}

func newMap(numberParticipants int, value interface{}) (channels cmap.ConcurrentMap) {
	channels = cmap.New()

	for id := 1; id <= numberParticipants; id++ {
		channels.Set(strconv.Itoa(id), value)
	}

	return
}
