package main

import (
	"fmt"
	"flag"
	"os"
	"stubborn"
	"log"
	mut "mutation"
	msg "message"
)

var (
	debug 			  bool
	value 			  string 
	consensusDecision string
	voters			  map[int] bool
	round 			  int
	phase 			  int
	peerID			  int
	nParticipants	  int
	estimate 		  *msg.Estimate
)

func run(port string, allPorts []string) {
	channel := stubborn.NewStubChannel(port, allPorts)
	defer channel.Close()
	configChannel(channel)

	peerID 	      = channel.GetPeerID()
	nParticipants = len(allPorts)
	value = "consensus1"

	go consensus(channel, value)

	handleMessages(channel)
	log.Println("Finish: " + consensusDecision)
}

func configChannel(channel stubborn.StubChannel) {
	channel.Init()
	channel.SetMaxTries(3)
	channel.SetDefaultDelta(3)

	mut := mut.NewMutation(channel, mut.GOSSIP)
	channel.SetDelta0(mut.Delta0)
	channel.SetDelta(mut.Delta)
}

//start the peer
func main() {
	debugFlag := flag.Bool("debug", false, "Debug mode")
	flag.Parse()
	
	debug = *debugFlag
	args  := flag.Args()	
	
	if len(args) < 2 { 
		fmt.Println("Less than arguments 2!")
		os.Exit(1) 
	}

	port := args[0]	
	run(port, args)
}


