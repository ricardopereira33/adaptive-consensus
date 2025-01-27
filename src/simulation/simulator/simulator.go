package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"time"
	"runtime/debug"

	cmap "github.com/orcaman/concurrent-map"
	con "simulation/consensus"
	ex "simulation/exception"
	fd "simulation/failuredetection"
	mut "simulation/mutation"
	stb "simulation/stubborn"
	ml "simulation/machinelearning"
)

var (
	withMetrics        bool
	mutationName       string
	mutationCode       int
	numberParticipants int
	maxTries           int
	bandwidth          int
	defaultDelta       float64
	percentageMiss     float64
	percentageFaults   float64
	latency            float64
	probabilityToFail  float64
	balancer		   *ml.Balancer
)

func propose(value string) {
	channels := stb.Channels(numberParticipants)
	responses := make(chan *con.Results)
	detectors := fd.NewDetectors(3.3, 10, numberParticipants, percentageFaults)
	bandwidthExceeded := false

	if mutationName == "adapted" {
		balancer = ml.NewBalancer(0, 8, numberParticipants)
		println("Finish.")
	}

	startTime := time.Now()

	for id := 1; id <= numberParticipants; id++ {
		go runPeer(id, value, responses, channels, detectors, startTime)
	}

	list := make(map[int]*con.Results)

	for id := 1; id <= numberParticipants; id++ {
		response := <-responses
		list[response.PeerID] = response

		if response.BandwidthExceeded {
			bandwidthExceeded = true
		}
	}

	endTime := time.Now()

	if withMetrics {
		drawResults(list, startTime, mutationName)
	}

	saveResult(endTime, startTime, bandwidthExceeded, list)
}

func runPeer(peerID int, value string, response chan *con.Results, channels cmap.ConcurrentMap, detectors *fd.Detectors, startTime time.Time) {
	peer := con.NewPeer(peerID, numberParticipants, channels, detectors, startTime)
	configurePeer(peer)

	consensus(peer, value)
	handleMessages(peer)

	channel := peer.GetChannel()
	metrics := peer.GetMetrics()
	sent, received, decisionTime, listOfBandwidthUsage, listOfRetransmission, bandwidthExceeded := channel.Results(startTime)

	response <- con.NewResults(sent, received, decisionTime, listOfBandwidthUsage, listOfRetransmission, bandwidthExceeded, metrics, peerID)
}

func configurePeer(peer *con.Peer) {
	channel := peer.GetChannel()
	mutation := mut.NewMutation(peer, mutationCode, balancer)

	channel.SetSuspectedFunc(suspected)
	channel.SetMaxTries(maxTries)
	channel.SetPercentageMiss(percentageMiss)
	channel.SetDelta0(mutation.Delta0)
	channel.SetDelta(mutation.Delta)
	channel.SetBandwidth(bandwidth)
	channel.SetLatency(latency)
	channel.SetSenderVoted(con.SenderVoted)

	if mutationName == "adapted" {
		channel.SetCacheQuerie(mutation.(*mut.Adapted).CacheQuerie)
	}

	peer.SetDefaultDelta(defaultDelta)
	peer.SetProbabilityToFail(probabilityToFail)
	peer.Init()
}

func argsInfo(nArgs int) {
	if nArgs < 8 {
		fmt.Println("simulation <MUTATION> <N_NODES> <DEFAULT_DELTA> <MAX_TRIES> <%_MISS> <LATENCY> <BANDWIDTH> <%_FAULTS> <PROB_TO_FAIL> <WITH_ALL_METRICS>")
		os.Exit(1)
	}
}

func main() {
	flag.Parse()
	args := flag.Args()
	argsInfo(len(args))

	var err error

	mutationName = args[0]
	mutationCode, err = mut.Find(mutationName)
	numberParticipants, err = strconv.Atoi(args[1])
	defaultDelta, err = strconv.ParseFloat(args[2], 64)
	maxTries, err = strconv.Atoi(args[3])
	percentageMiss, err = strconv.ParseFloat(args[4], 64)
	latency, err = strconv.ParseFloat(args[5], 64)
	bandwidth, err = strconv.Atoi(args[6])
	percentageFaults, err = strconv.ParseFloat(args[7], 64)
	probabilityToFail, err = strconv.ParseFloat(args[8], 64)
	withMetrics, err = strconv.ParseBool(args[9])
	ex.CheckError(err)

	println(mutationName                                   + " - " +
		strconv.Itoa(numberParticipants)                   + " - " +
		strconv.FormatFloat(defaultDelta, 'f', 2, 64)      + " - " +
		strconv.Itoa(maxTries)                             + " - " +
		strconv.FormatFloat(percentageMiss, 'f', 2, 64)    + " - " +
		strconv.FormatFloat(latency, 'f', 2, 64)           + " - " +
		strconv.Itoa(bandwidth)                            + " - " +
		strconv.FormatFloat(percentageFaults, 'f', 2, 64)  + " - " +
		strconv.FormatFloat(probabilityToFail, 'f', 2, 64) + " - " +
		strconv.FormatBool(withMetrics))

	debug.SetMaxThreads(100000)

	propose("accept")
}
