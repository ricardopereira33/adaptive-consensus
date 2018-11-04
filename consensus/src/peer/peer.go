package main

import (
	"fmt"
	"flag"
	"os"
	stubborn "stubbornSim"
	"log"
	"strconv"
	mut "mutation"
	msg "message"
	ex "exception"
)

var (
	debug 			  bool
	mutation     	  int
	nNodes			  int
	defaultDelta	  int
	maxTries		  int
	percentMiss		  int
	consensusDecision string
)

func propose(value string) {
	nParticipants = nNodes

	for id := 1; id < nParticipants; id++ {
		runPeer(id, value)
	}

	// wait for the answer
}

func runPeer(peerID int, value string) {
	channel := stubborn.NewStubChannelSim(peerID)
	defer channel.Close()
	configChannel(channel)

	go consensus(channel, value)
	handleMessages(channel)
}

func configChannel(channel stubborn.StubChannel) {
	channel.Init()
	channel.SetMaxTries(maxTries)
	channel.SetDefaultDelta(defaultDelta)

	mut := mut.NewMutation(channel, mutation)
	channel.SetDelta0(mut.Delta0)
	channel.SetDelta(mut.Delta)
}

func argsInfo(nArgs int) {
	if nArgs < 2 { 
		fmt.Println("peer <MUTATION> <N_NODES> <DEFAULT_DELTA> <MAX_TRIES> <%_MISS>")
		os.Exit(1) 
	}
}

/** Start the peer
  *
  * peer gossip 1000 3 3 0.01
  *	peer <MUTATION> <N_NODES> <DEFAULT_DELTA> <MAX_TRIES> <%_MISS> 
***/
func main() {
	debugFlag := flag.Bool("debug", false, "Debug mode")
	flag.Parse()
	
	debug  = *debugFlag
	args  := flag.Args()	
	argsInfo(len(args))

	mutation,     err := mut.Find(args[0])	
	nNodes,       err := strconv.Atoi(args[1])
	defaultDelta, err := strconv.Atoi(args[2])
	maxTries,     err := strconv.Atoi(args[3])
	percentMiss,  err := strconv.Atoi(args[4])
	ex.CheckError(err)

	propose("consensus")
}
