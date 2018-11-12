package main

import (
    "fmt"
    "flag"
    "os"
    "log"
    "strconv"
    "consensus/stubborn"
    ex "consensus/exception"
    mut "consensus/mutation"
    msg "consensus/message"
)

var (
    debug 			  bool
    consensusDecision string
    voters			  map[int] bool
    round 			  int
    phase 			  int
    peerID			  int
    nParticipants	  int
    estimate 		  *msg.Estimate
    mutation          int
    defaultDelta      int
    maxTries          int
)

func run(value, port string, allPorts []string) {
    channel := stubborn.NewStubChannel(port, allPorts, debug)
    defer channel.Close()
    configChannel(channel)

    peerID        = channel.GetPeerID()
    nParticipants = len(allPorts)

    go consensus(channel, value)

    handleMessages(channel)
    log.Println("Finish: " + consensusDecision)
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
    if nArgs < 4 { 
        fmt.Println("peer <MUTATION> <DEFAULT_DELTA> <MAX_TRIES> <OWN_PORT> <NEIGHBOR_PORT> ...")
        os.Exit(1) 
    }
}

//start the peer
func main() {
    debugFlag := flag.Bool("debug", false, "Debug mode")
    flag.Parse()
    
    debug = *debugFlag
    args := flag.Args()	
    argsInfo(len(args))

    var err error
    mutation,      err = mut.Find(args[0])	
    defaultDelta,  err = strconv.Atoi(args[1])
    maxTries,      err = strconv.Atoi(args[2])
    port              := args[3]	
    ex.CheckError(err)

    run("consensus", port, args[4:])
}