package mutation

import "stubborn"

type Gossip struct {
	channel stubborn.StubChannel
}

func NewGossip(channel stubborn.StubChannel) (g *Gossip) {
	g 		  = new(Gossip)
	g.channel = channel
	
	return g
}

func (g Gossip) Delta0(id int, message *stubborn.Package) bool {
	return true
}
	
func (g Gossip) Delta(id int) bool {
	return true	
}