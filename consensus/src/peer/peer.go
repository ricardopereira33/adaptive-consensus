package main

import (
	"fmt"
	"flag"
	"os"
	"stubborn"
)

// Debug mode
var (
	debug 			  bool
	value 			  string 
	consensusDecision string
	voters			  []int
	round 			  int
	phase 			  int
	peerID			  int
	nParticipants	  int
)

func run(port string, allPorts []string) {
	channel := stubborn.NewStubChannel(port, allPorts)
	defer channel.Close()

	channel.Init()
	channel.SetDelta0(delta0)
	channel.SetDelta(delta)
	
	peerID 	      = channel.GetPeerID()
	value  		  = "consensus"
	nParticipants = len(allPorts)

	go consensus(channel, value)

	handleMessages(channel)
}

func testConnection(channel stubborn.StubChannel, srcID int, destID int){
	if peerID == srcID {
		estimate := newEstimate("1", 1)
		vot 	 := make([]int, 1)
		vot 	 = append(vot, 1)
		message  := newMessage(1, 1, 1, vot, estimate)
		data     := messageToBytes(message)

		channel.SSend(destID, data)
	}
}

//start the peer
func main() {
	debugFlag := flag.Bool("debug", false, "Debug mode")
	flag.Parse()
	
	Debug = *debugFlag
	args := flag.Args()	
	
	if len(args) < 2 { 
		fmt.Println("Less than arguments 2!")
		os.Exit(1) 
	}

	port := args[0]	
	run(port, args)
}


