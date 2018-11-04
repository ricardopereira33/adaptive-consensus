package mutation

import (
	"log"
	"stubborn"
	"math/rand"
)

type Gossip struct {
	channel stubborn.StubChannel
	permut 	[]int
	c 		int
	turn    int
	fanout  int
}

func NewGossip(channel stubborn.StubChannel) (g *Gossip) {
	g 		  = new(Gossip)
	g.channel = channel
	g.permut  = perm(channel.GetNParticipants())
	log.Println(g.permut)
	g.c       = 1
	g.turn    = 0
	g.fanout  = 5
	return g
}

func (g Gossip) Delta0(id int, message *stubborn.Package) bool {
	return g.Delta(id)
}
	
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

	for i, _ := range list {
		list[i]++
	}

	return 
}
