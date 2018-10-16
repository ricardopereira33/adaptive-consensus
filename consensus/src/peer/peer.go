package main

import (
	"fmt"
	"flag"
	"log"
	"os"
	"strconv"
	"stubborn"
)

var (
	username string 								 // Username of the client
	port 	 string 								 // Port address
	debug  	 bool         							 // Debug mode
)

func run(port string, neighborsPorts []string) {
    if debug { log.Println("server") }
	
	id, err := strconv.Atoi(port)
	checkError(err, false)

	channel := stubborn.NewStubChannel(id, port, neighborsPorts)
	channel.Init()
	channel.SetDelta0(delta0)
	channel.SetDelta(delta)

	handleMessages(channel)
}

//start the peer
func main() {
	debugFlag := flag.Bool("debug", false, "Debug mode")
	flag.Parse()
	
	debug = *debugFlag
	args := flag.Args()	
	
	if len(args) < 3 { 
		fmt.Println("Less than arguments 3!")
		os.Exit(1) 
	}

	username = args[0] 							
	port 	 = args[1]			

	//starting server
	run(port, args[2:])
}


