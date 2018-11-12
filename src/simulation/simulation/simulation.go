package main

import (
    "fmt"
    "flag"
    "os"
    "log"
    "strconv"
    mut "simulation/mutation"
    ex "simulation/exception"
    stb "simulation/stubborn"
    cmap "github.com/orcaman/concurrent-map"
)

var (
    debug         bool
    mutation      int
    nParticipants int
    defaultDelta  int
    maxTries      int
    percentMiss   float64
)

func propose(value string) {
    channels  := stb.Channels(nParticipants)
    responses := make(chan string)
    
    for id := 1; id <= nParticipants; id++ {
        go runPeer(id, value, responses, channels)
    }

    for id := 1; id <= nParticipants; id++ {
        response := <-responses
        log.Println("END - " + response)
    }
}

func runPeer(peerID int, value string, response chan string, channels cmap.ConcurrentMap) {
    channel := stb.NewStubChannel(peerID, nParticipants, channels)
    configChannel(channel)

    go consensus(channel, value)
    handleMessages(channel)

    response <- "[Peer " + strconv.Itoa(channel.GetPeerID()) + "] " + channel.GetConsensusDecision()
}

func configChannel(channel stb.StubChannel) {
    channel.Init()
    channel.SetMaxTries(maxTries)
    channel.SetDefaultDelta(defaultDelta)
    channel.SetPercentageMiss(percentMiss)

    mut := mut.NewMutation(channel, mutation)
    channel.SetDelta0(mut.Delta0)
    channel.SetDelta(mut.Delta)
}

func argsInfo(nArgs int) {
    if nArgs < 2 { 
        fmt.Println("simulation <MUTATION> <N_NODES> <DEFAULT_DELTA> <MAX_TRIES> <%_MISS>")
        os.Exit(1) 
    }
}

/** Start the peer
  *
  * simulation gossip 1000 3 3 0.01
  *	simulation <MUTATION> <N_NODES> <DEFAULT_DELTA> <MAX_TRIES> <%_MISS> 
***/
func main() {
    debugFlag := flag.Bool("debug", false, "Debug mode")
    flag.Parse()
    
    debug  = *debugFlag
    args  := flag.Args()	
    argsInfo(len(args))

    var err error
    mutation,      err = mut.Find(args[0])	
    nParticipants, err = strconv.Atoi(args[1])
    defaultDelta,  err = strconv.Atoi(args[2])
    maxTries,      err = strconv.Atoi(args[3])
    percentMiss,   err = strconv.ParseFloat(args[4], 64)
    ex.CheckError(err)

    propose("accept")
}
