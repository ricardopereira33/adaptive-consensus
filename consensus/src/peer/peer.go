package main

import (
	"fmt"
	"flag"
	"log"
	"os"
)

var (
	username string 								 // Username of the client
	port 	 string 								 // Port address
	debug  	 bool         							 // Debug mode
)

func run(port string, neighborsPorts []string) {
    if debug { log.Println("server") }
	
	// add Stubborn channel
	// run channel

    //go writer(conn)
    //reader(conn)
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


