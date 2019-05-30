package stubborn

import (
	"strconv"
	"sync"
	"time"
	"fmt"
	// "log"

	cmap "github.com/orcaman/concurrent-map"
	lb "github.com/yangwenmai/ratelimit/leakybucket"
)

// Metrics contains all the metrics of a simulation for 1 peer
type Metrics struct {
	messagesReceived    cmap.ConcurrentMap
	messagesSent        cmap.ConcurrentMap
	delays           	cmap.ConcurrentMap
	bandwidthUsage      []*bandwidthUsage
	retransmissions     []*retransmission
	bandwidthMutex      *sync.Mutex
	retransmissionMutex *sync.Mutex

	startTime time.Time
	decision  time.Time
	flag bool
}

type bandwidthUsage struct {
	peerID           int
	timestamp        time.Time
	numberOfMessages uint
}

type retransmission struct {
	peerID    int
	timestamp time.Time
}

// Delay stuffs
type Delay struct {
    lastRegister time.Time
    value        time.Duration
}

// NewMetrics creates a new metrics struct
func NewMetrics(numberParticipants int, startTime time.Time) (metrics *Metrics) {
	metrics = new(Metrics)
	metrics.messagesReceived = newMap(numberParticipants, 0)
	metrics.messagesSent = newMap(numberParticipants, 0)
	metrics.bandwidthUsage = make([]*bandwidthUsage, 0)
	metrics.retransmissions = make([]*retransmission, 0)
	metrics.bandwidthMutex = new(sync.Mutex)
	metrics.retransmissionMutex = new(sync.Mutex)
	metrics.startTime = startTime
	metrics.delays = newMap(numberParticipants, defaultDelay(-1))
	metrics.flag = false

	return
}

func (metrics *Metrics) finish() {
	metrics.decision = time.Now()
}

// logDelay log a delay for a given peer
func (metrics *Metrics) initialDelay(peerID int) {
    strID := strconv.Itoa(peerID)
    metrics.delays.Set(strID, defaultDelay(0))
}

// logDelay log a delay for a given peer
func (metrics *Metrics) logDelay(peerID int, id int) {
	strID := strconv.Itoa(peerID)
    oldDelay, _ := metrics.delays.Get(strID)
	newDelay := defaultDelay(0)

	if oldDelay.(*Delay).value == time.Duration(-1) {
		newDelay.value = newDelay.lastRegister.Sub(oldDelay.(*Delay).lastRegister)
	} else {
		newDelay.value = oldDelay.(*Delay).value + newDelay.lastRegister.Sub(oldDelay.(*Delay).lastRegister)
	}

	metrics.delays.Set(strID, newDelay)

	// metrics.flag = true
}

func defaultDelay(value int) (delay *Delay) {
    delay = new(Delay)
    delay.lastRegister = time.Now()
    delay.value = time.Duration(value)

    return
}

func (metrics *Metrics) resultsOfDelays(peerID int) (delays []float64) {
	// if metrics.flag {
	// 	log.Printf("[resultOfDelays] PeerID (%d) | Memory address (%v)", peerID, &metrics)
	// }

	size := metrics.messagesSent.Count()
    delays = make([]float64, size)

    for _, id := range metrics.delays.Keys() {
		delay, _ := metrics.delays.Get(id)
		id, _ := strconv.Atoi(id)

		// if metrics.flag {
		// 	fmt.Println(delay)
		// }

		delays[id - 1] = float64(delay.(*Delay).value) / float64(time.Millisecond)
	}

    return
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

func (metrics *Metrics) checkBandwidth(peerID int, leakybucket lb.BucketI) {
	for metrics.decision.IsZero() {
		time.Sleep(500 * time.Millisecond)

		numberOfMessages := leakybucket.Capacity() - leakybucket.Remaining()
		newBandwidthUsage := &bandwidthUsage{peerID: peerID, timestamp: time.Now(), numberOfMessages: numberOfMessages}

		metrics.bandwidthMutex.Lock()

		metrics.bandwidthUsage = append(metrics.bandwidthUsage, newBandwidthUsage)

		metrics.bandwidthMutex.Unlock()
	}
}

func (metrics *Metrics) saveRetransmission(peerID int) {
	newRetransmission := &retransmission{peerID: peerID, timestamp: time.Now()}

	metrics.retransmissionMutex.Lock()
	defer metrics.retransmissionMutex.Unlock()

	metrics.retransmissions = append(metrics.retransmissions, newRetransmission)
}

func (metrics *Metrics) results(startTime time.Time) ([]float64, []float64, time.Time, []string, []string) {
	size := metrics.messagesReceived.Count()
	sent := make([]float64, size)
	received := make([]float64, size)
	listOfBandwidthUsage := make([]string, 0)
	listOfRetransmission := make([]string, 0)

	for _, id := range metrics.messagesReceived.Keys() {
		messageReceived, _ := metrics.messagesReceived.Get(id)
		messageSent, _ := metrics.messagesSent.Get(id)
		id, _ := strconv.Atoi(id)

		sent[id-1] = float64(messageSent.(int))
		received[id-1] = float64(messageReceived.(int))
	}

	for _, row := range metrics.bandwidthUsage {
		duration := float64(row.timestamp.Sub(startTime)) / float64(time.Millisecond)
		stringRow := fmt.Sprintf("%d,%f,%d", row.peerID, duration, row.numberOfMessages)

		listOfBandwidthUsage = append(listOfBandwidthUsage, stringRow)
	}

	for _, row := range metrics.retransmissions {
		duration := float64(row.timestamp.Sub(startTime)) / float64(time.Millisecond)
		stringRow := fmt.Sprintf("%d,%f", row.peerID, duration)

		listOfRetransmission = append(listOfRetransmission, stringRow)
	}

	return sent, received, metrics.decision, listOfBandwidthUsage, listOfRetransmission
}

func newMap(numberParticipants int, value interface{}) (channels cmap.ConcurrentMap) {
	channels = cmap.New()

	for id := 1; id <= numberParticipants; id++ {
		channels.Set(strconv.Itoa(id), value)
	}

	return
}
