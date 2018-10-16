package main

import (
	"fmt"
	"flag"
	"log"
	"os"
	"strconv"
	"stubborn"
)

// Debug mode
var Debug bool

func run(peerID int, port string, neighborsPorts []string) {
    if Debug { log.Println("server") }

	channel := stubborn.NewStubChannel(peerID, port, neighborsPorts)
	defer channel.Close()

	channel.Init()
	channel.SetDelta0(delta0)
	channel.SetDelta(delta)

	if peerID == 4000 {
		message := newMessage("peer1", "hello")
		data	:= messageToBytes(message)
		channel.SSend(3000, data)
	}

	handleMessages(channel)
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

	port   		:= args[0]	
	peerID, err := strconv.Atoi(args[0])
	checkError(err, false)

	//starting server
	run(peerID, port, args[1:])
}


