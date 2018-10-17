package main

import (
	"net"
	"fmt"
	"flag"
	"os"
	"stubborn"
)

// Debug mode
var (
	Debug 			  bool
	value 			  string 
	consensusDecision string
	voters			  map[int] *net.UDPAddr
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

	//go consensus(channel, value)

	testConnection(channel, 1, 2)

	handleMessages(channel)
}

func testConnection(channel stubborn.StubChannel, srcID int, destID int){
	if peerID == srcID {
		message := newMessage("peer1", "hello")
		data := messageToBytes(message)

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


