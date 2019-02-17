package main

import (
    "strconv"
    "time"
	"flag"
	"fmt"
	"log"
    "os"

	ex "simulation/exception"
    fd "simulation/failuredetection"
	mut "simulation/mutation"
    stb "simulation/stubborn"
	con "simulation/consensusInfo"
	cmap "github.com/orcaman/concurrent-map"
)

var (
	debug              bool
	mutation           string
	mutationCode       int
	numberParticipants int
	defaultDelta       int
	maxTries           int
	percentMiss        float64
)

func propose(value string) {
	channels := stb.Channels(numberParticipants)
    responses := make(chan *con.Results)
    detectors := fd.NewDetectors(3.3, 10, numberParticipants)
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

    if debug {
        log.Println("All received!")
        log.Println("--------------")
    }

	drawResults(list, startTime, mutation)
}

func runPeer(peerID int, value string, response chan *con.Results, channels cmap.ConcurrentMap, detectors *fd.Detectors) {
	channel := stb.NewStubChannel(peerID, numberParticipants, channels)
	configChannel(channel)

    go consensus(channel, value)
    go handleFailures(channel, detectors)

    handleMessages(channel)

	received, sent, decisionTime := channel.Results()
	response <- con.NewResults(received, sent, decisionTime, channel.GetPeerID())
}

func configChannel(channel stb.StubChannel) {
	channel.Init()
	channel.SetMaxTries(maxTries)
	channel.SetDefaultDelta(defaultDelta)
	channel.SetPercentageMiss(percentMiss)

	mut := mut.NewMutation(channel, mutationCode)
	channel.SetDelta0(mut.Delta0)
	channel.SetDelta(mut.Delta)
}

func argsInfo(nArgs int) {
	if nArgs < 2 {
		fmt.Println("simulation <MUTATION> <N_NODES> <DEFAULT_DELTA> <MAX_TRIES> <%_MISS>")
		os.Exit(1)
	}
}

/** Start the peer
  *
  * simulation gossip 1000 3 3 0.01
  *	simulation <MUTATION> <N_NODES> <DEFAULT_DELTA> <MAX_TRIES> <%_MISS>
***/
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
	defaultDelta, err = strconv.Atoi(args[2])
	maxTries, err = strconv.Atoi(args[3])
	percentMiss, err = strconv.ParseFloat(args[4], 64)
	ex.CheckError(err)

	propose("accept")
}
