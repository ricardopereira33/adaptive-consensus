package main

import (
	"fmt"
	"flag"
	"os"
	"stubborn"
	"log"
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
	estimate 		  *Estimate
)

func run(port string, allPorts []string) {
	channel := stubborn.NewStubChannel(port, allPorts)
	defer channel.Close()

	channel.Init()
	channel.SetDelta0(delta0)
	channel.SetDelta(delta)
	channel.SetMaxTries(3)
	channel.SetDefaultDelta(2)
	
	peerID 	      = channel.GetPeerID()
	value  		  = "consensus"
	nParticipants = len(allPorts)

	go consensus(channel, value)

	handleMessages(channel)

	if debug { 
		log.Println("Consensus decision: " + consensusDecision) 
	}
}

func testConnection(channel stubborn.StubChannel, srcID int, destID int){
	if peerID == srcID {
		estimate := newEstimate("1", 1)
		vot 	 := make(map[int] bool)
		vot[1]   = true
		message  := newMessage(1, 1, 1, vot, estimate)
		data     := message.messageToBytes()

		channel.SSend(destID, data)
	}
}

//start the peer
func main() {
	debugFlag := flag.Bool("debug", false, "Debug mode")
	flag.Parse()
	
	debug = *debugFlag
	args := flag.Args()	
	
	if len(args) < 2 { 
		fmt.Println("Less than arguments 2!")
		os.Exit(1) 
	}

	port := args[0]	
	run(port, args)
}


