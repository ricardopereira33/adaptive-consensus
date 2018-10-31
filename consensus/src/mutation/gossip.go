package mutation

import "stubborn"

type Gossip struct { }

func NewGossip() *Gossip {
	return new(Gossip)
}

func (g Gossip) Delta0(id int, message *stubborn.Package) bool {
	return true
}
	
func (g Gossip) Delta(id int) bool {
	return true	
}