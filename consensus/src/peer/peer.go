package main

import (
	"net"
	"fmt"
	"sync"
	"flag"
	"log"
	"os"
)

// MaxDatagramSize is the maximum size of a datagram packet
const MaxDatagramSize int = 2048 

var (
	listConnections = make(map[string] *net.UDPAddr) // List of users connections connected to me
	bufReader       = make(chan string)              // Channel waitin on the user to type something
	mutex           = new(sync.Mutex)
	
	username string 								 // Username of the client
	port 	 string 								 // Port address
	debug  	 bool         							 // Debug mode
)

func connectPeers(ports []string){
	for _, port := range ports {
		addr, err := net.ResolveUDPAddr("udp", "127.0.0.1:" + port)
		checkError(err, false)

		listConnections[port] = addr
	}
}

func run(port string) {
    if debug { log.Println("server") }

    addr, err := net.ResolveUDPAddr("udp", ":" + port)
    checkError(err, false)

    conn, err := net.ListenUDP("udp", addr)
    checkError(err, false)
    
    defer conn.Close()
    conn.SetReadBuffer(MaxDatagramSize)

    go writer(conn)
    reader(conn)
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

	connectPeers(args[2:])

	//starting server
	run(port)
}


