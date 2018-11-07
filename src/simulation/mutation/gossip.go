package mutation

import (
	"math/rand"
	stb "simulation/stubborn"
)

// Gossip is a mutation 
type Gossip struct {
	channel stb.StubChannel
	permut 	[]int
	c 		int
	turn    int
	fanout  int
}

// NewGossip creates a new gossip mutation
func NewGossip(channel stb.StubChannel) (g *Gossip) {
	g 		  = new(Gossip)
	g.channel = channel
	g.permut  = perm(channel.GetNParticipants())
	g.c       = 1
	g.turn    = 0
	g.fanout  = 5
	return g
}

// Delta0 is the delta0 implementation
func (g Gossip) Delta0(id int, message *stb.Package) bool {
	return g.Delta(id)
}

// Delta is the delta implementation	
func (g Gossip) Delta(id int) bool {
	nParticipants := len(g.permut)
	g.turn++

	if g.turn == nParticipants {
		g.c    = g.c + g.fanout
		g.turn = 0
	}
	l := 0

	for l < g.fanout {
		if g.permut[(l + g.c) % nParticipants] == id {
			return true
		}
		l++  
	}

	return false	
}


// Auxiliary Functions

func perm(nParticipants int) (list []int) {
	list = rand.Perm(nParticipants)

	for i := range list {
		list[i]++
	}

	return 
}
