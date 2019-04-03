package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	cmap "github.com/orcaman/concurrent-map"
	con "simulation/consensus"
	ex "simulation/exception"
	fd "simulation/failuredetection"
	mut "simulation/mutation"
	stb "simulation/stubborn"
)

var (
	debug              bool
	mutation           string
	mutationCode       int
	numberParticipants int
	maxTries           int
	defaultDelta       float64
	percentageMiss     float64
	percentageFaults   float64
	latency            float64
	withMetrics        bool
	withFaults         bool
)

func propose(value string) {
	channels := stb.Channels(numberParticipants)
	responses := make(chan *con.Results)
	detectors := fd.NewDetectors(3.3, 10, numberParticipants, percentageFaults)
	startTime := time.Now()

	for id := 1; id <= numberParticipants; id++ {
		go runPeer(id, value, responses, channels, detectors)
	}

	if debug {
		log.Println("--------------")
		log.Println("Running peers...")
	}

	list := make(map[int]*con.Results)

	for id := 1; id <= numberParticipants; id++ {
		response := <-responses
		list[response.PeerID] = response
	}

	endTime := time.Now()

	if debug {
		log.Println("All received!")
		log.Println("--------------")
	}

	if withMetrics {
		drawResults(list, startTime, mutation)
	}

	// save(list, mutation)
	saveResult(endTime, startTime)
}

func runPeer(peerID int, value string, response chan *con.Results, channels cmap.ConcurrentMap, detectors *fd.Detectors) {
	peer := con.NewPeer(peerID, numberParticipants, channels, detectors, latency)
	configurePeer(peer)

	go consensus(peer, value)
	handleMessages(peer)

	channel := peer.GetChannel()
	received, sent, decisionTime, delays := channel.Results()

	response <- con.NewResults(sent, received, decisionTime, delays, peerID)
}

func configurePeer(peer *con.Peer) {
	channel := peer.GetChannel()
	mut := mut.NewMutation(peer, mutationCode)

	channel.SetSuspectedFunc(suspected)
	channel.SetMaxTries(maxTries)
	channel.SetPercentageMiss(percentageMiss)
	channel.SetDelta0(mut.Delta0)
	channel.SetDelta(mut.Delta)

	peer.SetDefaultDelta(defaultDelta)
	peer.Init(withFaults)
}

func argsInfo(nArgs int) {
	if nArgs < 7 {
		fmt.Println("simulation <MUTATION> <N_NODES> <DEFAULT_DELTA> <MAX_TRIES> <%_MISS> <WITH_FAULTS> <WITH_ALL_METRICS>")
		os.Exit(1)
	}
}

func main() {
	debugFlag := flag.Bool("debug", false, "Debug mode")
	flag.Parse()

	debug = *debugFlag
	args := flag.Args()
	argsInfo(len(args))

	var err error

	mutation = args[0]
	mutationCode, err = mut.Find(mutation)
	numberParticipants, err = strconv.Atoi(args[1])
	defaultDelta, err = strconv.ParseFloat(args[2], 64)
	maxTries, err = strconv.Atoi(args[3])
	percentageMiss, err = strconv.ParseFloat(args[4], 64)
	withFaults, err = strconv.ParseBool(args[5])
	withMetrics, err = strconv.ParseBool(args[6])
	latency, err = strconv.ParseFloat(args[7], 64)
	percentageFaults, err = strconv.ParseFloat(args[8], 64)
	ex.CheckError(err)

	println(mutation + " - " +
		strconv.FormatFloat(defaultDelta, 'f', 2, 64) + " - " +
		strconv.Itoa(maxTries) + " - " +
		strconv.FormatBool(withFaults) + " - " +
		strconv.FormatFloat(latency, 'f', 2, 64) + " - " +
		strconv.FormatFloat(percentageMiss, 'f', 2, 64) + " - " +
		strconv.FormatFloat(percentageFaults, 'f', 2, 64))

	propose("accept")
}
